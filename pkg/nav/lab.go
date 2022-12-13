package nav

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/util"
)

/*
- Nodes can have a cost to control preference
- Nodes can move (ex: a shifting platform)
- Nodes & Edges can have types which can control what type of agents can use it
- Nodes & Edges can have size which is a rough representation of how close a follower needs to be to it to be considered on it
- Nodes & Edges can become intermittently unavailable or congested where the cost is driven up
- Nodes & Edges can be baked or dynamic (baked for unaligned 3d world, dynamic for aligned 2d/3d grid worlds)
- Edges have a cost which is often distance between edges
- Edges have a start and end node
- Paths are a collection of Edges
- Travel time should factor into path determination, a node/edge could be aware of approximate times agents will be using it so we can project if its a good path to use
- Nodes can have a shape (point, sphere, area)
- Edges have a size which is the space an agent has to go from one node to another
- Edges can be one way or bi-directional
- Edges can be thought of as the lines where two nodes touch
- Paths can be fully, partially, or unable to be computed
- Paths can be requested one frame but not answered until later
- Path requests can be prioritized dynamically (updated each frame before picking ones to try to solve)
- Map is configured to solve a maximum of X steps per frame
- Map + Space can return a random node that is not visible by other agents
- Map keeps track of agents
- Steering AI updates the agent position and notifies the Map
- AI (Space & Nav) needs to answer the following questions
	- What agents can be seen from a shape?
	- What agents can be heard from a shape?
	- A seek behavior looks for perceived agents periodically and goes towards highest priority
	- A hide behavior looks for perceived agents periodically and goes to a spot with the lowest perceive value
- Agents can perceive others (0=none,1=fully) and visibility, sound, and recent perception play into it.
*/

type Node[V any] struct {
	Shape Shape[V]
	Cost  int
	Types uint64
	Size  float32
	Edges []int
}

type Edge struct {
	Start int
	End   int
	Cost  int
	Types uint64
	Size  float32
}

type Path struct {
	Edges []int
	Cost  int
}

type ShapeDistance struct {
	Min float32
	Max float32
}

func (sr ShapeDistance) IsOutside() bool      { return sr.Min >= 0 && sr.Max >= 0 }
func (sr ShapeDistance) InIntersecting() bool { return sr.Min < 0 && sr.Max >= 0 }
func (sr ShapeDistance) IsInside() bool       { return sr.Min < 0 && sr.Max < 0 }

type Shape[V any] interface {
	Distance(shape Shape[V]) ShapeDistance
}

// Shapes: Point, Ray, AABB, Sphere, Plane

type SpaceKnn[V any] struct {
	Node     Node[V]
	Distance float32
	Near     bool
}

type SpaceListener[V any] func(shape Shape[V], dist ShapeDistance) bool

type Space[V any] interface {
	Add(node Node[V])
	Knn(shape Shape[V], out []SpaceKnn[V])
	Relative(shape Shape[V], dist ShapeDistance, listener SpaceListener[V])
}

type Mesh[V any] struct {
	Nodes ds.SparseList[Node[V]]
	Edges ds.SparseList[Edge]
}

func (m Mesh[V]) FindPath(start V, end V, match util.Match[uint64]) Path {
	return Path{}
}
