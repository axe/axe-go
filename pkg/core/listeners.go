package core

type ListenerEntry[L any] struct {
	listener  L
	remaining int
	id        int
}

func (entry *ListenerEntry[L]) Off() {
	entry.remaining = 0
}

type ListenerOff func()

type ListenerOffs []ListenerOff

func (offs *ListenerOffs) Add(off ListenerOff) {
	*offs = append(*offs, off)
}

func (offs *ListenerOffs) Off() {
	if offs != nil {
		for _, off := range *offs {
			off()
		}
		*offs = (*offs)[:0]
	}
}

type Listeners[L any] struct {
	entries []ListenerEntry[L]
	nextId  int
}

func NewListeners[L any]() *Listeners[L] {
	return &Listeners[L]{
		entries: make([]ListenerEntry[L], 0),
		nextId:  0,
	}
}

func (l *Listeners[L]) Once(listener L) ListenerOff {
	return l.OnCount(listener, 1)
}

func (l *Listeners[L]) On(listener L) ListenerOff {
	return l.OnCount(listener, -1)
}

func (l *Listeners[L]) OnCount(listener L, count int) ListenerOff {
	entry := ListenerEntry[L]{
		listener:  listener,
		remaining: count,
		id:        l.nextId,
	}
	l.nextId++
	l.entries = append(l.entries, entry)

	return func() {
		for i, e := range l.entries {
			if e.id == entry.id {
				l.entries = append(l.entries[:i], l.entries[i+1:]...)
				break
			}
		}
	}
}

func (l *Listeners[L]) Trigger(call func(listener L) bool) int {
	triggered := 0
	for i := range l.entries {
		entry := &l.entries[i]
		if call(entry.listener) {
			triggered++
			if entry.remaining > 0 {
				entry.remaining--
			}
		}
	}
	alive := 0
	for i := range l.entries {
		entry := &l.entries[i]
		if entry.remaining != 0 {
			l.entries[alive] = l.entries[i]
			alive++
		}
	}
	l.entries = l.entries[:alive]
	return triggered
}
