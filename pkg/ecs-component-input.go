package axe

import "github.com/axe/axe-go/pkg/ds"

type InputActionListener struct {
	Handler func(action *InputAction) bool
}

var INPUTACTION = DefineComponent("input-action", InputActionListener{})

type InputActionSystem struct {
	unhandled InputActionHandler
	queue     ds.Queue[*InputAction]
}

var _ EntityDataSystem[InputActionListener] = &InputActionSystem{}

func NewInputActionSystem(unhandled InputActionHandler) EntityDataSystem[InputActionListener] {
	queue := ds.NewCircularQueue[*InputAction](32)

	return &InputActionSystem{
		unhandled: unhandled,
		queue:     &queue,
	}
}

func (sys InputActionSystem) OnStage(data *InputActionListener, e *Entity, ctx EntityContext) {
}
func (sys InputActionSystem) OnLive(data *InputActionListener, e *Entity, ctx EntityContext) {
}
func (sys InputActionSystem) OnRemove(data *InputActionListener, e *Entity, ctx EntityContext) {
}
func (sys InputActionSystem) Init(ctx EntityContext) error {
	ctx.Game.Actions.Handler = func(action *InputAction) {
		sys.queue.Push(action)
	}
	return nil
}
func (sys InputActionSystem) Update(iter ds.Iterable[EntityValue[*InputActionListener]], ctx EntityContext) {
	if sys.queue.IsEmpty() {
		return
	}

	listenerIterator := iter.Iterator()
	currentListener := listenerIterator.Next()

	if currentListener == nil {
		sys.queue.Clear()
		return
	}

	for !sys.queue.IsEmpty() {
		action := sys.queue.Pop()
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
func (sys InputActionSystem) Destroy(ctx EntityContext) {
	ctx.Game.Actions.Handler = nil
}
