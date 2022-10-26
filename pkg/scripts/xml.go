package scripts

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

var ErrorEndUnexpected = errors.New("End of element")
var ErrorEndExpected = errors.New("End of element not found")
var ErrorWrongElement = errors.New("Wrong element")

type XmlEncodeError struct {
	Expr Expr
	Err  error
}

var _ error = &XmlEncodeError{}

func (e XmlEncodeError) Error() string {
	if e.Expr != nil {
		return fmt.Sprintf("%s: %v", e.Expr.Name(), e.Err)
	} else {
		return e.Err.Error()
	}
}

type XmlDecodeError struct {
	Element xmlElement
	Err     error
}

var _ error = &XmlDecodeError{}

func (e XmlDecodeError) Error() string {
	return fmt.Sprintf("<%s>: %v", e.Element.token.Name.Local, e.Err)
}

type XmlEntry struct {
	Expr   Expr
	Decode func(out *Expr, n xmlElement, x *Xml) error
	Encode func(in Expr, n *xmlElement, x *Xml) error
}

func NewXmlEntry[E Expr](
	decode func(n xmlElement, x *Xml) (E, error),
	encode func(in E, n *xmlElement, x *Xml) error,
) XmlEntry {
	var empty E
	return XmlEntry{
		Expr: empty,
		Decode: func(out *Expr, n xmlElement, x *Xml) error {
			expr, err := decode(n, x)
			if err != nil {
				return XmlDecodeError{n, err}
			}
			*out = expr
			return nil
		},
		Encode: func(in Expr, n *xmlElement, x *Xml) error {
			if casted, ok := in.(E); ok {
				err := encode(casted, n, x)
				if err != nil {
					return XmlEncodeError{in, err}
				}
				return nil
			} else {
				return errors.New("There was an unexpected type error with encoding Exprs to XML.")
			}
		},
	}
}

type Xml struct {
	Entries   map[string]XmlEntry
	Types     map[string]reflect.Type
	NameToTag func(string) string
}

func NewXml() *Xml {
	return &Xml{
		Entries:   make(map[string]XmlEntry),
		Types:     make(map[string]reflect.Type),
		NameToTag: strings.ToLower,
	}
}

func CreateXML(with func(x *Xml)) *Xml {
	x := NewXml()
	with(x)
	return x
}

func (x *Xml) Define(entry XmlEntry) {
	x.Entries[strings.ToLower(entry.Expr.Name())] = entry
}

func (x *Xml) Defines(entries []XmlEntry) {
	for _, entry := range entries {
		x.Define(entry)
	}
}

func (x *Xml) Decode(n xmlElement) (Expr, error) {
	entry, exists := x.Entries[strings.ToLower(n.token.Name.Local)]
	if !exists {
		return nil, nil
	}
	var out Expr
	err := entry.Decode(&out, n, x)
	return out, err
}

func (x *Xml) DecodeList(n xmlNodes) ([]Expr, error) {
	elements := n.getElements()
	exprs := make([]Expr, 0)
	for _, child := range elements {
		parsed, err := x.Decode(child)
		if err != nil {
			return exprs, err
		}
		if parsed == nil {
			break
		}
		exprs = append(exprs, parsed)
	}
	return exprs, nil
}

func (x *Xml) DecodeRangedList(n xmlNodes, min int, max int) ([]Expr, error) {
	list, err := x.DecodeList(n)
	if err != nil {
		return list, err
	}
	if min != -1 && len(list) < min {
		return list, fmt.Errorf("Expecting at least %d, found %d", min, len(list))
	}
	if max != -1 && len(list) > max {
		return list, fmt.Errorf("Expecting at most %d, found %d", max, len(list))
	}
	return list, nil
}

func (x *Xml) Encode(e Expr) (*xmlElement, error) {
	entry, exists := x.Entries[strings.ToLower(e.Name())]
	if !exists {
		return nil, XmlEncodeError{e, fmt.Errorf("No encoding logic exists for %s", e.Name())}
	}
	tag := e.Name()
	if x.NameToTag != nil {
		tag = x.NameToTag(tag)
	}
	el := newXmlElement(tag, xmlNodes{})
	return el, entry.Encode(e, el, x)
}

