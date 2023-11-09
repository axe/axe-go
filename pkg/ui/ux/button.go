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
						base.clicks.Add(1)
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
						base.enters.Add(1)
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

	clicks Counter
	enters Counter
}

var _ HasComponent = &ButtonBase{}

func (b *ButtonBase) GetComponent(theme *Theme) *ui.Base { return b.Component }
func (b *ButtonBase) Clicked() bool                      { return b.clicks.Changed() }
func (b *ButtonBase) Entered() bool                      { return b.enters.Changed() }
