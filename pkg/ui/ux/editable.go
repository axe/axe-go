package ux

type EditableStyles struct {
}

type Editable struct {
	Text      string
	TextValue Value[string]

	OnChange func(string)
}
