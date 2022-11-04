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
	Parent *Transform[A]
	Local  Matrix[A]
	World  Matrix[A]

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
func (t *Transform[A]) SetPosition(position A) {
	t.position = position
	t.dirty |= transformDirtyPosition
}
func (t *Transform[A]) SetRotation(rotation A) {
	t.rotation = rotation
	t.dirty |= transformDirtyRotation
}
func (t *Transform[A]) SetScale(scale A) {
	t.scale = scale
	t.dirty |= transformDirtyScale
}

func (t *Transform[A]) Update() {
	if t.dirty > transformDirtyLocal {
		if t.dirty == transformDirtyPosition {
			t.Local.SetRow(t.Local.Size()-1, t.position)
			t.dirty ^= transformDirtyPosition
		} else {
			// position, scale, rotation
		}
		t.dirty = transformDirtyLocal
	}
	if t.Parent != nil {
		isDirty := t.IsDirty() || t.Parent.IsDirty()
		if isDirty {
			t.World.Mul(t.Parent.World, t.Local)
		}
	} else if t.dirty != transformDirtyNone {
		t.World.Set(t.Local)
	}
	t.dirty = transformDirtyNone
}

var TRANSFORM2 = DefineComponent("transform2", Transform2f{}, true)
var TRANSFORM3 = DefineComponent("transform3", Transform3f{}, true)

type TransformSystem[A Attr[A]] struct {
	storage *ComponentStorage[Transform[A]]
}

func NewTransformSystem[A Attr[A]](storage *ComponentStorage[Transform[A]]) BaseSystem {
	return &TransformSystem[A]{storage}
}

var _ BaseSystem = &TransformSystem[Vec2f]{}

func (t *TransformSystem[A]) Init(game *Game) error { return nil }
func (t *TransformSystem[A]) Destroy()              {}
func (t *TransformSystem[A]) Update(game *Game) {
	for i := 0; i < t.storage.ValuesSize; i++ {
		t := &t.storage.Values[i]
		t.Update()
	}
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

var MESH = DefineComponent("mesh", Mesh{}, false)
