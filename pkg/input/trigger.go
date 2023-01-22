package input

import (
	"math"
	"time"
)

type Trigger interface {
	Init(inputs InputSystem)
	Update(inputs InputSystem) (Data, bool)
	InputCount() int
}

type TriggerNamed struct {
	Name  string
	input *Input
}

var _ Trigger = &TriggerNamed{}

func (i *TriggerNamed) Init(inputs InputSystem) {
	i.input = inputs.Get(i.Name)
}
func (i *TriggerNamed) Update(inputs InputSystem) (Data, bool) {
	if i.input == nil || i.input.ValueChanged.IsZero() {
		return Data{}, false
	}
	return Data{Value: i.input.Value, Inputs: []*Input{i.input}}, true
}
func (i *TriggerNamed) InputCount() int {
	if i.input == nil {
		return 0
	}
	return 1
}

type TriggerRange struct {
	TriggerMin        float32
	TriggerMinInclude bool
	TriggerMax        float32
	TriggerMaxInclude bool
	Trigger           Trigger
}

var _ Trigger = &TriggerRange{}

func (i *TriggerRange) Init(inputs InputSystem) {
	if i.Trigger != nil {
		i.Trigger.Init(inputs)
	}
}

func (i *TriggerRange) Update(inputs InputSystem) (Data, bool) {
	if i.Trigger == nil {
		return Data{}, false
	}
	value, triggered := i.Trigger.Update(inputs)
	if !triggered {
		return value, false
	}
	minSatisfied := (value.Value > i.TriggerMin || (value.Value >= i.TriggerMin && i.TriggerMinInclude))
	maxSatisfied := (value.Value < i.TriggerMax || (value.Value <= i.TriggerMax && i.TriggerMaxInclude))
	return value, minSatisfied && maxSatisfied
}

func (i *TriggerRange) InputCount() int {
	if i.Trigger == nil {
		return 0
	}
	return i.Trigger.InputCount()
}

type TriggerGroupOp string

const (
	TriggerGroupOpMin     TriggerGroupOp = "min"
	TriggerGroupOpMax     TriggerGroupOp = "max"
	TriggerGroupOpAverage TriggerGroupOp = "mavg"
)

type TriggerGroup struct {
	Op       TriggerGroupOp
	All      bool
	Triggers []Trigger
}

var _ Trigger = &TriggerGroup{}

func (i *TriggerGroup) Init(inputs InputSystem) {
	if i.Triggers != nil {
		for _, trigger := range i.Triggers {
			trigger.Init(inputs)
		}
	}
}

func (i *TriggerGroup) Update(inputs InputSystem) (Data, bool) {
	if i.Triggers == nil {
		return Data{}, false
	}
	triggeredCount := 0
	var group Data
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
				case TriggerGroupOpMin:
					if data.Value < group.Value {
						group.Value = data.Value
						use = true
					}
				case TriggerGroupOpMax:
					if data.Value > group.Value {
						group.Value = data.Value
						use = true
					}
				case TriggerGroupOpAverage:
					group.Value += data.Value
					use = true
				}

				if i.All || i.Op == TriggerGroupOpAverage {
					group.AddInputs(data.Inputs)
				} else if use {
					group.SetInputs(data.Inputs)
				}
			}
		}
	}
	if i.Op == TriggerGroupOpAverage && triggeredCount > 0 {
		group.Value /= float32(triggeredCount)
	}
	return group, triggeredCount > 0
}

func (i *TriggerGroup) InputCount() int {
	size := 0
	if i.Triggers != nil {
		for _, trigger := range i.Triggers {
			size += trigger.InputCount()
		}
	}
	return size
}

type TriggerValue struct {
	Value   float32
	Epsilon float32
	Trigger Trigger
}

var _ Trigger = &TriggerValue{}

func (i *TriggerValue) Init(inputs InputSystem) {
	if i.Trigger != nil {
		i.Trigger.Init(inputs)
	}
}

func (i *TriggerValue) Update(inputs InputSystem) (Data, bool) {
	if i.Trigger == nil {
		return Data{}, false
	}
	value, triggered := i.Trigger.Update(inputs)
	if !triggered {
		return value, false
	}
	return value, (math.Abs(float64(value.Value-i.Value)) <= float64(i.Epsilon))
}

func (i *TriggerValue) InputCount() int {
	if i.Trigger == nil {
		return 0
	}
	return i.Trigger.InputCount()
}

type TriggerOnce struct {
	Trigger   Trigger
	triggered bool
}

var _ Trigger = &TriggerOnce{}

