package fx

var (
	// Age in seconds
	Age = NewAttribute(1)
	// Total lifespan in seconds
	Lifespan = NewAttribute(1)

	// A value that can be randomly generated and used in initializers and modifiers.
	// Useful for generating particles on a shape with a velocity relative to the shape.
	Seed = NewAttribute(1)

	// A 2d position
	Pos2 = NewAttribute(2)
	// A 2d velocity
	Vel2 = NewAttribute(2)
	// A 2d acceleration
	Acc2 = NewAttribute(2)

	// A 3d position
	Pos3 = NewAttribute(3)
	// A 3d velocity
	Vel3 = NewAttribute(3)
	// A 3d acceleration
	Acc3 = NewAttribute(3)

	// A uniform size (same width & height)
	Size = NewAttribute(1)
	// A uniform size velocity
	SizeVel = NewAttribute(1)
	// A uniform size acceleration
	SizeAcc = NewAttribute(1)

	// A width & height
	Size2 = NewAttribute(2)
	// A width & height velocity
	SizeVel2 = NewAttribute(2)
	// A width & height acceleration
	SizeAcc2 = NewAttribute(2)

	// A uniform scale (both width & height)
	Scale = NewAttribute(1)
	// A uniform scale velocity
	ScaleVel = NewAttribute(1)
	// A uniform scale acceleration
	ScaleAcc = NewAttribute(1)

	// A 2d scale
	Scale2 = NewAttribute(2)
	// A 2d scale velocity
	ScaleVel2 = NewAttribute(2)
	// A 2d scale acceleration
	ScaleAcc2 = NewAttribute(2)

	// An angle/rotation in radians
	Angle = NewAttribute(1)
	// Angular velocity
	AngleVel = NewAttribute(1)
	// Angular acceleration
	AngleAcc = NewAttribute(1)

	// Rotation anchor, {0,0}=top left, {1,1}=bottom right, {.5,.5}=center
	Anchor = NewAttribute(2)

	// An RGB color
	Shade = NewAttribute(3)
	// An RGBA color
	Color = NewAttribute(4)
	// The alpha
	Alpha = NewAttribute(1)

	// Which texture defined in the system format to use.
	Texture = NewAttribute(1)
)

func init() {
	// inits are automatically added if the format has the attribute as data and an init for the attribute has not been set
	// modifys are automatically added if one has not been set

	Age.init = InitConstant{Attribute: Age, Constant: []float32{0}}
	Age.modify = ModifyAge{Age: Age}
	Seed.init = InitRandom{Attribute: Seed, Start: []float32{0}, End: []float32{1}}
	Vel2.init = InitConstant{Attribute: Vel2, Constant: []float32{0, 0}}
	Vel2.modify = ModifyAdder{Value: Pos2, Add: Vel2}
	Acc2.init = InitConstant{Attribute: Acc2, Constant: []float32{0, 0}}
	Acc2.modify = ModifyAdder{Value: Vel2, Add: Acc2}
	Vel3.init = InitConstant{Attribute: Vel3, Constant: []float32{0, 0, 0}}
	Vel3.modify = ModifyAdder{Value: Pos3, Add: Vel3}
	Acc3.init = InitConstant{Attribute: Acc3, Constant: []float32{0, 0, 0}}
	Acc3.modify = ModifyAdder{Value: Vel3, Add: Acc3}
	SizeVel.init = InitConstant{Attribute: SizeVel, Constant: []float32{0}}
	SizeVel.modify = ModifyAdder{Value: Size, Add: SizeVel}
	SizeAcc.init = InitConstant{Attribute: SizeAcc, Constant: []float32{0}}
	SizeAcc.modify = ModifyAdder{Value: SizeVel, Add: SizeAcc}
	SizeVel2.init = InitConstant{Attribute: SizeVel2, Constant: []float32{0, 0}}
	SizeVel2.modify = ModifyAdder{Value: Size2, Add: SizeVel2}
	SizeAcc2.init = InitConstant{Attribute: SizeAcc2, Constant: []float32{0, 0}}
	SizeAcc2.modify = ModifyAdder{Value: SizeVel2, Add: SizeAcc2}
	Scale.init = InitConstant{Attribute: Scale, Constant: []float32{1}}
	ScaleVel.init = InitConstant{Attribute: ScaleVel, Constant: []float32{0}}
	ScaleVel.modify = ModifyAdder{Value: Scale, Add: ScaleVel}
	ScaleAcc.init = InitConstant{Attribute: ScaleAcc, Constant: []float32{0}}
	ScaleAcc.modify = ModifyAdder{Value: ScaleVel, Add: ScaleAcc}
	Scale2.init = InitConstant{Attribute: Scale2, Constant: []float32{1, 1}}
	ScaleVel2.init = InitConstant{Attribute: ScaleAcc, Constant: []float32{0, 0}}
	ScaleVel2.modify = ModifyAdder{Value: Scale2, Add: ScaleVel2}
	ScaleAcc2.init = InitConstant{Attribute: ScaleAcc2, Constant: []float32{0, 0}}
	ScaleAcc2.modify = ModifyAdder{Value: ScaleVel2, Add: ScaleAcc2}
	Anchor.init = InitConstant{Attribute: Anchor, Constant: []float32{0.5, 0.5}}
	Shade.init = InitConstant{Attribute: Shade, Constant: []float32{1, 1, 1}}
	Color.init = InitConstant{Attribute: Color, Constant: []float32{1, 1, 1, 1}}
	Alpha.init = InitConstant{Attribute: Alpha, Constant: []float32{1}}
}