func (x *Xml) EncodeList(exprs []Expr) (xmlNodes, error) {
	nodes := xmlNodes{}
	for _, expr := range exprs {
		el, err := x.Encode(expr)
		if err != nil {
			return nodes, err
		}
		nodes = append(nodes, el.toAny())
	}
	return nodes, nil
}

func (x *Xml) Unmarshal(d *xml.Decoder) (Expr, error) {
	nodes, err := parseNodes(d)
	if err != nil {
		return nil, err
	}
	els := nodes.getElements()
	if len(els) == 0 {
		return nil, errors.New("No elements found for UnmarshalXML")
	}
	return x.Decode(els[0])
}

func (x *Xml) FromString(input string) (Expr, error) {
	decoder := xml.NewDecoder(strings.NewReader(input))
	return x.Unmarshal(decoder)
}

func (x *Xml) EncodeElement(e *xml.Encoder, expr Expr) (*xmlElement, error) {
	return XmlInstance.Encode(expr)
}

func (x *Xml) Marshal(e *xml.Encoder, expr Expr) error {
	encoded, err := x.Encode(expr)
	if err != nil {
		return XmlEncodeError{expr, err}
	}

	var encode func(el *xmlNode[any]) error
	encode = func(el *xmlNode[any]) error {
		err := e.EncodeToken(el.token)
		if err != nil {
			return err
		}
		if el.children != nil && len(el.children) > 0 {
			for _, child := range el.children {
				err := encode(child)
				if err != nil {
					return err
				}
			}
		}
		if hasEnd, ok := el.token.(xml.StartElement); ok {
			return e.EncodeToken(hasEnd.End())
		}
		return nil
	}

	return encode(encoded.toAny())
}

func (x *Xml) ToString(expr Expr, indent string) (string, error) {
	out := new(bytes.Buffer)
	enc := xml.NewEncoder(out)
	if indent != "" {
		enc.Indent("", indent)
	}
	err := x.Marshal(enc, expr)
	enc.Flush()
	return out.String(), err
}

func UnmarshalXML(d *xml.Decoder) (Expr, error) {
	return XmlInstance.Unmarshal(d)
}

func FromXmlString(x string) (Expr, error) {
	return XmlInstance.FromString(x)
}

func MarshalXML(e *xml.Encoder, expr Expr) error {
	return XmlInstance.Marshal(e, expr)
}

func ToXmlString(expr Expr, indent string) (string, error) {
	return XmlInstance.ToString(expr, indent)
}

