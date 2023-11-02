package main

import (
	"runtime"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/impl/opengl"
	"github.com/axe/axe-go/pkg/ui"
)

func main() {
	runtime.LockOSThread()

	game := axe.NewGame(axe.GameSettings{
		EnableDebug:        true,
		Name:               "Test Window",
		FixedDrawFrequency: time.Second / 60,
		FirstStage:         "win",
		WorldSettings: ecs.WorldSettings{
			EntityCapacity:            2048,
			EntityStageCapacity:       128,
			AverageComponentPerEntity: 4,
			DeleteOnDestroy:           true,
		},
		Windows: []axe.StageWindow{{
			Title:     "Test Window",
			Placement: ui.Centered(720, 480),
		}},
		Stages: []axe.Stage{{
			Name: "win",
			Assets: []asset.Ref{
				{Name: "roboto", URI: "../assets/roboto.fnt"},
			},
			Views2: []axe.View2f{{
				Camera: axe.NewCamera2d(),
			}},
			Scenes2: []axe.Scene2f{{
				Enable: func(scene *axe.Scene2f, game *axe.Game) {
					scene.World.Enable(
						ecs.DataSettings{Capacity: 1024, StageCapacity: 16},
						axe.INPUT, axe.UI,
					)
				},
				Load: func(scene *axe.Scene2f, game *axe.Game) {
					e := ecs.New()

					userInterface := axe.NewUserInterface()
					userInterface.Root = &ui.Base{
						Children: []*ui.Base{
							generateWindow("Test window", ui.Absolute(20, 20, 300, 250)),
						},
					}

					axe.UI.Set(e, userInterface)
					axe.INPUT.Set(e, userInterface.GetInputEventsHandler())
				},
			}},
		}},
	})

	opengl.Setup(game, opengl.Settings{})

	err := game.Run()
	if err != nil {
		panic(err)
	}
}

func generateWindow(title string, placement ui.Placement) *ui.Base {
	frameOutline := ui.OutlineRounded{
		Radius: ui.AmountCorners{
			TopLeft:     ui.Amount{Value: 8},
			TopRight:    ui.Amount{Value: 8},
			BottomLeft:  ui.Amount{Value: 8},
			BottomRight: ui.Amount{Value: 8},
		},
		UnitToPoints: 0.5,
	}
	frame := &ui.Base{
		Placement: placement,
		Layers: []ui.Layer{{
			Visual:     ui.VisualFilled{Outline: frameOutline},
			Background: ui.BackgroundColor{Color: ui.ColorGray},
		}},
	}

	barOutline := ui.OutlineRounded{
		Radius: ui.AmountCorners{
			TopLeft:  ui.Amount{Value: 8},
			TopRight: ui.Amount{Value: 8},
		},
		UnitToPoints: 0.5,
	}
	bar := &ui.Base{
		Placement: ui.Placement{
			Left:   ui.Anchor{Base: 0, Delta: 0},
			Right:  ui.Anchor{Base: 0, Delta: 1},
			Top:    ui.Anchor{Base: 0, Delta: 0},
			Bottom: ui.Anchor{Base: 24, Delta: 0},
		},
		Layers: []ui.Layer{{
			Visual: ui.VisualFilled{Outline: barOutline},
			Background: ui.BackgroundLinearGradient{
				StartColor: ui.ColorCornflowerBlue,
				EndColor:   ui.ColorCornflowerBlue.Lighten(0.1),
				End:        ui.Coord{X: 0, Y: 1},
			},
		}, {
			Placement: ui.Maximized().Shrink(2).Shift(4, 0),
			Visual:    ui.MustTextToVisual("{s:18}{pv:0.5}" + title),
		}},
		Draggable: true,
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if ev.Type == ui.DragEventMove {
					frame.SetPlacement(frame.Placement.Shift(ev.DeltaMove.X, ev.DeltaMove.Y))
				}
			},
		},
	}

	frame.Children = append(frame.Children, bar)

	return frame
}
