package ui

type Layer struct {
	Placement  Placement
	Bounds     Bounds
	Visual     Visual
	Background Background
	States     StateFn
}

func (l *Layer) Init(b *Base, init Init) {
	l.Placement.Init(Maximized())
	if l.Visual != nil {
		l.Visual.Init(b, init)
	}
	if l.Background != nil {
		l.Background.Init(b, init)
	}
}

func (l *Layer) Place(b *Base, parent Bounds) {
	l.Bounds = l.Placement.GetBoundsIn(parent)
}

func (l *Layer) Update(b *Base, update Update) Dirty {
	dirty := DirtyNone
	if l.Visual != nil {
		dirty.Add(l.Visual.Update(b, update))
	}
	if l.Background != nil {
		dirty.Add(l.Background.Update(b, update))
	}
	return dirty
}

func (l Layer) Render(b *Base, ctx *RenderContext, out *VertexBuffers) {
	layerCtx := ctx.WithBounds(l.Bounds)

	iter := NewVertexIterator(out)

	l.Visual.Visualize(b, l.Bounds, layerCtx, out)

	if l.Background != nil {
		for iter.HasNext() {
			l.Background.Backgroundify(b, l.Bounds, layerCtx, iter.Next())
		}
	}
}

func (l Layer) ForStates(s State) bool {
	return l.States == nil || l.States(s)
}
