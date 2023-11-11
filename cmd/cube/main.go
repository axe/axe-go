package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/id"
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
			Placement:  ui.Centered(720*2, 480*2),
			ClearColor: ui.ColorWhite,
		}},
		Stages: []axe.Stage{{
			Name: "cube",
			Assets: []asset.Ref{
				{Name: "cube model", URI: "../assets/cube.obj"},
				{Name: "sans-serif", URI: "../assets/sans-serif.fnt"},
				{Name: "warrior", URI: "../assets/warrior.fnt"},
				{Name: "roboto", URI: "../assets/roboto.fnt"},
				{Name: "cursors", URI: "../assets/cursors.png"},
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

					userInterface := axe.NewUserInterface()
					userInterface.Theme.TextStyles.Font = id.Get("roboto")

					// Global State effects
					userInterface.Theme.StatePostProcess[ui.StateDisabled] = ui.PostProcessVertex(func(v *ui.Vertex) {
						// v.Color.A *= 0.7
						v.Color.A *= 0.38 // MD
						// v.Color.R += (0.5 - v.Color.R) * 0.5
						// v.Color.G += (0.5 - v.Color.G) * 0.5
						// v.Color.B += (0.5 - v.Color.B) * 0.5
					})

					// Global Animations
					userInterface.Theme.Animations.ForEvent.Set(ui.AnimationEventEnabled, ui.StatelessAnimationFactory(WiggleAnimation))

					// Cursors
					cursors := ui.TileGrid(10, 8, 56, 56, 559, 449, 0, 0, "cursors")
					userInterface.Theme.DefaultCursor = id.Get("pointer")
					userInterface.Theme.Cursors.SetStringMap(map[string]ui.ExtentTile{
						"pointer":  ui.NewExtentTile(cursors[0][0], ui.NewBounds(-7, -7, 49, 49).Scale(0.75)),
						"drag":     ui.NewExtentTile(cursors[4][1], ui.NewBounds(-25, -25, 31, 31).Scale(0.75)),
						"dragging": ui.NewExtentTile(cursors[4][0], ui.NewBounds(-25, -25, 31, 31).Scale(0.75)),
						"text":     ui.NewExtentTile(cursors[0][9], ui.NewBounds(-23, -23, 33, 33).Scale(0.75)),
						"click":    ui.NewExtentTile(cursors[1][0], ui.NewBounds(-21, -6, 35, 50).Scale(0.75)),
						"clicking": ui.NewExtentTile(cursors[1][2], ui.NewBounds(-21, -13, 35, 43).Scale(0.75)),
						"resizebr": ui.NewExtentTile(cursors[4][7], ui.NewBounds(-43, -41, 13, 15).Scale(0.75)),
						"resizer":  ui.NewExtentTile(cursors[4][3], ui.NewBounds(-43, -27, 13, 29).Scale(0.75)),
						"resizeb":  ui.NewExtentTile(cursors[4][4], ui.NewBounds(-26, -42, 30, 24).Scale(0.75)),
						"resizebl": ui.NewExtentTile(cursors[4][8], ui.NewBounds(-11, -41, 45, 15).Scale(0.75)),
					})

					textCoordinates := ui.MustTextToVisual("{h:1}{pa:10}{pv:1}0,0")

					btnPress := newButton(ui.Absolute(20, 20, 200, 50), "{f:warrior}{s:24}{h:0.5}{pv:0.5}Press me", true, nil)

					btnToggle := newButton(ui.Absolute(20, 100, 200, 50), "{f:roboto}{s:18}{h:0.5}{pv:0.5}TOGGLE DISABLED", false, func() {
						btnPress.SetDisabled(!btnPress.IsDisabled())
					})

					textWindow := newWindow("Test Window", ui.Absolute(900, 20, 400, 300))
					textWindow.Children = append(textWindow.Children,
						&ui.Base{
							Placement: ui.MaximizeOffset(10, 34, 10, 10),
							Children: []*ui.Base{{
								Layers: []ui.Layer{{
									Visual: ui.MustTextToVisual(strings.Join([]string{
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
									}, "\n")).Clipped(),
								}},
							}},
						},
					)

					layoutColumnName := id.Get("layoutColumn")
					layoutColumnChange := func(change func(*ui.LayoutColumn)) {
						frame := userInterface.Named.Get(layoutColumnName).(*ui.Base)
						layout := frame.Layout.(*ui.LayoutColumn)
						change(layout)
						frame.Dirty(ui.DirtyPlacement)
					}
					layoutColumnWindow := newWindow("Layout Column", ui.Absolute(10, 500, 300, 300))
					layoutColumnWindow.Children = append(layoutColumnWindow.Children, &ui.Base{
						Name:      layoutColumnName,
						Placement: ui.MaximizeOffset(8, 44, 8, 8),
						// Clip:      ui.Maximized(),
						Layout: &ui.LayoutColumn{
							FullWidth:           false,
							HorizontalAlignment: ui.AlignmentCenter,
							Spacing:             ui.Amount{Value: 10},
						},
						Children: []*ui.Base{
							newButton(ui.Absolute(0, 0, 0, 0), "Toggle Alignment", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn) {
									lc.HorizontalAlignment = ui.Alignment(math.Mod(float64(lc.HorizontalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "{s:20}{h:0.5}{pv:0.5}Toggle FullWidth", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn) {
									lc.FullWidth = !lc.FullWidth
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle EqualWidths", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn) {
									lc.EqualWidths = !lc.EqualWidths
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle FullHeight", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn) {
									lc.FullHeight = !lc.FullHeight
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle FullHeight Weights", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn) {
									if len(lc.FullHeightWeights) == 0 {
										lc.FullHeightWeights = ui.LayoutWeights{1, 2, 3, 4}
									} else {
										lc.FullHeightWeights = nil
									}
								})
							}),
							newButton(ui.Placement{}, "Compact", false, func() {
								layoutColumnWindow.Compact()
							}),
						},
					})

					layoutRowName := id.Get("layoutRow")
					layoutRowChange := func(change func(*ui.LayoutRow)) {
						frame := userInterface.Named.Get(layoutRowName).(*ui.Base)
						layout := frame.Layout.(*ui.LayoutRow)
						change(layout)
						frame.Dirty(ui.DirtyPlacement)
					}
					layoutRowWindow := newWindow("Layout Row", ui.Absolute(800, 500, 300, 300))
					layoutRowWindow.Children = append(layoutRowWindow.Children, &ui.Base{
						Name:      layoutRowName,
						Placement: ui.MaximizeOffset(8, 44, 8, 8),
						// Clip:      ui.Maximized(),
						Layout: &ui.LayoutRow{
							FullHeight:        false,
							VerticalAlignment: ui.AlignmentCenter,
							Spacing:           ui.Amount{Value: 10},
						},
						Children: []*ui.Base{
							newButton(ui.Absolute(0, 0, 0, 0), "Toggle Alignment", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow) {
									lr.VerticalAlignment = ui.Alignment(math.Mod(float64(lr.VerticalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "{s:20}{h:0.5}{pv:0.5}Toggle FullHeight", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow) {
									lr.FullHeight = !lr.FullHeight
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle EqualHeights", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow) {
									lr.EqualHeights = !lr.EqualHeights
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle FullWidth", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow) {
									lr.FullWidth = !lr.FullWidth
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle FullWidth Weights", false, func() {
								layoutRowChange(func(lc *ui.LayoutRow) {
									if len(lc.FullWidthWeights) == 0 {
										lc.FullWidthWeights = ui.LayoutWeights{1, 2, 3, 4}
									} else {
										lc.FullWidthWeights = nil
									}
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle Spacing", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow) {
									lr.Spacing.Value = float32(int(lr.Spacing.Value+5) % 15)
								})
							}),
							newButton(ui.Placement{}, "Compact", false, func() {
								layoutRowWindow.Compact()
							}),
						},
					})

					layoutGridName := id.Get("layoutGrid")
					layoutGridChange := func(change func(*ui.LayoutGrid)) {
						frame := userInterface.Named.Get(layoutGridName).(*ui.Base)
						layout := frame.Layout.(*ui.LayoutGrid)
						change(layout)
						frame.Dirty(ui.DirtyPlacement)
					}
					layoutGridWindow := newWindow("Layout Grid", ui.Absolute(1000, 300, 300, 300))
					layoutGridWindow.Children = append(layoutGridWindow.Children, &ui.Base{
						Name:      layoutGridName,
						Placement: ui.MaximizeOffset(8, 44, 8, 8),
						// Clip:      ui.Maximized(),
						Layout: &ui.LayoutGrid{
							FullHeight:          false,
							FullWidth:           false,
							VerticalAlignment:   ui.AlignmentCenter,
							HorizontalAlignment: ui.AlignmentCenter,
							VerticalSpacing:     ui.Amount{Value: 10},
							HorizontalSpacing:   ui.Amount{Value: 10},
							Columns:             3,
							AspectRatio:         0,
						},
						TextStyles: &ui.TextStylesOverride{
							FontSize: &ui.Amount{Value: 18},
							ParagraphsStylesOverride: &ui.ParagraphsStylesOverride{
								VerticalAlignment: ui.Override(ui.AlignmentCenter),
							},
							ParagraphStylesOverride: &ui.ParagraphStylesOverride{
								HorizontalAlignment: ui.Override(ui.AlignmentCenter),
							},
						},
						Children: []*ui.Base{
							newButton(ui.Placement{}, "Toggle GridFullWidth", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.GridFullWidth = !lg.GridFullWidth
								})
							}),
							newButton(ui.Placement{}, "Toggle GridFullHeight", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.GridFullHeight = !lg.GridFullHeight
								})
							}),
							newButton(ui.Placement{}, "Toggle FullHeight", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.FullHeight = !lg.FullHeight
								})
							}),
							newButton(ui.Placement{}, "Toggle FullWidth", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.FullWidth = !lg.FullWidth
								})
							}),
							newButton(ui.Placement{}, "Toggle EqualHeights", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.EqualHeights = !lg.EqualHeights
								})
							}),
							newButton(ui.Placement{}, "Toggle EqualWidths", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.EqualWidths = !lg.EqualWidths
								})
							}),
							newButton(ui.Placement{}, "Toggle AspectRatio", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.AspectRatio = 1 - lg.AspectRatio
								})
							}),
							newButton(ui.Placement{}, "Toggle Vertical Alignment", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.VerticalAlignment = ui.Alignment(math.Mod(float64(lg.VerticalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Placement{}, "Toggle Horizontal Alignment", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.HorizontalAlignment = ui.Alignment(math.Mod(float64(lg.HorizontalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Placement{}, "Toggle Columns", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.Columns = (lg.Columns + 1) % 6
								})
							}),
							newButton(ui.Placement{}, "Toggle MinHeights", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									if len(lg.MinHeights) == 0 {
										lg.MinHeights = ui.LayoutDimensions{80}
									} else if len(lg.MinHeights) == 1 {
										lg.MinHeights = ui.LayoutDimensions{80, 160}
									} else {
										lg.MinHeights = nil
									}
								})
							}),
							newButton(ui.Placement{}, "Toggle MinWidths", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									if len(lg.MinWidths) == 0 {
										lg.MinWidths = ui.LayoutDimensions{100}
									} else if len(lg.MinWidths) == 1 {
										lg.MinWidths = ui.LayoutDimensions{100, 200}
									} else {
										lg.MinWidths = nil
									}
								})
							}),
							newButton(ui.Placement{}, "Compact", false, func() {
								layoutGridWindow.Compact()
							}),
						},
					})

					layoutInlineName := id.Get("layoutInline")
					layoutInlineChange := func(change func(*ui.LayoutInline)) {
						frame := userInterface.Named.Get(layoutInlineName).(*ui.Base)
						layout := frame.Layout.(*ui.LayoutInline)
						change(layout)
						frame.Dirty(ui.DirtyPlacement)
					}
					layoutInlineWindow := newWindow("Layout Inline", ui.Absolute(20, 700, 300, 300))
					layoutInlineWindow.Children = append(layoutInlineWindow.Children, &ui.Base{
						Name:      layoutInlineName,
						Placement: ui.MaximizeOffset(8, 44, 8, 8),
						Layout: &ui.LayoutInline{
							VerticalAlignment:   ui.AlignmentTop,
							HorizontalAlignment: ui.AlignmentLeft,
							VerticalSpacing:     ui.Amount{Value: 10},
							HorizontalSpacing:   ui.Amount{Value: 10},
						},
						TextStyles: &ui.TextStylesOverride{
							FontSize: &ui.Amount{Value: 20},
							ParagraphsStylesOverride: &ui.ParagraphsStylesOverride{
								VerticalAlignment: ui.Override(ui.AlignmentCenter),
							},
							ParagraphStylesOverride: &ui.ParagraphStylesOverride{
								HorizontalAlignment: ui.Override(ui.AlignmentCenter),
							},
						},
						Children: []*ui.Base{
							newButton(ui.Placement{}, "Toggle Vertical Alignment", false, func() {
								layoutInlineChange(func(lg *ui.LayoutInline) {
									lg.VerticalAlignment = ui.Alignment(math.Mod(float64(lg.VerticalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Placement{}, "Toggle Horizontal Alignment", false, func() {
								layoutInlineChange(func(lg *ui.LayoutInline) {
									lg.HorizontalAlignment = ui.Alignment(math.Mod(float64(lg.HorizontalAlignment)+0.5, 1.5))
								})
							}).Edit(func(b *ui.Base) {
								b.MaxSize.X = 120
							}),
							newButton(ui.Placement{}, "Toggle Vertical Spacing", false, func() {
								layoutInlineChange(func(lg *ui.LayoutInline) {
									lg.VerticalSpacing.Value = float32(math.Mod(float64(lg.VerticalSpacing.Value)+10, 30))
								})
							}),
							newButton(ui.Placement{}, "Toggle Horizontal Spacing", false, func() {
								layoutInlineChange(func(lg *ui.LayoutInline) {
									lg.HorizontalSpacing.Value = float32(math.Mod(float64(lg.HorizontalSpacing.Value)+10, 30))
								})
							}),
							newButton(ui.Placement{}, "Compact", false, func() {
								layoutInlineWindow.Compact()
							}),
						},
					})

					userInterface.Root = &ui.Base{
						Layout: ui.LayoutStatic{
							EnforcePreferredSize: true,
							KeepInside:           true,
						},
						Children: []*ui.Base{
							newDraggable(),
							btnPress,
							btnToggle,
							textWindow,
							layoutColumnWindow,
							layoutRowWindow,
							layoutGridWindow,
							layoutInlineWindow,
							{
								Placement: ui.Placement{
									Left:   ui.Anchor{Delta: 1, Base: -200},
									Right:  ui.Anchor{Delta: 1},
									Top:    ui.Anchor{Delta: 1, Base: -60},
									Bottom: ui.Anchor{Delta: 1},
								},
								Layers: []ui.Layer{{
									Visual: textCoordinates,
								}},
							},
						},
						Events: ui.Events{
							OnPointer: func(ev *ui.PointerEvent) {
								if ev.Type == ui.PointerEventMove {
									textCoordinates.SetText(fmt.Sprintf("{h:1}{pa:10}{pv:1}%.0f,%.0f", ev.Point.X, ev.Point.Y))
								}
							},
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

// Animations

var OriginCenter = ui.AmountPoint{
	X: ui.Amount{Value: 0.5, Unit: ui.UnitParent},
	Y: ui.Amount{Value: 0.5, Unit: ui.UnitParent},
}

var WiggleAnimation = ui.BasicAnimation{
	Duration: 1.0,
	Frames: []ui.BasicAnimationFrame{
		{Time: 0, Rotate: 0, Origin: OriginCenter},
		{Time: .125, Rotate: -45, Origin: OriginCenter},
		{Time: .375, Rotate: 45, Origin: OriginCenter},
		{Time: .583, Rotate: -30, Origin: OriginCenter},
		{Time: .75, Rotate: 30, Origin: OriginCenter},
		{Time: .875, Rotate: -15, Origin: OriginCenter},
		{Time: .9583, Rotate: 15, Origin: OriginCenter},
		{Time: 1, Rotate: 0, Origin: OriginCenter},
	},
}

var RevealAnimation = ui.BasicAnimation{
	Duration: 1.0,
	Frames: []ui.BasicAnimationFrame{
		{Time: 0, Scale: &ui.Coord{X: 1}, Origin: OriginCenter},
		{Time: 1, Scale: &ui.Coord{X: 1, Y: 1}, Origin: OriginCenter},
	},
}

var FadeInAnimation = ui.BasicAnimation{
	Duration: 0.5,
	Frames: []ui.BasicAnimationFrame{
		{Time: 0, Transparency: 1},
		{Time: 1, Transparency: 0},
	},
}

var FadeOutSlideUpAnimation = ui.BasicAnimation{
	Save:     true,
	Duration: 0.7,
	Frames: []ui.BasicAnimationFrame{
		{Time: 0, Origin: OriginCenter},
		{Time: 1, Translate: ui.AmountPoint{Y: ui.Amount{Value: -100}}, Origin: OriginCenter, Transparency: 1},
	},
}

var FadeInSlideDownAnimation = ui.BasicAnimation{
	Save:     true,
	Duration: 0.7,
	Frames: []ui.BasicAnimationFrame{
		{Time: 0, Transparency: 1, Translate: ui.AmountPoint{Y: ui.Amount{Value: -100}}, Origin: OriginCenter},
		{Time: 1, Origin: OriginCenter},
	},
}

var FadeInSlideRightAnimation = ui.BasicAnimation{
	Save:     true,
	Duration: 0.7,
	Frames: []ui.BasicAnimationFrame{
		{Time: 0, Translate: ui.AmountPoint{X: ui.Amount{Value: -100}}, Origin: OriginCenter, Transparency: 1},
		{Time: 1, Origin: OriginCenter},
	},
}

// Temporary component generators

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
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("drag"),
			ui.CursorEventDrag:  id.Get("dragging"),
		}),
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

var buttonColor = ui.ColorFromHex("#008080")
var buttonTemplate = &ui.Template{
	Animations: &ui.Animations{
		ForEvent: ds.NewEnumMap(map[ui.AnimationEvent]ui.AnimationFactory{
			ui.AnimationEventShow: ui.StatelessAnimationFactory(FadeInAnimation),
		}),
	},
	Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
		ui.CursorEventHover: id.Get("click"),
		ui.CursorEventDown:  id.Get("clicking"),
	}),
	Shape: ui.ShapeRounded{
		Radius: ui.AmountCorners{
			TopLeft:     ui.Amount{Value: 8},
			TopRight:    ui.Amount{Value: 8},
			BottomLeft:  ui.Amount{Value: 8},
			BottomRight: ui.Amount{Value: 8},
		},
		UnitToPoints: 0.5,
	},
	PreLayers: []ui.Layer{{
		// Shadow filled
		Placement:  ui.Maximized().Shrink(4).Shift(1, 4),
		Visual:     ui.VisualFilled{},
		Background: ui.BackgroundColor{Color: buttonColor.Darken(0.5).Alpha(0.5)},
		States:     (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Is,
	}, {
		// Shadow blur
		Placement: ui.Maximized().Shrink(4).Shift(1, 4),
		Visual: ui.VisualBordered{
			Width:         8,
			OuterColor:    ui.ColorTransparent,
			HasOuterColor: true,
			InnerColor:    buttonColor.Darken(0.5).Alpha(0.5),
			HasInnerColor: true,
		},
		States: (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Is,
	}, {
		// Shadow filled (default)
		Placement:  ui.Maximized().Shrink(4).Shift(1, 4),
		Visual:     ui.VisualFilled{},
		Background: ui.BackgroundColor{Color: buttonColor.Darken(0.5).Alpha(0.2)},
		States:     (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Not,
	}, {
		// Shadow blur (default)
		Placement: ui.Maximized().Shrink(4).Shift(1, 4),
		Visual: ui.VisualBordered{
			Width:         8,
			OuterColor:    ui.ColorTransparent,
			HasOuterColor: true,
			InnerColor:    buttonColor.Darken(0.5).Alpha(0.2),
			HasInnerColor: true,
		},
		States: (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Not,
	}, {
		// Background
		Placement:  ui.Maximized(),
		Visual:     ui.VisualFilled{},
		Background: ui.BackgroundColor{Color: buttonColor},
		States:     (ui.StateHover | ui.StatePressed).Not,
	}, {
		// Background on hover
		Placement:  ui.Maximized(),
		Visual:     ui.VisualFilled{},
		Background: ui.BackgroundColor{Color: buttonColor.Lighten(0.1)},
		States:     ui.StateHover.Is,
	}, {
		// Background on press
		Placement:  ui.Maximized(),
		Visual:     ui.VisualFilled{},
		Background: ui.BackgroundColor{Color: buttonColor.Darken(0.1)},
		States:     ui.StatePressed.Is,
	}},
	PostLayers: []ui.Layer{{
		// Ripple animation
		Placement: ui.Maximized().Shrink(2),
		Visual: &RippleLayer{
			StartRadius: ui.Amount{Value: 0},
			EndRadius:   ui.Amount{Value: 4, Unit: ui.UnitParent},
			StartColor:  ui.NewColor(1, 1, 1, 0.3),
			EndColor:    ui.ColorTransparent,
			Duration:    1,
		},
	}},
}

func newButton(place ui.Placement, text string, pulse bool, onClick func()) *ui.Base {
	var button *ui.Base

	textVisual := ui.MustTextToVisual(text)

	button = &ui.Base{
		Placement: place,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventDown && onClick != nil {
					onClick()
					ev.Stop = true
				}
			},
		},
		Layers: []ui.Layer{{
			Placement: ui.Maximized().Shrink(10),
			Visual:    textVisual,
		}},
	}

	button.ApplyTemplate(buttonTemplate)

	if pulse {
		button.Layers = append([]ui.Layer{{
			Visual: &PulseLayer{
				StartColor: buttonColor.Lighten(0.2),
				EndColor:   ui.ColorTransparent,
				Duration:   1.5,
				PulseTime:  0.6,
				Size:       12,
			},
			States: ui.StateDefault.Exactly,
		}}, button.Layers...)
	}

	return button
}

type RippleLayer struct {
	StartRadius, EndRadius ui.Amount
	StartColor, EndColor   ui.Color
	Duration               float32
	Time                   float32
	Center                 ui.Coord
	Animating              bool

	animatingOn *ui.Base
}

func (r *RippleLayer) Init(b *ui.Base, init ui.Init) {
	b.Events.OnPointer.Add(func(ev *ui.PointerEvent) {
		if ev.Capture && ev.Type == ui.PointerEventDown {
			r.Center.X, r.Center.Y = b.Bounds.Delta(ev.Point.X, ev.Point.Y)
			r.Animating = true
			r.Time = 0
			r.animatingOn = b
		}
	}, false)
}
func (r *RippleLayer) Update(b *ui.Base, update ui.Update) ui.Dirty {
	if r.Animating && b == r.animatingOn {
		r.Time += float32(update.DeltaTime.Seconds())
		if r.Time > r.Duration {
			r.Animating = false
		} else {
			return ui.DirtyVisual
		}
	}
	return ui.DirtyNone
}
func (r *RippleLayer) Visualize(b *ui.Base, bounds ui.Bounds, ctx *ui.RenderContext, out *ui.VertexBuffers) {
	if r.Animating && b == r.animatingOn {
		centerX, centerY := bounds.Lerp(r.Center.X, r.Center.Y)
		delta := r.Time / r.Duration
		radius := ui.Lerp(r.StartRadius.Get(ctx.AmountContext, true), r.EndRadius.Get(ctx.AmountContext, true), delta)
		color := r.StartColor.Lerp(r.EndColor, delta)

		background := ui.VisualFilled{
			Shape: ui.ShapeRounded{
				Radius: ui.AmountCorners{
					TopLeft:     ui.Amount{Value: radius},
					TopRight:    ui.Amount{Value: radius},
					BottomLeft:  ui.Amount{Value: radius},
					BottomRight: ui.Amount{Value: delta},
				},
				UnitToPoints: 0.5,
			},
		}
		rippleBounds := ui.Bounds{
			Left:   centerX - radius,
			Right:  centerX + radius,
			Top:    centerY - radius,
			Bottom: centerY + radius,
		}
		visualized := ui.NewVertexIterator(out)
		background.Visualize(b, rippleBounds, ctx, out)
		for visualized.HasNext() {
			v := visualized.Next()
			v.X = ui.Clamp(v.X, bounds.Left, bounds.Right)
			v.Y = ui.Clamp(v.Y, bounds.Top, bounds.Bottom)
			v.AddColor(color)
		}
	}
}
func (r *RippleLayer) PreferredSize(b *ui.Base, ctx *ui.RenderContext, maxWidth float32) ui.Coord {
	return ui.Coord{}
}

type PulseLayer struct {
	StartColor, EndColor ui.Color
	Duration             float32
	PulseTime            float32
	Size                 float32
	Time                 float32
	Shape                ui.Shape
}

func (r *PulseLayer) Init(b *ui.Base, init ui.Init) {}
func (r *PulseLayer) Update(b *ui.Base, update ui.Update) ui.Dirty {
	r.Time += float32(update.DeltaTime.Seconds())
	if r.Time > r.Duration {
		r.Time -= r.Duration
	}
	if r.Time < r.PulseTime {
		return ui.DirtyVisual
	}
	return ui.DirtyNone
}
func (r *PulseLayer) Visualize(b *ui.Base, bounds ui.Bounds, ctx *ui.RenderContext, out *ui.VertexBuffers) {
	if r.Time <= r.PulseTime {
		delta := r.Time / r.PulseTime
		size := ui.Lerp(0, r.Size, delta)
		pulseBounds := ui.Bounds{
			Left:   bounds.Left - size,
			Right:  bounds.Right + size,
			Top:    bounds.Top - size,
			Bottom: bounds.Bottom + size,
		}
		color := r.StartColor.Lerp(r.EndColor, delta)
		background := ui.VisualFilled{
			Shape: r.Shape,
		}
		visualized := ui.NewVertexIterator(out)
		background.Visualize(b, pulseBounds, ctx, out)
		for visualized.HasNext() {
			v := visualized.Next()
			v.AddColor(color)
		}
	}
}
func (r *PulseLayer) PreferredSize(b *ui.Base, ctx *ui.RenderContext, maxWidth float32) ui.Coord {
	return ui.Coord{}
}

func newWindow(title string, placement ui.Placement) *ui.Base {
	barSize := float32(36)
	var frame *ui.Base

	frame = &ui.Base{
		Placement: placement,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventDown {
					frame.BringToFront()
				}
			},
		},
		Shape: ui.ShapeRounded{
			Radius: ui.AmountCorners{
				TopLeft:     ui.Amount{Value: 8},
				TopRight:    ui.Amount{Value: 8},
				BottomLeft:  ui.Amount{Value: 8},
				BottomRight: ui.Amount{Value: 8},
			},
			UnitToPoints: 0.5,
		},
		Animations: &ui.Animations{
			ForEvent: ds.NewEnumMap(map[ui.AnimationEvent]ui.AnimationFactory{
				ui.AnimationEventShow: ui.StatelessAnimationFactory(FadeInSlideDownAnimation),
			}),
			Named: id.NewDenseKeyMap[ui.AnimationFactory, uint16, uint8](
				id.WithStringMap(map[string]ui.AnimationFactory{
					"hide": ui.StatelessAnimationFactory(FadeOutSlideUpAnimation),
					"show": ui.StatelessAnimationFactory(FadeInSlideRightAnimation),
				}),
			),
		},
		Layers: []ui.Layer{{
			Placement: ui.Maximized().Shrink(4).Shift(1, 4),
			Visual: ui.VisualBordered{
				Width:         8,
				OuterColor:    ui.ColorTransparent,
				HasOuterColor: true,
				InnerColor:    buttonColor.Darken(0.5).Alpha(0.2),
				HasInnerColor: true,
			},
		}, {
			Visual:     ui.VisualFilled{},
			Background: ui.BackgroundColor{Color: ui.ColorGray.Lighten(0.2)},
		}},
	}

	bar := &ui.Base{
		Placement: ui.Placement{
			Left:   ui.Anchor{Base: 0, Delta: 0},
			Right:  ui.Anchor{Base: 0, Delta: 1},
			Top:    ui.Anchor{Base: 0, Delta: 0},
			Bottom: ui.Anchor{Base: barSize, Delta: 0},
		},
		Shape: ui.ShapeRounded{
			Radius: ui.AmountCorners{
				TopLeft:  ui.Amount{Value: 8},
				TopRight: ui.Amount{Value: 8},
			},
			UnitToPoints: 0.5,
		},
		Layers: []ui.Layer{{
			Visual: ui.VisualFilled{},
			Background: ui.BackgroundLinearGradient{
				StartColor: buttonColor,
				EndColor:   buttonColor.Lighten(0.2),
				End:        ui.Coord{X: 0, Y: 1},
			},
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
					frame.BringToFront()
				case ui.DragEventMove:
					frame.SetPlacement(frame.Placement.Shift(ev.DeltaMove.X, ev.DeltaMove.Y))
				case ui.DragEventEnd:
					frame.Transparency.Set(0)
				}
			},
		},
		Layout: ui.LayoutRow{
			FullWidth:         true,
			FullWidthWeights:  ui.LayoutWeights{1, 0, 0, 0},
			VerticalAlignment: ui.AlignmentCenter,
		},
		Children: []*ui.Base{
			{
				Layers: []ui.Layer{{
					Placement: ui.Maximized().Shrink(2).Shift(6, 0),
					Visual:    ui.MustTextToVisual("{w:none}{s:20}{pv:0.5}{k:-1}" + title),
				}},
			},
			newWindowHide(frame, barSize),
			newWindowMinimizeMaximize(frame, barSize),
			newWindowClose(frame, barSize),
		},
	}

	frame.Children = append(
		frame.Children, bar,
		newWindowResizeRight(frame, barSize),
		newWindowResizeBottom(frame),
		newWindowResizeBottomRight(frame),
		newWindowResizeBottomLeft(frame),
	)

	return frame
}

func newWindowClose(win *ui.Base, barSize float32) *ui.Base {
	return &ui.Base{
		MinSize: ui.Coord{X: barSize, Y: barSize},
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
		Draggable: true,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventDown {
					win.Transparency.Set(1)

					go func() {
						time.Sleep(time.Second * 3)
						win.Transparency.Set(0)
					}()
				}
			},
			OnDrag: func(ev *ui.DragEvent) {
				ev.Cancel = true
				ev.Stop = true
			},
		},
	}
}

func newWindowMinimizeMaximize(win *ui.Base, barSize float32) *ui.Base {
	minimized := win.Placement
	maximized := false

	return &ui.Base{
		MinSize: ui.Coord{X: barSize, Y: barSize},
		Layers: []ui.Layer{{
			Background: ui.BackgroundColor{Color: ui.ColorLightGray.Alpha(0.3)},
			Visual:     ui.VisualFilled{Shape: ui.ShapeRectangle{}},
			States:     ui.StateHover.Is,
		}, {
			Placement: ui.Maximized().Shrink(10),
			Visual: ui.VisualBordered{
				Width: 3,
				Shape: ui.ShapeRectangle{},
				// Double bottom & right borders
				Scales: []ui.VisualBorderScale{
					{NormalX: 1, NormalY: 0, Weight: 2},
					{NormalX: 0, NormalY: 1, Weight: 2},
				},
			},
			Background: ui.BackgroundColor{Color: ui.ColorBlack},
		}},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("click"),
			ui.CursorEventDown:  id.Get("clicking"),
		}),
		Draggable: true,
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
			OnDrag: func(ev *ui.DragEvent) {
				ev.Cancel = true
				ev.Stop = true
			},
		},
	}
}

