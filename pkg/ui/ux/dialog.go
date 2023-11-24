package ux

import "github.com/axe/axe-go/pkg/ui"

type DialogSettings struct {
	BarHeight        ui.Amount
	BarPadding       ui.AmountBounds
	ChildrenPadding  ui.AmountBounds
	DragTransparency float32
}

type Dialog struct {
	Title      string
	TitleValue Value[string]

	Children []HasComponent

	RetainOrderOnDrag        bool
	HideMaximize             bool
	HideClose                bool
	AllowOutside             bool
	DisableDragging          bool
	DisableResize            bool
	DisableResizeLeft        bool
	DisableResizeBottomLeft  bool
	DisableResizeBottom      bool
	DisableResizeBottomRight bool
	DisableResizeRight       bool

	OnResize   func(b ui.Bounds)
	OnDrag     func(b ui.Bounds)
	OnMaximize func()
	OnMinimize func()
	OnClose    func()
}
