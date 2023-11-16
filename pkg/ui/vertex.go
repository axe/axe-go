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

func NewVertexIterator(iterable buf.Iterable[Vertex, VertexBuffer]) VertexIterator {
	return buf.NewDataIterator[Vertex, VertexBuffer](iterable)
}

type IndexIterator = buf.IndexIterator[Vertex, VertexBuffer]

func NewIndexIterator(iterable buf.Iterable[Vertex, VertexBuffer]) IndexIterator {
	return buf.NewIndexIterator[Vertex, VertexBuffer](iterable)
}

var (
	BufferPoolCreate = func() *VertexBuffers {
		return NewVertexBuffers(BufferCapacity, Buffers)
	}
	ClipPoolCreate = func() *VertexBuffers {
		return NewVertexBuffers(ClipCapacity, Buffers)
	}
	BufferPoolQueueCreate = func() *VertexQueue {
		return NewVertexQueue(BufferQueueCapacity)
	}
	BufferPoolReset = func(vbs *VertexBuffers) *VertexBuffers {
		vbs.Clear()
		return vbs
	}
	BufferPoolQueueReset = func(vbs *VertexQueue) *VertexQueue {
		vbs.Clear()
		return vbs
	}

	BufferPoolSize      = 256
	BufferCapacity      = 64
	ClipPoolSize        = 24
	ClipCapacity        = 1024
	BufferQueueCapacity = 12
	BufferQueuePoolSize = 256
	Buffers             = 1

	BufferPool      = ds.NewPool(BufferPoolSize, BufferPoolCreate, BufferPoolReset)
	BufferQueuePool = ds.NewPool(BufferQueuePoolSize, BufferPoolQueueCreate, BufferPoolQueueReset)

	ClipPool   = ds.NewPool(ClipPoolSize, ClipPoolCreate, BufferPoolReset)
	ClipMemory = ds.NewStack[*VertexBuffers](32)
)

func NewVertexBuffers(capacity int, buffers int) *VertexBuffers {
	vbs := &VertexBuffers{}
	vbs.Buffers = *buf.NewBuffers[Vertex](capacity, buffers, vbs.initBuffer, vbs.resetBuffer)
	return vbs
}

func NewVertexQueue(capacity int) *VertexQueue {
	return &VertexQueue{
		Queue: *buf.NewQueue[Vertex, VertexBuffer](capacity),
	}
}

type VertexIterable = buf.Iterable[Vertex, VertexBuffer]

type VertexQueue struct {
	buf.Queue[Vertex, VertexBuffer]
}

func (vbs *VertexQueue) Clip(bounds Bounds, render func(clippable *VertexQueue, clipping bool)) bool {
	if bounds.IsEmpty() {
		render(vbs, false)

		return false
	} else {
		clippable := BufferQueuePool.Get()
		render(clippable, true)
		clipped := ClipPool.Get()
		clipIterable(bounds, clippable, clipped)
		BufferQueuePool.Free(clippable)
		ClipMemory.Push(clipped)
		vbs.Add(clipped)

		return true
	}
}

func (vbs *VertexQueue) Clone(start, endExclusive int) {
	if start == endExclusive {
		return
	}
	buffers := vbs.GetBuffers()
	clones := ClipPool.Get()
	clones.Reserve(endExclusive - start)
	for i := start; i < endExclusive; i++ {
		buffers[i] = *buffers[i].CloneTo(clones.ReservedNext())
	}
	ClipMemory.Push(clones)
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
		clippable := BufferPool.Get()
		render(clippable)
		clipIterable(bounds, clippable, vbs)
		BufferPool.Free(clippable)
	}
}

