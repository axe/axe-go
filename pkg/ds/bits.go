package ds

import (
	"math/bits"
)

const (
	bitsShift uint32 = 6
	bitsAnd   uint32 = 63
	bitsOne   uint64 = uint64(1)
)

type index interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type IBits[I index] interface {
	Iterable[I]

	Max() I
	Set(index I, on bool)
	Get(index I) bool
	Ons() uint32
	IsEmpty() bool
	FirstOn() int32
	TakeFirst() int32
	LastOn() int32
	TakeLast() int32
	OnAfter(index int32) int32
	OffAfter(index int32) int32
	Clear()
	Clone() IBits[I]
}

type Bits[I index] struct {
	values []uint64
	ons    uint32
}

func NewBits[I index](max uint32) Bits[I] {
	return Bits[I]{
		values: make([]uint64, (max>>bitsShift)+1),
		ons:    0,
	}
}

func (b Bits[I]) Max() I {
	return I((uint32(len(b.values)) << bitsShift) - 1)
}

func (b *Bits[I]) EnsureMax(index I) {
	if b.Max() < index {
		len := (index >> bitsShift) + 1
		values := make([]uint64, len)
		copy(values, b.values)
		b.values = values
	}
}

func (b *Bits[I]) SafeSet(index I, on bool) {
	b.EnsureMax(index)
	b.Set(index, on)
}

func (b *Bits[I]) Set(index I, on bool) {
	v := index >> bitsShift
	vi := uint32(index) & bitsAnd
	vbits := bitsOne << vi
	isOn := b.values[v]&vbits != 0
	if isOn != on {
		if on {
			b.values[v] |= vbits
			b.ons++
		} else {
			b.values[v] &= ^vbits
			b.ons--
		}
	}
}

func (b Bits[I]) Get(index I) bool {
	v := index >> bitsShift
	vi := uint32(index) & bitsAnd
	vbits := bitsOne << vi
	return (b.values[v] & vbits) != 0
}

func (b Bits[I]) Ons() uint32 {
	return b.ons
}

func (b Bits[I]) IsEmpty() bool {
	return b.ons == 0
}

func (b *Bits[I]) Clear() {
	for i := range b.values {
		b.values[i] = 0
	}
	b.ons = 0
}

func (b Bits[I]) FirstOn() int32 {
	for vi, v := range b.values {
		if v != 0 {
			return int32((vi << bitsShift) + bits.TrailingZeros64(v))
		}
	}
	return -1
}

func (b *Bits[I]) TakeFirst() int32 {
	first := b.FirstOn()
	if first != -1 {
		b.Set(I(first), false)
	}
	return first
}

func (b Bits[I]) LastOn() int32 {
	for vi, v := range b.values {
		if v != 0 {
			return int32((vi << bitsShift) + bits.TrailingZeros64(v))
		}
	}
	return -1
}

func (b *Bits[I]) TakeLast() int32 {
	first := b.LastOn()
	if first != -1 {
		b.Set(I(first), false)
	}
	return first
}

func (b Bits[I]) OnAfter(index int32) int32 {
	max := b.Max()
	start := uint32(index + 1)
	if start > uint32(max) {
		return -1
	}
	startIndex := start >> bitsShift
	startBits := start & bitsAnd
	valueIndex := startIndex
	value := b.values[valueIndex]
	value &= ^((bitsOne << startBits) - 1)
	valuesMax := uint32(len(b.values))
	for value == 0 {
		valueIndex++
		if valueIndex == valuesMax {
			return -1
		}
		value = b.values[valueIndex]
	}
	next := (valueIndex - startIndex) << bitsShift
	next += start
	next += uint32(bits.TrailingZeros64(value))
	next -= startBits
	return int32(next)
}

func (b Bits[I]) OffAfter(index int32) int32 {
	max := b.Max()
	start := uint32(index + 1)
	if start > uint32(max) {
		return -1
	}
	startIndex := start >> bitsShift
	startBits := start & bitsAnd
	valueIndex := startIndex
	value := ^b.values[valueIndex]
	value &= ^((bitsOne << startBits) - 1)
	valuesMax := uint32(len(b.values))
	for value == 0 {
		valueIndex++
		if valueIndex == valuesMax {
			return -1
		}
		value = ^b.values[valueIndex]
	}
	next := (valueIndex - startIndex) << bitsShift
	next += start
	next += uint32(bits.TrailingZeros64(value))
	next -= startBits
	return int32(next)
}

