package axe

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ecs"
)

type InputActionListener struct {
	Handler func(action *InputAction) bool
}

var ACTION = ecs.DefineComponent("input-action", InputActionListener{}).SetSystem(NewInputActionSystem(nil))

type InputActionSystem struct {
	unhandled InputActionHandler
	queue     ds.Queue[*InputAction]
}

var _ ecs.DataSystem[InputActionListener] = &InputActionSystem{}

func NewInputActionSystem(unhandled InputActionHandler) ecs.DataSystem[InputActionListener] {
	queue := ds.NewCircularQueue[*InputAction](32)

	return &InputActionSystem{
		unhandled: unhandled,
		queue:     &queue,
	}
}

func (sys InputActionSystem) OnStage(data *InputActionListener, e *ecs.Entity, ctx ecs.Context) {
}
func (sys InputActionSystem) OnLive(data *InputActionListener, e *ecs.Entity, ctx ecs.Context) {
}
func (sys InputActionSystem) OnRemove(data *InputActionListener, e *ecs.Entity, ctx ecs.Context) {
}
func (sys *InputActionSystem) Init(ctx ecs.Context) error {
	return nil
}
func (sys *InputActionSystem) Update(iter ds.Iterable[ecs.Value[*InputActionListener]], ctx ecs.Context) {
	current := ActiveGame().Stages.Current
	if current == nil {
		return
	}

	triggeredIterator := current.Actions.Iterable().Iterator()
	if !triggeredIterator.HasNext() {
		return
	}

	listenerIterator := iter.Iterator()
	if !listenerIterator.HasNext() {
		return
	}

	currentListener := listenerIterator.Next()
	for triggeredIterator.HasNext() {
		action := *triggeredIterator.Next()
		startListener := currentListener
		handled := false

		for {
			if currentListener.Data.Handler(action) {
				handled = true
				break
			}
			currentListener = listenerIterator.Next()
			if currentListener == nil {
				listenerIterator.Reset()
				currentListener = listenerIterator.Next()
			}
			if currentListener == startListener {
				break
			}
		}

		if !handled && sys.unhandled != nil {
			sys.unhandled(action)
		}
	}
}
func (sys *InputActionSystem) Destroy(ctx ecs.Context) {

}

type InputEvents = InputSystemEvents

var INPUT = ecs.DefineComponent("input", InputEvents{}).SetSystem(NewInputEventsSystem())

type InputEventsSystem struct {
	connectedDevices    ds.Stack[InputDevice]
	disconnectedDevices ds.Stack[InputDevice]
	connectedInput      ds.Stack[Input]
	disconnectedInput   ds.Stack[Input]
	changedInput        ds.Stack[Input]
	connectedPoint      ds.Stack[InputPoint]
	disconnectedPoint   ds.Stack[InputPoint]
	changedPoint        ds.Stack[InputPoint]

	off ListenerOff
}

var _ ecs.DataSystem[InputEvents] = &InputEventsSystem{}

func NewInputEventsSystem() ecs.DataSystem[InputEvents] {
	return &InputEventsSystem{
		connectedDevices:    ds.NewStack[InputDevice](4),
		disconnectedDevices: ds.NewStack[InputDevice](4),
		connectedInput:      ds.NewStack[Input](64),
		disconnectedInput:   ds.NewStack[Input](64),
		changedInput:        ds.NewStack[Input](64),
		connectedPoint:      ds.NewStack[InputPoint](8),
		disconnectedPoint:   ds.NewStack[InputPoint](8),
		changedPoint:        ds.NewStack[InputPoint](8),
	}
}

func (sys InputEventsSystem) OnStage(data *InputEvents, e *ecs.Entity, ctx ecs.Context) {
}
func (sys InputEventsSystem) OnLive(data *InputEvents, e *ecs.Entity, ctx ecs.Context) {
}
func (sys InputEventsSystem) OnRemove(data *InputEvents, e *ecs.Entity, ctx ecs.Context) {
}
func (sys *InputEventsSystem) Init(ctx ecs.Context) error {
	sys.off = ActiveGame().Input.Events().On(InputSystemEvents{
		DeviceConnected: func(newDevice InputDevice) {
			sys.connectedDevices.Push(newDevice)
		},
		DeviceDisconnected: func(oldDevice InputDevice) {
			sys.disconnectedDevices.Push(oldDevice)
		},
		InputConnected: func(newInput Input) {
			sys.connectedInput.Push(newInput)
		},
		InputDisconnected: func(oldInput Input) {
			sys.disconnectedInput.Push(oldInput)
		},
		InputChange: func(input Input) {
			sys.changedInput.Push(input)
		},
		PointConnected: func(newPoint InputPoint) {
			sys.connectedPoint.Push(newPoint)
		},
		PointDisconnected: func(oldPoint InputPoint) {
			sys.disconnectedPoint.Push(oldPoint)
		},
		PointChange: func(point InputPoint) {
			sys.changedPoint.Push(point)
		},
	})
	return nil
}
func (sys *InputEventsSystem) Update(iter ds.Iterable[ecs.Value[*InputEvents]], ctx ecs.Context) {
	i := iter.Iterator()
	for i.HasNext() {
		evts := i.Next().Data

		if evts.DeviceConnected != nil {
			for i := 0; i < sys.connectedDevices.Count; i++ {
				evts.DeviceConnected(sys.connectedDevices.Items[i])
			}
		}
		if evts.DeviceDisconnected != nil {
			for i := 0; i < sys.disconnectedDevices.Count; i++ {
				evts.DeviceDisconnected(sys.disconnectedDevices.Items[i])
			}
		}
		if evts.InputConnected != nil {
			for i := 0; i < sys.connectedInput.Count; i++ {
				evts.InputConnected(sys.connectedInput.Items[i])
			}
		}
		if evts.InputDisconnected != nil {
			for i := 0; i < sys.disconnectedInput.Count; i++ {
				evts.InputDisconnected(sys.disconnectedInput.Items[i])
			}
		}
		if evts.InputChange != nil {
			for i := 0; i < sys.changedInput.Count; i++ {
				evts.InputChange(sys.changedInput.Items[i])
			}
		}
		if evts.PointConnected != nil {
			for i := 0; i < sys.connectedPoint.Count; i++ {
				evts.PointConnected(sys.connectedPoint.Items[i])
			}
		}
		if evts.PointDisconnected != nil {
			for i := 0; i < sys.disconnectedPoint.Count; i++ {
				evts.PointDisconnected(sys.disconnectedPoint.Items[i])
			}
		}
		if evts.PointChange != nil {
			for i := 0; i < sys.changedPoint.Count; i++ {
				evts.PointChange(sys.changedPoint.Items[i])
			}
		}
		if evts.InputChangeMap != nil {
			for i := 0; i < sys.changedInput.Count; i++ {
				input := sys.changedInput.Items[i]
				if handler, ok := evts.InputChangeMap[input.Name]; ok {
					handler(input)
				}
			}
		}
	}

	sys.connectedDevices.Clear()
	sys.disconnectedDevices.Clear()
	sys.connectedInput.Clear()
	sys.disconnectedInput.Clear()
	sys.changedInput.Clear()
	sys.connectedPoint.Clear()
	sys.disconnectedPoint.Clear()
	sys.changedPoint.Clear()
}
func (sys *InputEventsSystem) Destroy(ctx ecs.Context) {
	sys.off()
}
