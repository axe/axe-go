package axe

import "math"

type Scalar[D Numeric] struct {
	Value D
}

type Scalarf = Scalar[float32]
type Scalari = Scalar[int]

var _ Attr[Scalarf] = Scalarf{}
var _ Attr[Scalari] = Scalari{}

func (s Scalar[D]) Distance(value Scalar[D]) float32 {
	return float32(math.Abs(float64(s.Value - value.Value)))
}
func (s Scalar[D]) DistanceSq(value Scalar[D]) float32 {
	d := s.Distance(value)
	return d * d
}
func (s Scalar[D]) Length() float32 {
	return float32(math.Abs(float64(s.Value)))
}
func (s Scalar[D]) LengthSq() float32 {
	return float32(s.Value * s.Value)
}
func (s Scalar[D]) Dot(value Scalar[D]) float32 {
	return float32(s.Value * value.Value)
}
func (s Scalar[D]) Components() int {
	return 1
}
func (s Scalar[D]) GetComponent(index int) float32 {
	return float32(s.Value)
}
func (s Scalar[D]) SetComponent(index int, value float32, out *Scalar[D]) {
	out.Value = D(value)
}
func (s Scalar[D]) SetComponents(value float32, out *Scalar[D]) {
	out.Value = D(value)
}
func (s Scalar[D]) Set(out *Scalar[D]) {
	out.Value = s.Value
}
func (s Scalar[D]) Add(addend Scalar[D], out *Scalar[D]) {
	out.Value = s.Value + addend.Value
}
func (s Scalar[D]) AddScaled(value Scalar[D], scale float32, out *Scalar[D]) {
	out.Value = D(float32(s.Value) + float32(value.Value)*scale)
}
func (s Scalar[D]) Sub(subtrahend Scalar[D], out *Scalar[D]) {
	out.Value = s.Value - subtrahend.Value
}
func (s Scalar[D]) Mul(factor Scalar[D], out *Scalar[D]) {
	out.Value = s.Value * factor.Value
}
func (s Scalar[D]) Div(factor Scalar[D], out *Scalar[D]) {
	if factor.Value == 0 {
		out.Value = 0
	} else {
		out.Value = s.Value / factor.Value
	}
}
func (s Scalar[D]) Scale(scale float32, out *Scalar[D]) {
	out.Value = D(float32(s.Value) * scale)
}
func (start Scalar[D]) Lerp(end Scalar[D], delta float32, out *Scalar[D]) {
	out.Value = D(float32(end.Value-start.Value)*delta + float32(start.Value))
}
