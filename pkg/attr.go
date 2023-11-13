package axe

import (
	"math"

	"github.com/axe/axe-go/pkg/util"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// An attribute in the game that has basic math operations.
type Attr[V any] interface {
	// distance between this and value
	Distance(value V) float32
	// squared distance between this and value
	DistanceSq(value V) float32
	// distance between this and zero
	Length() float32
	// squared distance between this and zero
	LengthSq() float32
	// dot product between this and value
	Dot(value V) float32
	// the number of float components that make up this attribute
	Components() int
	// gets the float component at the given index
	GetComponent(index int) float32
	// out[index] = value
	SetComponent(index int, value float32, out *V)
	// out[all] = value
	SetComponents(value float32, out *V)
	// out = this
	Set(out *V)
	// out = this + value
	Add(addend V, out *V)
	// out = this + value * scale
	AddScaled(value V, scale float32, out *V)
	// out = this - value
	Sub(subtrahend V, out *V)
	// out = this * value
	Mul(factor V, out *V)
	// out = this / value
	Div(factor V, out *V)
	// out = this * scale
	Scale(scale float32, out *V)
	// out = (end - start) * delta + start
	Lerp(end V, delta float32, out *V)
}

// Computes a value with the given length.
func Lengthen[A Attr[A]](value A, length float32) A {
	var out A
	len := value.LengthSq()
	if len != 0 && len != length*length {
		value.Scale(length/util.Sqrt(len), &out)
	} else {
		value.Set(&out)
	}
	return out
}

// Computes a value whose length is clamped between the given min and max.
func ClampLength[A Attr[A]](value A, min float32, max float32) A {
	var out A
	lenSq := value.LengthSq()
	if lenSq != 0 {
		if lenSq < min*min {
			len := util.Sqrt(lenSq)
			value.Scale(min/len, &out)
		} else if lenSq > max*max {
			len := util.Sqrt(lenSq)
			value.Scale(max/len, &out)
		} else {
			value.Set(&out)
		}
	}
	return out
}

// Performs a slerp between two values. Imagine the two values define a curve that lies
// on an ellipse where the zero value is the center. The returned value will exist on that
// curve where 0 is at start and 1 is at end and 0.5 is halfway between.
func Slerp[A Attr[A]](start A, end A, t float32) A {
	slength := start.Length()
	elength := end.Length()
	lengthSq := float64(slength * elength)
	if lengthSq == 0 {
		return start
	}
	angle := float32(math.Acos(float64(start.Dot(end)) / lengthSq))

	return SlerpAngle(start, end, angle, t)
}

// Performs a slerp between two normals given a t. The returned value will exist on the shortest
// curve from the end of start and end and t between those two values (t is 0 to 1).
func SlerpNormal[A Attr[A]](start A, end A, t float32) A {
	angle := float32(math.Acos(float64(start.Dot(end))))

	return SlerpAngle(start, end, angle, t)
}

// Performs a slerp between the start and end value given the curves "angle" between the two values and a t value.
// When t is 0 the returned value will equal start and when 1 will match end. Angle is in radians.
func SlerpAngle[A Attr[A]](start A, end A, angle float32, t float32) A {
	if angle == 0 {
		return start
	}
	denom := util.Div(1, util.Sin(angle))
	d0 := util.Sin((1-t)*angle) * denom
	d1 := util.Sin(t*angle) * denom

	var out A
	end.Scale(d1, &out)
	out.AddScaled(start, d0, &out)
	return out
}

// A line is defined by start and end and there is a point on it that's closest to the given point.
// The delta value defines how close that casted point is between start and end where 0 is on start,
// 1 is on end, 0.5 is halfway between, and -1 is before start on the line the same distance
// between start and end.
func Delta[A Attr[A]](start A, end A, point A) float32 {
	var p0, p1 A
	end.Sub(start, &p0)
	point.Sub(start, &p1)

	delta := util.Div(p0.Dot(p1), p0.LengthSq())
	return delta
}

// Returns the closest value to the given point on the line or segment defined by start and end.
func Closest[A Attr[A]](start A, end A, point A, line bool) A {
	delta := Delta(start, end, point)
	if !line {
		delta = util.Clamp(delta, 0, 1)
	}
	var out A
	start.Lerp(end, delta, &out)
	return out
}

// Sets normal to the given value but with a length of 1 or 0.
func Normalize[A Attr[A]](value A, normal *A) float32 {
	d := value.LengthSq()
	if d != 0 && d != 1 {
		d = util.Sqrt(d)
		value.Scale(1/d, normal)
	}
	return d
}

// Returns if the given value is normalized (has a length of 1).
func IsNormal[A Attr[A]](value A) bool {
	return util.Abs(value.LengthSq()-1) < util.EPSILON
}

// Returns the shortest distance from the point to the line or segment defined by start and end.
func DistanceFrom[A Attr[A]](start A, end A, point A, line bool) float32 {
	closest := Closest(start, end, point, line)
	return point.Distance(closest)
}

// Calculates the height triangle given it's base length and two sides.
func GetTriangleHeight(base float32, side1 float32, side2 float32) float32 {
	p := (base + side1 + side2) * 0.5
	area := util.Sqrt(p * (p - base) * (p - side1) * (p - side2))
	height := area * 2.0 / base

	return height
}

// Returns true if the point is in the defined conical view.
func IsPointInView[A Attr[A]](viewOrigin A, viewDirection A, fovCos float32, point A) bool {
	var temp A
	point.Sub(viewOrigin, &temp)
	return temp.Dot(viewDirection) > fovCos
}

// Returns true if the circle is fully or partially in the defined conical view.
func IsCircleInView[A Attr[A]](viewOrigin A, viewDirection A, fovTan float32, fovCos float32, circle A, circleRadius float32, entirely bool) bool {
	// http://www.cbloom.com/3d/techdocs/culling.txt
	var circleToOrigin A
	circle.Sub(viewOrigin, &circleToOrigin)
	distanceAlongDirection := circleToOrigin.Dot(viewDirection)
	coneRadius := distanceAlongDirection * fovTan
	distanceFromAxis := util.Sqrt(circleToOrigin.LengthSq() - distanceAlongDirection*distanceAlongDirection)
	distanceFromCenterToCone := distanceFromAxis - coneRadius
	shortestDistance := distanceFromCenterToCone * fovCos

	if entirely {
		shortestDistance += circleRadius
	} else {
		shortestDistance -= circleRadius
	}

	return shortestDistance <= 0
}

type FieldOfView string

const (
	FieldOfViewIgnore FieldOfView = "ignore"
	FieldOfViewHalf   FieldOfView = "half"
	FieldOfViewFull   FieldOfView = "full"
)

// Returns true if the circle is in the defined view based on FOV option.
func IsCircleInViewType[A Attr[A]](viewOrigin A, viewDirection A, fovTan float32, fovCos float32, circle A, circleRadius float32, fovType FieldOfView) bool {
	if fovType == FieldOfViewIgnore {
		return true
	}

	if fovType == FieldOfViewHalf {
		circleRadius = 0
	}

	return IsCircleInView(viewOrigin, viewDirection, fovTan, fovCos, circle, circleRadius, fovType == FieldOfViewFull)
}

// Calculates the value on the cubic curve given a delta between 0 and 1, the 4 control points, the matrix weights, and if its an inverse.
func CubicCurve[A Attr[A]](delta float32, p0 A, p1 A, p2 A, p3 A, matrix [4][4]float32, inverse bool) A {
	d0 := float32(1.0)
	d1 := delta
	d2 := d1 * d1
	d3 := d2 * d1

	ts := [4]float32{d0, d1, d2, d3}
	if inverse {
		ts[0] = d3
		ts[1] = d2
		ts[2] = d1
		ts[3] = d0
	}

	var out A

	for i := 0; i < 4; i++ {
		var temp A
		temp.AddScaled(p0, matrix[i][0], &temp)
		temp.AddScaled(p1, matrix[i][1], &temp)
		temp.AddScaled(p2, matrix[i][2], &temp)
		temp.AddScaled(p3, matrix[i][3], &temp)
		temp.AddScaled(out, ts[i], &out)
	}
	return out
}

// Calculates a parametric cubic curve given a delta between 0 and 1, the points of the curve, the weights, if its inversed, and if the curve loops.
func ParametricCubicCurve[A Attr[A]](delta float32, points []A, matrix [4][4]float32, weight float32, inverse bool, loop bool) A {
	n := len(points) - 1
	a := delta * float32(n)
	i := util.Clamp(util.Floor(a), 0, float32(n-1))
	d := a - i
	index := int(i)

	p0 := points[0]
	if i == 0 {
		if loop {
			p0 = points[n]
		}
	} else {
		p0 = points[index-1]
	}

	p1 := points[index]
	p2 := points[index+1]

	p3 := points[0]
	if index == n-1 {
		if !loop {
			p3 = points[n]
		}
	} else {
		p3 = points[index+2]
	}

	out := CubicCurve(d, p0, p1, p2, p3, matrix, inverse)
	out.Scale(weight, &out)
	return out
}

// Calculates the time an interceptor could intercept a target given the interceptors
// position and possible speed and the targets current position and velocity. If no
// intercept exists based on the parameters then -1 is returned. Otherwise a value is
// returned that can be used to calculate the interception point (targetPosition+(targetVelocity*time)).
func InterceptTime[A Attr[A]](interceptor A, interceptorSpeed float32, targetPosition A, targetVelocity A) float32 {
	var tvec A
	targetPosition.Sub(interceptor, &tvec)

	a := targetVelocity.LengthSq() - (interceptorSpeed * interceptorSpeed)
	b := 2 * targetVelocity.Dot(tvec)
	c := tvec.LengthSq()

	return util.QuadraticFormula(a, b, c, -1)
}

// Reflects the direction across a given normal. Imagine the normal is on a plane
// pointing away from it and a reflection is a ball with the given direction bouncing off of it.
func Reflect[A Attr[A]](dir A, normal A) A {
	scale := 2 * dir.Dot(normal)
	dir.AddScaled(normal, -scale, &dir)
	return dir
}

// Refracts the direction across the given normal. Like reflect except through the
// plane defined by the normal.
func Refract[A Attr[A]](dir A, normal A) A {
	scale := 2 * dir.Dot(normal)
	normal.Scale(scale, &normal)
	normal.Sub(dir, &dir)
	return dir
}
