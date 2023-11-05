package axe

import (
	"time"

	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/input"
	"github.com/axe/axe-go/pkg/ui"
)

var UI = ecs.DefineComponent("ui", UserInterface{}).SetSystem(NewUserInterfaceSystem())

type UserInterface struct {
	*ui.UI
}

func NewUserInterface() UserInterface {
	return UserInterface{
		UI: ui.NewUI(),
	}
}

func UserInterfaceInputEventsFor(e *ecs.Entity) InputEvents {
	return InputEvents{
		InputChange: func(i input.Input) {
			inputSystem := ActiveGame().Input
			switch i.Device.Type {
			case input.DeviceTypeKeyboard:
				info := input.KeyInfos[i.Name]

				if info.Lower != 0 {
					value := info.Lower
					leftShift := inputSystem.Get(input.KeyLeftShift)
					rightShift := inputSystem.Get(input.KeyRightShift)
					if (leftShift != nil && leftShift.Value > 0) || (rightShift != nil && rightShift.Value > 0) {
						value = info.Upper
					}
					keyType := ui.KeyEventDown
					if i.Value == 0 {
						keyType = ui.KeyEventUp
					}
					gui := UI.Get(e)
					gui.ProcessKeyEvent(ui.KeyEvent{
						Event: ui.Event{
							Time: time.Now(),
						},
						Key:  i.Name,
						Char: value,
						Type: keyType,
					})
				}

			case input.DeviceTypeMouse:
				p := *inputSystem.Points()[0]
				pointerType := ui.PointerEventDown
				if i.Value == 0 {
					pointerType = ui.PointerEventUp
				}

				gui := UI.Get(e)
				switch i.Name {
				case "MouseButton0": // left
					gui.ProcessPointerEvent(newPointerEvent(p, pointerType, 0, 0))
				case "MouseButton1": // right
					gui.ProcessPointerEvent(newPointerEvent(p, pointerType, 1, 0))
				case "MouseButton2": // middle
					gui.ProcessPointerEvent(newPointerEvent(p, ui.PointerEventWheel, 2, int(i.Value)))
				}
			}
		},
		PointChange: func(p input.Point) {
			gui := UI.Get(e)
			gui.ProcessPointerEvent(newPointerEvent(p, ui.PointerEventMove, 0, 0))
		},
		PointLeave: func(p input.Point) {
			gui := UI.Get(e)
			gui.ProcessPointerEvent(newPointerEvent(p, ui.PointerEventLeave, 0, 0))
		},
		PointEnter: func(p input.Point) {
			gui := UI.Get(e)
			gui.ProcessPointerEvent(newPointerEvent(p, ui.PointerEventEnter, 0, 0))
		},
	}
}

func newPointerEvent(p input.Point, eventType ui.PointerEventType, button, amount int) ui.PointerEvent {
	return ui.PointerEvent{
		Type: eventType,
		Event: ui.Event{
			Time: time.Now(),
		},
		Point: ui.Coord{
			X: float32(p.X),
			Y: float32(p.Y),
		},
		Button: button,
		Amount: amount,
	}
}

type UserInterfaceSystem struct{}

func NewUserInterfaceSystem() ecs.DataSystem[UserInterface] {
	return &UserInterfaceSystem{}
}

func (sys UserInterfaceSystem) OnStage(data *UserInterface, e *ecs.Entity, ctx ecs.Context) {
}
func (sys UserInterfaceSystem) OnLive(data *UserInterface, e *ecs.Entity, ctx ecs.Context) {
	game := ActiveGame()

	for _, a := range game.Assets.Assets.Assets {
		if font, ok := a.Data.(*ui.Font); ok {
			fontIdentifier := id.Get(font.Name)
			data.Theme.Fonts.Set(fontIdentifier, font)

			if data.Theme.TextStyles.Font.Empty() {
				data.Theme.TextStyles.Font = fontIdentifier
			}
		}
	}

	data.Init(ui.Init{
		Theme: data.Theme,
	})
}
func (sys UserInterfaceSystem) OnRemove(data *UserInterface, e *ecs.Entity, ctx ecs.Context) {
}
func (sys UserInterfaceSystem) Init(ctx ecs.Context) error {
	return nil
}
func (sys UserInterfaceSystem) Update(iter ds.Iterable[ecs.Value[*UserInterface]], ctx ecs.Context) {
	game := ActiveGame()
	update := ui.Update{
		DeltaTime: game.State.UpdateTimer.Elapsed,
	}
	guis := iter.Iterator()
	for guis.HasNext() {
		guis.Next().Data.Update(update)
	}
}
func (sys UserInterfaceSystem) Destroy(ctx ecs.Context) {
}
