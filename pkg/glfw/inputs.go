package glfw

import (
	"fmt"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/util"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewInputSystem() axe.InputSystem {
	return &inputSystem{
		events:    axe.NewListeners[axe.InputSystemEvents](),
		inputMap:  make(map[string]*axe.Input),
		inputs:    make([]*axe.Input, 0),
		devices:   make([]*axe.InputDevice, 0),
		joysticks: make(map[glfw.Joystick]*axe.InputDevice),
		points:    make([]*axe.InputPoint, 0),
		keys:      make(map[glfw.Key]*axe.Input),
		buttons:   make(map[glfw.MouseButton]*axe.Input),
		offs:      make(axe.ListenerOffs, 0),
	}
}

type inputSystem struct {
	events    *axe.Listeners[axe.InputSystemEvents]
	inputs    []*axe.Input
	devices   []*axe.InputDevice
	joysticks map[glfw.Joystick]*axe.InputDevice
	inputMap  map[string]*axe.Input
	points    []*axe.InputPoint
	now       time.Time
	keys      map[glfw.Key]*axe.Input
	buttons   map[glfw.MouseButton]*axe.Input
	offs      axe.ListenerOffs
}

func (in *inputSystem) Init(game *axe.Game) error {
	in.now = game.State.StartTime
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
func (in *inputSystem) deviceConnected(newDevice *axe.InputDevice) {
	in.devices = append(in.devices, newDevice)

	in.events.Trigger(func(listener axe.InputSystemEvents) bool {
		if listener.DeviceConnected != nil {
			listener.DeviceConnected(*newDevice)
			return true
		}
		return false
	})

	for _, input := range newDevice.Inputs {
		in.inputConnected(input)
	}
}
func (in *inputSystem) inputConnected(newInput *axe.Input) {
	in.inputs = append(in.inputs, newInput)
	in.inputMap[newInput.Name] = newInput

	in.events.Trigger(func(listener axe.InputSystemEvents) bool {
		if listener.InputConnected != nil {
			listener.InputConnected(*newInput)
			return true
		}
		return false
	})
}
func (in *inputSystem) deviceDisconnected(oldDevice *axe.InputDevice) {
	in.devices = util.SliceRemove(in.devices, oldDevice)

	in.events.Trigger(func(listener axe.InputSystemEvents) bool {
		if listener.DeviceDisconnected != nil {
			listener.DeviceDisconnected(*oldDevice)
			return true
		}
		return false
	})

	for _, input := range oldDevice.Inputs {
		in.inputDisconnected(input)
	}
}
func (in *inputSystem) inputDisconnected(oldInput *axe.Input) {
	in.inputs = util.SliceRemove(in.inputs, oldInput)
	if in.inputMap[oldInput.Name] == oldInput {
		delete(in.inputMap, oldInput.Name)
	}

	in.events.Trigger(func(listener axe.InputSystemEvents) bool {
		if listener.InputDisconnected != nil {
			listener.InputDisconnected(*oldInput)
			return true
		}
		return false
	})
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
	in.deviceDisconnected(dev)
}
func (in *inputSystem) addJoystick(joy glfw.Joystick) {
	dev := axe.NewInputDevice(joy.GetName(), axe.InputDeviceTypeController)

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
	in.deviceConnected(dev)
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

			in.onInputValue(axis, axisValue)
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

			in.onInputValue(hatX, x)
			in.onInputValue(hatY, y)
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
	keyboard := axe.NewInputDevice("keyboard", axe.InputDeviceTypeKeyboard)

	for key, keyName := range KEYS {
		keyCode := glfw.GetKeyScancode(key)
		if keyCode == -1 {
			continue
		}
		if in.inputMap[keyName] == nil {
			keyInput := keyboard.Add(keyName)
			in.keys[key] = keyInput
		}
	}

	in.deviceConnected(keyboard)
}
func (in *inputSystem) initMouse(game *axe.Game) {
	mouse := axe.NewInputDevice("mouse", axe.InputDeviceTypeMouse)

	in.points = append(in.points, &axe.InputPoint{})

	for button := glfw.MouseButton1; button < glfw.MouseButtonLast; button++ {
		buttonName := fmt.Sprintf("MouseButton%d", button)
		buttonInput := mouse.Add(buttonName)
		in.buttons[button] = buttonInput
	}

	in.deviceConnected(mouse)
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
			in.onInputPoint(in.points[0], int(xpos), int(ypos))
		})
		for button, buttonInput := range in.buttons {
			if glfwWindow.GetMouseButton(button) == glfw.Press {
				buttonInput.Set(1, in.now)
			}
		}
		for key, keyInput := range in.keys {
			if glfwWindow.GetKey(key) == glfw.Press {
				keyInput.Set(1, in.now)
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
	}
}
func (in *inputSystem) onInputValue(ia *axe.Input, newValue float32) {
	if ia != nil {
		if !ia.Set(newValue, in.now) {
			return
		}

		in.events.Trigger(func(listener axe.InputSystemEvents) bool {
			handled := false
			if listener.InputChange != nil {
				listener.InputChange(*ia)
				handled = true
			}
			if listener.InputChangeMap != nil {
				inputListener := listener.InputChangeMap[ia.Name]
				if inputListener != nil {
					inputListener(*ia)
					handled = true
				}
			}
			return handled
		})
	}
}
func (in *inputSystem) onInputAction(ia *axe.Input, action glfw.Action) {
	newValue := float32(1)
	if action == glfw.Release {
		newValue = 0
	}

	in.onInputValue(ia, newValue)
}
func (in *inputSystem) onInputPoint(ia *axe.InputPoint, x int, y int) {
	if ia != nil {
		ia.X = x
		ia.Y = y
		in.events.Trigger(func(listener axe.InputSystemEvents) bool {
			if listener.PointChange != nil {
				listener.PointChange(*ia)
				return true
			}
			return false
		})
	}
}
func (in *inputSystem) Update(game *axe.Game) {
	in.now = time.Now()

	glfw.PollEvents()

	for joy := range in.joysticks {
		in.updateJoystick(joy)
	}

	for _, input := range in.inputs {
		input.UpdateDuration(in.now)
	}
}
func (in *inputSystem) Destroy() {
	in.offs.Off()
}
func (in *inputSystem) Devices() []*axe.InputDevice                   { return nil }
func (in *inputSystem) Inputs() []*axe.Input                          { return in.inputs }
func (in *inputSystem) Get(name string) *axe.Input                    { return in.inputMap[name] }
func (in *inputSystem) Points() []*axe.InputPoint                     { return in.points }
func (in *inputSystem) Events() *axe.Listeners[axe.InputSystemEvents] { return in.events }
func (in *inputSystem) InputTime() time.Time                          { return in.now }

