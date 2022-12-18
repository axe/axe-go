package ui

import (
	"reflect"
)

type Bounds struct {
	Left, Top, Right, Bottom float32
}

func (b Bounds) Width() float32  { return b.Right - b.Left }
func (b Bounds) Height() float32 { return b.Bottom - b.Top }

type Color struct {
	R, G, B, A int
}

type Quad struct {
	Coords    Bounds
	Texture   string
	Placement Placement
}

type Tile struct {
	Coords  Bounds
	Texture string
}

type FocusTile struct {
	Tile
	Focus Coord
}

func (ft FocusTile) GetExtents() Bounds {
	l := ft.Focus.X / ft.Coords.Width()
	t := ft.Focus.Y / ft.Coords.Height()
	return Bounds{
		Left:   l,
		Right:  1 - l,
		Top:    t,
		Bottom: 1 - t,
	}
}

type Theme struct {
	Components  map[string]*ComponentTheme
	Fonts       map[string]*Font
	TextFormats map[string]*TextFormat
}

type ComponentType string

const (
	ComponentTypeContainer ComponentType = "container" // has children organized based on grid rules
	ComponentTypeButton    ComponentType = "button"    // clickable
	ComponentTypeList      ComponentType = "list"      // list of selectable things, tabs
	ComponentTypeText      ComponentType = "text"      // labels, editable text
	ComponentTypePopover   ComponentType = "popover"   // dropdown list, tooltip
	ComponentTypeCheckbox  ComponentType = "checkbox"  // checkbox, radio button
	ComponentTypeDynamic   ComponentType = "dynamic"   // tab, hidden panels,
)

type State string

const (
	StateDefault  State = "default"
	StateHover    State = "hover"
	StatePressed  State = "pressed"
	StateFocused  State = "focused"
	StateDisabled State = "disabled"
	StateSelected State = "selected" // checked, chosen option
)

type ComponentTheme struct {
	Visuals map[State][]Quad
}

type ComponentData struct {
}

type Base struct {
	Parent       Component
	Children     []Component
	Placement    Placement
	Visible      Value[bool]
	Disabled     Value[bool]
	State        State
	Events       ComponentEvents
	Visuals      map[State][]Quad
	LocalBounds  Bounds
	GlobalBounds Bounds
	Dirty        bool
	MaxSize      Coord
	MinSize      Coord
}

// TODO
// Set padding does placement
// min size from placement and on component
// trigger actions in event handling
// component templates
// fonts
// root data & component tree
// cursors
// popovers

// Composed types: textbox, tabs

type Value[V any] struct {
	Field    string
	Method   string
	Getter   string
	Setter   string
	Dynamic  bool
	Computed func() V
	Constant V
}

type Container struct {
	Padding Placement
}

type Button struct {
	Container
	Label   Text
	Pressed Value[bool]
}

type List struct {
	Container
	Items    Value[[]any]
	Selected Value[[]any]
}

type CheckboxState string

const (
	CheckboxStateChecked       = "checked"
	CheckboxStateUnchecked     = "unchecked"
	CheckboxStateIndeterminate = "indeterminate"
)

type Checkbox struct {
	Container
	Label   Text
	Checked Value[CheckboxState]
	Value   Value[any]
	Values  map[CheckboxState]any
}

type TextWrap string

const (
	TextWrapNone TextWrap = "none"
	TextWrapWord TextWrap = "word"
	TextWrapChar TextWrap = "char"
)

type Text struct {
	Content Value[string]
	Format  *TextFormat
}

type Input struct { // without visuals, just event handling and text
	Value      Value[string]
	Selectable bool
	Editable   bool
}

type Tooltip struct {
	Content Value[string]
}

type Image struct {
	Image Quad
}

type Scrollable struct {
}

type ComponentEvents struct {
	OnPointerEvent func(ev *PointerEvent)
	OnKeyEvent     func(ev *KeyEvent)
	OnFocus        func(ev *ComponentEvent)
	OnBlur         func(ev *ComponentEvent)
}

func NewButton(text string) *Base {
	return &Base{
		Events: ComponentEvents{
			OnPointerEvent: func(ev *PointerEvent) {
				if !ev.Capture {
					ev.Stop = true

				}
			},
		},
	}
}

type Component interface {
	Parent() Component
	At(pt Coord) Component
	IsFocusable() bool
	OnPointerEvent(ev *PointerEvent)
	OnKeyEvent(ev *KeyEvent)
	OnFocus(ev *ComponentEvent)
	OnBlur(ev *ComponentEvent)
}

func toPtr(x any) uintptr {
	return reflect.ValueOf(x).Pointer()
}

type ComponentMap map[uintptr]Component

func (cm ComponentMap) Add(c Component) {
	cm[toPtr(c)] = c
}

func (cm ComponentMap) AddMany(c []Component) {
	for _, m := range c {
		cm.Add(m)
	}
}

func (cm ComponentMap) Has(c Component) bool {
	_, exists := cm[toPtr(c)]
	return exists
}

func (cm ComponentMap) AddLineage(c Component) {
	curr := c
	for curr != nil {
		cm.Add(curr)
		curr = curr.Parent()
	}
}

func (old ComponentMap) Compare(new ComponentMap) (inOld []Component, inBoth []Component, inNew []Component) {
	inOld = make([]Component, 0)
	inBoth = make([]Component, 0)
	inNew = make([]Component, 0)

	for _, oldOverAncestor := range old {
		if !new.Has(oldOverAncestor) {
			inOld = append(inOld, oldOverAncestor)
		} else {
			inBoth = append(inBoth, oldOverAncestor)
		}
	}
	for _, newOverAncestor := range new {
		if !new.Has(newOverAncestor) {
			inNew = append(inNew, newOverAncestor)
		}
	}
	return
}
