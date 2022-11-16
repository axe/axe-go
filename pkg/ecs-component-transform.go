package axe

import "github.com/axe/axe-go/pkg/ds"

var TRANSFORM2 = DefineComponent("transform2", Transform3f{dirty: transformDirtyLocal})
var TRANSFORM3 = DefineComponent("transform3", Transform4f{dirty: transformDirtyLocal})

type Transform2f = Transform[Vec2f]
type Transform3f = Transform[Vec3f]
type Transform4f = Transform[Vec4f]

type transformDirty uint8

const (
	transformDirtyNone transformDirty = (1 << iota) - 1
	transformDirtyLocal
	transformDirtyPosition
	transformDirtyRotation
	transformDirtyScale
)

type Transform[A Attr[A]] struct {
	Tree EntityTree

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
		if t.Tree.parent != nil {
			parentTransform := Get[Transform[A]](t.Tree.parent)
			t.world.Mul(parentTransform.world, t.local)
		} else {
			t.world.Set(t.local)
		}
	}
	if len(t.Tree.children) > 0 {
		for _, child := range t.Tree.children {
			childTransform := Get[Transform[A]](child)
			childTransform.Update(updateWorld)
		}
	}
}

func (t *Transform[A]) SetParent(parent *Entity) {
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

var _ EntityDataSystem[Transform2f] = &TransformSystem[Vec2f]{}

func NewTransformSystem[A Attr[A]]() EntityDataSystem[Transform[A]] {
	return &TransformSystem[A]{}
}

func (sys *TransformSystem[A]) OnStage(data *Transform[A], e *Entity, ctx EntityContext) {

}
func (sys *TransformSystem[A]) OnLive(data *Transform[A], e *Entity, ctx EntityContext) {
	if data.local.columns == nil {
		InitMatrix(data.local)
	}
	if data.world.columns == nil {
		InitMatrix(data.world)
	}
}
func (sys *TransformSystem[A]) OnRemove(data *Transform[A], e *Entity, ctx EntityContext) {

}
func (sys *TransformSystem[A]) Init(ctx EntityContext) error {
	return nil
}
func (sys *TransformSystem[A]) Update(iter ds.Iterable[EntityValue[*Transform[A]]], ctx EntityContext) {
	i := iter.Iterator()
	for i.HasNext() {
		ev := i.Next()

		if ev.Data.Tree.parent == nil {
			ev.Data.Update(false)
		}
	}
}
func (sys *TransformSystem[A]) Destroy(ctx EntityContext) {

}
