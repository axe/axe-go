package ui

import "github.com/axe/axe-go/pkg/buf"

type Coord struct {
	X float32
	Y float32
}

func (mp Coord) Equals(other Coord) bool {
	return mp.X == other.X && mp.Y == other.Y
}

type TexCoord struct {
	Coord
	Texture string
}

type Tile struct {
	Coords  Bounds
	Texture string
}

func (t Tile) Coord(dx, dy float32) TexCoord {
	return TexCoord{
		Texture: t.Texture,
		Coord: Coord{
			X: lerp(t.Coords.Left, t.Coords.Right, dx),
			Y: lerp(t.Coords.Top, t.Coords.Bottom, dy),
		},
	}
}

type ExtentTile struct {
	Tile
	Extent Bounds
}

type Vertex struct {
	X, Y     float32
	Coord    TexCoord
	HasCoord bool
	Color    Color
	HasColor bool
}

type VertexModifier func(*Vertex)

func (v *Vertex) AddColor(r, g, b, a float32) {
	if v.HasColor {
		v.Color.R *= r
		v.Color.G *= g
		v.Color.B *= b
		v.Color.A *= a
	} else {
		v.Color.R = r
		v.Color.G = g
		v.Color.B = b
		v.Color.A = a
		v.HasColor = true
	}
}

func (v *Vertex) SetCoord(texture string, x, y float32) {
	v.Coord.Texture = texture
	v.Coord.X = x
	v.Coord.Y = y
	v.HasCoord = true
}

type VertexIterator = buf.DataIterator[Vertex, VertexBuffer]

func NewVertexIterator(vb *VertexBuffers) VertexIterator {
	return buf.NewDataIterator(&vb.Buffers)
}

type IndexIterator = buf.IndexIterator[Vertex, VertexBuffer]

func NewIndexIterator(vb *VertexBuffers) IndexIterator {
	return buf.NewIndexIterator(&vb.Buffers)
}

func InitVertexBuffer(vb *VertexBuffer, capacity int) {
	vb.Init(capacity)
}

func NewVertexBuffers(capacity int, buffers int) *VertexBuffers {
	return &VertexBuffers{
		Buffers: *buf.NewBuffers[Vertex](4096, 4, InitVertexBuffer),
	}
}

type VertexBuffers struct {
	buf.Buffers[Vertex, VertexBuffer]
}

func (vb *VertexBuffers) ClipStart(bounds Bounds) *VertexBuffer {
	added := vb.AddBuffer()
	added.clip = bounds
	return added
}

func (vb *VertexBuffers) ClipEnd() *VertexBuffer {
	prev := vb.At(vb.Current() - 1)
	ended := vb.AddBuffer()
	if prev != nil {
		ended.clip = prev.clip
	}
	return ended
}

func (vb *VertexBuffers) ClipMaybe(bounds Bounds, render func(vb *VertexBuffers)) {
	if bounds.IsZero() {
		render(vb)
	} else {
		vb.ClipStart(bounds)
		render(vb)
		vb.ClipEnd()
	}
}

type VertexBuffer struct {
	buf.Buffer[Vertex]
	clip Bounds
}

func (b *VertexBuffer) Init(capacity int) {
	b.Buffer.Init(capacity)
	b.clip = Bounds{}
}

func (b VertexBuffer) Empty() bool {
	return b.Buffer.Empty()
}

func (b VertexBuffer) Remaining() int {
	return b.Buffer.Remaining()
}

func (b *VertexBuffer) AddIndexQuad(i int) {
	b.AddIndex(i, i+1, i+2, i+2, i+3, i)
}

var relativeQuad = []int{0, 1, 2, 2, 3, 0}

func (b *VertexBuffer) AddQuad(v ...Vertex) {
	b.AddRelative(v, relativeQuad)
}

func (b *VertexBuffer) Clear() {
	b.Buffer.Clear()
	b.clip = Bounds{}
}

func (b *VertexBuffer) Clip() Bounds {
	return b.clip
}