func clipIterable(bounds Bounds, clippable VertexIterable, clipped *VertexBuffers) {
	childCount := clippable.Len()
	vb := clipped.Buffer()

	for bufferIndex := 0; bufferIndex < childCount; bufferIndex++ {
		child := clippable.At(bufferIndex)
		if !child.ClipCompatible(vb) {
			vb = clipped.Add()
			vb.Blend = child.Blend
		}
		child.Clip(bounds, vb)
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
		vbs.Add()
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
	added := vbs.Add()
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

func (b *VertexBuffer) CloneTo(out *VertexBuffer) *VertexBuffer {
	out.Buffer = b.Buffer.CloneTo(out.Buffer)
	out.Blend = b.Blend
	out.Primitive = b.Primitive
	return out
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

	clip := clipper{
		out:    out,
		bounds: bounds,
	}

	for i := 0; i < indexCount; i += 3 {
		a := indices.At(i)
		b := indices.At(i + 1)
		c := indices.At(i + 2)
		clip.addTriangle(*a, *b, *c)
	}
}

func (b *VertexBuffer) clipLines(bounds Bounds, out *VertexBuffer) {
	indices := b.IndexSpanAt(0)
	indexCount := indices.Len()

	out.ReserveLines(indexCount/2 + 1)

	for i := 0; i < indexCount; i += 2 {
		a := indices.At(i)
		b := indices.At(i + 1)
		line := bounds.ClipLine(a.X, a.Y, b.X, b.Y)

		if !line.Outside {
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

	out.ReserveTriangles(indexCount/4*3 + 1)

	clip := clipper{
		out:    out,
		bounds: bounds,
	}

	for i := 0; i < indexCount; i += 4 {
		a := indices.At(i)
		b := indices.At(i + 1)
		c := indices.At(i + 2)
		d := indices.At(i + 3)
		clip.addTriangle(*a, *b, *c)
		clip.addTriangle(*c, *d, *a)
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

func (b *VertexBuffer) AddQuad() []Vertex {
	b.ReserveQuads(1)
	return b.GetReservedQuad()
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

func (b *VertexBuffer) AddTriangle() []Vertex {
	b.ReserveTriangles(1)
	return b.GetReservedTriangle()
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

func (b *VertexBuffer) AddLine() []Vertex {
	b.ReserveLines(1)
	return b.GetReservedLine()
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

type clipper struct {
	bounds Bounds
	out    *VertexBuffer
	i      int
	tri    []Vertex
	first  Vertex
	last   Vertex
}

func (c *clipper) add(v Vertex) {
	if c.i == 0 {
		c.first = v
	}
	if c.i < 3 {
		c.tri[c.i] = v
	} else {
		c.tri = c.out.AddTriangle()
		c.tri[0] = c.last
		c.tri[1] = v
		c.tri[2] = c.first
	}
	c.last = v
	c.i++
}

func (c *clipper) addInterpolate(a, b Vertex, delta, x, y float32) {
	if delta == 0 {
		c.add(a)
	} else if delta == 1 {
		c.add(b)
	} else {
		c.add(a.LerpWith(b, delta, x, y))
	}
}

func (c *clipper) addTriangularInterpolate(v1, v2, v3 Vertex, p Coord) {
	// https://codeplea.com/triangular-interpolation
	dy23 := v2.Y - v3.Y
	dxp3 := p.X - v3.X
	dx32 := v3.X - v2.X
	dyp3 := p.Y - v3.Y
	dy31 := v3.Y - v1.Y
	dx13 := v1.X - v3.X
	dy13 := v1.Y - v3.Y
	weight0 := (dy23*dxp3 + dx32*dyp3) / (dy31*dxp3 + dx13*dy13)
	weight1 := (dy31*dxp3 + dx13*dyp3) / (dy23*dx13 + dx32*dy13)
	weight2 := 1 - weight0 - weight1
	c.add(v1.Scale(weight0).Add(v2.Scale(weight1)).Add(v3.Scale(weight2)))
}

func (c *clipper) addLine(line ClippedLine, a, b Vertex, only bool, other Vertex) {
	if !line.Outside {
		c.addInterpolate(a, b, line.StartDelta, line.Start.X, line.Start.Y)
		if line.EndDelta < 1 {
			c.addInterpolate(a, b, line.EndDelta, line.End.X, line.End.Y)
		}
		if only {
			clipped := c.bounds.ClipCoord(Coord{X: other.X, Y: other.Y})
			c.addTriangularInterpolate(a, b, other, clipped)
		}
	}
}

func (c *clipper) addTriangle(v1, v2, v3 Vertex) {
	line0 := c.bounds.ClipLine(v1.X, v1.Y, v2.X, v2.Y)
	line1 := c.bounds.ClipLine(v2.X, v2.Y, v3.X, v3.Y)
	line2 := c.bounds.ClipLine(v3.X, v3.Y, v1.X, v1.Y)

	if line0.Outside && line1.Outside && line2.Outside {
		return
	}

	c.i = 0
	c.tri = c.out.GetReservedTriangle()
	c.addLine(line0, v1, v2, line1.Outside && line2.Outside, v3)
	c.addLine(line1, v2, v3, line0.Outside && line2.Outside, v1)
	c.addLine(line2, v3, v1, line0.Outside && line1.Outside, v2)
}
