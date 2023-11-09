package ui

import "github.com/axe/axe-go/pkg/buf"

type Coord struct {
	X float32
	Y float32
}

func (mp Coord) Equals(other Coord) bool {
	return mp.X == other.X && mp.Y == other.Y
}

func (c Coord) Max(other Coord) Coord {
	return Coord{X: max(c.X, other.X), Y: max(c.Y, other.Y)}
}

func (c Coord) Min(other Coord) Coord {
	return Coord{X: min(c.X, other.X), Y: min(c.Y, other.Y)}
}

func (c *Coord) Set(x, y float32) {
	c.X = x
	c.Y = y
}

func (c Coord) IsZero() bool {
	return c.X == 0 && c.Y == 0
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
			X: Lerp(t.Coords.Left, t.Coords.Right, dx),
			Y: Lerp(t.Coords.Top, t.Coords.Bottom, dy),
		},
	}
}

func TileGrid(columns, rows, columnWidth, rowHeight, textureWidth, textureHeight, offsetX, offsetY int, texture string) [][]Tile {
	tiles := make([][]Tile, rows)
	sx := 1.0 / float32(textureWidth)
	sy := 1.0 / float32(textureHeight)

	for y := 0; y < rows; y++ {
		tiles[y] = make([]Tile, columns)
		for x := 0; x < columns; x++ {
			tiles[y][x] = Tile{
				Texture: texture,
				Coords: Bounds{
					Left:   float32(x*columnWidth+offsetX) * sx,
					Right:  float32(x*columnWidth+columnWidth-1+offsetX) * sx,
					Top:    float32(y*rowHeight+offsetY) * sy,
					Bottom: float32(y*rowHeight+rowHeight-1+offsetY) * sy,
				},
			}
		}
	}
	return tiles
}

type ExtentTile struct {
	Tile
	Extent Bounds
}

func NewExtentTile(tile Tile, extent Bounds) ExtentTile {
	return ExtentTile{Tile: tile, Extent: extent}
}

type Vertex struct {
	X, Y     float32
	Coord    TexCoord
	HasCoord bool
	Color    Color
	HasColor bool
}

type VertexModifier func(*Vertex)

func (v *Vertex) AddColor(c Color) {
	if v.HasColor {
		v.Color = v.Color.Multiply(c)
	} else {
		v.Color = c
		v.HasColor = true
	}
}

func (v *Vertex) SetCoord(texture string, x, y float32) {
	v.Coord.Texture = texture
	v.Coord.X = x
	v.Coord.Y = y
	v.HasCoord = true
}

func (v Vertex) Lerp(to Vertex, delta float32) Vertex {
	return Vertex{
		X:        Lerp(v.X, to.X, delta),
		Y:        Lerp(v.Y, to.Y, delta),
		HasCoord: v.HasCoord && to.HasCoord,
		Coord: TexCoord{
			Coord: Coord{
				X: Lerp(v.Coord.X, to.Coord.X, delta),
				Y: Lerp(v.Coord.Y, to.Coord.Y, delta),
			},
			Texture: v.Coord.Texture,
		},
		HasColor: v.HasColor && to.HasColor,
		Color:    v.Color.Lerp(to.Color, delta),
	}
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
		Buffers: *buf.NewBuffers[Vertex](capacity, buffers, InitVertexBuffer),
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

func (b *VertexBuffer) AddIndexTriangle(i int) {
	b.AddIndex(i, i+1, i+2)
}

func (b *VertexBuffer) AddTriangle(v ...Vertex) {
	if b.clip.IsZero() {
		i := b.Add(v...)
		b.AddIndex(i, i+1, i+2)
	} else {
		inside := 0
		for i := 0; i < 3; i++ {
			if b.clip.Inside(v[i].X, v[i].Y) {
				inside++
			}
		}
		if inside == 3 {
			i := b.Add(v...)
			b.AddIndex(i, i+1, i+2)
		} else if inside > 0 {
			b.addClippedTriangle(v[0], v[1], v[2])
		}
	}
}

var relativeQuad = []int{0, 1, 2, 2, 3, 0}

func (b *VertexBuffer) AddQuad(v ...Vertex) {
	if b.clip.IsZero() {
		b.AddRelative(v, relativeQuad)
	} else {
		inside := 0
		for i := 0; i < 4; i++ {
			if b.clip.Inside(v[i].X, v[i].Y) {
				inside++
			}
		}
		if inside == 4 {
			b.AddRelative(v, relativeQuad)
		} else if inside > 0 {
			b.addClippedTriangle(v[0], v[1], v[2])
			b.addClippedTriangle(v[2], v[3], v[0])
		}
	}
}

func (buffer *VertexBuffer) addClippedTriangle(a, b, c Vertex) {

}

func (b *VertexBuffer) Clear() {
	b.Buffer.Clear()
	b.clip = Bounds{}
}

func (b *VertexBuffer) Clip() Bounds {
	return b.clip
}

type PrimitiveAdder interface {
	AddQuad(out *VertexBuffer, a Vertex, b Vertex, c Vertex, d Vertex)
	AddTriangle(out *VertexBuffer, a Vertex, b Vertex, c Vertex)
}

type PrimitiveAdderUnclipped struct{}

func (va PrimitiveAdderUnclipped) AddQuad(out *VertexBuffer, a, b, c, d Vertex) {
	i := out.Add(a, b, c, d)
	out.AddIndexQuad(i)
}
func (va PrimitiveAdderUnclipped) AddTriangle(out *VertexBuffer, a, b, c Vertex) {
	i := out.Add(a, b, c)
	out.AddIndexTriangle(i)
}

type PrimitiveAdderBoundsClipped struct {
	Bounds Bounds
}

func (va PrimitiveAdderBoundsClipped) AddQuad(out *VertexBuffer, a, b, c, d Vertex) {
	ain := va.Bounds.Inside(a.X, a.Y)
	bin := va.Bounds.Inside(b.X, b.Y)
	cin := va.Bounds.Inside(c.X, c.Y)
	din := va.Bounds.Inside(d.X, d.Y)
	if ain && bin && cin && din {
		i := out.Add(a, b, c, d)
		out.AddIndexQuad(i)
	} else if ain || bin || cin || din {
		va.addClippedTriangle(out, a, b, c, ain, bin, cin)
		va.addClippedTriangle(out, c, d, a, cin, din, ain)
	}
}
func (va PrimitiveAdderBoundsClipped) AddTriangle(out *VertexBuffer, a, b, c Vertex) {
	ain := va.Bounds.Inside(a.X, a.Y)
	bin := va.Bounds.Inside(b.X, b.Y)
	cin := va.Bounds.Inside(c.X, c.Y)
	if ain && bin && cin {
		i := out.Add(a, b, c)
		out.AddIndexTriangle(i)
	} else if ain || bin || cin {
		va.addClippedTriangle(out, a, b, c, ain, bin, cin)
	}
}

func (va PrimitiveAdderBoundsClipped) addClippedTriangle(out *VertexBuffer, a, b, c Vertex, ain, bin, cin bool) {

}

var clipPoints [6]ClippedPoint

func clipTriangle(bounds Bounds, a, b, c Vertex, out [12]Vertex) int {
	n := bounds.ClipTriangle(a.X, a.Y, b.X, b.Y, c.X, c.Y, clipPoints)
	switch n {
	case 0:
		return 0
	case 3:
		out[0] = a
		out[1] = b
		out[2] = c
	case 4:
		// 0,1,2 - 2,3,0
	case 5:
		// 0,1,2 - 2,3,0 - 3,4,0
	case 6:
		// stop
	}
	return n
}
