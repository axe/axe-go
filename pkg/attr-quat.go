package axe

import "math"

type Quat struct {
	X, Y, Z, W float32
}

var _ Attr[Quat] = Quat{}

func (q Quat) Set(out *Quat) {
	out.X = q.X
	out.Y = q.Y
	out.Z = q.Z
	out.W = q.W
}

func (q *Quat) Sets(X, Y, Z, W float32) {
	q.X = X
	q.Y = Y
	q.Z = Z
	q.W = W
}

func (q *Quat) SetAngle(dir Vec3f, angle float32) {
	half := angle * 0.5
	sin := Sin(half) / dir.Length()
	cos := Cos(half)
	q.X = dir.X * sin
	q.Y = dir.Y * sin
	q.Z = dir.Z * sin
	q.W = cos
}

func (q Quat) GetAngle() (dir Vec3f, angle float32) {
	invW := Div(1, Sqrt(1-q.W*q.W))
	dir.X = q.X * invW
	dir.Y = q.Y * invW
	dir.Z = q.Z * invW
	angle = Acos(q.W) * 2
	return
}

func (q *Quat) SetEuler(yaw float32, pitch float32, roll float32) {
	hr := roll * 0.5
	shr := Sin(hr)
	chr := Cos(hr)
	hp := pitch * 0.5
	shp := Sin(hp)
	chp := Cos(hp)
	hy := yaw * 0.5
	shy := Sin(hy)
	chy := Cos(hy)
	chy_shp := chy * shp
	shy_chp := shy * chp
	chy_chp := chy * chp
	shy_shp := shy * shp

	q.X = ((chy_shp * chr) + (shy_chp * shr))
	q.Y = ((shy_chp * chr) - (chy_shp * shr))
	q.Z = ((chy_chp * shr) - (shy_shp * chr))
	q.W = ((chy_chp * chr) + (shy_shp * shr))
}

func (q Quat) GetEuler() Vec3f {
	roll := Atan2(2.0*(q.W*q.X+q.Y*q.Z), 1.0-2.0*(q.X*q.X+q.Y*q.Y))
	sinp := 2.0 * (q.W*q.Y - q.Z*q.X)
	var pitch float32
	if Abs(sinp) >= 1.0 {
		pitch = CopySign(math.Pi/2.0, sinp)
	} else {
		pitch = Asin(pitch)
	}
	yaw := Atan2(2.0*(q.W*q.Z+q.X*q.Y), 1.0-2.0*(q.Y*q.Y+q.Z*q.Z))
	return Vec3f{roll, pitch, yaw}
}

func (q Quat) LengthSq() float32 {
	return q.Dot(q)
}

func (q Quat) Length() float32 {
	return Sqrt(q.LengthSq())
}

func (q *Quat) Normal() float32 {
	l := q.Length()
	if l != 0 {
		inv := 1.0 / l
		q.X *= inv
		q.Y *= inv
		q.Z *= inv
	}
	return l
}

func (q *Quat) Conjugate() {
	q.X = -q.X
	q.Y = -q.Y
	q.Z = -q.Z
}

func (q *Quat) Invert() {
	l := -q.LengthSq()
	if l != 0 {
		q.X /= l
		q.Y /= l
		q.Z /= l
		q.W /= -l
	}
}

func (a Quat) Mul(b Quat, out *Quat) {
	out.Multiply(a, b)
}

func (a Quat) Div(b Quat, out *Quat) {
	inv := b
	inv.Invert()
	out.Multiply(a, inv)
}

func (q *Quat) Multiply(a Quat, b Quat) {
	q.X = a.W*b.X + a.X*b.W + a.Y*b.Z - a.Z*b.Y
	q.Y = a.W*b.Y + a.Y*b.W + a.Z*b.X - a.X*b.Z
	q.Z = a.W*b.Z + a.Z*b.W + a.X*b.Y - a.Y*b.X
	q.W = a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z
}

func (q *Quat) MultiplyValues(aX, aY, aZ, aW, bX, bY, bZ, bW float32) {
	q.X = aW*bX + aX*bW + aY*bZ - aZ*bY
	q.Y = aW*bY + aY*bW + aZ*bX - aX*bZ
	q.Z = aW*bZ + aZ*bW + aX*bY - aY*bX
	q.W = aW*bW - aX*bX - aY*bY - aZ*bZ
}

func (q Quat) Rotate(v Vec3f, out *Vec3f) {
	m := v.Normal(&v)
	rot := Quat{}
	rot.MultiplyValues(v.X, v.Y, v.Z, 0, -q.X, -q.Y, -q.Z, q.W)
	rot.MultiplyValues(q.X, q.Y, q.Z, q.W, rot.X, rot.Y, rot.Z, rot.W)
	v.Scale(m, out)
}

