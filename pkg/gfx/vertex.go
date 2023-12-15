package gfx

import "github.com/axe/axe-go/pkg/data"

type DataType int

const (
	Float DataType = iota
	Double
	Byte
	Ubyte
	Int16
	Uint16
	Int32
	Uint32
)

func (d DataType) Bytes() int {
	switch d {
	case Double:
		return 8
	case Float, Int32, Uint32:
		return 4
	case Int16, Uint16:
		return 2
	case Byte, Ubyte:
		return 1
	}
	return 0
}

type AttributeType int

const (
	TypeVertex AttributeType = iota
	TypeColor
	TypeNormal
	TypeTexture
	TypeGeneric
)

type Frequency int

const (
	Never Frequency = iota
	Sometimes
	Always
)

type BufferAttribute struct {
	Name     string
	Data     DataType
	Size     int
	Type     AttributeType
	Offset   int
	Location int
}

func (ba *BufferAttribute) SetName(name string) {
	ba.Name = name
}

type BufferFormat struct {
	Attributes []BufferAttribute
	Stride     int
	Frequency  Frequency
	Draw, Read bool
}

func (f *BufferFormat) Add(data DataType, size int, attr AttributeType) *BufferAttribute {
	i := len(f.Attributes)

	f.Attributes = append(f.Attributes, BufferAttribute{
		Data:   data,
		Size:   size,
		Type:   attr,
		Offset: f.Stride,
	})
	f.Stride += data.Bytes() * size

	return &f.Attributes[i]
}

func (f *BufferFormat) AddVertex(data DataType, size int) *BufferAttribute {
	return f.Add(data, size, TypeVertex)
}

func (f *BufferFormat) AddColor(data DataType, size int) *BufferAttribute {
	return f.Add(data, size, TypeColor)
}

func (f *BufferFormat) AddNormal() *BufferAttribute {
	return f.Add(Float, 3, TypeNormal)
}

func (f *BufferFormat) AddTexture() *BufferAttribute {
	return f.Add(Float, 2, TypeTexture)
}

func (f *BufferFormat) AddGeneric(data DataType, size int) *BufferAttribute {
	return f.Add(data, size, TypeGeneric)
}

type Buffer struct {
	Format   *BufferFormat
	Data     data.Bytes
	Metadata any
}
