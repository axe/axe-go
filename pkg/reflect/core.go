package reflect

import "golang.org/x/exp/constraints"

// type Value = []float32

type ID = uint16

type Type struct {
	id    ID
	name  string
	key   *Type
	value *Type
	props []Prop
}

type Prop struct {
	t    *Type
	name string
	set  bool
	get  bool
}

func (t Type) ID() ID        { return t.id }
func (t Type) Name() string  { return t.name }
func (t Type) Key() *Type    { return t.key }
func (t Type) Value() *Type  { return t.value }
func (t Type) Props() []Prop { return t.props }

type Value interface {
	Type() *Type
	Get(prop string) Value
	Set(values []PropValue)
}

type PropValue struct {
	name  string
	value Value
}

var (
	typeID = NewIncrementor[ID](0, 1)

	Converters = ConversionRegistry{
		converters: make(map[conversion]Converter),
	}
	Calculators = CalculatorRegistry{}
)

func NewType(name string) *Type {
	t := &Type{id: typeID.Get(), name: name}
	return t
}

func (t *Type) AsMapOf(mapOf *Type) *Type {
	c := *t
	c.key = mapOf
	return &c
}

func (t Type) AsListOf(listOf *Type) *Type {
	c := &t
	c.value = listOf
	return c
}

func (t Type) WithProps(props []Prop) *Type {
	c := &t
	c.props = props
	return c
}

func (t Type) WithSelfProps(props func(self *Type) []Prop) *Type {
	c := &t
	c.props = props(c)
	return c
}

func (t Type) Copy() *Type {
	c := t
	c.id = typeID.Get()
	if t.props != nil {
		c.props = make([]Prop, len(t.props))
		for k, v := range t.props {
			c.props[k] = v
		}
	}
	return &c
}

func (t Type) Alias(name string) *Type {
	c := t.Copy()
	c.name = name
	return c
}

type Incrementor[V constraints.Ordered] struct {
	current   V
	increment V
}

func NewIncrementor[V constraints.Ordered](start V, increment V) Incrementor[V] {
	return Incrementor[V]{current: start, increment: increment}
}

func (i Incrementor[V]) Get() V {
	c := i.current
	i.current = i.current + i.increment
	return c
}

/*
type Instance interface {
	MapElements() Iterator[string, Instance]
	ListElements() Iterator[int, Instance]
	Type() *Type
	Prop(name string) *Instance
}

type PropertyDescriptor struct {
	Name string
	Type *Type
}

type Property struct {
	Descriptor PropertyDescriptor
	Index      int
	Value      Value
}

func (p Property) Exists() bool { return p.Descriptor.Type != nil }

type Object interface {
	GetProperties() []PropertyDescriptor
	SetProperties(props []Property) error
	GetProperty(name string) Property
	GetPropertyAt(name string, index int) Property
}
*/

type Calculator interface {
	Size() int
	Zero() Value
	DistanceSq(a Value, b Value) float32
	AddScaled(out Value, a Value, b Value, bScale float32)
	Mul(out Value, a Value, scale Value)
}

type CalculatorRegistry map[ID]Calculator

func (r CalculatorRegistry) Set(t *Type, c Calculator) {
	r[t.id] = c
}
func (r CalculatorRegistry) Get(t *Type) Calculator {
	return r[t.id]
}

type Converter func(from Value, out Value)

type conversion struct{ from, to ID }

type ConversionRegistry struct {
	converters map[conversion]Converter
}

func (r ConversionRegistry) Get(from *Type, to *Type) Converter {
	return r.converters[conversion{from: from.id, to: to.id}]
}

func (r ConversionRegistry) Set(from *Type, to *Type, converter Converter) {
	r.converters[conversion{from: from.id, to: to.id}] = converter
}

// EXAMPLE

// Scalar Types
// - Integer = any whole number, positive or negative
// - Float = any floating precision number with 32-bits of precision
// - Index = whole positive number
// - Angle = a number representing an angle. internally stored as radians but can be translated to and from degrees
// Dimensional Types
// - Point = somewhere in space
// - Vector = points in a direction with a magnitude
// - UnitVector = vector with length of 1
// - Scale = scales something, default value of 1
// - Alignment = values typically between 0->1 where 0.5 is middle
// Object Types
// - Sprite
// - Bone
// - Skeleton
// - Mesh

