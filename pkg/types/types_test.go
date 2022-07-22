package types

import (
	"fmt"
	"reflect"
	"testing"
)

var TFloat = New("float", float32(0))

type Vec3f struct {
	X, Y, Z float32
}

var MetaJsonName = NewMeta("jsonName")

var TVec3f = New("vec3f", Vec3f{}).Add([]GameTypeProperty{
	{
		Name: "x",
		Type: TFloat,
		Copy: func(source any) any { return source.(Vec3f).X },
		Ref:  func(source any) any { return &source.(*Vec3f).X },
		Set: func(source any, value any) {
			source.(*Vec3f).X = CastNumeric[float32](value, 0)
		},
		Meta: map[GameTypeMeta]any{
			MetaJsonName: "$x",
		},
	}, {
		Name: "y",
		Type: TFloat,
		Copy: func(source any) any { return source.(Vec3f).Y },
		Ref:  func(source any) any { return &source.(*Vec3f).Y },
		Set: func(source any, value any) {
			source.(*Vec3f).Y = CastNumeric[float32](value, 0)
		},
	}, {
		Name: "z",
		Type: TFloat,
		Copy: func(source any) any { return source.(Vec3f).Z },
		Ref:  func(source any) any { return &source.(*Vec3f).Z },
		Set: func(source any, value any) {
			source.(*Vec3f).Z = CastNumeric[float32](value, 0)
		},
	},
})

func TestTypes(m *testing.T) {
	v := Vec3f{1, 2, 3}
	x := Copy(v, []string{"x"})

	fmt.Printf("%v: %v\n", x, reflect.TypeOf(x))

	y := Copy(v, []string{"y"})

	fmt.Printf("%v: %v\n", y, reflect.TypeOf(y))

	xptr := Ref(&v, []string{"x"})

	if xx, ok := xptr.(*float32); ok {
		fmt.Printf("x before %v = %v\n", v.X, *xx)
		*xx = 3.4
		fmt.Printf("x after %v = %v\n", v.X, *xx)
	}

	if jsonName, ok := TVec3f.Props.Get("x").Meta[MetaJsonName].(string); ok {
		fmt.Println(jsonName)
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
