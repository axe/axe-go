package ui

import (
	"github.com/axe/axe-go/pkg/buf"
	"github.com/axe/axe-go/pkg/color"
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/util"
)

type ExtentTile struct {
	gfx.Tile
	Extent Bounds
}

func NewExtentTile(tile gfx.Tile, extent Bounds) ExtentTile {
	return ExtentTile{Tile: tile, Extent: extent}
}

type Vertex struct {
	X, Y     float32
	Tex      gfx.TextureCoord
	HasCoord bool
	Color    color.Color
	HasColor bool
}

func (v *Vertex) AddColor(c color.Color) {
	if v.HasColor {
		v.Color = v.Color.Multiply(c)
	} else {
		v.Color = c
		v.HasColor = true
	}
}

func (v *Vertex) InitColor() {
	if !v.HasColor {
		v.HasColor = true
		v.Color = color.White
	}
}

func (vert *Vertex) SetCoord(coord gfx.TextureCoord) {
	vert.Tex = coord
	vert.HasCoord = true
}

func (from Vertex) Lerp(to Vertex, delta float32) Vertex {
	return from.LerpWith(to, delta, util.Lerp(from.X, to.X, delta), util.Lerp(from.Y, to.Y, delta))
}

func (from Vertex) LerpWith(to Vertex, delta, x, y float32) Vertex {
	return Vertex{
		X:        x,
		Y:        y,
		HasCoord: from.HasCoord && to.HasCoord,
		Tex:      from.Tex.Lerp(to.Tex, delta),
		HasColor: from.HasColor && to.HasColor,
		Color:    from.Color.Lerp(to.Color, delta),
	}
}

type VertexModifier func(*Vertex)

type VertexIterator = buf.DataIterator[Vertex, VertexBuffer]

func NewVertexIterator(iterable buf.Iterable[Vertex, VertexBuffer], beginning bool) VertexIterator {
	return buf.NewDataIterator[Vertex, VertexBuffer](iterable, beginning)
}

type IndexIterator = buf.IndexIterator[Vertex, VertexBuffer]

func NewIndexIterator(iterable buf.Iterable[Vertex, VertexBuffer], beginning bool) IndexIterator {
	return buf.NewIndexIterator[Vertex, VertexBuffer](iterable, beginning)
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

func (vbs *VertexQueue) Clip(bounds Bounds, clipOut **VertexBuffers, render func(clippable *VertexQueue, clipping bool)) {
	if bounds.IsEmpty() {
		if *clipOut != nil {
			ClipPool.Free(*clipOut)
			*clipOut = nil
		}

		render(vbs, false)
	} else {
		clippable := BufferQueuePool.Get()

		render(clippable, true)

		if *clipOut == nil {
			*clipOut = ClipPool.Get()
		} else {
			(*clipOut).Clear()
		}

		(*clipOut).ClipInto(bounds, clippable)

		BufferQueuePool.Free(clippable)
	}
}

func (vbs *VertexQueue) ToBuffers() *VertexBuffers {
	if vbs == nil {
		return nil
	}
	return &VertexBuffers{Buffers: vbs.Queue.ToBuffers()}
}

type VertexBuffers struct {
	buf.Buffers[Vertex, VertexBuffer]

	Blend     gfx.Blend
	Primitive gfx.Primitive
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
		vbs.ClipInto(bounds, clippable)
		BufferPool.Free(clippable)
	}
}

