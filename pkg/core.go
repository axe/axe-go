package axe

import (
	"time"

	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/geom"
	"github.com/axe/axe-go/pkg/job"
	"github.com/axe/axe-go/pkg/react"
	"github.com/axe/axe-go/pkg/ui"
)

type ProjectionOutside int

const (
	ProjectionOutsideIgnore ProjectionOutside = iota
	ProjectionOutsideClamp
	ProjectionOutsideRelative
)

type Vec[V any] interface {
	Get() V
	Set(value V)
}

type Scene[D Numeric, V Attr[V]] struct { // & GameSystem
	Name  string
	Jobs  job.JobRunner
	World ecs.World
	Space Space[D, V, SpaceComponent[D, V]]
}

type SpaceComponent[D Numeric, V Attr[V]] struct {
	Shape          Shape[D, V]
	Offset         V
	WorldTransform *Matrix[V]
	Flags          uint64
	Static         bool
	Inert          bool
}
type SpaceQuery[D Numeric, V Attr[V]] struct {
	Point   V
	End     V
	Shape   Shape[D, V]
	Maximum int
	Flags   uint64
	Match   ecs.FlagMatch
}
type SpaceNearest[E any] struct {
	Entity   E
	Distance float32
}
type SpaceSearchCallback[D Numeric, V Attr[V], E any] func(entity E, overlap float32, index int, query SpaceQuery[D, V]) bool
type SpaceCollisionCallback[E any] func(subject E, otherSubject E, overlap float32, index int, second bool)
type Space[D Numeric, V Attr[V], E any] interface {
	Collisions(flags uint32, match ecs.FlagMatch, callback SpaceCollisionCallback[E])
	Intersects(query SpaceQuery[D, V], callback SpaceSearchCallback[D, V, E]) int
	Contains(query SpaceQuery[D, V], callback SpaceSearchCallback[D, V, E]) int
	Raytrace(query SpaceQuery[D, V], callback SpaceSearchCallback[D, V, E]) int
	KNN(query SpaceQuery[D, V], nearest []SpaceNearest[E], nearestCount *int)
}

type Matrix[V Attr[V]] interface{}
type Plane[V Attr[V]] struct {
}
type Camera[V Attr[V]] interface {
	GameSystem
	Planes() []Plane[V]
	Intersects(shape Shape[float32, V], position V) bool
}
type RenderTarget interface { // window or texture
	Size() geom.Vec2i
}

type View[D Numeric, V Attr[V]] interface {
	GameSystem
	Name() string
	Scene() Scene[D, V]
	Camera() Camera[V]
	ProjectionMatrix() Matrix[V]
	ViewMatrix() Matrix[V]
	CombinedMatrix() Matrix[V]
	ProjectPoint(mouse geom.Vec2i, outside ProjectionOutside) V
	Project(outside ProjectionOutside) V
	ProjectIgnore() V
	UnprojectPoint(point V, outside ProjectionOutside) V
	UnprojectIgnore(point V) V
	Placement() ui.Placement
	Target() RenderTarget
	Draw()
}

type EventSystem struct {
}

var _ GameSystem = &EventSystem{}

func (events *EventSystem) Init(game *Game) error { return nil }
func (events *EventSystem) Update(game *Game)     {}
func (events *EventSystem) Destroy()              {}

type DebugSystem struct {
	Logs      []DebugLog
	Snapshots []DebugSnapshot
	Graphs    []DebugGraph
	Events    []DebugEvent
	Enabled   bool
}

var _ GameSystem = &DebugSystem{}

func (ds *DebugSystem) Trigger(ev DebugEvent) {}
func (ds *DebugSystem) Init(game *Game) error { return nil }
func (ds *DebugSystem) Update(game *Game)     {}
func (ds *DebugSystem) Destroy()              {}

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

type AudioSystem interface { // & GameSystem
	GameSystem
}

type GraphicsSystem interface { // & GameSystem
	GameSystem
}

type WindowSystem interface {
	GameSystem
	MainWindow() Window
	GetScreens() []Screen
	Events() *Listeners[WindowSystemEvents]
}
type WindowSystemEvents struct {
	MouseScreenChange  func(oldMouse geom.Vec2i, oldScreen Screen, newMouse geom.Vec2i, newScreen Screen)
	ScreenConnected    func(newScreen Screen)
	ScreenDisconnected func(oldScreen Screen)
	ScreenResize       func(screen Screen, oldSize geom.Vec2i, newSize geom.Vec2i)
	WindowResize       func(window Window, oldSize geom.Vec2i, newSize geom.Vec2i)
}

type StageManager struct {
	Stages  map[string]*Stage
	Current *Stage
	Next    *Stage

	events Listeners[StageManagerEvents]
}

var _ GameSystem = &StageManager{}

func (sm *StageManager) Events() *Listeners[StageManagerEvents] {
	return &sm.events
}
func (sm *StageManager) Init(game *Game) error { return nil }
func (sm *StageManager) Update(game *Game)     {}
func (sm *StageManager) Destroy()              {}

type StageManagerEvents struct {
	StageStarted func(current *Stage)
	StageExiting func(current *Stage, next *Stage)
	StageExited  func(previous *Stage, current *Stage)
}

type Stage struct {
	Name    string
	Assets  []AssetRef
	Windows []StageWindow
	Scenes  []Scene[float32, Vec2[float32]]
	Views   []View[float32, Vec2[float32]]
	Actions InputActionSets
}

type StageWindow struct {
	Name       string
	Title      string
	Placement  ui.Placement
	Fullscreen bool
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
type Calculator[T Attr[T]] interface {
	Add(a T, b T, out *T)
}

type Path[T Attr[T]] interface {
	Set(out *T, delta float32)
	PointCount() int
	Point(index int) T
	GetCalculator() Calculator[T]
	SetCalculator(calc Calculator[T])
}
