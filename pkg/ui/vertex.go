package ui

type UIVertex struct {
	X, Y     float32
	Coord    TexCoord
	HasCoord bool
	Color    Color
	HasColor bool
}

func (v *UIVertex) AddColor(r, g, b, a float32) {
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

func (v *UIVertex) SetCoord(texture string, x, y float32) {
	v.Coord.Texture = texture
	v.Coord.X = x
	v.Coord.Y = y
	v.HasCoord = true
}

type UIVertexBuffer struct {
	Data    []UIVertex
	Indices []int
}

func NewVertexBuffer(capacity int) *UIVertexBuffer {
	return &UIVertexBuffer{
		Data:    make([]UIVertex, 0, capacity),
		Indices: make([]int, 0, capacity*3/2),
	}
}

func (b UIVertexBuffer) Pos() int {
	return len(b.Data)
}
func (b *UIVertexBuffer) Add(v ...UIVertex) int {
	i := len(b.Data)
	b.Data = append(b.Data, v...)
	return i
}
func (b *UIVertexBuffer) AddIndices(index ...int) {
	b.Indices = append(b.Indices, index...)
}
func (b *UIVertexBuffer) Clear() {
	b.Data = b.Data[:0]
	b.Indices = b.Indices[:0]
}

// assuming the vertex at the given index is the top left corner and the remaining vertices are clockwise this adds two triangles that make a quad.
func (b *UIVertexBuffer) AddQuad(i int) {
	b.Indices = append(b.Indices, i, i+1, i+2, i+2, i+3, i)
}
