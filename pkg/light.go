package axe

type LightAttenuate int

const (
	LightAttenuateNone LightAttenuate = iota
	LightAttenuateConstant
	LightAttenuateLinear
	LightAttenuateQuadratic
)

type Light struct {
	Ambient       Colorf
	Diffuse       Colorf
	Specular      Colorf
	Position      Vec4f
	SpotDirection Vec4f
	SpotExponent  uint8
	SpotCutOff    uint8
	Attenuate     LightAttenuate
}