var XmlEntries []XmlEntry = []XmlEntry{
	// And
	// <and>[condition]{1,}</and>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (And, error) {
			conditions, err := x.DecodeRangedList(e.children, 1, -1)
			return And{Conditions: conditions}, err
		},
		func(in And, e *xmlElement, x *Xml) error {
			conditions, err := x.EncodeList(in.Conditions)
			if err != nil {
				return err
			}
			e.append(conditions)
			return nil
		},
	),
	// Or
	// <or>[condition]{1,}</or>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Or, error) {
			conditions, err := x.DecodeRangedList(e.children, 1, -1)
			return Or{Conditions: conditions}, err
		},
		func(in Or, e *xmlElement, x *Xml) error {
			conditions, err := x.EncodeList(in.Conditions)
			if err != nil {
				return err
			}
			e.append(conditions)
			return nil
		},
	),
	// Not
	// <not>[condition]</not>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Not, error) {
			els := e.children.getElements()
			if len(els) != 1 {
				return Not{}, errors.New("Not can only have one expression")
			}
			condition, err := x.Decode(els[0])
			return Not{Condition: condition}, err
		},
		func(in Not, e *xmlElement, x *Xml) error {
			condition, err := x.Encode(in.Condition)
			if err != nil {
				return err
			}
			e.addElement(condition)
			return nil
		},
	),
	// Body
	// <body>[line]{0,}</body>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Body, error) {
			lines, err := x.DecodeList(e.children)
			return Body{Lines: lines}, err
		},
		func(in Body, e *xmlElement, x *Xml) error {
			lines, err := x.EncodeList(in.Lines)
			if err != nil {
				return err
			}
			e.append(lines)
			return nil
		},
	),
	// If
	// <if>
	//   <case>[condition][then]</case>{1,}
	//   <else>[else]</else>{0,1}
	// </if>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (If, error) {
			out := If{[]IfCase{}, nil}

			err := e.children.sequential([]nodeSequence{
				{
					name: "case",
					min:  1,
					handle: handleList(x, 2, func(inner []Expr) error {
						out.Cases = append(out.Cases, IfCase{
							Condition: inner[0],
							Then:      inner[1],
						})
						return nil
					}),
				},
				{
					name: "else",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Else = inner[0]
						return nil
					}),
				},
			})

			return out, err
		},
		func(in If, e *xmlElement, x *Xml) error {
			for _, ifCase := range in.Cases {
				condition, err := x.Encode(ifCase.Condition)
				if err != nil {
					return nil
				}
				then, err := x.Encode(ifCase.Then)
				if err != nil {
					return err
				}
				e.addNamedElement("case", xmlNodes{
					condition.toAny(),
					then.toAny(),
				})
			}
			if in.Else != nil {
				els, err := x.Encode(in.Else)
				if err != nil {
					return err
				}
				e.addNamedElement("else", els.toNodes())
			}
			return nil
		},
	),
	// Switch
	// <switch>
	//   <value>[value]</value>
	//   <case>
	//     <expected>[expected]{1,}</expected>
	//     <then>[then]</then>
	//   </case>{1,}
	//   <default>[default]</default>{0,1}
	// </switch>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Switch, error) {
			out := Switch{nil, []SwitchCase{}, nil}

			err := e.children.sequential([]nodeSequence{
				{
					name: "value",
					min:  1,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Value = inner[0]
						return nil
					}),
				},
				{
					name: "case",
					min:  1,
					handle: func(c xmlElement, comments []xmlNode[xml.Comment]) error {
						switchCase := SwitchCase{}

						err := c.children.sequential([]nodeSequence{
							{
								name: "expected",
								min:  1,
								max:  1,
								handle: func(c xmlElement, comments []xmlNode[xml.Comment]) error {
									inner, err := x.DecodeRangedList(c.children, 1, -1)
									if err != nil {
										return err
									}
									switchCase.Expected = inner
									return nil
								},
							},
							{
								name: "then",
								min:  1,
								max:  1,
								handle: handleList(x, 1, func(inner []Expr) error {
									switchCase.Then = inner[0]
									return nil
								}),
							},
						})

						out.Cases = append(out.Cases, switchCase)
						return err
					},
				},
				{
					name: "default",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Default = inner[0]
						return nil
					}),
				},
			})

			return out, err
		},
		func(in Switch, e *xmlElement, x *Xml) error {
			value, err := x.Encode(in.Value)
			if err != nil {
				return err
			}
			e.addNamedElement("value", value.toNodes())
			for _, switchCase := range in.Cases {
				then, err := x.Encode(switchCase.Then)
				if err != nil {
					return err
				}

				expected, err := x.EncodeList(switchCase.Expected)
				if err != nil {
					return err
				}

				e.addNamedElement("case", xmlNodes{
					newXmlElement("expected", expected).toAny(),
					newXmlElement("then", then.toNodes()).toAny(),
				})
			}
			if in.Default != nil {
				elDefault, err := x.Encode(in.Default)
				if err != nil {
					return err
				}
				e.addNamedElement("default", xmlNodes{elDefault.toAny()})
			}
			return nil
		},
	),
	// Loop
	// <loop>
	//   <start>[start]</start>{1,}
	//   <while>[while]</while>{1,}
	//   <then>[then]</then>{1,}
	//   <next>[next]</next>{1,}
	// </loop>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Loop, error) {
			out := Loop{}

			err := e.children.sequential([]nodeSequence{
				{
					name: "start",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Start = inner[0]
						return nil
					}),
				},
				{
					name: "while",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.While = inner[0]
						return nil
					}),
				},
				{
					name: "then",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Then = inner[0]
						return nil
					}),
				},
				{
					name: "next",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Next = inner[0]
						return nil
					}),
				},
			})

			return out, err
		},
		func(in Loop, e *xmlElement, x *Xml) error {
			if in.Start != nil {
				start, err := x.Encode(in.Start)
				if err != nil {
					return err
				}
				e.addNamedElement("start", xmlNodes{start.toAny()})
			}
			if in.While != nil {
				while, err := x.Encode(in.While)
				if err != nil {
					return err
				}
				e.addNamedElement("while", while.toNodes())
			}
			if in.Then != nil {
				then, err := x.Encode(in.Then)
				if err != nil {
					return err
				}
				e.addNamedElement("then", then.toNodes())
			}
			if in.Next != nil {
				next, err := x.Encode(in.Next)
				if err != nil {
					return err
				}
				e.addNamedElement("next", next.toNodes())
			}
			return nil
		},
	),
	// Break
	// <break></break>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Break, error) {
			out := Break{}

			_, err := x.DecodeRangedList(e.children, 0, 0)
			if err != nil {
				return out, err
			}

			return out, nil
		},
		func(in Break, e *xmlElement, x *Xml) error {
			return nil
		},
	),
	// Return
	// <return>[value]{0,1}</return>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Return, error) {
			out := Return{}

			inner, err := x.DecodeRangedList(e.children, 0, 1)
			if err != nil {
				return out, err
			}
			if len(inner) == 1 {
				out.Value = inner[0]
			}

			return out, nil
		},
		func(in Return, e *xmlElement, x *Xml) error {
			if in.Value != nil {
				value, err := x.Encode(in.Value)
				if err != nil {
					return err
				}
				e.addElement(value)
			}
			return nil
		},
	),
	// Throw
	// <throw>[error]{0,1}</throw>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Throw, error) {
			out := Throw{}

			inner, err := x.DecodeRangedList(e.children, 0, 1)
			if err != nil {
				return out, err
			}
			if len(inner) == 1 {
				out.Error = inner[0]
			}

			return out, nil
		},
		func(in Throw, e *xmlElement, x *Xml) error {
			if in.Error != nil {
				er, err := x.Encode(in.Error)
				if err != nil {
					return err
				}
				e.addElement(er)
			}
			return nil
		},
	),
	// Try
	// <try>
	//   <body>[body]</body>
	//   <catch>[catch]</catch>{0,1}
	//   <finally>[finally]</finally>{0,1}
	// </try>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Try, error) {
			out := Try{}

			err := e.children.sequential([]nodeSequence{
				{
					name: "body",
					min:  1,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Body = inner[0]
						return nil
					}),
				},
				{
					name: "catch",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Catch = inner[0]
						return nil
					}),
				},
				{
					name: "finally",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Finally = inner[0]
						return nil
					}),
				},
			})

			return out, err
		},
		func(in Try, e *xmlElement, x *Xml) error {
			body, err := x.Encode(in.Body)
			if err != nil {
				return err
			}
			e.addNamedElement("body", body.toNodes())
			if in.Catch != nil {
				catch, err := x.Encode(in.Catch)
				if err != nil {
					return err
				}
				e.addNamedElement("catch", catch.toNodes())
			}
			if in.Finally != nil {
				finally, err := x.Encode(in.Finally)
				if err != nil {
					return err
				}
				e.addNamedElement("finally", finally.toNodes())
			}
			return nil
		},
	),
	// Assert
	// <assert>
	//   <expect>[body]</expect>
	//   <error>[catch]</error>{0,1}
	// </assert>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Assert, error) {
			out := Assert{}

			err := e.children.sequential([]nodeSequence{
				{
					name: "expect",
					min:  1,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Expect = inner[0]
						return nil
					}),
				},
				{
					name: "error",
					min:  0,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Error = inner[0]
						return nil
					}),
				},
			})

			return out, err
		},
		func(in Assert, e *xmlElement, x *Xml) error {
			expect, err := x.Encode(in.Expect)
			if err != nil {
				return err
			}
			e.addNamedElement("expect", expect.toNodes())
			if in.Error != nil {
				er, err := x.Encode(in.Error)
				if err != nil {
					return err
				}
				e.addNamedElement("error", er.toNodes())
			}
			return nil
		},
	),
	// Constant
	// <constant type="[value type]">
	//   [value]
	// </constant>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Constant, error) {
			out := Constant{}

			typeName := e.attr("type")
			if typeName == "" {
				return out, fmt.Errorf("Constant type not given.")
			}
			typeKind, exists := x.Types[typeName]
			if !exists {
				return out, fmt.Errorf("Type not defined in XML: %s", typeName)
			}

			textValue := ""
			for _, n := range e.children {
				if textNode, ok := n.token.(xml.CharData); ok {
					textValue += string(textNode)
				}
			}

			value, err := fromString(strings.TrimSpace(textValue), typeKind)
			if err != nil {
				return out, err
			}

			out.Value = value.Interface()

			return out, nil
		},
		func(in Constant, e *xmlElement, x *Xml) error {
			valueType := reflect.TypeOf(in.Value)
			for typeName, typeKind := range x.Types {
				if valueType == typeKind {
					e.token.Attr = []xml.Attr{{
						Name:  xml.Name{Local: "type"},
						Value: typeName,
					}}
					break
				}
			}
			if e.token.Attr == nil || len(e.token.Attr) == 0 {
				return fmt.Errorf("Error saving constant, unregistered type: %v", valueType)
			}

			value, err := toString(in.Value)
			if err != nil {
				return err
			}

			e.children = append(e.children, &xmlNode[any]{
				token: xml.CharData(value),
			})
			return nil
		},
	),
	// Compare
	// <compare op="=">
	//   [left]
	//   [right]
	// </compare>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Compare, error) {
			out := Compare{}

			opName := e.attr("op")
			if opName == "" {
				return out, fmt.Errorf("Compare op not given.")
			}
			out.Type = CompareOp(opName)

			inner, err := x.DecodeRangedList(e.children, 2, 2)
			if err != nil {
				return out, err
			}

			out.Left = inner[0]
			out.Right = inner[1]

			return out, nil
		},
		func(in Compare, e *xmlElement, x *Xml) error {
			e.token.Attr = []xml.Attr{{
				Name: xml.Name{
					Local: "op",
				},
				Value: string(in.Type),
			}}
			left, err := x.Encode(in.Left)
			if err != nil {
				return err
			}
			e.addElement(left)
			right, err := x.Encode(in.Right)
			if err != nil {
				return err
			}
			e.addElement(right)
			return nil
		},
	),
	// Binary
	// <binary op="=">
	//   [left]
	//   [right]
	// </binary>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Binary, error) {
			out := Binary{}

			opName := e.attr("op")
			if opName == "" {
				return out, fmt.Errorf("Binary op not given.")
			}
			out.Type = BinaryOp(opName)

			inner, err := x.DecodeRangedList(e.children, 2, 2)
			if err != nil {
				return out, err
			}

			out.Left = inner[0]
			out.Right = inner[1]

			return out, nil
		},
		func(in Binary, e *xmlElement, x *Xml) error {
			e.token.Attr = []xml.Attr{{
				Name: xml.Name{
					Local: "op",
				},
				Value: string(in.Type),
			}}
			left, err := x.Encode(in.Left)
			if err != nil {
				return err
			}
			e.addElement(left)
			right, err := x.Encode(in.Right)
			if err != nil {
				return err
			}
			e.addElement(right)
			return nil
		},
	),
	// Unary
	// <unary op="=">
	//   [value]
	// </unary>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Unary, error) {
			out := Unary{}

			opName := e.attr("op")
			if opName == "" {
				return out, fmt.Errorf("Unary op not given.")
			}
			out.Type = UnaryOp(opName)

			inner, err := x.DecodeRangedList(e.children, 1, 1)
			if err != nil {
				return out, err
			}

			out.Value = inner[0]

			return out, nil
		},
		func(in Unary, e *xmlElement, x *Xml) error {
			e.token.Attr = []xml.Attr{{
				Name: xml.Name{
					Local: "op",
				},
				Value: string(in.Type),
			}}
			encoded, err := x.Encode(in.Value)
			if err != nil {
				return err
			}
			e.addElement(encoded)
			return nil
		},
	),
	// Get
	// <get>
	//   [path]
	// </get>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Get, error) {
			out := Get{[]any{}}

			for _, c := range e.children {
				switch tt := c.token.(type) {
				case xml.CharData:
					nodes := strings.Split(string(tt), ",")
					for _, n := range nodes {
						out.Path = append(out.Path, strings.TrimSpace(n))
					}
				case xml.StartElement:
					decoded, err := x.Decode(xmlElement{
						token:    tt,
						children: c.children,
					})
					if err != nil {
						return out, err
					}
					out.Path = append(out.Path, decoded)
				}
			}

			return out, nil
		},
		func(in Get, e *xmlElement, x *Xml) error {
			for _, n := range in.Path {
				if expr, ok := n.(Expr); ok {
					encoded, err := x.Encode(expr)
					if err != nil {
						return nil
					}
					e.addElement(encoded)
				} else {
					asString, err := toString(n)
					if err != nil {
						return err
					}
					e.addNode(&xmlNode[any]{
						token: xml.CharData(asString),
					})
				}
			}
			return nil
		},
	),
	// Set
	// <set>
	//   <path>[path]</path>
	//   <value>[path]</value>
	// </get>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Set, error) {
			out := Set{[]any{}, nil}

			e.children.sequential([]nodeSequence{
				{
					name: "path",
					min:  1,
					max:  1,
					handle: func(e xmlElement, comments []xmlNode[xml.Comment]) error {
						for _, c := range e.children {
							switch tt := c.token.(type) {
							case xml.CharData:
								nodes := strings.Split(string(tt), ",")
								for _, n := range nodes {
									out.Path = append(out.Path, strings.TrimSpace(n))
								}
							case xml.StartElement:
								decoded, err := x.Decode(xmlElement{
									token:    tt,
									children: c.children,
								})
								if err != nil {
									return err
								}
								out.Path = append(out.Path, decoded)
							}
						}
						return nil
					},
				},
				{
					name: "value",
					min:  1,
					max:  1,
					handle: handleList(x, 1, func(inner []Expr) error {
						out.Value = inner[0]
						return nil
					}),
				},
			})

			return out, nil
		},
		func(in Set, e *xmlElement, x *Xml) error {

			paths := newXmlElement("path", xmlNodes{})
			for _, n := range in.Path {
				if expr, ok := n.(Expr); ok {
					encoded, err := x.Encode(expr)
					if err != nil {
						return nil
					}
					e.addElement(encoded)
				} else {
					asString, err := toString(n)
					if err != nil {
						return err
					}
					paths.addNode(&xmlNode[any]{
						token: xml.CharData(asString),
					})
				}
			}
			e.addElement(paths)

			value, err := x.Encode(in.Value)
			if err != nil {
				return err
			}
			e.addNamedElement("value", value.toNodes())

			return nil
		},
	),
	// Invoke
	// <invoke func="funcName">
	//   <paramX>[value]</paramX>
	// 	 <paramY>[value]</paramY>
	// </invoke>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Invoke, error) {
			out := Invoke{
				Params: make(map[string]Expr),
			}

			functionName := e.attr("func")
			if functionName == "" {
				return out, fmt.Errorf("Invoke function not given.")
			}
			out.Function = functionName

			els := e.children.getElements()
			for _, el := range els {
				decoded, err := x.DecodeRangedList(el.children, 1, 1)
				if err != nil {
					return out, err
				}
				out.Params[el.token.Name.Local] = decoded[0]
			}

			return out, nil
		},
		func(in Invoke, e *xmlElement, x *Xml) error {
			e.token.Attr = []xml.Attr{{
				Name: xml.Name{
					Local: "func",
				},
				Value: in.Function,
			}}
			for key, value := range in.Params {
				valueEncoded, err := x.Encode(value)
				if err != nil {
					return err
				}
				e.addNamedElement(key, valueEncoded.toNodes())
			}
			return nil
		},
	),
	// Define
	// <define>
	//   <vars>
	//    <x>[value]</x>
	//    <y>[value]</y>
	//   </vars>
	//	 [body]
	// </define>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Define, error) {
			out := Define{
				Vars: make([]DefineVar, 0),
			}

			e.children.sequential([]nodeSequence{{
				name: "vars",
				min:  1,
				max:  -1,
				handle: func(e xmlElement, comments []xmlNode[xml.Comment]) error {
					for _, n := range e.children.getElements() {
						encoded, err := x.Decode(n)
						if err != nil {
							return nil
						}
						out.Vars = append(out.Vars, DefineVar{
							Name:  n.token.Name.Local,
							Value: encoded,
						})
					}
					return nil
				},
			}, {
				min: 1,
				max: 1,
				handle: handleList(x, 1, func(inner []Expr) error {
					out.Body = inner[0]
					return nil
				}),
			}})

			return out, nil
		},
		func(in Define, e *xmlElement, x *Xml) error {

			vars := newXmlElement("vars", xmlNodes{})

			for _, v := range in.Vars {
				valueEncoded, err := x.Encode(v.Value)
				if err != nil {
					return err
				}
				vars.addNamedElement(v.Name, valueEncoded.toNodes())
			}

			e.addElement(vars)

			body, err := x.Encode(in.Body)
			if err != nil {
				return err
			}
			e.addElement(body)

			return nil
		},
	),
	// Template
	// <template format="Hello {{.name}}!">
	//   <format><constant type="string">Hello {{.name}}!</constant></format>
	//   <vars>
	//     <name>[value]</name>
	//   </vars>
	// </template>
	NewXmlEntry(
		func(e xmlElement, x *Xml) (Template, error) {
			out := Template{
				Vars: make(map[string]Expr),
			}

			format := e.attr("format")
			if format != "" {
				out.Format = Constant{format}
			}

			e.children.sequential([]nodeSequence{{
				name: "format",
				min:  0,
				max:  1,
				handle: handleList(x, 1, func(inner []Expr) error {
					out.Format = inner[0]
					return nil
				}),
			}, {
				name: "vars",
				min:  0,
				max:  -1,
				handle: func(e xmlElement, comments []xmlNode[xml.Comment]) error {
					for _, n := range e.children.getElements() {
						encoded, err := x.Decode(n)
						if err != nil {
							return nil
						}
						out.Vars[n.token.Name.Local] = encoded
					}
					return nil
				},
			}})

			return out, nil
		},
		func(in Template, e *xmlElement, x *Xml) error {
			if formatConstant, ok := in.Format.(Constant); ok {
				if formatString, ok := formatConstant.Value.(string); ok {
					e.token.Attr = []xml.Attr{{
						Name: xml.Name{
							Local: "format",
						},
						Value: formatString,
					}}
				}
			}
			if e.token.Attr == nil || len(e.token.Attr) == 0 {
				formatEncoded, err := x.Encode(in.Format)
				if err != nil {
					return err
				}
				e.addNamedElement("format", formatEncoded.toNodes())
			}

			vars := newXmlElement("vars", xmlNodes{})

			for k, v := range in.Vars {
				valueEncoded, err := x.Encode(v)
				if err != nil {
					return err
				}
				vars.addNamedElement(k, valueEncoded.toNodes())
			}

			e.addElement(vars)

			return nil
		},
	),
}

