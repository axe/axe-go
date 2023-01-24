package anim

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/axe/axe-go/pkg/core"
)

type Easing func(x float32) float32

type EasingType func(easing Easing) Easing

var EasingTypes = map[string]EasingType{
	"linear": func(easing Easing) Easing {
		return easing
	},
	"out": func(easing Easing) Easing {
		return func(x float32) float32 {
			return 1.0 - easing(1.0-x)
		}
	},
	"inout": func(easing Easing) Easing {
		return func(x float32) float32 {
			if x < 0.5 {
				return easing(2.0*x) * 0.5
			} else {
				return 1.0 - (easing(2.0-2.0*x) * 0.5)
			}
		}
	},
	"yoyo": func(easing Easing) Easing {
		return func(x float32) float32 {
			if x < 0.5 {
				return easing(2.0 * x)
			} else {
				return easing(2.0 - 2.0*x)
			}
		}
	},
	"mirror": func(easing Easing) Easing {
		return func(x float32) float32 {
			if x < 0.5 {
				return easing(2.0 * x)
			} else {
				return 1.0 - easing(2.0-2.0*x)
			}
		}
	},
	"reverse": func(easing Easing) Easing {
		return func(x float32) float32 {
			return easing(1.0 - x)
		}
	},
	"flip": func(easing Easing) Easing {
		return func(x float32) float32 {
			return 1.0 - easing(x)
		}
	},
}

var Easings = map[string]Easing{
	"": func(x float32) float32 {
		return x
	},
	"linear": func(x float32) float32 {
		return x
	},
	"quad": func(x float32) float32 {
		return x * x
	},
	"ease": func(x float32) float32 {
		i := (1.0 - x)
		i2 := i * i
		x2 := x * x
		eq1 := (0.3 * i2 * x) + (3.0 * i * x2) + (x2 * x)
		eq2 := 1.0 - i2*i2

		return eq1*i + eq2*x
	},
	"cubic": func(x float32) float32 {
		return x * x * x
	},
	"quartic": func(x float32) float32 {
		x2 := x * x
		return x2 * x2
	},
	"quintic": func(x float32) float32 {
		x2 := x * x
		return x2 * x2 * x
	},
	"back": func(x float32) float32 {
		x2 := x * x
		x3 := x2 * x
		return x3 + x2 - x
	},
	"sine": func(x float32) float32 {
		return core.Sin(x * 1.57079632679)
	},
	"overshot": func(x float32) float32 {
		return (1.0 - x*(7.0/10)) * x * (10.0 / 3.0)
	},
	"elastic": func(x float32) float32 {
		x2 := x * x
		x3 := x2 * x
		scale := x2 * ((2.0 * x3) + x2 - (4.0 * x) + 2.0)
		wave := -core.Sin(x * 10.9955742876)
		return scale * wave
	},
	"revisit": func(x float32) float32 {
		return core.Abs(x - core.Sin(x*3.14159265359))
	},
	"lasso": func(x float32) float32 {
		return (1.0 - core.Cos(x*x*x*36.0)*(1.0-x))
	},
	"slowbounce": func(x float32) float32 {
		x2 := x * x
		return (1.0 - core.Abs((1.0-x2)*core.Cos(x2*x*14.8044066016)))
	},
	"bounce": func(x float32) float32 {
		return (1.0 - core.Abs((1.0-x)*core.Cos(x*x*14.8044066016)))
	},
	"smallbounce": func(x float32) float32 {
		inv := 1.0 - x
		return (1.0 - core.Abs(inv*inv*core.Cos(x*x*14.8044066016)))
	},
	"tinybounce": func(x float32) float32 {
		inv := 1.0 - x
		return (1.0 - core.Abs(inv*inv*core.Cos(x*x*7.0)))
	},
	"hesitant": func(x float32) float32 {
		return (core.Cos(x*x*12.0)*x*(1.0-x) + x)
	},
	"sqrt": func(x float32) float32 {
		return core.Sqrt(x)
	},
	"sqrtf": func(x float32) float32 {
		i := (1.0 - x)
		i2 := i * i
		return ((1.0 - i2*i2) + x) * 0.5
	},
	"log10": func(x float32) float32 {
		return (core.Log10(x+0.01) + 2.0) * 0.5 / 1.0021606868913213
	},
	"slingshot": func(x float32) float32 {
		if x < 0.7 {
			return (x * -0.357)
		} else {
			d := x - 0.7
			return ((d*d*27.5 - 0.5) * 0.5)
		}
	},
	"circular": func(x float32) float32 {
		return 1.0 - core.Sqrt(1-x*x)
	},
	"gentle": func(x float32) float32 {
		return (3.0 * (1.0 - x) * x * x) + (x * x * x)
	},
}

