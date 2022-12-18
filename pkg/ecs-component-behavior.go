package axe

import "github.com/axe/axe-go/pkg/ds"

type BehaviorOnStage interface {
	OnStage(e *Entity, ctx EntityContext)
}

type BehaviorOnLive interface {
	OnLive(e *Entity, ctx EntityContext)
}

type BehaviorOnRemove interface {
	OnRemove(e *Entity, ctx EntityContext)
}

type BehaviorOnUpdate interface {
	Update(e *Entity, ctx EntityContext)
}

type BehaviorSystem[B any] struct{}

var _ EntityDataSystem[InputActionListener] = &InputActionSystem{}

func NewBehaviorSystem[B any]() EntityDataSystem[B] {
	return &BehaviorSystem[B]{}
}

func (sys BehaviorSystem[B]) OnStage(data *B, e *Entity, ctx EntityContext) {
	if onStage, ok := any(data).(BehaviorOnStage); ok {
		onStage.OnStage(e, ctx)
	}
}
func (sys BehaviorSystem[B]) OnLive(data *B, e *Entity, ctx EntityContext) {
	if onLive, ok := any(data).(BehaviorOnLive); ok {
		onLive.OnLive(e, ctx)
	}
}
func (sys BehaviorSystem[B]) OnRemove(data *B, e *Entity, ctx EntityContext) {
	if onRemove, ok := any(data).(BehaviorOnRemove); ok {
		onRemove.OnRemove(e, ctx)
	}
}
func (sys BehaviorSystem[B]) Init(ctx EntityContext) error {
	return nil
}
func (sys BehaviorSystem[B]) Update(iter ds.Iterable[EntityValue[*B]], ctx EntityContext) {
	var test *B
	if _, ok := any(test).(BehaviorOnUpdate); !ok {
		return
	}

	i := iter.Iterator()
	for i.HasNext() {
		v := i.Next()
		u := any(v.Data).(BehaviorOnUpdate)
		u.Update(v.ID.Entity(), ctx)
	}
}
func (sys BehaviorSystem[B]) Destroy(ctx EntityContext) {
}
