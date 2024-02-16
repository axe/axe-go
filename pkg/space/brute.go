package space

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/util"
)

type Brute[D Dimensional] struct {
	live     ds.List[*Entity[D]]
	disabled ds.List[*Entity[D]]
}

var _ Space[emptyDimensional] = &Brute[emptyDimensional]{}

func (s *Brute[D]) Add(entity *Entity[D]) {
	if entity.Expired {
		return
	}
	if entity.Disabled {
		s.addDisabled(entity)
	} else {
		s.addLive(entity)
	}
}
func (s *Brute[D]) addDisabled(entity *Entity[D]) {
	s.disabled.Reserve(1)
	s.disabled.Add(entity)
}
func (s *Brute[D]) addLive(entity *Entity[D]) {
	s.live.Reserve(1)
	s.live.Add(entity)
}
func (s *Brute[D]) Update() {
	disabledCount := 0
	for i := 0; i < s.disabled.Size; i++ {
		d := s.disabled.Get(i)
		if d.Expired {
			// let it be overwritten
		} else if !d.Disabled {
			s.addLive(d)
		} else {
			s.disabled.Items[disabledCount] = d
			disabledCount++
		}
	}
	for s.disabled.Size > disabledCount {
		s.disabled.Size--
		s.disabled.Items[s.disabled.Size] = nil
	}
	liveCount := 0
	for i := 0; i < s.live.Size; i++ {
		l := s.live.Get(i)
		if l.Expired {
			// let it be overwritten
		} else if l.Disabled {
			s.addDisabled(l)
		} else {
			s.live.Items[liveCount] = l
			liveCount++
		}
	}
	for s.live.Size > liveCount {
		s.live.Size--
		s.live.Items[s.live.Size] = nil
	}
}
func (s *Brute[D]) Collisions(flags util.Match[Flags], callback CollisionCallback[D]) {

}
func (s *Brute[D]) Intersects(query Query[D], callback SearchCallback[D]) int {
	return 0
}
func (s *Brute[D]) Contains(query Query[D], callback SearchCallback[D]) int {
	return 0
}
func (s *Brute[D]) Raytrace(query Query[D], callback SearchCallback[D]) int {
	return 0
}
func (s *Brute[D]) Nearest(query Query[D], nearest []Nearest[D]) int {
	return 0
}
