package input

import (
	"time"

	"github.com/axe/axe-go/pkg/react"
)

type Input struct {
	Name             string
	Value            float32
	ValueChanged     time.Time
	ValueDuration    time.Duration
	PreviousValue    float32
	PreviousChanged  time.Time
	PreviousDuration time.Duration
	Device           *Device
	Action           *Action
}

func New(name string) *Input {
	return &Input{Name: name}
}

func (i *Input) Set(value float32, now time.Time) bool {
	if i.Value != value {
		i.PreviousValue = i.Value
		i.PreviousChanged = i.ValueChanged
		i.PreviousDuration = now.Sub(i.ValueChanged)
		i.Value = value
		i.ValueChanged = now
		i.ValueDuration = 0
		return true
	}
	return false
}

func (i *Input) UpdateDuration(now time.Time) {
	i.ValueDuration = now.Sub(i.ValueChanged)
}

func (i *Input) Cancel() {
	i.Value = i.PreviousValue
	i.ValueChanged = i.PreviousChanged
}

func (i *Input) HasCancel() bool {
	return i.Value == i.PreviousValue && i.ValueChanged == i.PreviousChanged
}

type Data struct {
	Value  float32
	Inputs []*Input
}

func (data *Data) SetInputs(inputs []*Input) {
	if len(inputs) > 0 {
		data.Inputs = inputs
	} else if len(data.Inputs) > 0 {
		data.Inputs = data.Inputs[:0]
	}
}

func (data *Data) AddInputs(inputs []*Input) {
	if len(inputs) > 0 {
		if data.Inputs == nil {
			data.Inputs = inputs
		} else {
			data.Inputs = append(data.Inputs, inputs...)
		}
	}
}

type Gesture struct {
	Inputs  []*Input
	MinTime time.Duration
	Ignore  []*Input
}

type Vector struct {
	X *Input
	Y *Input
}

type Point struct {
	X int
	Y int
	// Window *Window
	// Screen *Screen
}

type DeviceType string

const (
	DeviceTypeKeyboard   DeviceType = "keyboard"
	DeviceTypeMouse      DeviceType = "mouse"
	DeviceTypeController DeviceType = "controller"
	DeviceTypeTouch      DeviceType = "touch"
)

type Device struct {
	Name      string
	Type      DeviceType
	Inputs    []*Input
	Connected react.Value[bool]
}

func (device *Device) Add(name string) *Input {
	in := New(name)
	in.Device = device
	device.Inputs = append(device.Inputs, in)
	return in
}

func NewDevice(name string, deviceType DeviceType) *Device {
	return &Device{
		Name:      name,
		Type:      deviceType,
		Inputs:    make([]*Input, 0, 64),
		Connected: react.Val(true),
	}
}

const (
	KeySpace        string = "space"
	KeyApostrophe   string = "apostrophe"
	KeyComma        string = "comma"
	KeyMinus        string = "minus"
	KeyPeriod       string = "period"
	KeySlash        string = "slash"
	Key0            string = "0"
	Key1            string = "1"
	Key2            string = "2"
	Key3            string = "3"
	Key4            string = "4"
	Key5            string = "5"
	Key6            string = "6"
	Key7            string = "7"
	Key8            string = "8"
	Key9            string = "9"
	KeySemicolon    string = "semicolon"
	KeyEqual        string = "equal"
	KeyA            string = "a"
	KeyB            string = "b"
	KeyC            string = "c"
	KeyD            string = "d"
	KeyE            string = "e"
	KeyF            string = "f"
	KeyG            string = "g"
	KeyH            string = "h"
	KeyI            string = "i"
	KeyJ            string = "j"
	KeyK            string = "k"
	KeyL            string = "l"
	KeyM            string = "m"
	KeyN            string = "n"
	KeyO            string = "o"
	KeyP            string = "p"
	KeyQ            string = "q"
	KeyR            string = "r"
	KeyS            string = "s"
	KeyT            string = "t"
	KeyU            string = "u"
	KeyV            string = "v"
	KeyW            string = "w"
	KeyX            string = "x"
	KeyY            string = "y"
	KeyZ            string = "z"
	KeyLeftBracket  string = "["
	KeyBackslash    string = "/"
	KeyRightBracket string = "]"
	KeyGraveAccent  string = "`"
	KeyWorld1       string = "world1"
	KeyWorld2       string = "world2"
	KeyEscape       string = "escape"
	KeyEnter        string = "enter"
	KeyTab          string = "tab"
	KeyBackspace    string = "backspace"
	KeyInsert       string = "insert"
	KeyDelete       string = "delete"
	KeyRight        string = "right"
	KeyLeft         string = "left"
	KeyDown         string = "down"
	KeyUp           string = "up"
	KeyPageUp       string = "page up"
	KeyPageDown     string = "page down"
	KeyHome         string = "home"
	KeyEnd          string = "end"
	KeyCapsLock     string = "caps lock"
	KeyScrollLock   string = "scroll lock"
	KeyNumLock      string = "num lock"
	KeyPrintScreen  string = "print"
	KeyPause        string = "pause"
	KeyF1           string = "f1"
	KeyF2           string = "f2"
	KeyF3           string = "f3"
	KeyF4           string = "f4"
	KeyF5           string = "f5"
	KeyF6           string = "f6"
	KeyF7           string = "f7"
	KeyF8           string = "f8"
	KeyF9           string = "f9"
	KeyF10          string = "f10"
	KeyF11          string = "f11"
	KeyF12          string = "f12"
	KeyF13          string = "f13"
	KeyF14          string = "f14"
	KeyF15          string = "f15"
	KeyF16          string = "f16"
	KeyF17          string = "f17"
	KeyF18          string = "f18"
	KeyF19          string = "f19"
	KeyF20          string = "f20"
	KeyF21          string = "f21"
	KeyF22          string = "f22"
	KeyF23          string = "f23"
	KeyF24          string = "f24"
	KeyF25          string = "f25"
	KeyKP0          string = "numpad 0"
	KeyKP1          string = "numpad 1"
	KeyKP2          string = "numpad 2"
	KeyKP3          string = "numpad 3"
	KeyKP4          string = "numpad 4"
	KeyKP5          string = "numpad 5"
	KeyKP6          string = "numpad 6"
	KeyKP7          string = "numpad 7"
	KeyKP8          string = "numpad 8"
	KeyKP9          string = "numpad 9"
	KeyKPDecimal    string = "numpad decimal"
	KeyKPDivide     string = "numpad divide"
	KeyKPMultiply   string = "numpad multiply"
	KeyKPSubtract   string = "numpad subtract"
	KeyKPAdd        string = "numpad add"
	KeyKPEnter      string = "numpad enter"
	KeyKPEqual      string = "numpad equal"
	KeyLeftShift    string = "left shift"
	KeyLeftControl  string = "left ctrl"
	KeyLeftAlt      string = "left alt"
	KeyLeftSuper    string = "left super"
	KeyRightShift   string = "right shift"
	KeyRightControl string = "right control"
	KeyRightAlt     string = "right alt"
	KeyRightSuper   string = "right super"
	KeyMenu         string = "menu"
)
