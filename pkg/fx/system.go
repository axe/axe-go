package fx

import (
	"math/rand"

	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/util"
)

type SystemType struct {
	Format       *Format
	Initializers Inits
	Initializer  Init
	Modifiers    Modifys
	Modifier     Modify
	Emitter      EmitterType
	Random       *rand.Rand
	Renderer     Renderer
	MaxParticles int
	Trail        bool
	Spread       bool
	Ordered      bool
}

func (st SystemType) Inits(attr Attribute) bool {
	for _, init := range st.Initializers {
		if init.Inits(attr) {
			return true
		}
	}
	return false
}

func (st SystemType) Init(particle []float32, format *Format) {
	for _, init := range st.Initializers {
		init.Init(particle, format)
	}
}

func (st SystemType) Modifies(attr Attribute) bool {
	for _, modifier := range st.Modifiers {
		if modifier.Modifies(attr) {
			return true
		}
	}
	return false
}

func (st SystemType) Modify(particle []float32, format *Format, dt float32) {
	for _, modifier := range st.Modifiers {
		modifier.Modify(particle, format, dt)
	}
}

func (st *SystemType) GetModifier() Modify {
	if st.Modifier == nil {
		modifiers := ModifyList{List: st.Modifiers[:]}
		for _, attr := range st.Format.Attributes {
			if attr.Access != nil && st.Modifier != nil && attr.Attribute.modify != nil && !modifiers.Modifies(attr.Attribute) {
				modifiers.List = append(modifiers.List, attr.Attribute.modify)
			}
		}

		switch len(modifiers.List) {
		case 0:
			st.Modifier = ModifyNone{}
		case 1:
			st.Modifier = modifiers.List[0]
		default:
			st.Modifier = modifiers
		}
	}

	return st.Modifier
}

func (st *SystemType) GetInitializer() Init {
	if st.Initializer == nil {
		inits := InitList{List: st.Initializers[:]}
		for _, attr := range st.Format.Attributes {
			if attr.Access != nil && attr.Attribute.init != nil && !inits.Inits(attr.Attribute) && st.Format.HasData(attr.Attribute) {
				inits.List = append(inits.List, attr.Attribute.init)
			}
		}

		switch len(inits.List) {
		case 0:
			st.Initializer = InitNone{}
		case 1:
			st.Initializer = inits.List[0]
		default:
			st.Initializer = inits
		}
	}

	return st.Initializer
}

func (st *SystemType) Setup() {
	st.GetInitializer()
	st.GetModifier()
	if st.Random == nil {
		st.Random = rand.New(rand.NewSource(rand.Int63()))
	}
}

func (st *SystemType) Create() System {
	st.Setup()

	return System{
		Type:    st,
		Random:  st.Random,
		Data:    NewData(st.Format, st.MaxParticles),
		Emitter: st.Emitter.Create(),
		Buffer:  st.Renderer.CreateBuffer(st.MaxParticles),
	}
}

type System struct {
	Type    *SystemType
	Random  *rand.Rand
	Data    Data
	Emitter Emitter
	Buffer  *gfx.Buffer
}

func (s *System) Render() {
	s.Buffer.Data.Clear()
	stop := s.Data.Count
	for i := 0; i < stop; i++ {
		// p :=
	}
}

func (s *System) Update(dt float32) {
	s.updateParticles(dt)
	s.updateEmitter(dt)
}

func (s *System) updateParticles(dt float32) {
	stop := s.Data.Count
	alive := 0
	if s.Type.Ordered {
		for i := 0; i < stop; i++ {
			if s.UpdateAt(i, dt) {
				s.Data.Move(i, alive)
				alive++
			}
		}
	} else {
		for i := 0; i < stop; i++ {
			if s.UpdateAt(i, dt) {
				alive++
			} else {
				stop--
				i--
				s.Data.Move(stop, alive)
			}
		}
	}
	s.Data.Count = alive
}

func (s *System) updateEmitter(dt float32) {
	count, overTime := s.Emitter.Update(dt, s.Random)
	add := util.Min(count, s.Data.Available())

	if add > 0 {
		timeCurrent := util.If(s.Type.Trail, -overTime, 0)
		timeAdd := util.If(s.Type.Spread && s.Type.Trail, overTime/float32(add), 0)

		for i := 0; i < add; i++ {
			s.Emit(timeCurrent)
			timeCurrent += timeAdd
		}
	}
}

func (s *System) Done() bool {
	return s.Data.Count == 0 && s.Emitter.Done()
}

func (s *System) Emit(update float32) {
	i := s.Data.Count
	s.InitAt(i)
	if update != 0 {
		s.UpdateAt(i, update)
	}
	s.Data.Count++
}

func (s *System) InitAt(i int) {
	s.Type.Initializer.Init(s.Data.At(i), s.Data.Format)
}

func (s *System) UpdateAt(i int, dt float32) bool {
	particle := s.Data.At(i)
	format := s.Data.Format

	s.Type.Modifier.Modify(particle, format, dt)

	return Life(particle, format) < 1
}
