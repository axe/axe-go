package ux

import (
	"github.com/axe/axe-go/pkg/input"
	"github.com/axe/axe-go/pkg/ui"
)

type ButtonStyles struct {
}

type Button struct {
	*ui.Template

	Text      string
	TextValue Value[string]
	LeftIcon  Icon
	RightIcon Icon

	OnClick   func()
	OnEnter   func()
	OnAction  func()
	OnPointer func(ev *ui.PointerEvent)
}

var _ HasComponent = Button{}

func (b Button) Build(theme *Theme) *ButtonBase {

	text := CoalesceValue(b.TextValue, b.Text)
	textVisual := &ui.VisualText{}

	var base *ButtonBase
	base = &ButtonBase{
		Text:       text,
		TextVisual: textVisual,
		Options:    b,
		Component: &ui.Base{
			Layers: []ui.Layer{{
				Visual: textVisual,
			}},
			Hooks: ui.Hooks{
				OnUpdate: func(b *ui.Base, update ui.Update) ui.Dirty {
					if changed := base.Text.GetChanged(); changed != nil {
						textVisual.SetText(*changed)
					}
					return ui.DirtyNone
				},
			},
			Events: ui.Events{
				OnPointer: func(ev *ui.PointerEvent) {
					if b.OnPointer != nil {
						b.OnPointer(ev)
						if ev.Stop {
							return
						}
					}
					if !ev.Capture && ev.Type == ui.PointerEventDown {
						if b.OnClick != nil {
							b.OnClick()
							ev.Stop = true
						} else if b.OnAction != nil {
							b.OnAction()
							ev.Stop = true
						}
						base.clickTrigger.Set(1)
						base.actionTrigger.Set(1)
						base.clicks.Add(1)
					}
					if !ev.Capture && ev.Type == ui.PointerEventUp {
						base.clickTrigger.Set(0)
						base.actionTrigger.Set(0)
					}
				},
				OnKey: func(ev *ui.KeyEvent) {
					if !ev.Capture && ev.Key == input.KeyEnter && ev.Type == ui.KeyEventDown {
						if b.OnEnter != nil {
							b.OnEnter()
							ev.Stop = true
						} else if b.OnAction != nil {
							b.OnAction()
							ev.Stop = true
						}
						base.enterTrigger.Set(1)
						base.actionTrigger.Set(1)
						base.enters.Add(1)
					}
					if !ev.Capture && ev.Key == input.KeyEnter && ev.Type == ui.KeyEventUp {
						base.enterTrigger.Set(0)
						base.actionTrigger.Set(0)
					}
				},
			},
		},
	}

	base.Component.ApplyTemplate(b.Template)

	return base
}

func (b Button) GetComponent(theme *Theme) *ui.Base {
	return b.Build(theme).GetComponent(theme)
}

type ButtonBase struct {
	Component  *ui.Base
	Text       Value[string]
	Options    Button
	TextVisual *ui.VisualText

	clickTrigger  Trigger
	enterTrigger  Trigger
	actionTrigger Trigger
	clicks        Counter
	enters        Counter
}

var _ HasComponent = &ButtonBase{}

func (b *ButtonBase) GetComponent(theme *Theme) *ui.Base { return b.Component }
func (b *ButtonBase) Clicked() bool                      { return b.clicks.Changed() }
func (b *ButtonBase) Entered() bool                      { return b.enters.Changed() }
func (b *ButtonBase) ClickAction(name string) *input.Action {
	return input.NewAction(name, b.clickTrigger)
}
func (b *ButtonBase) EnterAction(name string) *input.Action {
	return input.NewAction(name, b.enterTrigger)
}
func (b *ButtonBase) Action(name string) *input.Action {
	return input.NewAction(name, b.actionTrigger)
}
