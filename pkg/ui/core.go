package ui

import (
	"github.com/axe/axe-go/pkg/color"
	"github.com/axe/axe-go/pkg/id"
)

// The Theme of a UI. Contains the base text styles, state specific post processing, the collection
// of supported fonts & cursors, the default cursor, the theme wide named animations, and the theme
// wide colors.
type Theme struct {
	TextStyles       TextStyles
	StatePostProcess map[State]PostProcess
	Fonts            id.DenseMap[*Font, uint16, uint8]
	Cursors          id.DenseMap[ExtentTile, uint16, uint8]
	DefaultCursor    id.Identifier
	Animations       Animations
	Colors           color.Colors
}

// One or more states of a component. The states can define which layers or post processing is done.
type State = Flags

const (
	// Default enabled state
	StateDefault State = 1 << iota
	// Enabled state with pointer over component
	StateHover
	// Enabled state with pointer down on component. Can be other states as well.
	StatePressed
	// Enabled state when component is the currently focused one.
	StateFocused
	// Component was disabled by user and should receive no input events.
	StateDisabled
	// Enabled state when component is being drug.
	StateDragging
	// Enabled state when component accepts drops and there's a drag element over it.
	StateDragOver
	// General state for more complex components (checked, chosen option, etc).
	StateSelected
	// Disabled state when component is hidden from render parent and is not placed.
	StateHidden
	// Disabled state when component has removing animation that's playing.
	StateRemoving
	// Disabled state when component has hiding animation that's playing.
	StateHiding
	// Disabled state when component has showing animation that's playing.
	StateShowing

	// All states that can't receive input events.
	StatesDisabled = StateDisabled | StateRemoving | StateShowing | StateHiding | StateHidden
)

// A function which accepts a state and returns true/false.
// Used to control which layers are updated & rendered for a component's state.
type StateFn = func(s State) bool

// The possible dirty states of a component. Dirty states are used to keep track of what
// operations needs to be done (if at all) on component Place and Render. If a component
// tree is not animating or has a post processing state (for example) then render/clipping
// logic can be avoided for the current frame and cached render data can be used.
type Dirty = Flags

const (
	// No dirty state
	DirtyNone Dirty = (1 << iota) >> 1
	// The components placement in its parent has changed. After placement is done if the
	// bounds of the component has changed the visual is also marked dirty.
	DirtyPlacement
	// The components layers need to be rendered.
	DirtyVisual
	// The components children had placement changed and need to be layout and placed again.
	DirtyChildPlacement
	// The components children had its layer, children, or post processing changed.
	DirtyChildVisual
	// The component's post processing has changed and needs to be processed again.
	DirtyPostProcess
)

// Converts this Dirty into what should be added to the parent's Dirty state.
// If the placement is dirty, then the parent state will have DirtyChildPlacement.
// If the visual is dirty, then the parent state will have DirtyChildVisual.
// If the post process is dirty, then the parent state will have DirtyChildVisual.
func (d Dirty) ParentDirty() Dirty {
	parentDirty := d
	if d.Is(DirtyPlacement) {
		parentDirty.Remove(DirtyPlacement)
		parentDirty.Add(DirtyChildPlacement)
	}
	if d.Is(DirtyVisual) {
		parentDirty.Remove(DirtyVisual)
		parentDirty.Add(DirtyChildVisual)
	}
	if d.Is(DirtyPostProcess) {
		parentDirty.Remove(DirtyPostProcess)
		parentDirty.Add(DirtyChildVisual)
	}
	return parentDirty
}

// Provides an override for a value, which is really just a pointer to a copy of the value.
// This is useful function for the simple properties in the types which have concrete
// and override types (like Text/Paragraph/Paragraphs Styles)
func Override[V any](value V) *V {
	return &value
}
