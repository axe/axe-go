package fx

import (
	"github.com/axe/axe-go/pkg/ease"
	"github.com/axe/axe-go/pkg/util"
)

type Access interface {
	Get(particle []float32, format *Format) []float32
}

var _ Access = AccessData{}
var _ Access = AccessConstant{}
var _ Access = AccessDynamic{}
var _ Access = AccessLerp{}

type AccessData struct {
	Offset int
	Size   int
}

func (a AccessData) Get(particle []float32, format *Format) []float32 {
	return particle[a.Offset : a.Offset+a.Size]
}

type AccessConstant struct {
	Constant []float32
}

func (a AccessConstant) Get(particle []float32, format *Format) []float32 {
	return a.Constant
}

type Dynamic func(particle []float32, format *Format) []float32

type AccessDynamic struct {
	Dynamic Dynamic
}

func (a AccessDynamic) Get(particle []float32, format *Format) []float32 {
	return a.Dynamic(particle, format)
}

type AccessLerp struct {
	Data   [][]float32
	Easing ease.Easing

	temp []float32
}

func (a AccessLerp) Get(particle []float32, format *Format) []float32 {
	delta := Life(particle, format)
	deltaEased := ease.Get(delta, a.Easing)
	n := len(a.Data) - 1
	deltaScaled := deltaEased * float32(n)
	frameStart := util.Clamp(int(deltaScaled), 0, n-1)
	frameDelta := deltaScaled - float32(frameStart)
	Lerp(a.Data[frameStart], a.Data[frameStart+1], frameDelta, a.temp)
	return a.temp
}
