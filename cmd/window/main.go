package main

import (
	"runtime"
	"strings"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/id"
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
				{Name: "cursors", URI: "../assets/cursors.png"},
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

					// Cursors
					cursors := ui.TileGrid(10, 8, 56, 56, 559, 449, 0, 0, "cursors")
					userInterface.Theme.DefaultCursor = id.Get("pointer")
					userInterface.Theme.Cursors.SetStringMap(map[string]ui.ExtentTile{
						"pointer":      ui.NewExtentTile(cursors[0][0], ui.NewBounds(-7, -7, 49, 49).Scale(0.75)),
						"drag":         ui.NewExtentTile(cursors[4][1], ui.NewBounds(-25, -25, 31, 31).Scale(0.75)),
						"dragging":     ui.NewExtentTile(cursors[4][0], ui.NewBounds(-25, -25, 31, 31).Scale(0.75)),
						"text":         ui.NewExtentTile(cursors[0][9], ui.NewBounds(-23, -23, 33, 33).Scale(0.75)),
						"click":        ui.NewExtentTile(cursors[1][0], ui.NewBounds(-21, -6, 35, 50).Scale(0.75)),
						"clicking":     ui.NewExtentTile(cursors[1][2], ui.NewBounds(-21, -13, 35, 43).Scale(0.75)),
						"resizecorner": ui.NewExtentTile(cursors[4][7], ui.NewBounds(-38, -38, 18, 18).Scale(0.75)),
					})

					userInterface.Root = &ui.Base{
						Children: []*ui.Base{
							generateWindow("Test window", ui.Absolute(20, 20, 300, 250)),
						},
					}

					axe.UI.Set(e, userInterface)
					axe.INPUT.Set(e, axe.UserInterfaceInputEventsFor(e))
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
	barSize := float32(36)
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
			Bottom: ui.Anchor{Base: barSize, Delta: 0},
		},
		Layers: []ui.Layer{{
			Visual: ui.VisualFilled{Shape: barShape},
			Background: ui.BackgroundLinearGradient{
				StartColor: ui.ColorCornflowerBlue,
				EndColor:   ui.ColorCornflowerBlue.Lighten(0.2),
				End:        ui.Coord{X: 0, Y: 1},
			},
		}, {
			Placement: ui.Maximized().Shrink(2).Shift(6, 0),
			Visual:    ui.MustTextToVisual("{s:20}{pv:0.5}" + title),
		}},
		Draggable: true,
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if ev.Capture {
					return
				}
				switch ev.Type {
				case ui.DragEventStart:
					frame.Transparency.Set(0.2)
				case ui.DragEventMove:
					shifted := frame.Placement.Shift(ev.DeltaMove.X, ev.DeltaMove.Y)
					// shifted.Constrain(frame.Parent())
					frame.SetPlacement(shifted)
				case ui.DragEventEnd:
					frame.Transparency.Set(0)
				}
			},
		},
		Children: []*ui.Base{
			newWindowClose(frame, barSize),
			newWindowMinimizeMaximize(frame, barSize),
		},
	}

	lines := []string{
		"{c:black}{s:150%f}{ls:100%f}{ps:100%f}Dear Reader,",
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
		Placement: ui.MaximizeOffset(10, 34, -10, -10),
		Children: []*ui.Base{{
			Layers: []ui.Layer{{
				Visual: ui.MustTextToVisual(strings.Join(lines, "\n")).Clipped(),
			}},
		}},
	}

	frame.Children = append(frame.Children, bar, text, newWindowResize(frame))

	return frame
}

func newWindowClose(win *ui.Base, barSize float32) *ui.Base {
	return &ui.Base{
		Placement: ui.Placement{
			Left:   ui.Anchor{Base: -barSize, Delta: 1},
			Right:  ui.Anchor{Delta: 1},
			Bottom: ui.Anchor{Base: barSize},
		},
		Layers: []ui.Layer{{
			Background: ui.BackgroundColor{Color: ui.ColorLightGray.Alpha(0.3)},
			Visual:     ui.VisualFilled{Shape: ui.ShapeRectangle{}},
			States:     ui.StateHover.Is,
		}, {
			Placement: ui.Maximized().Shrink(8),
			Visual: ui.VisualFilled{
				Shape: ui.ShapePolygon{
					Points: []ui.Coord{
						{X: 0, Y: 0}, {X: 0.1, Y: 0}, {X: 0.5, Y: 0.4}, {X: 0.9, Y: 0},
						{X: 1, Y: 0}, {X: 1.0, Y: 0.1}, {X: 0.6, Y: 0.5}, {X: 1, Y: 0.9},
						{X: 1, Y: 1}, {X: 0.9, Y: 1}, {X: 0.5, Y: 0.6}, {X: 0.1, Y: 1},
						{X: 0, Y: 1}, {X: 0, Y: 0.9}, {X: 0.4, Y: 0.5}, {X: 0, Y: 0.1},
					},
				},
			},
			Background: ui.BackgroundColor{Color: ui.ColorBlack},
		}},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("click"),
			ui.CursorEventDown:  id.Get("clicking"),
		}),
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventDown {
					win.Transparency.Set(1)
				}
			},
		},
	}
}

func newWindowMinimizeMaximize(win *ui.Base, barSize float32) *ui.Base {
	minimized := win.Placement
	maximized := false

	return &ui.Base{
		Placement: ui.Placement{
			Left:   ui.Anchor{Base: -barSize * 2, Delta: 1},
			Right:  ui.Anchor{Base: -barSize, Delta: 1},
			Bottom: ui.Anchor{Base: barSize},
		},
		Layers: []ui.Layer{{
			Background: ui.BackgroundColor{Color: ui.ColorLightGray.Alpha(0.3)},
			Visual:     ui.VisualFilled{Shape: ui.ShapeRectangle{}},
			States:     ui.StateHover.Is,
		}, {
			Placement: ui.Maximized().Shrink(10),
			Visual: ui.VisualBordered{
				Width: 3,
				Shape: ui.ShapeRectangle{},
			},
			Background: ui.BackgroundColor{Color: ui.ColorBlack},
		}},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("click"),
			ui.CursorEventDown:  id.Get("clicking"),
		}),
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventDown {
					if maximized {
						win.SetPlacement(minimized)
					} else {
						minimized = win.Placement
						win.SetPlacement(ui.Maximized())
					}
					maximized = !maximized
				}
			},
		},
	}
}

func newWindowResize(win *ui.Base) *ui.Base {
	start := win.Placement

	return &ui.Base{
		Draggable: true,
		Placement: ui.Placement{
			Left:   ui.Anchor{Base: -8, Delta: 1},
			Right:  ui.Anchor{Base: 0, Delta: 1},
			Top:    ui.Anchor{Base: -8, Delta: 1},
			Bottom: ui.Anchor{Base: 0, Delta: 1},
		},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("resizecorner"),
		}),
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if !ev.Capture {
					ev.Stop = true
					switch ev.Type {
					case ui.DragEventStart:
						win.Transparency.Set(0.2)
						start = win.Placement
					case ui.DragEventMove:
						current := win.Placement
						current.Right.Base += ev.DeltaMove.X
						current.Bottom.Base += ev.DeltaMove.Y
						win.SetPlacement(current)
					case ui.DragEventCancel:
						win.SetPlacement(start)
					case ui.DragEventEnd:
						win.Transparency.Set(0)
					}
				}
			},
		},
	}
}
