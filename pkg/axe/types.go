package axe

import (
	"reflect"
	"time"

	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/geom"
	"github.com/axe/axe-go/pkg/input"
	"github.com/axe/axe-go/pkg/job"
	"github.com/axe/axe-go/pkg/react"
	"github.com/axe/axe-go/pkg/ui"
	"github.com/axe/axe-go/pkg/util"
)

type GameSystem interface {
	Init(game *Game) error
	Update(game *Game)
	Destroy()
}

type GameTime struct {
	Name        string
	DayDuration time.Duration
	Enabled     bool
	Scale       float32
	DateTime    time.Time
	Elapsed     time.Duration
	StartTime   time.Time
}

type Timer struct {
	LastTick  time.Time
	Current   time.Time
	Elapsed   time.Duration
	Frequency time.Duration
	Ticks     int64
}

type GameState struct {
	StartTime   time.Time
	Times       []GameTime
	UpdateTimer Timer
	DrawTimer   Timer
}

type Game struct {
	Debug    DebugSystem
	Assets   AssetSystem
	Windows  WindowSystem
	Graphics GraphicsSystem
	Input    InputSystem
	Actions  input.ActionSets
	Audio    AudioSystem
	Events   EventSystem
	Stages   StageManager
	State    GameState
	Settings GameSettings
	Running  bool
}

type AssetSystem interface{}

type EventSystem interface{}

type WindowSystem interface {
	GameSystem
	MainWindow() Window
	Windows() []Window
	Screens() []Screen
	// Events() *core.Listeners[WindowSystemEvents]
}

type Window interface {
	Name() string
	Title() react.Value[string]
	Placement() ui.Placement
	Screen() Screen
	Size() geom.Vec2i
}

type Screen interface {
	Size() geom.Vec2i
	Position() geom.Vec2i
}

type GraphicsSystem interface { // & GameSystem
	GameSystem
}

type InputSystem interface {
	GameSystem
	input.InputSystem
}

type AudioSystem interface { // & GameSystem
	GameSystem

	// Listeners() []AudioListener
	// Instances() []AudioInstance
	// Settings() map[string]AudioSettings
	// Sources() []AudioSource
	// EntitySystem() ecs.DataSystem[AudioEmitter]
}

type GameSettings struct {
	Name                 string
	EnableDebug          bool
	FixedUpdateFrequency time.Duration
	FixedDrawFrequency   time.Duration
	FirstStage           string
	JobGroups            int
	JobBudget            int
	Stages               []Stage
	Assets               []asset.Ref
	Windows              []StageWindow
	WorldName            string
	WorldSettings        ecs.WorldSettings
}

type StageManager struct {
	Stages  map[string]*Stage
	Current *Stage
	Next    *Stage
}

type StageWindow struct {
	Name       string
	Title      string
	Placement  ui.Placement
	Fullscreen bool
}

type Stage struct {
	Name    string
	Assets  []asset.Ref
	Windows []StageWindow
	Scenes  []Scene
	Views   []View
	Actions input.ActionSets
}

type View interface {
	GameSystem
	Name() string
	Scene() Scene
	Camera() Camera
	Placement() ui.Placement
}

type Scene struct {
	Name   string
	Jobs   *job.JobRunner
	World  *ecs.World
	Space  Space
	Enable func(scene *Scene, game *Game)
	Load   func(scene *Scene, game *Game)
}

type Camera interface {
}

type Space interface {
	GameSystem
	Collisions(flags util.Match[int], callback SpaceCollisionCallback)
	Intersects(query SpaceQuery, callback SpaceSearchCallback) int
	Contains(query SpaceQuery, callback SpaceSearchCallback) int
	Raytrace(query SpaceQuery, callback SpaceSearchCallback) int
	KNN(query SpaceQuery, nearest []SpaceNearest, nearestCount *int)
}

type Matrix struct{}

type SpaceComponent struct {
	Shape          Shape
	Offset         SpaceCoord
	WorldTransform *Matrix
	Flags          int
	Static         bool
	Inert          bool
}

type SpaceQuery struct {
	Point   SpaceCoord
	End     SpaceCoord
	Shape   Shape
	Maximum int
	Flags   util.Match[int]
}

type SpaceNearest struct {
	Entity   any
	Distance float32
}

type SpaceSearchCallback func(entity any, overlap float32, index int, query SpaceQuery) bool

type SpaceCollisionCallback func(subject any, otherSubject any, overlap float32, index int, second bool)

type Shape interface {
	Finite() bool
	Distance(point SpaceCoord) float32
	Normal(point SpaceCoord, out *SpaceCoord) bool
	Raytrace(point SpaceCoord, direction SpaceCoord) bool
	// Bounds(bounds *Bounds[A]) bool
	// Round(round *Round[A]) bool
	// PlaneSign(plane Plane[A]) PlaneSign
	// Transform(shapeTransform Trans[A]) Shape[A]
}

type DebugSystem struct {
	Logs      []DebugLog
	Snapshots []DebugSnapshot
	Graphs    []DebugGraph
	Events    []DebugEvent
	Enabled   bool
}

type DebugLog struct {
	Severity int
	Message  string
	Data     any
}
type DebugSnapshot struct {
	At      time.Time
	Elapsed time.Duration
}
type DebugGraph struct {
	Placement ui.Placement
	Database  *StatDatabase
	Set       *StatSet
}
type DebugProfiler struct{}

func (prof *DebugProfiler) Begin(ev string) {}
func (prof *DebugProfiler) End()            {}

type StatDatabase struct {
	Name    string
	Updated time.Time
	Visible bool
	Event   *DebugEvent
	Enabled bool
	Sets    []StatSet
}
type StatPoint struct {
	Total int
	Sum   float64
	Min   float64
	Max   float64
}
type StatSet struct {
	Index        int
	Description  string
	Interval     time.Duration
	PointerTime  time.Time
	PointerIndex int
	Points       []StatPoint
}
type DebugEvent struct {
	Id       int
	Name     string
	Parent   *DebugEvent
	Depth    int
	Children []*DebugEvent
	Sibling  *DebugEvent
}

type SpaceCoord interface {
	To2d() (x float32, y float32)
	To3d() (x float32, y float32, z float32)
}

type ReflectRegistry interface {
	Get(t reflect.Type) any
	Set(t reflect.Type, value any)
}

// times, ok := axe.Get[axe.Times](game)
func Get[V any](r ReflectRegistry) (V, bool) {
	t := TypeOf[V]()
	v := r.Get(t)
	if cast, ok := v.(V); ok {
		return cast, true
	}
	var empty V
	return empty, false
}

func TypeOf[V any]() reflect.Type {
	return reflect.TypeOf((*V)(nil)).Elem()
}

// type X[V any] func(value V) V
