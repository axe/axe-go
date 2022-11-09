package ecs

import "math/bits"

type ComponentSet uint64

func (set ComponentSet) Max() int {
	return 64 - bits.LeadingZeros64(uint64(set))
}
func (set ComponentSet) Size() int {
	return bits.OnesCount64(uint64(set))
}
func (set ComponentSet) Empty() bool {
	return set == 0
}
func (set *ComponentSet) Set(index uint8) {
	*set |= 1 << index
}
func (set *ComponentSet) Unset(index uint8) {
	*set &= ^(1 << index)
}
func (set ComponentSet) Has(index uint8) bool {
	return set&(1<<index) != 0
}
func (set *ComponentSet) Take() uint8 {
	x := uint8(bits.TrailingZeros64(uint64(*set)))
	*set ^= 1 << x
	return x
}