func newWindowHide(win *ui.Base, barSize float32) *ui.Base {
	return &ui.Base{
		MinSize: ui.Coord{X: barSize, Y: barSize},
		Layers: []ui.Layer{{
			Background: ui.BackgroundColor{Color: ui.ColorLightGray.Alpha(0.3)},
			Visual:     ui.VisualFilled{Shape: ui.ShapeRectangle{}},
			States:     ui.StateHover.Is,
		}, {
			Placement: ui.Absolute(8, barSize-10, barSize-14, 3),
			Visual: ui.VisualFilled{
				Shape: ui.ShapeRectangle{},
			},
			Background: ui.BackgroundColor{Color: ui.ColorBlack},
		}},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("click"),
			ui.CursorEventDown:  id.Get("clicking"),
		}),
		Draggable: true,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventDown {
					win.Play(id.Maybe("hide"))
					go func() {
						time.Sleep(time.Second * 3)
						win.Play(id.Maybe("show"))
					}()
				}
			},
			OnDrag: func(ev *ui.DragEvent) {
				ev.Cancel = true
				ev.Stop = true
			},
		},
	}
}

func newWindowResizeBottomRight(win *ui.Base) *ui.Base {
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
			ui.CursorEventHover: id.Get("resizebr"),
		}),
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if !ev.Capture && !win.Placement.IsMaximized() {
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

func newWindowResizeRight(win *ui.Base, barSize float32) *ui.Base {
	start := win.Placement

	return &ui.Base{
		Draggable: true,
		Placement: ui.Placement{
			Left:   ui.Anchor{Base: -8, Delta: 1},
			Right:  ui.Anchor{Base: 0, Delta: 1},
			Top:    ui.Anchor{Base: barSize, Delta: 0},
			Bottom: ui.Anchor{Base: -8, Delta: 1},
		},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("resizer"),
		}),
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if !ev.Capture && !win.Placement.IsMaximized() {
					ev.Stop = true
					switch ev.Type {
					case ui.DragEventStart:
						win.Transparency.Set(0.2)
						start = win.Placement
					case ui.DragEventMove:
						current := win.Placement
						current.Right.Base += ev.DeltaMove.X
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

func newWindowResizeBottom(win *ui.Base) *ui.Base {
	start := win.Placement

	return &ui.Base{
		Draggable: true,
		Placement: ui.Placement{
			Left:   ui.Anchor{Base: 8, Delta: 0},
			Right:  ui.Anchor{Base: -8, Delta: 1},
			Top:    ui.Anchor{Base: -8, Delta: 1},
			Bottom: ui.Anchor{Base: 0, Delta: 1},
		},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("resizeb"),
		}),
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if !ev.Capture && !win.Placement.IsMaximized() {
					ev.Stop = true
					switch ev.Type {
					case ui.DragEventStart:
						win.Transparency.Set(0.2)
						start = win.Placement
					case ui.DragEventMove:
						current := win.Placement
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

func newWindowResizeBottomLeft(win *ui.Base) *ui.Base {
	start := win.Placement

	return &ui.Base{
		Draggable: true,
		Placement: ui.Placement{
			Left:   ui.Anchor{Base: 0, Delta: 0},
			Right:  ui.Anchor{Base: 8, Delta: 0},
			Top:    ui.Anchor{Base: -8, Delta: 1},
			Bottom: ui.Anchor{Base: 0, Delta: 1},
		},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("resizebl"),
		}),
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if !ev.Capture && !win.Placement.IsMaximized() {
					ev.Stop = true
					switch ev.Type {
					case ui.DragEventStart:
						win.Transparency.Set(0.2)
						start = win.Placement
					case ui.DragEventMove:
						current := win.Placement
						current.Bottom.Base += ev.DeltaMove.Y
						current.Left.Base += ev.DeltaMove.X
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
