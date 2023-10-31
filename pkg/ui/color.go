package ui

import (
	"fmt"
	"strconv"
)

//go:generate go run colors.go

type Color struct {
	R, G, B, A float32
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

func (c Color) IsZero() bool {
	return c.R == 0 && c.G == 0 && c.B == 0 && c.A == 0
}

func (c Color) ToInts() (r, g, b, a int) {
	r = int(clamp(c.R*255, 0, 255))
	g = int(clamp(c.G*255, 0, 255))
	b = int(clamp(c.B*255, 0, 255))
	a = int(clamp(c.A*255, 0, 255))
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
