package geom

type Bounds[C BoundsCorner] struct {
	Min C
	Max C
}

type Bounds2f Bounds[Vec2f]
type Bounds2i Bounds[Vec2i]
type Bounds3f Bounds[Vec3f]
type Bounds3i Bounds[Vec3i]

func (b Bounds2f) Width() float32  { return b.Max.X - b.Min.X }
func (b Bounds2f) Height() float32 { return b.Max.Y - b.Min.Y }

func (b Bounds2i) Width() int32  { return b.Max.X - b.Min.X }
func (b Bounds2i) Height() int32 { return b.Max.Y - b.Min.Y }

func (b Bounds3f) Width() float32  { return b.Max.X - b.Min.X }
func (b Bounds3f) Height() float32 { return b.Max.Y - b.Min.Y }
func (b Bounds3f) Depth() float32  { return b.Max.Z - b.Min.Z }

func (b Bounds3i) Width() int32  { return b.Max.X - b.Min.X }
func (b Bounds3i) Height() int32 { return b.Max.Y - b.Min.Y }
func (b Bounds3i) Depth() int32  { return b.Max.Z - b.Min.Z }

type BoundsCorner interface {
}
