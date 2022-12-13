package axe

import "math"

type Camera[A Attr[A]] interface {
	GameSystem

	GetTransform() Trans[A]
	GetPlanes() []Plane[A]
	GetMatrix() Mat[A]
	// Intersects(shape Shape[A], position A) bool
}

type Camera2d struct {
	Transform Trans2d
	Size      Vec2d
	Near      Scalarf
	Far       Scalarf

	Up     Vec2d
	Right  Vec2d
	Left   Vec2d
	Down   Vec2d
	Min    Vec2d
	Max    Vec2d
	Aspect Scalarf
	Planes [4]Plane2d
	Matrix Mat2d
}

var _ Camera[Vec2d] = &Camera2d{}

func NewCamera2d() Camera2d {
	return Camera2d{
		Transform: Trans2d{
			Scale: Vec2f{X: 1, Y: 1},
		},
		Far: Scalarf{1},
	}
}

func (c *Camera2d) Init(game *Game) error {
	return nil
}
func (c *Camera2d) Update(game *Game) {
	hx := c.Size.X * 0.5
	hy := c.Size.Y * 0.5

	c.Aspect.Value = c.Size.X / c.Size.Y

	c.Min = c.Transform.Transform(Vec2d{X: -hx, Y: -hy})
	c.Max = c.Transform.Transform(Vec2d{X: hx, Y: hy})

	c.Up = c.Transform.TransformVector(Vec2d{X: 0, Y: -1})
	c.Down = c.Transform.TransformVector(Vec2d{X: 0, Y: 1})
	c.Right = c.Transform.TransformVector(Vec2d{X: 1, Y: 0})
	c.Left = c.Transform.TransformVector(Vec2d{X: -1, Y: 0})

	c.Planes[0].SetPointNormal(c.Min, c.Left)
	c.Planes[1].SetPointNormal(c.Min, c.Up)
	c.Planes[2].SetPointNormal(c.Max, c.Right)
	c.Planes[3].SetPointNormal(c.Max, c.Down)

	values := c.Matrix.Get()
	c.Transform.ToMatrix(values)
	c.Matrix.Set(values)
}
func (c *Camera2d) Destroy() {

}
func (c *Camera2d) GetTransform() Trans[Vec2d] {
	return &c.Transform
}
func (c *Camera2d) GetPlanes() []Plane[Vec2d] {
	return []Plane[Vec2d]{
		&c.Planes[0],
		&c.Planes[1],
		&c.Planes[2],
		&c.Planes[3],
	}
}
func (c *Camera2d) GetMatrix() Mat[Vec2d] {
	return &c.Matrix
}

type Camera3d struct {
	Transform Trans3d
	Size      Vec2d
	Near      Scalarf
	Far       Scalarf
	Fov       Scalarf

	Up       Vec3d
	Right    Vec3d
	Left     Vec3d
	Down     Vec3d
	Forward  Vec3d
	Backward Vec3d
	Min      Vec3d
	Max      Vec3d
	Aspect   Scalarf
	Planes   [6]Plane3d
	Matrix   Mat3d
}

var _ Camera[Vec3d] = &Camera3d{}

func NewCamera3d() Camera3d {
	return Camera3d{
		Transform: Trans3d{
			Scale: Vec3d{X: 1, Y: 1, Z: 1},
		},
		Near: Scalarf{0.001},
		Far:  Scalarf{1000.0},
		Fov:  Scalarf{60},
	}
}

