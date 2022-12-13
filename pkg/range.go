package axe

import "math/rand"

type Range[A Attr[A]] struct {
	Min A
	Max A
}

func (r Range[A]) At(delta float32) A {
	var at A
	at.Interpolate(r.Min, r.Max, delta, &at)
	return at
}

func (r Range[A]) Random(rnd rand.Rand) A {
	return r.At(rnd.Float32())
}

type NumericRange[D Numeric] struct {
	Min D
	Max D
}

func (r NumericRange[D]) At(delta float32) D {
	return D(float32(r.Max-r.Min) * delta)
}

func (r NumericRange[D]) Random(rnd rand.Rand) D {
	return r.At(rnd.Float32())
}
