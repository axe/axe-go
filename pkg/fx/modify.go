package fx

import (
	"github.com/axe/axe-go/pkg/util"
)

type Modify interface {
	Modify(particle []float32, format *Format, dt float32)
	Modifies(attr Attribute) bool
}

type ModifyAge struct {
	Age Attribute
}

func (m ModifyAge) Modify(particle []float32, format *Format, dt float32) {
	format.Get(m.Age, particle)[0] += dt
}
func (m ModifyAge) Modifies(attr Attribute) bool {
	return m.Age.id == attr.id
}

type ModifyAdder struct {
	Value Attribute
	Add   Attribute
}

func (m ModifyAdder) Modify(particle []float32, format *Format, dt float32) {
	Add(format.Get(m.Value, particle), format.Get(m.Add, particle))
}
func (m ModifyAdder) Modifies(attr Attribute) bool {
	return m.Value.id == attr.id
}

type ModifyScalar struct {
	Value  Attribute
	Scalar Attribute
}

func (m ModifyScalar) Modify(particle []float32, format *Format, dt float32) {
	value := format.Get(m.Value, particle)
	scalar := format.Get(m.Scalar, particle)[0]
	scalarScaled := util.Pow(scalar, dt)

	MultiplyScalar(value, scalarScaled)
}
func (m ModifyScalar) Modifies(attr Attribute) bool {
	return m.Value.id == attr.id
}

type ModifyLength struct {
	Value Attribute
	Add   Attribute
}

func (m ModifyLength) Modify(particle []float32, format *Format, dt float32) {
	value := format.Get(m.Value, particle)
	lengthSq := LengthSq(value)
	if lengthSq == 0 {
		return
	}
	add := format.Get(m.Add, particle)[0]
	addScaled := add * dt
	length := util.Sqrt(lengthSq)
	newLength := util.Max(0, length+addScaled)
	scale := util.Div(newLength, length)

	MultiplyScalar(value, scale)
}
func (m ModifyLength) Modifies(attr Attribute) bool {
	return m.Value.id == attr.id
}

type ModifyNone struct{}

func (m ModifyNone) Modify(particle []float32, format *Format, dt float32) {

}
func (m ModifyNone) Modifies(attr Attribute) bool {
	return false
}

type ModifyList struct {
	List []Modify
}

func (m ModifyList) Modify(particle []float32, format *Format, dt float32) {
	for _, modify := range m.List {
		modify.Modify(particle, format, dt)
	}
}
func (m ModifyList) Modifies(attr Attribute) bool {
	for _, modify := range m.List {
		if modify.Modifies(attr) {
			return true
		}
	}
	return false
}

type Modifys []Modify

func (m Modifys) Add(mod Modify) Modifys {
	return append(m, mod)
}

func (m Modifys) Age(attr Attribute) Modifys {
	return append(m, ModifyAge{Age: attr})
}

func (m Modifys) Adder(attr Attribute, add Attribute) Modifys {
	return append(m, ModifyAdder{Value: attr, Add: add})
}

func (m Modifys) Scalar(attr Attribute, scalar Attribute) Modifys {
	return append(m, ModifyScalar{Value: attr, Scalar: scalar})
}

func (m Modifys) Length(attr Attribute, add Attribute) Modifys {
	return append(m, ModifyLength{Value: attr, Add: add})
}
