package axe

import "github.com/axe/axe-go/pkg/color"

type LightAttenuate int

const (
	LightAttenuateNone LightAttenuate = iota
	LightAttenuateConstant
	LightAttenuateLinear
	LightAttenuateQuadratic
)

type Light struct {
	Ambient       color.Color
	Diffuse       color.Color
	Specular      color.Color
	Position      Vec4f
	SpotDirection Vec4f
	SpotExponent  uint8
	SpotCutOff    uint8
	Attenuate     LightAttenuate
}
