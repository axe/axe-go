package axe

import "time"

type Timer struct {
	LastTick  time.Time
	Current   time.Time
	Elapsed   time.Duration
	Frequency time.Duration
	Ticks     int64
}

func (e *Timer) Tick() bool {
	e.Current = time.Now()
	e.Elapsed = e.Current.Sub(e.LastTick)

	ticking := e.Elapsed >= e.Frequency
	if ticking {
		if e.Frequency == 0 {
			e.LastTick = e.Current
		} else {
			e.LastTick = e.LastTick.Add(e.Frequency)
		}
		e.Ticks++
	}
	return ticking
}
func (e *Timer) NextTick() time.Duration {
	return e.Frequency - e.Elapsed
}

func (e *Timer) Reset() {
	e.LastTick = time.Now()
	e.Current = e.LastTick
	e.Elapsed = 0
	e.Ticks = 0
}
