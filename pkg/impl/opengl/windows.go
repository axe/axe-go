package opengl

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/core"
	"github.com/axe/axe-go/pkg/geom"
	"github.com/axe/axe-go/pkg/react"
	"github.com/axe/axe-go/pkg/ui"
)

func NewWindowSystem(settings Settings) *windowSystem {
	return &windowSystem{
		settings: settings,
		screens:  make([]axe.Screen, 0),
		events:   core.NewListeners[axe.WindowSystemEvents](),
		windows:  make([]axe.Window, 0),
	}
}

type windowSystem struct {
	settings Settings
	main     *window
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

	for _, monitor := range glfw.GetMonitors() {
		x, y, w, h := monitor.GetWorkarea()
		ws.screens = append(ws.screens, &screen{
			size:     geom.Vec2i{X: w, Y: h},
			position: geom.Vec2i{X: x, Y: y},
		})
	}

	primary := glfw.GetPrimaryMonitor()
	_, _, primaryWidth, primaryHeight := primary.GetWorkarea()

	win := newWindow()
	winBounds := win.placement.GetBoundsi(float32(primaryWidth), float32(primaryHeight))
	winWidth := winBounds.Max.X - winBounds.Min.X
	winHeight := winBounds.Min.Y - winBounds.Max.Y
	win.size.X = winWidth
	win.size.Y = winHeight

	// glfw.WindowHint(glfw.Resizable, glfw.False)
	ws.settings.apply()

	gwin, err := glfw.CreateWindow(winWidth, winHeight, game.Settings.Name, nil, nil)
	if err != nil {
		return err
	}
	gwin.MakeContextCurrent()
	gwin.SetPos(winBounds.Min.X, winBounds.Max.Y)
	win.window = gwin

	ws.main = win
	ws.windows = append(ws.windows, win)

	gwin.SetSizeCallback(func(w *glfw.Window, width, height int) {
		ws.events.Trigger(func(listener axe.WindowSystemEvents) bool {
			if listener.WindowResize != nil {
				oldSize := win.size
				win.size.X = width
				win.size.Y = height
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

type window struct {
	name      string
	title     react.Value[string]
	placement ui.Placement
	window    *glfw.Window
	size      geom.Vec2i
}

var _ axe.Window = &window{}

func newWindow() *window {
	place := ui.Placement{}
	place.Center(512, 512)

	return &window{
		title:     react.Val(""),
		placement: place,
		size:      geom.Vec2i{X: 512, Y: 512},
	}
}

func (w *window) Name() string               { return w.name }
func (w *window) Title() react.Value[string] { return w.title }
func (w *window) Placement() ui.Placement    { return w.placement }
func (w *window) Screen() axe.Screen         { return nil }
func (w *window) Size() geom.Vec2i           { return w.size }

type screen struct {
	size     geom.Vec2i
	position geom.Vec2i
}

var _ axe.Screen = &screen{}

func (s *screen) Size() geom.Vec2i     { return s.size }
func (s *screen) Position() geom.Vec2i { return s.position }