var (
	TypeFloat   = NewType("float")
	TypeInteger = NewType("integer")
	TypeBoolean = NewType("boolean")
	TypeMap     = NewType("map").WithProps([]Prop{
		{name: "size", t: TypeInteger},
	})
	TypeList = NewType("list").WithProps([]Prop{
		{name: "size", t: TypeInteger},
	})
	TypePoint2D = NewType("point2d").WithProps([]Prop{
		{name: "x", t: TypeFloat},
		{name: "y", t: TypeFloat},
	})
	TypePoint3D = NewType("point3d").WithProps([]Prop{
		{name: "x", t: TypeFloat},
		{name: "y", t: TypeFloat},
		{name: "z", t: TypeFloat},
	})
	TypeScale2D  = TypePoint2D.Alias("scale2d")
	TypeScale3D  = TypePoint3D.Alias("scale3d")
	TypeVector2D = NewType("vector2d").WithProps([]Prop{
		{name: "x", t: TypeFloat},
		{name: "y", t: TypeFloat},
		{name: "length", t: TypeFloat},
	})
	TypeVector3D = NewType("vector3d").WithProps([]Prop{
		{name: "x", t: TypeFloat},
		{name: "y", t: TypeFloat},
		{name: "z", t: TypeFloat},
		{name: "length", t: TypeFloat},
	})
	TypeAngle = NewType("angle").WithProps([]Prop{
		{name: "radians", t: TypeFloat},
		{name: "degrees", t: TypeFloat},
	})
	TypeQuaternion = NewType("quaternion").WithProps([]Prop{
		{name: "v", t: TypeVector3D},
		{name: "t", t: TypeAngle},
	})
	TypeSprite = NewType("sprite").WithProps([]Prop{
		{name: "angle", t: TypeAngle},
		{name: "position", t: TypePoint2D},
		{name: "alignment", t: TypePoint2D},
		{name: "size", t: TypePoint2D},
		{name: "scale", t: TypeScale2D},
	})
	TypeBone = NewType("bone").WithSelfProps(func(bone *Type) []Prop {
		return []Prop{
			{name: "offset", t: TypeVector3D},
			{name: "rotation", t: TypeQuaternion},
			{name: "children", t: TypeMap.AsMapOf(bone)},
		}
	})
	TypeSkeleton = NewType("skeleton").WithProps([]Prop{
		{name: "bones", t: TypeMap.AsMapOf(TypeBone)},
	})
	TypeMesh = NewType("mesh").WithProps([]Prop{
		{name: "position", t: TypePoint3D},
		{name: "rotation", t: TypeQuaternion},
		{name: "scale", t: TypeScale3D},
		{name: "skeleton", t: TypeSkeleton},
	})
)

func ptr[V any](value V) *V {
	return &value
}

type Float float32

var _ Value = Float(0)

func (f Float) Type() *Type            { return TypeFloat }
func (f Float) Get(prop string) Value  { return nil }
func (f Float) Set(values []PropValue) {}

type Vector3D struct {
	X float32
	Y float32
}

var _ Value = &Vector3D{}

func (v Vector3D) Type() *Type { return TypeVector3D }
func (v Vector3D) Get(prop string) Value {
	switch prop {
	case "x":
		return Float(v.X)
	case "y":
		return Float(v.Y)
	}
	return nil
}
func (v *Vector3D) Set(values []PropValue) {
	for _, p := range values {
		switch p.name {
		case "x":
			v.X = float32(p.value.(Float))
		case "y":
			v.Y = float32(p.value.(Float))
		}
	}
}

// Type {id, name, props{name, type, get, set}, values{valueType, keyType, get, set}}

/**

type Sprite struct {
	Angle      float32
	Positition [2]float32
	Offset     [2]float32
	Children   []Sprite
}

var _ Object = &Sprite{}

func (s Sprite) GetProperties() []PropertyDescriptor {
	return []PropertyDescriptor{
		{Type: Angle, Name: "angle"},
		{Type: Point2D, Name: "position"},
		{Type: Point2D, Name: "offset"},
		{Type: SpriteType, Name: "children", Count: len(s.Children)},
	}
}

func (s *Sprite) SetProperties(props []Property) error {
	for _, p := range props {
		switch p.Descriptor.Name {
		case "angle":
			s.Angle = p.Value[0]
		case "position":
			s.Positition = [2]float32(p.Value)
		case "offset":
			s.Offset = [2]float32(p.Value)
		case "children":
			// s.Children[p.Index] = p TODO
		}
	}
	return nil
}

func (s Sprite) GetProperty(name string) Property {
	switch name {
	case "angle":
		return Property{PropertyDescriptor{Type: Angle, Name: "angle"}, 0, []float32{s.Angle}}
	case "position":
		return Property{PropertyDescriptor{Type: Point2D, Name: "position"}, 0, s.Positition[:]}
	case "offset":
		return Property{PropertyDescriptor{Type: Point2D, Name: "offset"}, 0, s.Offset[:]}
	}
}
func (s Sprite) GetPropertyAt(name string, index int) Property {
	switch name {
	case "children":
		return Property{Index: index, Descriptor: PropertyDescriptor{Name: "children", Type: SpriteType, Count: len(s.Children)}}
	default:
		return s.GetProperty(name)
	}
}

*/
