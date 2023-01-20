package axe

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ecs"
)

type BehaviorOnStage interface {
	OnStage(e *ecs.Entity, ctx ecs.Context)
}

type BehaviorOnLive interface {
	OnLive(e *ecs.Entity, ctx ecs.Context)
}

type BehaviorOnRemove interface {
	OnRemove(e *ecs.Entity, ctx ecs.Context)
}

type BehaviorOnUpdate interface {
	Update(e *ecs.Entity, ctx ecs.Context)
}

type BehaviorSystem[B any] struct{}

var _ ecs.DataSystem[InputActionListener] = &InputActionSystem{}

func NewBehaviorSystem[B any]() ecs.DataSystem[B] {
	return &BehaviorSystem[B]{}
}

func (sys BehaviorSystem[B]) OnStage(data *B, e *ecs.Entity, ctx ecs.Context) {
	if onStage, ok := any(data).(BehaviorOnStage); ok {
		onStage.OnStage(e, ctx)
	}
}
func (sys BehaviorSystem[B]) OnLive(data *B, e *ecs.Entity, ctx ecs.Context) {
	if onLive, ok := any(data).(BehaviorOnLive); ok {
		onLive.OnLive(e, ctx)
	}
}
func (sys BehaviorSystem[B]) OnRemove(data *B, e *ecs.Entity, ctx ecs.Context) {
	if onRemove, ok := any(data).(BehaviorOnRemove); ok {
		onRemove.OnRemove(e, ctx)
	}
}
func (sys BehaviorSystem[B]) Init(ctx ecs.Context) error {
	return nil
}
func (sys BehaviorSystem[B]) Update(iter ds.Iterable[ecs.Value[*B]], ctx ecs.Context) {
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
func (sys BehaviorSystem[B]) Destroy(ctx ecs.Context) {
}
