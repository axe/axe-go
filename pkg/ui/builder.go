package ui

import "github.com/axe/axe-go/pkg/ds"

type Builder struct {
	shape      ds.Stack[Shape]
	placement  ds.Stack[Placement]
	background ds.Stack[Background]
	visual     ds.Stack[Visual]
	states     ds.Stack[StateFn]

	unitsToPoint float32
	borderRadius AmountCorners
	borderWidth  AmountBounds

	base Base
}

func NewBuilder() Builder {
	return Builder{
		shape:        ds.NewStack[Shape](8),
		placement:    ds.NewStack[Placement](8),
		background:   ds.NewStack[Background](8),
		visual:       ds.NewStack[Visual](8),
		states:       ds.NewStack[StateFn](8),
		unitsToPoint: 0.5,
		base: Base{
			States: StateDefault,
		},
	}
}

func (b Builder) Shape(shape Shape) Builder {
	b.shape.Push(shape)
	return b
}
func (b Builder) ShapePop() Builder {
	b.shape.Pop()
	return b
}
func (b Builder) ShapeRectangle() Builder {
	return b.Shape(ShapeRectangle{})
}
func (b Builder) Radius(radius float32) Builder {
	b.borderRadius.Set(radius, UnitConstant)
	return b
}
func (b Builder) RadiusOval() Builder {
	b.borderRadius.Set(0.5, UnitParent)
	return b
}
func (b Builder) ShapeRounded() Builder {
	return b.Shape(ShapeRounded{
		Radius:       b.borderRadius,
		UnitToPoints: b.unitsToPoint,
	})
}

func (b Builder) Place(placement Placement) Builder {
	if b.placement.Count == 0 {
		b.base.Placement = placement
	}
	b.placement.Push(placement)
	return b
}
func (b Builder) PlacePop() Builder {
	b.placement.Pop()
	return b
}
func (b Builder) Maximized() Builder {
	return b.Place(Maximized())
}
func (b Builder) Grow(amount float32) Builder {
	return b.Place(b.placement.Peek().Grow(amount))
}
func (b Builder) Shrink(amount float32) Builder {
	return b.Place(b.placement.Peek().Shrink(amount))
}
func (b Builder) Shift(dx, dy float32) Builder {
	return b.Place(b.placement.Peek().Shift(dx, dy))
}

func (b Builder) Background(background Background) Builder {
	b.background.Push(background)
	return b
}
func (b Builder) BackgroundPop() Builder {
	b.background.Pop()
	return b
}
func (b Builder) BackgroundColor(color Color) Builder {
	return b.Background(BackgroundColor{Color: color})
}
func (b Builder) BackgroundLinearGradient(startColor Color, start Coord, endColor Color, end Coord) Builder {
	return b.Background(BackgroundLinearGradient{StartColor: startColor, Start: start, EndColor: endColor, End: end})
}
func (b Builder) BackgroundImage(tile Tile) Builder {
	return b.Background(BackgroundImage{Tile: tile})
}

func (b Builder) Visual(visual Visual) Builder {
	b.visual.Push(visual)
	return b
}
func (b Builder) VisualPop() Builder {
	b.visual.Pop()
	return b
}
func (b Builder) Filled() Builder {
	return b.Visual(VisualFilled{Shape: b.shape.Peek()})
}
func (b Builder) Bordered(width float32, color Color, outerColor *Color) Builder {
	vb := VisualBordered{Width: width, Shape: b.shape.Peek(), InnerColor: color, HasInnerColor: true, OuterColor: color, HasOuterColor: true}
	if outerColor != nil {
		vb.OuterColor = *outerColor
	}
	return b.Visual(vb)
}

func (b Builder) States(states StateFn) Builder {
	b.states.Push(states)
	return b
}
func (b Builder) StatesPop() Builder {
	b.states.Pop()
	return b
}

func (b Builder) Layer() Builder {
	b.base.Layers = append(b.base.Layers, Layer{
		Placement:  b.placement.Peek(),
		Visual:     b.visual.Peek(),
		Background: b.background.Peek(),
		States:     b.states.Peek(),
	})
	return b
}

func (b Builder) Focusable() Builder {
	b.base.Focusable = true
	return b
}

func (b Builder) OnClick(handle func(ev *PointerEvent)) Builder {
	return b.OnPointerButtonEvent(PointerEventDown, 0, handle)
}
func (b Builder) OnRightClick(handle func(ev *PointerEvent)) Builder {
	return b.OnPointerButtonEvent(PointerEventDown, 1, handle)
}
func (b Builder) OnScroll(handle func(ev *PointerEvent)) Builder {
	return b.OnPointerButtonEvent(PointerEventDown, 2, handle)
}
func (b Builder) OnPointerUp(handle func(ev *PointerEvent)) Builder {
	return b.OnPointerMoveEvent(PointerEventUp, handle)
}
func (b Builder) OnPointerButtonEvent(pointerEvent PointerEventType, button int, handle func(ev *PointerEvent)) Builder {
	return b.AddPointerEvent(func(ev *PointerEvent) {
		if !ev.Capture && ev.Type == pointerEvent && ev.Button == button {
			handle(ev)
		}
	})
}
func (b Builder) OnEnter(handle func(ev *PointerEvent)) Builder {
	return b.OnPointerMoveEvent(PointerEventEnter, handle)
}
func (b Builder) OnLeave(handle func(ev *PointerEvent)) Builder {
	return b.OnPointerMoveEvent(PointerEventLeave, handle)
}
func (b Builder) OnPointerMoveEvent(pointerEvent PointerEventType, handle func(ev *PointerEvent)) Builder {
	return b.AddPointerEvent(func(ev *PointerEvent) {
		if !ev.Capture && ev.Type == pointerEvent {
			handle(ev)
		}
	})
}
func (b Builder) AddPointerEvent(handler func(ev *PointerEvent)) Builder {
	existing := b.base.Events.OnPointer
	if existing == nil {
		b.base.Events.OnPointer = handler
	} else {
		b.base.Events.OnPointer = func(ev *PointerEvent) {
			existing(ev)
			if !ev.Stop {
				handler(ev)
			}
		}
	}
	return b
}
func (b Builder) OnFocus(handle func(ev *ComponentEvent)) Builder {
	b.base.Events.OnFocus = func(ev *ComponentEvent) {
		if !ev.Capture {
			handle(ev)
		}
	}
	return b
}
func (b Builder) OnBlur(handle func(ev *ComponentEvent)) Builder {
	b.base.Events.OnBlur = func(ev *ComponentEvent) {
		if !ev.Capture {
			handle(ev)
		}
	}
	return b
}
func (b Builder) OnKeyDown(key string, handler func(ev *KeyEvent)) Builder {
	return b.AddKeyEvent(func(ev *KeyEvent) {
		if !ev.Capture && ev.Key == key {
			handler(ev)
		}
	})
}
func (b Builder) AddKeyEvent(handler func(ev *KeyEvent)) Builder {
	existing := b.base.Events.OnKey
	if existing == nil {
		b.base.Events.OnKey = handler
	} else {
		b.base.Events.OnKey = func(ev *KeyEvent) {
			existing(ev)
			if !ev.Stop {
				handler(ev)
			}
		}
	}
	return b
}

func (b Builder) End() Component {
	return &b.base
}
