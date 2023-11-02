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

func (l *Layer) Update(update Update) Dirty {
	dirty := DirtyNone
	if l.Visual != nil {
		dirty.Add(l.Visual.Update(update))
	}
	if l.Background != nil {
		dirty.Add(l.Background.Update(update))
	}
	return dirty
}

func (l Layer) Render(ctx AmountContext, out *VertexBuffer) {
	layerCtx := ctx.WithParent(l.Bounds.Width(), l.Bounds.Height())
	span := out.DataSpan()
	l.Visual.Visualize(l.Bounds, layerCtx, out)
	if l.Background != nil {
		for i := 0; i < span.Len(); i++ {
			l.Background.Backgroundify(l.Bounds, layerCtx, span.At(i))
		}
	}
}

func (l Layer) ForStates(s State) bool {
	return l.States == nil || l.States(s)
}
