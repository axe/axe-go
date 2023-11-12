package opengl

import (
	"sort"

	"github.com/go-gl/glfw/v3.3/glfw"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/core"
	"github.com/axe/axe-go/pkg/geom"
	"github.com/axe/axe-go/pkg/react"
	"github.com/axe/axe-go/pkg/ui"
	"github.com/axe/axe-go/pkg/util"
)

func NewWindowSystem(settings Settings) *windowSystem {
	return &windowSystem{
		settings: settings,
		screens:  make([]axe.Screen, 0),
		events:   core.NewListeners[axe.WindowSystemEvents](),
		windows:  make([]axe.Window, 0),
		loaded:   make(map[string]*window),
	}
}

type windowSystem struct {
	settings Settings
	main     *window
	loaded   map[string]*window
	windows  []axe.Window
	events   *core.Listeners[axe.WindowSystemEvents]
	screens  []axe.Screen
}

var _ axe.WindowSystem = &windowSystem{}

func (ws *windowSystem) MainWindow() axe.Window                          { return ws.main }
func (ws *windowSystem) Windows() []axe.Window                           { return ws.windows }
func (ws *windowSystem) Screens() []axe.Screen                           { return ws.screens }
func (ws *windowSystem) Events() *core.Listeners[axe.WindowSystemEvents] { return ws.events }

func (ws *windowSystem) Init(game *axe.Game) error {
	err := glfw.Init()
	if err != nil {
		return err
	}

	primary := glfw.GetPrimaryMonitor()
	monitors := glfw.GetMonitors()

	for i := range monitors {
		monitor := monitors[i]
		x, y, w, h := monitor.GetWorkarea()
		ws.screens = append(ws.screens, &screen{
			size:     geom.Vec2i{X: int32(w), Y: int32(h)},
			position: geom.Vec2i{X: int32(x), Y: int32(y)},
			primary:  primary.GetName() == monitor.GetName(),
			monitor:  monitor,
		})
	}

	sort.Slice(ws.screens, func(i, j int) bool {
		return ws.screens[i].Position().X < ws.screens[j].Position().X
	})

	game.Stages.Events().On(axe.StageManagerEvents{
		StageStarting: ws.addStageWindows,
		StageExited:   ws.removeDroppedWindows,
	})

	stage := game.Stages.GetStage(game)
	if stage != nil {
		ws.addWindows(stage.Windows)
	} else {
		panic("no stage ready, cannot create a window")
	}
	if len(ws.windows) == 0 {
		ws.addWindows(game.Settings.Windows)
	}

	return nil
}

func (ws *windowSystem) Update(game *axe.Game) {
	if ws.main.window.ShouldClose() {
		game.Running = false
	}
}

func (ws *windowSystem) Destroy() {
	glfw.Terminate()
}

func (ws *windowSystem) addStageWindows(stage *axe.Stage) {
	ws.addWindows(stage.Windows)
}

func (ws *windowSystem) addWindows(stageWindows []axe.StageWindow) {
	for _, win := range stageWindows {
		if ws.loaded[win.Name] != nil {
			continue
		}
		ws.createWindow(win)
	}
}

func (ws *windowSystem) createWindow(stageWindow axe.StageWindow) *window {
	screen := ws.getScreen(stageWindow.Screen)

	win := newWindow()
	var fullscreen *glfw.Monitor

	if !stageWindow.ClearColor.IsZero() {
		win.clear = stageWindow.ClearColor
	}
	if stageWindow.Placement.Defined() {
		win.placement = stageWindow.Placement
	}
	if stageWindow.Title != "" {
		win.title.Set(stageWindow.Title)
	}
	if stageWindow.Mode == axe.WindowModeFixed {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	if stageWindow.Mode == axe.WindowModeBorderless {
		glfw.WindowHint(glfw.Decorated, glfw.False)
	}
	if stageWindow.Mode == axe.WindowModeFullscreen {
		fullscreen = screen.monitor
		win.placement = ui.Maximized()
	}

	winBounds := win.placement.GetBoundsi(float32(screen.size.X), float32(screen.size.Y))
	win.size.X = winBounds.Width()
	win.size.Y = winBounds.Height()

	ws.settings.apply()

	gwin, err := glfw.CreateWindow(int(win.size.X), int(win.size.Y), win.title.Get(), fullscreen, nil)
	if err != nil {
		panic(err)
	}
	gwin.MakeContextCurrent()
	gwin.SetPos(int(winBounds.Min.X), int(winBounds.Min.Y))
	win.window = gwin
	win.screen = screen

	if ws.main == nil {
		ws.main = win
	}

	ws.windows = append(ws.windows, win)
	ws.loaded[stageWindow.Name] = win

	gwin.SetSizeCallback(func(w *glfw.Window, width, height int) {
		if width == 0 || height == 0 {
			return
		}
		ws.events.Trigger(func(listener axe.WindowSystemEvents) bool {
			if listener.WindowResize != nil {
				oldSize := win.size
				win.size.X = int32(width)
				win.size.Y = int32(height)
				listener.WindowResize(win, oldSize, win.size)
				return true
			}
			return false
		})
	})

	ws.events.Trigger(func(listener axe.WindowSystemEvents) bool {
		if listener.WindowAdded != nil {
			listener.WindowAdded(win)
			return true
		}
		return false
	})

	return win
}

func (ws *windowSystem) removeDroppedWindows(prev *axe.Stage, curr *axe.Stage) {
	if prev == nil {
		return
	}

	skip := map[string]bool{}
	for _, win := range curr.Windows {
		skip[win.Name] = true
	}

	for _, win := range prev.Windows {
		if !skip[win.Name] && ws.loaded[win.Name] != nil {
			close := ws.loaded[win.Name]

			ws.events.Trigger(func(listener axe.WindowSystemEvents) bool {
				if listener.WindowRemoved != nil {
					listener.WindowRemoved(close)
					return true
				}
				return false
			})

			close.window.Destroy()

			delete(ws.loaded, win.Name)
			ws.windows = util.SliceRemove[axe.Window](ws.windows, close)
		}
	}
}

func (ws *windowSystem) getScreen(relativeIndex int) *screen {
	primaryIndex := 0
	for i, s := range ws.screens {
		if s.(*screen).primary {
			primaryIndex = i
			break
		}
	}
	desired := primaryIndex + relativeIndex
	actual := core.Clamp(desired, 0, len(ws.screens)-1)
	return ws.screens[actual].(*screen)
}

type window struct {
	name      string
	title     react.Value[string]
	placement ui.Placement
	window    *glfw.Window
	size      geom.Vec2i
	clear     ui.Color
	screen    *screen
}

var _ axe.Window = &window{}

func newWindow() *window {
	place := ui.Centered(512, 512)

	return &window{
		title:     react.Val(""),
		placement: place,
		size:      geom.Vec2i{X: 512, Y: 512},
		clear:     ui.ColorBlack,
	}
}

func (w *window) Name() string               { return w.name }
func (w *window) Title() react.Value[string] { return w.title }
func (w *window) Placement() ui.Placement    { return w.placement }
func (w *window) Screen() axe.Screen         { return w.screen }
func (w *window) Size() geom.Vec2i           { return w.size }

type screen struct {
	size     geom.Vec2i
	position geom.Vec2i
	primary  bool
	monitor  *glfw.Monitor
}

var _ axe.Screen = &screen{}

func (s *screen) Size() geom.Vec2i     { return s.size }
func (s *screen) Position() geom.Vec2i { return s.position }
