package axe

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ecs"
)

type Logic func(e *ecs.Entity, ctx ecs.Context)

var LOGIC = ecs.DefineComponent("Logic", Logic(nil)).SetSystem(NewLogicSystem())

type LogicSystem struct{}

var _ ecs.DataSystem[InputActionListener] = &InputActionSystem{}

func NewLogicSystem() ecs.DataSystem[Logic] {
	return &LogicSystem{}
}

func (sys LogicSystem) OnStage(data *Logic, e *ecs.Entity, ctx ecs.Context) {
}
func (sys LogicSystem) OnLive(data *Logic, e *ecs.Entity, ctx ecs.Context) {
}
func (sys LogicSystem) OnRemove(data *Logic, e *ecs.Entity, ctx ecs.Context) {
}
func (sys LogicSystem) Init(ctx ecs.Context) error {
	return nil
}
func (sys LogicSystem) Update(iter ds.Iterable[ecs.Value[*Logic]], ctx ecs.Context) {
	i := iter.Iterator()
	for i.HasNext() {
		logic := i.Next()
		if logic.Data != nil {
			(*logic.Data)(logic.ID.Entity(), ctx)
		}
	}
}
func (sys LogicSystem) Destroy(ctx ecs.Context) {
}