func (c *Camera3d) Init(game *Game) error {
	return nil
}
func (c *Camera3d) Update(game *Game) {
	hx := c.Size.X * 0.5
	hy := c.Size.Y * 0.5

	c.Aspect.Value = c.Size.X / c.Size.Y

	c.Min = c.Transform.Transform(Vec3d{X: -hx, Y: -hy})
	c.Max = c.Transform.Transform(Vec3d{X: hx, Y: hy})

	c.Up = c.Transform.TransformVector(Vec3d{X: 0, Y: 1, Z: 0})
	c.Down = c.Transform.TransformVector(Vec3d{X: 0, Y: -1, Z: 0})
	c.Right = c.Transform.TransformVector(Vec3d{X: 1, Y: 0, Z: 0})
	c.Left = c.Transform.TransformVector(Vec3d{X: -1, Y: 0, Z: 0})
	c.Forward = c.Transform.TransformVector(Vec3d{X: 0, Y: 0, Z: 1})
	c.Backward = c.Transform.TransformVector(Vec3d{X: 0, Y: 0, Z: -1})

	var nearCenter, farCenter Vec3d
	c.Transform.Position.AddScaled(c.Forward, c.Near.Value, &nearCenter)
	c.Transform.Position.AddScaled(c.Forward, c.Far.Value, &farCenter)

	fovTheta := c.Fov.Value / 360 * math.Pi
	cos := Cos(fovTheta)
	sin := Sin(fovTheta)

	planeLeft := c.Transform.TransformVector(Vec3d{X: -cos, Y: 0, Z: -sin})
	planeUp := c.Transform.TransformVector(Vec3d{X: 0, Y: cos, Z: -sin})
	planeRight := c.Transform.TransformVector(Vec3d{X: cos, Y: 0, Z: -sin})
	planeDown := c.Transform.TransformVector(Vec3d{X: 0, Y: -cos, Z: -sin})

	c.Planes[0].SetPointNormal(c.Min, planeLeft)
	c.Planes[1].SetPointNormal(c.Min, planeUp)
	c.Planes[2].SetPointNormal(c.Max, planeRight)
	c.Planes[3].SetPointNormal(c.Max, planeDown)
	c.Planes[4].SetPointNormal(nearCenter, c.Backward)
	c.Planes[5].SetPointNormal(farCenter, c.Forward)

	// values := c.Matrix.Get()
	// c.Transform.ToMatrix(values)
	// c.Matrix.Set(values)
}
func (c *Camera3d) Destroy() {

}
func (c Camera3d) GetTransform() Trans[Vec3d] {
	return &c.Transform
}
func (c Camera3d) GetPlanes() []Plane[Vec3d] {
	return []Plane[Vec3d]{
		&c.Planes[0],
		&c.Planes[1],
		&c.Planes[2],
		&c.Planes[3],
		&c.Planes[4],
		&c.Planes[5],
	}
}
func (c *Camera3d) GetMatrix() Mat[Vec3d] {
	return &c.Matrix
}

type Trans[A Attr[A]] interface {
	GetScale() A
	GetPosition() A
	ToMatrix(m []float32)
	Transform(v A) A
	TransformVector(v A) A
}

type Trans2d struct {
	Scale    Vec2d
	Position Vec2d
	Rotation Vec2d
}

var _ Trans[Vec2d] = &Trans2d{}

func (t Trans2d) GetScale() Vec2d {
	return t.Scale
}
func (t Trans2d) GetPosition() Vec2d {
	return t.Position
}
func (t Trans2d) ToMatrix(m []float32) {
	m[0] = t.Rotation.X * t.Scale.X
	m[1] = t.Rotation.Y * t.Scale.X
	m[2] = 0
	m[3] = -t.Rotation.Y * t.Scale.Y
	m[4] = t.Rotation.X * t.Scale.Y
	m[5] = 0
	m[6] = t.Position.X
	m[7] = t.Position.Y
	m[8] = 1
}
func (t Trans2d) Transform(v Vec2d) Vec2d {
	return Vec2d{
		X: v.X*t.Rotation.X*t.Scale.X - v.Y*t.Rotation.Y*t.Scale.Y + t.Position.X,
		Y: v.X*t.Rotation.Y*t.Scale.X + v.Y*t.Rotation.X*t.Scale.Y + t.Position.Y,
	}
}
func (t Trans2d) TransformVector(v Vec2d) Vec2d {
	return Vec2d{
		X: v.X*t.Rotation.X*t.Scale.X - v.Y*t.Rotation.Y*t.Scale.Y,
		Y: v.X*t.Rotation.Y*t.Scale.X + v.Y*t.Rotation.X*t.Scale.Y,
	}
}

type Trans3d struct {
	Scale    Vec3d
	Position Vec3d
	Rotation Quat
}

var _ Trans[Vec3d] = &Trans3d{}