func DefineEasingType(name string, easing EasingType) EasingType {
	EasingTypes[strings.ToLower(name)] = easing
	return easing
}

func DefineEasing(name string, easing Easing) Easing {
	Easings[strings.ToLower(name)] = easing
	return easing
}

type EasingScaled struct {
	Scale  float32
	Easing Easing
}

func (easing EasingScaled) Ease(x float32) float32 {
	i := easing.Easing(x)
	return easing.Scale*i + (1-easing.Scale)*x
}

func (easing EasingScaled) GetEasing() Easing {
	return easing.Ease
}

type EasingBezier struct {
	MX1, MY1, MX2, MY2 float32
}

func (easing EasingBezier) A(aA1, aA2 float32) float32 { return 1.0 - 3.0*aA2 + 3.0*aA1 }
func (easing EasingBezier) B(aA1, aA2 float32) float32 { return 3.0*aA2 - 6.0*aA1 }
func (easing EasingBezier) C(aA1 float32) float32      { return 3.0 * aA1 }
func (easing EasingBezier) CalcBezier(aT, aA1, aA2 float32) float32 {
	return ((easing.A(aA1, aA2)*aT+easing.B(aA1, aA2))*aT + easing.C(aA1)) * aT
}
func (easing EasingBezier) GetSlope(aT, aA1, aA2 float32) float32 {
	return 3.0*easing.A(aA1, aA2)*aT*aT + 2.0*easing.B(aA1, aA2)*aT + easing.C(aA1)
}
func (easing EasingBezier) GetTForX(aX float32) float32 {
	var aGuessT = aX
	for i := 0; i < 4; i++ {
		var currentSlope = easing.GetSlope(aGuessT, easing.MX1, easing.MX2)
		if currentSlope == 0.0 {
			return aGuessT
		}
		var currentX = easing.CalcBezier(aGuessT, easing.MX1, easing.MX2) - aX
		aGuessT -= currentX / currentSlope
	}
	return aGuessT
}

func (easing EasingBezier) Ease(x float32) float32 {
	return easing.CalcBezier(easing.GetTForX(x), easing.MY1, easing.MY2)
}

func (easing EasingBezier) GetEasing() Easing {
	return easing.Ease
}

// Formats:
// easing
// easing-easingtype
// easing*scale
// easing-easingtype*scale
// bezier(mx1,my1,mx2,my2)
// bezier(mx1,my1,mx2,my2)*scale
var EasingRegex = regexp.MustCompile(`^(?i)(bezier\(\s*(-?\d*\.?\d*)\s*,\s*(-?\d*\.?\d*)\s*,\s*(-?\d*\.?\d*)\s*,\s*(-?\d*\.?\d*)\s*\)|([^-]*))(|\s*-\s*([^*]+))(|\s*\*\s*(-?\d*\.?\d*))`)

func ParseEasing(s string) Easing {
	matches := EasingRegex.FindStringSubmatch(s)
	if matches == nil {
		panic(fmt.Sprintf("invalid easing %s", s))
	}

	var easing Easing
	if matches[1] != "" {
		mx1, _ := strconv.ParseFloat(matches[2], 32)
		my1, _ := strconv.ParseFloat(matches[3], 32)
		mx2, _ := strconv.ParseFloat(matches[4], 32)
		my2, _ := strconv.ParseFloat(matches[5], 32)

		easing = EasingBezier{
			MX1: float32(mx1),
			MY1: float32(my1),
			MX2: float32(mx2),
			MY2: float32(my2),
		}.GetEasing()
	} else {
		easing = Easings[strings.ToLower(matches[6])]
		if easing == nil {
			panic(fmt.Sprintf("easing not found %s", matches[6]))
		}
	}

	if matches[7] != "" {
		easingType := EasingTypes[strings.ToLower(matches[7])]
		if easingType == nil {
			panic(fmt.Sprintf("easing type not found %s", matches[7]))
		}

		easing = easingType(easing)
	}

	if matches[9] != "" {
		scale, _ := strconv.ParseFloat(matches[9], 32)

		easing = EasingScaled{
			Scale:  float32(scale),
			Easing: easing,
		}.GetEasing()
	}

	return easing
}
