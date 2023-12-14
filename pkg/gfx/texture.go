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
	Name string
	Info TextureInfo
}

func (tex *Texture) IsDefined() bool {
	return tex != nil && tex.Info.IsDefined()
}

func (tex *Texture) Texel(x, y float32) Texel {
	xy := TexelXY{X: x, Y: y}
	if tex.IsDefined() {
		x, y = xy.UV(tex)
		return TexelUV{X: x, Y: y}
	}
	return xy
}

func (tex *Texture) Coord(x, y float32) TextureCoord {
	return TextureCoord{Texture: tex, Texel: tex.Texel(x, y)}
}

func (tex *Texture) Grid(columns, rows, columnWidth, rowHeight, offsetX, offsetY int) [][]Tile {
	tiles := make([][]Tile, rows)

	for y := 0; y < rows; y++ {
		tiles[y] = make([]Tile, columns)
		for x := 0; x < columns; x++ {
			left := x*columnWidth + offsetX
			top := y*rowHeight + offsetY

			tiles[y][x] = Tile{
				Texture:     tex,
				TopLeft:     tex.Texel(float32(left), float32(top)),
				BottomRight: tex.Texel(float32(left+columnWidth-1), float32(top+rowHeight-1)),
			}
		}
	}
	return tiles
}

func (tex *Texture) Frames(width, height float32, topLeftCorners []Coord) []Tile {
	tiles := make([]Tile, len(topLeftCorners))

	for i, corner := range topLeftCorners {
		tiles[i] = Tile{
			Texture:     tex,
			TopLeft:     tex.Texel(corner.X, corner.Y),
			BottomRight: tex.Texel(corner.X+width-1, corner.Y+height-1),
		}
	}

	return tiles
}

type TextureInfo struct {
	Width, Height       int
	InvWidth, InvHeight float32
	Metadata            any
}

func (info TextureInfo) IsDefined() bool {
	return info.Width != 0 && info.Height != 0
}

func (td *TextureInfo) SetDimensions(width, height int) {
	td.Width = width
	td.InvWidth = 1.0 / float32(width)
	td.Height = height
	td.InvHeight = 1.0 / float32(height)
}

type Texel interface {
	UV(tex *Texture) (u, v float32)
}

type TexelUV Coord

func (t TexelUV) UV(tex *Texture) (u, v float32) {
	return t.X, t.Y
}

type TexelXY Coord

func (t TexelXY) UV(tex *Texture) (u, v float32) {
	if tex != nil {
		u = (t.X + 0.5) * tex.Info.InvWidth
		v = (t.Y + 0.5) * tex.Info.InvHeight
	}
	return
}

func (tex *Texture) Tile() Tile {
	return Tile{
		Texture:     tex,
		TopLeft:     TexelUV{X: 0, Y: 0},
		BottomRight: TexelUV{X: 1, Y: 1},
	}
}

type Tile struct {
	Texture              *Texture
	TopLeft, BottomRight Texel
}

func (t Tile) Coord(dx, dy float32) TextureCoord {
	u0, v0 := t.TopLeft.UV(t.Texture)
	u1, v1 := t.BottomRight.UV(t.Texture)

	return TextureCoord{
		Texture: t.Texture,
		Texel: TexelUV{
			X: util.Lerp(u0, u1, dx),
			Y: util.Lerp(v0, v1, dy),
		},
	}
}

type TextureCoord struct {
	Texture *Texture
	Texel
}

func (tc TextureCoord) UV() (u, v float32) {
	if tc.Texel != nil {
		u, v = tc.Texel.UV(tc.Texture)
	}
	return
}

func (from TextureCoord) Lerp(to TextureCoord, delta float32) TextureCoord {
	u0, v0 := from.UV()
	u1, v1 := to.UV()

	return TextureCoord{
		Texture: from.Texture,
		Texel: TexelUV{
			X: util.Lerp(u0, u1, delta),
			Y: util.Lerp(v0, v1, delta),
		},
	}
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
