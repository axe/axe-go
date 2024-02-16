package shape

import (
	"math"

	"github.com/axe/axe-go/pkg/util"
)

func Identity() Transform2 {
	return Transform2{sx: 1, ry: 0, rx: 0, sy: 1, tx: 0, ty: 0}
}

type Transform2 struct {
	sx, ry, rx, sy, tx, ty float32
}

func (t *Transform2) Set(sx, ry, rx, sy, tx, ty float32) {
	t.sx = sx
	t.ry = ry
	t.rx = rx
	t.sy = sy
	t.tx = tx
	t.ty = ty
}

func (t *Transform2) Identity() {
	t.Set(1, 0, 0, 1, 0, 0)
}

func (t Transform2) IsIdentity() bool {
	return t.rx == 0 && t.ry == 0 && t.sx == 1 && t.sy == 1 && t.tx == 0 && t.ty == 0
}

func (t Transform2) IsEffectivelyIdentity() bool {
	return util.Equal(t.rx, 0) && util.Equal(t.ry, 0) && util.Equal(t.sx, 1) && util.Equal(t.sy, 1) && util.Equal(t.tx, 0) && util.Equal(t.ty, 0)
}

func (t Transform2) IsZero() bool {
	return t.rx == 0 && t.ry == 0 && t.sx == 0 && t.sy == 0 && t.tx == 0 && t.ty == 0
}

func (t Transform2) HasAffect() bool {
	return !(t.IsZero() || t.IsIdentity())
}

func (t *Transform2) Transform(x, y float32) (float32, float32) {
	return t.sx*x + t.rx*y + t.tx, t.ry*x + t.sy*y + t.ty
}

func (t *Transform2) TransformPoint(c Point2) Point2 {
	x, y := t.Transform(c.X, c.Y)
	return Point2{X: x, Y: y}
}

func (t *Transform2) TransformVector(vx, vy float32) (float32, float32) {
	return t.sx*vx + t.rx*vy, t.ry*vx + t.sy*vy
}

func (t *Transform2) TransformVectorPoint(v Point2) Point2 {
	x, y := t.TransformVector(v.X, v.Y)
	return Point2{X: x, Y: y}
}

func (t *Transform2) Determinant() float32 {
	return t.sx*t.sy - t.ry*t.rx
}

func (t *Transform2) Multiply(b Transform2) {
	t.Set(
		t.sx*b.sx+t.rx*b.ry,
		t.sx*b.rx+t.rx*b.sy,
		t.ry*b.sx+t.sy*b.ry,
		t.ry*b.rx+t.sy*b.sy,
		t.tx*b.sx+t.ty*b.ry+b.tx,
		t.tx*b.rx+t.ty*b.sy+b.ty,
	)
}

func (t *Transform2) Invert() {
	d := t.Determinant()

	if d == 0 {
		return
	}

	invd := 1.0 / d

	t.Set(
		invd*t.sy,
		invd*-t.ry,
		invd*-t.rx,
		invd*t.sx,
		invd*(t.ry*t.ty-t.tx*t.sy),
		invd*(t.tx*t.rx-t.sx*t.ty),
	)
}

func (t *Transform2) GetInvert() Transform2 {
	copy := *t
	copy.Invert()
	return copy
}

func (t *Transform2) GetRadians() float32 {
	return util.Atan2(t.ry, t.rx)
}

func (t *Transform2) GetDegrees() float32 {
	return t.GetRadians() * 180 / math.Pi
}

func (t *Transform2) GetScale() (x, y float32) {
	x = float32(math.Sqrt(float64(t.sx*t.sx + t.rx*t.rx)))
	y = float32(math.Sqrt(float64(t.ry*t.ry + t.sy*t.sy)))
	return
}

func (t *Transform2) GetTranslate() (float32, float32) {
	return t.tx, t.ty
}

func (t *Transform2) Translate(tx, ty float32) {
	t.tx += tx*t.sx + ty*t.rx
	t.ty += tx*t.ry + ty*t.sy
}

func (t *Transform2) PreTranslate(tx, ty float32) {
	t.tx += tx
	t.ty += ty
}

func (t *Transform2) SetTranslate(tx, ty float32) {
	t.Set(1, 0, 0, 1, tx, ty)
}

func (t *Transform2) Scale(sx, sy float32) {
	t.sx *= sx
	t.sy *= sy
	t.rx *= sy
	t.ry *= sx
}

func (t *Transform2) SetScale(sx, sy float32) {
	t.Set(sx, 0, 0, sy, 0, 0)
}

func (t *Transform2) SetScaleAround(sx, sy, anchorX, anchorY float32) {
	t.Set(sx, 0, 0, sy, anchorX-anchorX*sx, anchorY-anchorY*sy)
}

func (t *Transform2) PreScale(sx, sy float32) {
	t.sx *= sx
	t.sy *= sy
}

func (t *Transform2) Shear(sx, sy float32) {
	tsx := t.sx
	try := t.ry
	t.sx = tsx + t.rx*sy
	t.rx = tsx*sx + t.rx
	t.ry = try + t.sy*sy
	t.sy = try*sx + t.sy
}

func (t *Transform2) SetShear(sx, sy float32) {
	t.Set(1, sy, sx, 1, 0, 0)
}

func (t *Transform2) SetShearAround(sx, sy, anchorX, anchorY float32) {
	t.Set(1, sy, sx, 1, -anchorY*sx, -anchorX*sy)
}

func (t *Transform2) RotateDegrees(degrees float32) {
	t.Rotate(degrees * math.Pi / 180)
}

func (t *Transform2) Rotate(radians float32) {
	cos, sin := util.CosSin(radians)
	tsx := t.sx
	try := t.ry
	t.sx = tsx*cos + t.rx*sin
	t.rx = tsx*-sin + t.rx*cos
	t.ry = try*cos + t.sy*sin
	t.sy = try*-sin + t.sy*cos
}

func (t *Transform2) SetRotateDegrees(degrees float32) {
	t.SetRotate(degrees * math.Pi / 180)
}

func (t *Transform2) SetRotate(radians float32) {
	cos, sin := util.CosSin(radians)
	t.Set(cos, sin, -sin, cos, 0, 0)
}

func (t *Transform2) SetRotateDegreesAround(radians, anchorX, anchorY float32) {
	t.SetRotateAround(radians*math.Pi/180, anchorX, anchorY)
}

func (t *Transform2) SetRotateAround(radians, anchorX, anchorY float32) {
	cos, sin := util.CosSin(radians)
	tx := anchorX - anchorX*cos + anchorY*sin
	ty := anchorY - anchorX*sin - anchorY*cos
	t.Set(cos, sin, -sin, cos, tx, ty)
}

func (t *Transform2) SetRotateDegreesScaleAround(degrees, scaleX, scaleY, anchorX, anchorY float32) {
	t.SetRotateScaleAround(degrees*math.Pi/180, scaleX, scaleY, anchorX, anchorY)
}

func (t *Transform2) SetRotateScaleAround(radians, scaleX, scaleY, anchorX, anchorY float32) {
	cos, sin := util.CosSin(radians)
	sx := cos * scaleX
	ry := sin * scaleX
	rx := -sin * scaleY
	sy := cos * scaleY
	tx := anchorX - anchorX*sx - anchorY*rx
	ty := anchorY - anchorX*ry - anchorY*sy
	t.Set(sx, ry, rx, sy, tx, ty)
}
