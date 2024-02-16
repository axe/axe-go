package fx_test

import (
	"fmt"
	"testing"

	"github.com/axe/axe-go/pkg/color"
	"github.com/axe/axe-go/pkg/fx"
)

func TestFormat(t *testing.T) {
	pf := fx.NewFormat().
		AddData(fx.Age).
		AddData(fx.Lifespan).
		AddData(fx.Pos2).
		AddData(fx.Vel2).
		AddConstant(fx.Acc2, 0, -10).
		AddLerp(fx.Alpha, [][]float32{{1}, {0}}, nil).
		AddLerp(fx.Shade, [][]float32{color.Red.ToFloats(), color.Orange.ToFloats(), color.Yellow.ToFloats()}, nil)

	pd := fx.NewData(pf, 100)

	pd.Get(0, fx.Lifespan)[0] = 1
	pd.Get(0, fx.Age)[0] = 0.5
	alpha := pd.Get(0, fx.Alpha)
	shade := pd.Get(0, fx.Shade)

	fmt.Printf("alpha: %v, shade: %v\n", alpha, shade)

	sf := fx.SystemType{
		Format: pf,
		Initializers: fx.Inits{}.
			Random(fx.Lifespan, []float32{5}, []float32{10}).
			Constant(fx.Pos2, 10, 10),
	}
	sf.Setup()

	fmt.Printf("SystemFormat: %+v\n", sf)
}
