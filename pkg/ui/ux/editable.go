package ux

type EditableSettings struct {
}

type Editable struct {
	Text      string
	TextValue Value[string]

	OnChange func(string)
}
