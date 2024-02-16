package util

import (
	"math"
)

var EPSILON = float32(0.00001)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Divides a by b, unless b is zero then zero is returned.
func Div[D Numeric](a D, b D) D {
	if b == 0 {
		return 0
	}
	return a / b
}

// Returns v but no larger than the max or smaller than the min.
func Clamp[D Numeric](v D, min D, max D) D {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Returns the smallest number between a & b.
func Min[D Numeric](a D, b D) D {
	if a < b {
		return a
	}
	return b
}

// Returns the largest number between a & b.
func Max[D Numeric](a D, b D) D {
	if a > b {
		return a
	}
	return b
}

// Returns the smallest then largest number between a & b.
func MinMax[D Numeric](a D, b D) (D, D) {
	if a < b {
		return a, b
	}
	return b, a
}

// Constrains the absolute value of `value` to be no larger than max.
func MaxMagnitude[D Numeric](value D, max D) D {
	if value < 0 && value < -max {
		value = -max
	} else if value > 0 && value > max {
		value = max
	}
	return value
}

// Constrains the absolute value of `value` to be no smaller than min (unless value is zero)
func MinMagnitude[D Numeric](value D, min D) D {
	if value < 0 && value > -min {
		value = -min
	} else if value > 0 && value < min {
		value = min
	}
	return value
}

// Computes the absolute value of the given value.
func Abs[D Numeric](v D) D {
	return D(math.Abs(float64(v)))
}

// Computes the floor of the given value.
func Floor[D Numeric](v D) D {
	return D(math.Floor(float64(v)))
}

// Computes the ceil of the given value.
func Ceil[D Numeric](v D) D {
	return D(math.Ceil(float64(v)))
}

// Rounds the value to the largest absolute value.
// Up(4.5)=5, Up(-4.5)=-5, Up(3)=3
func Up[D Numeric](v D) D {
	if v < 0 {
		return Floor(v)
	}
	return Ceil(v)
}

// Rounds the value to the smallest absolute value.
// Down(4.5)=4, Down(-4.5)=-4, Down(3)=3
func Down[D Numeric](v D) D {
	if v < 0 {
		return Ceil(v)
	}
	return Floor(v)
}

// Computes the square of the value
func Sq[D Numeric](x D) D {
	return x * x
}

// Computes the sine of the given radians.
func Sin[D Numeric](rad D) D {
	return D(math.Sin(float64(rad)))
}

// Computes the arcsine of the given radians.
func Asin[D Numeric](rad D) D {
	return D(math.Asin(float64(rad)))
}

// Computes the cosine of the given radians.
func Cos[D Numeric](rad D) D {
	return D(math.Cos(float64(rad)))
}

// Computes the arcosine of the given radians.
func Acos[D Numeric](rad D) D {
	return D(math.Acos(float64(rad)))
}

// Computes the square root of the given number.
func Sqrt[D Numeric](v D) D {
	return D(math.Sqrt(float64(v)))
}

// Computes the Atan2 of the given radians.
func Atan2[D Numeric](y D, x D) D {
	return D(math.Atan2(float64(y), float64(x)))
}

// Computes the sign of the given number where 0 has a sign of 0,
// a negative number returns -1 and a positive number returns 1.
func Sign[D Numeric](v D) int {
	if v == 0 {
		return 0
	}
	if v < 0 {
		return -1
	}
	return 1
}

// Returns "value" with the sign of the value "sign"
func CopySign[D Numeric](value D, sign D) D {
	return D(math.Copysign(float64(value), float64(sign)))
}

// Does a lineary interpolation between
func Lerp[D Numeric](start, end D, delta float32) D {
	return D(float32(end-start)*delta) + start
}

// Computes the delta of v between a start and end value where a delta of 0 is at start and a delta of 1 is at end.
func Delta[D Numeric](start, end, value D) float32 {
	return Div(float32(value-start), float32(end-start))
}

// Returns if two numbers are equal enough (based on EPSILON)
func Equal[D Numeric](a, b D) bool {
	return float32(Abs(a-b)) < EPSILON
}

// Returns the cosine & sine of the given radians.
func CosSin[D Numeric](rad D) (cos D, sin D) {
	cos = Cos[D](rad)
	sin = Sin[D](rad)
	return
}

// Returns the cosine & sine of the given radians.
func Pow[D Numeric](x, y D) D {
	return D(math.Pow(float64(x), float64(y)))
}

// Calculates the greatest common denominator between two integer numbers.
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

// Computes the factorial of the given number.
// ex: Factorial(5) = 5*4*3*2*1
func Factorial[D Numeric](x D) D {
	n := x
	x--
	for x > 1 {
		n *= x
		x--
	}
	return n
}

// Computes the number of combinations of size m that can be made with n things.
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

// Computes the quadratic formula between a, b, and c. If it can't be computed
// then none is returned.
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

	return none
}
