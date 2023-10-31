package ui

type Anchor struct {
	Base  float32
	Delta float32
}

func (a Anchor) IsZero() bool {
	return a.Base == 0 && a.Delta == 0
}

func (a *Anchor) Set(base float32, delta float32) {
	a.Base = base
	a.Delta = delta
}

func (a *Anchor) SetRelative(value float32, relative bool) {
	if relative {
		a.Base = 0
		a.Delta = value
	} else {
		a.Base = value
		a.Delta = 0
	}
}

func (a Anchor) Get(total float32) float32 {
	return a.Base + total*a.Delta
}

func (a Anchor) Is(base float32, delta float32) bool {
	return a.Base == base && a.Delta == delta
}