func (t Trans3d) GetScale() Vec3d {
	return t.Scale
}
func (t Trans3d) GetPosition() Vec3d {
	return t.Position
}
func (t Trans3d) ToMatrix(m []float32) {
	m[0] = t.Rotation.X * t.Scale.X
	m[1] = t.Rotation.Y
	m[2] = 0
	m[3] = -t.Rotation.Y
	m[4] = t.Rotation.X * t.Scale.Y
	m[5] = 0
	m[6] = t.Position.X
	m[7] = t.Position.Y
	m[8] = 1
}
func (t Trans3d) Transform(v Vec3d) Vec3d {
	return Vec3d{}
}
func (t Trans3d) TransformVector(v Vec3d) Vec3d {
	return Vec3d{}
}

type Mat[A Attr[A]] interface {
	Get() []float32
	Set(values []float32)
	Identity()
	Transform(v A) A
	Multiply(other Mat[A])
	Invert()
}

type Mat2d struct {
	v [9]float32
}

var _ Mat[Vec2d] = &Mat2d{}

func (m *Mat2d) Get() []float32 {
	return m.v[:]
}
func (m *Mat2d) Set(values []float32) {
	n := Min(len(values), 9)
	for i := 0; i < n; i++ {
		m.v[i] = values[i]
	}
}
func (m *Mat2d) Identity() {
	m.v[0] = 1
	m.v[1] = 0
	m.v[2] = 0
	m.v[3] = 0
	m.v[4] = 1
	m.v[5] = 0
	m.v[6] = 0
	m.v[7] = 0
	m.v[8] = 1
}
func (m *Mat2d) Transform(v Vec2d) Vec2d {
	invW := Div(1, m.v[2]*v.X+m.v[5]*v.Y+m.v[8])

	return Vec2d{
		X: (m.v[0]*v.X + m.v[3]*v.Y + m.v[6]) * invW,
		Y: (m.v[1]*v.X + m.v[4]*v.Y + m.v[7]) * invW,
	}
}
func (m *Mat2d) Multiply(other Mat[Vec2d]) {
	// v := other.Get()

}
func (m *Mat2d) Invert() {

}
func (m *Mat2d) Determinant() float32 {
	return m.v[0]*m.v[4]*m.v[8] +
		m.v[3]*m.v[6]*m.v[2] +
		m.v[6]*m.v[1]*m.v[5] -
		m.v[0]*m.v[6]*m.v[5] -
		m.v[3]*m.v[1]*m.v[8] -
		m.v[6]*m.v[4]*m.v[2]
}
func (m *Mat2d) Ortho(left, top, right, bottom float32) {
	dx := (right - left)
	dy := (top - bottom)
	xo := Div(2, dx)
	yo := Div(2, dy)
	tx := -(right + left) / dx
	ty := -(top + bottom) / dy

	m.v[0] = xo
	m.v[1] = 0
	m.v[2] = 0
	m.v[3] = 0
	m.v[4] = yo
	m.v[5] = 0
	m.v[6] = tx
	m.v[7] = ty
	m.v[8] = 1
}

type Mat3d struct {
	v [16]float32
}

var _ Mat[Vec3d] = &Mat3d{}

