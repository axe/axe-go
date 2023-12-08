package data

import (
	"encoding/binary"
	"math"

	"github.com/axe/axe-go/pkg/util"
)

type Bytes struct {
	bytes  []byte
	offset int
	limit  int
	order  binary.ByteOrder
}

type ByteWritable interface {
	Write(b *Bytes)
}
type ByteSizable interface {
	Bytes() int
}
type ByteReadable interface {
	Read(b *Bytes)
}

func NewBytes(capacity int) *Bytes {
	return &Bytes{
		bytes: make([]byte, capacity),
		order: binary.NativeEndian,
	}
}

func (b *Bytes) Data() []byte {
	if b.limit == 0 {
		return b.bytes[0:b.offset]
	} else {
		return b.bytes[b.offset:b.limit]
	}
}

func (b *Bytes) Take(n int) []byte {
	taken := b.bytes[b.offset : b.offset+n]
	b.offset += n
	return taken
}

func (b *Bytes) Flip() {
	b.limit = b.offset
	b.offset = 0
}

func (b *Bytes) Limit() int {
	return b.limit
}

func (b *Bytes) Offset() int {
	return b.offset
}

func (b *Bytes) Reset(offset int) {
	b.offset = offset
}

func (b *Bytes) Clear() {
	b.limit = 0
	b.offset = 0
}

func (b *Bytes) Remaining() int {
	return b.limit - b.offset
}

func (b *Bytes) HasRemaining() bool {
	return b.offset < b.limit
}

// Ensures there are at least n bytes available in the buffer.
func (b *Bytes) Reserve(n int) {
	b.bytes = util.SliceEnsureSize(b.bytes, b.offset+n)
}

func (b Bytes) IsNative() bool {
	return b.order == binary.NativeEndian
}

func (b Bytes) IsLittle() bool {
	return b.order == binary.LittleEndian
}

func (b Bytes) IsBig() bool {
	return b.order == binary.BigEndian
}

func (b *Bytes) SetNative() {
	b.order = binary.NativeEndian
}

func (b *Bytes) SetLittle() {
	b.order = binary.LittleEndian
}

func (b *Bytes) SetBig() {
	b.order = binary.BigEndian
}

func (b *Bytes) putLen(len int) {
	b.PutCompressed(int64(len))
}

func (b *Bytes) getLen() int {
	return int(b.GetCompressed())
}

func (b *Bytes) Put(v byte) {
	b.Take(1)[0] = v
}

func (b *Bytes) Puts(v []byte) {
	putSlice(v, b.Put)
}

func (b *Bytes) PutDynamic(v []byte) {
	putSliceDynamic(v, b.Put, b.putLen)
}

func (b *Bytes) PutBool(v bool) {
	b.Take(1)[0] = util.If[byte](v, 1, 0)
}

func (b *Bytes) PutBools(v []bool) {
	putSlice(v, b.PutBool)
}

func (b *Bytes) PutBoolDynamic(v []bool) {
	putSliceDynamic(v, b.PutBool, b.putLen)
}

func (b *Bytes) PutInt8(v int8) {
	b.Take(1)[0] = byte(v)
}

func (b *Bytes) PutInt8s(v []int8) {
	putSlice(v, b.PutInt8)
}

func (b *Bytes) PutInt8Dynamic(v []int8) {
	putSliceDynamic(v, b.PutInt8, b.putLen)
}

func (b *Bytes) PutInt16(v int16) {
	b.order.PutUint16(b.Take(2), uint16(v))
}

func (b *Bytes) PutInt16s(v []int16) {
	putSlice(v, b.PutInt16)
}

func (b *Bytes) PutInt16Dynamic(v []int16) {
	putSliceDynamic(v, b.PutInt16, b.putLen)
}

func (b *Bytes) PutInt32(v int32) {
	b.order.PutUint32(b.Take(4), uint32(v))
}

func (b *Bytes) PutInt32s(v []int32) {
	putSlice(v, b.PutInt32)
}

func (b *Bytes) PutInt32Dynamic(v []int32) {
	putSliceDynamic(v, b.PutInt32, b.putLen)
}

func (b *Bytes) PutInt64(v int64) {
	b.order.PutUint64(b.Take(8), uint64(v))
}

func (b *Bytes) PutInt64s(v []int64) {
	putSlice(v, b.PutInt64)
}

func (b *Bytes) PutInt64Dynamic(v []int64) {
	putSliceDynamic(v, b.PutInt64, b.putLen)
}

func (b *Bytes) PutUint8(v uint8) {
	b.Take(1)[0] = byte(v)
}

func (b *Bytes) PutUint8s(v []uint8) {
	putSlice(v, b.PutUint8)
}

