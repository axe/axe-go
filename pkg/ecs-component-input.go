package axe

import (
	"github.com/axe/axe-go/pkg/core"
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/input"
)

type InputActionListener struct {
	Handler func(action *input.Action) bool
}

var ACTION = ecs.DefineComponent("input-action", InputActionListener{}).SetSystem(NewInputActionSystem(nil))

type InputActionSystem struct {
	unhandled input.ActionHandler
	queue     ds.Queue[*input.Action]
}

var _ ecs.DataSystem[InputActionListener] = &InputActionSystem{}

func NewInputActionSystem(unhandled input.ActionHandler) ecs.DataSystem[InputActionListener] {
	queue := ds.NewCircularQueue[*input.Action](32)

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

type InputEvents = input.SystemEvents

var INPUT = ecs.DefineComponent("input", InputEvents{}).SetSystem(NewInputEventsSystem())

type InputEventsSystem struct {
	connectedDevices    ds.Stack[input.Device]
	disconnectedDevices ds.Stack[input.Device]
	connectedInput      ds.Stack[input.Input]
	disconnectedInput   ds.Stack[input.Input]
	changedInput        ds.Stack[input.Input]
	connectedPoint      ds.Stack[input.Point]
	disconnectedPoint   ds.Stack[input.Point]
	changedPoint        ds.Stack[input.Point]

	off core.ListenerOff
}

var _ ecs.DataSystem[InputEvents] = &InputEventsSystem{}

func NewInputEventsSystem() ecs.DataSystem[InputEvents] {
	return &InputEventsSystem{
		connectedDevices:    ds.NewStack[input.Device](4),
		disconnectedDevices: ds.NewStack[input.Device](4),
		connectedInput:      ds.NewStack[input.Input](64),
		disconnectedInput:   ds.NewStack[input.Input](64),
		changedInput:        ds.NewStack[input.Input](64),
		connectedPoint:      ds.NewStack[input.Point](8),
		disconnectedPoint:   ds.NewStack[input.Point](8),
		changedPoint:        ds.NewStack[input.Point](8),
	}
}

func (sys InputEventsSystem) OnStage(data *InputEvents, e *ecs.Entity, ctx ecs.Context) {
}
func (sys InputEventsSystem) OnLive(data *InputEvents, e *ecs.Entity, ctx ecs.Context) {
}
func (sys InputEventsSystem) OnRemove(data *InputEvents, e *ecs.Entity, ctx ecs.Context) {
}
func (sys *InputEventsSystem) Init(ctx ecs.Context) error {
	if sys.off != nil {
		return nil
	}

	sys.off = ActiveGame().Input.Events().On(input.SystemEvents{
		DeviceConnected: func(newDevice input.Device) {
			sys.connectedDevices.Push(newDevice)
		},
		DeviceDisconnected: func(oldDevice input.Device) {
			sys.disconnectedDevices.Push(oldDevice)
		},
		InputConnected: func(newInput input.Input) {
			sys.connectedInput.Push(newInput)
		},
		InputDisconnected: func(oldInput input.Input) {
			sys.disconnectedInput.Push(oldInput)
		},
		InputChange: func(input input.Input) {
			sys.changedInput.Push(input)
		},
		PointConnected: func(newPoint input.Point) {
			sys.connectedPoint.Push(newPoint)
		},
		PointDisconnected: func(oldPoint input.Point) {
			sys.disconnectedPoint.Push(oldPoint)
		},
		PointChange: func(point input.Point) {
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