func (m *Mat3d) Get() []float32 {
	return m.v[:]
}
func (m *Mat3d) Set(values []float32) {
	n := Min(len(values), 9)
	for i := 0; i < n; i++ {
		m.v[i] = values[i]
	}
}
func (m *Mat3d) Identity() {
	m.v[0] = 1
	m.v[1] = 0
	m.v[2] = 0
	m.v[3] = 0
	m.v[4] = 0
	m.v[5] = 1
	m.v[6] = 0
	m.v[7] = 0
	m.v[8] = 0
	m.v[9] = 0
	m.v[10] = 1
	m.v[11] = 0
	m.v[12] = 0
	m.v[13] = 0
	m.v[14] = 0
	m.v[15] = 1
}
func (m *Mat3d) Transform(v Vec3d) Vec3d {
	return Vec3d{
		X: (v.X * m.v[0]) + (v.Y * m.v[4]) + (v.Z * m.v[8]) + m.v[12],
		Y: (v.X * m.v[1]) + (v.Y * m.v[5]) + (v.Z * m.v[9]) + m.v[13],
		Z: (v.X * m.v[2]) + (v.Y * m.v[6]) + (v.Z * m.v[10]) + m.v[14],
	}
}
func (m *Mat3d) TransformVector(v Vec3d) Vec3d {
	return Vec3d{
		X: (v.X * m.v[0]) + (v.Y * m.v[4]) + (v.Z * m.v[8]),
		Y: (v.X * m.v[1]) + (v.Y * m.v[5]) + (v.Z * m.v[9]),
		Z: (v.X * m.v[2]) + (v.Y * m.v[6]) + (v.Z * m.v[10]),
	}
}
func (m *Mat3d) Multiply(other Mat[Vec3d]) {
	// a := m.v
	// b := other.(*Mat3d).v

	// m.v[0] = (a.v[0] * b00) + (a.v[4] * b10) + (a02 * b20)
	// m.v[1] = (a10 * b00) + (a11 * b10) + (a12 * b20)
	// m.v[2] = (a20 * b00) + (a21 * b10) + (a22 * b20)
	// m.v[4] = (a.v[0] * b01) + (a.v[4] * b11) + (a02 * b21)
	// m.v[5] = (a10 * b01) + (a11 * b11) + (a12 * b21)
	// m.v[6] = (a20 * b01) + (a21 * b11) + (a22 * b21)
	// m.v[8] = (a.v[0] * b02) + (a.v[4] * b12) + (a02 * b22)
	// m.v[9] = (a10 * b02) + (a11 * b12) + (a12 * b22)
	// m.v[10] = (a20 * b02) + (a21 * b12) + (a22 * b22)
	// m.v[12] = (a.v[0] * b.tx) + (a.v[4] * b.ty) + (am02 * b.tz) + atx
	// m.v[13] = (am10 * b.tx) + (am11 * b.ty) + (am12 * b.tz) + aty
	// m.v[14] = (am20 * b.tx) + (am21 * b.ty) + (am22 * b.tz) + atz
}
func (m *Mat3d) Invert() {

}
func (m *Mat3d) Determinant() float32 {
	return m.v[0]*m.v[4]*m.v[8] +
		m.v[3]*m.v[6]*m.v[2] +
		m.v[6]*m.v[1]*m.v[5] -
		m.v[0]*m.v[6]*m.v[5] -
		m.v[3]*m.v[1]*m.v[8] -
		m.v[6]*m.v[4]*m.v[2]
}
func (m *Mat3d) Perpsective(fov, aspectRatio, near, far float32) {
	fovRad := fov * 2.0 * math.Pi / 360.0
	focalLength := float32(1.0 / math.Tan(float64(fovRad/2.0)))

	x := focalLength / aspectRatio
	y := -focalLength
	A := near / (far - near)
	B := far * A

	m.v = [16]float32{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, A, -1,
		0, 0, B, 0,
	}
}
func (m *Mat3d) Ortho(left, right, bottom, top float32) {
	x := 2 / (right - left)
	y := 2 / (top - bottom)
	A := -(right + left) / (right - left)
	B := -(top + bottom) / (top - bottom)

	m.v = [16]float32{
		x, 0, 0, A,
		0, y, 0, B,
		0, 0, -1, 0,
		0, 0, 0, 1,
	}
}

type Plane[A Attr[A]] interface {
	SetPointNormal(point A, normal A)
	SetPoints(points []A)
	Distance(point A) float32
	GetPoint() A
	GetNormal() A
	Sign(point A) PlaneSign
}

type PlaneSign int

const (
	PlaneSignBack PlaneSign = iota - 1
	PlaneSignOn
	PlaneSignFront
	PlaneSignIntersects
)

type Plane2d struct {
	A, B, C float32
}

var _ Plane[Vec2d] = &Plane2d{}

