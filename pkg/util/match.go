package util

import "golang.org/x/exp/constraints"

// A match is a function which is given a single integer and returns whether it matches some criteria.
type Match[T constraints.Integer] func(value T) bool

func MatchAlways[T constraints.Integer]() Match[T] {
	return func(value T) bool {
		return true
	}
}

// Creates a match which returns true if all bits in test are on in a given value.
func MatchAll[T constraints.Integer](test T) Match[T] {
	return func(value T) bool {
		return value&test == test
	}
}

// Creates a match which returns true if all bits in a given value are on in test.
func MatchOnly[T constraints.Integer](test T) Match[T] {
	return func(value T) bool {
		return value&test == value
	}
}

// Creates a match which returns true if a given value and test are equal.
func MatchExact[T constraints.Integer](test T) Match[T] {
	return func(value T) bool {
		return value&value == test
	}
}

// Creates a match which returns true if any bits in test are in a given value.
func MatchAny[T constraints.Integer](test T) Match[T] {
	return func(value T) bool {
		return value&test != 0
	}
}

// Creates a match which returns true if there are no bits in common with a given value.
func MatchNone[T constraints.Integer](test T) Match[T] {
	return func(value T) bool {
		return value&test == 0
	}
}

// Creates a match which returns true when a given value is zero.
func MatchEmpty[T constraints.Integer]() Match[T] {
	return func(value T) bool {
		return value&value == 0
	}
}

// Creaes a match which returns the negation of given match.
func MatchNot[T constraints.Integer](not Match[T]) Match[T] {
	return func(value T) bool {
		return !not(value)
	}
}

// Creates a match which returns true if all given matches return true.
func MatchAnd[T constraints.Integer](ands ...Match[T]) Match[T] {
	return func(value T) bool {
		for _, and := range ands {
			if !and(value) {
				return false
			}
		}
		return true
	}
}

// Creates a match which returns true if any given matches return true.
func MatchOr[T constraints.Integer](ors ...Match[T]) Match[T] {
	return func(value T) bool {
		for _, or := range ors {
			if or(value) {
				return true
			}
		}
		return false
	}
}