func (b *Bytes) PutUint8Dynamic(v []uint8) {
	putSliceDynamic(v, b.PutUint8, b.putLen)
}

func (b *Bytes) PutUint16(v uint16) {
	b.order.PutUint16(b.Take(2), v)
}

func (b *Bytes) PutUint16s(v []uint16) {
	putSlice(v, b.PutUint16)
}

func (b *Bytes) PutUint16Dynamic(v []uint16) {
	putSliceDynamic(v, b.PutUint16, b.putLen)
}

func (b *Bytes) PutUint32(v uint32) {
	b.order.PutUint32(b.Take(4), v)
}

func (b *Bytes) PutUint32s(v []uint32) {
	putSlice(v, b.PutUint32)
}

func (b *Bytes) PutUint32Dynamic(v []uint32) {
	putSliceDynamic(v, b.PutUint32, b.putLen)
}

func (b *Bytes) PutUint64(v uint64) {
	b.order.PutUint64(b.Take(8), v)
}

func (b *Bytes) PutUint64s(v []uint64) {
	putSlice(v, b.PutUint64)
}

func (b *Bytes) PutUint64Dynamic(v []uint64) {
	putSliceDynamic(v, b.PutUint64, b.putLen)
}

func (b *Bytes) PutFloat32(v float32) {
	b.order.PutUint32(b.Take(4), math.Float32bits(v))
}

func (b *Bytes) PutFloat32s(v []float32) {
	putSlice(v, b.PutFloat32)
}

func (b *Bytes) PutFloat32Dynamic(v []float32) {
	putSliceDynamic(v, b.PutFloat32, b.putLen)
}

func (b *Bytes) PutFloat64(v float64) {
	b.order.PutUint64(b.Take(8), math.Float64bits(v))
}

func (b *Bytes) PutFloat64s(v []float64) {
	putSlice(v, b.PutFloat64)
}

func (b *Bytes) PutFloat64Dynamic(v []float64) {
	putSliceDynamic(v, b.PutFloat64, b.putLen)
}

func (b *Bytes) Get() byte {
	return b.Take(1)[0]
}

func (b *Bytes) Gets(v []byte) {
	getSlice(v, b.Get)
}

func (b *Bytes) GetDynamic() []byte {
	return getSliceDynamic(b.Get, b.getLen)
}

func (b *Bytes) GetBool() bool {
	return b.Take(1)[0] != 0
}

func (b *Bytes) GetBools(v []bool) {
	getSlice(v, b.GetBool)
}

func (b *Bytes) GetBoolDynamic() []bool {
	return getSliceDynamic(b.GetBool, b.getLen)
}

func (b *Bytes) GetInt8() int8 {
	return int8(b.Take(1)[0])
}

func (b *Bytes) GetInt8s(v []int8) {
	getSlice(v, b.GetInt8)
}

func (b *Bytes) GetInt8Dynamic() []int8 {
	return getSliceDynamic(b.GetInt8, b.getLen)
}

func (b *Bytes) GetInt16() int16 {
	return int16(b.order.Uint16(b.Take(2)))
}

func (b *Bytes) GetInt16s(v []int16) {
	getSlice(v, b.GetInt16)
}

func (b *Bytes) GetInt16Dynamic() []int16 {
	return getSliceDynamic(b.GetInt16, b.getLen)
}

func (b *Bytes) GetInt32() int32 {
	return int32(b.order.Uint32(b.Take(4)))
}

func (b *Bytes) GetInt32s(v []int32) {
	getSlice(v, b.GetInt32)
}

func (b *Bytes) GetInt32Dynamic() []int32 {
	return getSliceDynamic(b.GetInt32, b.getLen)
}

func (b *Bytes) GetInt64() int64 {
	return int64(b.order.Uint64(b.Take(8)))
}

func (b *Bytes) GetInt64s(v []int64) {
	getSlice(v, b.GetInt64)
}

func (b *Bytes) GetInt64Dynamic() []int64 {
	return getSliceDynamic(b.GetInt64, b.getLen)
}

func (b *Bytes) GetUint8() uint8 {
	return uint8(b.Take(1)[0])
}

func (b *Bytes) GetUint8s(v []uint8) {
	getSlice(v, b.GetUint8)
}

func (b *Bytes) GetUint8Dynamic() []uint8 {
	return getSliceDynamic(b.GetUint8, b.getLen)
}

func (b *Bytes) GetUint16() uint16 {
	return uint16(b.order.Uint16(b.Take(2)))
}

func (b *Bytes) GetUint16s(v []uint16) {
	getSlice(v, b.GetUint16)
}

