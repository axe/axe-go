package ux

import (
	"github.com/axe/axe-go/pkg/input"
	"github.com/axe/axe-go/pkg/ui"
	"github.com/axe/axe-go/pkg/util"
)

type HasComponent interface {
	GetComponent(theme *Theme) *ui.Base
}

type Kind int

const (
	KindDefault Kind = iota
	KindNone
	KindButton
	KindCheckbox
	KindEditable
	KindScrollable
	KindDropdown
	KindInputText
	KindInputDropdown
	KindInputArea
	KindInputNumber
	KindCollapsible
	KindDialog
)

func (k Kind) Get(defaultKind Kind) Kind {
	if k == KindDefault {
		return defaultKind
	}
	return k
}

type Base struct {
	*ui.Template
	Kind Kind
}

// A value that can be tracked for changes.
type Value[V any] interface {
	// If the value has changed since the last time this method was called
	// a non-nil value will be returned pointing to the new value.
	GetChanged() *V
	// Sets the value.
	Set(value V)
}

// Returns a Value[V] given an optional value and its constant alternative
// if a value was not given.
func CoalesceValue[V any](value Value[V], constant V) Value[V] {
	if value != nil {
		return value
	}
	return NewConstant(constant)
}

var _ Value[int] = &Computed[int]{}
var _ Value[int] = &Constant[int]{}
var _ Value[int] = &Dynamic[int]{}
var _ Value[int] = &Live[int]{}

// A value that is changed by set.
type Constant[V any] struct {
	value   V
	changed bool
}

func NewConstant[V any](value V) *Constant[V] {
	return &Constant[V]{value: value, changed: true}
}

func (c *Constant[V]) GetChanged() *V {
	if !c.changed {
		return nil
	}
	c.changed = false
	return &c.value
}
func (c *Constant[V]) Set(value V) {
	c.value = value
	c.changed = true
}
func (c *Constant[V]) Changed() bool {
	return c.changed
}

// A value that is computed when GetChanged is invoked and it returns
// a new value then last time. Set has no effect.
type Computed[V comparable] struct {
	lastValue V
	getValue  func() V
}

func NewComputed[V comparable](getValue func() V) *Computed[V] {
	return &Computed[V]{getValue: getValue}
}

func (w *Computed[V]) GetChanged() *V {
	newValue := w.getValue()
	if w.lastValue == newValue {
		return nil
	}
	w.lastValue = newValue
	return &w.lastValue
}
func (w *Computed[V]) Set(value V) {
}

// A value that is computed when GetChanged is invoked but also
// supports setting.
type Dynamic[V comparable] struct {
	lastValue V
	getValue  func() V
	setValue  func(V)
}

func NewDynamic[V comparable](get func() V, set func(V)) *Dynamic[V] {
	return &Dynamic[V]{getValue: get, setValue: set}
}

func (d *Dynamic[V]) GetChanged() *V {
	newValue := d.getValue()
	if d.lastValue == newValue {
		return nil
	}
	d.lastValue = newValue
	return &d.lastValue
}
func (d *Dynamic[V]) Set(value V) {
	d.setValue(value)
}

// A value that is expected to change every time GetChanged is invoked.
type Live[V any] struct {
	getValue func() V
	setValue func(V)
}

func NewLive[V any](get func() V, set func(V)) *Live[V] {
	return &Live[V]{getValue: get, setValue: set}
}

func (d *Live[V]) GetChanged() *V {
	value := d.getValue()
	return &value
}
func (d *Live[V]) Set(value V) {
	d.setValue(value)
}

// A counter for the number of times something occurs.
type Counter int

func (c *Counter) Add(amount int) { *c += Counter(amount) }
func (c *Counter) Changed() bool {
	changed := *c > 0
	*c = 0
	return changed
}

// A trigger used for input actions.
type Trigger float32

func (t *Trigger) Set(value float32)            { *t = Trigger(value) }
func (t Trigger) Init(inputs input.InputSystem) {}
func (t Trigger) Update(inputs input.InputSystem) (input.Data, bool) {
	return input.Data{Value: float32(t)}, true
}
func (t Trigger) InputCount() int { return 0 }

// A listener is a function where multiple can be added.
type Listener[E any] func(ev E)

func (l Listener[E]) Trigger(ev E) {
	if l != nil {
		l(ev)
	}
}

func listenerNil[E any](a Listener[E]) bool {
	return a == nil
}
func listenerJoin[E any](first Listener[E], second Listener[E]) Listener[E] {
	return func(ev E) {
		first(ev)
		second(ev)
	}
}

func (a *Listener[E]) Add(b Listener[E], before bool) {
	*a = util.CoalesceJoin(*a, b, before, listenerNil[E], listenerJoin[E])
}

func (a *Listener[E]) Clear() {
	*a = nil
}
