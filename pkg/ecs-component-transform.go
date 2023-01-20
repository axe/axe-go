package axe

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ecs"
)

var TRANSFORM2 = ecs.DefineComponent("transform2", Transform3f{dirty: transformDirtyLocal}).SetSystem(NewTransformSystem3f())
var TRANSFORM3 = ecs.DefineComponent("transform3", Transform4f{dirty: transformDirtyLocal}).SetSystem(NewTransformSystem4f())

type Transform2f = Transform[Vec2f]
type Transform3f = Transform[Vec3f]
type Transform4f = Transform[Vec4f]

type TransformCreate2f = TransformCreate[Vec2f]
type TransformCreate3f = TransformCreate[Vec3f]
type TransformCreate4f = TransformCreate[Vec4f]

type transformDirty uint8

const (
	transformDirtyNone transformDirty = (1 << iota) - 1
	transformDirtyLocal
	transformDirtyPosition
	transformDirtyRotation
	transformDirtyScale
)

func NewTransform[A Attr[A]](create TransformCreate[A]) Transform[A] {
	return Transform[A]{
		local:    NewMatrix[A](),
		world:    NewMatrix[A](),
		position: create.Position,
		rotation: create.Rotation,
		scale:    create.Scale,
		dirty:    transformDirtyPosition | transformDirtyRotation | transformDirtyScale,
	}
}

func NewTransform2(create TransformCreate2f) Transform2f {
	return NewTransform(create)
}

func NewTransform3(create TransformCreate3f) Transform3f {
	return NewTransform(create)
}

func NewTransform4(create TransformCreate4f) Transform4f {
	return NewTransform(create)
}

type TransformCreate[A Attr[A]] struct {
	Position A
	Rotation A
	Scale    A
}

type Transform[A Attr[A]] struct {
	Tree ecs.Tree

	local    Matrix[A]
	world    Matrix[A]
	dirty    transformDirty
	position A
	rotation A
	scale    A
}

func (t Transform[A]) IsDirty() bool {
	return t.dirty != transformDirtyNone
}

func (t *Transform[A]) World() Matrix[A] {
	return t.world
}

func (t *Transform[A]) SetLocal(local Matrix[A]) {
	t.local = local
	t.dirty |= transformDirtyLocal
}

func (t *Transform[A]) Local() Matrix[A] {
	t.updateLocal()
	return t.local
}

func (t *Transform[A]) SetPosition(position A) {
	t.position = position
	t.dirty |= transformDirtyPosition
}

func (t *Transform[A]) GetPosition() A {
	return t.position
}

func (t *Transform[A]) SetRotation(rotation A) {
	t.rotation = rotation
	t.dirty |= transformDirtyRotation
}

func (t *Transform[A]) GetRotation() A {
	return t.rotation
}

func (t *Transform[A]) SetScale(scale A) {
	t.scale = scale
	t.dirty |= transformDirtyScale
}

func (t *Transform[A]) GetScale() A {
	return t.scale
}

func (t *Transform[A]) Update(updateWorld bool) {
	updateWorld = updateWorld || t.dirty != transformDirtyNone
	t.updateLocal()
	t.dirty = transformDirtyNone

	if updateWorld {
		if t.Tree.Parent() != nil {
			parentTransform := ecs.Get[Transform[A]](t.Tree.Parent())
			t.world.Mul(parentTransform.world, t.local)
		} else {
			t.world.Set(t.local)
		}
	}
	if len(t.Tree.Children()) > 0 {
		for _, child := range t.Tree.Children() {
			childTransform := ecs.Get[Transform[A]](child)
			childTransform.Update(updateWorld)
		}
	}
}

func (t *Transform[A]) SetParent(parent *ecs.Entity) {
	t.Tree.SetParent(parent)
}

func (t *Transform[A]) updateLocal() {
	if t.dirty > transformDirtyLocal {
		t.local.SetRotaton(t.rotation, false)
		t.local.Scale(t.scale)
		t.local.PostTranslate(t.position)
		t.dirty = transformDirtyLocal
	}
}

type TransformSystem[A Attr[A]] struct{}

var _ ecs.DataSystem[Transform2f] = &TransformSystem[Vec2f]{}

func NewTransformSystem[A Attr[A]]() ecs.DataSystem[Transform[A]] {
	return &TransformSystem[A]{}
}
func NewTransformSystem3f() ecs.DataSystem[Transform3f] {
	return &TransformSystem[Vec3f]{}
}
func NewTransformSystem4f() ecs.DataSystem[Transform4f] {
	return &TransformSystem[Vec4f]{}
}

func (sys *TransformSystem[A]) OnStage(data *Transform[A], e *ecs.Entity, ctx ecs.Context) {
	if data.local.columns == nil {
		InitMatrix(data.local)
	}
	if data.world.columns == nil {
		InitMatrix(data.world)
	}
}
func (sys *TransformSystem[A]) OnLive(data *Transform[A], e *ecs.Entity, ctx ecs.Context) {

}
func (sys *TransformSystem[A]) OnRemove(data *Transform[A], e *ecs.Entity, ctx ecs.Context) {

}
func (sys *TransformSystem[A]) Init(ctx ecs.Context) error {
	return nil
}
func (sys *TransformSystem[A]) Update(iter ds.Iterable[ecs.Value[*Transform[A]]], ctx ecs.Context) {
	i := iter.Iterator()
	for i.HasNext() {
		ev := i.Next()

		if ev.Data.Tree.Parent() == nil {
			ev.Data.Update(false)
		}
	}
}
func (sys *TransformSystem[A]) Destroy(ctx ecs.Context) {

}