func (vbs *VertexBuffers) With(primitive gfx.Primitive, blend gfx.Blend, render func(out *VertexBuffers)) {
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

func (vbs *VertexBuffers) Get(primitive gfx.Primitive, blend gfx.Blend) *VertexBuffer {
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

func (clipped *VertexBuffers) ClipInto(bounds Bounds, source VertexIterable) {
	childCount := source.Len()
	vb := clipped.Buffer()

	for bufferIndex := 0; bufferIndex < childCount; bufferIndex++ {
		child := source.At(bufferIndex)
		if !child.ClipCompatible(vb) {
			vb = clipped.Add()
			vb.Blend = child.Blend
		}

		child.Clip(bounds, vb)
	}
}

func (vbs *VertexBuffers) CloneTo(target *VertexBuffers) {
	n := vbs.Len()
	target.Clear()
	target.Reserve(n)
	for i := 0; i < n; i++ {
		vbs.At(i).CloneTo(target.ReservedNext())
	}
}

type VertexBuffer struct {
	buf.Buffer[Vertex]

	Blend     gfx.Blend
	Primitive gfx.Primitive
}

func (b *VertexBuffer) CloneTo(out *VertexBuffer) *VertexBuffer {
	out.Buffer = b.Buffer.CloneTo(out.Buffer)
	out.Blend = b.Blend
	out.Primitive = b.Primitive
	return out
}

func (b *VertexBuffer) Clip(bounds Bounds, out *VertexBuffer) {
	switch b.Primitive {
	case gfx.PrimitiveTriangle:
		b.clipTriangles(bounds, out)
	case gfx.PrimitiveLine:
		b.clipLines(bounds, out)
	case gfx.PrimitiveQuad:
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
			if line.Inside() {
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
	b.Blend = gfx.BlendAlpha
	b.Primitive = gfx.PrimitiveTriangle
}

func (b *VertexBuffer) Clear() {
	b.Buffer.Clear()
	b.Blend = gfx.BlendAlpha
	b.Primitive = gfx.PrimitiveTriangle
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
	case gfx.PrimitiveQuad:
		b.Reserve(quads*4, quads*4)
	case gfx.PrimitiveTriangle:
		b.Reserve(quads*4, quads*6)
	case gfx.PrimitiveLine:
		b.Reserve(quads*4, quads*8)
	}
}

func (b *VertexBuffer) ReserveTriangles(triangles int) {
	switch b.Primitive {
	case gfx.PrimitiveQuad:
		b.Reserve(triangles*3, triangles*4)
	case gfx.PrimitiveTriangle:
		b.Reserve(triangles*3, triangles*3)
	case gfx.PrimitiveLine:
		b.Reserve(triangles*3, triangles*6)
	}
}

func (b *VertexBuffer) ReserveLines(lines int) {
	switch b.Primitive {
	case gfx.PrimitiveQuad:
		b.Reserve(lines*2, lines*4)
	case gfx.PrimitiveTriangle:
		b.Reserve(lines*2, lines*3)
	case gfx.PrimitiveLine:
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
	case gfx.PrimitiveQuad:
		dataIndex, data, index = b.Reserved(4, 4)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 2
		index[3] = dataIndex + 3
	case gfx.PrimitiveTriangle:
		dataIndex, data, index = b.Reserved(4, 6)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 2
		index[3] = dataIndex + 2
		index[4] = dataIndex + 3
		index[5] = dataIndex
	case gfx.PrimitiveLine:
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
	case gfx.PrimitiveQuad:
		dataIndex, data, index = b.Reserved(3, 4)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 2
		index[3] = dataIndex + 2
	case gfx.PrimitiveTriangle:
		dataIndex, data, index = b.Reserved(3, 3)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 2
	case gfx.PrimitiveLine:
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
	case gfx.PrimitiveQuad:
		dataIndex, data, index = b.Reserved(2, 4)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 1
		index[3] = dataIndex + 1
	case gfx.PrimitiveTriangle:
		dataIndex, data, index = b.Reserved(2, 3)
		index[0] = dataIndex
		index[1] = dataIndex + 1
		index[2] = dataIndex + 1
	case gfx.PrimitiveLine:
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
}

var (
	tempvs0 = make([]Vertex, 12)
	tempvs1 = make([]Vertex, 12)
)

func (c *clipper) addTriangle(v1, v2, v3 Vertex) {
	side0 := c.bounds.Side(v1.X, v1.Y)
	side1 := c.bounds.Side(v2.X, v2.Y)
	side2 := c.bounds.Side(v3.X, v3.Y)
	sideAll := side0 | side1 | side2

	// If all are inside, no clipping necessary
	if sideAll == 0 {
		tri := c.out.GetReservedTriangle()
		tri[0] = v1
		tri[1] = v2
		tri[2] = v3
		return
	}

	// If all lines are outside in the same quadrant then we can
	// discard this triangle. We can't just exclude it if all lines
	// are outside because they might form a large triangle around
	// the bounds.
	sideCommon := side0 & side1 & side2
	if sideCommon != 0 {
		return
	}

	// Simplified "polygon" (bounds) & polygon clipping:
	// https://gist.github.com/alenaksu/89105882bb106b228a0e850a00becba7?ref=gorillasun.de
	result, points := tempvs0, tempvs1
	resultCount := 3
	result[0] = v1
	result[1] = v2
	result[2] = v3
	for sideIndex := 0; sideIndex < 4; sideIndex++ {
		side := BoundsSide(1 << sideIndex)

		points, result = result, points
		pointsCount := resultCount
		resultCount = 0

		p0 := points[pointsCount-1]
		for i := 0; i < pointsCount; i++ {
			p1 := points[i]
			if c.bounds.SideInside(p1.X, p1.Y, side) {
				if !c.bounds.SideInside(p0.X, p0.Y, side) {
					x, y, d := c.bounds.SideIntersect(p0.X, p0.Y, p1.X, p1.Y, side)
					result[resultCount] = p0.LerpWith(p1, d, x, y)
					resultCount++
				}
				result[resultCount] = p1
				resultCount++
			} else if c.bounds.SideInside(p0.X, p0.Y, side) {
				x, y, d := c.bounds.SideIntersect(p0.X, p0.Y, p1.X, p1.Y, side)
				result[resultCount] = p0.LerpWith(p1, d, x, y)
				resultCount++
			}
			p0 = p1
		}
		if resultCount == 0 {
			break
		}
	}

	if resultCount >= 3 {
		tri := c.out.GetReservedTriangle()
		tri[0] = result[0]
		tri[1] = result[1]
		tri[2] = result[2]
		c.out.ReserveTriangles(resultCount - 3)
		for i := 3; i < resultCount; i++ {
			tri = c.out.GetReservedTriangle()
			tri[0] = result[i-1]
			tri[1] = result[i]
			tri[2] = result[0]
		}
	}
}
