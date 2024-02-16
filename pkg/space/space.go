package space

import (
	"github.com/axe/axe-go/pkg/util"
)

type Flags uint64

type Dimensional interface {
	Dimensions() int
	Get(v []*float32)
}

type Shape[D Dimensional] interface {
	GetExtent(position D) (min, max D)
	GetCircle(position D) (center D, radius float32)
}

type Entity[D Dimensional] struct {
	Shape    Shape[D]
	Position D
	Flags    Flags
	Static   bool
	Disabled bool
	Expired  bool
	Update   func()
	Data     any
}

type Query[D Dimensional] struct {
	Point   D
	End     D
	Shape   Shape[D]
	Maximum int
	Flags   util.Match[Flags]
}

type Nearest[D Dimensional] struct {
	Entity   *Entity[D]
	Distance float32
}

type SearchCallback[D Dimensional] func(entity *Entity[D], overlap float32, index int, query Query[D]) bool

type CollisionCallback[D Dimensional] func(subject *Entity[D], otherSubject *Entity[D], overlap float32, index int, second bool)

type Space[D Dimensional] interface {
	Add(entity *Entity[D])
	Update()
	Collisions(flags util.Match[Flags], callback CollisionCallback[D])
	Intersects(query Query[D], callback SearchCallback[D]) int
	Contains(query Query[D], callback SearchCallback[D]) int
	Raytrace(query Query[D], callback SearchCallback[D]) int
	Nearest(query Query[D], nearest []Nearest[D]) int
}

type emptyDimensional struct{}

func (emptyDimensional) Dimensions() int  { return 0 }
func (emptyDimensional) Get(v []*float32) {}
