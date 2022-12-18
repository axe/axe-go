package axe

import "github.com/axe/axe-go/pkg/ds"

type Logic func(e *Entity, ctx EntityContext)

var LOGIC = DefineComponent("Logic", Logic(nil)).SetSystem(NewLogicSystem())

type LogicSystem struct{}

var _ EntityDataSystem[InputActionListener] = &InputActionSystem{}

func NewLogicSystem() EntityDataSystem[Logic] {
	return &LogicSystem{}
}

func (sys LogicSystem) OnStage(data *Logic, e *Entity, ctx EntityContext) {
}
func (sys LogicSystem) OnLive(data *Logic, e *Entity, ctx EntityContext) {
}
func (sys LogicSystem) OnRemove(data *Logic, e *Entity, ctx EntityContext) {
}
func (sys LogicSystem) Init(ctx EntityContext) error {
	return nil
}
func (sys LogicSystem) Update(iter ds.Iterable[EntityValue[*Logic]], ctx EntityContext) {
	i := iter.Iterator()
	for i.HasNext() {
		logic := i.Next()
		if logic.Data != nil {
			(*logic.Data)(logic.ID.Entity(), ctx)
		}
	}
}
func (sys LogicSystem) Destroy(ctx EntityContext) {
}
