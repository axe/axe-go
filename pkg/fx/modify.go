package fx

type Modify interface {
	Modify(particle []float32, format *Format, dt float32)
	Modifies(attr Attribute) bool
}

type ModifyAge struct {
	Age Attribute
}

func (m ModifyAge) Modify(particle []float32, format *Format, dt float32) {
	format.Get(m.Age, particle)[0] += dt
}
func (m ModifyAge) Modifies(attr Attribute) bool {
	return m.Age.id == attr.id
}

type ModifyAdder struct {
	Value Attribute
	Add   Attribute
}

func (m ModifyAdder) Modify(particle []float32, format *Format, dt float32) {
	Add(format.Get(m.Value, particle), format.Get(m.Add, particle))
}
func (m ModifyAdder) Modifies(attr Attribute) bool {
	return m.Value.id == attr.id
}

type Modifys []Modify

func (m Modifys) Add(mod Modify) Modifys {
	return append(m, mod)
}

func (m Modifys) Age(attr Attribute) Modifys {
	return append(m, ModifyAge{Age: attr})
}

func (m Modifys) Adder(attr Attribute, add Attribute) Modifys {
	return append(m, ModifyAdder{Value: attr, Add: add})
}
