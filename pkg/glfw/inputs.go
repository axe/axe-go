package glfw

import (
	"fmt"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewInputSystem() axe.InputSystem {
	return &inputSystem{
		events: axe.NewListeners[axe.InputSystemEvents](),
		inputs: make([]axe.Input, 0),
		points: make([]axe.InputPoint, 0),
	}
}

type inputSystem struct {
	events *axe.Listeners[axe.InputSystemEvents]
	inputs []axe.Input
	points []axe.InputPoint
}

func (in *inputSystem) Init(game *axe.Game) error {
	inputMap := make(map[string]*input)
	keys := make(map[glfw.Key]*input)
	buttons := make(map[glfw.MouseButton]*input)

	for key, keyName := range KEYS {
		keyCode := glfw.GetKeyScancode(key)
		if keyCode == -1 {
			continue
		}
		// keyName := glfw.GetKeyName(key, keyCode)

		if inputMap[keyName] == nil {
			keyInput := &input{
				data: axe.InputData{
					Name:          keyName,
					Value:         0,
					ValueChanged:  time.Time{},
					ValueDuration: 0,
				},
			}
			inputMap[keyName] = keyInput
			in.inputs = append(in.inputs, keyInput)
			keys[key] = keyInput
		}
	}

	for button := glfw.MouseButton1; button < glfw.MouseButtonLast; button++ {
		buttonName := fmt.Sprintf("MouseButton%d", button)
		buttonInput := &input{
			data: axe.InputData{
				Name:          buttonName,
				Value:         0,
				ValueChanged:  time.Time{},
				ValueDuration: 0,
			},
		}
		inputMap[buttonName] = buttonInput
		in.inputs = append(in.inputs, buttonInput)
		buttons[button] = buttonInput
	}

	in.points = append(in.points, axe.InputPoint{
		X: 0,
		Y: 0,
	})

	handleInputAction := func(ia *input, action glfw.Action) {
		if ia != nil {
			if action == glfw.Press {
				ia.Set(1)
			} else if action == glfw.Release {
				ia.Set(0)
			}
			in.events.Trigger(func(listener axe.InputSystemEvents) bool {
				handled := false
				if listener.InputChange != nil {
					listener.InputChange(ia)
					handled = true
				}
				if listener.InputChangeMap != nil {
					inputListener := listener.InputChangeMap[ia.data.Name]
					if inputListener != nil {
						inputListener(ia)
						handled = true
					}
				}
				return handled
			})
		}
	}
	handlePointAction := func(ia *axe.InputPoint, x int, y int) {
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

	if win, ok := game.Windows.MainWindow().(*window); ok {
		win.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			handleInputAction(keys[key], action)
		})
		win.window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
			handleInputAction(buttons[button], action)
		})
		win.window.SetCursorPosCallback(func(w *glfw.Window, xpos, ypos float64) {
			handlePointAction(&in.points[0], int(xpos), int(ypos))
		})
		for key, keyInput := range keys {
			if win.window.GetKey(key) == glfw.Press {
				keyInput.Set(1)
			}
		}
		for button, buttonInput := range buttons {
			if win.window.GetMouseButton(button) == glfw.Press {
				buttonInput.Set(1)
			}
		}
	}
	return nil
}
func (in *inputSystem) Update(game *axe.Game) {
	glfw.PollEvents()

	now := time.Now()

	for _, in := range in.inputs {
		i := in.(*input)
		i.data.ValueDuration = now.Sub(i.data.ValueChanged)
	}
}
func (in *inputSystem) Destroy() {

}
func (in *inputSystem) Devices() []axe.InputDevice                    { return nil }
func (in *inputSystem) Inputs() []axe.Input                           { return in.inputs }
func (in *inputSystem) Points() []axe.InputPoint                      { return in.points }
func (in *inputSystem) Events() *axe.Listeners[axe.InputSystemEvents] { return in.events }

type input struct {
	data axe.InputData
}

var _ axe.Input = &input{}

func (i *input) Data() axe.InputData {
	return i.data
}

func (i *input) Set(value float32) {
	if i.data.Value != value {
		i.data.Value = value
		i.data.ValueChanged = time.Now()
		i.data.ValueDuration = 0
	}
}

