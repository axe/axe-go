package axe

import (
	"context"
	"encoding/xml"
	"errors"
)

var XmlExprRegistry = make(map[string]XmlExpr)

type XmlExpr interface {
	Run(context.Context) any
	Clone() XmlExpr
}

var ErrorEndUnexpected = errors.New("End of element")
var ErrorEndExpected = errors.New("End of element not found")

func ParseExpr(d *xml.Decoder) (XmlExpr, error) {
	for {
		t, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch tt := t.(type) {
		case xml.StartElement:
			registered, exists := XmlExprRegistry[tt.Name.Local]
			if !exists {
				break
			}
			copy := registered.Clone()
			err = d.DecodeElement(&copy, &tt)
			return copy, err
		case xml.EndElement:
			return nil, ErrorEndUnexpected
		}
	}
}

func ParseEnd(d *xml.Decoder) error {
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