func (b Bits[I]) SameOnes(a Bits[I]) bool {
	if a.ons != b.ons {
		return false
	}
	if a.ons == 0 {
		return true
	}
	large := a.values
	small := b.values
	if len(large) < len(small) {
		large = b.values
		small = a.values
	}
	for i := range small {
		if small[i] != large[i] {
			return false
		}
	}
	for i := len(small); i < len(large); i++ {
		if large[i] != 0 {
			return false
		}
	}
	return true
}

func (b Bits[I]) Clone() IBits[I] {
	values := make([]uint64, len(b.values))
	copy(values, b.values)
	return &Bits[I]{values, b.ons}
}

func (b *Bits[I]) Iterator() Iterator[I] {
	return &bitsIterator[I]{b, 0}
}

type bitsIterator[I index] struct {
	bits  IBits[I]
	index int32
}

func (i *bitsIterator[I]) Reset() {
	i.index = -1
}
func (i *bitsIterator[I]) HasNext() bool {
	return i.nextOn() != -1
}
func (i *bitsIterator[I]) Next() *I {
	i.index = i.nextOn()
	if i.index == -1 {
		return nil
	}
	ui := I(i.index)
	return &ui
}
func (i *bitsIterator[I]) nextOn() int32 {
	return i.bits.OnAfter(i.index)
}
func (i *bitsIterator[I]) Remove() {
	if i.index != -1 {
		i.bits.Set(I(i.index), false)
	}
}

type Bits64 = Bits64Indexed[uint32]

type Bits64Indexed[I index] uint64

func (b Bits64Indexed[I]) Max() I {
	return I(bitsAnd)
}
func (b *Bits64Indexed[I]) Set(index I, on bool) {
	indexMask := Bits64Indexed[I](bitsOne << index)
	if on {
		*b |= indexMask
	} else {
		*b &= ^indexMask
	}
}
func (b Bits64Indexed[I]) Get(index I) bool {
	return uint64(b)&(bitsOne<<index) != 0
}
func (b Bits64Indexed[I]) Ons() uint32 {
	return uint32(bits.OnesCount64(uint64(b)))
}
func (b Bits64Indexed[I]) IsEmpty() bool {
	return b == 0
}
func (b Bits64Indexed[I]) FirstOn() int32 {
	if b == 0 {
		return -1
	}
	return int32(bits.TrailingZeros64(uint64(b)))
}
func (b *Bits64Indexed[I]) TakeFirst() int32 {
	first := b.FirstOn()
	if first != -1 {
		b.Set(I(first), false)
	}
	return first
}
func (b Bits64Indexed[I]) LastOn() int32 {
	if b == 0 {
		return -1
	}
	return int32(64 - bits.LeadingZeros64(uint64(b)))
}
func (b *Bits64Indexed[I]) TakeLast() int32 {
	last := b.LastOn()
	if last != -1 {
		b.Set(I(last), false)
	}
	return last
}
func (b Bits64Indexed[I]) OnAfter(index int32) int32 {
	max := b.LastOn()
	after := index + 1
	if after > max {
		return -1
	}
	beforeMask := ^((bitsOne << after) - 1)
	withoutBefore := uint64(b) & beforeMask
	return int32(bits.TrailingZeros64(withoutBefore))
}
func (b Bits64Indexed[I]) OffAfter(index int32) int32 {
	max := b.LastOn()
	after := index + 1
	if after > max {
		return -1
	}
	beforeMask := ^((bitsOne << after) - 1)
	withoutBefore := uint64(b) & beforeMask
	return int32(bits.TrailingZeros64(^withoutBefore))
}
func (b Bits64Indexed[I]) All(all Bits64Indexed[I]) bool {
	return b&all == all
}
func (b *Bits64Indexed[I]) Flip(flip Bits64Indexed[I]) {
	*b ^= flip
}
func (b *Bits64Indexed[I]) Clear() {
	*b = 0
}
func (b Bits64Indexed[I]) Clone() IBits[I] {
	clone := b
	return &clone
}
func (b *Bits64Indexed[I]) Iterator() Iterator[I] {
	return &bitsIterator[I]{b, -1}
}