func (q *Quat) RotateAround(point Vec3f, angle float32, axis Vec3f) {
	p := Quat{}
	p.X = point.X
	p.Y = point.Y
	p.Z = point.Z

	q.SetAngle(axis, angle)

	inv := *q
	inv.Invert()

	q.MultiplyValues(q.X, q.Y, q.Z, q.W, p.X, p.Y, p.Z, p.W)
	q.MultiplyValues(q.X, q.Y, q.Z, q.W, inv.X, inv.Y, inv.Z, inv.W)
}

func (q *Quat) ToMatrix(m Matrix4f) {
	xx := q.X * q.X
	xy := q.X * q.Y
	xz := q.X * q.Z
	yy := q.Y * q.Y
	zz := q.Z * q.Z
	yz := q.Y * q.Z
	wx := q.W * q.X
	wy := q.W * q.Y
	wz := q.W * q.Z

	m.columns[0].X = 1.0 - 2.0*(yy+zz)
	m.columns[0].Y = 2.0 * (xy - wz)
	m.columns[0].Z = 2.0 * (xz + wy)
	m.columns[0].W = 0.0

	m.columns[1].X = 2.0 * (xy + wz)
	m.columns[1].Y = 1.0 - 2.0*(xx+zz)
	m.columns[1].Z = 2.0 * (yz - wx)
	m.columns[1].W = 0.0

	m.columns[2].X = 2.0 * (xz - wy)
	m.columns[2].Y = 2.0 * (yz + wx)
	m.columns[2].Z = 1.0 - 2.0*(xx+yy)
	m.columns[2].W = 0.0

	m.columns[3].X = 0
	m.columns[3].Y = 0
	m.columns[3].Z = 0
	m.columns[3].W = 1
}

func (q1 Quat) Dot(q2 Quat) float32 {
	return q1.X*q2.X + q1.Y*q2.Y + q1.Z*q2.Z + q1.W*q2.W
}

func (q *Quat) Slerp(q1 Quat, q2 Quat, delta float32) {
	theta := q1.AngleBetween(q2)
	delta = delta / 2.0
	st := Sin(theta)
	sut := Sin(delta * theta)
	sout := Sin((1 - delta) * theta)
	coeff1 := sout / st
	coeff2 := sut / st
	q.X = coeff1*q1.X + coeff2*q2.Z
	q.Y = coeff1*q1.Y + coeff2*q2.Y
	q.Z = coeff1*q1.Z + coeff2*q2.Z
	q.W = coeff1*q1.W + coeff2*q2.W
	q.Normal()
}

func (start Quat) Lerp(end Quat, delta float32, out *Quat) {
	out.Slerp(start, end, delta)
}

func (q1 Quat) AngleBetween(q2 Quat) float32 {
	return Abs(Acos(q1.Dot(q2)))
}

func (q Quat) Distance(other Quat) float32 {
	return q.AngleBetween(other)
	// var i, qd Quat
	// i = q
	// i.Invert()
	// qd.Multiply(i, q)
	// lenSq := qd.X*qd.X + qd.Y*qd.Y + qd.Z*qd.Z
	// angle := 2 * Atan2(Sqrt(lenSq), qd.W)
	// return Abs(angle)
}

func (q Quat) DistanceSq(other Quat) float32 {
	return Sq(q.Distance(other))
}

func (q Quat) Add(value Quat, out *Quat) {
	out.X = q.X + value.X
	out.Y = q.Y + value.Y
	out.Z = q.Z + value.Z
	out.W = q.W + value.W
}

func (q Quat) Sub(subtrahend Quat, out *Quat) {
	q.Invert()
	out.Multiply(q, subtrahend)
}

func (q Quat) Scale(amount float32, out *Quat) {
	out.X = q.X * amount
	out.Y = q.Y * amount
	out.Z = q.Z * amount
	out.W = q.W * amount
}

func (q Quat) AddScaled(value Quat, amount float32, out *Quat) {
	out.X = q.X + value.X*amount
	out.Y = q.Y + value.Y*amount
	out.Z = q.Z + value.Z*amount
	out.W = q.W + value.W*amount
}

func (q Quat) Components() int {
	return 4
}

func (q Quat) GetComponent(index int) float32 {
	switch index {
	case 0:
		return q.X
	case 1:
		return q.Y
	case 2:
		return q.Z
	case 3:
		return q.W
	}
	return 0
}

func (q Quat) SetComponent(index int, value float32, out *Quat) {
	switch index {
	case 0:
		out.X = value
	case 1:
		out.Y = value
	case 2:
		out.Z = value
	case 3:
		out.W = value
	}
}

func (q Quat) SetComponents(value float32, out *Quat) {
	out.X = value
	out.Y = value
	out.Z = value
	out.W = value
}
