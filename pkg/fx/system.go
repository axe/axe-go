package fx

type SystemFormat struct {
	Format       *Format
	Initializers Inits
	Modifiers    Modifys
}

func (sf SystemFormat) Inits(attr Attribute) bool {
	for _, init := range sf.Initializers {
		if init.Inits(attr) {
			return true
		}
	}
	return false
}

func (sf SystemFormat) Modifies(attr Attribute) bool {
	for _, modifier := range sf.Modifiers {
		if modifier.Modifies(attr) {
			return true
		}
	}
	return false
}

func (sf *SystemFormat) Setup() {
	for _, attr := range sf.Format.Attributes {
		if attr.Access == nil {
			continue
		}

		if attr.Attribute.init != nil && !sf.Inits(attr.Attribute) && sf.Format.HasData(attr.Attribute) {
			sf.Initializers = sf.Initializers.Add(attr.Attribute.init)
		}

		if attr.Attribute.modify != nil && !sf.Modifies(attr.Attribute) {
			sf.Modifiers = sf.Modifiers.Add(attr.Attribute.modify)
		}
	}
}

type System struct {
	Format *SystemFormat
	Data   Data
}
