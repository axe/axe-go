package color

type Themeable interface {
	GetColorable(t Themed) Able
}

type Able interface {
	GetColor(b Themeable) Color
}

func Get(colorable Able, b Themeable) (Color, bool) {
	if colorable == nil {
		return Color{}, false
	} else {
		return colorable.GetColor(b), true
	}
}

type Themed uint8

var _ Able = Themed(0)

func (t Themed) GetColor(b Themeable) Color {
	if b == nil {
		return White
	}
	colorable := b.GetColorable(t)
	if colorable != nil {
		return colorable.GetColor(b)
	}
	return White
}

type modifiedThemeColor struct {
	modify     Modify
	themeColor Themed
}

var _ Able = modifiedThemeColor{}

func (m modifiedThemeColor) GetColor(b Themeable) Color {
	return m.modify(m.themeColor.GetColor(b))
}

func (t Themed) Modify(modify Modify) Able {
	return modifiedThemeColor{modify: modify, themeColor: t}
}
