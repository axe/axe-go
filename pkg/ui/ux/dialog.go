package ux

import "github.com/axe/axe-go/pkg/ui"

type DialogStyles struct {
	BarHeight       ui.Amount
	BarPadding      ui.AmountBounds
	ChildrenPadding ui.AmountBounds
}

type Dialog struct {
	Title      string
	TitleValue Value[string]

	Children []HasComponent

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
