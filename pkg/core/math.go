package core

import "math"

func Div[D Numeric](a D, b D) D {
	if b == 0 {
		return 0
	}
	return a / b
}

func Clamp[D Numeric](v D, min D, max D) D {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func Min[D Numeric](a D, b D) D {
	if a < b {
		return a
	}
	return b
}

func Max[D Numeric](a D, b D) D {
	if a > b {
		return a
	}
	return b
}

func Abs[D Numeric](v D) D {
	return D(math.Abs(float64(v)))
}

func Floor[D Numeric](v D) D {
	return D(math.Floor(float64(v)))
}

func Sin[D Numeric](rad D) D {
	return D(math.Sin(float64(rad)))
}

func Cos[D Numeric](rad D) D {
	return D(math.Cos(float64(rad)))
}

func Sqrt[D Numeric](v D) D {
	return D(math.Sqrt(float64(v)))
}

func Log10[D Numeric](v D) D {
	return D(math.Log10(float64(v)))
}

func Sign[D Numeric](v D) int {
	if v == 0 {
		return 0
	}
	if v < 0 {
		return -1
	}
	return 1
}

func Gcd[D Integer](a D, b D) D {
	shift := 0

	if a == 0 || b == 0 {
		return (a | b)
	}

	for shift := 0; ((a | b) & 1) == 0; shift++ {
		a >>= 1
		b >>= 1
	}

	for (a & 1) == 0 {
		a >>= 1
	}

	for {
		for (b & 1) == 0 {
			b >>= 1
		}
		if a < b {
			b -= a
		} else {
			d := a - b
			a = b
			b = d
		}
		b >>= 1

		if b == 0 {
			break
		}
	}

	return (a << shift)
}

func Factorial[D Numeric](x D) D {
	n := x
	x--
	for x >= 1 {
		n *= x
		x--
	}
	return n
}

func Choose[D Integer](n D, m D) D {
	num := D(1)
	den := D(1)

	if m > (n >> 1) {
		m = n - m
	}

	for m >= 1 {
		num *= n
		n--
		den *= m
		m--
		gcd := Gcd(num, den)
		num /= gcd
		den /= gcd
	}

	return num
}

var EPSILON = float32(0.00001)

func QuadraticFormula(a, b, c, none float32) float32 {
	t0 := float32(math.SmallestNonzeroFloat32)
	t1 := float32(math.SmallestNonzeroFloat32)

	if Abs(a) < EPSILON {
		if Abs(b) < EPSILON {
			if Abs(c) < EPSILON {
				t0 = 0.0
				t1 = 0.0
			}
		} else {
			t0 = -c / b
			t1 = -c / b
		}
	} else {
		disc := b*b - 4*a*c

		if disc >= 0 {
			disc = Sqrt(disc)
			a = 2 * a
			t0 = (-b - disc) / a
			t1 = (-b + disc) / a
		}
	}

	if t0 != math.SmallestNonzeroFloat32 {
		t := Min(t0, t1)

		if t < 0 {
			t = Max(t0, t1)
		}

		if t > 0 {
			return t
		}
	}

	return -1
}
