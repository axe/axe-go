package color

import (
	"fmt"
	"strconv"

	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/util"
)

//go:generate go run colors.go

type Color struct {
	R, G, B, A float32
}

var _ Able = Color{}

type Colors struct {
	ds.EnumMap[Themed, Able]
}

func NewColors(m map[Themed]Able) Colors {
	return Colors{
		EnumMap: ds.NewEnumMap(m),
	}
}

func New(r, g, b, a float32) Color {
	return Color{R: r, B: b, G: g, A: a}
}

func FromInts(r, g, b, a int) Color {
	return Color{R: float32(r) / 255, G: float32(g) / 255, B: float32(b) / 255, A: float32(a) / 255}
}

func MustParse(input string) Color {
	color, err := Parse(input)
	if err != nil {
		panic(err)
	}
	return color
}

func Parse(input string) (Color, error) {
	named, exists := Named(input)
	if exists {
		return named, nil
	}

	if input[0] == '#' {
		input = input[1:]
	}

	switch len(input) {
	case 1:
		rgbHalf, err := strconv.ParseInt(input[0:1], 16, 32)
		if err != nil {
			return White, err
		}
		rgb := rgbHalf | (rgbHalf << 8)
		return FromInts(int(rgb), int(rgb), int(rgb), 255), nil
	case 2:
		rgb, err := strconv.ParseInt(input[0:2], 16, 32)
		if err != nil {
			return White, err
		}
		return FromInts(int(rgb), int(rgb), int(rgb), 255), nil
	case 3:
		rHalf, err := strconv.ParseInt(input[0:1], 16, 32)
		if err != nil {
			return White, err
		}
		gHalf, err := strconv.ParseInt(input[1:2], 16, 32)
		if err != nil {
			return White, err
		}
		bHalf, err := strconv.ParseInt(input[2:3], 16, 32)
		if err != nil {
			return White, err
		}
		r := rHalf | (rHalf << 8)
		g := gHalf | (gHalf << 8)
		b := bHalf | (bHalf << 8)
		return FromInts(int(r), int(g), int(b), 255), nil
	case 4:
		rHalf, err := strconv.ParseInt(input[0:1], 16, 32)
		if err != nil {
			return White, err
		}
		gHalf, err := strconv.ParseInt(input[1:2], 16, 32)
		if err != nil {
			return White, err
		}
		bHalf, err := strconv.ParseInt(input[2:3], 16, 32)
		if err != nil {
			return White, err
		}
		aHalf, err := strconv.ParseInt(input[3:4], 16, 32)
		if err != nil {
			return White, err
		}
		r := rHalf | (rHalf << 8)
		g := gHalf | (gHalf << 8)
		b := bHalf | (bHalf << 8)
		a := aHalf | (aHalf << 8)
		return FromInts(int(r), int(g), int(b), int(a)), nil
	case 6:
		r, err := strconv.ParseInt(input[0:2], 16, 32)
		if err != nil {
			return White, err
		}
		g, err := strconv.ParseInt(input[2:4], 16, 32)
		if err != nil {
			return White, err
		}
		b, err := strconv.ParseInt(input[4:6], 16, 32)
		if err != nil {
			return White, err
		}
		return FromInts(int(r), int(g), int(b), 255), nil
	case 8:
		r, err := strconv.ParseInt(input[0:2], 16, 32)
		if err != nil {
			return White, err
		}
		g, err := strconv.ParseInt(input[2:4], 16, 32)
		if err != nil {
			return White, err
		}
		b, err := strconv.ParseInt(input[4:6], 16, 32)
		if err != nil {
			return White, err
		}
		a, err := strconv.ParseInt(input[6:8], 16, 32)
		if err != nil {
			return White, err
		}
		return FromInts(int(r), int(g), int(b), int(a)), nil
	}

	return White, fmt.Errorf("Color %s could not be parsed", input)
}

func (c Color) GetColor(b Themeable) Color {
	return c
}

func (c Color) Ptr() *Color {
	return &c
}

func (c Color) Alpha(a float32) Color {
	c.A = a
	return c
}

func (c Color) AlphaScale(a float32) Color {
	c.A *= a
	return c
}

func (c Color) Red(v float32) Color {
	c.R = v
	return c
}

func (c Color) Green(v float32) Color {
	c.G = v
	return c
}

func (c Color) Blue(v float32) Color {
	c.B = v
	return c
}

func (c Color) Darken(scale float32) Color {
	c.R -= c.R * scale
	c.G -= c.G * scale
	c.B -= c.B * scale
	return c
}

func (c Color) Lighten(scale float32) Color {
	c.R += (1 - c.R) * scale
	c.G += (1 - c.G) * scale
	c.B += (1 - c.B) * scale
	return c
}

func (c Color) Lerp(to Color, delta float32) Color {
	return Color{
		R: util.Lerp(c.R, to.R, delta),
		G: util.Lerp(c.G, to.G, delta),
		B: util.Lerp(c.B, to.B, delta),
		A: util.Lerp(c.A, to.A, delta),
	}
}

func (c Color) Multiply(o Color) Color {
	return Color{
		R: c.R * o.R,
		G: c.G * o.G,
		B: c.B * o.B,
		A: c.A * o.A,
	}
}

func (c Color) IsZero() bool {
	return c.R == 0 && c.G == 0 && c.B == 0 && c.A == 0
}

func (c Color) ToInts() (r, g, b, a int) {
	r = int(util.Clamp(c.R*255, 0, 255))
	g = int(util.Clamp(c.G*255, 0, 255))
	b = int(util.Clamp(c.B*255, 0, 255))
	a = int(util.Clamp(c.A*255, 0, 255))
	return
}

func (c Color) ToFloats() []float32 {
	return []float32{c.R, c.G, c.B, c.A}
}

func (c Color) MarshalText() ([]byte, error) {
	r, g, b, a := c.ToInts()
	if a == 255 {
		return []byte(fmt.Sprintf("#%.2X%.2X%.2X", r, g, b)), nil
	}
	return []byte(fmt.Sprintf("#%.2X%.2X%.2X%.2X", r, g, b, a)), nil
}

func (c *Color) UnmarshalText(text []byte) error {
	s := string(text)
	named, exists := Named(s)
	if exists {
		*c = named
		return nil
	}
	*c = MustParse(s)
	return nil
}
