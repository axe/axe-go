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
	Name  string
	Jobs  *job.JobRunner
	World World
	Space Space[A, SpaceComponent[A]]
	Load  func(scene *Scene[A], game *Game)
}

type Scene2f = Scene[Vec2f]
type Scene3f = Scene[Vec3f]

var _ GameSystem = &Scene2f{}

func (scene *Scene[A]) Init(game *Game) error {
	if scene.Jobs == nil {
		scene.Jobs = job.NewRunner(game.Settings.JobGroups, game.Settings.JobBudget)
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
	scene.Jobs.Run()
	scene.World.Update(game)
	if scene.Space != nil {
		scene.Space.Update(game)
	}
}

func (scene *Scene[A]) Destroy() {
	scene.World.Destroy()
	if scene.Space != nil {
		scene.Space.Destroy()
	}
}

type Camera[A Attr[A]] interface {
	GameSystem
	Planes() []Plane[A]
	Intersects(shape Shape[A], position A) bool
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
	Draw()
}

type View2f = View[Vec2f]
type View3f = View[Vec3f]

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

type Plane[A Attr[A]] struct {
}
