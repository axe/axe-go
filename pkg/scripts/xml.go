package scripts

import (
	"encoding/xml"
	"errors"
	"strings"
)

var ErrorEndUnexpected = errors.New("End of element")
var ErrorEndExpected = errors.New("End of element not found")
var ErrorWrongElement = errors.New("Wrong element")

type XmlEntry struct {
	Expr   Expr
	Decode func(out *Expr, el xml.StartElement, d *xml.Decoder, x *Xml) error
	Encode func(in Expr, e *xml.Encoder, x *Xml) error
}

func NewXmlEntry[E Expr](
	decode func(el xml.StartElement, d *xml.Decoder, x *Xml) (E, error),
	encode func(in E, e *xml.Encoder, x *Xml) error,
) XmlEntry {
	var empty E
	return XmlEntry{
		Expr: empty,
		Decode: func(out *Expr, el xml.StartElement, d *xml.Decoder, x *Xml) error {
			expr, err := decode(el, d, x)
			if err != nil {
				return err
			}
			*out = expr
			return nil
		},
		Encode: func(in Expr, e *xml.Encoder, x *Xml) error {
			if casted, ok := in.(E); ok {
				return encode(casted, e, x)
			} else {
				return errors.New("There was an unexpected type error with encoding Exprs to XML.")
			}
		},
	}
}

type Xml struct {
	Entries map[string]XmlEntry
}

func NewXml() *Xml {
	return &Xml{
		Entries: make(map[string]XmlEntry),
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

func (x *Xml) Decode(d *xml.Decoder) (Expr, error) {
	var out Expr
	for {
		t, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch tt := t.(type) {
		case xml.StartElement:
			entry, exists := x.Entries[strings.ToLower(tt.Name.Local)]
			if !exists {
				break
			}
			err = entry.Decode(&out, tt, d, x)
			if err != nil {
				return out, err
			}
			err = x.decodeEnd(d)
			return out, err
		case xml.EndElement:
			return nil, ErrorEndUnexpected
		}
	}
}

func (x *Xml) decodeEnd(d *xml.Decoder) error {
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch t.(type) {
		case xml.StartElement:
			return ErrorEndExpected
		case xml.EndElement:
			return nil
		}
	}
}

func (x *Xml) decodeNext(d *xml.Decoder, name string, inside func() error) error {
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch tt := t.(type) {
		case xml.StartElement:
			if !strings.EqualFold(tt.Name.Local, name) {
				return ErrorWrongElement
			}
			err = inside()
			if err != nil {
				return err
			}
			err = x.decodeEnd(d)
			return err
		case xml.EndElement:
			return ErrorEndUnexpected
		}
	}
}

func (x *Xml) DecodeList(d *xml.Decoder) ([]Expr, error) {
	parsed, err := x.Decode(d)
	list := make([]Expr, 0)
	for err != nil && parsed != nil {
		list = append(list, parsed)
		parsed, err = x.Decode(d)
	}
	if err != nil && err != ErrorEndUnexpected {
		return list, err
	}
	return list, nil
}

var XmlEntries []XmlEntry = []XmlEntry{
	// And
	NewXmlEntry(
		func(el xml.StartElement, d *xml.Decoder, x *Xml) (And, error) {
			conditions, err := x.DecodeList(d)
			return And{Conditions: conditions}, err
		},
		func(in And, d *xml.Encoder, x *Xml) error {
			return nil
		},
	),
	// Or
	NewXmlEntry(
		func(el xml.StartElement, d *xml.Decoder, x *Xml) (Or, error) {
			conditions, err := x.DecodeList(d)
			return Or{Conditions: conditions}, err
		},
		func(in Or, d *xml.Encoder, x *Xml) error {
			return nil
		},
	),
	// Not
	NewXmlEntry(
		func(el xml.StartElement, d *xml.Decoder, x *Xml) (Not, error) {
			condition, err := x.Decode(d)
			return Not{Condition: condition}, err
		},
		func(in Not, d *xml.Encoder, x *Xml) error {
			return nil
		},
	),
	// Body
	NewXmlEntry(
		func(el xml.StartElement, d *xml.Decoder, x *Xml) (Body, error) {
			lines, err := x.DecodeList(d)
			return Body{Lines: lines}, err
		},
		func(in Body, d *xml.Encoder, x *Xml) error {
			return nil
		},
	),
	// If
	NewXmlEntry(
		func(el xml.StartElement, d *xml.Decoder, x *Xml) (If, error) {
			out := If{[]IfCase{}, nil}

			condition, err := x.Decode(d)
			then, err := x.Decode(d)

			out.Cases = append(out.Cases, IfCase{
				Condition: condition,
				Then:      then,
			})
			return out, err
		},
		func(in If, d *xml.Encoder, x *Xml) error {
			return nil
		},
	),
}

var XmlInstance *Xml = CreateXML(func(x *Xml) {
	x.Defines(XmlEntries)
})

func UnmarshalXML(d *xml.Decoder) (Expr, error) {
	return XmlInstance.Decode(d)
}
