package util

import "golang.org/x/exp/constraints"

// Utility struct for typed bitwise operations.
type Flags[T constraints.Integer] struct {
	value T
}

// Adds the flags to this set.
func (f *Flags[T]) Set(flags T) {
	f.value = f.value | flags
}

// Removes the flags from this set.
func (f *Flags[T]) Remove(flags T) {
	f.value = f.value & ^flags
}

// Changes this set to be only values that are shared between it and the flags.
func (f *Flags[T]) Only(flags T) {
	f.value = f.value & flags
}

// Toggles the flags in this set.
func (f *Flags[T]) Toggle(flags T) {
	f.value = f.value ^ flags
}

// Clears all flags from this set.
func (f *Flags[T]) Clear() {
	f.value = 0
}

// Returns whether this set is empty.
func (f Flags[T]) IsEmpty() bool {
	return f.value == 0
}

// Returns the current integer value in this set.
func (f Flags[T]) Get() T {
	return f.value
}

// Returns whether this set of flags matches the given Match function.
func (f Flags[T]) Is(match Match[T]) bool {
	return match(f.value)
}
