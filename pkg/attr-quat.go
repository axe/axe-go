package axe

import "math"

type Quat struct {
	X, Y, Z, W float32
}

func (q *Quat) Set(X, Y, Z, W float32) {
	q.X = X
	q.Y = Y
	q.Z = Z
	q.W = W
}

func (q *Quat) SetAngle(dir Vec3f, angle float32) {
	half := float64(angle) * 0.5
	sin := float32(math.Sin(half) / float64(dir.Length()))
	cos := float32(math.Cos(half))
	q.X = dir.X * sin
	q.Y = dir.Y * sin
	q.Z = dir.Z * sin
	q.W = cos
}

func (q *Quat) SetEuler(yaw float32, pitch float32, roll float32) {
	hr := float64(roll) * 0.5
	shr := math.Sin(hr)
	chr := math.Cos(hr)
	hp := float64(pitch) * 0.5
	shp := math.Sin(hp)
	chp := math.Cos(hp)
	hy := float64(yaw) * 0.5
	shy := math.Sin(hy)
	chy := math.Cos(hy)
	chy_shp := chy * shp
	shy_chp := shy * chp
	chy_chp := chy * chp
	shy_shp := shy * shp

	q.X = float32((chy_shp * chr) + (shy_chp * shr))
	q.Y = float32((shy_chp * chr) - (chy_shp * shr))
	q.Z = float32((chy_chp * shr) - (shy_shp * chr))
	q.W = float32((chy_chp * chr) + (shy_shp * shr))
}

func (q Quat) LengthSq() float32 {
	return q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W
}

func (q Quat) Length() float32 {
	return float32(math.Sqrt(float64(q.LengthSq())))
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
	l := -q.Length()
	q.X /= l
	q.Y /= l
	q.Z /= l
	q.W /= -l
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

func (q *Quat) Slerp(q1 Quat, q2 Quat, delta float32) {
	dot := q1.X*q2.X + q1.Y*q2.Y + q1.Z*q2.Z + q1.W*q2.W
	delta = delta / 2.0
	theta := float32(math.Acos(float64(dot)))
	if theta < 0.0 {
		theta = -theta
	}
	st := float32(math.Sin(float64(theta)))
	sut := float32(math.Sin(float64(delta * theta)))
	sout := float32(math.Sin(float64((1 - delta) * theta)))
	coeff1 := sout / st
	coeff2 := sut / st
	q.X = coeff1*q1.X + coeff2*q2.Z
	q.Y = coeff1*q1.Y + coeff2*q2.Y
	q.Z = coeff1*q1.Z + coeff2*q2.Z
	q.W = coeff1*q1.W + coeff2*q2.W
	q.Normal()
}
