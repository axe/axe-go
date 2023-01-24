package core

import "math"

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
	Interpolate(start V, end V, delta float32, out *V)
}

func Lengthen[A Attr[A]](value A, length float32) A {
	var out A
	len := value.LengthSq()
	if len != 0 && len != length*length {
		value.Scale(length/Sqrt(len), &out)
	}
	return out
}

func ClampLength[A Attr[A]](value A, min float32, max float32) A {
	var out A
	lenSq := value.LengthSq()
	if lenSq != 0 {
		if lenSq < min*min {
			len := Sqrt(lenSq)
			value.Scale(min/len, &out)
		} else if lenSq > max*max {
			len := Sqrt(lenSq)
			value.Scale(max/len, &out)
		}
	}
	return out
}

func Slerp[A Attr[A]](start A, end A, t float32) A {
	slength := start.Length()
	elength := end.Length()
	angle := float32(math.Acos(float64(start.Dot(end)) / float64(slength*elength)))

	return SlerpAngle(start, end, angle, t)
}

func SlerpNormal[A Attr[A]](start A, end A, t float32) A {
	angle := float32(math.Acos(float64(start.Dot(end))))

	return SlerpAngle(start, end, angle, t)
}

func SlerpAngle[A Attr[A]](start A, end A, angle float32, t float32) A {
	denom := Div(1, Sin(angle))
	d0 := Sin((1-t)*angle) * denom
	d1 := Sin(t*angle) * denom

	var out A
	end.Scale(d1, &out)
	out.AddScaled(start, d0, &out)
	return out
}

func Delta[A Attr[A]](start A, end A, point A) float32 {
	var p0, p1 A
	end.Sub(start, &p0)
	point.Sub(start, &p1)

	delta := Div(p0.Dot(p1), p0.LengthSq())
	return delta
}

func Closest[A Attr[A]](start A, end A, point A, line bool) A {
	delta := Delta(start, end, point)
	if !line {
		delta = Clamp(delta, 0, 1)
	}
	var out A
	out.Interpolate(start, end, delta, &out)
	return out
}

func Normalize[A Attr[A]](value A, normal *A) float32 {
	d := value.LengthSq()
	if d != 0 && d != 1 {
		d = Sqrt(d)
		value.Scale(1/d, normal)
	}
	return d
}

func IsNormal[A Attr[A]](value A) bool {
	return Abs(value.LengthSq()-1) < EPSILON
}

func DistanceFrom[A Attr[A]](start A, end A, point A, line bool) float32 {
	closest := Closest(start, end, point, line)
	return point.Distance(closest)
}

func GetTriangleHeight(base float32, side1 float32, side2 float32) float32 {
	p := (base + side1 + side2) * 0.5
	area := Sqrt(p * (p - base) * (p - side1) * (p - side2))
	height := area * 2.0 / base

	return height
}

func IsPointInView[A Attr[A]](origin A, direction A, fovCos float32, point A) bool {
	var temp A
	point.Sub(origin, &temp)
	return temp.Dot(direction) > fovCos
}

func IsCircleInView[A Attr[A]](viewOrigin A, viewDirection A, fovTan float32, fovCos float32, circle A, circleRadius float32, entirely bool) bool {
	// http://www.cbloom.com/3d/techdocs/culling.txt
	var circleToOrigin A
	circle.Sub(viewOrigin, &circleToOrigin)
	distanceAlongDirection := circleToOrigin.Dot(viewDirection)
	coneRadius := distanceAlongDirection * fovTan
	distanceFromAxis := Sqrt(circleToOrigin.LengthSq() - distanceAlongDirection*distanceAlongDirection)
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

func IsCircleInViewType[A Attr[A]](viewOrigin A, viewDirection A, fovTan float32, fovCos float32, circle A, circleRadius float32, fovType FieldOfView) bool {
	if fovType == FieldOfViewIgnore {
		return true
	}

	if fovType == FieldOfViewHalf {
		circleRadius = 0
	}

	return IsCircleInView(viewOrigin, viewDirection, fovTan, fovCos, circle, circleRadius, fovType == FieldOfViewFull)
}

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

func ParametricCubicCurve[A Attr[A]](delta float32, points []A, matrix [4][4]float32, weight float32, inverse bool, loop bool) A {
	n := len(points) - 1
	a := delta * float32(n)
	i := Clamp(Floor(a), 0, float32(n-1))
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

func InterceptTime[A Attr[A]](shooter A, shooterSpeed float32, targetPosition A, targetVelocity A) float32 {
	var tvec A
	targetPosition.Sub(shooter, &tvec)

	a := targetVelocity.LengthSq() - (shooterSpeed * shooterSpeed)
	b := 2 * targetVelocity.Dot(tvec)
	c := tvec.LengthSq()

	return QuadraticFormula(a, b, c, -1)
}

func Reflect[A Attr[A]](dir A, normal A) A {
	scale := 2 * dir.Dot(normal)
	dir.AddScaled(normal, -scale, &dir)
	return dir
}

func Refract[A Attr[A]](dir A, normal A) A {
	scale := 2 * dir.Dot(normal)
	normal.Scale(scale, &normal)
	normal.Sub(dir, &dir)
	return dir
}
