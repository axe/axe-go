package opengl

import (
	"fmt"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/core"
	"github.com/axe/axe-go/pkg/input"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewInputSystem() axe.InputSystem {
	return &inputSystem{
		System:    *input.NewSystem(),
		joysticks: make(map[glfw.Joystick]*input.Device),
		keys:      make(map[glfw.Key]*input.Input),
		buttons:   make(map[glfw.MouseButton]*input.Input),
		offs:      make(core.ListenerOffs, 0),
	}
}

type inputSystem struct {
	input.System
	joysticks map[glfw.Joystick]*input.Device
	keys      map[glfw.Key]*input.Input
	buttons   map[glfw.MouseButton]*input.Input
	offs      core.ListenerOffs
}

func (in *inputSystem) Init(game *axe.Game) error {
	in.SetInputTime(game.State.StartTime)
	in.initKeys(game)
	in.initMouse(game)
	in.initJoysticks(game)

	for _, w := range game.Windows.Windows() {
		in.listenToWindow(w)
	}
	off := game.Windows.Events().On(axe.WindowSystemEvents{
		WindowAdded: func(window axe.Window) {
			in.listenToWindow(window)
		},
		WindowRemoved: func(window axe.Window) {
			in.unlistenToWindow(window)
		},
	})
	in.offs.Add(off)

	return nil
}

func (in *inputSystem) initJoysticks(game *axe.Game) {
	for joy := glfw.Joystick1; joy < glfw.JoystickLast; joy++ {
		if joy.Present() {
			in.addJoystick(joy)
		}
	}

	glfw.SetJoystickCallback(func(joy glfw.Joystick, event glfw.PeripheralEvent) {
		if event == glfw.Connected {
			in.addJoystick(joy)
		} else if event == glfw.Disconnected {
			in.removeJoystick(joy)
		}
	})
}

func (in *inputSystem) removeJoystick(joy glfw.Joystick) {
	dev := in.joysticks[joy]

	if dev == nil {
		return
	}

	delete(in.joysticks, joy)
	in.DisconnectDevice(dev)
}

func (in *inputSystem) addJoystick(joy glfw.Joystick) {
	dev := input.NewDevice(joy.GetName(), input.DeviceTypeController)

	joyAxis := joy.GetAxes()
	if len(joyAxis) > 0 {
		for i := range joyAxis {
			dev.Add(fmt.Sprintf("Axis %d", i))
		}
	}

	joyButtons := joy.GetButtons()
	if len(joyButtons) > 0 {
		for i := range joyButtons {
			dev.Add(fmt.Sprintf("Button %d", i))
		}
	}

	joyHats := joy.GetHats()
	if len(joyHats) > 0 {
		for i := range joyHats {
			dev.Add(fmt.Sprintf("Hat %d X", i))
			dev.Add(fmt.Sprintf("Hat %d Y", i))
		}
	}

	in.joysticks[joy] = dev
	in.updateJoystick(joy)
	in.ConnectDevice(dev)
}

func (in *inputSystem) updateJoystick(joy glfw.Joystick) {
	dev := in.joysticks[joy]

	if dev == nil {
		return
	}

	inputIndex := 0

	joyAxis := joy.GetAxes()
	if len(joyAxis) > 0 {
		for _, axisValue := range joyAxis {
			axis := dev.Inputs[inputIndex]
			inputIndex++

			in.SetInputValue(axis, axisValue)
		}
	}

	joyButtons := joy.GetButtons()
	if len(joyButtons) > 0 {
		for _, buttonValue := range joyButtons {
			btn := dev.Inputs[inputIndex]
			inputIndex++

			in.onInputAction(btn, buttonValue)
		}
	}

	joyHats := joy.GetHats()
	if len(joyHats) > 0 {
		for _, hatValue := range joyHats {
			hatX := dev.Inputs[inputIndex]
			inputIndex++
			hatY := dev.Inputs[inputIndex]
			inputIndex++
			x, y := in.getHatXY(hatValue)

			in.SetInputValue(hatX, x)
			in.SetInputValue(hatY, y)
		}
	}
}

func (in *inputSystem) getHatXY(hat glfw.JoystickHatState) (float32, float32) {
	switch hat {
	case glfw.HatCentered:
		return 0, 0
	case glfw.HatDown:
		return 0, 1
	case glfw.HatLeft:
		return -1, 0
	case glfw.HatLeftDown:
		return -1, 1
	case glfw.HatLeftUp:
		return -1, -1
	case glfw.HatRight:
		return 1, 0
	case glfw.HatRightDown:
		return 1, 1
	case glfw.HatRightUp:
		return 1, -1
	case glfw.HatUp:
		return 0, -1
	}
	return 0, 0
}

func (in *inputSystem) initKeys(game *axe.Game) {
	keyboard := input.NewDevice("keyboard", input.DeviceTypeKeyboard)

	for key, keyName := range KEYS {
		keyCode := glfw.GetKeyScancode(key)
		if keyCode == -1 {
			continue
		}
		if in.Get(keyName) == nil {
			keyInput := keyboard.Add(keyName)
			in.keys[key] = keyInput
		}
	}

	in.ConnectDevice(keyboard)
}

func (in *inputSystem) initMouse(game *axe.Game) {
	mouse := input.NewDevice("mouse", input.DeviceTypeMouse)

	in.AddPoint(&input.Point{})

	for button := glfw.MouseButton1; button < glfw.MouseButtonLast; button++ {
		buttonName := fmt.Sprintf("MouseButton%d", button)
		buttonInput := mouse.Add(buttonName)
		in.buttons[button] = buttonInput
	}

	in.ConnectDevice(mouse)
}

func (in *inputSystem) listenToWindow(axeWindow axe.Window) {
	if win, ok := axeWindow.(*window); ok {
		glfwWindow := win.window
		glfwWindow.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			in.onInputAction(in.keys[key], action)
		})
		glfwWindow.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
			in.onInputAction(in.buttons[button], action)
		})
		glfwWindow.SetCursorPosCallback(func(w *glfw.Window, xpos, ypos float64) {
			in.SetInputPoint(in.Points()[0], float32(xpos), float32(ypos))
		})
		glfwWindow.SetCursorEnterCallback(func(w *glfw.Window, entered bool) {
			if entered {
				in.SetInputEnter(in.Points()[0])
			} else {
				in.SetInputLeave(in.Points()[0])
			}
		})
		for button, buttonInput := range in.buttons {
			if glfwWindow.GetMouseButton(button) == glfw.Press {
				buttonInput.Set(1, in.InputTime())
			}
		}
		for key, keyInput := range in.keys {
			if glfwWindow.GetKey(key) == glfw.Press {
				keyInput.Set(1, in.InputTime())
			}
		}
	}
}

