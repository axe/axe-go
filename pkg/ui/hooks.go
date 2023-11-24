package ui

import "github.com/axe/axe-go/pkg/util"

type Hooks struct {
	OnInit                HookBase
	OnPlace               HookPlace
	OnUpdate              HookUpdate
	OnRender              HookRender
	OnRemove              HookBase
	OnPostProcess         PostProcess
	OnPostProcessLayers   PostProcess
	OnPostProcessChildren PostProcess
}

func (h *Hooks) Add(add Hooks, before bool) {
	h.OnInit.Add(add.OnInit, before)
	h.OnPlace.Add(add.OnPlace, before)
	h.OnUpdate.Add(add.OnUpdate, before)
	h.OnRender.Add(add.OnRender, before)
	h.OnRemove.Add(add.OnRemove, before)
	h.OnPostProcess.Add(add.OnPostProcess, before)
	h.OnPostProcessLayers.Add(add.OnPostProcessLayers, before)
	h.OnPostProcessChildren.Add(add.OnPostProcessChildren, before)
}

func (h *Hooks) Clear() {
	h.OnInit.Clear()
	h.OnPlace.Clear()
	h.OnUpdate.Clear()
	h.OnRender.Clear()
	h.OnRemove.Clear()
	h.OnPostProcess.Clear()
	h.OnPostProcessLayers.Clear()
	h.OnPostProcessChildren.Clear()
}

type HookBase func(b *Base)

func hookBaseNil(h HookBase) bool {
	return h == nil
}
func hookBaseJoin(first, second HookBase) HookBase {
	return func(b *Base) {
		first(b)
		second(b)
	}
}

func (h HookBase) Run(b *Base) {
	if h != nil {
		h(b)
	}
}

func (h *HookBase) Add(add HookBase, before bool) {
	*h = util.CoalesceJoin(*h, add, before, hookBaseNil, hookBaseJoin)
}

func (h *HookBase) Clear() {
	*h = nil
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

func (h HookPlace) Run(b *Base, parent Bounds, ctx *RenderContext, force bool) {
	if h != nil {
		h(b, parent, ctx, force)
	}
}

func (h *HookPlace) Add(add HookPlace, before bool) {
	*h = util.CoalesceJoin(*h, add, before, hookPlaceNil, hookPlaceJoin)
}

func (h *HookPlace) Clear() {
	*h = nil
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

func (h HookUpdate) Run(b *Base, update Update) {
	if h != nil {
		h(b, update)
	}
}

func (h *HookUpdate) Add(add HookUpdate, before bool) {
	*h = util.CoalesceJoin(*h, add, before, hookUpdateNil, hookUpdateJoin)
}

func (h *HookUpdate) Clear() {
	*h = nil
}

type HookRender func(b *Base, ctx *RenderContext, out *VertexQueue)

func hookRenderNil(h HookRender) bool {
	return h == nil
}
func hookRenderJoin(first, second HookRender) HookRender {
	return func(b *Base, ctx *RenderContext, out *VertexQueue) {
		first(b, ctx, out)
		second(b, ctx, out)
	}
}

func (h *HookRender) Add(add HookRender, before bool) {
	*h = util.CoalesceJoin(*h, add, before, hookRenderNil, hookRenderJoin)
}

func (h *HookRender) Clear() {
	*h = nil
}
