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
type GameSystem interface {
	Init(game *Game) error
	Update(game *Game)
	Destroy()
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

type Game struct {
	Debug    DebugSystem
	Windows  WindowSystem
	Graphics GraphicsSystem
	Input    InputSystem
	Assets   AssetSystem
	Audio    AudioSystem
	Events   EventSystem
	Stages   StageManager
	State    GameState
	Running  bool
}

func (game *Game) Run() error {
	err := game.Init()
	if err != nil {
		return err
	}
	defer game.Destroy()

	for game.Running {
		err = game.Tick()
		if err != nil {
			return err
		}
	}

	return nil
}
func (game *Game) Init() error {
	err := game.Debug.Init(game)
	if err != nil {
		return err
	}
	err = game.Windows.Init(game)
	if err != nil {
		return err
	}
	err = game.Graphics.Init(game)
	if err != nil {
		return err
	}
	err = game.Input.Init(game)
	if err != nil {
		return err
	}
	err = game.Assets.Init(game)
	if err != nil {
		return err
	}
	err = game.Audio.Init(game)
	if err != nil {
		return err
	}
	err = game.Events.Init(game)
	if err != nil {
		return err
	}
	err = game.Stages.Init(game)
	if err != nil {
		return err
	}

	game.Running = true
	game.State.StartTime = time.Now()
	game.State.UpdateTimer.Reset()
	game.State.DrawTimer.Reset()
	return nil
}
func (game *Game) Destroy() {
	game.Stages.Destroy()
	game.Events.Destroy()
	game.Audio.Destroy()
	game.Assets.Destroy()
	game.Input.Destroy()
	game.Graphics.Destroy()
	game.Windows.Destroy()
	game.Debug.Destroy()
}
func (game *Game) Tick() error {
	doUpdate := game.State.UpdateTimer.Tick()

	game.Windows.Update(game)
	game.Input.Update(game)
	game.Assets.Update(game)
	if doUpdate {
		game.Stages.Update(game)
	}
	game.Audio.Update(game)
	game.Debug.Update(game)

	doDraw := game.State.DrawTimer.Tick()
	if doDraw {
		game.Graphics.Update(game)
	}

	sleepUpdate := game.State.UpdateTimer.NextTick()
	sleepDraw := game.State.DrawTimer.NextTick()
	sleep := sleepUpdate
	if sleepDraw < sleep {
		sleep = sleepDraw
	}
	if sleep > 0 {
		time.Sleep(sleep)
	}

	return nil
}

type Timer struct {
	LastTick  time.Time
	Current   time.Time
	Elapsed   time.Duration
	Frequency time.Duration
	Ticks     int64
}

func (e *Timer) Tick() bool {
	e.Current = time.Now()
	e.Elapsed = e.Current.Sub(e.LastTick)

	ticking := e.Elapsed >= e.Frequency
	if ticking {
		if e.Frequency == 0 {
			e.LastTick = e.Current
		} else {
			e.LastTick = e.LastTick.Add(e.Frequency)
		}
		e.Ticks++
	}
	return ticking
}
func (e *Timer) NextTick() time.Duration {
	return e.Frequency - e.Elapsed
}

func (e *Timer) Reset() {
	e.LastTick = time.Now()
	e.Current = e.LastTick
	e.Elapsed = 0
	e.Ticks = 0
}

type GameState struct {
	StartTime   time.Time
	UpdateTimer Timer
	DrawTimer   Timer
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

type AssetSystem struct {
	LoaderMap map[AssetType]AssetLoader
	Loaders   []AssetLoader
	Sources   []AssetSource
	Loaded    []Asset
}

var _ GameSystem = &AssetSystem{}

func (assets *AssetSystem) Init(game *Game) error { return nil }
func (assets *AssetSystem) Update(game *Game)     {}
func (assets *AssetSystem) Destroy()              {}

type AudioSystem interface { // & GameSystem
	GameSystem
}

type GraphicsSystem interface { // & GameSystem
	GameSystem
}

type Input interface {
	Data() InputData
}
type InputData struct {
	Name          string
	Value         float32
	ValueChanged  time.Time
	ValueDuration time.Duration
	Device        *InputDevice
}
type InputFilter struct {
	Input          Input
	Min            float32
	Max            float32
	NoiseReduction float32
}
type InputAny struct {
	Inputs []Input
}
type InputAll struct {
	Inputs []Input
}
type InputGesture struct {
	Inputs  []Input
	MinTime time.Duration
	Ignore  []Input
}
type InputVector struct {
	X      Input
	Y      Input
	Filter InputFilter
}
type InputTrigger struct {
	Start      float32
	End        float32
	MinElapsed time.Duration
}
type InputDevice struct {
	Name      string
	Type      string
	Inputs    []Input
	Connected react.Value[bool]
}
type InputSystem interface { // & GameSystem
	GameSystem
	Devices() []InputDevice
	Inputs() []Input
	Points() []InputPoint
	Events() *Listeners[InputSystemEvents]
}
type InputSystemEvents struct {
	DeviceConnected    func(newDevice InputDevice)
	DeviceDisconnected func(oldDevice InputDevice)
	InputConnected     func(newInput Input)
	InputDisconnected  func(oldInput Input)
	InputChange        func(input Input)
	PointConnected     func(newPoint InputPoint)
	PointDisconnected  func(oldPoint InputPoint)
	PointChange        func(point InputPoint)
	InputChangeMap     map[string]func(input Input)
}
type InputPoint struct {
	X      int
	Y      int
	Window *Window
	Screen *Screen
}
type InputAction struct {
	Name  string
	Input *Input
}
type InputActionSet struct {
	Actions []InputAction
	Enabled react.Value[bool]
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

type AssetType string
type AssetLoadStatus struct {
	Progress float32
	Loaded   bool
}
type AssetLoader interface {
	Handles(ref AssetRef) bool
	Load(asset Asset) AssetLoadStatus
	Types() []AssetType
}
type AssetSource interface {
	Handles(ref AssetRef) bool
	Create(ref AssetRef) Asset
}
type AssetRef struct {
	Name string
	URI  string
	Type AssetType
}
type Asset interface {
	Ref() AssetRef
	Source() AssetSource
	Status() AssetLoadStatus
	Loader() AssetLoader
	Data() any
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

type ListenerEntry[L any] struct {
	listener  L
	remaining int
	id        int
}

func (entry ListenerEntry[L]) Off() {
	entry.remaining = 0
}

type ListenerOff func()

type Listeners[L any] struct {
	entries []ListenerEntry[L]
	nextId  int
}

func NewListeners[L any]() *Listeners[L] {
	return &Listeners[L]{
		entries: make([]ListenerEntry[L], 0),
		nextId:  0,
	}
}

func (l *Listeners[L]) Once(listener L) ListenerOff {
	return l.OnCount(listener, 1)
}
func (l *Listeners[L]) On(listener L) ListenerOff {
	return l.OnCount(listener, -1)
}
func (l *Listeners[L]) OnCount(listener L, count int) ListenerOff {
	entry := ListenerEntry[L]{
		listener:  listener,
		remaining: count,
		id:        l.nextId,
	}
	l.nextId++
	l.entries = append(l.entries, entry)

	return func() {
		for i, e := range l.entries {
			if e.id == entry.id {
				l.entries = append(l.entries[:i], l.entries[i+1:]...)
				break
			}
		}
	}
}

func (l *Listeners[L]) Trigger(call func(listener L) bool) int {
	triggered := 0
	for i := range l.entries {
		entry := &l.entries[i]
		if call(entry.listener) {
			triggered++
			if entry.remaining > 0 {
				entry.remaining--
			}
		}
	}
	alive := 0
	for i := range l.entries {
		entry := &l.entries[i]
		if entry.remaining != 0 {
			l.entries[alive] = l.entries[i]
			alive++
		}
	}
	l.entries = l.entries[:alive]
	return triggered
}
