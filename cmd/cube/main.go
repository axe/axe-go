package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/impl/opengl"
	"github.com/axe/axe-go/pkg/input"
	"github.com/axe/axe-go/pkg/ui"
)

func main() {
	runtime.LockOSThread()

	logInput := false
	itb := input.TriggerBuilder{}

	game := axe.NewGame(axe.GameSettings{
		EnableDebug:        true,
		Name:               "Test GLFW",
		FixedDrawFrequency: time.Second / 60,
		FirstStage:         "cube",
		WorldSettings: ecs.WorldSettings{
			EntityCapacity:            2048,
			EntityStageCapacity:       128,
			AverageComponentPerEntity: 4,
			DeleteOnDestroy:           true,
		},
		Windows: []axe.StageWindow{{
			Title:      "Test GLFW Main Window",
			Placement:  ui.Centered(720, 480),
			ClearColor: ui.ColorCornflowerBlue,
		}},
		Stages: []axe.Stage{{
			Name: "cube",
			Assets: []asset.Ref{
				{Name: "cube model", URI: "../assets/cube.obj"},
				{Name: "sans-serif", URI: "../assets/sans-serif.fnt"},
				{Name: "warrior", URI: "../assets/warrior.fnt"},
				{Name: "roboto", URI: "../assets/roboto.fnt"},
			},
			Actions: input.CreateActionSets(input.ActionSetsInput{
				"main": {
					"close":    itb.Key(input.KeyTrigger{Key: input.KeyEscape}),
					"down":     itb.Key(input.KeyTrigger{Key: input.KeyZ}),
					"undo":     itb.Key(input.KeyTrigger{Key: input.KeyZ, CmdCtrl: true}),
					"pasteUp":  itb.Key(input.KeyTrigger{Key: input.KeyV, Ctrl: true, UpOnly: true}),
					"pressA":   itb.Key(input.KeyTrigger{Key: input.KeyA, PressInterval: time.Second / 4, FirstPressDelay: time.Second}),
					"logInput": itb.Key(input.KeyTrigger{Key: input.KeyC}),
					"delete":   itb.Key(input.KeyTrigger{Key: input.KeyBackspace}),
				},
			}),
			Views3: []axe.View3f{{
				Name:      "main",
				Camera:    axe.NewCamera3d(),
				Placement: ui.Maximized(),
			}},
			Scenes3: []axe.Scene3f{{
				Enable: func(scene *axe.Scene3f, game *axe.Game) {
					// Add components & systems
					scene.World.Enable(
						// Component data settings
						ecs.DataSettings{Capacity: 2048, StageCapacity: 128},
						// Components
						axe.TAG, axe.MESH, axe.TRANSFORM3, axe.AUDIO, axe.ACTION, axe.LIGHT, axe.LOGIC, axe.INPUT,
					)
				},
				Load: func(scene *axe.Scene3f, game *axe.Game) {
					// Entities
					e := ecs.New()

					axe.TAG.Set(e, axe.Tag("cube"))

					axe.MESH.Set(e, axe.Mesh{Ref: asset.Ref{Name: "cube model"}})

					axe.TRANSFORM3.Set(e, axe.NewTransform4(axe.TransformCreate4f{
						Position: axe.Vec4f{X: 0, Y: 0, Z: -3, W: 0},
						Scale:    axe.Vec4f{X: 1, Y: 1, Z: 1, W: 0},
					}))

					axe.LOGIC.Set(e, func(e *ecs.Entity, ctx ecs.Context) {
						dt := game.State.UpdateTimer.Elapsed.Seconds()
						transform := axe.TRANSFORM3.Get(e)
						rot := transform.GetRotation()

						rot.X += float32(dt * 6)
						rot.Y += float32(dt * 4)
						transform.SetRotation(rot)
					})

					axe.LIGHT.Set(e, axe.Light{
						Diffuse:  axe.Colorf{R: 1, G: 1, B: 1, A: 1},
						Ambient:  axe.Colorf{R: 0.5, G: 0.5, B: 0.5, A: 1},
						Position: axe.Vec4f{X: -5, Y: 5, Z: 10},
					})

					axe.ACTION.Set(e, axe.InputActionListener{
						Handler: func(action *input.Action) bool {
							switch action.Name {
							case "close":
								game.Running = false
							case "logInput":
								logInput = !logInput
							case "delete":
								e.Delete()
							default:
								inputNames := []string{}
								if action.Data.Inputs != nil {
									for _, in := range action.Data.Inputs {
										inputNames = append(inputNames, in.Name)
									}
								}
								fmt.Printf("%s %0.1f (priority=%d, inputs=%s)\n", action.Name, action.Data.Value, action.Priority(), strings.Join(inputNames, ","))
							}
							return true
						},
					})

					axe.INPUT.Set(e, input.SystemEvents{
						InputChange: func(input input.Input) {
							if logInput {
								fmt.Printf("%s changed to %v\n", input.Name, input.Value)
							}
						},
					})
				},
			}},
			Views2: []axe.View2f{{
				Name:      "ui",
				SceneName: "ui",
				Camera:    axe.NewCamera2d(),
				Placement: ui.Maximized(),
			}},
			Scenes2: []axe.Scene2f{{
				Name: "ui",
				Enable: func(scene *axe.Scene2f, game *axe.Game) {
					scene.World.Enable(
						ecs.DataSettings{Capacity: 1024, StageCapacity: 16},
						axe.TAG, axe.TRANSFORM2, axe.AUDIO, axe.LOGIC, axe.INPUT, axe.UI,
					)
				},
				Load: func(scene *axe.Scene2f, game *axe.Game) {
					e := ecs.New()

					disabled := func(v *ui.Vertex) {
						// v.Color.A *= 0.7
						v.Color.R += (0.5 - v.Color.R) * 0.5
						v.Color.G += (0.5 - v.Color.G) * 0.5
						v.Color.B += (0.5 - v.Color.B) * 0.5
					}

					userInterface := axe.NewUserInterface()
					userInterface.Theme.StateModifier[ui.StateDisabled] = disabled

					// userInterface.UI.Root = ui.NewBuilder().
					// 	Place(ui.Absolute(20, 20, 200, 50)).
					// 	Radius(4).
					// 	ShapeRounded().
					// 	Shrink(4).Shift(1, 4).
					// 	Filled().
					// 	States(ui.StateAny(ui.StateHover|ui.StatePressed|ui.StateFocused|ui.StateSelected)).
					// 	BackgroundColor(ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.3)).
					// 	Layer().
					// 	Bordered(4, ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.3), ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.1).Ptr()).
					// 	Layer().
					// 	States(ui.StateNot(ui.StateHover|ui.StatePressed|ui.StateFocused|ui.StateSelected)).
					// 	BackgroundColor(ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.1)).
					// 	Layer().
					// 	Bordered(4, ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.1), ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.0).Ptr()).
					// 	Layer().
					// 	Maximized().
					// 	BackgroundColor(ui.ColorFromHex("#008080")).
					// 	States(ui.StateNot(ui.StateHover | ui.StatePressed)).
					// 	Layer().
					// 	BackgroundColor(ui.ColorFromHex("#008080").Lighten(0.1)).
					// 	States(ui.StateAll(ui.StateHover)).
					// 	Layer().
					// 	BackgroundColor(ui.ColorFromHex("#008080").Darken(0.1)).
					// 	States(ui.StateAll(ui.StatePressed)).
					// 	Layer().
					// 	Shrink(10).
					// 	Visual(&ui.VisualText{
					// 		Glyphs: simpleTextGlyphs("Press me", "roboto", 16, ui.ColorWhite),
					// 	}).
					// 	End()

					btnPress := newButton(ui.Absolute(20, 20, 200, 50), "{f:warrior}{s:24}{h:0.5}{pv:0.5}Press me", nil)

					btnToggle := newButton(ui.Absolute(20, 100, 200, 50), "{f:roboto}{s:18}{h:0.5}{pv:0.5}TOGGLE DISABLED", func() {
						btnPress.SetDisabled(!btnPress.IsDisabled())
					})

					userInterface.Root = &ui.Base{
						Children: []*ui.Base{
							newDraggable(),
							btnPress,
							btnToggle,
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

func newDraggable() *ui.Base {
	var draggable *ui.Base

	shape := []ui.Coord{
		{X: 0, Y: 0.5},
		{X: 0.5, Y: 0},
		{X: 1, Y: 0.5},
		{X: 0.5, Y: 1},
	}

	draggable = &ui.Base{
		Placement: ui.Absolute(10, 200, 80, 80),
		Draggable: true,
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if ev.Type == ui.DragEventMove {
					draggable.SetPlacement(draggable.Placement.Shift(ev.DeltaMove.X, ev.DeltaMove.Y))
					draggable.BringToFront()
				}
			},
		},
		OverShape: shape,
		Layers: []ui.Layer{{
			Visual:     ui.VisualFilled{Shape: ui.ShapePolygon{Points: shape}},
			Background: ui.BackgroundColor{Color: ui.ColorBlack},
			States:     ui.StateHover.Not,
		}, {
			Visual:     ui.VisualFilled{Shape: ui.ShapePolygon{Points: shape}},
			Background: ui.BackgroundColor{Color: ui.ColorBlack.Lighten(0.3)},
			States:     ui.StateHover.Is,
		}, {
			Visual: ui.MustTextToVisual("{f:roboto}{s:14}{c:white}{h:0.5}{pv:0.5}drag me"),
		}},
	}

	return draggable
}

func newButton(place ui.Placement, text string, onClick func()) *ui.Base {
	shape := ui.ShapeRounded{
		Radius: ui.AmountCorners{
			TopLeft:     ui.Amount{Value: 8},
			TopRight:    ui.Amount{Value: 8},
			BottomLeft:  ui.Amount{Value: 8},
			BottomRight: ui.Amount{Value: 8},
		},
		UnitToPoints: 0.5,
	}

	var button *ui.Base

	button = &ui.Base{
		Placement: place,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				// fmt.Printf("OnPointer: %+v\n", *ev)
				if !ev.Capture && ev.Type == ui.PointerEventDown && onClick != nil {
					// fmt.Printf("event: %+v\n", *ev)
					onClick()
				}
			},
			OnKey: func(ev *ui.KeyEvent) {
				// fmt.Printf("OnKey: %+v\n", *ev)
			},
			OnFocus: func(ev *ui.Event) {
				// fmt.Printf("OnFocus: %+v\n", *ev)
			},
			OnBlur: func(ev *ui.Event) {
				// fmt.Printf("OnBlur: %+v\n", *ev)
			},
		},
		Layers: []ui.Layer{{
			// Shadow filled
			Placement: ui.Maximized().Shrink(4).Shift(1, 4),
			Visual: ui.VisualFilled{
				Shape: shape,
			},
			Background: ui.BackgroundColor{
				Color: ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.5),
			},
			States: (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Is,
		}, {
			// Shadow blur
			Placement: ui.Maximized().Shrink(4).Shift(1, 4),
			Visual: ui.VisualBordered{
				Width:         8,
				OuterColor:    ui.ColorTransparent,
				HasOuterColor: true,
				InnerColor:    ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.5),
				HasInnerColor: true,
				Shape:         shape,
			},
			States: (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Is,
		}, {
			// Shadow filled (default)
			Placement: ui.Maximized().Shrink(4).Shift(1, 4),
			Visual: ui.VisualFilled{
				Shape: shape,
			},
			Background: ui.BackgroundColor{
				Color: ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.2),
			},
			States: (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Not,
		}, {
			// Shadow blur (default)
			Placement: ui.Maximized().Shrink(4).Shift(1, 4),
			Visual: ui.VisualBordered{
				Width:         8,
				OuterColor:    ui.ColorTransparent,
				HasOuterColor: true,
				InnerColor:    ui.ColorFromHex("#008080").Darken(0.5).Alpha(0.2),
				HasInnerColor: true,
				Shape:         shape,
			},
			States: (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Not,
		}, {
			// Background
			Placement: ui.Maximized(),
			Visual: ui.VisualFilled{
				Shape: shape,
			},
			Background: ui.BackgroundColor{
				Color: ui.ColorFromHex("#008080"),
			},
			States: (ui.StateHover | ui.StatePressed).Not,
		}, {
			// Background on hover
			Placement: ui.Maximized(),
			Visual: ui.VisualFilled{
				Shape: shape,
			},
			Background: ui.BackgroundColor{
				Color: ui.ColorFromHex("#008080").Lighten(0.1),
			},
			States: ui.StateHover.Is,
		}, {
			// Background on press
			Placement: ui.Maximized(),
			Visual: ui.VisualFilled{
				Shape: shape,
			},
			Background: ui.BackgroundColor{
				Color: ui.ColorFromHex("#008080").Darken(0.1),
			},
			States: ui.StatePressed.Is,
		}, {
			// Text content
			Placement: ui.Maximized().Shrink(10),
			Visual:    ui.MustTextToVisual(text),
		}},
	}

	return button
}