var XmlInstance *Xml = CreateXML(func(x *Xml) {
	x.Defines(XmlEntries)

	x.Types["int"] = reflect.TypeOf(int(0))
	x.Types["int8"] = reflect.TypeOf(int8(0))
	x.Types["int16"] = reflect.TypeOf(int16(0))
	x.Types["int32"] = reflect.TypeOf(int32(0))
	x.Types["int64"] = reflect.TypeOf(int64(0))
	x.Types["uint"] = reflect.TypeOf(uint(0))
	x.Types["uint8"] = reflect.TypeOf(uint8(0))
	x.Types["uint16"] = reflect.TypeOf(uint16(0))
	x.Types["uint32"] = reflect.TypeOf(uint32(0))
	x.Types["uint64"] = reflect.TypeOf(uint64(0))
	x.Types["bool"] = reflect.TypeOf(true)
	x.Types["string"] = reflect.TypeOf("")
	x.Types["float32"] = reflect.TypeOf(float32(0))
	x.Types["float64"] = reflect.TypeOf(float64(0))
})

type xmlNode[T any] struct {
	token    T
	children xmlNodes
}

type xmlElement xmlNode[xml.StartElement]

func newXmlElement(name string, children xmlNodes) *xmlElement {
	return &xmlElement{
		token: xml.StartElement{
			Name: xml.Name{
				Local: name,
			},
		},
		children: children,
	}
}

