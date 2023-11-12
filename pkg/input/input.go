package input

import (
	"time"

	"github.com/axe/axe-go/pkg/react"
)

type Input struct {
	Name             string
	Value            float32
	ValueChanged     time.Time
	ValueDuration    time.Duration
	PreviousValue    float32
	PreviousChanged  time.Time
	PreviousDuration time.Duration
	Device           *Device
	Action           *Action
}

func New(name string) *Input {
	return &Input{Name: name}
}

func (i *Input) Set(value float32, now time.Time) bool {
	if i.Value != value {
		i.PreviousValue = i.Value
		i.PreviousChanged = i.ValueChanged
		i.PreviousDuration = now.Sub(i.ValueChanged)
		i.Value = value
		i.ValueChanged = now
		i.ValueDuration = 0
		return true
	}
	return false
}

func (i *Input) UpdateDuration(now time.Time) {
	i.ValueDuration = now.Sub(i.ValueChanged)
}

func (i *Input) Cancel() {
	i.Value = i.PreviousValue
	i.ValueChanged = i.PreviousChanged
}

func (i *Input) HasCancel() bool {
	return i.Value == i.PreviousValue && i.ValueChanged == i.PreviousChanged
}

type Data struct {
	Value  float32
	Inputs []*Input
}

func (data *Data) SetInputs(inputs []*Input) {
	if len(inputs) > 0 {
		data.Inputs = inputs
	} else if len(data.Inputs) > 0 {
		data.Inputs = data.Inputs[:0]
	}
}

func (data *Data) AddInputs(inputs []*Input) {
	if len(inputs) > 0 {
		if data.Inputs == nil {
			data.Inputs = inputs
		} else {
			data.Inputs = append(data.Inputs, inputs...)
		}
	}
}

type Gesture struct {
	Inputs  []*Input
	MinTime time.Duration
	Ignore  []*Input
}

type Vector struct {
	X *Input
	Y *Input
}

type Point struct {
	X     float32
	Y     float32
	Index int
	// Window *Window
	// Screen *Screen
}

type DeviceType string

const (
	DeviceTypeKeyboard   DeviceType = "keyboard"
	DeviceTypeMouse      DeviceType = "mouse"
	DeviceTypeController DeviceType = "controller"
	DeviceTypeTouch      DeviceType = "touch"
)

type Device struct {
	Name      string
	Type      DeviceType
	Inputs    []*Input
	Connected react.Value[bool]
}

func (device *Device) Add(name string) *Input {
	in := New(name)
	in.Device = device
	device.Inputs = append(device.Inputs, in)
	return in
}

func NewDevice(name string, deviceType DeviceType) *Device {
	return &Device{
		Name:      name,
		Type:      deviceType,
		Inputs:    make([]*Input, 0, 64),
		Connected: react.Val(true),
	}
}
