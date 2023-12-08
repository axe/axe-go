package gfx

import "github.com/axe/axe-go/pkg/util"

type Coord struct {
	X float32
	Y float32
}

func (mp Coord) Equals(other Coord) bool {
	return mp.X == other.X && mp.Y == other.Y
}

func (c Coord) Max(other Coord) Coord {
	return Coord{X: util.Max(c.X, other.X), Y: util.Max(c.Y, other.Y)}
}

func (c Coord) Min(other Coord) Coord {
	return Coord{X: util.Min(c.X, other.X), Y: util.Min(c.Y, other.Y)}
}

func (c *Coord) Set(x, y float32) {
	c.X = x
	c.Y = y
}

func (c Coord) IsZero() bool {
	return c.X == 0 && c.Y == 0
}

func (a Coord) Lerp(b Coord, delta float32) Coord {
	return Coord{
		X: util.Lerp(a.X, b.X, delta),
		Y: util.Lerp(a.Y, b.Y, delta),
	}
}

type Texture struct {
	Name     string
	Metadata any
}

func (tex Texture) Tile() Tile {
	return Tile{
		Texture: tex,
		Left:    0,
		Right:   1,
		Top:     0,
		Bottom:  1,
	}
}

func (tex Texture) Grid(columns, rows, columnWidth, rowHeight, textureWidth, textureHeight, offsetX, offsetY int) [][]Tile {
	tiles := make([][]Tile, rows)
	sx := 1.0 / float32(textureWidth)
	sy := 1.0 / float32(textureHeight)

	for y := 0; y < rows; y++ {
		tiles[y] = make([]Tile, columns)
		for x := 0; x < columns; x++ {
			tiles[y][x] = Tile{
				Texture: tex,
				Left:    float32(x*columnWidth+offsetX) * sx,
				Right:   float32(x*columnWidth+columnWidth-1+offsetX) * sx,
				Top:     float32(y*rowHeight+offsetY) * sy,
				Bottom:  float32(y*rowHeight+rowHeight-1+offsetY) * sy,
			}
		}
	}
	return tiles
}

func (tex Texture) Frames(width, height, textureWidth, textureHeight float32, topLeftCorners []Coord) []Tile {
	return nil
}

type Tile struct {
	Texture                  Texture
	Left, Top, Right, Bottom float32
}

func (t Tile) Coord(dx, dy float32) TextureCoord {
	return TextureCoord{
		Texture: t.Texture,
		Coord: Coord{
			X: util.Lerp(t.Left, t.Right, dx),
			Y: util.Lerp(t.Top, t.Bottom, dy),
		},
	}
}

type TextureCoord struct {
	Texture Texture
	Coord
}

type Primitive uint8

const (
	PrimitiveTriangle Primitive = iota
	PrimitiveLine
	PrimitiveQuad
	PrimitiveNone
)

type Blend uint8

const (
	BlendAlpha Blend = iota
	BlendAlphaAdd
	BlendAdd
	BlendColor
	BlendMinus
	BlendPremultAlpha
	BlendModulate
	BlendXor
	BlendNone
)
