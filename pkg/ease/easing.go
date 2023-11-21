package ease

import (
	"fmt"

	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/util"
)

type Easing interface {
	Ease(x float32) float32
	Name() id.Identifier
	String() string
}

func Get(x float32, easing Easing) float32 {
	if easing != nil {
		return easing.Ease(x)
	}
	return x
}

func NameOf(easing Easing) id.Identifier {
	if easing != nil {
		return easing.Name()
	}
	return id.None
}

func StringOf(easing Easing) string {
	if easing != nil {
		return easing.String()
	}
	return ""
}

// Fn

type Fn func(x float32) float32

var _ Easing = Fn(nil)

func (fn Fn) Ease(x float32) float32 {
	return fn(x)
}
func (fn Fn) Name() id.Identifier {
	return id.None
}
func (fn Fn) String() string {
	return ""
}

// Named

type Named struct {
	easing Easing
	name   id.Identifier
}

var _ Easing = Named{}

func (en Named) Ease(x float32) float32 {
	return en.easing.Ease(x)
}
func (en Named) Name() id.Identifier {
	return en.name
}
func (en Named) String() string {
	return en.name.String()
}

func Name(name string, easing Easing) Easing {
	return Named{easing: easing, name: id.Get(name)}
}

// Modified

type ModifierFn func(easing Fn) Fn

type Modifier struct {
	Fn   ModifierFn
	Name id.Identifier
}

func NewModifier(name string, fn ModifierFn) Modifier {
	return Modifier{Name: id.Get(name), Fn: fn}
}

func (m Modifier) Modify(easing Easing) Modified {
	return Modified{Modifier: m, Easing: easing, Fn: m.Fn(easing.Ease)}
}

type Modified struct {
	Modifier Modifier
	Easing   Easing
	Fn       Fn
}

var _ Easing = Modified{}

func (m Modified) Ease(x float32) float32 {
	return m.Fn(x)
}
func (m Modified) Name() id.Identifier {
	return id.None
}
func (m Modified) String() string {
	return m.Easing.String() + "-" + m.Modifier.Name.String()
}

// Scaled

type Scaled struct {
	Scale  float32
	Easing Easing
}

var _ Easing = Scaled{}

func (s Scaled) Ease(x float32) float32 {
	i := s.Easing.Ease(x)
	return s.Scale*i + (1-s.Scale)*x
}
func (s Scaled) Name() id.Identifier {
	return id.None
}
func (s Scaled) String() string {
	return fmt.Sprintf("%s*%f", s.Easing.String(), s.Scale)
}

// Bezier

type Bezier struct {
	MX1, MY1, MX2, MY2 float32
}

var _ Easing = Bezier{}

func NewBezier(MX1, MY1, MX2, MY2 float32) Bezier {
	return Bezier{MX1: MX1, MY1: MY1, MX2: MX2, MY2: MY2}
}

func (b Bezier) A(aA1, aA2 float32) float32 { return 1.0 - 3.0*aA2 + 3.0*aA1 }
func (b Bezier) B(aA1, aA2 float32) float32 { return 3.0*aA2 - 6.0*aA1 }
func (b Bezier) C(aA1 float32) float32      { return 3.0 * aA1 }
func (b Bezier) CalcBezier(aT, aA1, aA2 float32) float32 {
	return ((b.A(aA1, aA2)*aT+b.B(aA1, aA2))*aT + b.C(aA1)) * aT
}
func (b Bezier) GetSlope(aT, aA1, aA2 float32) float32 {
	return 3.0*b.A(aA1, aA2)*aT*aT + 2.0*b.B(aA1, aA2)*aT + b.C(aA1)
}
func (b Bezier) GetTForX(aX float32) float32 {
	var aGuessT = aX
	for i := 0; i < 4; i++ {
		var currentSlope = b.GetSlope(aGuessT, b.MX1, b.MX2)
		if currentSlope == 0.0 {
			return aGuessT
		}
		var currentX = b.CalcBezier(aGuessT, b.MX1, b.MX2) - aX
		aGuessT -= currentX / currentSlope
	}
	return aGuessT
}
func (b Bezier) Ease(x float32) float32 {
	return b.CalcBezier(b.GetTForX(x), b.MY1, b.MY2)
}
func (b Bezier) Name() id.Identifier {
	return id.None
}
func (b Bezier) String() string {
	return fmt.Sprintf("bezier(%f,%f,%f,%f)", b.MX1, b.MY1, b.MX2, b.MY2)
}

// Subset

type Subset struct {
	Easing     Easing
	Start, End float32

	startValue, endValue float32
}

var _ Easing = Subset{}

func NewSubset(easing Easing, start, end float32) Subset {
	if easing == nil {
		return Subset{}
	}

	return Subset{
		Easing: easing,
		Start:  start,
		End:    end,

		startValue: easing.Ease(start),
		endValue:   easing.Ease(end),
	}
}
func (s Subset) Ease(x float32) float32 {
	easingDelta := util.Lerp(s.Start, s.End, x)
	rangeDelta := s.Easing.Ease(easingDelta)
	easedDelta := util.Delta(s.startValue, s.endValue, rangeDelta)

	return easedDelta
}
func (s Subset) Name() id.Identifier {
	return id.None
}
func (s Subset) String() string {
	return fmt.Sprintf("%s{%f,%f}", s.Easing.String(), s.Start, s.End)
}
