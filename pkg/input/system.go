package input

import (
	"time"

	"github.com/axe/axe-go/pkg/core"
	"github.com/axe/axe-go/pkg/util"
)

type SystemEvents struct {
	DeviceConnected    func(newDevice Device)
	DeviceDisconnected func(oldDevice Device)
	InputConnected     func(newInput Input)
	InputDisconnected  func(oldInput Input)
	InputChange        func(input Input)
	PointConnected     func(newPoint Point)
	PointDisconnected  func(oldPoint Point)
	PointChange        func(point Point)
	PointLeave         func(point Point)
	PointEnter         func(point Point)
	InputChangeMap     map[string]func(input Input)
}

type InputSystem interface {
	Devices() []*Device
	Inputs() []*Input
	Points() []*Point
	Events() *core.Listeners[SystemEvents]
	Get(inputName string) *Input
	InputTime() time.Time
}

type System struct {
	events   core.Listeners[SystemEvents]
	inputs   []*Input
	devices  []*Device
	inputMap map[string]*Input
	points   []*Point
	now      time.Time
}

func NewSystem() *System {
	return &System{
		events:   *core.NewListeners[SystemEvents](),
		inputs:   make([]*Input, 0, 256),
		devices:  make([]*Device, 0, 8),
		inputMap: make(map[string]*Input, 256),
		points:   make([]*Point, 0, 4),
		now:      time.Now(),
	}
}

func (sys *System) Devices() []*Device {
	return sys.devices
}

func (sys *System) Inputs() []*Input {
	return sys.inputs
}

func (sys *System) Points() []*Point {
	return sys.points
}

func (sys *System) Events() *core.Listeners[SystemEvents] {
	return &sys.events
}

func (sys *System) Get(inputName string) *Input {
	return sys.inputMap[inputName]
}

func (sys *System) InputTime() time.Time {
	return sys.now
}

func (sys *System) SetInputTime(now time.Time) {
	sys.now = now
}

func (sys *System) AddPoint(point *Point) {
	sys.points = append(sys.points, point)
}

func (in *System) ConnectDevice(newDevice *Device) {
	in.devices = append(in.devices, newDevice)

	in.events.Trigger(func(listener SystemEvents) bool {
		if listener.DeviceConnected != nil {
			listener.DeviceConnected(*newDevice)
			return true
		}
		return false
	})

	for _, input := range newDevice.Inputs {
		in.ConnectInput(input)
	}
}

func (in *System) ConnectInput(newInput *Input) {
	in.inputs = append(in.inputs, newInput)
	in.inputMap[newInput.Name] = newInput

	in.events.Trigger(func(listener SystemEvents) bool {
		if listener.InputConnected != nil {
			listener.InputConnected(*newInput)
			return true
		}
		return false
	})
}

func (in *System) DisconnectDevice(oldDevice *Device) {
	in.devices = util.SliceRemove(in.devices, oldDevice)

	in.events.Trigger(func(listener SystemEvents) bool {
		if listener.DeviceDisconnected != nil {
			listener.DeviceDisconnected(*oldDevice)
			return true
		}
		return false
	})

	for _, input := range oldDevice.Inputs {
		in.DisconnectInput(input)
	}
}

func (in *System) DisconnectInput(oldInput *Input) {
	in.inputs = util.SliceRemove(in.inputs, oldInput)
	if in.inputMap[oldInput.Name] == oldInput {
		delete(in.inputMap, oldInput.Name)
	}

	in.events.Trigger(func(listener SystemEvents) bool {
		if listener.InputDisconnected != nil {
			listener.InputDisconnected(*oldInput)
			return true
		}
		return false
	})
}

func (in *System) SetInputValue(ia *Input, newValue float32) {
	if ia != nil {
		if !ia.Set(newValue, in.now) {
			return
		}

		in.events.Trigger(func(listener SystemEvents) bool {
			handled := false
			if listener.InputChange != nil {
				listener.InputChange(*ia)
				handled = true
			}
			if listener.InputChangeMap != nil {
				inputListener := listener.InputChangeMap[ia.Name]
				if inputListener != nil {
					inputListener(*ia)
					handled = true
				}
			}
			return handled
		})
	}
}

func (in *System) SetInputPoint(ia *Point, x int, y int) {
	if ia != nil {
		ia.X = x
		ia.Y = y
		in.events.Trigger(func(listener SystemEvents) bool {
			if listener.PointChange != nil {
				listener.PointChange(*ia)
				return true
			}
			return false
		})
	}
}

func (in *System) SetInputLeave(ia *Point) {
	if ia != nil {
		in.events.Trigger(func(listener SystemEvents) bool {
			if listener.PointLeave != nil {
				listener.PointLeave(*ia)
				return true
			}
			return false
		})
	}
}

func (in *System) SetInputEnter(ia *Point) {
	if ia != nil {
		in.events.Trigger(func(listener SystemEvents) bool {
			if listener.PointEnter != nil {
				listener.PointEnter(*ia)
				return true
			}
			return false
		})
	}
}
