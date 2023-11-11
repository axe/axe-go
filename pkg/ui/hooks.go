package ui

import "github.com/axe/axe-go/pkg/util"

type Hooks struct {
	OnInit                HookInit
	OnPlace               HookPlace
	OnUpdate              HookUpdate
	OnRender              HookRender
	OnPostProcess         PostProcess
	OnPostProcessLayers   PostProcess
	OnPostProcessChildren PostProcess
}

func (h *Hooks) Add(add Hooks, before bool) {
	h.OnInit.Add(add.OnInit, before)
	h.OnPlace.Add(add.OnPlace, before)
	h.OnUpdate.Add(add.OnUpdate, before)
	h.OnRender.Add(add.OnRender, before)
	h.OnPostProcess.Add(add.OnPostProcess, before)
	h.OnPostProcessLayers.Add(add.OnPostProcessLayers, before)
	h.OnPostProcessChildren.Add(add.OnPostProcessChildren, before)
}

type HookInit func(b *Base, init Init)

func hookInitNil(h HookInit) bool {
	return h == nil
}
func hookInitJoin(first, second HookInit) HookInit {
	return func(b *Base, init Init) {
		first(b, init)
		second(b, init)
	}
}

func (h *HookInit) Add(add HookInit, before bool) {
	*h = util.CoalesceJoin(*h, add, before, hookInitNil, hookInitJoin)
}

type HookPlace func(b *Base, parent Bounds, ctx *RenderContext, force bool)

func hookPlaceNil(h HookPlace) bool {
	return h == nil
}
func hookPlaceJoin(first, second HookPlace) HookPlace {
	return func(b *Base, parent Bounds, ctx *RenderContext, force bool) {
		first(b, parent, ctx, force)
		second(b, parent, ctx, force)
	}
}

func (h *HookPlace) Add(add HookPlace, before bool) {
	*h = util.CoalesceJoin(*h, add, before, hookPlaceNil, hookPlaceJoin)
}

type HookUpdate func(b *Base, update Update) Dirty

func hookUpdateNil(h HookUpdate) bool {
	return h == nil
}
func hookUpdateJoin(first, second HookUpdate) HookUpdate {
	return func(b *Base, update Update) Dirty {
		dirty := DirtyNone
		dirty.Add(first(b, update))
		dirty.Add(second(b, update))
		return dirty
	}
}

func (h *HookUpdate) Add(add HookUpdate, before bool) {
	*h = util.CoalesceJoin(*h, add, before, hookUpdateNil, hookUpdateJoin)
}

type HookRender func(b *Base, ctx *RenderContext, out *VertexBuffers)

func hookRenderNil(h HookRender) bool {
	return h == nil
}
func hookRenderJoin(first, second HookRender) HookRender {
	return func(b *Base, ctx *RenderContext, out *VertexBuffers) {
		first(b, ctx, out)
		second(b, ctx, out)
	}
}

func (h *HookRender) Add(add HookRender, before bool) {
	*h = util.CoalesceJoin(*h, add, before, hookRenderNil, hookRenderJoin)
}
