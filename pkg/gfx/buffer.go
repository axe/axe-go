package gfx

import "github.com/axe/axe-go/pkg/data"

type DataType int

const (
	Float DataType = iota
	Double
	Byte
	UByte
	Short
	UShort
	Int
	UInt
)

func (d DataType) Bytes() int {
	switch d {
	case Double:
		return 8
	case Float, Int, UInt:
		return 4
	case Short, UShort:
		return 2
	case Byte, UByte:
		return 1
	}
	return 0
}

func (d DataType) Write(out *data.Bytes, v float32) {
	switch d {
	case Double:
		out.PutFloat64(float64(v))
	case Float:
		out.PutFloat32(v)
	case Int:
		out.PutInt32(int32(v))
	case UInt:
		out.PutUint32(uint32(v))
	case Short:
		out.PutInt16(int16(v))
	case UShort:
		out.PutUint16(uint16(v))
	case Byte:
		out.PutInt8(int8(v))
	case UByte:
		out.PutUint8(uint8(v))
	}
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
	ByType     [5]*BufferAttribute
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

	a := &f.Attributes[i]
	f.ByType[a.Type] = a

	return a
}

func (f *BufferFormat) Get(attr AttributeType) *BufferAttribute {
	return f.ByType[attr]
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

func (f *BufferFormat) Create(capacity int) *Buffer {
	return &Buffer{
		Format: f,
		Data:   data.NewBytes(capacity),
	}
}

type Buffer struct {
	Format   *BufferFormat
	Data     *data.Bytes
	Metadata any
}
