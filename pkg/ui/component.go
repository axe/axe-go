package ui

import "time"

type Init struct {
	Theme *Theme
}

type Update struct {
	DeltaTime time.Duration
}

type Component interface {
	Init(init Init)
	Place(parent Bounds, force bool)
	Update(update Update)
	Render(ctx AmountContext, out *VertexBuffer)
	GetDirty() Dirty
	Dirty(dirty Dirty)
	Parent() Component
	At(pt Coord) Component

	IsFocusable() bool
	IsDraggable() bool
	IsDroppable() bool

	OnPointer(ev *PointerEvent)
	OnKey(ev *KeyEvent)
	OnFocus(ev *ComponentEvent)
	OnBlur(ev *ComponentEvent)
	OnDrag(ev *DragEvent)
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
