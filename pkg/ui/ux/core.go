package ux

import "github.com/axe/axe-go/pkg/ui"

type HasComponent interface {
	GetComponent(theme *Theme) *ui.Base
}

type Kind int

const (
	KindButton Kind = iota
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

type Value[V any] interface {
	GetChanged() *V
	Set(value V)
}

func CoalesceValue[V any](value Value[V], constant V) Value[V] {
	if value != nil {
		return value
	}
	return NewConstant(constant)
}

var _ Value[int] = &Computed[int]{}
var _ Value[int] = &Constant[int]{}
var _ Value[int] = &Dynamic[int]{}

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

type Counter int

func (c *Counter) Add(amount int) { *c += Counter(amount) }
func (c *Counter) Changed() bool {
	changed := *c > 0
	*c = 0
	return changed
}
