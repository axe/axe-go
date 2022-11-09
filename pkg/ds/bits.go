package ds

import "math/bits"

const (
	bitsShift uint32 = 6
	bitsAnd   uint32 = 63
)

type Bits struct {
	values []uint64
	ons    uint32
}

func NewBits(max uint32) Bits {
	return Bits{
		values: make([]uint64, (max>>bitsShift)+1),
		ons:    0,
	}
}

func (b Bits) Max() uint32 {
	return (uint32(len(b.values)) << bitsShift) - 1
}

func (b *Bits) EnsureMax(index uint32) {
	if b.Max() < index {
		len := (index >> bitsShift) + 1
		values := make([]uint64, len)
		copy(values, b.values)
		b.values = values
	}
}

func (b *Bits) SafeSet(index uint32, on bool) {
	b.EnsureMax(index)
	b.Set(index, on)
}

func (b *Bits) Set(index uint32, on bool) {
	v := index >> bitsShift
	vi := index & bitsAnd
	vbits := uint64(1) << vi
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

func (b Bits) Get(index uint32) bool {
	v := index >> bitsShift
	vi := index & bitsAnd
	vbits := uint64(1) << vi
	return (b.values[v] & vbits) != 0
}

func (b Bits) Ons() uint32 {
	return b.ons
}

func (b Bits) IsEmpty() bool {
	return b.ons == 0
}

func (b *Bits) Clear() {
	for i := range b.values {
		b.values[i] = 0
	}
	b.ons = 0
}

func (b Bits) FirstOn() int32 {
	for vi, v := range b.values {
		if v != 0 {
			return int32((vi << bitsShift) + bits.TrailingZeros64(v))
		}
	}
	return -1
}

func (b *Bits) TakeFirst() int32 {
	first := b.FirstOn()
	if first != -1 {
		b.Set(uint32(first), false)
	}
	return first
}

func (b Bits) SameOnes(a Bits) bool {
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

func (b Bits) Clone() Bits {
	values := make([]uint64, len(b.values))
	copy(values, b.values)
	return Bits{values, b.ons}
}
