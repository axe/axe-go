package color

func Lighten(scale float32) Modify {
	return func(c Color) Color {
		return c.Lighten(scale)
	}
}

func Darken(scale float32) Modify {
	return func(c Color) Color {
		return c.Darken(scale)
	}
}

func Alpha(alpha float32) Modify {
	return func(c Color) Color {
		return c.AlphaScale(alpha)
	}
}

func Multiply(shade Color) Modify {
	return func(c Color) Color {
		return c.Multiply(shade)
	}
}

type Modify func(Color) Color

func (a Modify) Then(b Modify) Modify {
	return func(c Color) Color {
		return b(a(c))
	}
}

func (a Modify) Lerp(b Modify, delta float32) Modify {
	if a == nil || delta >= 1 {
		return b
	}
	if b == nil || delta <= 0 {
		return a
	}
	return func(c Color) Color {
		start := a(c)
		end := b(c)
		return start.Lerp(end, delta)
	}
}

func (a Modify) Modify(c Color) Color {
	if a != nil {
		return a(c)
	}
	return c
}

func (a Modify) GetEffective() Modify {
	if a.HasAffect() {
		return a
	}
	return nil
}

func (a Modify) HasAffect() bool {
	if a == nil {
		return false
	}

	return a(White) != White || a(Transparent) != Transparent
}

func (a Modify) Equals(b Modify) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if a == nil {
		return true
	}

	return a(White) == b(White) && a(Transparent) == b(Transparent)
}
