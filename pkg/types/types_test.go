package types

import (
	"fmt"
	"strconv"
	"testing"
)

type Vec2f struct {
	X, Y float32
}

var MetaJsonName = NewMeta("jsonName")

var (
	TFloat = New("float")
	TVec2f = New("vec2f")
)

func init() {
	TFloat.Define(TypeDef{
		Create: func() Value {
			f := float32(0)
			return &f
		},
		ToString: func(v Value) string {
			return fmt.Sprintf("%v", v)
		},
		FromString: func(x string) (Value, error) {
			f, err := strconv.ParseFloat(x, 32)
			return &f, err
		},
	})

	TVec2f.Define(TypeDef{
		Create: func() Value { return &Vec2f{} },
		Props: Props([]Prop{
			NewProp("x", TFloat, func(v *Vec2f) *float32 { return &v.X }),
			NewProp("y", TFloat, func(v *Vec2f) *float32 { return &v.Y }),
		}),
	})
}

func TestTypes(m *testing.T) {
	v := Vec2f{1, 2}
	a, _ := TVec2f.Access(NewPathFromTokens([]string{"x"}))
	a.Set(&v, float32(3))

	fmt.Printf("%+v\n", v.X)
}

func NewProp[S any, V any](name string, t *Type, ref func(source *S) *V) Prop {
	return Prop{
		Name: name,
		Type: t,
		Ref: func(source Value) Value {
			if s, ok := source.(*S); ok {
				return ref(s)
			}
			return nil
		},
		Get: func(source Value, args []Value) Value {
			if s, ok := source.(*S); ok {
				if v := ref(s); v != nil {
					return *v
				}
			}
			return nil
		},
		Set: func(source, value Value) bool {
			if s, ok := source.(*S); ok {
				r := ref(s)
				if r == nil {
					return false
				}
				if v, ok := value.(V); ok {
					*r = v
					return true
				}
				if v, ok := value.(*V); ok {
					*r = *v
					return true
				}
			}
			return false
		},
	}
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func CastNumeric[T Numeric](value any, otherwise T) T {
	switch v := value.(type) {
	case float32:
		return T(v)
	case float64:
		return T(v)
	case int:
		return T(v)
	case int8:
		return T(v)
	case int16:
		return T(v)
	case int32:
		return T(v)
	case int64:
		return T(v)
	case uint:
		return T(v)
	case uint8:
		return T(v)
	case uint16:
		return T(v)
	case uint32:
		return T(v)
	case uint64:
		return T(v)
	case uintptr:
		return T(v)
	}
	return otherwise
}
