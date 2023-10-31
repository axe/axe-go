package ui

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Init struct {
	Theme *Theme
}

type Update struct {
	DeltaTime time.Duration
}

type Bounds struct {
	Left, Top, Right, Bottom float32
}

func (b Bounds) Width() float32 {
	return b.Right - b.Left
}
func (b Bounds) Height() float32 {
	return b.Bottom - b.Top
}
func (b Bounds) Dimensions() (float32, float32) {
	return b.Width(), b.Height()
}
func (b *Bounds) Translate(x, y float32) {
	b.Left += x
	b.Right += x
	b.Top += y
	b.Bottom += y
}
func (b Bounds) Dx(x float32) float32 {
	return (x - b.Left) / b.Width()
}
func (b Bounds) Dy(y float32) float32 {
	return (y - b.Top) / b.Height()
}
func (b Bounds) Contains(c Coord) bool {
	return !(c.X < b.Left || c.X > b.Right || c.Y < b.Top || c.Y > b.Bottom)
}

type Tile struct {
	Coords  Bounds
	Texture string
}

func (t Tile) Coord(dx, dy float32) TexCoord {
	return TexCoord{
		Texture: t.Texture,
		Coord: Coord{
			X: lerp(t.Coords.Left, t.Coords.Right, dx),
			Y: lerp(t.Coords.Top, t.Coords.Bottom, dy),
		},
	}
}

type ExtentTile struct {
	Tile
	Extent Bounds
}

type Theme struct {
	DefaultFontSize  float32
	DefaultFontColor Color
	DefaultFont      string

	Components map[string]*ComponentTheme
	Fonts      map[string]*Font
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

type State uint8

const (
	StateDefault State = 1 << iota
	StateHover
	StatePressed
	StateFocused
	StateDisabled
	StateDragging
	StateDragOver
	StateSelected // checked, chosen option
)

func (s State) Has(t State) bool { return (s & t) != 0 }

type StateFn = func(s State) bool

func StateAny(e State) StateFn {
	return func(s State) bool {
		return (s & e) != 0
	}
}
func StateAll(e State) StateFn {
	return func(s State) bool {
		return (s & e) == e
	}
}
func StateExact(e State) StateFn {
	return func(s State) bool {
		return s == e
	}
}
func StateEvery() StateFn {
	return func(s State) bool {
		return true
	}
}
func StateNot(e State) StateFn {
	return func(s State) bool {
		return (e & s) == 0
	}
}
func StateAnyNot(anyState, notState State) StateFn {
	return func(s State) bool {
		return (anyState&s) != 0 && (notState&s) == 0
	}
}

type ComponentTheme struct {
}

type TextWrap string

const (
	TextWrapNone TextWrap = "none"
	TextWrapWord TextWrap = "word"
	TextWrapChar TextWrap = "char"
)

func (w TextWrap) MarshalText() ([]byte, error) {
	return []byte(w), nil
}

func (w *TextWrap) UnmarshalText(text []byte) error {
	s := strings.ToLower(string(text))
	switch s {
	case "none", "n", "":
		*w = TextWrapNone
	case "word", "w":
		*w = TextWrapWord
	case "char", "c", "letter":
		*w = TextWrapChar
	default:
		return fmt.Errorf("invalid text wrap: " + s)
	}
	return nil
}

type Coord struct {
	X float32
	Y float32
}

type TexCoord struct {
	Coord
	Texture string
}

func (mp Coord) Equals(other Coord) bool {
	return mp.X == other.X && mp.Y == other.Y
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
		if !old.Has(newOverAncestor) {
			inNew = append(inNew, newOverAncestor)
		}
	}
	return
}