func (in *inputSystem) unlistenToWindow(w axe.Window) {
	if win, ok := w.(*window); ok {
		glfwWindow := win.window
		glfwWindow.SetKeyCallback(nil)
		glfwWindow.SetMouseButtonCallback(nil)
		glfwWindow.SetCursorPosCallback(nil)
		glfwWindow.SetCursorEnterCallback(nil)
	}
}

func (in *inputSystem) onInputAction(ia *input.Input, action glfw.Action) {
	newValue := float32(1)
	if action == glfw.Release {
		newValue = 0
	}

	in.SetInputValue(ia, newValue)
}

func (in *inputSystem) Update(game *axe.Game) {
	in.SetInputTime(time.Now())

	glfw.PollEvents()

	for joy := range in.joysticks {
		in.updateJoystick(joy)
	}

	for _, input := range in.Inputs() {
		input.UpdateDuration(in.InputTime())
	}
}

func (in *inputSystem) Destroy() {
	in.offs.Off()
}

var KEYS = map[glfw.Key]string{
	glfw.KeySpace:        input.KeySpace,
	glfw.KeyApostrophe:   input.KeyApostrophe,
	glfw.KeyComma:        input.KeyComma,
	glfw.KeyMinus:        input.KeyMinus,
	glfw.KeyPeriod:       input.KeyPeriod,
	glfw.KeySlash:        input.KeySlash,
	glfw.Key0:            input.Key0,
	glfw.Key1:            input.Key1,
	glfw.Key2:            input.Key2,
	glfw.Key3:            input.Key3,
	glfw.Key4:            input.Key4,
	glfw.Key5:            input.Key5,
	glfw.Key6:            input.Key6,
	glfw.Key7:            input.Key7,
	glfw.Key8:            input.Key8,
	glfw.Key9:            input.Key9,
	glfw.KeySemicolon:    input.KeySemicolon,
	glfw.KeyEqual:        input.KeyEqual,
	glfw.KeyA:            input.KeyA,
	glfw.KeyB:            input.KeyB,
	glfw.KeyC:            input.KeyC,
	glfw.KeyD:            input.KeyD,
	glfw.KeyE:            input.KeyE,
	glfw.KeyF:            input.KeyF,
	glfw.KeyG:            input.KeyG,
	glfw.KeyH:            input.KeyH,
	glfw.KeyI:            input.KeyI,
	glfw.KeyJ:            input.KeyJ,
	glfw.KeyK:            input.KeyK,
	glfw.KeyL:            input.KeyL,
	glfw.KeyM:            input.KeyM,
	glfw.KeyN:            input.KeyN,
	glfw.KeyO:            input.KeyO,
	glfw.KeyP:            input.KeyP,
	glfw.KeyQ:            input.KeyQ,
	glfw.KeyR:            input.KeyR,
	glfw.KeyS:            input.KeyS,
	glfw.KeyT:            input.KeyT,
	glfw.KeyU:            input.KeyU,
	glfw.KeyV:            input.KeyV,
	glfw.KeyW:            input.KeyW,
	glfw.KeyX:            input.KeyX,
	glfw.KeyY:            input.KeyY,
	glfw.KeyZ:            input.KeyZ,
	glfw.KeyLeftBracket:  input.KeyLeftBracket,
	glfw.KeyBackslash:    input.KeyBackslash,
	glfw.KeyRightBracket: input.KeyRightBracket,
	glfw.KeyGraveAccent:  input.KeyGraveAccent,
	glfw.KeyWorld1:       input.KeyWorld1,
	glfw.KeyWorld2:       input.KeyWorld2,
	glfw.KeyEscape:       input.KeyEscape,
	glfw.KeyEnter:        input.KeyEnter,
	glfw.KeyTab:          input.KeyTab,
	glfw.KeyBackspace:    input.KeyBackspace,
	glfw.KeyInsert:       input.KeyInsert,
	glfw.KeyDelete:       input.KeyDelete,
	glfw.KeyRight:        input.KeyRight,
	glfw.KeyLeft:         input.KeyLeft,
	glfw.KeyDown:         input.KeyDown,
	glfw.KeyUp:           input.KeyUp,
	glfw.KeyPageUp:       input.KeyPageUp,
	glfw.KeyPageDown:     input.KeyPageDown,
	glfw.KeyHome:         input.KeyHome,
	glfw.KeyEnd:          input.KeyEnd,
	glfw.KeyCapsLock:     input.KeyCapsLock,
	glfw.KeyScrollLock:   input.KeyScrollLock,
	glfw.KeyNumLock:      input.KeyNumLock,
	glfw.KeyPrintScreen:  input.KeyPrintScreen,
	glfw.KeyPause:        input.KeyPause,
	glfw.KeyF1:           input.KeyF1,
	glfw.KeyF2:           input.KeyF2,
	glfw.KeyF3:           input.KeyF3,
	glfw.KeyF4:           input.KeyF4,
	glfw.KeyF5:           input.KeyF5,
	glfw.KeyF6:           input.KeyF6,
	glfw.KeyF7:           input.KeyF7,
	glfw.KeyF8:           input.KeyF8,
	glfw.KeyF9:           input.KeyF9,
	glfw.KeyF10:          input.KeyF10,
	glfw.KeyF11:          input.KeyF11,
	glfw.KeyF12:          input.KeyF12,
	glfw.KeyF13:          input.KeyF13,
	glfw.KeyF14:          input.KeyF14,
	glfw.KeyF15:          input.KeyF15,
	glfw.KeyF16:          input.KeyF16,
	glfw.KeyF17:          input.KeyF17,
	glfw.KeyF18:          input.KeyF18,
	glfw.KeyF19:          input.KeyF19,
	glfw.KeyF20:          input.KeyF20,
	glfw.KeyF21:          input.KeyF21,
	glfw.KeyF22:          input.KeyF22,
	glfw.KeyF23:          input.KeyF23,
	glfw.KeyF24:          input.KeyF24,
	glfw.KeyF25:          input.KeyF25,
	glfw.KeyKP0:          input.KeyKP0,
	glfw.KeyKP1:          input.KeyKP1,
	glfw.KeyKP2:          input.KeyKP2,
	glfw.KeyKP3:          input.KeyKP3,
	glfw.KeyKP4:          input.KeyKP4,
	glfw.KeyKP5:          input.KeyKP5,
	glfw.KeyKP6:          input.KeyKP6,
	glfw.KeyKP7:          input.KeyKP7,
	glfw.KeyKP8:          input.KeyKP8,
	glfw.KeyKP9:          input.KeyKP9,
	glfw.KeyKPDecimal:    input.KeyKPDecimal,
	glfw.KeyKPDivide:     input.KeyKPDivide,
	glfw.KeyKPMultiply:   input.KeyKPMultiply,
	glfw.KeyKPSubtract:   input.KeyKPSubtract,
	glfw.KeyKPAdd:        input.KeyKPAdd,
	glfw.KeyKPEnter:      input.KeyKPEnter,
	glfw.KeyKPEqual:      input.KeyKPEqual,
	glfw.KeyLeftShift:    input.KeyLeftShift,
	glfw.KeyLeftControl:  input.KeyLeftControl,
	glfw.KeyLeftAlt:      input.KeyLeftAlt,
	glfw.KeyLeftSuper:    input.KeyLeftSuper,
	glfw.KeyRightShift:   input.KeyRightShift,
	glfw.KeyRightControl: input.KeyRightControl,
	glfw.KeyRightAlt:     input.KeyRightAlt,
	glfw.KeyRightSuper:   input.KeyRightSuper,
	glfw.KeyMenu:         input.KeyMenu,
}
