package axe

import (
	"github.com/axe/axe-go/pkg/geom"
)

// TRANSFORM

type transformDirty uint8

const (
	transformDirtyNone transformDirty = (1 << iota) - 1
	transformDirtyLocal
	transformDirtyPosition
	transformDirtyRotation
	transformDirtyScale
)

type Transform[A Attr[A]] struct {
	Parent   *Transform[A]
	Local    Matrix[A]
	World    Matrix[A]
	Children []*Transform[A]

	dirty    transformDirty
	position A
	rotation A
	scale    A
}

type Transform2f = Transform[Vec2f]
type Transform3f = Transform[Vec3f]

func (t Transform[A]) IsDirty() bool {
	return t.dirty != transformDirtyNone
}

func (t *Transform[A]) SetLocal(local Matrix[A]) {
	t.Local = local
	t.dirty |= transformDirtyLocal
}
func (t *Transform[A]) GetLocal() Matrix[A] {
	t.updateLocal()
	return t.Local
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
		if t.Parent != nil {
			t.World.Mul(t.Parent.World, t.Local)
		} else {
			t.World.Set(t.Local)
		}
	}
	if len(t.Children) > 0 {
		for _, child := range t.Children {
			child.Update(updateWorld)
		}
	}
}

func (t *Transform[A]) SetParent(parent *Transform[A]) {
	if t.Parent != nil {
		for i, child := range t.Parent.Children {
			if child == t {
				t.Parent.Children = sliceRemoveAt(t.Parent.Children, i)
				break
			}
		}
	}
	t.Parent = parent
	if parent != nil {
		parent.Children = append(parent.Children, t)
	}
}

func (t *Transform[A]) updateLocal() {
	if t.dirty > transformDirtyLocal {
		t.Local.SetRotaton(t.rotation, false)
		t.Local.Scale(t.scale)
		t.Local.PostTranslate(t.position)
		t.dirty = transformDirtyLocal
	}
}

var TRANSFORM2 = DefineComponent("transform2", Transform2f{})
var TRANSFORM3 = DefineComponent("transform3", Transform3f{})

type TransformSystem[A Attr[A]] struct{}

var _ EntityDataSystem[Transform2f] = &TransformSystem[Vec2f]{}

func NewTransformSystem[A Attr[A]]() EntityDataSystem[Transform[A]] {
	return &TransformSystem[A]{}
}

func (sys *TransformSystem[A]) OnStage(e *Entity, data *Transform[A])  {}
func (sys *TransformSystem[A]) OnAdd(e *Entity, data *Transform[A])    {}
func (sys *TransformSystem[A]) OnRemove(e *Entity, data *Transform[A]) {}
func (sys *TransformSystem[A]) PreUpdate(game *Game) bool              { return true }
func (sys *TransformSystem[A]) PostUpdate(game *Game)                  {}
func (sys *TransformSystem[A]) Update(game *Game, e *Entity, data *Transform[A]) bool {
	if data.Parent == nil {
		data.Update(false)
	}
	return true
}

// MESH

type MeshVertex struct {
	Color  Color
	Coord  TextureCoord
	Point  geom.Vec3f
	Normal geom.Vec3f
}

type Mesh struct {
	Texture  Texture
	Vertices []MeshVertex
	Indices  []int
}

var MESH = DefineComponent("mesh", Mesh{})
