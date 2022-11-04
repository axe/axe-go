package axe

import (
	"github.com/axe/axe-go/pkg/geom"
	"github.com/axe/axe-go/pkg/react"
	"github.com/axe/axe-go/pkg/ui"
)

type WindowSystem interface {
	GameSystem
	MainWindow() Window
	Windows() []Window
	Screens() []Screen
	Events() *Listeners[WindowSystemEvents]
}

type WindowSystemEvents struct {
	MouseScreenChange  func(oldMouse geom.Vec2i, oldScreen Screen, newMouse geom.Vec2i, newScreen Screen)
	ScreenConnected    func(newScreen Screen)
	ScreenDisconnected func(oldScreen Screen)
	ScreenResize       func(screen Screen, oldSize geom.Vec2i, newSize geom.Vec2i)
	WindowResize       func(window Window, oldSize geom.Vec2i, newSize geom.Vec2i)
	WindowAdded        func(window Window)
	WindowRemoved      func(window Window)
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
