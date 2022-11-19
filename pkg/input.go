package axe

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/react"
	"github.com/axe/axe-go/pkg/util"
)

type Input struct {
	Name             string
	Value            float32
	ValueChanged     time.Time
	ValueDuration    time.Duration
	PreviousValue    float32
	PreviousChanged  time.Time
	PreviousDuration time.Duration
	Device           *InputDevice
	Action           *InputAction
}

func NewInput(name string) *Input {
	return &Input{Name: name}
}
func (i *Input) Set(value float32, now time.Time) {
	if i.Value != value {
		i.PreviousValue = i.Value
		i.PreviousChanged = i.ValueChanged
		i.PreviousDuration = now.Sub(i.ValueChanged)
		i.Value = value
		i.ValueChanged = now
		i.ValueDuration = 0
	}
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

type InputData struct {
	Value  float32
	Inputs []*Input
}

func (data *InputData) SetInputs(inputs []*Input) {
	if len(inputs) > 0 {
		data.Inputs = inputs
	} else if len(data.Inputs) > 0 {
		data.Inputs = data.Inputs[:0]
	}
}
func (data *InputData) AddInputs(inputs []*Input) {
	if len(inputs) > 0 {
		if data.Inputs == nil {
			data.Inputs = inputs
		} else {
			data.Inputs = append(data.Inputs, inputs...)
		}
	}
}

type InputAction struct {
	Name             string
	Enabled          react.Value[bool]
	Trigger          InputTrigger
	Data             InputData
	Triggered        bool
	OverridePriority int
}

func NewInputAction(name string, trigger InputTrigger) *InputAction {
	return &InputAction{
		Name:      name,
		Trigger:   trigger,
		Enabled:   react.Val(true),
		Data:      InputData{},
		Triggered: false,
	}
}

func (action *InputAction) String() string {
	return fmt.Sprintf("%s: %.1f", action.Name, action.Data.Value)
}

func (action *InputAction) Init(inputs InputSystem) {
	if action.Trigger != nil {
		action.Trigger.Init(inputs)
	}
}

func (action *InputAction) Update(inputs InputSystem) {
	if action.Enabled.Get() && action.Trigger != nil {
		data, triggered := action.Trigger.Update(inputs)

		if triggered && len(data.Inputs) > 0 {
			for _, in := range data.Inputs {
				if in.Action != nil {
					triggered = false
					break
				}
			}
			if triggered {
				for _, in := range data.Inputs {
					in.Action = action
				}
			}
		}

		action.Data = data
		action.Triggered = triggered
	}
}
func (action InputAction) Priority() int {
	if action.OverridePriority > 0 || action.Trigger == nil {
		return action.OverridePriority
	}
	return action.Trigger.InputCount()
}

type InputActions []*InputAction

func (actions InputActions) Less(i, j int) bool {
	return actions[i].Priority() > actions[j].Priority()
}
func (actions InputActions) Swap(i, j int) {
	t := actions[i]
	actions[i] = actions[j]
	actions[j] = t
}
func (actions InputActions) Len() int {
	return len(actions)
}
func (actions InputActions) Sort() {
	sort.Sort(actions)
}

type InputActionSet struct {
	Name      string
	Enabled   react.Value[bool]
	Actions   InputActions
	Triggered ds.CircularQueue[*InputAction]
}

func NewInputActionSet(name string) *InputActionSet {
	return &InputActionSet{
		Name:      name,
		Enabled:   react.Val(true),
		Actions:   make(InputActions, 0, 64),
		Triggered: ds.NewCircularQueue[*InputAction](64),
	}
}

func CreateInputActionSet(name string, actions map[string]InputTrigger) *InputActionSet {
	set := NewInputActionSet(name)
	for name, trigger := range actions {
		set.Add(NewInputAction(name, trigger))
	}
	return set
}

func (set *InputActionSet) Init(inputs InputSystem) {
	for _, action := range set.Actions {
		action.Init(inputs)
	}

	set.Actions.Sort()
}
func (set *InputActionSet) Update(inputs InputSystem) {
	set.Triggered.Clear()

	if !set.Enabled.Get() {
		return
	}
	if set.Actions != nil {
		for i := range set.Actions {
			action := set.Actions[i]
			action.Update(inputs)

			if action.Triggered {
				set.Triggered.Push(action)
			}
		}
	}
}

func (set *InputActionSet) Add(action *InputAction) {
	set.Actions = append(set.Actions, action)
}

func (set *InputActionSet) Iterator() ds.Iterator[*InputAction] {
	return set.Triggered.Iterator()
}

type InputActionHandler func(action *InputAction)

type InputActionSets struct {
	Sets    map[string]*InputActionSet
	Handler InputActionHandler
}

func NewInputActionSets() InputActionSets {
	return InputActionSets{
		Sets:    make(map[string]*InputActionSet),
		Handler: nil,
	}
}

func CreateInputActionSets(actionSets map[string]map[string]InputTrigger) InputActionSets {
	sets := NewInputActionSets()
	for name, actions := range actionSets {
		sets.Sets[name] = CreateInputActionSet(name, actions)
	}
	return sets
}

func (sets *InputActionSets) Init(inputs InputSystem) {
	if sets.Sets == nil {
		return
	}
	for _, set := range sets.Sets {
		set.Init(inputs)
	}
}
func (sets *InputActionSets) Update(inputs InputSystem) {
	if sets.Sets == nil {
		return
	}
	for _, input := range inputs.Inputs() {
		input.Action = nil
	}
	for _, set := range sets.Sets {
		set.Update(inputs)
		if sets.Handler != nil {
			triggeredIterator := set.Triggered.Iterator()
			for triggeredIterator.HasNext() {
				triggered := triggeredIterator.Next()
				sets.Handler(*triggered)
			}
		}
	}
}
func (sets *InputActionSets) Add(set *InputActionSet) {
	sets.Sets[set.Name] = set
}

func (sets *InputActionSets) Iterable() ds.Iterable[*InputAction] {
	return ds.NewMultiIterable(
		util.SliceMap(util.MapValues(sets.Sets), func(s *InputActionSet) ds.Iterable[*InputAction] {
			return s
		}),
	)
}

type InputTrigger interface {
	Init(inputs InputSystem)
	Update(inputs InputSystem) (InputData, bool)
	InputCount() int
}

type InputTriggerNamed struct {
	Name  string
	input *Input
}

var _ InputTrigger = &InputTriggerNamed{}

func (i *InputTriggerNamed) Init(inputs InputSystem) {
	i.input = inputs.Get(i.Name)
}
func (i *InputTriggerNamed) Update(inputs InputSystem) (InputData, bool) {
	if i.input == nil || i.input.ValueChanged.IsZero() {
		return InputData{}, false
	}
	return InputData{Value: i.input.Value, Inputs: []*Input{i.input}}, true
}
func (i *InputTriggerNamed) InputCount() int {
	if i.input == nil {
		return 0
	}
	return 1
}

type InputTriggerRange struct {
	TriggerMin        float32
	TriggerMinInclude bool
	TriggerMax        float32
	TriggerMaxInclude bool
	Trigger           InputTrigger
}

var _ InputTrigger = &InputTriggerRange{}

func (i *InputTriggerRange) Init(inputs InputSystem) {
	if i.Trigger != nil {
		i.Trigger.Init(inputs)
	}
}
func (i *InputTriggerRange) Update(inputs InputSystem) (InputData, bool) {
	if i.Trigger == nil {
		return InputData{}, false
	}
	value, triggered := i.Trigger.Update(inputs)
	if !triggered {
		return value, false
	}
	minSatisfied := (value.Value > i.TriggerMin || (value.Value >= i.TriggerMin && i.TriggerMinInclude))
	maxSatisfied := (value.Value < i.TriggerMax || (value.Value <= i.TriggerMax && i.TriggerMaxInclude))
	return value, minSatisfied && maxSatisfied
}

func (i *InputTriggerRange) InputCount() int {
	if i.Trigger == nil {
		return 0
	}
	return i.Trigger.InputCount()
}

type InputTriggerGroupOp string

const (
	InputTriggerGroupOpMin     InputTriggerGroupOp = "min"
	InputTriggerGroupOpMax     InputTriggerGroupOp = "max"
	InputTriggerGroupOpAverage InputTriggerGroupOp = "mavg"
)

type InputTriggerGroup struct {
	Op       InputTriggerGroupOp
	All      bool
	Triggers []InputTrigger
}

var _ InputTrigger = &InputTriggerGroup{}

func (i *InputTriggerGroup) Init(inputs InputSystem) {
	if i.Triggers != nil {
		for _, trigger := range i.Triggers {
			trigger.Init(inputs)
		}
	}
}
func (i *InputTriggerGroup) Update(inputs InputSystem) (InputData, bool) {
	if i.Triggers == nil {
		return InputData{}, false
	}
	triggeredCount := 0
	var group InputData
	for _, trigger := range i.Triggers {
		data, triggered := trigger.Update(inputs)
		if !triggered && i.All {
			return group, false
		}
		if triggered {
			triggeredCount++

			if triggeredCount == 1 {
				group = data
			} else {
				use := false

				switch i.Op {
				case InputTriggerGroupOpMin:
					if data.Value < group.Value {
						group.Value = data.Value
						use = true
					}
				case InputTriggerGroupOpMax:
					if data.Value > group.Value {
						group.Value = data.Value
						use = true
					}
				case InputTriggerGroupOpAverage:
					group.Value += data.Value
					use = true
				}

				if i.All || i.Op == InputTriggerGroupOpAverage {
					group.AddInputs(data.Inputs)
				} else if use {
					group.SetInputs(data.Inputs)
				}
			}
		}
	}
	if i.Op == InputTriggerGroupOpAverage && triggeredCount > 0 {
		group.Value /= float32(triggeredCount)
	}
	return group, triggeredCount > 0
}

func (i *InputTriggerGroup) InputCount() int {
	size := 0
	if i.Triggers != nil {
		for _, trigger := range i.Triggers {
			size += trigger.InputCount()
		}
	}
	return size
}

type InputTriggerValue struct {
	Value   float32
	Epsilon float32
	Trigger InputTrigger
}

var _ InputTrigger = &InputTriggerValue{}

func (i *InputTriggerValue) Init(inputs InputSystem) {
	if i.Trigger != nil {
		i.Trigger.Init(inputs)
	}
}
func (i *InputTriggerValue) Update(inputs InputSystem) (InputData, bool) {
	if i.Trigger == nil {
		return InputData{}, false
	}
	value, triggered := i.Trigger.Update(inputs)
	if !triggered {
		return value, false
	}
	return value, (math.Abs(float64(value.Value-i.Value)) <= float64(i.Epsilon))
}

func (i *InputTriggerValue) InputCount() int {
	if i.Trigger == nil {
		return 0
	}
	return i.Trigger.InputCount()
}

type InputTriggerOnce struct {
	Trigger   InputTrigger
	triggered bool
}

var _ InputTrigger = &InputTriggerOnce{}

func (i *InputTriggerOnce) Init(inputs InputSystem) {
	if i.Trigger != nil {
		i.Trigger.Init(inputs)
	}
}
func (i *InputTriggerOnce) Update(inputs InputSystem) (InputData, bool) {
	if i.Trigger == nil {
		return InputData{}, false
	}
	previouslyTriggered := i.triggered
	value, triggered := i.Trigger.Update(inputs)
	i.triggered = triggered
	return value, triggered && !previouslyTriggered
}

func (i *InputTriggerOnce) InputCount() int {
	if i.Trigger == nil {
		return 0
	}
	return i.Trigger.InputCount()
}

type InputTriggerInterval struct {
	Trigger       InputTrigger
	FirstInterval time.Duration
	Interval      time.Duration
	Strict        bool

	triggered      time.Time
	triggeredCount int
}

var _ InputTrigger = &InputTriggerInterval{}

func (i *InputTriggerInterval) Init(inputs InputSystem) {
	if i.Trigger != nil {
		i.Trigger.Init(inputs)
	}
}
func (i *InputTriggerInterval) Update(inputs InputSystem) (InputData, bool) {
	if i.Trigger == nil {
		return InputData{}, false
	}
	value, triggered := i.Trigger.Update(inputs)
	onInterval := false
	if triggered {
		now := inputs.InputTime()
		interval := i.Interval
		if i.triggeredCount == 1 && i.FirstInterval != 0 {
			interval = i.FirstInterval
		}
		if now.Sub(i.triggered) >= interval {
			i.triggered = now
			i.triggeredCount++
			onInterval = true
		}
	} else {
		i.triggeredCount = 0
		if !i.Strict {
			i.triggered = time.Time{}
		}
	}
	return value, onInterval
}

func (i *InputTriggerInterval) InputCount() int {
	if i.Trigger == nil {
		return 0
	}
	return i.Trigger.InputCount()
}

type InputTriggerNone struct{}

var _ InputTrigger = &InputTriggerNone{}

func (i *InputTriggerNone) Init(inputs InputSystem)                     {}
func (i *InputTriggerNone) Update(inputs InputSystem) (InputData, bool) { return InputData{}, false }
func (i *InputTriggerNone) InputCount() int                             { return 0 }

// TODO future
// InputTriggerFiltered: InputTrigger with min, max, noise reduction
// InputTriggerGesture

type InputTriggerBuilder struct{}

func (builder *InputTriggerBuilder) Once(trigger InputTrigger) InputTrigger {
	return &InputTriggerOnce{Trigger: trigger}
}
func (builder *InputTriggerBuilder) Named(name string) InputTrigger {
	return &InputTriggerNamed{Name: name}
}
func (builder *InputTriggerBuilder) And(triggers ...InputTrigger) InputTrigger {
	return &InputTriggerGroup{
		Op:       InputTriggerGroupOpMin,
		All:      true,
		Triggers: triggers,
	}
}
func (builder *InputTriggerBuilder) Or(triggers ...InputTrigger) InputTrigger {
	return &InputTriggerGroup{
		Op:       InputTriggerGroupOpMax,
		All:      false,
		Triggers: triggers,
	}
}
func (builder *InputTriggerBuilder) Value(value float32, trigger InputTrigger) InputTrigger {
	return &InputTriggerValue{
		Value:   value,
		Trigger: trigger,
	}
}
func (builder *InputTriggerBuilder) On(name string) InputTrigger {
	return &InputTriggerValue{
		Value:   1.0,
		Trigger: &InputTriggerNamed{Name: name},
	}
}
func (builder *InputTriggerBuilder) Off(name string) InputTrigger {
	return &InputTriggerValue{
		Value:   0,
		Trigger: &InputTriggerNamed{Name: name},
	}
}
func (builder *InputTriggerBuilder) Interval(interval time.Duration, trigger InputTrigger) InputTrigger {
	return &InputTriggerInterval{
		Interval: interval,
		Trigger:  trigger,
	}
}
func (builder *InputTriggerBuilder) FirstInterval(firstInterval time.Duration, interval time.Duration, trigger InputTrigger) InputTrigger {
	return &InputTriggerInterval{
		FirstInterval: firstInterval,
		Interval:      interval,
		Trigger:       trigger,
	}
}
func (builder *InputTriggerBuilder) Key(key InputKeyTrigger) InputTrigger {
	return NewInputKeyTrigger(key)
}

type InputKeyTrigger struct {
	Key             string
	Keys            []string
	Shift           bool
	Ctrl            bool
	Alt             bool
	CmdCtrl         bool
	Cmd             bool
	PressInterval   time.Duration
	FirstPressDelay time.Duration
	Up              bool
	UpOnly          bool
	Down            bool
}

func NewInputKeyTrigger(trigger InputKeyTrigger) InputTrigger {
	itb := InputTriggerBuilder{}
	keys := make([]InputTrigger, 0)

	if trigger.Key != "" {
		keys = append(keys, itb.Named(trigger.Key))
	}
	if len(trigger.Keys) > 0 {
		for _, key := range trigger.Keys {
			keys = append(keys, itb.Named(key))
		}
	}
	if trigger.Shift {
		keys = append(keys, itb.Or(itb.Named(InputKeyLeftShift), itb.Named(InputKeyRightShift)))
	}
	if trigger.Alt {
		keys = append(keys, itb.Or(itb.Named(InputKeyLeftAlt), itb.Named(InputKeyRightAlt)))
	}
	if trigger.CmdCtrl {
		keys = append(keys, itb.Or(itb.Named(InputKeyLeftControl), itb.Named(InputKeyRightControl), itb.Named(InputKeyLeftSuper), itb.Named(InputKeyRightSuper)))
	} else if trigger.Ctrl {
		keys = append(keys, itb.Or(itb.Named(InputKeyLeftControl), itb.Named(InputKeyRightControl)))
	} else if trigger.Cmd {
		keys = append(keys, itb.Or(itb.Named(InputKeyLeftSuper), itb.Named(InputKeyRightSuper)))
	}
	if len(keys) == 0 {
		return &InputTriggerNone{}
	}
	allKeys := keys[0]
	if len(keys) > 1 {
		allKeys = itb.And(keys...)
	}

	top := make([]InputTrigger, 0)
	if trigger.PressInterval > 0 {
		if trigger.FirstPressDelay > 0 {
			top = append(top, itb.FirstInterval(trigger.PressInterval+trigger.FirstPressDelay, trigger.PressInterval, itb.Value(1, allKeys)))
		} else {
			top = append(top, itb.Interval(trigger.PressInterval, itb.Value(1, allKeys)))
		}
	} else if trigger.Down {
		top = append(top, itb.Value(1, allKeys))
	} else if !trigger.UpOnly {
		top = append(top, itb.Once(itb.Value(1, allKeys)))
	}
	if trigger.Up || trigger.UpOnly {
		top = append(top, itb.Once(itb.Value(0, allKeys)))
	}

	if len(top) == 1 {
		return top[0]
	} else {
		return itb.Or(top...)
	}
}

type InputGesture struct {
	Inputs  []*Input
	MinTime time.Duration
	Ignore  []*Input
}
type InputVector struct {
	X *Input
	Y *Input
}
type InputPoint struct {
	X      int
	Y      int
	Window *Window
	Screen *Screen
}
type InputDeviceType string

const (
	InputDeviceTypeKeyboard   InputDeviceType = "keyboard"
	InputDeviceTypeMouse      InputDeviceType = "mouse"
	InputDeviceTypeController InputDeviceType = "controller"
	InputDeviceTypeTouch      InputDeviceType = "touch"
)

type InputDevice struct {
	Name      string
	Type      InputDeviceType
	Inputs    []*Input
	Connected react.Value[bool]
}

func (device *InputDevice) Add(name string) *Input {
	in := NewInput(name)
	in.Device = device
	device.Inputs = append(device.Inputs, in)
	return in
}

func NewInputDevice(name string, deviceType InputDeviceType) *InputDevice {
	return &InputDevice{
		Name:      name,
		Type:      deviceType,
		Inputs:    make([]*Input, 0, 64),
		Connected: react.Val(true),
	}
}

type InputSystem interface { // & GameSystem
	GameSystem
	Devices() []*InputDevice
	Inputs() []*Input
	Points() []*InputPoint
	Events() *Listeners[InputSystemEvents]
	Get(inputName string) *Input
	InputTime() time.Time
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

const (
	InputKeySpace        string = "space"
	InputKeyApostrophe   string = "apostrophe"
	InputKeyComma        string = "comma"
	InputKeyMinus        string = "minus"
	InputKeyPeriod       string = "period"
	InputKeySlash        string = "slash"
	InputKey0            string = "0"
	InputKey1            string = "1"
	InputKey2            string = "2"
	InputKey3            string = "3"
	InputKey4            string = "4"
	InputKey5            string = "5"
	InputKey6            string = "6"
	InputKey7            string = "7"
	InputKey8            string = "8"
	InputKey9            string = "9"
	InputKeySemicolon    string = "semicolon"
	InputKeyEqual        string = "equal"
	InputKeyA            string = "a"
	InputKeyB            string = "b"
	InputKeyC            string = "c"
	InputKeyD            string = "d"
	InputKeyE            string = "e"
	InputKeyF            string = "f"
	InputKeyG            string = "g"
	InputKeyH            string = "h"
	InputKeyI            string = "i"
	InputKeyJ            string = "j"
	InputKeyK            string = "k"
	InputKeyL            string = "l"
	InputKeyM            string = "m"
	InputKeyN            string = "n"
	InputKeyO            string = "o"
	InputKeyP            string = "p"
	InputKeyQ            string = "q"
	InputKeyR            string = "r"
	InputKeyS            string = "s"
	InputKeyT            string = "t"
	InputKeyU            string = "u"
	InputKeyV            string = "v"
	InputKeyW            string = "w"
	InputKeyX            string = "x"
	InputKeyY            string = "y"
	InputKeyZ            string = "z"
	InputKeyLeftBracket  string = "["
	InputKeyBackslash    string = "/"
	InputKeyRightBracket string = "]"
	InputKeyGraveAccent  string = "`"
	InputKeyWorld1       string = "world1"
	InputKeyWorld2       string = "world2"
	InputKeyEscape       string = "escape"
	InputKeyEnter        string = "enter"
	InputKeyTab          string = "tab"
	InputKeyBackspace    string = "backspace"
	InputKeyInsert       string = "insert"
	InputKeyDelete       string = "delete"
	InputKeyRight        string = "right"
	InputKeyLeft         string = "left"
	InputKeyDown         string = "down"
	InputKeyUp           string = "up"
	InputKeyPageUp       string = "page up"
	InputKeyPageDown     string = "page down"
	InputKeyHome         string = "home"
	InputKeyEnd          string = "end"
	InputKeyCapsLock     string = "caps lock"
	InputKeyScrollLock   string = "scroll lock"
	InputKeyNumLock      string = "num lock"
	InputKeyPrintScreen  string = "print"
	InputKeyPause        string = "pause"
	InputKeyF1           string = "f1"
	InputKeyF2           string = "f2"
	InputKeyF3           string = "f3"
	InputKeyF4           string = "f4"
	InputKeyF5           string = "f5"
	InputKeyF6           string = "f6"
	InputKeyF7           string = "f7"
	InputKeyF8           string = "f8"
	InputKeyF9           string = "f9"
	InputKeyF10          string = "f10"
	InputKeyF11          string = "f11"
	InputKeyF12          string = "f12"
	InputKeyF13          string = "f13"
	InputKeyF14          string = "f14"
	InputKeyF15          string = "f15"
	InputKeyF16          string = "f16"
	InputKeyF17          string = "f17"
	InputKeyF18          string = "f18"
	InputKeyF19          string = "f19"
	InputKeyF20          string = "f20"
	InputKeyF21          string = "f21"
	InputKeyF22          string = "f22"
	InputKeyF23          string = "f23"
	InputKeyF24          string = "f24"
	InputKeyF25          string = "f25"
	InputKeyKP0          string = "numpad 0"
	InputKeyKP1          string = "numpad 1"
	InputKeyKP2          string = "numpad 2"
	InputKeyKP3          string = "numpad 3"
	InputKeyKP4          string = "numpad 4"
	InputKeyKP5          string = "numpad 5"
	InputKeyKP6          string = "numpad 6"
	InputKeyKP7          string = "numpad 7"
	InputKeyKP8          string = "numpad 8"
	InputKeyKP9          string = "numpad 9"
	InputKeyKPDecimal    string = "numpad decimal"
	InputKeyKPDivide     string = "numpad divide"
	InputKeyKPMultiply   string = "numpad multiply"
	InputKeyKPSubtract   string = "numpad subtract"
	InputKeyKPAdd        string = "numpad add"
	InputKeyKPEnter      string = "numpad enter"
	InputKeyKPEqual      string = "numpad equal"
	InputKeyLeftShift    string = "left shift"
	InputKeyLeftControl  string = "left ctrl"
	InputKeyLeftAlt      string = "left alt"
	InputKeyLeftSuper    string = "left super"
	InputKeyRightShift   string = "right shift"
	InputKeyRightControl string = "right control"
	InputKeyRightAlt     string = "right alt"
	InputKeyRightSuper   string = "right super"
	InputKeyMenu         string = "menu"
)
