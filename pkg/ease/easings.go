package ease

import (
	"math"

	"github.com/axe/axe-go/pkg/core"
	"github.com/axe/axe-go/pkg/id"
)

func init() {
	AddModifier(In)
	AddModifier(Out)
	AddModifier(InOut)
	AddModifier(YoYo)
	AddModifier(Mirror)
	AddModifier(Reverse)
	AddModifier(Flip)

	Add(Linear)
	Add(Quad)
	Add(Ease)
	Add(Cubic)
	Add(Quartic)
	Add(Quintic)
	Add(Back)
	Add(Sine)
	Add(Cos)
	Add(Overshot)
	Add(Elastic)
	Add(Revisit)
	Add(Lasso)
	Add(SlowBounce)
	Add(Bounce)
	Add(SmallBounce)
	Add(TinyBounce)
	Add(Hesitant)
	Add(Sqrt)
	Add(Sqrtf)
	Add(Log10)
	Add(Slingshot)
	Add(Circular)
	Add(Gentle)
	Add(CssEase)
	Add(CssEaseIn)
	Add(CssEaseInOut)
	Add(CssEaseOut)
	Add(CssLinear)
}

func AddModifier(modifier Modifier) {
	Modifiers.Set(modifier.Name, modifier)
}

func Add(easing Easing) {
	Easings.Set(easing.Name(), easing)
}

// Modifiers
var (
	Modifiers = id.NewDenseKeyMap[Modifier, uint16, uint8]()

	In = NewModifier("in", func(easing Fn) Fn {
		return easing
	})
	Out = NewModifier("out", func(easing Fn) Fn {
		return func(x float32) float32 {
			return 1 - easing(1-x)
		}
	})
	InOut = NewModifier("inout", func(easing Fn) Fn {
		return func(x float32) float32 {
			if x < 0.5 {
				return easing(2.0*x) * 0.5
			} else {
				return 1.0 - (easing(2.0-2.0*x) * 0.5)
			}
		}
	})
	YoYo = NewModifier("yoyo", func(easing Fn) Fn {
		return func(x float32) float32 {
			if x < 0.5 {
				return easing(2.0 * x)
			} else {
				return easing(2.0 - 2.0*x)
			}
		}
	})
	Mirror = NewModifier("mirror", func(easing Fn) Fn {
		return func(x float32) float32 {
			if x < 0.5 {
				return easing(2.0 * x)
			} else {
				return 1.0 - easing(2.0-2.0*x)
			}
		}
	})
	Reverse = NewModifier("reverse", func(easing Fn) Fn {
		return func(x float32) float32 {
			return easing(1.0 - x)
		}
	})
	Flip = NewModifier("flip", func(easing Fn) Fn {
		return func(x float32) float32 {
			return 1.0 - easing(x)
		}
	})
)

// Easings
var (
	Easings = id.NewDenseKeyMap[Easing, uint16, uint8]()

	Linear = Name("linear", Fn(func(x float32) float32 {
		return x
	}))
	Quad = Name("quad", Fn(func(x float32) float32 {
		return x * x
	}))
	Ease = Name("ease", Fn(func(x float32) float32 {
		i := (1.0 - x)
		i2 := i * i
		x2 := x * x
		eq1 := (0.3 * i2 * x) + (3.0 * i * x2) + (x2 * x)
		eq2 := 1.0 - i2*i2

		return eq1*i + eq2*x
	}))
	Cubic = Name("cubic", Fn(func(x float32) float32 {
		return x * x * x
	}))
	Quartic = Name("quartic", Fn(func(x float32) float32 {
		x2 := x * x
		return x2 * x2
	}))
	Quintic = Name("quintic", Fn(func(x float32) float32 {
		x2 := x * x
		return x2 * x2 * x
	}))
	Back = Name("back", Fn(func(x float32) float32 {
		x2 := x * x
		x3 := x2 * x
		return x3 + x2 - x
	}))
	Sine = Name("sine", Fn(func(x float32) float32 {
		return core.Sin(x * 1.57079632679)
	}))
	Overshot = Name("overshot", Fn(func(x float32) float32 {
		return (1.0 - x*(7.0/10)) * x * (10.0 / 3.0)
	}))
	Elastic = Name("elastic", Fn(func(x float32) float32 {
		x2 := x * x
		x3 := x2 * x
		scale := x2 * ((2.0 * x3) + x2 - (4.0 * x) + 2.0)
		wave := -core.Sin(x * 10.9955742876)
		return scale * wave
	}))
	Revisit = Name("revisit", Fn(func(x float32) float32 {
		return core.Abs(x - core.Sin(x*3.14159265359))
	}))
	Lasso = Name("lasso", Fn(func(x float32) float32 {
		return (1.0 - core.Cos(x*x*x*36.0)*(1.0-x))
	}))
	SlowBounce = Name("slowbounce", Fn(func(x float32) float32 {
		x2 := x * x
		return (1.0 - core.Abs((1.0-x2)*core.Cos(x2*x*14.8044066016)))
	}))
	Bounce = Name("bounce", Fn(func(x float32) float32 {
		return (1.0 - core.Abs((1.0-x)*core.Cos(x*x*14.8044066016)))
	}))
	SmallBounce = Name("smallbounce", Fn(func(x float32) float32 {
		inv := 1.0 - x
		return (1.0 - core.Abs(inv*inv*core.Cos(x*x*14.8044066016)))
	}))
	TinyBounce = Name("tinybounce", Fn(func(x float32) float32 {
		inv := 1.0 - x
		return (1.0 - core.Abs(inv*inv*core.Cos(x*x*7.0)))
	}))
	Hesitant = Name("hesitant", Fn(func(x float32) float32 {
		return (core.Cos(x*x*12.0)*x*(1.0-x) + x)
	}))
	Sqrt = Name("sqrt", Fn(func(x float32) float32 {
		return core.Sqrt(x)
	}))
	Cos = Name("cos", Fn(func(x float32) float32 {
		return 0.5 - core.Cos(x*math.Pi)*0.5
	}))
	Sqrtf = Name("sqrtf", Fn(func(x float32) float32 {
		i := (1.0 - x)
		i2 := i * i
		return ((1.0 - i2*i2) + x) * 0.5
	}))
	Log10 = Name("log10", Fn(func(x float32) float32 {
		return (core.Log10(x+0.01) + 2.0) * 0.5 / 1.0021606868913213
	}))
	Slingshot = Name("slingshot", Fn(func(x float32) float32 {
		if x < 0.7 {
			return (x * -0.357)
		} else {
			d := x - 0.7
			return ((d*d*27.5 - 0.5) * 0.5)
		}
	}))
	Circular = Name("circular", Fn(func(x float32) float32 {
		return 1.0 - core.Sqrt(1-x*x)
	}))
	Gentle = Name("gentle", Fn(func(x float32) float32 {
		return (3.0 * (1.0 - x) * x * x) + (x * x * x)
	}))

	// CSS easing functions
	CssEase      = Name("cssease", Ease)
	CssEaseIn    = Name("csseasein", Quad)
	CssEaseOut   = Name("csseaseout", Out.Modify(Quad))
	CssEaseInOut = Name("csseaseinout", InOut.Modify(Quad))
	CssLinear    = Name("csslinear", Linear)
)
