package ui

type Layer struct {
	Placement  Placement
	Bounds     Bounds
	Visual     Visual
	Background Background
	States     StateFn
}

func (l *Layer) Init(init Init) {
	l.Placement.Init(Maximized())
	if l.Visual != nil {
		l.Visual.Init(init)
	}
	if l.Background != nil {
		l.Background.Init(init)
	}
}

func (l *Layer) Place(parent Bounds) {
	l.Bounds = l.Placement.GetBoundsIn(parent)
}

func (l *Layer) Update(update Update) {
	if l.Visual != nil {
		l.Visual.Update(update)
	}
	if l.Background != nil {
		l.Background.Update(update)
	}
}

func (l Layer) Render(ctx AmountContext, out *UIVertexBuffer) {
	layerCtx := ctx.WithParent(l.Bounds.Width(), l.Bounds.Height())
	start := out.Pos()
	l.Visual.Visualize(l.Bounds, layerCtx, out)
	end := out.Pos()
	if l.Background != nil {
		for i := start; i < end; i++ {
			l.Background.Backgroundify(l.Bounds, layerCtx, &out.Data[i])
		}
	}
}

func (l Layer) ForStates(s State) bool {
	return l.States == nil || l.States(s)
}
