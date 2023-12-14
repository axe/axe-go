package ui

import (
	"math"

	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/util"
)

func Identity() Transform {
	return Transform{sx: 1, ry: 0, rx: 0, sy: 1, tx: 0, ty: 0}
}

type Transform struct {
	sx, ry, rx, sy, tx, ty float32
}

func (t *Transform) Set(sx, ry, rx, sy, tx, ty float32) {
	t.sx = sx
	t.ry = ry
	t.rx = rx
	t.sy = sy
	t.tx = tx
	t.ty = ty
}

func (t *Transform) Identity() {
	t.Set(1, 0, 0, 1, 0, 0)
}

func (t Transform) IsIdentity() bool {
	return t.rx == 0 && t.ry == 0 && t.sx == 1 && t.sy == 1 && t.tx == 0 && t.ty == 0
}

func (t Transform) IsEffectivelyIdentity() bool {
	return util.Equal(t.rx, 0) && util.Equal(t.ry, 0) && util.Equal(t.sx, 1) && util.Equal(t.sy, 1) && util.Equal(t.tx, 0) && util.Equal(t.ty, 0)
}

func (t Transform) IsZero() bool {
	return t.rx == 0 && t.ry == 0 && t.sx == 0 && t.sy == 0 && t.tx == 0 && t.ty == 0
}

func (t Transform) HasAffect() bool {
	return !(t.IsZero() || t.IsIdentity())
}

func (t *Transform) Transform(x, y float32) (float32, float32) {
	return t.sx*x + t.rx*y + t.tx, t.ry*x + t.sy*y + t.ty
}

func (t *Transform) TransformCoord(c gfx.Coord) gfx.Coord {
	x, y := t.Transform(c.X, c.Y)
	return gfx.Coord{X: x, Y: y}
}

func (t *Transform) TransformVector(vx, vy float32) (float32, float32) {
	return t.sx*vx + t.rx*vy, t.ry*vx + t.sy*vy
}

func (t *Transform) TransformVectorCoord(v gfx.Coord) gfx.Coord {
	x, y := t.TransformVector(v.X, v.Y)
	return gfx.Coord{X: x, Y: y}
}

func (t *Transform) Determinant() float32 {
	return t.sx*t.sy - t.ry*t.rx
}

func (t *Transform) Multiply(b Transform) {
	t.Set(
		t.sx*b.sx+t.rx*b.ry,
		t.sx*b.rx+t.rx*b.sy,
		t.ry*b.sx+t.sy*b.ry,
		t.ry*b.rx+t.sy*b.sy,
		t.tx*b.sx+t.ty*b.ry+b.tx,
		t.tx*b.rx+t.ty*b.sy+b.ty,
	)
}

func (t *Transform) Invert() {
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

func (t *Transform) GetInvert() Transform {
	copy := *t
	copy.Invert()
	return copy
}

func (t *Transform) GetRadians() float32 {
	return util.Atan2(t.ry, t.rx)
}

func (t *Transform) GetDegrees() float32 {
	return t.GetRadians() * 180 / math.Pi
}

func (t *Transform) GetScale() (x, y float32) {
	x = float32(math.Sqrt(float64(t.sx*t.sx + t.rx*t.rx)))
	y = float32(math.Sqrt(float64(t.ry*t.ry + t.sy*t.sy)))
	return
}

func (t *Transform) GetTranslate() (float32, float32) {
	return t.tx, t.ty
}

func (t *Transform) Translate(tx, ty float32) {
	t.tx += tx*t.sx + ty*t.rx
	t.ty += tx*t.ry + ty*t.sy
}

func (t *Transform) PreTranslate(tx, ty float32) {
	t.tx += tx
	t.ty += ty
}

func (t *Transform) SetTranslate(tx, ty float32) {
	t.Set(1, 0, 0, 1, tx, ty)
}

func (t *Transform) Scale(sx, sy float32) {
	t.sx *= sx
	t.sy *= sy
	t.rx *= sy
	t.ry *= sx
}

func (t *Transform) SetScale(sx, sy float32) {
	t.Set(sx, 0, 0, sy, 0, 0)
}

func (t *Transform) SetScaleAround(sx, sy, anchorX, anchorY float32) {
	t.Set(sx, 0, 0, sy, anchorX-anchorX*sx, anchorY-anchorY*sy)
}

func (t *Transform) PreScale(sx, sy float32) {
	t.sx *= sx
	t.sy *= sy
}

func (t *Transform) Shear(sx, sy float32) {
	tsx := t.sx
	try := t.ry
	t.sx = tsx + t.rx*sy
	t.rx = tsx*sx + t.rx
	t.ry = try + t.sy*sy
	t.sy = try*sx + t.sy
}

func (t *Transform) SetShear(sx, sy float32) {
	t.Set(1, sy, sx, 1, 0, 0)
}

func (t *Transform) SetShearAround(sx, sy, anchorX, anchorY float32) {
	t.Set(1, sy, sx, 1, -anchorY*sx, -anchorX*sy)
}

func (t *Transform) RotateDegrees(degrees float32) {
	t.Rotate(degrees * math.Pi / 180)
}

func (t *Transform) Rotate(radians float32) {
	cos, sin := util.CosSin(radians)
	tsx := t.sx
	try := t.ry
	t.sx = tsx*cos + t.rx*sin
	t.rx = tsx*-sin + t.rx*cos
	t.ry = try*cos + t.sy*sin
	t.sy = try*-sin + t.sy*cos
}

func (t *Transform) SetRotateDegrees(degrees float32) {
	t.SetRotate(degrees * math.Pi / 180)
}

func (t *Transform) SetRotate(radians float32) {
	cos, sin := util.CosSin(radians)
	t.Set(cos, sin, -sin, cos, 0, 0)
}

func (t *Transform) SetRotateDegreesAround(radians, anchorX, anchorY float32) {
	t.SetRotateAround(radians*math.Pi/180, anchorX, anchorY)
}

func (t *Transform) SetRotateAround(radians, anchorX, anchorY float32) {
	cos, sin := util.CosSin(radians)
	tx := anchorX - anchorX*cos + anchorY*sin
	ty := anchorY - anchorX*sin - anchorY*cos
	t.Set(cos, sin, -sin, cos, tx, ty)
}

func (t *Transform) SetRotateDegreesScaleAround(degrees, scaleX, scaleY, anchorX, anchorY float32) {
	t.SetRotateScaleAround(degrees*math.Pi/180, scaleX, scaleY, anchorX, anchorY)
}

func (t *Transform) SetRotateScaleAround(radians, scaleX, scaleY, anchorX, anchorY float32) {
	cos, sin := util.CosSin(radians)
	sx := cos * scaleX
	ry := sin * scaleX
	rx := -sin * scaleY
	sy := cos * scaleY
	tx := anchorX - anchorX*sx - anchorY*rx
	ty := anchorY - anchorX*ry - anchorY*sy
	t.Set(sx, ry, rx, sy, tx, ty)
}
