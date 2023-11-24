package ui

import (
	"time"

	"github.com/axe/axe-go/pkg/util"
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

func listenerNil[E CanStop](a Listener[E]) bool {
	return a == nil
}
func listenerJoin[E CanStop](first Listener[E], second Listener[E]) Listener[E] {
	return func(ev E) {
		first(ev)
		if !ev.Stopped() {
			second(ev)
		}
	}
}

func (a *Listener[E]) Add(b Listener[E], before bool) {
	*a = util.CoalesceJoin(*a, b, before, listenerNil[E], listenerJoin[E])
}

func (a *Listener[E]) Clear() {
	*a = nil
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

func (e *Events) Clear() {
	e.OnPointer.Clear()
	e.OnKey.Clear()
	e.OnFocus.Clear()
	e.OnBlur.Clear()
	e.OnDrag.Clear()
}

type Event struct {
	Time    time.Time
	Stop    bool
	Capture bool
	Cancel  bool
	Target  *Base
}

var _ CanStop = &Event{}

func (e Event) withTarget(target *Base) Event {
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
	Dragging   *Base
	*HasCursor
}

func (ev DragEvent) as(dragType DragEventType) *DragEvent {
	ev.Event = Event{Capture: false}
	ev.Type = dragType
	return &ev
}

type PointerEventType int

const (
	// Press down (might not be followed by up)
	PointerEventDown PointerEventType = iota
	// Press up (might not be followed by )
	PointerEventUp
	// When a component previously received a down but the latest down excludes them.
	PointerEventDownOut
	// When a component previously received an up but the latest up excludes them.
	PointerEventUpOut
	// A pointer has leaved a component.
	PointerEventLeave
	// A pointer has entered a component.
	PointerEventEnter
	// A scroll/wheel event.
	PointerEventWheel
	// The pointer has moved.
	PointerEventMove
)

type PointerEvent struct {
	Event
	Point  Coord
	Button int
	Wheel  Coord
	Type   PointerEventType
	*HasCursor
}

func (e PointerEvent) as(eventType PointerEventType) PointerEvent {
	e.Event = Event{Capture: true}
	e.Type = eventType
	return e
}

func (e PointerEvent) withTarget(target *Base) PointerEvent {
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
