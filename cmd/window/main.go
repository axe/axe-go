package main

import (
	"runtime"
	"strings"
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
	frameShape := ui.ShapeRounded{
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
			Visual:     ui.VisualFilled{Shape: frameShape},
			Background: ui.BackgroundColor{Color: ui.ColorGray},
		}},
	}

	barShape := ui.ShapeRounded{
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
			Visual: ui.VisualFilled{Shape: barShape},
			Background: ui.BackgroundLinearGradient{
				StartColor: ui.ColorCornflowerBlue,
				EndColor:   ui.ColorCornflowerBlue.Lighten(0.2),
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

	lines := []string{
		"{c:white}{s:150%f}{ls:100%f}{ps:100%f}Dear Reader,",
		"{p}{h:0.5}This is centered.",
		"{v:0.5}And {s:300%f}{f:warrior}THIS{s:150%f}{f} is big!",
		"{v:1}This is bottom & center {s:300%f}aligned?",
		"{p}{h:0}{v:0}Top{s:150%f} and left aligned.",
		"{p}{h:0.5}{c:red}And {c:orange}this {c:yellow}line {c:green}is {c:blue}super {c:indigo}duper {c:violet}gay!",
		"{p}{h:1}{c:white}Right aligned!",
		"{p}{h:0.25}25% aligned?",
		"{p}{h}{w:word}This should wrap at the word and not at the character and should take up at least two lines. Resize the window!",
		"{p}{pt:20}{h:0.5}{w:char}This should wrap at the character and not at the word and be centered.",
	}
	text := &ui.Base{
		Placement: ui.MaximizeOffset(10, 34, 10, 10),
		Clip:      ui.Maximized(),
		Children: []*ui.Base{{
			Layers: []ui.Layer{{
				Visual: ui.MustTextToVisual(strings.Join(lines, "\n")),
			}},
		}},
	}

	frame.Children = append(frame.Children, bar, text)

	return frame
}
