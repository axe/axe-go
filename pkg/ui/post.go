package ui

import "github.com/axe/axe-go/pkg/util"

type PostProcess func(b *Base, ctx *RenderContext, out *VertexBuffers)

func PostProcessVertex(modify VertexModifier) PostProcess {
	return func(b *Base, ctx *RenderContext, out *VertexBuffers) {
		vertices := NewVertexIterator(out, true)
		for vertices.HasNext() {
			modify(vertices.Next())
		}
	}
}

func (pp PostProcess) Run(b *Base, ctx *RenderContext, out *VertexBuffers) {
	if pp != nil {
		pp(b, ctx, out)
	}
}

func postProcessNil(a PostProcess) bool {
	return a == nil
}

func postProcessJoin(first, second PostProcess) PostProcess {
	return func(b *Base, ctx *RenderContext, out *VertexBuffers) {
		first(b, ctx, out)
		second(b, ctx, out)
	}
}

func (pp *PostProcess) Add(other PostProcess, before bool) {
	*pp = util.CoalesceJoin(*pp, other, before, postProcessNil, postProcessJoin)
}

func (pp *PostProcess) Clear() {
	*pp = nil
}
