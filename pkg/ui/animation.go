package ui

import (
	"sort"

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

func (start BasicAnimationFrame) PreLerp(end BasicAnimationFrame, delta float32) BasicAnimationFrame {
	if delta == 0 {
		return start
	}
	if delta == 1 {
		return end
	}

	startScale := getScale(start.Scale)
	endScale := getScale(end.Scale)

	lerped := BasicAnimationFrame{
		Translate:    start.Translate.Lerp(end.Translate, delta),
		Scale:        Override(startScale.Lerp(endScale, delta)),
		Rotate:       util.Lerp(start.Rotate, end.Rotate, delta),
		Origin:       start.Origin.Lerp(end.Origin, delta),
		Time:         util.Lerp(start.Time, end.Time, delta),
		Transparency: util.Lerp(start.Time, end.Time, delta),
		Easing:       ease.NewSubset(start.Easing, delta, 1),
		Color:        start.Color.Lerp(end.Color, delta),
	}

	if start.Scale == nil && end.Scale == nil {
		lerped.Scale = nil
	}
	if start.Easing == nil {
		lerped.Easing = nil
	}

	return lerped
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

type BasicAnimationProp int

const (
	BasicAnimationPropScale BasicAnimationProp = 1 << iota
	BasicAnimationPropOrigin
	BasicAnimationPropTranslate
	BasicAnimationPropRotate
	BasicAnimationPropColor
	BasicAnimationPropTransparency
	BasicAnimationPropEasing
)

type BasicAnimation struct {
	Duration             float32
	Easing               ease.Easing
	Save                 bool
	SaveSkipColor        bool
	SaveSkipTransparency bool
	SaveSkipTransform    bool
	Frames               []BasicAnimationFrame
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
	copy.SaveSkipTransparency = skipTransparent
	copy.SaveSkipTransform = skipTransform
	return copy
}
func (a BasicAnimation) WithNoSave() BasicAnimation {
	copy := a
	copy.Save = false
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
	i := a.IndexForTime(animationEasingDelta)
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
		if !a.SaveSkipTransparency {
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
func (a BasicAnimation) IndexForTime(time float32) int {
	i := len(a.Frames) - 2
	for i > 0 && a.Frames[i].Time > time {
		i--
	}
	return i
}
func (a BasicAnimation) PreLerpForTime(time float32) BasicAnimationFrame {
	i := a.IndexForTime(time)
	start := a.Frames[i]
	end := a.Frames[i+1]
	delta := util.Delta(start.Time, end.Time, time)
	return start.PreLerp(end, delta)
}
func (a BasicAnimation) GetProps() BasicAnimationProp {
	var props BasicAnimationProp
	for _, frame := range a.Frames {
		if frame.Color != nil {
			props |= BasicAnimationPropColor
		}
		if frame.Easing != nil {
			props |= BasicAnimationPropEasing
		}
		if frame.Rotate != 0 {
			props |= BasicAnimationPropRotate
		}
		if !frame.Origin.IsZero() {
			props |= BasicAnimationPropOrigin
		}
		if frame.Scale != nil {
			props |= BasicAnimationPropScale
		}
		if !frame.Translate.IsZero() {
			props |= BasicAnimationPropTranslate
		}
		if frame.Transparency > 0 {
			props |= BasicAnimationPropTransparency
		}
	}
	return props
}
func (a BasicAnimation) Copy() BasicAnimation {
	copy := a
	copy.Frames = util.SliceCopy(a.Frames)
	for i := range copy.Frames {
		frame := &copy.Frames[i]
		if frame.Scale != nil {
			frame.Scale = util.Clone(frame.Scale)
		}
	}
	return copy
}
func (a BasicAnimation) Only(props BasicAnimationProp) BasicAnimation {
	copy := a.Copy()
	for i := range copy.Frames {
		frame := &copy.Frames[i]
		if props&BasicAnimationPropColor == 0 {
			frame.Color = nil
		}
		if props&BasicAnimationPropEasing == 0 {
			frame.Easing = nil
		}
		if props&BasicAnimationPropOrigin == 0 {
			frame.Origin = AmountPoint{}
		}
		if props&BasicAnimationPropRotate == 0 {
			frame.Rotate = 0
		}
		if props&BasicAnimationPropScale == 0 {
			frame.Scale = nil
		}
		if props&BasicAnimationPropTranslate == 0 {
			frame.Translate = AmountPoint{}
		}
		if props&BasicAnimationPropTransparency == 0 {
			frame.Transparency = 0
		}
	}
	return copy
}
func (a BasicAnimation) Without(props BasicAnimationProp) BasicAnimation {
	return a.Only(^props)
}
func (a BasicAnimation) Merge(b BasicAnimation) BasicAnimation {
	times := map[float32]struct{}{}
	for _, frame := range a.Frames {
		times[frame.Time] = struct{}{}
	}
	for _, frame := range b.Frames {
		times[frame.Time] = struct{}{}
	}
	timesOrdered := make([]float32, 0, len(times))
	for time := range times {
		timesOrdered = append(timesOrdered, time)
	}
	sort.Slice(timesOrdered, func(i, j int) bool {
		return timesOrdered[i] < timesOrdered[j]
	})

	bProps := b.GetProps()

	frames := make([]BasicAnimationFrame, len(timesOrdered))
	for index, time := range timesOrdered {
		aFrame := a.PreLerpForTime(time)
		bFrame := b.PreLerpForTime(time)

		frames[index] = BasicAnimationFrame{
			Time:         time,
			Color:        util.If(bProps&BasicAnimationPropColor == 0, aFrame.Color, bFrame.Color),
			Easing:       util.If(bProps&BasicAnimationPropEasing == 0, aFrame.Easing, bFrame.Easing),
			Origin:       util.If(bProps&BasicAnimationPropOrigin == 0, aFrame.Origin, bFrame.Origin),
			Rotate:       util.If(bProps&BasicAnimationPropRotate == 0, aFrame.Rotate, bFrame.Rotate),
			Scale:        util.If(bProps&BasicAnimationPropScale == 0, aFrame.Scale, bFrame.Scale),
			Translate:    util.If(bProps&BasicAnimationPropTranslate == 0, aFrame.Translate, bFrame.Translate),
			Transparency: util.If(bProps&BasicAnimationPropTransparency == 0, aFrame.Transparency, bFrame.Transparency),
		}
	}

	return BasicAnimation{
		Duration:             util.Max(a.Duration, b.Duration),
		Easing:               a.Easing,
		Save:                 a.Save || b.Save,
		SaveSkipColor:        a.SaveSkipColor && b.SaveSkipColor,
		SaveSkipTransparency: a.SaveSkipTransparency && b.SaveSkipTransparency,
		SaveSkipTransform:    a.SaveSkipTransform && b.SaveSkipTransform,
		Frames:               frames,
	}
}

func (a BasicAnimation) Reverse() BasicAnimation {
	copy := a.Copy()
	for i := range copy.Frames {
		frame := &copy.Frames[i]
		frame.Time = 1 - frame.Time
	}
	util.SliceReverse(copy.Frames)
	return copy
}
