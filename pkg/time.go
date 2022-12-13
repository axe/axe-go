package axe

import "time"

type TimeType uint8

const (
	TimeTypeReal TimeType = iota
	TimeTypeWorld
)

type Time struct {
	Name        string
	DayDuration time.Duration
	Enabled     bool
	Scale       float32
	DateTime    time.Time
	Elapsed     time.Duration
	StartTime   time.Time

	resume bool
}

var day = time.Hour * 24

func NewTime(name string) Time {
	now := time.Now()

	return Time{
		Name:        name,
		DayDuration: day,
		Enabled:     true,
		Scale:       1.0,
		Elapsed:     0,
		DateTime:    now,
		StartTime:   now,
	}
}

func (t *Time) Update() {
	now := time.Now()

	if t.resume {
		t.resume = false
		t.Enabled = true
		t.StartTime = now
	}

	if !t.Enabled {
		return
	}

	if t.DayDuration != day {
		now = t.ToTimeNow(now)
	}

	t.Elapsed = now.Sub(t.DateTime)
	t.DateTime = now
}

func (t *Time) Pause() {
	t.Enabled = false
}

func (t *Time) Resume() {
	t.resume = true
}

func (t *Time) ToTimeNow(now time.Time) time.Time {
	realElapsedSinceStart := now.Sub(t.StartTime)
	timeElapsedSinceStart := realElapsedSinceStart * day / t.DayDuration
	timeNow := t.StartTime.Add(timeElapsedSinceStart)

	return timeNow
}

func (t *Time) ElapsedScaled() time.Duration {
	if !t.Enabled {
		return 0
	}
	if t.Scale == 1 {
		return t.Elapsed
	}
	return time.Duration(float32(t.Elapsed) * t.Scale)
}
