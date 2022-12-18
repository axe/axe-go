package ui

type UI struct {
	PointerButtons []PointerButtons
	PointerPoint   Coord
	Root           Component
	PointerOver    Component
	Focused        Component
	Theme          *Theme
}

func getPath(c Component) []Component {
	return getPathWhere(c, nil)
}

func getFocusablePath(c Component) []Component {
	return getPathWhere(c, func(c Component) bool {
		return c.IsFocusable()
	})
}

func getPathWhere(c Component, where func(Component) bool) []Component {
	path := make([]Component, 0)
	curr := c
	for curr != nil {
		if where == nil || where(curr) {
			path = append(path, curr)
		}

		curr = curr.Parent()
	}
	return path
}

func triggerPointerEvent(path []Component, ev PointerEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnPointerEvent(&ev)
	})
}

func triggerKeyEvent(path []Component, ev KeyEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnKeyEvent(&ev)
	})
}

func triggerFocusEvent(path []Component, ev ComponentEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnFocus(&ev)
	})
}

func triggerBlurEvent(path []Component, ev ComponentEvent) {
	triggerEvent(path, &ev.Event, func(c Component) {
		c.OnBlur(&ev)
	})
}

func triggerEvent(path []Component, ev *Event, trigger func(Component)) {
	ev.Capture = true
	ev.Stop = false
	for i := len(path) - 1; i >= 0; i-- {
		trigger(path[i])
		if ev.Stop {
			return
		}
	}
	ev.Capture = false
	for i := 0; i < len(path); i++ {
		trigger(path[i])
		if ev.Stop {
			return
		}
	}
}

func (ui *UI) ProcessKeyEvent(ev KeyEvent) error {
	if ui.Focused != nil {
		triggerKeyEvent(getPath(ui.Focused), ev)
	}

	return nil
}

func (ui *UI) ProcessPointerEvent(ev PointerEvent) error {
	if ui.Root == nil {
		return nil
	}

	// If leave event
	if ev.Type == PointerEventLeave {
		// For every component in ui.MouseOver trigger leave
		if ui.PointerOver != nil {
			triggerPointerEvent(getPath(ui.PointerOver), ev)
		}
		ui.PointerOver = nil
		return nil
	}

	// Handle mouse moving & enter/leave/move events
	if !ui.PointerPoint.Equals(ev.Point) {
		ui.PointerPoint = ev.Point
		over := ui.Root.At(ui.PointerPoint)

		if over != ui.PointerOver {
			oldOver := ComponentMap{}
			if ui.PointerOver != nil {
				oldOver.AddLineage(ui.PointerOver)
			}
			newOver := ComponentMap{}
			if over != nil {
				newOver.AddLineage(over)
			}

			leavePath, movePath, enterPath := oldOver.Compare(newOver)

			// For every component in ui.MouseOver not in over we need to trigger leave
			triggerPointerEvent(leavePath, ev.OfType(PointerEventLeave))

			// For every component in over not in ui.MouseOver we need to trigger enter
			triggerPointerEvent(enterPath, ev.OfType(PointerEventEnter))

			// For every component in both and ev.Type = move we need to trigger move
			triggerPointerEvent(movePath, ev.OfType(PointerEventMove))

			ui.PointerOver = over
		}
	}

	// Handle down/up/wheel event
	if (ev.Type == PointerEventDown || ev.Type == PointerEventUp || ev.Type == PointerEventWheel) && ui.PointerOver != nil {
		triggerPointerEvent(getPath(ui.PointerOver), ev)
	}

	// Change focus on down
	if ev.Type == PointerEventDown {
		if ui.Focused != ui.PointerOver {
			old := getFocusablePath(ui.Focused)
			new := getFocusablePath(ui.PointerOver)

			oldMap := ComponentMap{}
			oldMap.AddMany(old)
			newMap := ComponentMap{}
			newMap.AddMany(new)

			ui.Focused = ui.PointerOver
			ev := ComponentEvent{Target: ui.Focused}

			inOld, _, inNew := oldMap.Compare(newMap)
			triggerBlurEvent(inOld, ev)
			triggerFocusEvent(inNew, ev)
		}
	}

	return nil
}
