package ui

type Layer struct {
	Placement  Placement
	Bounds     Bounds
	Visual     Visual
	Background Background
	States     StateFn
}

func (l *Layer) Init(b *Base) {
	l.Placement.Init(Maximized())
	l.Visual.Init(b)
	if l.Background != nil {
		l.Background.Init(b)
	}
}

func (l *Layer) Place(b *Base, parent Bounds) {
	l.Bounds = l.Placement.GetBoundsIn(parent)
}

func (l *Layer) Update(b *Base, update Update) Dirty {
	dirty := DirtyNone
	dirty.Add(l.Visual.Update(b, update))
	if l.Background != nil {
		dirty.Add(l.Background.Update(b, update))
	}
	return dirty
}

func (l Layer) PreferredSize(b *Base, ctx *RenderContext, maxWidth float32) Coord {
	layerCtx := ctx.WithBounds(l.Bounds)
	padding := l.Placement.Padding()
	size := l.Visual.PreferredSize(b, layerCtx, maxWidth-padding.X)
	size.X += padding.X
	size.Y += padding.Y

	// if text, ok := l.Visual.(*VisualText); ok {
	// 	if text.Paragraphs.String() == "Toggle Vertical Spacing" {
	// 		fmt.Printf("PreferredSize Bounds %v, Padding: %v, maxWidth: %v, size: %v\n", l.Bounds, padding, maxWidth, size)
	// 	}
	// }

	return size
}

func (l Layer) Render(b *Base, ctx *RenderContext, out *VertexBuffers) {
	layerCtx := ctx.WithBounds(l.Bounds)
	iter := NewVertexIterator(out)

	// if text, ok := l.Visual.(*VisualText); ok {
	// 	if text.Paragraphs.String() == "Toggle Vertical Spacing" {
	// 		fmt.Printf("Render Bounds %v, Padding: %v\n", l.Bounds, l.Placement.Padding())
	// 	}
	// }

	l.Visual.Visualize(b, l.Bounds, layerCtx, out)
	if l.Background != nil {
		for iter.HasNext() {
			l.Background.Backgroundify(b, l.Bounds, layerCtx, iter.Next())
		}
	}
}

func (l Layer) ComputeRenderContext(b *Base) *RenderContext {
	return b.ComputeRenderContext().WithBounds(l.Bounds)
}

func (l Layer) ForStates(s State) bool {
	return l.States == nil || l.States(s)
}