func (b *Bytes) GetUint16Dynamic() []uint16 {
	return getSliceDynamic(b.GetUint16, b.getLen)
}

func (b *Bytes) GetUint32() uint32 {
	return b.order.Uint32(b.Take(4))
}

func (b *Bytes) GetUint32s(v []uint32) {
	getSlice(v, b.GetUint32)
}

func (b *Bytes) GetUint32Dynamic() []uint32 {
	return getSliceDynamic(b.GetUint32, b.getLen)
}

func (b *Bytes) GetUint64() uint64 {
	return b.order.Uint64(b.Take(8))
}

func (b *Bytes) GetUint64s(v []uint64) {
	getSlice(v, b.GetUint64)
}

func (b *Bytes) GetUint64Dynamic() []uint64 {
	return getSliceDynamic(b.GetUint64, b.getLen)
}

func (b *Bytes) GetFloat32() float32 {
	return math.Float32frombits(b.order.Uint32(b.Take(4)))
}

func (b *Bytes) GetFloat32s(v []float32) {
	getSlice(v, b.GetFloat32)
}

func (b *Bytes) GetFloat32Dynamic() []float32 {
	return getSliceDynamic(b.GetFloat32, b.getLen)
}

func (b *Bytes) GetFloat64() float64 {
	return math.Float64frombits(b.order.Uint64(b.Take(8)))
}

func (b *Bytes) GetFloat64s(v []float64) {
	getSlice(v, b.GetFloat64)
}

func (b *Bytes) GetFloat64Dynamic() []float64 {
	return getSliceDynamic(b.GetFloat64, b.getLen)
}

const (
	compressedShift = 7
	compressedMask  = 0x7F
	compressedMore  = 0x80
)

func (b *Bytes) GetCompressed() int64 {
	v := int64(0)
	byteValue := byte(compressedMore)
	shift := 0

	for (byteValue & compressedMore) == compressedMore {
		byteValue = b.Get()
		v |= (int64(byteValue&compressedMask) << shift)
		shift += compressedShift
	}

	return v
}

func (b *Bytes) GetCompresseds(v []int64) {
	getSlice(v, b.GetCompressed)
}

func (b *Bytes) PutCompressed(v int64) int {
	bytes := 0

	for {
		byteValue := v & compressedMask
		if v > compressedMask {
			byteValue |= compressedMore
		}
		b.Put(byte(byteValue))
		v >>= compressedShift
		bytes++
		if v == 0 {
			break
		}
	}

	return bytes
}

func (b *Bytes) PutCompresseds(v []int64) {
	putSlice(v, func(i int64) {
		b.PutCompressed(i)
	})
}

func BytesCompressed(v uint64) int {
	bytes := 0
	for {
		v >>= compressedShift
		bytes++
		if v == 0 {
			break
		}
	}

	return bytes
}

func (b *Bytes) Write(writable ByteWritable) {
	if size, ok := writable.(ByteSizable); ok {
		b.Reserve(size.Bytes())
	}
	writable.Write(b)
}

func Read[B ByteReadable](b *Bytes) B {
	var instance B
	instance.Read(b)
	return instance
}

func Reads[B ByteReadable](b *Bytes, count int) []B {
	instances := make([]B, count)
	for i := 0; i < count; i++ {
		instances[i].Read(b)
	}
	return instances
}

func ReadDynamic[B ByteReadable](b *Bytes) []B {
	return Reads[B](b, b.getLen())
}

func Write[B ByteWritable](b *Bytes, writable B) {
	b.Write(writable)
}

func Writes[B ByteWritable](b *Bytes, writables []B) {
	n := len(writables)
	if n == 0 {
		return
	}
	reserve := 0
	for _, w := range writables {
		if size, ok := any(w).(ByteSizable); ok {
			reserve += size.Bytes()
		}
	}

	b.Reserve(reserve)
	for _, w := range writables {
		w.Write(b)
	}
}

func WriteDynamic[B ByteWritable](b *Bytes, writables []B) {
	b.putLen(len(writables))
	Writes(b, writables)
}

func putSlice[V any](slice []V, put func(V)) {
	for i := range slice {
		put(slice[i])
	}
}

func putSliceDynamic[V any](slice []V, put func(V), putLen func(int)) {
	putLen(len(slice))
	putSlice(slice, put)
}

func getSlice[V any](slice []V, get func() V) {
	for i := range slice {
		slice[i] = get()
	}
}

func getSliceDynamic[V any](get func() V, getLen func() int) []V {
	n := getLen()
	slice := make([]V, n)
	for i := range slice {
		slice[i] = get()
	}
	return slice
}
