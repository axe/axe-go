package ui

import "github.com/axe/axe-go/pkg/id"

type Theme struct {
	TextStyles       TextStyles
	StatePostProcess map[State]PostProcess
	Fonts            id.DenseMap[*Font, uint16, uint8]
	Cursors          id.DenseMap[ExtentTile, uint16, uint8]
	DefaultCursor    id.Identifier
	Animations       Animations
	Colors           Colors
}

type State = Flags

const (
	StateDefault State = 1 << iota
	StateHover
	StatePressed
	StateFocused
	StateDisabled
	StateDragging
	StateDragOver
	StateSelected // checked, chosen option
	StateHidden
	// States that only occur when there is an animation for Show/Hide/Remove. The component is considered disabled.
	StateRemoving
	StateHiding
	StateShowing

	StatesDisabled = StateDisabled | StateRemoving | StateShowing | StateHiding | StateHidden
)

type StateFn = func(s State) bool

type Dirty = Flags

const (
	DirtyNone Dirty = (1 << iota) >> 1
	DirtyPlacement
	DirtyVisual
	DirtyChildPlacement
	DirtyChildVisual
	DirtyPostProcess
)

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

func Override[V any](value V) *V {
	return &value
}
