package ui

import (
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ease"
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/util"
)

type Animation interface {
	Init(base *Base)
	Update(base *Base, animationTime float32, update Update) Dirty
	IsDone(base *Base, animationTime float32) bool
	PostProcess(base *Base, animationTime float32, ctx *RenderContext, out *VertexBuffers)
}

type AnimationFactory interface {
	GetAnimation(base *Base) Animation
}

type AnimationEvent uint

const (
	AnimationEventNone AnimationEvent = iota
	AnimationEventShow
	AnimationEventHide
	AnimationEventRemove
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

func (as *AnimationState) PostProcess(base *Base, ctx *RenderContext, out *VertexBuffers) {
	if as.Current != nil {
		as.Current.PostProcess(base, as.CurrentTime, ctx, out)
		if as.lastPostProcess {
			as.Stop(true)
		}
	}
}

func (as *AnimationState) IsAnimating() bool {
	return as.Current != nil
}

func (b *Base) Play(factory AnimationFactory) bool {
	return b.playFactory(factory, AnimationEventNone)
}

func (b *Base) PlayMaybe(name string) bool {
	return b.PlayName(id.Maybe(name))
}

func (b *Base) PlayName(name id.Identifier) bool {
	return b.playFactory(b.AnimationNamed(name), AnimationEventNone)
}

func (b *Base) PlayEvent(ev AnimationEvent) bool {
	if b.Animation.CurrentEvent == ev {
		return true
	}
	return b.playFactory(b.AnimationFor(ev), ev)
}

func (b *Base) AnimationFor(ev AnimationEvent) AnimationFactory {
	var factory AnimationFactory
	if b.Animations != nil {
		factory = b.Animations.ForEvent.Get(ev)
	}
	if factory == nil {
		factory = b.ui.Theme.Animations.ForEvent.Get(ev)
	}
	return factory
}

func (b *Base) AnimationNamed(name id.Identifier) AnimationFactory {
	var named AnimationFactory
	if b.Animations != nil {
		named = b.Animations.Named.Get(name)
	}
	if named == nil {
		named = b.ui.Theme.Animations.Named.Get(name)
	}
	return named
}

func (b *Base) playFactory(factory AnimationFactory, ev AnimationEvent) bool {
	if factory == nil {
		return false
	}
	animation := factory.GetAnimation(b)
	if animation != nil {
		b.Animation.Set(animation, ev)
		b.Dirty(DirtyPostProcess)
		return true
	}
	return false
}

type BasicAnimationFrame struct {
	Translate    AmountPoint
	Scale        *Coord
	Rotate       float32
	Origin       AmountPoint
	Time         float32
	Color        ColorModify
	Transparency float32
	Easing       ease.Easing
}

func (start BasicAnimationFrame) Lerp(end BasicAnimationFrame, delta float32, ctx *AmountContext, x, y float32) BasicAnimationFrameInterpolated {
	startTx, startTy := start.Translate.Get(ctx)
	startOx, startOy := start.Origin.Get(ctx)
	endTx, endTy := end.Translate.Get(ctx)
	endOx, endOy := end.Origin.Get(ctx)
	startSc := getScale(start.Scale)
	endSc := getScale(end.Scale)

	inter := BasicAnimationFrameInterpolated{}
	inter.ScaleX = util.Lerp(startSc.X, endSc.X, delta)
	inter.ScaleY = util.Lerp(startSc.Y, endSc.Y, delta)
	inter.OriginX = util.Lerp(startOx, endOx, delta) + x
	inter.OriginY = util.Lerp(startOy, endOy, delta) + y
	inter.TranslateX = util.Lerp(startTx, endTx, delta)
	inter.TranslateY = util.Lerp(startTy, endTy, delta)
	inter.Rotation = util.Lerp(start.Rotate, end.Rotate, delta)
	inter.Color = start.Color.Lerp(end.Color, delta)
	inter.Transparency = util.Lerp(start.Transparency, end.Transparency, delta)

	return inter
}

func getScale(c *Coord) Coord {
	if c == nil {
		return Coord{X: 1, Y: 1}
	}
	return *c
}

type BasicAnimationFrameInterpolated struct {
	ScaleX, ScaleY         float32
	OriginX, OriginY       float32
	TranslateX, TranslateY float32
	Rotation               float32
	Color                  ColorModify
	Transparency           float32
}

func (inter BasicAnimationFrameInterpolated) Transform() Transform {
	transform := Transform{}
	transform.SetRotateDegreesScaleAround(inter.Rotation, inter.ScaleX, inter.ScaleY, inter.OriginX, inter.OriginY)
	transform.PreTranslate(inter.TranslateX, inter.TranslateY)
	return transform
}

type BasicAnimation struct {
	Duration            float32
	Easing              ease.Easing
	Save                bool
	SaveSkipColor       bool
	SaveSkipTransparent bool
	SaveSkipTransform   bool
	Frames              []BasicAnimationFrame
}

func (a BasicAnimation) GetAnimation(b *Base) Animation {
	return a
}

func (a BasicAnimation) WithDuration(duration float32) BasicAnimation {
	copy := a
	copy.Duration = duration
	return copy
}
func (a BasicAnimation) WithEasing(easing ease.Easing) BasicAnimation {
	copy := a
	copy.Easing = easing
	return copy
}
func (a BasicAnimation) WithSave(save, skipColor, skipTransparent, skipTransform bool) BasicAnimation {
	copy := a
	copy.Save = save
	copy.SaveSkipColor = skipColor
	copy.SaveSkipTransparent = skipTransparent
	copy.SaveSkipTransform = skipTransform
	return copy
}
func (a BasicAnimation) Init(base *Base) {}
func (a BasicAnimation) Update(base *Base, animationTime float32, update Update) Dirty {
	return DirtyPostProcess
}
func (a BasicAnimation) IsDone(base *Base, animationTime float32) bool {
	return animationTime > a.Duration
}
func (a BasicAnimation) PostProcess(base *Base, animationTime float32, ctx *RenderContext, out *VertexBuffers) {
	animationDelta := util.Min(animationTime/a.Duration, 1)
	animationEasingDelta := ease.Get(animationDelta, a.Easing)

	i := len(a.Frames) - 2
	for i > 0 && a.Frames[i].Time > animationEasingDelta {
		i--
	}

	start := a.Frames[i]
	end := a.Frames[i+1]

	timeDelta := util.Delta(start.Time, end.Time, animationEasingDelta)
	timeEasingDelta := ease.Get(timeDelta, start.Easing)

	inter := start.Lerp(end, timeEasingDelta, ctx.AmountContext, base.Bounds.Left, base.Bounds.Top)
	transform := inter.Transform()

	if a.Save {
		if !a.SaveSkipColor {
			base.SetColor(inter.Color)
		}
		if !a.SaveSkipColor {
			base.SetTransparency(inter.Transparency)
		}
		if !a.SaveSkipTransform {
			base.SetTransform(transform)
		}
	} else {
		alphaMultiplier := 1 - inter.Transparency
		colorModify := inter.Color
		if colorModify == nil {
			colorModify = func(c Color) Color { return c }
		}

		vertices := NewVertexIterator(out, true)
		for vertices.HasNext() {
			v := vertices.Next()
			v.X, v.Y = transform.Transform(v.X, v.Y)
			if !v.HasColor {
				v.Color = ColorWhite
				v.HasColor = true
			}
			v.Color = colorModify(v.Color)
			v.Color.A *= alphaMultiplier
		}
	}
}
