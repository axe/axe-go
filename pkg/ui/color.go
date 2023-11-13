package ui

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

var _ Colorable = Color{}

type Colorable interface {
	GetColor(b *Base) Color
}

func GetColor(colorable Colorable, b *Base) (Color, bool) {
	if colorable == nil {
		return Color{}, false
	} else {
		return colorable.GetColor(b), true
	}
}

func NewColor(r, g, b, a float32) Color {
	return Color{R: r, B: b, G: g, A: a}
}

func ColorFromInts(r, g, b, a int) Color {
	return Color{R: float32(r) / 255, G: float32(g) / 255, B: float32(b) / 255, A: float32(a) / 255}
}

func ColorFromHex(hex string) Color {
	if hex[0] == '#' {
		hex = hex[1:]
	}
	switch len(hex) {
	case 1:
		rgbHalf, _ := strconv.ParseInt(hex[0:1], 16, 32)
		rgb := rgbHalf | (rgbHalf << 8)
		return ColorFromInts(int(rgb), int(rgb), int(rgb), 255)
	case 2:
		rgb, _ := strconv.ParseInt(hex[0:2], 16, 32)
		return ColorFromInts(int(rgb), int(rgb), int(rgb), 255)
	case 3:
		rHalf, _ := strconv.ParseInt(hex[0:1], 16, 32)
		gHalf, _ := strconv.ParseInt(hex[1:2], 16, 32)
		bHalf, _ := strconv.ParseInt(hex[2:3], 16, 32)
		r := rHalf | (rHalf << 8)
		g := gHalf | (gHalf << 8)
		b := bHalf | (bHalf << 8)
		return ColorFromInts(int(r), int(g), int(b), 255)
	case 4:
		rHalf, _ := strconv.ParseInt(hex[0:1], 16, 32)
		gHalf, _ := strconv.ParseInt(hex[1:2], 16, 32)
		bHalf, _ := strconv.ParseInt(hex[2:3], 16, 32)
		aHalf, _ := strconv.ParseInt(hex[3:4], 16, 32)
		r := rHalf | (rHalf << 8)
		g := gHalf | (gHalf << 8)
		b := bHalf | (bHalf << 8)
		a := aHalf | (aHalf << 8)
		return ColorFromInts(int(r), int(g), int(b), int(a))
	case 6:
		r, _ := strconv.ParseInt(hex[0:2], 16, 32)
		g, _ := strconv.ParseInt(hex[2:4], 16, 32)
		b, _ := strconv.ParseInt(hex[4:6], 16, 32)
		return ColorFromInts(int(r), int(g), int(b), 255)
	case 8:
		r, _ := strconv.ParseInt(hex[0:2], 16, 32)
		g, _ := strconv.ParseInt(hex[2:4], 16, 32)
		b, _ := strconv.ParseInt(hex[4:6], 16, 32)
		a, _ := strconv.ParseInt(hex[6:8], 16, 32)
		return ColorFromInts(int(r), int(g), int(b), int(a))
	}
	panic("Invalid hex color: " + hex)
}

func (c Color) GetColor(b *Base) Color {
	return c
}

func (c Color) Ptr() *Color {
	return &c
}

func (c Color) Alpha(a float32) Color {
	c.A = a
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

func (c Color) MarshalText() ([]byte, error) {
	r, g, b, a := c.ToInts()
	if a == 255 {
		return []byte(fmt.Sprintf("#%.2X%.2X%.2X", r, g, b)), nil
	}
	return []byte(fmt.Sprintf("#%.2X%.2X%.2X%.2X", r, g, b, a)), nil
}

func (c *Color) UnmarshalText(text []byte) error {
	s := string(text)
	named := ColorNamed(s)
	if !named.IsZero() {
		*c = named
		return nil
	}
	*c = ColorFromHex(s)
	return nil
}

type ThemeColor uint8

var _ Colorable = ThemeColor(0)

func (t ThemeColor) GetColor(b *Base) Color {
	if b == nil {
		return ColorWhite
	}
	colorable := b.Colors.Get(t)
	if colorable == nil {
		colorable = b.ui.Theme.Colors.Get(t)
	}
	if colorable != nil {
		return colorable.GetColor(b)
	}
	return ColorWhite
}

type ColorModify func(Color) Color

func (a ColorModify) Then(b ColorModify) ColorModify {
	return func(c Color) Color {
		return b(a(c))
	}
}

type modifiedThemeColor struct {
	modify     ColorModify
	themeColor ThemeColor
}

var _ Colorable = modifiedThemeColor{}

func (m modifiedThemeColor) GetColor(b *Base) Color {
	return m.modify(m.themeColor.GetColor(b))
}

func (t ThemeColor) Modify(modify ColorModify) Colorable {
	return modifiedThemeColor{modify: modify, themeColor: t}
}

func Lighten(scale float32) ColorModify {
	return func(c Color) Color {
		return c.Lighten(scale)
	}
}

func Darken(scale float32) ColorModify {
	return func(c Color) Color {
		return c.Darken(scale)
	}
}

func Alpha(alpha float32) ColorModify {
	return func(c Color) Color {
		return c.Alpha(alpha)
	}
}

type Colors struct {
	ds.EnumMap[ThemeColor, Colorable]
}

func NewColors(m map[ThemeColor]Colorable) Colors {
	return Colors{
		EnumMap: ds.NewEnumMap(m),
	}
}

func colorableNil(c Colorable) bool {
	return c == nil
}
