// The steer package contains steering behaviors for autonomous movement.
//
// Concepts:
// - Target
// - Selector
//
// Behaviors
// - Match velocity = given a query around the subject and all applicable entities, try to match their velocity magnitude
// - Alignment = given a query around the subject
// - Cohesion
// -
//
// Accumulators
package steer

import (
	"github.com/axe/axe-go/pkg/data"
	"github.com/axe/axe-go/pkg/id"
)

type Force struct {
	Scale    float32
	Velocity []float32
	Weight   float32
}

type ForceGenerator interface {
	GenerateForces(subject Subject, dt float32, gen func(Force))
}

type Behavior interface {
	GetForce(subject Subject, dt float32) Force
}

type Object struct {
}

type Subject struct {
	Object

	MaximumVelocity float32
}

type Filter func(subject Subject, object Object) bool

type Query struct {
	Position []float32
}

type SubjectSource interface {
	GetSubject() Subject
}

type ObjectSource interface {
	GetObject() Object
}

// CustomData (option #1)

type CustomDataType int

const (
	CustomDataTypeInt CustomDataType = iota
	CustomDataTypeBool
	CustomDataTypeFloat
	CustomDataTypeString
	CustomDataTypeCount
)

type CustomDataProperty struct {
	Name      id.Identifier
	Type      CustomDataType
	TypeIndex int
}

type CustomDataDefinition struct {
	Name       id.Identifier
	Properties id.SparseMap[CustomDataProperty, uint16]
	Counts     []int
}

func NewCustomDataDefinition(name id.Identifier) *CustomDataDefinition {
	return &CustomDataDefinition{
		Name:       name,
		Properties: id.NewSparseMap[CustomDataProperty, uint16](),
		Counts:     make([]int, CustomDataTypeCount),
	}
}

func (def *CustomDataDefinition) Add(name id.Identifier, dataType CustomDataType) {
	def.Properties.Set(name, CustomDataProperty{
		Name:      name,
		Type:      dataType,
		TypeIndex: def.Counts[dataType],
	})
	def.Counts[dataType]++
}

func (def *CustomDataDefinition) New() CustomData {
	return NewCustomData(def)
}

type CustomData struct {
	Definition *CustomDataDefinition
	Ints       []int
	Bools      []bool
	Floats     []float32
	Strings    []string
}

func NewCustomData(def *CustomDataDefinition) CustomData {
	return CustomData{
		Definition: def,
		Ints:       make([]int, def.Counts[CustomDataTypeInt]),
		Bools:      make([]bool, def.Counts[CustomDataTypeBool]),
		Floats:     make([]float32, def.Counts[CustomDataTypeFloat]),
		Strings:    make([]string, def.Counts[CustomDataTypeString]),
	}
}
func getCustomDataPointer[V any](cd *CustomData, name id.Identifier, values []V, valueType CustomDataType) *V {
	prop := cd.Definition.Properties.Get(name)
	if prop.Name == name && prop.Type == valueType {
		return &values[prop.TypeIndex]
	}
	return nil
}
func (cd *CustomData) Int(name id.Identifier) *int {
	return getCustomDataPointer(cd, name, cd.Ints, CustomDataTypeInt)
}
func (cd *CustomData) Bool(name id.Identifier) *bool {
	return getCustomDataPointer(cd, name, cd.Bools, CustomDataTypeBool)
}
func (cd *CustomData) Float(name id.Identifier) *float32 {
	return getCustomDataPointer(cd, name, cd.Floats, CustomDataTypeFloat)
}
func (cd *CustomData) String(name id.Identifier) *string {
	return getCustomDataPointer(cd, name, cd.Strings, CustomDataTypeString)
}

// ExtraData (option #2)

// Most types are defined by the engine
// The user can define custom types in editor
// A type implements another if the type has all the same things on the other
// Scripted entities are a dynamic type where all components are properties but there's also an inner custom type where the user's scripted variables go to
// A type can have values with zero or more arguments.
// Every value must be uniquely named
// Value with no arguments are like properties
// Values with arguments are like methods
// Values can have arguments with defaults that feel like a property
// Collection type has len, get, set, add, contains, and iterator

type Param struct {
	Name    id.Identifier
	Type    Typ
	Default *Val
}

type Field struct {
	Name   id.Identifier
	Params []Param
	Type   Typ
}

type Typ struct {
	Name id.Identifier
}

type Val struct {
	Type string
	Data *data.Bytes
}
