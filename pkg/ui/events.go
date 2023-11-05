package ui

import (
	"time"
)

type CanStop interface {
	Stopped() bool
}

type Listener[E CanStop] func(ev E)

func (l Listener[E]) Trigger(ev E) {
	if l != nil {
		l(ev)
	}
}

func (a *Listener[E]) Add(b Listener[E], before bool) {
	if *a == nil {
		*a = b
	} else if b != nil {
		var first, second Listener[E]
		if before {
			first = b
			second = *a
		} else {
			first = *a
			second = b
		}

		*a = func(ev E) {
			first(ev)
			if !ev.Stopped() {
				second(ev)
			}
		}
	}
}

type Events struct {
	OnPointer Listener[*PointerEvent]
	OnKey     Listener[*KeyEvent]
	OnFocus   Listener[*Event]
	OnBlur    Listener[*Event]
	OnDrag    Listener[*DragEvent]
}

func (e *Events) Add(add Events, before bool) {
	e.OnPointer.Add(add.OnPointer, before)
	e.OnKey.Add(add.OnKey, before)
	e.OnFocus.Add(add.OnFocus, before)
	e.OnBlur.Add(add.OnBlur, before)
	e.OnDrag.Add(add.OnDrag, before)
}

type Event struct {
	Time    time.Time
	Stop    bool
	Capture bool
	Cancel  bool
	Target  Component
}

var _ CanStop = &Event{}

func (e Event) withTarget(target Component) Event {
	e.Target = target
	return e
}

func (e *Event) StopPropagation() {
	e.Stop = true
}

func (e *Event) PreventDefault() {
	e.Cancel = true
}

func (e *Event) Stopped() bool {
	return e.Stop
}

type DragEventType int

const (
	DragEventStart DragEventType = iota
	DragEventMove
	DragEventEnd
	DragEventCancel
	DragEventOver
	DragEventDrop
)

type DragEvent struct {
	Event
	Point      Coord
	Start      Coord
	DeltaStart Coord
	DeltaMove  Coord
	Type       DragEventType
	Dragging   Component
	*HasCursor
}

func (ev DragEvent) as(dragType DragEventType) *DragEvent {
	ev.Event = Event{Capture: false}
	ev.Type = dragType
	return &ev
}

type PointerEventType int

const (
	PointerEventDown PointerEventType = iota
	PointerEventUp
	PointerEventLeave
	PointerEventEnter
	PointerEventWheel
	PointerEventMove
)

type PointerEvent struct {
	Event
	Point  Coord
	Button int
	Amount int
	Type   PointerEventType
	*HasCursor
}

func (e PointerEvent) as(eventType PointerEventType) PointerEvent {
	e.Event = Event{Capture: true}
	e.Type = eventType
	return e
}

func (e PointerEvent) withTarget(target Component) PointerEvent {
	e.Target = target
	return e
}

type PointerButtons struct {
	Down     bool
	DownTime int64
}

type KeyEventType int

const (
	KeyEventDown KeyEventType = iota
	KeyEventUp
	KeyEventPress
)

type KeyEvent struct {
	Event
	Key  string
	Char rune
	Type KeyEventType
}
