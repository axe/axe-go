package fx

import "github.com/axe/axe-go/pkg/util"

type SystemFormat struct {
	Format       *Format
	Initializers Inits
	Initializer  Init
	Modifiers    Modifys
	Modifier     Modify
	Emitter      EmitterType
	MaxParticles int
	Trail        bool
	Spread       bool
	Ordered      bool
}

func (sf SystemFormat) Inits(attr Attribute) bool {
	for _, init := range sf.Initializers {
		if init.Inits(attr) {
			return true
		}
	}
	return false
}

func (sf SystemFormat) Init(particle []float32, format *Format) {
	for _, init := range sf.Initializers {
		init.Init(particle, format)
	}
}

func (sf SystemFormat) Modifies(attr Attribute) bool {
	for _, modifier := range sf.Modifiers {
		if modifier.Modifies(attr) {
			return true
		}
	}
	return false
}

func (sf SystemFormat) Modify(particle []float32, format *Format, dt float32) {
	for _, modifier := range sf.Modifiers {
		modifier.Modify(particle, format, dt)
	}
}

func (sf *SystemFormat) Setup() {
	if sf.Initializer != nil && sf.Modifier != nil {
		return
	}

	if sf.Initializer == nil {
		inits := InitList{List: sf.Initializers[:]}
		for _, attr := range sf.Format.Attributes {
			if attr.Access != nil && attr.Attribute.init != nil && !inits.Inits(attr.Attribute) && sf.Format.HasData(attr.Attribute) {
				inits.List = append(inits.List, attr.Attribute.init)
			}
		}

		switch len(sf.Initializers) {
		case 0:
			sf.Initializer = InitNone{}
		case 1:
			sf.Initializer = inits.List[0]
		default:
			sf.Initializer = inits
		}
	}

	if sf.Modifier == nil {
		modifiers := ModifyList{List: sf.Modifiers[:]}
		for _, attr := range sf.Format.Attributes {
			if attr.Access != nil && sf.Modifier != nil && attr.Attribute.modify != nil && !modifiers.Modifies(attr.Attribute) {
				modifiers.List = append(modifiers.List, attr.Attribute.modify)
			}
		}

		switch len(sf.Modifiers) {
		case 0:
			sf.Modifier = ModifyNone{}
		case 1:
			sf.Modifier = modifiers.List[0]
		default:
			sf.Modifier = modifiers
		}
	}
}

func (sf *SystemFormat) Create() System {
	sf.Setup()

	return System{
		Format:  sf,
		Data:    NewData(sf.Format, sf.MaxParticles),
		Emitter: sf.Emitter.Create(),
	}
}

type System struct {
	Format  *SystemFormat
	Data    Data
	Emitter Emitter
}

func (s *System) Update(dt float32) {
	stop := s.Data.Count
	alive := 0
	if s.Format.Ordered {
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

	count, overTime := s.Emitter.Update(dt)
	add := util.Min(count, s.Data.Available())
	if add > 0 {
		timeCurrent := util.If(s.Format.Trail, -overTime, 0)
		timeAdd := util.If(s.Format.Spread, overTime/float32(add), 0)

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
	s.Format.Initializer.Init(s.Data.At(i), s.Data.Format)
}

func (s *System) UpdateAt(i int, dt float32) bool {
	particle := s.Data.At(i)
	format := s.Data.Format

	s.Format.Modifier.Modify(particle, format, dt)

	return Life(particle, format) < 1
}
