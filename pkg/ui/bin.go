package ui

type Value[V any] struct {
	Field    string
	Method   string
	Getter   string
	Setter   string
	Dynamic  bool
	Computed func() V
	Constant V
}

type Container struct {
	Padding Placement
}

type Button struct {
	Container
	Label   Text
	Pressed Value[bool]
}

type List struct {
	Container
	Items    Value[[]any]
	Selected Value[[]any]
}

type CheckboxState string

const (
	CheckboxStateChecked       = "checked"
	CheckboxStateUnchecked     = "unchecked"
	CheckboxStateIndeterminate = "indeterminate"
)

type Checkbox struct {
	Container
	Label   Text
	Checked Value[CheckboxState]
	Value   Value[any]
	Values  map[CheckboxState]any
}

type Text struct {
	Content Value[string]
}

type Input struct { // without visuals, just event handling and text
	Value      Value[string]
	Selectable bool
	Editable   bool
}

type Tooltip struct {
	Content Value[string]
}

type Image struct {
	// Image Quad
}

type Scrollable struct {
}
