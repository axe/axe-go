package ui

import "github.com/axe/axe-go/pkg/util"

type PostProcess func(b *Base, ctx *RenderContext, out *VertexBuffers, index IndexIterator, vertex VertexIterator)

func PostProcessVertex(modify VertexModifier) PostProcess {
	return func(b *Base, ctx *RenderContext, out *VertexBuffers, index IndexIterator, vertex VertexIterator) {
		for vertex.HasNext() {
			modify(vertex.Next())
		}
	}
}

func (pp PostProcess) Run(b *Base, ctx *RenderContext, out *VertexBuffers, index IndexIterator, vertex VertexIterator) {
	if pp != nil {
		pp(b, ctx, out, index, vertex)
	}
}

func postProcessNil(a PostProcess) bool {
	return a == nil
}

func postProcessJoin(first, second PostProcess) PostProcess {
	return func(b *Base, ctx *RenderContext, out *VertexBuffers, index IndexIterator, vertex VertexIterator) {
		first(b, ctx, out, index, vertex)
		second(b, ctx, out, index, vertex)
	}
}

func (pp *PostProcess) Add(other PostProcess, before bool) {
	*pp = util.CoalesceJoin(*pp, other, before, postProcessNil, postProcessJoin)
}
