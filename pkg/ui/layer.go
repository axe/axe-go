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

func (l Layer) Render(ctx *RenderContext, out *VertexBuffers) {
	layerCtx := ctx.WithBounds(l.Bounds)

	iter := NewVertexIterator(out)

	l.Visual.Visualize(l.Bounds, layerCtx, out)

	if l.Background != nil {
		for iter.HasNext() {
			l.Background.Backgroundify(l.Bounds, layerCtx, iter.Next())
		}
	}
}

func (l Layer) ForStates(s State) bool {
	return l.States == nil || l.States(s)
}