var KEYS map[glfw.Key]string = map[glfw.Key]string{
	glfw.KeySpace:        axe.InputKeySpace,
	glfw.KeyApostrophe:   axe.InputKeyApostrophe,
	glfw.KeyComma:        axe.InputKeyComma,
	glfw.KeyMinus:        axe.InputKeyMinus,
	glfw.KeyPeriod:       axe.InputKeyPeriod,
	glfw.KeySlash:        axe.InputKeySlash,
	glfw.Key0:            axe.InputKey0,
	glfw.Key1:            axe.InputKey1,
	glfw.Key2:            axe.InputKey2,
	glfw.Key3:            axe.InputKey3,
	glfw.Key4:            axe.InputKey4,
	glfw.Key5:            axe.InputKey5,
	glfw.Key6:            axe.InputKey6,
	glfw.Key7:            axe.InputKey7,
	glfw.Key8:            axe.InputKey8,
	glfw.Key9:            axe.InputKey9,
	glfw.KeySemicolon:    axe.InputKeySemicolon,
	glfw.KeyEqual:        axe.InputKeyEqual,
	glfw.KeyA:            axe.InputKeyA,
	glfw.KeyB:            axe.InputKeyB,
	glfw.KeyC:            axe.InputKeyC,
	glfw.KeyD:            axe.InputKeyD,
	glfw.KeyE:            axe.InputKeyE,
	glfw.KeyF:            axe.InputKeyF,
	glfw.KeyG:            axe.InputKeyG,
	glfw.KeyH:            axe.InputKeyH,
	glfw.KeyI:            axe.InputKeyI,
	glfw.KeyJ:            axe.InputKeyJ,
	glfw.KeyK:            axe.InputKeyK,
	glfw.KeyL:            axe.InputKeyL,
	glfw.KeyM:            axe.InputKeyM,
	glfw.KeyN:            axe.InputKeyN,
	glfw.KeyO:            axe.InputKeyO,
	glfw.KeyP:            axe.InputKeyP,
	glfw.KeyQ:            axe.InputKeyQ,
	glfw.KeyR:            axe.InputKeyR,
	glfw.KeyS:            axe.InputKeyS,
	glfw.KeyT:            axe.InputKeyT,
	glfw.KeyU:            axe.InputKeyU,
	glfw.KeyV:            axe.InputKeyV,
	glfw.KeyW:            axe.InputKeyW,
	glfw.KeyX:            axe.InputKeyX,
	glfw.KeyY:            axe.InputKeyY,
	glfw.KeyZ:            axe.InputKeyZ,
	glfw.KeyLeftBracket:  axe.InputKeyLeftBracket,
	glfw.KeyBackslash:    axe.InputKeyBackslash,
	glfw.KeyRightBracket: axe.InputKeyRightBracket,
	glfw.KeyGraveAccent:  axe.InputKeyGraveAccent,
	glfw.KeyWorld1:       axe.InputKeyWorld1,
	glfw.KeyWorld2:       axe.InputKeyWorld2,
	glfw.KeyEscape:       axe.InputKeyEscape,
	glfw.KeyEnter:        axe.InputKeyEnter,
	glfw.KeyTab:          axe.InputKeyTab,
	glfw.KeyBackspace:    axe.InputKeyBackspace,
	glfw.KeyInsert:       axe.InputKeyInsert,
	glfw.KeyDelete:       axe.InputKeyDelete,
	glfw.KeyRight:        axe.InputKeyRight,
	glfw.KeyLeft:         axe.InputKeyLeft,
	glfw.KeyDown:         axe.InputKeyDown,
	glfw.KeyUp:           axe.InputKeyUp,
	glfw.KeyPageUp:       axe.InputKeyPageUp,
	glfw.KeyPageDown:     axe.InputKeyPageDown,
	glfw.KeyHome:         axe.InputKeyHome,
	glfw.KeyEnd:          axe.InputKeyEnd,
	glfw.KeyCapsLock:     axe.InputKeyCapsLock,
	glfw.KeyScrollLock:   axe.InputKeyScrollLock,
	glfw.KeyNumLock:      axe.InputKeyNumLock,
	glfw.KeyPrintScreen:  axe.InputKeyPrintScreen,
	glfw.KeyPause:        axe.InputKeyPause,
	glfw.KeyF1:           axe.InputKeyF1,
	glfw.KeyF2:           axe.InputKeyF2,
	glfw.KeyF3:           axe.InputKeyF3,
	glfw.KeyF4:           axe.InputKeyF4,
	glfw.KeyF5:           axe.InputKeyF5,
	glfw.KeyF6:           axe.InputKeyF6,
	glfw.KeyF7:           axe.InputKeyF7,
	glfw.KeyF8:           axe.InputKeyF8,
	glfw.KeyF9:           axe.InputKeyF9,
	glfw.KeyF10:          axe.InputKeyF10,
	glfw.KeyF11:          axe.InputKeyF11,
	glfw.KeyF12:          axe.InputKeyF12,
	glfw.KeyF13:          axe.InputKeyF13,
	glfw.KeyF14:          axe.InputKeyF14,
	glfw.KeyF15:          axe.InputKeyF15,
	glfw.KeyF16:          axe.InputKeyF16,
	glfw.KeyF17:          axe.InputKeyF17,
	glfw.KeyF18:          axe.InputKeyF18,
	glfw.KeyF19:          axe.InputKeyF19,
	glfw.KeyF20:          axe.InputKeyF20,
	glfw.KeyF21:          axe.InputKeyF21,
	glfw.KeyF22:          axe.InputKeyF22,
	glfw.KeyF23:          axe.InputKeyF23,
	glfw.KeyF24:          axe.InputKeyF24,
	glfw.KeyF25:          axe.InputKeyF25,
	glfw.KeyKP0:          axe.InputKeyKP0,
	glfw.KeyKP1:          axe.InputKeyKP1,
	glfw.KeyKP2:          axe.InputKeyKP2,
	glfw.KeyKP3:          axe.InputKeyKP3,
	glfw.KeyKP4:          axe.InputKeyKP4,
	glfw.KeyKP5:          axe.InputKeyKP5,
	glfw.KeyKP6:          axe.InputKeyKP6,
	glfw.KeyKP7:          axe.InputKeyKP7,
	glfw.KeyKP8:          axe.InputKeyKP8,
	glfw.KeyKP9:          axe.InputKeyKP9,
	glfw.KeyKPDecimal:    axe.InputKeyKPDecimal,
	glfw.KeyKPDivide:     axe.InputKeyKPDivide,
	glfw.KeyKPMultiply:   axe.InputKeyKPMultiply,
	glfw.KeyKPSubtract:   axe.InputKeyKPSubtract,
	glfw.KeyKPAdd:        axe.InputKeyKPAdd,
	glfw.KeyKPEnter:      axe.InputKeyKPEnter,
	glfw.KeyKPEqual:      axe.InputKeyKPEqual,
	glfw.KeyLeftShift:    axe.InputKeyLeftShift,
	glfw.KeyLeftControl:  axe.InputKeyLeftControl,
	glfw.KeyLeftAlt:      axe.InputKeyLeftAlt,
	glfw.KeyLeftSuper:    axe.InputKeyLeftSuper,
	glfw.KeyRightShift:   axe.InputKeyRightShift,
	glfw.KeyRightControl: axe.InputKeyRightControl,
	glfw.KeyRightAlt:     axe.InputKeyRightAlt,
	glfw.KeyRightSuper:   axe.InputKeyRightSuper,
	glfw.KeyMenu:         axe.InputKeyMenu,
}
