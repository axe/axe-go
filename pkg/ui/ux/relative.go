package ux

import "github.com/axe/axe-go/pkg/ui"

// relatively placed, different render parent. hidden by default
// desired placement
// allowed placements (when it can't fit)
// how to determine render parent (by default, root)

type RelativeOnEvent int

const (
	RelativeOnNone RelativeOnEvent = (1 << iota) >> 1
	RelativeOnPointerEnter
	RelativeOnPointerLeave
	RelativeOnFocus
	RelativeOnBlur
	RelativeOnClick
	RelativeOnEscape
	RelativeOnSpace
	RelativeOnOutside
	RelativeOnKeyDown
)

type RelativePlacement struct {
	Anchor                  ui.AmountPoint
	Align                   ui.AmountPoint
	Offset                  ui.AmountPoint
	Distance                float32
	MatchWidth, MatchHeight bool
}

type RelativeSettings struct {
	HideOn        RelativeOnEvent
	ShowOn        RelativeOnEvent
	ToggleOn      RelativeOnEvent
	ShowAnimation ui.AnimationFactory
	HideAnimation ui.AnimationFactory
}

type Relative struct {
	RelativeSettings

	Placement  RelativePlacement
	RelativeTo HasComponent
	Content    HasComponent

	OnShow Listener[*RelativeBase]
	OnHide Listener[*RelativeBase]
}

var _ HasComponent = Relative{}

type RelativeBase struct {
	Component  *ui.Base
	RelativeTo *ui.Base
	Content    *ui.Base
	Enabled    Value[bool]
	Visible    Value[bool]
	Options    Relative

	visible bool
}

var _ HasComponent = &RelativeBase{}

func (r Relative) Build(theme *Theme) *RelativeBase {
	relativeTo := r.RelativeTo.GetComponent(theme)
	// content := r.Content.GetComponent(theme)
	contentWrapper := &ui.Base{}

	component := &ui.Base{
		Placement: relativeTo.Placement,
		Layout: ui.LayoutStatic{
			EnforcePreferredSize: true,
		},
		Children: []*ui.Base{
			relativeTo,
			contentWrapper,
		},
	}
	return &RelativeBase{
		Component: component,
		Options:   r,
	}
}

func (b Relative) GetComponent(theme *Theme) *ui.Base {
	return b.Build(theme).GetComponent(theme)
}

func (rb *RelativeBase) GetComponent(theme *Theme) *ui.Base { return rb.Component }
func (rb *RelativeBase) Show() {
	if !rb.visible {
		if rb.Content.Parent() != rb.Component {
			rb.Component.AddChildren(rb.Content)
			rb.Content.SetRenderParent(rb.Component.UI().Root)
		} else {
			rb.Content.Show()
		}
		rb.visible = true
		rb.Options.OnShow.Trigger(rb)
	}
}
func (rb *RelativeBase) Hide() {
	if rb.visible {
		rb.Content.Hide()
		rb.visible = false
		rb.Options.OnHide.Trigger(rb)
	}
}
func (rb *RelativeBase) Toggle() {
	if rb.visible {
		rb.Hide()
	} else {
		rb.Show()
	}
}
func (rb *RelativeBase) Trigger(on RelativeOnEvent) {
	if rb.Options.ShowOn&on != 0 {
		rb.Show()
	} else if rb.Options.HideOn&on != 0 {
		rb.Hide()
	} else if rb.Options.ToggleOn&on != 0 {
		rb.Toggle()
	}
}
