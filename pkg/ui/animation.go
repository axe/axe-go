package ui

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/id"
)

type Animation interface {
	Init(base *Base)
	Update(base *Base, animationTime float32, update Update) Dirty
	IsDone(base *Base, animationTime float32) bool
	PostProcess(base *Base, animationTime float32, ctx *RenderContext, out *VertexBuffers, index IndexIterator, vertex VertexIterator)
}

type AnimationFactory func(base *Base) Animation

type AnimationEvent uint

const (
	AnimationEventNone AnimationEvent = iota
	AnimationEventShow
	AnimationEventHide
	AnimationEventFocus
	AnimationEventBlur
	AnimationEventPointerEnter
	AnimationEventPointerLeave
	AnimationEventPointerDown
	AnimationEventPointerUp
	AnimationEventDragStart
	AnimationEventDragStop
	AnimationEventDragCancel
	AnimationEventDragEnter
	AnimationEventDrop
	AnimationEventDisabled
	AnimationEventEnabled
)

type Animations struct {
	ForEvent ds.EnumMap[AnimationEvent, AnimationFactory]
	Named    id.DenseKeyMap[AnimationFactory, uint16, uint8]
}

func animationFactoryNil(af AnimationFactory) bool {
	return af == nil
}

func (a *Animations) Merge(other *Animations, replace bool) {
	a.ForEvent.Merge(other.ForEvent, replace, animationFactoryNil)
	a.Named.Merge(other.Named, replace, animationFactoryNil)
}

type AnimationState struct {
	Current      Animation
	CurrentTime  float32
	CurrentEvent AnimationEvent

	lastPostProcess bool
}

func (as *AnimationState) Set(a Animation, ev AnimationEvent) {
	as.Current = a
	as.CurrentTime = 0
	as.CurrentEvent = ev
	as.lastPostProcess = false
}

func (as *AnimationState) Stop(now bool) {
	if now {
		as.Current = nil
		as.CurrentTime = 0
		as.CurrentEvent = AnimationEventNone
		as.lastPostProcess = false
	} else {
		as.lastPostProcess = true
	}
}

func (as *AnimationState) Update(base *Base, update Update) Dirty {
	dirty := DirtyNone
	if as.Current != nil {
		if as.CurrentTime == 0 {
			as.Current.Init(base)
		}
		dirty = as.Current.Update(base, as.CurrentTime, update)
		if as.Current.IsDone(base, as.CurrentTime) {
			as.Stop(false)
		} else {
			as.CurrentTime += float32(update.DeltaTime.Seconds())
		}
	}
	return dirty
}

func (as *AnimationState) PostProcess(base *Base, ctx *RenderContext, out *VertexBuffers, index IndexIterator, vertex VertexIterator) {
	if as.Current != nil {
		as.Current.PostProcess(base, as.CurrentTime, ctx, out, index, vertex)
		if as.lastPostProcess {
			as.Stop(true)
		}
	}
}

func (as *AnimationState) IsAnimating() bool {
	return as.Current != nil
}

func (c *Base) PlayMaybe(name string) bool {
	return c.Play(id.Maybe(name))
}

func (c *Base) Play(name id.Identifier) bool {
	var named AnimationFactory
	if c.Animations != nil {
		named = c.Animations.Named.Get(name)
	}
	if named == nil {
		named = c.ui.Theme.Animations.Named.Get(name)
	}
	return c.playFactory(named, AnimationEventNone)
}

func (c *Base) PlayEvent(ev AnimationEvent) bool {
	if c.Animation.CurrentEvent == ev {
		return true
	}
	var factory AnimationFactory
	if c.Animations != nil {
		factory = c.Animations.ForEvent.Get(ev)
	}
	if factory == nil {
		factory = c.ui.Theme.Animations.ForEvent.Get(ev)
	}
	return c.playFactory(factory, ev)
}

func (c *Base) playFactory(factory AnimationFactory, ev AnimationEvent) bool {
	if factory == nil {
		return false
	}
	animation := factory(c)
	if animation != nil {
		c.Animation.Set(animation, ev)
		return true
	}
	return false
}

func StatelessAnimationFactory(a Animation) AnimationFactory {
	return func(base *Base) Animation {
		return a
	}
}

type BasicAnimationFrame struct {
	Translate    AmountPoint
	Scale        *Coord
	Rotate       float32
	Origin       AmountPoint
	Time         float32
	Transparency float32
	Easing       func(float32) float32
}

type BasicAnimation struct {
	Duration float32
	Easing   func(float32) float32
	Save     bool
	Frames   []BasicAnimationFrame
}

func (a BasicAnimation) WithDuration(duration float32) BasicAnimation {
	copy := a
	copy.Duration = duration
	return copy
}
func (a BasicAnimation) WithEasing(easing func(float32) float32) BasicAnimation {
	copy := a
	copy.Easing = easing
	return copy
}
func (a BasicAnimation) Init(base *Base) {}
func (a BasicAnimation) Update(base *Base, animationTime float32, update Update) Dirty {
	return DirtyVisual
}
func (a BasicAnimation) IsDone(base *Base, animationTime float32) bool {
	return animationTime > a.Duration
}
func (a BasicAnimation) PostProcess(base *Base, animationTime float32, ctx *RenderContext, out *VertexBuffers, index IndexIterator, vertex VertexIterator) {
	animationDelta := min(animationTime/a.Duration, 1)
	animationEasingDelta := Ease(animationDelta, a.Easing)

	i := len(a.Frames) - 2
	for i > 0 && a.Frames[i].Time > animationEasingDelta {
		i--
	}

	start := a.Frames[i]
	end := a.Frames[i+1]

	timeDelta := Delta(start.Time, end.Time, animationEasingDelta)
	timeEasingDelta := Ease(timeDelta, start.Easing)

	startTx, startTy := start.Translate.Get(ctx.AmountContext)
	startOx, startOy := start.Origin.Get(ctx.AmountContext)
	endTx, endTy := end.Translate.Get(ctx.AmountContext)
	endOx, endOy := end.Origin.Get(ctx.AmountContext)

	scaleX := float32(1)
	scaleY := float32(1)
	if start.Scale != nil && end.Scale != nil {
		scaleX = Lerp(start.Scale.X, end.Scale.X, timeEasingDelta)
		scaleY = Lerp(start.Scale.Y, end.Scale.Y, timeEasingDelta)
	}
	origX := Lerp(startOx, endOx, timeEasingDelta) + base.Bounds.Left
	origY := Lerp(startOy, endOy, timeEasingDelta) + base.Bounds.Top
	transX := Lerp(startTx, endTx, timeEasingDelta)
	transY := Lerp(startTy, endTy, timeEasingDelta)

	rotation := Lerp(start.Rotate, end.Rotate, timeEasingDelta)
	transparency := Lerp(start.Transparency, end.Transparency, timeEasingDelta)

	transform := Transform{}
	transform.SetRotateDegreesScaleAround(rotation, scaleX, scaleY, origX, origY)
	transform.PreTranslate(transX, transY)

	if a.Save {
		base.Transparency.Set(transparency)
		if transform.IsEffectivelyIdentity() {
			base.Transform.Identity()
		} else {
			base.Transform = transform
		}
	} else {
		alphaScalar := 1 - transparency

		for vertex.HasNext() {
			v := vertex.Next()
			v.X, v.Y = transform.Transform(v.X, v.Y)
			v.Color.A *= alphaScalar
		}
	}
}
