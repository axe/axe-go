package ui

import (
	"github.com/axe/axe-go/pkg/buf"
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/util"
)

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

type TexCoord struct {
	Coord
	Texture string
}

type Tile struct {
	Coords  Bounds
	Texture string
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

func (t Tile) Coord(dx, dy float32) TexCoord {
	return TexCoord{
		Texture: t.Texture,
		Coord: Coord{
			X: util.Lerp(t.Coords.Left, t.Coords.Right, dx),
			Y: util.Lerp(t.Coords.Top, t.Coords.Bottom, dy),
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
	Tex      TexCoord
	HasCoord bool
	Color    Color
	HasColor bool
}

func (v *Vertex) AddColor(c Color) {
	if v.HasColor {
		v.Color = v.Color.Multiply(c)
	} else {
		v.Color = c
		v.HasColor = true
	}
}

func (v *Vertex) SetCoord(texture string, x, y float32) {
	v.Tex.Texture = texture
	v.Tex.X = x
	v.Tex.Y = y
	v.HasCoord = true
}

func (v Vertex) Lerp(to Vertex, delta float32) Vertex {
	return v.LerpWith(to, delta, util.Lerp(v.X, to.X, delta), util.Lerp(v.Y, to.Y, delta))
}

func (v Vertex) LerpWith(to Vertex, delta, x, y float32) Vertex {
	return Vertex{
		X:        x,
		Y:        y,
		HasCoord: v.HasCoord && to.HasCoord,
		Tex: TexCoord{
			Coord:   v.Tex.Lerp(to.Tex.Coord, delta),
			Texture: v.Tex.Texture,
		},
		HasColor: v.HasColor && to.HasColor,
		Color:    v.Color.Lerp(to.Color, delta),
	}
}

func (v Vertex) Scale(s float32) Vertex {
	copy := v
	copy.X *= s
	copy.Y *= s
	copy.Tex.X *= s
	copy.Tex.Y *= s
	copy.Color.R *= s
	copy.Color.G *= s
	copy.Color.B *= s
	copy.Color.A *= s
	return copy
}

func (v Vertex) Add(b Vertex) Vertex {
	copy := v
	copy.X += b.X
	copy.Y += b.Y
	copy.Tex.X += b.Tex.X
	copy.Tex.Y += b.Tex.Y
	copy.Color.R += b.Color.R
	copy.Color.G += b.Color.G
	copy.Color.B += b.Color.B
	copy.Color.A += b.Color.A
	return copy
}

type VertexModifier func(*Vertex)

type VertexIterator = buf.DataIterator[Vertex, VertexBuffer]

func NewVertexIterator(vb *VertexBuffers) VertexIterator {
	return buf.NewDataIterator(&vb.Buffers)
}

type IndexIterator = buf.IndexIterator[Vertex, VertexBuffer]

func NewIndexIterator(vb *VertexBuffers) IndexIterator {
	return buf.NewIndexIterator(&vb.Buffers)
}

var (
	BufferPoolCreate = func() *VertexBuffers {
		return NewVertexBuffers(BufferCapacity, Buffers)
	}
	BufferPoolQueueCreate = func() *VertexBuffers {
		return NewVertexBuffers(0, 0)
	}
	BufferPoolReset = func(vbs *VertexBuffers) *VertexBuffers {
		vbs.Clear()
		return vbs
	}

	BufferPoolSize = 256
	BufferCapacity = 64
	Buffers        = 1

	BufferPool      = ds.NewPool(BufferPoolSize, BufferPoolCreate, BufferPoolReset)
	BufferQueuePool = ds.NewPool(BufferPoolSize, BufferPoolQueueCreate, BufferPoolReset)
)

func NewVertexBuffers(capacity int, buffers int) *VertexBuffers {
	vbs := &VertexBuffers{}
	vbs.Buffers = *buf.NewBuffers[Vertex](capacity, buffers, vbs.initBuffer, vbs.resetBuffer)
	return vbs
}

type VertexBuffers struct {
	buf.Buffers[Vertex, VertexBuffer]

	Blend     Blend
	Primitive Primitive
}

func (vbs *VertexBuffers) initBuffer(vb *VertexBuffer, capacity int) {
	vb.Init(capacity)
	vb.Blend = vbs.Blend
	vb.Primitive = vbs.Primitive
}

func (vbs *VertexBuffers) resetBuffer(vb *VertexBuffer, pos buf.Position) {
	vb.Reset(pos)
	vb.Blend = vbs.Blend
	vb.Primitive = vbs.Primitive
}

func (vbs *VertexBuffers) Clip(bounds Bounds, render func(clippable *VertexBuffers)) {
	if bounds.IsEmpty() {
		render(vbs)
	} else {
		children := vbs.NewLike()
		render(children)
		childCount := children.Len()
		out := vbs.Buffer()

		for bufferIndex := 0; bufferIndex < childCount; bufferIndex++ {
			child := children.At(bufferIndex)
			if !child.ClipCompatible(out) {
				out = vbs.AddBuffer()
				out.Blend = child.Blend
			}
			child.Clip(bounds, out)
		}

		BufferPool.Free(children)
	}
}

func (vbs *VertexBuffers) With(primitive Primitive, blend Blend, render func(out *VertexBuffers)) {
	current := vbs.Buffer()
	currentPrimitive := vbs.Primitive
	currentBlend := vbs.Blend
	if current.Primitive == primitive && current.Blend == blend {
		render(vbs)
	} else {
		vbs.Blend = blend
		vbs.Primitive = primitive
		vbs.AddBuffer()
		render(vbs)
		vbs.Blend = currentBlend
		vbs.Primitive = currentPrimitive
	}
}

func (vbs *VertexBuffers) Get(primitive Primitive, blend Blend) *VertexBuffer {
	current := vbs.Buffer()
	if current.Primitive == primitive && current.Blend == blend {
		return current
	}
	added := vbs.AddBuffer()
	added.Blend = blend
	added.Primitive = primitive
	return added
}

func (vbs *VertexBuffers) NewLike() *VertexBuffers {
	like := BufferPool.Get()
	like.Blend = vbs.Blend
	like.Primitive = vbs.Primitive
	return like
}

type VertexBuffer struct {
	buf.Buffer[Vertex]

	Blend     Blend
	Primitive Primitive
}

func (b *VertexBuffer) Clip(bounds Bounds, out *VertexBuffer) {
	switch b.Primitive {
	case PrimitiveTriangle:
		b.clipTriangles(bounds, out)
	case PrimitiveLine:
		b.clipLines(bounds, out)
	case PrimitiveQuad:
		b.clipQuads(bounds, out)
	}
}

func (vb *VertexBuffer) clipTriangles(bounds Bounds, out *VertexBuffer) {
	indices := vb.IndexSpanAt(0)
	indexCount := indices.Len()

	out.ReserveTriangles(indexCount / 3)

	for i := 0; i < indexCount; i += 3 {
		a := indices.At(i)
		b := indices.At(i + 1)
		c := indices.At(i + 2)
		clipTriangle(bounds, out, *a, *b, *c)
	}
}

func (b *VertexBuffer) clipLines(bounds Bounds, out *VertexBuffer) {
	indices := b.IndexSpanAt(0)
	indexCount := indices.Len()

	out.ReserveLines(indexCount / 2)

	for i := 0; i < indexCount; i += 2 {
		a := indices.At(i)
		b := indices.At(i + 1)
		line := bounds.ClipLine(a.X, a.Y, b.X, b.Y)

		if line.Inside {
			if line.StartDelta == 0 && line.EndDelta == 1 {
				out.AddReservedLine(*a, *b)
			} else {
				v0 := a.Lerp(*b, line.StartDelta)
				v1 := a.Lerp(*b, line.EndDelta)
				out.AddReservedLine(v0, v1)
			}
		}
	}
}

func (vb *VertexBuffer) clipQuads(bounds Bounds, out *VertexBuffer) {
	indices := vb.IndexSpanAt(0)
	indexCount := indices.Len()

	out.ReserveTriangles(indexCount / 4 * 3)

	for i := 0; i < indexCount; i += 4 {
		a := indices.At(i)
		b := indices.At(i + 1)
		c := indices.At(i + 2)
		d := indices.At(i + 3)
		clipTriangle(bounds, out, *a, *b, *c)
		clipTriangle(bounds, out, *c, *d, *a)
	}
}

func (b *VertexBuffer) Init(capacity int) {
	b.Buffer.Init(capacity)
	b.Blend = BlendAlpha
	b.Primitive = PrimitiveTriangle
}

func (b *VertexBuffer) Clear() {
	b.Buffer.Clear()
	b.Blend = BlendAlpha
	b.Primitive = PrimitiveTriangle
}

func (b *VertexBuffer) Compatible(vb *VertexBuffer) bool {
	return b.Blend == vb.Blend && b.Primitive == vb.Primitive
}

func (b *VertexBuffer) ClipCompatible(vb *VertexBuffer) bool {
	return b.Blend == vb.Blend
}

func (b VertexBuffer) Empty() bool {
	return b.Buffer.Empty()
}

func (b VertexBuffer) Remaining() int {
	return b.Buffer.Remaining()
}

func (b *VertexBuffer) ReserveQuads(quads int) {
	switch b.Primitive {
	case PrimitiveQuad:
		b.Reserve(quads*4, quads*4)
	case PrimitiveTriangle:
		b.Reserve(quads*4, quads*6)
	case PrimitiveLine:
		b.Reserve(quads*4, quads*8)
	}
}

func (b *VertexBuffer) ReserveTriangles(triangles int) {
	switch b.Primitive {
	case PrimitiveQuad:
		b.Reserve(triangles*3, triangles*4)
	case PrimitiveTriangle:
		b.Reserve(triangles*3, triangles*3)
	case PrimitiveLine:
		b.Reserve(triangles*3, triangles*6)
	}
}

func (b *VertexBuffer) ReserveLines(lines int) {
	switch b.Primitive {
	case PrimitiveQuad:
		b.Reserve(lines*2, lines*4)
	case PrimitiveTriangle:
		b.Reserve(lines*2, lines*3)
	case PrimitiveLine:
		b.Reserve(lines*2, lines*2)
	}
}

func (b *VertexBuffer) AddReservedQuadSlice(quad []Vertex) {
	b.AddReservedQuad(quad[0], quad[1], quad[2], quad[3])
}

func (b *VertexBuffer) GetReservedQuad() (data []Vertex) {
	var index []int
	var dataIndex int
	switch b.Primitive {
	case PrimitiveQuad:
		dataIndex, data, index = b.Reserved(4, 4)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 2
		index[3] = dataIndex + 3
	case PrimitiveTriangle:
		dataIndex, data, index = b.Reserved(4, 6)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 2
		index[3] = dataIndex + 2
		index[4] = dataIndex + 3
		index[5] = dataIndex
	case PrimitiveLine:
		dataIndex, data, index = b.Reserved(4, 8)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 1
		index[3] = dataIndex + 2
		index[4] = dataIndex + 2
		index[5] = dataIndex + 3
		index[6] = dataIndex + 3
		index[7] = dataIndex + 0
	}
	return
}

func (b *VertexBuffer) AddReservedQuad(v0, v1, v2, v3 Vertex) {
	quad := b.GetReservedQuad()
	quad[0] = v0
	quad[1] = v1
	quad[2] = v2
	quad[3] = v3
}

func (b *VertexBuffer) GetReservedTriangle() (data []Vertex) {
	var index []int
	var dataIndex int
	switch b.Primitive {
	case PrimitiveQuad:
		dataIndex, data, index = b.Reserved(3, 4)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 2
		index[3] = dataIndex + 2
	case PrimitiveTriangle:
		dataIndex, data, index = b.Reserved(3, 3)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 2
	case PrimitiveLine:
		dataIndex, data, index = b.Reserved(3, 6)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 1
		index[3] = dataIndex + 2
		index[4] = dataIndex + 2
		index[5] = dataIndex + 0
	}
	return
}

func (b *VertexBuffer) AddReservedTriangle(v0, v1, v2 Vertex) {
	tri := b.GetReservedTriangle()
	tri[0] = v0
	tri[1] = v1
	tri[2] = v2
}

func (b *VertexBuffer) GetReservedLine() (data []Vertex) {
	var index []int
	var dataIndex int
	switch b.Primitive {
	case PrimitiveQuad:
		dataIndex, data, index = b.Reserved(2, 4)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 1
		index[3] = dataIndex + 1
	case PrimitiveTriangle:
		dataIndex, data, index = b.Reserved(2, 3)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 1
	case PrimitiveLine:
		dataIndex, data, index = b.Reserved(2, 2)
		index[0] = dataIndex
		index[1] = dataIndex + 1
	}
	return
}

func (b *VertexBuffer) AddReservedLine(v0, v1 Vertex) {
	line := b.GetReservedLine()
	line[0] = v0
	line[1] = v1
}

func clipTriangle(bounds Bounds, out *VertexBuffer, a, b, c Vertex) {
	line0 := bounds.ClipLine(a.X, a.Y, b.X, b.Y)
	line1 := bounds.ClipLine(b.X, b.Y, c.X, c.Y)
	line2 := bounds.ClipLine(c.X, c.Y, a.X, a.Y)

	if !line0.Inside && !line1.Inside && !line2.Inside {
		return
	}

	tri := out.GetReservedTriangle()
	var last, first Vertex
	i := 0

	add := func(v Vertex) {
		if i == 0 {
			first = v
		}
		if i < 3 {
			tri[i] = v
		} else {
			out.ReserveTriangles(1)
			tri = out.GetReservedTriangle()
			tri[0] = last
			tri[1] = v
			tri[2] = first
		}
		last = v
		i++
	}
	addInterpolate := func(a, b Vertex, delta, x, y float32) {
		if delta == 0 {
			add(a)
		} else if delta == 1 {
			add(b)
		} else {
			add(a.LerpWith(b, delta, x, y))
		}
	}
	addLine := func(line ClippedLine, a, b Vertex, only bool, c Vertex) {
		if line.Inside {
			addInterpolate(a, b, line.StartDelta, line.Start.X, line.Start.Y)
			if line.EndDelta < 1 {
				addInterpolate(a, b, line.EndDelta, line.End.X, line.End.Y)
			}
			if only {
				// https://codeplea.com/triangular-interpolation
				clipped := bounds.ClipCoord(Coord{X: c.X, Y: c.Y})
				dy23 := b.Y - c.Y
				dxp3 := clipped.X - c.X
				dx32 := c.X - b.X
				dyp3 := clipped.Y - c.Y
				dy31 := c.Y - a.Y
				dx13 := a.X - c.X
				dy13 := a.Y - c.Y
				weight0 := (dy23*dxp3 + dx32*dyp3) / (dy31*dxp3 + dx13*dy13)
				weight1 := (dy31*dxp3 + dx13*dyp3) / (dy23*dx13 + dx32*dy13)
				weight2 := 1 - weight0 - weight1
				add(a.Scale(weight0).Add(b.Scale(weight1)).Add(c.Scale(weight2)))
			}
		}
	}

	addLine(line0, a, b, !line1.Inside && !line2.Inside, c)
	addLine(line1, b, c, !line0.Inside && !line2.Inside, a)
	addLine(line2, c, a, !line0.Inside && !line1.Inside, b)
}
