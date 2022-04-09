package ecs

type FlagMatch func(expected uint64, actual uint64) bool

var (
	MATCH_AND = func(a FlagMatch, b FlagMatch) FlagMatch {
		return func(x uint64, y uint64) bool {
			return a(x, y) && b(x, y)
		}
	}
	MATCH_OR = func(a FlagMatch, b FlagMatch) FlagMatch {
		return func(x uint64, y uint64) bool {
			return a(x, y) || b(x, y)
		}
	}
	MATCH_NOT = func(a FlagMatch) FlagMatch {
		return func(x uint64, y uint64) bool {
			return !a(x, y)
		}
	}
	MATCH_ANY = func(expected uint64, actual uint64) bool {
		return true
	}
	MATCH_NOTHING = func(expected uint64, actual uint64) bool {
		return false
	}
	MATCH_ALL = func(expected uint64, actual uint64) bool {
		return (expected & actual) == expected
	}
	MATCH_SOME = func(expected uint64, actual uint64) bool {
		return (expected & actual) != 0
	}
	MATCH_NONE = func(expected uint64, actual uint64) bool {
		return (expected & actual) == 0
	}
	MATCH_NOT_ALL = func(expected uint64, actual uint64) bool {
		return (expected & actual) != expected
	}
	MATCH_EXACTLY = func(expected uint64, actual uint64) bool {
		return expected == actual
	}
)
