package geom

type Bounds2f Bounds[Vec2f]
type Bounds2i Bounds[Vec2i]
type Bounds3f Bounds[Vec3f]
type Bounds3i Bounds[Vec3i]

type Bounds[C BoundsCorner] struct {
	Min C
	Max C
}

type BoundsCorner interface {
}
