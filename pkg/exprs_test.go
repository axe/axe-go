package axe

import (
	"context"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

type Add struct {
	First  XmlExpr
	Second XmlExpr
}

var _ XmlExpr = &Add{}

func (x *Add) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	x.First, err = ParseExpr(d)
	if err != nil {
		return err
	}
	x.Second, err = ParseExpr(d)
	if err != nil {
		return err
	}
	err = ParseEnd(d)
	return err
}

func (a Add) Run(ctx context.Context) any {
	if firstValue, ok := a.First.Run(ctx).(float64); ok {
		if secondValue, ok := a.Second.Run(ctx).(float64); ok {
			return firstValue + secondValue
		}
	}

	panic("Invalid inputs for add")
}

func (a Add) Clone() XmlExpr {
	return &Add{}
}

type Number struct {
	Value float64 `xml:"value,attr"`
}

var _ XmlExpr = &Number{}

func (a Number) Run(ctx context.Context) any {
	return a.Value
}

func (a Number) Clone() XmlExpr {
	return &Number{}
}

type BinaryOperation struct {
	First    XmlExpr
	Second   XmlExpr
	Operator string
}

var _ XmlExpr = &BinaryOperation{}

func (x *BinaryOperation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "op" {
			x.Operator = a.Value
		}
	}
	var err error
	x.First, err = ParseExpr(d)
	if err != nil {
		return err
	}
	x.Second, err = ParseExpr(d)
	if err != nil {
		return err
	}
	err = ParseEnd(d)
	return err
}

func (a BinaryOperation) Run(ctx context.Context) any {
	if firstValue, ok := a.First.Run(ctx).(float64); ok {
		if secondValue, ok := a.Second.Run(ctx).(float64); ok {
			switch a.Operator {
			case "=":
				return firstValue == secondValue
			case "!=":
				return firstValue != secondValue
			case ">":
				return firstValue > secondValue
			case ">=":
				return firstValue >= secondValue
			case "<=":
				return firstValue <= secondValue
			case "<":
				return firstValue < secondValue
			}
		}
	}
	return nil
}

func (a BinaryOperation) Clone() XmlExpr {
	return &BinaryOperation{}
}

func TestExprs(t *testing.T) {
	XmlExprRegistry["add"] = &Add{}
	XmlExprRegistry["num"] = &Number{}
	XmlExprRegistry["binop"] = &BinaryOperation{}

	x := `
<binop op="=">
	<add>
		<num value="2"></num>
		<num value="3"></num>
	</add>
	<num value="5"></num>
</binop>
	`

	decoder := xml.NewDecoder(strings.NewReader(x))
	expr, err := ParseExpr(decoder)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	result := expr.Run(ctx)

	fmt.Printf("%#v\n", result)
}
