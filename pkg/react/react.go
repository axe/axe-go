package react

// A reactive value that can be watched and possibly set.
type Value[T any] interface {
	// Gets the value. This may be a variable or computed value. When this is called
	// within a watch or computed function it ties the value to that function - telling
	// it that when this value changes that function can re-execute to get an updated value.
	Get() T
	// Sets the value and returns true if successful. Computed values cannot be set. If set
	// is called on a value with the same value (by == comparison) it has no affect.
	Set(value T) bool
	// Detaches the value from the reactivity system.
	Detach()
}

// A watcher allows control over a watched function and the function invoked when the watched
// function has a new value. A watcher can be made lazy which means it only processes on Get
// when there's an outdate value.
type Watcher[T any] interface {
	// Stops watching the current values. Watching can be resumed with Update.
	Stop()
	// Updates the watcher forcefully, calling the get function and update function.
	Update()
	// Returns whether this watcher is lazy. A lazy watcher only processes on Get.
	IsLazy() bool
	// Sets whether this watcher is lazy. A lazy watcher only processes on Get.
	SetLazy(lazy bool)
	// Gets the latest value. If this is a lazy watcher and it has an outdated value it
	// will recompute the new value and call the passed update function.
	Get() T
}

var watchersLive map[int]watcher = map[int]watcher{}
var watcherIds int
var valueIds int

// Creates a reactive watcher. The get function is called and all reactive values
// referenced within are tracked. When any of those values are changed get is
// called again. Any time get is called and a new value is returned updated will be
// called as well. The returned watcher can be used to stop watching, resume, be
// marked as lazy, and get the most recent value.
func Watch[T any](get func() T, updated func(value T)) Watcher[T] {
	w := &watcherImpl[T]{
		watcherId: watcherIds,
		get:       get,
		updated:   updated,
		values:    make(map[int]value),
		lazy:      false,
		dirty:     false,
	}
	watcherIds++
	w.Update()
	return w
}

// Creates a reactive value. This value can be referenced in watchers and computed values
// and can trigger those watch functions or computed values to recalculate.
func Val[T comparable](value T) Value[T] {
	v := &valueImpl[T]{
		valueId:  valueIds,
		value:    value,
		watchers: make(map[int]watcher),
	}
	valueIds++
	return v
}

// Creates a computed value. A computed value is lazily calculated on Get. If no dependent
// reactive values have changed it's not recomputed.
func Computed[T comparable](get func() T) Value[T] {
	var empty T
	value := Val(empty)
	watcher := Watch(get, nil)
	watcher.SetLazy(true)
	return &computedImpl[T]{value, watcher}
}

type watcher interface {
	id() int
	notify()
	addValue(v value)
	removeValue(v value)
}

type value interface {
	id() int
	removeWatcher(w watcher)
}

type valueImpl[T comparable] struct {
	valueId  int
	value    T
	watchers map[int]watcher
}

var _ value = &valueImpl[int]{}
var _ Value[int] = &valueImpl[int]{}

func (v valueImpl[T]) id() int {
	return v.valueId
}

func (v *valueImpl[T]) removeWatcher(w watcher) {
	delete(v.watchers, w.id())
}

func (v *valueImpl[T]) Get() T {
	for _, watcher := range watchersLive {
		watcher.addValue(v)
		v.watchers[watcher.id()] = watcher
	}
	return v.value
}

func (v *valueImpl[T]) Set(value T) bool {
	changed := value != v.value
	if changed {
		v.value = value
		for _, n := range v.watchers {
			n.notify()
		}
	}
	return changed
}

func (v *valueImpl[T]) Detach() {
	v.watchers = make(map[int]watcher)
}

type watcherImpl[T any] struct {
	watcherId int
	lastValue T
	lazy      bool
	dirty     bool
	get       func() T
	updated   func(value T)
	values    map[int]value
}

var _ watcher = &watcherImpl[any]{}
var _ Watcher[any] = &watcherImpl[any]{}

func (w watcherImpl[T]) id() int {
	return w.watcherId
}

func (wn *watcherImpl[T]) notify() {
	wn.Update()
}

func (w *watcherImpl[T]) addValue(v value) {
	w.values[v.id()] = v
}

func (w *watcherImpl[T]) removeValue(v value) {
	delete(w.values, v.id())
}

func (w *watcherImpl[T]) updateLastValue() {
	w.Stop()
	watchersLive[w.watcherId] = w
	w.lastValue = w.get()
	w.dirty = false
	delete(watchersLive, w.watcherId)
	if w.updated != nil {
		w.updated(w.lastValue)
	}
}

func (w watcherImpl[T]) IsLazy() bool {
	return w.lazy
}

func (w watcherImpl[T]) SetLazy(lazy bool) {
	w.lazy = lazy
}

func (w *watcherImpl[T]) Stop() {
	for key, v := range w.values {
		v.removeWatcher(w)
		delete(w.values, key)
	}
}

func (w *watcherImpl[T]) Update() {
	if w.lazy {
		w.dirty = true
	} else {
		w.updateLastValue()
	}
}

func (w watcherImpl[T]) Get() T {
	if w.dirty {
		w.updateLastValue()
	}
	return w.lastValue
}

type computedImpl[T any] struct {
	value   Value[T]
	watcher Watcher[T]
}

var _ Value[any] = &computedImpl[any]{}

func (c computedImpl[T]) Set(value T) bool {
	return false
}

func (c computedImpl[T]) Get() T {
	watched := c.watcher.Get()
	c.value.Set(watched)
	return c.value.Get()
}

func (c computedImpl[T]) Detach() {
	c.value.Detach()
	c.watcher.Stop()
}
