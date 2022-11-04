package axe

import "github.com/axe/axe-go/pkg/ecs"

type SpaceComponent[A Attr[A]] struct {
	Shape          Shape[A]
	Offset         A
	WorldTransform *Matrix[A]
	Flags          uint64
	Static         bool
	Inert          bool
}

type SpaceQuery[A Attr[A]] struct {
	Point   A
	End     A
	Shape   Shape[A]
	Maximum int
	Flags   uint64
	Match   ecs.FlagMatch
}

type SpaceNearest[E any] struct {
	Entity   E
	Distance float32
}

type SpaceSearchCallback[A Attr[A], E any] func(entity E, overlap float32, index int, query SpaceQuery[A]) bool

type SpaceCollisionCallback[E any] func(subject E, otherSubject E, overlap float32, index int, second bool)

type Space[A Attr[A], E any] interface {
	GameSystem
	Collisions(flags uint32, match ecs.FlagMatch, callback SpaceCollisionCallback[E])
	Intersects(query SpaceQuery[A], callback SpaceSearchCallback[A, E]) int
	Contains(query SpaceQuery[A], callback SpaceSearchCallback[A, E]) int
	Raytrace(query SpaceQuery[A], callback SpaceSearchCallback[A, E]) int
	KNN(query SpaceQuery[A], nearest []SpaceNearest[E], nearestCount *int)
}
