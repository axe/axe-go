package geom

type Vec2f Vec2[float32]
type Vec2i Vec2[int]
type Vec3f Vec3[float32]
type Vec3i Vec3[int]

type Vec2[D comparable] struct {
	X D
	Y D
}

type Vec3[D comparable] struct {
	X D
	Y D
	Z D
}
