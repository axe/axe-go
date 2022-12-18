package ui

import "time"

type Coord struct {
	X float32
	Y float32
}

func (mp Coord) Equals(other Coord) bool {
	return mp.X == other.X && mp.Y == other.Y
}

type Event struct {
	Time    time.Time
	Stop    bool
	Capture bool
}

type PointerEventType string

const (
	PointerEventDown  PointerEventType = "down"
	PointerEventUp    PointerEventType = "up"
	PointerEventLeave PointerEventType = "leave"
	PointerEventEnter PointerEventType = "enter"
	PointerEventWheel PointerEventType = "wheel"
	PointerEventMove  PointerEventType = "move"
)

type PointerEvent struct {
	Event
	Point  Coord
	Button int
	Amount int
	Type   PointerEventType
}

func (ev PointerEvent) OfType(eventType PointerEventType) PointerEvent {
	return PointerEvent{
		Type:   eventType,
		Point:  ev.Point,
		Button: ev.Button,
		Amount: ev.Amount,
		Event: Event{
			Time:    ev.Time,
			Stop:    false,
			Capture: true,
		},
	}
}

type PointerButtons struct {
	Down     bool
	DownTime int64
}

type KeyEventType string

const (
	KeyEventDown  PointerEventType = "down"
	KeyEventUp    PointerEventType = "up"
	KeyEventPress PointerEventType = "press"
)

type KeyEvent struct {
	Event
	Key  int
	Char rune
	Type KeyEventType
}

type ComponentEvent struct {
	Event
	Target Component
}