func (n xmlElement) attr(name string) string {
	for _, attr := range n.token.Attr {
		if strings.EqualFold(attr.Name.Local, name) {
			return attr.Value
		}
	}
	return ""
}

func (n *xmlElement) addNode(node *xmlNode[any]) {
	n.children = append(n.children, node)
}

func (n *xmlElement) append(nodes xmlNodes) {
	for i := range nodes {
		n.addNode(nodes[i])
	}
}

func (n *xmlElement) addElement(node *xmlElement) {
	n.children = append(n.children, node.toAny())
}

func (n *xmlElement) addNamedElement(name string, children xmlNodes) {
	n.addElement(newXmlElement(name, children))
}
func (n *xmlElement) toAny() *xmlNode[any] {
	return &xmlNode[any]{
		token:    n.token,
		children: n.children,
	}
}
func (n *xmlElement) toNodes() xmlNodes {
	return xmlNodes{n.toAny()}
}

type xmlNodes []*xmlNode[any]

func parseNodes(tokener xml.TokenReader) (xmlNodes, error) {
	token, err := tokener.Token()
	nodes := xmlNodes{}
	stack := xmlNodes{}
	for err == nil {
		switch tt := token.(type) {
		case xml.ProcInst:
			continue
		case xml.StartElement:
			stackItem := &xmlNode[any]{
				token:    xml.CopyToken(tt),
				children: xmlNodes{},
			}
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				top.children = append(top.children, stackItem)
			} else {
				nodes = append(nodes, stackItem)
			}
			stack = append(stack, stackItem)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		default:
			child := &xmlNode[any]{
				token: xml.CopyToken(tt),
			}
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				top.children = append(top.children, child)
			} else {
				nodes = append(nodes, child)
			}
		}
		token, err = tokener.Token()
		if token == nil {
			break
		}
		if err != nil && err != io.EOF {
			return nodes, err
		}
	}
	return nodes, nil
}

