package axe

import (
	"github.com/axe/axe-go/pkg/geom"
	"github.com/axe/axe-go/pkg/job"
	"github.com/axe/axe-go/pkg/ui"
)

type ProjectionOutside int

const (
	ProjectionOutsideIgnore ProjectionOutside = iota
	ProjectionOutsideClamp
	ProjectionOutsideRelative
)

type Scene[A Attr[A]] struct {
	Name   string
	Jobs   *job.JobRunner
	World  *World
	Space  Space[A, SpaceComponent[A]]
	Enable func(scene *Scene[A], game *Game)
	Load   func(scene *Scene[A], game *Game)
}

type Scene2f = Scene[Vec2f]
type Scene3f = Scene[Vec3f]

var _ GameSystem = &Scene2f{}

func (scene *Scene[A]) Init(game *Game) error {
	if scene.Jobs == nil {
		scene.Jobs = job.NewRunner(game.Settings.JobGroups, game.Settings.JobBudget)
	}
	if scene.World == nil {
		scene.World = NewWorld(game.Settings.WorldName, game.Settings.WorldSettings)
	}
	if scene.Enable != nil {
		scene.Enable(scene, game)
	}
	if scene.Load != nil {
		scene.Load(scene, game)
	}
	err := scene.World.Init(game)
	if err != nil {
		return err
	}
	if scene.Space != nil {
		err := scene.Space.Init(game)
		if err != nil {
			return err
		}
	}
	return nil
}

func (scene *Scene[A]) Update(game *Game) {
	if scene.Jobs != nil {
		scene.Jobs.Run()
	}
	if scene.World != nil {
		scene.World.Update(game)
	}
	if scene.Space != nil {
		scene.Space.Update(game)
	}
}

func (scene *Scene[A]) Destroy() {
	if scene.World != nil {
		scene.World.Destroy()
	}
	if scene.Space != nil {
		scene.Space.Destroy()
	}
}

type RenderTarget interface { // window or texture
	Size() geom.Vec2i
	Texture() *Texture
	Window() *Window
}

type View[A Attr[A]] interface {
	GameSystem

	Name() string
	Scene() Scene[A]
	Camera() Camera[A]
	ProjectionMatrix() Matrix[A]
	ViewMatrix() Matrix[A]
	CombinedMatrix() Matrix[A]
	ProjectPoint(mouse geom.Vec2i, outside ProjectionOutside) A
	Project(outside ProjectionOutside) A
	ProjectIgnore() A
	UnprojectPoint(point A, outside ProjectionOutside) A
	UnprojectIgnore(point A) A
	Placement() ui.Placement
	Target() RenderTarget
}

// type View2f = View[Vec2f]
// type View3f = View[Vec3f]

type View3f struct {
	Name             string
	SceneName        string
	Camera           Camera3d
	ProjectionMatrix Matrix4f
	ViewMatrix       Matrix4f
	CombinedMatrix   Matrix4f
	Placement        ui.Placement
	Target           RenderTarget

	OnInit    func(view *View3f, game *Game)
	OnUpdate  func(view *View3f, game *Game)
	OnDestroy func(view *View3f)

	scene *Scene3f
}

func (v *View3f) Init(game *Game) error {
	v.scene = game.Stages.Next.GetScene3f(v.SceneName)
	v.Placement.Init(ui.Maximized())
	v.Camera.Init(game)
	if v.OnInit != nil {
		v.OnInit(v, game)
	}
	return nil
}
func (v *View3f) Update(game *Game) {
	v.Camera.Update(game)
	if v.OnUpdate != nil {
		v.OnUpdate(v, game)
	}
}
func (v *View3f) Destroy() {
	v.Camera.Destroy()
	if v.OnDestroy != nil {
		v.OnDestroy(v)
	}
}
func (v View3f) Scene() *Scene3f                                                { return v.scene }
func (v View3f) ProjectPoint(mouse geom.Vec2i, outside ProjectionOutside) Vec3f { return Vec3f{} }
func (v View3f) Project(outside ProjectionOutside) Vec3f                        { return Vec3f{} }
func (v View3f) ProjectIgnore() Vec3f                                           { return Vec3f{} }
func (v View3f) UnprojectPoint(point Vec3f, outside ProjectionOutside) Vec3f    { return Vec3f{} }
func (v View3f) UnprojectIgnore(point Vec3f) Vec3f                              { return Vec3f{} }

type View2f struct {
	Name             string
	SceneName        string
	Camera           Camera2d
	ProjectionMatrix Matrix2f
	ViewMatrix       Matrix2f
	CombinedMatrix   Matrix2f
	Placement        ui.Placement
	Target           RenderTarget

	OnInit    func(view *View2f, game *Game)
	OnUpdate  func(view *View2f, game *Game)
	OnDestroy func(view *View2f)

	scene *Scene2f
}

func (v *View2f) Init(game *Game) error {
	v.scene = game.Stages.Next.GetScene2f(v.SceneName)
	v.Placement.Init(ui.Maximized())
	v.Camera.Init(game)
	if v.OnInit != nil {
		v.OnInit(v, game)
	}
	return nil
}
func (v *View2f) Update(game *Game) {
	v.Camera.Update(game)
	if v.OnUpdate != nil {
		v.OnUpdate(v, game)
	}
}
func (v *View2f) Destroy() {
	v.Camera.Destroy()
	if v.OnDestroy != nil {
		v.OnDestroy(v)
	}
}
func (v View2f) Scene() *Scene2f                                                { return v.scene }
func (v View2f) ProjectPoint(mouse geom.Vec2i, outside ProjectionOutside) Vec2f { return Vec2f{} }
func (v View2f) Project(outside ProjectionOutside) Vec2f                        { return Vec2f{} }
func (v View2f) ProjectIgnore() Vec2f                                           { return Vec2f{} }
func (v View2f) UnprojectPoint(point Vec2f, outside ProjectionOutside) Vec2f    { return Vec2f{} }
func (v View2f) UnprojectIgnore(point Vec2f) Vec2f                              { return Vec2f{} }

type EventSystem struct {
}

var _ GameSystem = &EventSystem{}

func (events *EventSystem) Init(game *Game) error { return nil }
func (events *EventSystem) Update(game *Game)     {}
func (events *EventSystem) Destroy()              {}

type GraphicsSystem interface { // & GameSystem
	GameSystem
}

type SpaceCoord interface {
	To2d() (x, y float32)
	To3d() (x, y, z float32)
}

type Watched[V any] struct {
	value V

	Changed bool
}

func (w *Watched[V]) Get() V {
	return w.value
}

func (w *Watched[V]) Ptr() *V {
	w.Changed = true
	return &w.value
}

func (w *Watched[V]) Set(value V) {
	w.value = value
	w.Changed = true
}