func (p *Plane2d) SetPointNormal(point Vec2d, normal Vec2d) {
	p.A = normal.X
	p.B = normal.Y
	p.C = -((normal.X * point.X) + (normal.Y * point.Y))
}
func (p *Plane2d) SetPoints(points []Vec2d) {
	if len(points) < 2 {
		return
	}
	start := points[0]
	end := points[1]
	dx := (end.X - start.X)
	dy := (end.Y - start.Y)
	d := Div(1.0, Sqrt(dx*dx+dy*dy))

	p.A = -dy * d
	p.B = dx * d
	p.C = -(p.A*start.X + p.B*start.Y)
}
func (p *Plane2d) Distance(point Vec2d) float32 {
	return (p.A*point.X + p.B*point.Y + p.C)
}
func (p *Plane2d) Sign(point Vec2d) PlaneSign {
	return PlaneSign(Sign(p.Distance(point)))
}
func (p *Plane2d) Intersection(plane Plane[Vec2d], out *Vec2d) bool {
	other := plane.(*Plane2d)
	div := (p.A*other.B - p.B*other.A)

	if div == 0 {
		return false
	}

	div = 1 / div
	out.X = (-p.C*other.B + p.B*other.C) * div
	out.Y = (-p.A*other.C + p.C*other.A) * div
	return true
}
func (p *Plane2d) GetPoint() Vec2d {
	return Vec2d{X: p.A * p.C, Y: p.B * p.C}
}
func (p *Plane2d) GetNormal() Vec2d {
	return Vec2d{X: p.A, Y: p.B}
}

type Plane3d struct {
	A, B, C, D float32
}

var _ Plane[Vec3d] = &Plane3d{}

func (p *Plane3d) SetPointNormal(point Vec3d, normal Vec3d) {
	p.A = normal.X
	p.B = normal.Y
	p.C = normal.Z
	p.D = -((normal.X * point.X) + (normal.Y * point.Y) + (normal.Y * point.Z))
}
func (p *Plane3d) SetPoints(points []Vec3d) {
	if len(points) < 3 {
		return
	}
	var d0, d1, cross Vec3d
	p0 := points[0]
	p1 := points[1]
	p2 := points[2]
	p0.Sub(p1, &d0)
	p2.Sub(p1, &d1)
	d0.Cross(d1, &cross)
	Normalize(cross, &cross)

	p.SetPointNormal(p1, cross)
}
func (p *Plane3d) Distance(point Vec3d) float32 {
	return p.A*point.X + p.B*point.Y + p.C*point.Z + p.D
}
func (p *Plane3d) DistanceRound(round Round[Vec3d]) float32 {
	d := p.Distance(round.Center)
	return d - float32(Sign(d))*round.Radius
}
func (p *Plane3d) DistanceLine(line Line[Vec3d]) float32 {
	ds := p.Distance(line.Start)
	de := p.Distance(line.End)
	ss := float32(Sign(ds))
	se := float32(Sign(de))
	if ss != se {
		return 0
	}
	return Min(ds*ss, de*se)
}
func (p *Plane3d) Sign(point Vec3d) PlaneSign {
	return PlaneSign(Sign(p.Distance(point)))
}
func (p *Plane3d) SignPoints(points []Vec3d) PlaneSign {
	n := len(points)
	if n < 1 {
		return PlaneSignBack
	}
	s := p.Sign(points[0])
	for i := 1; i < n; i++ {
		if p.Sign(points[i]) != s {
			return PlaneSignIntersects
		}
	}
	return s
}
func (p *Plane3d) Intersection(plane Plane[Vec3d], out *Line[Vec3d]) bool {
	// other := plane.(*Plane3d)
	return true
}
func (p *Plane3d) GetPoint() Vec3d {
	return Vec3d{X: p.A * p.D, Y: p.B * p.D, Z: p.C * p.D, W: 1}
}
func (p *Plane3d) GetNormal() Vec3d {
	return Vec3d{X: p.A, Y: p.B, Z: p.C, W: 0}
}
func (p *Plane3d) Clip(line Line[Vec3d], side PlaneSign) *Line[Vec3d] {
	dir := line.Diff()
	normal := p.GetNormal()

	d := normal.Dot(dir)
	sd := p.Distance(line.Start)
	q := PlaneSign(Sign(sd))

	if d == 0.0 {
		if q == side || p.Sign(line.End) == side {
			return &line
		}
		return nil
	}

	u := -sd / d

	if u >= 1.0 || u <= 0.0 {
		if q == side || p.Sign(line.End) == side {
			return &line
		}
		return nil
	}

	if q == side || q == PlaneSignOn {
		line.End.Interpolate(line.Start, line.End, u, &line.End)
	} else {
		line.Start.Interpolate(line.Start, line.End, u, &line.Start)
	}

	return &line
}