func (n xmlNodes) getElements() []xmlElement {
	els := make([]xmlElement, 0)
	if n != nil {
		for _, e := range n {
			if el, ok := e.token.(xml.StartElement); ok {
				els = append(els, xmlElement{
					token:    el,
					children: e.children,
				})
			}
		}
	}
	return els
}

func (n xmlNodes) split(names []string) (map[string][]xmlElement, []xmlElement) {
	namesMap := make(map[string]struct{})
	for _, n := range names {
		namesMap[strings.ToLower(n)] = struct{}{}
	}
	namedMap := make(map[string][]xmlElement)
	unnamed := make([]xmlElement, 0)
	for _, e := range n.getElements() {
		key := strings.ToLower(e.token.Name.Local)
		if _, named := namesMap[key]; named {
			namedMap[key] = append(namedMap[key], e)
		} else {
			unnamed = append(unnamed, e)
		}
	}
	return namedMap, unnamed
}

type nodeSequenceHandle func(e xmlElement, comments []xmlNode[xml.Comment]) error

type nodeSequence struct {
	min    int
	max    int
	name   string
	handle nodeSequenceHandle
}

func handleList(x *Xml, count int, handle func(inner []Expr) error) nodeSequenceHandle {
	return func(e xmlElement, comments []xmlNode[xml.Comment]) error {
		parsed, err := x.DecodeList(e.children)
		if err != nil {
			return err
		}
		if len(parsed) != count {
			return fmt.Errorf("Expected %d expressions but got %d.", count, len(parsed))
		}
		return handle(parsed)
	}
}

