package ux

import "github.com/axe/axe-go/pkg/ui"

// icons, options, searching, etc
// a function is used to generate the items given a data value

type DropdownSettings struct {
	TriggerOnDown  bool
	TriggerOnEnter bool
	TriggerOnFocus bool
	ShowAnimation  ui.AnimationFactory
	HideAnimation  ui.AnimationFactory
}

type Dropdown[I any] struct {
	Items           []I
	PlaceholderText string
	Placeholder     HasComponent
	Value           I
	ItemText        func(item I) string
	ItemComponent   func(item I) HasComponent
	OnChange        func(item I)
	ValueCompare    func(a, b I) bool

	DropdownSettings
}

type DropdownBase[I any] struct {
}
