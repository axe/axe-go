package fx

import "github.com/axe/axe-go/pkg/util"

func Life(particle []float32, format *Format) float32 {
	return util.Div(format.Get(Age, particle)[0], format.Get(Lifespan, particle)[0])
}

func Lerp(start, end []float32, delta float32, out []float32) {
	switch len(out) {
	case 1:
		out[0] = util.Lerp(start[0], end[0], delta)
	case 2:
		out[0] = util.Lerp(start[0], end[0], delta)
		out[1] = util.Lerp(start[1], end[1], delta)
	case 3:
		out[0] = util.Lerp(start[0], end[0], delta)
		out[1] = util.Lerp(start[1], end[1], delta)
		out[2] = util.Lerp(start[2], end[2], delta)
	case 4:
		out[0] = util.Lerp(start[0], end[0], delta)
		out[1] = util.Lerp(start[1], end[1], delta)
		out[2] = util.Lerp(start[2], end[2], delta)
		out[3] = util.Lerp(start[3], end[3], delta)
	}
}

func Add(out, add []float32) {
	switch len(out) {
	case 1:
		out[0] += add[0]
	case 2:
		out[0] += add[0]
		out[1] += add[1]
	case 3:
		out[0] += add[0]
		out[1] += add[1]
		out[2] += add[2]
	case 4:
		out[0] += add[0]
		out[1] += add[1]
		out[2] += add[2]
		out[3] += add[3]
	}
}

func MultiplyScalar(out []float32, scalar float32) {
	switch len(out) {
	case 1:
		out[0] *= scalar
	case 2:
		out[0] *= scalar
		out[1] *= scalar
	case 3:
		out[0] *= scalar
		out[1] *= scalar
		out[2] *= scalar
	case 4:
		out[0] *= scalar
		out[1] *= scalar
		out[2] *= scalar
		out[3] *= scalar
	}
}

func Length(out []float32) float32 {
	switch len(out) {
	case 1:
		return out[0]
	default:
		return util.Sqrt(LengthSq(out))
	}
}

func LengthSq(out []float32) float32 {
	switch len(out) {
	case 1:
		return out[0] * out[0]
	case 2:
		return out[0]*out[0] + out[1]*out[1]
	case 3:
		return out[0]*out[0] + out[1]*out[1] + out[2]*out[2]
	case 4:
		return out[0]*out[0] + out[1]*out[1] + out[2]*out[2] + out[3]*out[3]
	}
	return 0
}