func (n xmlNodes) sequential(sequence []nodeSequence) error {
	comments := make([]xmlNode[xml.Comment], 0)
	current := 0
	currentHandled := 0
	for _, node := range n {
		switch tt := node.token.(type) {
		case xml.Comment:
			comments = append(comments, xmlNode[xml.Comment]{token: tt})
		case xml.StartElement:
			key := strings.ToLower(tt.Name.Local)
			curr := sequence[current]
			for len(sequence) > current+1 {
				next := sequence[current+1]
				if strings.ToLower(next.name) == key {
					if curr.min != -1 && currentHandled < curr.min {
						return fmt.Errorf("Not enough <%s> elements found.", curr.name)
					}
					current++
					currentHandled = 0
					curr = next
				} else {
					break
				}
			}
			if curr.name != "" && key != curr.name {
				return fmt.Errorf("Unexpected element <%s> where <%s> expected.", key, curr.name)
			}
			if curr.handle != nil {
				err := curr.handle(xmlElement{token: tt, children: node.children}, comments)
				if err != nil {
					return err
				}
			}
			comments = make([]xmlNode[xml.Comment], 0)
			currentHandled++
			if curr.max > 0 && currentHandled > curr.max {
				return fmt.Errorf("Too many <%s> elements found.", curr.name)
			}
		}
	}
	return nil
}
