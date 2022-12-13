package axe

import (
	"fmt"
	"math"
	"testing"

	"github.com/axe/axe-go/pkg/ds"
)

/*
1. Define components
2. Define types
3. Create world
4. Enable components in world
5. Enable types in world
6. Add systems
7. Populate & update world
*/

var testTransform = DefineComponent("transform", testComponentTransform{})
var testMesh = DefineComponent("mesh", testComponentMesh{})
var testSpatial = DefineComponent("spatial", testComponentSpatial{})
var testTransformSpatial = DefineType("transform-spatial", testTypeTransformSpatial{}, func(t *EntityData[testTypeTransformSpatial]) {
	DefineTypeComponent(t, testTransform, func(data *testTypeTransformSpatial) *testComponentTransform { return &data.Transform })
	DefineTypeComponent(t, testSpatial, func(data *testTypeTransformSpatial) *testComponentSpatial { return &data.Spatial })
})

type testComponentTransform struct {
	X, Y             float32
	Angle            float32
	GlobalX, GlobalY float32
	Parent           *Entity
	Children         []*Entity
}

func (t *testComponentTransform) SetParent(e *Entity, parent *Entity) {
	parentTransform := testTransform.Get(parent)

	t.Parent = parent
	parentTransform.Children = append(parentTransform.Children, e)
}

func (t *testComponentTransform) Update() {
	if t.Parent == nil {
		t.GlobalX = t.X
		t.GlobalY = t.Y
	} else {
		parentTransform := testTransform.Get(t.Parent)
		theta := float64(parentTransform.Angle) / 180 * math.Pi
		cos := math.Cos(theta)
		sin := math.Sin(theta)
		x := float64(t.X)
		y := float64(t.Y)
		t.GlobalX = float32(x*cos - y*sin + float64(parentTransform.GlobalX))
		t.GlobalY = float32(y*cos + x*sin + float64(parentTransform.GlobalY))
	}
	if t.Children != nil {
		for _, child := range t.Children {
			childTransform := testTransform.Get(child)
			childTransform.Update()
		}
	}
}

type testSystemTransform struct{}

var _ EntityDataSystem[testComponentTransform] = &testSystemTransform{}

func (sys *testSystemTransform) OnStage(data *testComponentTransform, e *Entity, ctx EntityContext) {}
func (sys *testSystemTransform) OnLive(data *testComponentTransform, e *Entity, ctx EntityContext)  {}
func (sys *testSystemTransform) OnRemove(data *testComponentTransform, e *Entity, ctx EntityContext) {
}
func (sys *testSystemTransform) Init(ctx EntityContext) error { return nil }
func (sys *testSystemTransform) Destroy(ctx EntityContext)    {}
func (sys *testSystemTransform) Update(iter ds.Iterable[EntityValue[*testComponentTransform]], ctx EntityContext) {
	i := iter.Iterator()
	for i.HasNext() {
		t := i.Next()
		if t.Data.Parent == nil {
			t.Data.Update()
		}
	}
}

type testComponentMesh struct {
	Data [][]float32
}

type testComponentSpatial struct {
	Radius float32
}

type testTypeTransformSpatial struct {
	Transform testComponentTransform
	Spatial   testComponentSpatial
}

var testSettings = WorldSettings{
	EntityCapacity:            16,
	EntityStageCapacity:       8,
	AverageComponentPerEntity: 3,
}
var testDataSettings = EntityDataSettings{
	Capacity:      16,
	StageCapacity: 8,
}

func TestWorld(t *testing.T) {
	w := NewWorld("test", testSettings)
	w.Enable(testDataSettings, testTransform, testMesh, testSpatial, testTransformSpatial)

	testTransform.AddSystem(&testSystemTransform{})

	e1 := w.New()
	e1t := testTransform.Add(e1)
	e1t.Angle = 90
	e1s := testSpatial.Add(e1)
	e1s.Radius = 5

	e2 := w.New()
	e2t := testTransform.Add(e2)
	e2t.X = 10
	e2t.SetParent(e2, e1)

	w.Init(nil)
	w.Update(nil)

	transforms := testTransform.Iterable().Iterator()
	for transforms.HasNext() {
		t := transforms.Next()
		fmt.Printf("%+v\n", t.Data)
	}

	w.Delete(e2)
	w.Delete(e1)
	w.Destroy()
}
