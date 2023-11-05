package id

import "strings"

// Aliases

type LargeArea = Area[uint32, uint32]
type MediumArea = Area[uint32, uint16]
type SmallArea = Area[uint32, uint8]

// Identifier

type Identifier int32

func (i Identifier) Exists() bool   { return i >= 0 }
func (i Identifier) Empty() bool    { return i <= 0 }
func (i Identifier) String() string { return getString(i) }

var (
	mapping = make(map[string]int32, 256)
	list    = make([]string, 0, 256)
)

// Returns the identifier for the given string.
func Get(s string) Identifier {
	if id, exists := mapping[s]; exists {
		return Identifier(id)
	} else {
		nextId := len(list)
		list = append(list, s)
		mapping[s] = int32(nextId)
		return Identifier(nextId)
	}
}

// Returns the identifier for the lowercase version of the given string.
func Lower(s string) Identifier {
	return Get(strings.ToLower(s))
}

// Adds the empty string as identifier 0
func init() {
	Get("")
}

// Returns an identifier and if the string has not been added as one then -1 is returned.
func Maybe(s string) Identifier {
	id, exists := mapping[s]
	if exists {
		return Identifier(id)
	}
	return Identifier(-1)
}

func getString(i Identifier) string {
	if i >= 0 && int(i) < len(list) {
		return list[i]
	}
	return ""
}