var KEYS map[glfw.Key]string = map[glfw.Key]string{
	glfw.KeySpace:        "space",
	glfw.KeyApostrophe:   "apostrophe",
	glfw.KeyComma:        "comma",
	glfw.KeyMinus:        "minus",
	glfw.KeyPeriod:       "period",
	glfw.KeySlash:        "slash",
	glfw.Key0:            "0",
	glfw.Key1:            "1",
	glfw.Key2:            "2",
	glfw.Key3:            "3",
	glfw.Key4:            "4",
	glfw.Key5:            "5",
	glfw.Key6:            "6",
	glfw.Key7:            "7",
	glfw.Key8:            "8",
	glfw.Key9:            "9",
	glfw.KeySemicolon:    "semicolon",
	glfw.KeyEqual:        "equal",
	glfw.KeyA:            "a",
	glfw.KeyB:            "b",
	glfw.KeyC:            "c",
	glfw.KeyD:            "d",
	glfw.KeyE:            "e",
	glfw.KeyF:            "f",
	glfw.KeyG:            "g",
	glfw.KeyH:            "h",
	glfw.KeyI:            "i",
	glfw.KeyJ:            "j",
	glfw.KeyK:            "k",
	glfw.KeyL:            "l",
	glfw.KeyM:            "m",
	glfw.KeyN:            "n",
	glfw.KeyO:            "o",
	glfw.KeyP:            "p",
	glfw.KeyQ:            "q",
	glfw.KeyR:            "r",
	glfw.KeyS:            "s",
	glfw.KeyT:            "t",
	glfw.KeyU:            "u",
	glfw.KeyV:            "v",
	glfw.KeyW:            "w",
	glfw.KeyX:            "x",
	glfw.KeyY:            "y",
	glfw.KeyZ:            "z",
	glfw.KeyLeftBracket:  "[",
	glfw.KeyBackslash:    "/",
	glfw.KeyRightBracket: "]",
	glfw.KeyGraveAccent:  "`",
	glfw.KeyWorld1:       "world1",
	glfw.KeyWorld2:       "world2",
	glfw.KeyEscape:       "escape",
	glfw.KeyEnter:        "enter",
	glfw.KeyTab:          "tab",
	glfw.KeyBackspace:    "backspace",
	glfw.KeyInsert:       "insert",
	glfw.KeyDelete:       "delete",
	glfw.KeyRight:        "right",
	glfw.KeyLeft:         "left",
	glfw.KeyDown:         "down",
	glfw.KeyUp:           "up",
	glfw.KeyPageUp:       "page up",
	glfw.KeyPageDown:     "page down",
	glfw.KeyHome:         "home",
	glfw.KeyEnd:          "end",
	glfw.KeyCapsLock:     "caps lock",
	glfw.KeyScrollLock:   "scroll lock",
	glfw.KeyNumLock:      "num lock",
	glfw.KeyPrintScreen:  "print",
	glfw.KeyPause:        "pause",
	glfw.KeyF1:           "f1",
	glfw.KeyF2:           "f2",
	glfw.KeyF3:           "f3",
	glfw.KeyF4:           "f4",
	glfw.KeyF5:           "f5",
	glfw.KeyF6:           "f6",
	glfw.KeyF7:           "f7",
	glfw.KeyF8:           "f8",
	glfw.KeyF9:           "f9",
	glfw.KeyF10:          "f10",
	glfw.KeyF11:          "f11",
	glfw.KeyF12:          "f12",
	glfw.KeyF13:          "f13",
	glfw.KeyF14:          "f14",
	glfw.KeyF15:          "f15",
	glfw.KeyF16:          "f16",
	glfw.KeyF17:          "f17",
	glfw.KeyF18:          "f18",
	glfw.KeyF19:          "f19",
	glfw.KeyF20:          "f20",
	glfw.KeyF21:          "f21",
	glfw.KeyF22:          "f22",
	glfw.KeyF23:          "f23",
	glfw.KeyF24:          "f24",
	glfw.KeyF25:          "f25",
	glfw.KeyKP0:          "numpad 0",
	glfw.KeyKP1:          "numpad 1",
	glfw.KeyKP2:          "numpad 2",
	glfw.KeyKP3:          "numpad 3",
	glfw.KeyKP4:          "numpad 4",
	glfw.KeyKP5:          "numpad 5",
	glfw.KeyKP6:          "numpad 6",
	glfw.KeyKP7:          "numpad 7",
	glfw.KeyKP8:          "numpad 8",
	glfw.KeyKP9:          "numpad 9",
	glfw.KeyKPDecimal:    "numpad decimal",
	glfw.KeyKPDivide:     "numpad divide",
	glfw.KeyKPMultiply:   "numpad multiply",
	glfw.KeyKPSubtract:   "numpad subtract",
	glfw.KeyKPAdd:        "numpad add",
	glfw.KeyKPEnter:      "numpad enter",
	glfw.KeyKPEqual:      "numpad equal",
	glfw.KeyLeftShift:    "left shift",
	glfw.KeyLeftControl:  "left ctrl",
	glfw.KeyLeftAlt:      "left alt",
	glfw.KeyLeftSuper:    "left super",
	glfw.KeyRightShift:   "right shift",
	glfw.KeyRightControl: "right control",
	glfw.KeyRightAlt:     "right alt",
	glfw.KeyRightSuper:   "right super",
	glfw.KeyMenu:         "menu",
}