func (i *TriggerOnce) Init(inputs InputSystem) {
	if i.Trigger != nil {
		i.Trigger.Init(inputs)
	}
}
func (i *TriggerOnce) Update(inputs InputSystem) (Data, bool) {
	if i.Trigger == nil {
		return Data{}, false
	}
	previouslyTriggered := i.triggered
	value, triggered := i.Trigger.Update(inputs)
	i.triggered = triggered
	return value, triggered && !previouslyTriggered
}

func (i *TriggerOnce) InputCount() int {
	if i.Trigger == nil {
		return 0
	}
	return i.Trigger.InputCount()
}

type TriggerInterval struct {
	Trigger       Trigger
	FirstInterval time.Duration
	Interval      time.Duration
	Strict        bool

	triggered      time.Time
	triggeredCount int
}

var _ Trigger = &TriggerInterval{}

func (i *TriggerInterval) Init(inputs InputSystem) {
	if i.Trigger != nil {
		i.Trigger.Init(inputs)
	}
}

func (i *TriggerInterval) Update(inputs InputSystem) (Data, bool) {
	if i.Trigger == nil {
		return Data{}, false
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

func (i *TriggerInterval) InputCount() int {
	if i.Trigger == nil {
		return 0
	}
	return i.Trigger.InputCount()
}

type TriggerNone struct{}

var _ Trigger = &TriggerNone{}

func (i *TriggerNone) Init(inputs InputSystem)                {}
func (i *TriggerNone) Update(inputs InputSystem) (Data, bool) { return Data{}, false }
func (i *TriggerNone) InputCount() int                        { return 0 }

type TriggerBuilder struct{}

func (builder *TriggerBuilder) Once(trigger Trigger) Trigger {
	return &TriggerOnce{Trigger: trigger}
}

func (builder *TriggerBuilder) Named(name string) Trigger {
	return &TriggerNamed{Name: name}
}

func (builder *TriggerBuilder) And(triggers ...Trigger) Trigger {
	return &TriggerGroup{
		Op:       TriggerGroupOpMin,
		All:      true,
		Triggers: triggers,
	}
}

func (builder *TriggerBuilder) Or(triggers ...Trigger) Trigger {
	return &TriggerGroup{
		Op:       TriggerGroupOpMax,
		All:      false,
		Triggers: triggers,
	}
}

func (builder *TriggerBuilder) Value(value float32, trigger Trigger) Trigger {
	return &TriggerValue{
		Value:   value,
		Trigger: trigger,
	}
}

func (builder *TriggerBuilder) On(name string) Trigger {
	return &TriggerValue{
		Value:   1.0,
		Trigger: &TriggerNamed{Name: name},
	}
}

func (builder *TriggerBuilder) Off(name string) Trigger {
	return &TriggerValue{
		Value:   0,
		Trigger: &TriggerNamed{Name: name},
	}
}

func (builder *TriggerBuilder) Interval(interval time.Duration, trigger Trigger) Trigger {
	return &TriggerInterval{
		Interval: interval,
		Trigger:  trigger,
	}
}

func (builder *TriggerBuilder) FirstInterval(firstInterval time.Duration, interval time.Duration, trigger Trigger) Trigger {
	return &TriggerInterval{
		FirstInterval: firstInterval,
		Interval:      interval,
		Trigger:       trigger,
	}
}

func (builder *TriggerBuilder) Key(key KeyTrigger) Trigger {
	return NewKeyTrigger(key)
}

type KeyTrigger struct {
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

func NewKeyTrigger(trigger KeyTrigger) Trigger {
	itb := TriggerBuilder{}
	keys := make([]Trigger, 0)

	if trigger.Key != "" {
		keys = append(keys, itb.Named(trigger.Key))
	}
	if len(trigger.Keys) > 0 {
		for _, key := range trigger.Keys {
			keys = append(keys, itb.Named(key))
		}
	}
	if trigger.Shift {
		keys = append(keys, itb.Or(itb.Named(KeyLeftShift), itb.Named(KeyRightShift)))
	}
	if trigger.Alt {
		keys = append(keys, itb.Or(itb.Named(KeyLeftAlt), itb.Named(KeyRightAlt)))
	}
	if trigger.CmdCtrl {
		keys = append(keys, itb.Or(itb.Named(KeyLeftControl), itb.Named(KeyRightControl), itb.Named(KeyLeftSuper), itb.Named(KeyRightSuper)))
	} else if trigger.Ctrl {
		keys = append(keys, itb.Or(itb.Named(KeyLeftControl), itb.Named(KeyRightControl)))
	} else if trigger.Cmd {
		keys = append(keys, itb.Or(itb.Named(KeyLeftSuper), itb.Named(KeyRightSuper)))
	}
	if len(keys) == 0 {
		return &TriggerNone{}
	}
	allKeys := keys[0]
	if len(keys) > 1 {
		allKeys = itb.And(keys...)
	}

	top := make([]Trigger, 0)
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
