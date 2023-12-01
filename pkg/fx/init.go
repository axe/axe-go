package fx

import "math/rand"

type Init interface {
	Init(particle []float32, format *Format)
	Inits(attr Attribute) bool
}

type InitConstant struct {
	Attribute Attribute
	Constant  []float32
}

func (i InitConstant) Init(particle []float32, format *Format) {
	copy(format.Get(i.Attribute, particle), i.Constant)
}
func (i InitConstant) Inits(attr Attribute) bool {
	return i.Attribute.id == attr.id
}

type InitDynamic struct {
	Attribute Attribute
	Dynamic   Dynamic
}

func (i InitDynamic) Init(particle []float32, format *Format) {
	copy(format.Get(i.Attribute, particle), i.Dynamic(particle, format))
}
func (i InitDynamic) Inits(attr Attribute) bool {
	return i.Attribute.id == attr.id
}

type InitRandom struct {
	Attribute Attribute
	Start     []float32
	End       []float32
}

func (i InitRandom) Init(particle []float32, format *Format) {
	Lerp(i.Start, i.End, rand.Float32(), format.Get(i.Attribute, particle))
}
func (i InitRandom) Inits(attr Attribute) bool {
	return i.Attribute.id == attr.id
}

type InitNone struct{}

func (i InitNone) Init(particle []float32, format *Format) {

}
func (i InitNone) Inits(attr Attribute) bool {
	return false
}

type InitList struct {
	List []Init
}

func (i InitList) Init(particle []float32, format *Format) {
	for _, init := range i.List {
		init.Init(particle, format)
	}
}
func (i InitList) Inits(attr Attribute) bool {
	for _, init := range i.List {
		if init.Inits(attr) {
			return true
		}
	}
	return false
}

type Inits []Init

func (i Inits) Add(init Init) Inits {
	return append(i, init)
}

func (i Inits) Constant(attr Attribute, constant ...float32) Inits {
	return append(i, InitConstant{Attribute: attr, Constant: constant})
}

func (i Inits) Dynamic(attr Attribute, dynamic Dynamic) Inits {
	return append(i, InitDynamic{Attribute: attr, Dynamic: dynamic})
}

func (i Inits) Random(attr Attribute, start, end []float32) Inits {
	return append(i, InitRandom{Attribute: attr, Start: start, End: end})
}
