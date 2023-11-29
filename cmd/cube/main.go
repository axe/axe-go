package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/color"
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/ease"
	"github.com/axe/axe-go/pkg/ecs"
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/impl/opengl"
	"github.com/axe/axe-go/pkg/input"
	"github.com/axe/axe-go/pkg/ui"
	"github.com/axe/axe-go/pkg/ui/ua"
	"github.com/axe/axe-go/pkg/util"
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
			ClearColor: color.White,
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
						Diffuse:  color.White,
						Ambient:  color.StormGray,
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
					userInterface.TransformPointer = true
					userInterface.TransparencyThreshold = 0.9
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
					userInterface.Theme.Animations.ForEvent.Set(ui.AnimationEventEnabled, ua.Wiggle)

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
						"resizel":  ui.NewExtentTile(cursors[4][5], ui.NewBounds(-10, -27, 46, 29).Scale(0.75)),
					})

					// Colors
					userInterface.Theme.TextStyles.Color = TextColor
					userInterface.Theme.Colors.Set(PrimaryColor, color.CornflowerBlue /* color.FromHex("#009090")*/)
					userInterface.Theme.Colors.Set(SecondaryColor, color.Purple)
					userInterface.Theme.Colors.Set(BackgroundColor, color.White)
					userInterface.Theme.Colors.Set(TextColor, color.Black)

					textCoordinates := ui.MustTextToVisual("{h:1}{pa:10}{pv:1}0,0")

					btnPress := newButton(ui.Absolute(20, 20, 200, 50), "{f:warrior}{s:24}{h:0.5}{pv:0.5}Press me", true, nil)
					newDropdown(
						btnPress,
						util.SliceSorted(ua.Named.Keys(), func(a, b id.Identifier) bool {
							return a.String() < b.String()
						}),
						func(id id.Identifier) string {
							return id.String()
						},
						func(id id.Identifier) {
							anim := ua.Named.Get(id)
							if basic, ok := anim.(ui.BasicAnimation); ok {
								anim = basic.WithNoSave()
							}
							btnPress.Play(anim)
						},
					)

					btnToggle := newButton(ui.Absolute(20, 100, 200, 50), "{f:roboto}{s:18}{h:0.5}{pv:0.5}TOGGLE DISABLED", false, func() {
						btnPress.SetDisabled(!btnPress.IsDisabled())
					})
					newTooltip("{w:none}This is a tooltip", 1, 5, btnToggle)

					textAnimation := ui.BasicTextAnimation{
						Settings: []ui.BasicTextAnimationSettings{{
							Kind: ui.BasicTextAnimationKindChar,
							Frames: []ui.BasicAnimationFrame{{
								Translate:    ui.NewAmountPoint(0, 40),
								Scale:        &ui.Coord{X: 4, Y: 4},
								Origin:       ui.NewAmountPointUnit(0.5, 1, ui.UnitParent),
								Transparency: 1,
								Time:         0,
							}, {
								Scale:        &ui.Coord{X: 1, Y: 1},
								Origin:       ui.NewAmountPointUnit(0.5, 1, ui.UnitParent),
								Transparency: 0,
								Time:         1,
							}},
							Duration: 0.5,
							Delay:    0.2,
						}, {
							Start: 11,
							Kind:  ui.BasicTextAnimationKindWord,
							Frames: []ui.BasicAnimationFrame{{
								Translate:    ui.NewAmountPoint(0, -10),
								Transparency: 1,
								Time:         0,
							}, {
								Transparency: 0,
								Time:         1,
							}},
							Duration: 1,
							Delay:    0.3,
						}, {
							Start: 39,
							Kind:  ui.BasicTextAnimationKindLine,
							Frames: []ui.BasicAnimationFrame{{
								Translate:    ui.NewAmountPoint(-100, 0),
								Transparency: 1,
								Time:         0,
							}, {
								Transparency: 0,
								Time:         1,
							}},
							Duration: 2,
							Delay:    1,
						}, {
							Start: 159,
							Kind:  ui.BasicTextAnimationKindColumn,
							Frames: []ui.BasicAnimationFrame{{
								Translate:    ui.NewAmountPoint(-10, 0),
								Transparency: 1,
								Time:         0,
							}, {
								Transparency: 0,
								Time:         1,
							}},
							Duration: 0.3,
							Delay:    0.3,
						}},
					}

					textWindow := newWindow("Test Window", ui.Absolute(900, 20, 500, 400))
					textWindowVisual := ui.MustTextToVisual(strings.Join([]string{
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
					}, "\n"))
					textWindow.Children = append(textWindow.Children,
						&ui.Base{
							Placement: ui.MaximizeOffset(10, 44, 10, 10),
							Children: []*ui.Base{
								newScrollingSection(20, &ui.Base{
									Layout: ui.LayoutStatic{
										EnforcePreferredSize: true,
									},
									Children: []*ui.Base{{
										Layers: []ui.Layer{{
											Visual: textWindowVisual.Play(textAnimation),
										}},
										Events: ui.Events{
											OnPointer: func(ev *ui.PointerEvent) {
												if !ev.Capture && ev.Type == ui.PointerEventDown {
													text := ev.Target.Layers[0].Visual.(*ui.VisualText)
													rendered := text.Rendered()
													closest := rendered.ClosestByLine(ev.Point.X, ev.Point.Y)
													glyph := rendered.Glyphs[closest]
													input := rendered.Paragraphs.Paragraphs[glyph.Paragraph].Glyphs[glyph.ParagraphIndex]
													fmt.Printf("Clicked on %s in paragraph %d and index %d\n", input.String(), glyph.Paragraph, glyph.Index)
												}
											},
										},
									},
										newButton(ui.Placement{}.Attach(1, 0, 0, 0), "{h:center}{w:none}Do a barrel roll!", false, func() {
											textWindow.Play(ui.BasicAnimation{
												Duration: 1.0,
												Easing:   ease.TinyBounce,
												Save:     true,
												Frames: []ui.BasicAnimationFrame{
													{Rotate: 360, Time: 0, Origin: ui.NewAmountPointParent(0.5, 0.5)},
													{Rotate: 0, Time: 1, Origin: ui.NewAmountPointParent(0.5, 0.5)},
												},
											})
										}),
										newButton(ui.Placement{}.Attach(1, 0, 0, 0).Shift(0, 40), "{h:center}{w:none}Animate text!", false, func() {
											textWindowVisual.Play(textAnimation)
										}),
									},
								}),
							},
						},
					)
					textWindow.Colors.Set(BackgroundColor, color.Gray)

					layoutColumnName := id.Get("layoutColumn")
					layoutColumnChange := func(change func(*ui.LayoutColumn, *ui.Base)) {
						frame := userInterface.Named.Get(layoutColumnName)
						layout := frame.Layout.(*ui.LayoutColumn)
						change(layout, frame)
						frame.Relayout()
					}
					layoutColumnWindow := newWindow("Layout Column", ui.Absolute(10, 500, 300, 300))
					layoutColumnWindow.Children = append(layoutColumnWindow.Children, &ui.Base{
						Name:      layoutColumnName,
						Placement: ui.MaximizeOffset(8, 44, 8, 8),
						Layout: &ui.LayoutColumn{
							FullWidth:           false,
							HorizontalAlignment: ui.AlignmentCenter,
							Spacing:             ui.Amount{Value: 10},
						},
						Children: []*ui.Base{
							newButton(ui.Absolute(0, 0, 0, 0), "Toggle Alignment", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn, p *ui.Base) {
									lc.HorizontalAlignment = ui.Alignment(math.Mod(float64(lc.HorizontalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "{s:20}{h:0.5}{pv:0.5}Toggle FullWidth", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn, p *ui.Base) {
									lc.FullWidth = !lc.FullWidth
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle EqualWidths", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn, p *ui.Base) {
									lc.EqualWidths = !lc.EqualWidths
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle FullHeight", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn, p *ui.Base) {
									lc.FullHeight = !lc.FullHeight
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle FullHeight Weights", false, func() {
								layoutColumnChange(func(lc *ui.LayoutColumn, p *ui.Base) {
									if p.Children[0].LayoutData.HeightWeight == 0 {
										p.Children[0].LayoutData.HeightWeight = 1
										p.Children[1].LayoutData.HeightWeight = 2
										p.Children[2].LayoutData.HeightWeight = 3
										p.Children[3].LayoutData.HeightWeight = 4
									} else {
										p.Children[0].LayoutData.HeightWeight = 0
										p.Children[1].LayoutData.HeightWeight = 0
										p.Children[2].LayoutData.HeightWeight = 0
										p.Children[3].LayoutData.HeightWeight = 0
									}
								})
							}),
							newButton(ui.Placement{}, "MinWidth", false, func() {
								layoutColumnWindow.PlaceMinWidth()
							}),
							newButton(ui.Placement{}, "MinHeight", false, func() {
								layoutColumnWindow.PlaceMinHeight()
							}),
						},
					})

					layoutRowName := id.Get("layoutRow")
					layoutRowChange := func(change func(*ui.LayoutRow, *ui.Base)) {
						frame := userInterface.Named.Get(layoutRowName)
						layout := frame.Layout.(*ui.LayoutRow)
						change(layout, frame)
						frame.Relayout()
					}
					layoutRowWindow := newWindow("Layout Row", ui.Absolute(800, 500, 300, 300))
					layoutRowWindow.Children = append(layoutRowWindow.Children, &ui.Base{
						Name:      layoutRowName,
						Placement: ui.MaximizeOffset(8, 44, 8, 8),
						Layout: &ui.LayoutRow{
							FullHeight:        false,
							VerticalAlignment: ui.AlignmentCenter,
							Spacing:           ui.Amount{Value: 10},
						},
						Children: []*ui.Base{
							newButton(ui.Absolute(0, 0, 0, 0), "Toggle Alignment", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow, p *ui.Base) {
									lr.VerticalAlignment = ui.Alignment(math.Mod(float64(lr.VerticalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "{s:20}{h:0.5}{pv:0.5}Toggle FullHeight", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow, p *ui.Base) {
									lr.FullHeight = !lr.FullHeight
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle EqualHeights", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow, p *ui.Base) {
									lr.EqualHeights = !lr.EqualHeights
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle FullWidth", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow, p *ui.Base) {
									lr.FullWidth = !lr.FullWidth
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle FullWidth Weights", false, func() {
								layoutRowChange(func(lc *ui.LayoutRow, p *ui.Base) {
									if p.Children[0].LayoutData.WidthWeight == 0 {
										p.Children[0].LayoutData.WidthWeight = 1
										p.Children[1].LayoutData.WidthWeight = 2
										p.Children[2].LayoutData.WidthWeight = 3
										p.Children[3].LayoutData.WidthWeight = 4
									} else {
										p.Children[0].LayoutData.WidthWeight = 0
										p.Children[1].LayoutData.WidthWeight = 0
										p.Children[2].LayoutData.WidthWeight = 0
										p.Children[3].LayoutData.WidthWeight = 0
									}
								})
							}),
							newButton(ui.Absolute(0, 0, 150, 60), "Toggle Spacing", false, func() {
								layoutRowChange(func(lr *ui.LayoutRow, p *ui.Base) {
									lr.Spacing.Value = float32(int(lr.Spacing.Value+5) % 15)
								})
							}),
							newButton(ui.Placement{}, "MinWidth", false, func() {
								layoutRowWindow.PlaceMinWidth()
							}),
							newButton(ui.Placement{}, "MinHeight", false, func() {
								layoutRowWindow.PlaceMinHeight()
							}),
						},
					})

					layoutGridName := id.Get("layoutGrid")
					layoutGridChange := func(change func(*ui.LayoutGrid)) {
						frame := userInterface.Named.Get(layoutGridName)
						layout := frame.Layout.(*ui.LayoutGrid)
						change(layout)
						frame.Relayout()
					}
					layoutGridWindow := newWindow("Layout Grid", ui.Absolute(1000, 300, 300, 300))
					layoutGridWindow.Children = append(layoutGridWindow.Children, &ui.Base{
						Name:      layoutGridName,
						Placement: ui.MaximizeOffset(8, 44, 8, 8),
						Layout: &ui.LayoutGrid{
							Rows: []ui.LayoutGridRow{{
								FullHeight:        false,
								VerticalAlignment: ui.AlignmentCenter,
							}},
							Columns: []ui.LayoutGridColumn{{
								FullWidth:           false,
								HorizontalAlignment: ui.AlignmentCenter,
							}},
							VerticalSpacing:   ui.Amount{Value: 10},
							HorizontalSpacing: ui.Amount{Value: 10},
							ColumnsMin:        3,
							AspectRatio:       0,
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
									lg.FullWidth = !lg.FullWidth
								})
							}),
							newButton(ui.Placement{}, "Toggle GridFullHeight", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.FullHeight = !lg.FullHeight
								})
							}),
							newButton(ui.Placement{}, "Toggle FullHeight", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.Rows[0].FullHeight = !lg.Rows[0].FullHeight
								})
							}),
							newButton(ui.Placement{}, "Toggle FullWidth", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.Columns[0].FullWidth = !lg.Columns[0].FullWidth
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
									lg.Rows[0].VerticalAlignment = ui.Alignment(math.Mod(float64(lg.Rows[0].VerticalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Placement{}, "Toggle Horizontal Alignment", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.Columns[0].HorizontalAlignment = ui.Alignment(math.Mod(float64(lg.Columns[0].HorizontalAlignment)+0.5, 1.5))
								})
							}),
							newButton(ui.Placement{}, "Toggle Columns", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									lg.ColumnsMin = (lg.ColumnsMin + 1) % 6
									lg.ColumnsDynamic = lg.ColumnsMin == 0
								})
							}),
							newButton(ui.Placement{}, "Toggle MinHeights", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									if lg.Rows[0].Min == 0 {
										lg.Rows[0].Min = 80
									} else if lg.Rows[0].Min == 80 {
										lg.Rows[0].Min = 160
									} else {
										lg.Rows[0].Min = 0
									}
								})
							}),
							newButton(ui.Placement{}, "Toggle MinWidths", false, func() {
								layoutGridChange(func(lg *ui.LayoutGrid) {
									if lg.Columns[0].Min == 0 {
										lg.Columns[0].Min = 100
									} else if lg.Columns[0].Min == 100 {
										lg.Columns[0].Min = 200
									} else {
										lg.Columns[0].Min = 0
									}
								})
							}),
							newButton(ui.Placement{}, "MinWidth", false, func() {
								layoutGridWindow.PlaceMinWidth()
							}),
							newButton(ui.Placement{}, "MinHeight", false, func() {
								layoutGridWindow.PlaceMinHeight()
							}),
							newButton(ui.Placement{}, "Hide & Show Animation", false, nil).Edit(func(b *ui.Base) {
								b.Colors.Set(BackgroundColor, color.Orange)
								b.Colors.Set(TextColor, color.Black)
								b.Animations.ForEvent.Set(ui.AnimationEventShow, ua.FadeIn)
								b.Animations.ForEvent.Set(ui.AnimationEventHide, ua.FadeOut)
								b.Events.OnPointer.Add(func(ev *ui.PointerEvent) {
									if !ev.Capture && ev.Type == ui.PointerEventDown {
										b.Hide()
										go func() {
											time.Sleep(time.Second * 3)
											b.Show()
										}()
									}
								}, false)
							}),
							newButton(ui.Placement{}, "Remove Instantly", false, nil).Edit(func(b *ui.Base) {
								b.Events.OnPointer.Add(func(ev *ui.PointerEvent) {
									if !ev.Capture && ev.Type == ui.PointerEventDown {
										b.Remove()
									}
								}, false)
							}),
							newButton(ui.Placement{}, "Remove Animating", false, nil).Edit(func(b *ui.Base) {
								b.Animations.ForEvent.Set(ui.AnimationEventRemove, ua.Explode.WithDuration(0.2))
								b.Events.OnPointer.Add(func(ev *ui.PointerEvent) {
									if !ev.Capture && ev.Type == ui.PointerEventDown {
										b.Remove()
									}
								}, false)
							}),
						},
					})

					layoutInlineName := id.Get("layoutInline")
					layoutInlineChange := func(change func(*ui.LayoutInline)) {
						frame := userInterface.Named.Get(layoutInlineName)
						layout := frame.Layout.(*ui.LayoutInline)
						change(layout)
						frame.Relayout()
					}
					layoutInlineWindow := newWindow("Layout Inline", ui.Absolute(20, 700, 300, 300))
					layoutInlineWindow.Children = append(layoutInlineWindow.Children,
						newCollapsibleSection("{w:none}Show/Hide Inline Components",
							&ui.Base{
								Name: layoutInlineName,
								Layout: &ui.LayoutInline{
									LineVerticalAlignment:   ui.AlignmentTop,
									LineHorizontalAlignment: ui.AlignmentLeft,
									VerticalSpacing:         ui.Amount{Value: 10},
									HorizontalSpacing:       ui.Amount{Value: 10},
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
									newButton(ui.Placement{}, "Toggle Line Vertical Alignment", false, func() {
										layoutInlineChange(func(lg *ui.LayoutInline) {
											lg.LineVerticalAlignment = ui.Alignment(math.Mod(float64(lg.LineVerticalAlignment)+0.5, 1.5))
										})
									}),
									newButton(ui.Placement{}, "Toggle Line Horizontal Alignment", false, func() {
										layoutInlineChange(func(lg *ui.LayoutInline) {
											lg.LineHorizontalAlignment = ui.Alignment(math.Mod(float64(lg.LineHorizontalAlignment)+0.5, 1.5))
										})
									}),
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
										b.MaxSize.X.Value = 120
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
									newButton(ui.Placement{}, "Toggle Line Full Height", false, func() {
										layoutInlineChange(func(lg *ui.LayoutInline) {
											lg.LineFullHeight = !lg.LineFullHeight
										})
									}),
									newButton(ui.Placement{}, "Toggle Line Full Width", false, func() {
										layoutInlineChange(func(lg *ui.LayoutInline) {
											lg.LineFullWidth = !lg.LineFullWidth
										})
									}),
									newButton(ui.Placement{}, "Toggle Equal Heights", false, func() {
										layoutInlineChange(func(lg *ui.LayoutInline) {
											lg.EqualHeights = !lg.EqualHeights
										})
									}),
									newButton(ui.Placement{}, "Toggle Line Equal Widths", false, func() {
										layoutInlineChange(func(lg *ui.LayoutInline) {
											lg.LineEqualWidths = !lg.LineEqualWidths
										})
									}),
									newButton(ui.Placement{}, "Toggle Full Height", false, func() {
										layoutInlineChange(func(lg *ui.LayoutInline) {
											lg.FullHeight = !lg.FullHeight
										})
									}),
									newButton(ui.Placement{}, "Toggle Full Width", false, func() {
										layoutInlineChange(func(lg *ui.LayoutInline) {
											lg.FullWidth = !lg.FullWidth
										})
									}),
									newButton(ui.Placement{}, "MinWidth", false, func() {
										layoutInlineWindow.PlaceMinWidth()
									}).Edit(func(b *ui.Base) {
										b.Colors.Set(BackgroundColor, color.Purple)
										b.Colors.Set(TextColor, color.White)
									}),
									newButton(ui.Placement{}, "MinHeight", false, func() {
										layoutInlineWindow.PlaceMinHeight()
									}),
								},
							},
						).Edit(func(b *ui.Base) {
							b.Placement = ui.MaximizeOffset(8, 44, 8, 8)
						}),
					)

					userInterface.Root = &ui.Base{
						Name: id.Get("root"),
						Layout: ui.LayoutStatic{
							EnforcePreferredSize: true,
							KeepInside:           true,
						},
						ChildrenOrderless: true,
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

// Colors

const (
	PrimaryColor color.Themed = iota
	SecondaryColor
	BackgroundColor
	TextColor
)

// Temporary component generators

func newScrollingSection(sensitivity float32, children ...*ui.Base) *ui.Base {
	section := &ui.Base{
		Placement: ui.Maximized(),
		Children:  children,
		Layout: ui.LayoutStatic{
			EnforcePreferredSize: true,
		},
	}

	var scrollArea *ui.Base
	scrollArea = &ui.Base{
		Placement:                   ui.Maximized(),
		IgnoreLayoutPreferredHeight: true,
		Clip:                        ui.Maximized(),
		Children:                    []*ui.Base{section},
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventWheel {
					size := section.PreferredSize(section.ComputeRenderContext(), scrollArea.Bounds.Width())
					dx := ev.Wheel.X * sensitivity
					dy := ev.Wheel.Y * sensitivity

					bounds := section.Bounds
					bounds.Translate(-scrollArea.Bounds.Left, -scrollArea.Bounds.Top)
					bounds.Bottom = bounds.Top + size.Y
					bounds.Right = bounds.Left + size.X
					bounds.Translate(dx, dy)
					overRight := scrollArea.Bounds.Width() - bounds.Right
					if overRight > 0 {
						bounds.Translate(overRight, 0)
					}
					overBottom := scrollArea.Bounds.Height() - bounds.Bottom
					if overBottom > 0 {
						bounds.Translate(0, overBottom)
					}
					if bounds.Left > 0 {
						bounds.Translate(-bounds.Left, 0)
					}
					if bounds.Top > 0 {
						bounds.Translate(0, -bounds.Top)
					}

					section.SetPlacement(ui.Absolute(bounds.Left, bounds.Top, size.X, size.Y))

					ev.Stop = true
				}
			},
		},
	}

	return scrollArea
}

func newCollapsibleSection(text string, children ...*ui.Base) *ui.Base {
	section := &ui.Base{
		Animations: &ui.Animations{
			ForEvent: ds.NewEnumMap(map[ui.AnimationEvent]ui.AnimationFactory{
				ui.AnimationEventShow: ua.RevealDown.Merge(ua.FadeIn).WithDuration(0.3),
				ui.AnimationEventHide: ua.RevealDown.Merge(ua.FadeIn).Reverse().WithDuration(0.3),
			}),
		},
		Children: children,
		LayoutData: ui.LayoutData{
			HeightWeight: 1,
		},
	}

	return &ui.Base{
		Layout: ui.LayoutColumn{
			Spacing:    ui.Amount{Value: 8},
			FullWidth:  true,
			FullHeight: true,
		},
		Children: []*ui.Base{
			{
				Layers: []ui.Layer{{
					Visual: ui.MustTextToVisual(text),
				}},
				Events: ui.Events{
					OnPointer: func(ev *ui.PointerEvent) {
						if !ev.Capture && ev.Type == ui.PointerEventDown {
							section.SetVisible(section.IsHidden())
							ev.Stop = true
						}
					},
				},
				Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
					ui.CursorEventHover: id.Get("click"),
					ui.CursorEventDown:  id.Get("clicking"),
				}),
			},
			section,
		},
	}
}

func newDropdown[T any](parent *ui.Base, items []T, toText func(T) string, onItem func(T)) *ui.Base {
	var dropdown *ui.Base
	var visible = false
	dropdown = &ui.Base{
		Name: id.Get("dropdown"),
		// centered below parent with bottom being -10px above parent's top
		Placement: ui.Placement{}.Relative(0, 0, 1, 0),
		// 0px below the parent
		RelativePlacement: ui.Placement{
			Left:   ui.Anchor{Delta: 0},
			Right:  ui.Anchor{Delta: 1},
			Top:    ui.Anchor{Delta: 1},
			Bottom: ui.Anchor{Delta: 1},
		},
		MaxSize: ui.AmountPoint{
			X: ui.Amount{Value: 1, Unit: ui.UnitParent},
			// Y: ui.Amount{Value: 200},
		},
		MinSize: ui.AmountPoint{
			Y: ui.Amount{Value: float32(util.Min(10, len(items))) * 24},
		},
		Shape: ui.ShapeRectangle{},
		TextStyles: &ui.TextStylesOverride{
			Color:    ui.Override[color.Able](color.Black),
			FontSize: &ui.Amount{Value: 16},
			ParagraphStylesOverride: &ui.ParagraphStylesOverride{
				ParagraphPadding: ui.Override(ui.NewAmountBoundsUniform(4, ui.UnitConstant)),
			},
			ParagraphsStylesOverride: &ui.ParagraphsStylesOverride{
				VerticalAlignment: ui.Override(ui.AlignmentCenter),
			},
		},
		Animations: &ui.Animations{
			ForEvent: ds.NewEnumMap(map[ui.AnimationEvent]ui.AnimationFactory{
				ui.AnimationEventShow: ua.RevealDown.WithDuration(0.2),
				ui.AnimationEventHide: ua.RevealDown.Reverse().WithDuration(0.2),
			}),
		},
		Layers: []ui.Layer{{
			Visual: ui.VisualShadow{
				Blur:    ui.NewAmountBoundsUniform(8, ui.UnitConstant),
				Offsets: ui.NewAmountBounds(2, 2, -2, -2),
			},
			Background: ui.BackgroundColor{Color: color.Black},
		}, {
			Visual:     ui.VisualFilled{},
			Background: ui.BackgroundColor{Color: color.White},
		}},
		Layout: ui.LayoutStatic{
			EnforcePreferredSize: true,
		},
		Children: []*ui.Base{
			newScrollingSection(100, &ui.Base{
				Name:   id.Get("dropdown-scroll-section"),
				Layout: ui.LayoutColumn{},
				Children: util.SliceMap(items, func(item T) *ui.Base {
					return &ui.Base{
						Placement: ui.Maximized().Shrink(4),
						Focusable: true,
						Layers: []ui.Layer{{
							Visual: ui.MustTextToVisual(toText(item)),
						}},
						Events: ui.Events{
							OnPointer: func(ev *ui.PointerEvent) {
								if !ev.Capture && ev.Type == ui.PointerEventDown {
									onItem(item)
									dropdown.Hide()
									visible = false
									ev.Stop = true
								}
							},
						},
					}
				}),
			}),
		},
	}

	parent.Events.OnPointer.Add(func(ev *ui.PointerEvent) {
		if !ev.Capture && ev.Type == ui.PointerEventDown {
			if dropdown.Parent() != parent {
				parent.AddChildren(dropdown)
				dropdown.SetRenderParent(parent.UI().Root)
			} else {
				if !visible {
					dropdown.Show()
				} else {
					dropdown.Hide()
				}

			}
			visible = !visible
			ev.Stop = true
		}
		if ev.Capture && ev.Type == ui.PointerEventDownOut && visible {
			dropdown.Hide()
			ev.Stop = true
		}
	}, true)

	return dropdown
}

func newTooltip(text string, delayTime float32, hideTime float32, around *ui.Base) *ui.Base {
	showAt := time.Time{}
	hideAt := time.Time{}

	tooltip := &ui.Base{
		// centered above parent with bottom being -10px above parent's top
		Placement: ui.Placement{}.Relative(0.5, 1, 0.5, 1),
		// -10px above the parent
		RelativePlacement: ui.Placement{
			Left:   ui.Anchor{Delta: 0.5},
			Right:  ui.Anchor{Delta: 0.5},
			Top:    ui.Anchor{Base: -10},
			Bottom: ui.Anchor{Base: -10},
		},
		MaxSize: ui.AmountPoint{
			X: ui.Amount{Value: 1, Unit: ui.UnitParent},
		},
		Shape: ui.ShapeRounded{
			Radius: ui.NewAmountCornersUniform(6, ui.UnitConstant),
		},
		TextStyles: &ui.TextStylesOverride{
			Color: ui.Override[color.Able](color.White),
			ParagraphStylesOverride: &ui.ParagraphStylesOverride{
				HorizontalAlignment: ui.Override(ui.AlignmentCenter),
				ParagraphPadding:    ui.Override(ui.NewAmountBoundsUniform(6, ui.UnitConstant)),
			},
			ParagraphsStylesOverride: &ui.ParagraphsStylesOverride{
				VerticalAlignment: ui.Override(ui.AlignmentCenter),
			},
		},
		Animations: &ui.Animations{
			ForEvent: ds.NewEnumMap(map[ui.AnimationEvent]ui.AnimationFactory{
				ui.AnimationEventShow:   ua.FadeIn,
				ui.AnimationEventRemove: ua.FadeOut,
			}),
		},
		Layers: []ui.Layer{{
			Visual: ui.VisualShadow{
				Blur:    ui.NewAmountBoundsUniform(6, ui.UnitConstant),
				Offsets: ui.NewAmountBounds(2, 2, -2, -2),
			},
			Background: ui.BackgroundColor{Color: color.White},
		}, {
			Visual:     ui.VisualFilled{},
			Background: ui.BackgroundColor{Color: color.Black},
		}, {
			Visual: ui.MustTextToVisual(text),
		}},
	}

	around.Events.OnPointer.Add(func(ev *ui.PointerEvent) {
		if ev.Capture {
			if ev.Type == ui.PointerEventEnter {
				showAt = time.Now().Add(time.Duration(float32(time.Second) * delayTime))
				hideAt = showAt.Add(time.Duration(float32(time.Second) * hideTime))
			} else if ev.Type == ui.PointerEventLeave {
				showAt = time.Time{}
				if tooltip.Parent() != nil {
					tooltip.Remove()
				}
			}
		}
	}, false)

	around.Hooks.OnUpdate.Add(func(b *ui.Base, update ui.Update) ui.Dirty {
		if tooltip.Parent() == nil && !showAt.IsZero() && time.Now().After(showAt) && !around.IsDisabled() {
			around.AddChildren(tooltip)
			tooltip.SetRenderParent(around.UI().Root)
		}
		if tooltip.Parent() != nil && hideAt != showAt && time.Now().After(hideAt) {
			tooltip.Remove()
			showAt = time.Time{}
			hideAt = time.Time{}
		}
		return ui.DirtyNone
	}, false)

	return tooltip
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
		Colors: color.NewColors(map[color.Themed]color.Able{
			BackgroundColor: color.Black,
		}),
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
			Background: ui.BackgroundColor{Color: BackgroundColor},
			States:     ui.StateHover.Not,
		}, {
			Visual:     ui.VisualFilled{Shape: ui.ShapePolygon{Points: shape}},
			Background: ui.BackgroundColor{Color: BackgroundColor.Modify(color.Lighten(0.3))},
			States:     ui.StateHover.Is,
		}, {
			Visual: ui.MustTextToVisual("{f:roboto}{s:14}{c:white}{h:0.5}{pv:0.5}drag me"),
		}},
	}

	return draggable
}

var buttonTemplate = &ui.Template{
	Animations: &ui.Animations{
		ForEvent: ds.NewEnumMap(map[ui.AnimationEvent]ui.AnimationFactory{
			ui.AnimationEventShow: ua.FadeIn,
		}),
	},
	Colors: color.NewColors(map[color.Themed]color.Able{
		BackgroundColor: PrimaryColor,
	}),
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
		Visual: ui.VisualShadow{
			Blur:    ui.NewAmountBoundsUniform(6, ui.UnitConstant),
			Offsets: ui.NewAmountBounds(5, 8, -3, 0),
		},
		Background: ui.BackgroundColor{Color: BackgroundColor.Modify(color.Darken(0.5).Then(color.Alpha(0.2)))},
		States:     (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Not,
	}, {
		Visual: ui.VisualShadow{
			Blur:    ui.NewAmountBoundsUniform(6, ui.UnitConstant),
			Offsets: ui.NewAmountBounds(5, 8, -3, 0),
		},
		Background: ui.BackgroundColor{Color: BackgroundColor.Modify(color.Darken(0.5).Then(color.Alpha(0.5)))},
		States:     (ui.StateHover | ui.StatePressed | ui.StateFocused | ui.StateSelected).Is,
	}, {
		// Background
		Placement:  ui.Maximized(),
		Visual:     ui.VisualFilled{},
		Background: ui.BackgroundColor{Color: BackgroundColor},
		States:     (ui.StateHover | ui.StatePressed).Not,
	}, {
		// Background on hover
		Placement:  ui.Maximized(),
		Visual:     ui.VisualFilled{},
		Background: ui.BackgroundColor{Color: BackgroundColor.Modify(color.Lighten(0.1))},
		States:     ui.StateHover.Is,
	}, {
		// Background on press
		Placement:  ui.Maximized(),
		Visual:     ui.VisualFilled{},
		Background: ui.BackgroundColor{Color: BackgroundColor.Modify(color.Darken(0.1))},
		States:     ui.StatePressed.Is,
	}},
	PostLayers: []ui.Layer{{
		// Ripple animation
		Placement: ui.Maximized().Shrink(2),
		Visual: &RippleLayer{
			StartRadius: ui.Amount{Value: 0},
			EndRadius:   ui.Amount{Value: 4, Unit: ui.UnitParent},
			StartColor:  color.New(1, 1, 1, 0.3),
			EndColor:    color.Transparent,
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
			Placement:  ui.Maximized().Shrink(10),
			Visual:     textVisual,
			Background: ui.BackgroundColor{Color: TextColor},
		}},
	}

	button.ApplyTemplate(buttonTemplate)

	if pulse {
		button.Layers = append([]ui.Layer{{
			Visual: &PulseLayer{
				StartColor: BackgroundColor.Modify(color.Lighten(0.2)),
				EndColor:   color.Transparent,
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
	StartColor, EndColor   color.Able
	Duration               float32
	Time                   float32
	Center                 ui.Coord
	Animating              bool

	animatingOn *ui.Base
}

func (r *RippleLayer) Init(b *ui.Base) {
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
		radius := util.Lerp(r.StartRadius.Get(ctx.AmountContext, true), r.EndRadius.Get(ctx.AmountContext, true), delta)
		startColor := r.StartColor.GetColor(b)
		endColor := r.EndColor.GetColor(b)
		color := startColor.Lerp(endColor, delta)

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
		visualized := ui.NewVertexIterator(out, false)
		background.Visualize(b, rippleBounds, ctx, out)
		for visualized.HasNext() {
			v := visualized.Next()
			v.X = util.Clamp(v.X, bounds.Left, bounds.Right)
			v.Y = util.Clamp(v.Y, bounds.Top, bounds.Bottom)
			v.AddColor(color)
		}
	}
}
func (r *RippleLayer) PreferredSize(b *ui.Base, ctx *ui.RenderContext, maxWidth float32) ui.Coord {
	return ui.Coord{}
}

type PulseLayer struct {
	StartColor, EndColor color.Able
	Duration             float32
	PulseTime            float32
	Size                 float32
	Time                 float32
	Shape                ui.Shape
	Disabled             bool
	Times                int
}

func (r *PulseLayer) Start() {
	r.Time = 0
	r.Disabled = false
}

func (r *PulseLayer) Stop() {
	r.Disabled = true
}

func (r *PulseLayer) Init(b *ui.Base) {}
func (r *PulseLayer) Update(b *ui.Base, update ui.Update) ui.Dirty {
	if r.Disabled {
		return ui.DirtyNone
	}
	r.Time += float32(update.DeltaTime.Seconds())
	if r.Time > r.Duration {
		r.Time -= r.Duration
		r.Times--
		if r.Times == 0 {
			r.Disabled = true
		}
	}
	if r.Time < r.PulseTime {
		return ui.DirtyVisual
	}
	return ui.DirtyNone
}
func (r *PulseLayer) Visualize(b *ui.Base, bounds ui.Bounds, ctx *ui.RenderContext, out *ui.VertexBuffers) {
	if !r.Disabled && r.Time <= r.PulseTime {
		delta := r.Time / r.PulseTime
		size := util.Lerp(0, r.Size, delta)
		pulseBounds := ui.Bounds{
			Left:   bounds.Left - size,
			Right:  bounds.Right + size,
			Top:    bounds.Top - size,
			Bottom: bounds.Bottom + size,
		}
		startColor := r.StartColor.GetColor(b)
		endColor := r.EndColor.GetColor(b)
		color := startColor.Lerp(endColor, delta)
		background := ui.VisualFilled{
			Shape: r.Shape,
		}
		if background.Shape == nil {
			background.Shape = b.Shape
		}
		visualized := ui.NewVertexIterator(out, false)
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

	activePulse := &PulseLayer{
		StartColor: PrimaryColor.Modify(color.Darken(0.5)),
		EndColor:   PrimaryColor.Modify(color.Darken(0.5).Then(color.Alpha(0))),
		Duration:   0.3,
		PulseTime:  0.3,
		Size:       10,
		Disabled:   true,
	}

	frame = &ui.Base{
		Placement: placement,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if ev.Capture && ev.Type == ui.PointerEventDown {
					frame.BringToFront()
				}
			},
			OnFocus: func(ev *ui.Event) {
				activePulse.Times = 1
				activePulse.Start()
			},
			OnBlur: func(ev *ui.Event) {
				activePulse.Stop()
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
		Focusable: true,
		Animations: &ui.Animations{
			ForEvent: ds.NewEnumMap(map[ui.AnimationEvent]ui.AnimationFactory{
				ui.AnimationEventShow: ua.FadeInDown,
			}),
			Named: id.NewDenseKeyMap[ui.AnimationFactory, uint16, uint8](
				id.WithStringMap(map[string]ui.AnimationFactory{
					"hide": ua.FadeOutUp,
					"show": ua.FadeInDown,
				}),
			),
		},
		Colors: color.NewColors(map[color.Themed]color.Able{
			BackgroundColor: color.LightGray,
		}),
		Layers: []ui.Layer{{
			Visual: activePulse,
			States: ui.StateFocused.Is,
		}, {
			Visual: ui.VisualShadow{
				Blur:    ui.NewAmountBoundsUniform(8, ui.UnitConstant),
				Offsets: ui.NewAmountBounds(5, 8, -3, 0),
			},
			Background: ui.BackgroundColor{Color: PrimaryColor.Modify(color.Darken(0.5).Then(color.Alpha(0.2)))},
		}, {
			Visual:     ui.VisualFilled{},
			Background: ui.BackgroundColor{Color: BackgroundColor},
		}},
	}

	bar := &ui.Base{
		Placement: ui.Placement{}.TopFixedHeight(0, barSize, 0, 0),
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
				StartColor: PrimaryColor,
				EndColor:   PrimaryColor.Modify(color.Lighten(0.2)),
				End:        ui.Coord{X: 0, Y: 1},
			},
		}},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("drag"),
			ui.CursorEventDrag:  id.Get("dragging"),
		}),
		Draggable: true,
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if ev.Capture {
					return
				}
				switch ev.Type {
				case ui.DragEventStart:
					frame.SetTransparency(0.2)
					frame.BringToFront()
				case ui.DragEventMove:
					frame.SetPlacement(frame.Placement.Shift(ev.DeltaMove.X, ev.DeltaMove.Y))
				case ui.DragEventEnd:
					frame.SetTransparency(0)
				}
			},
		},
		Layout: ui.LayoutRow{
			FullWidth:         true,
			VerticalAlignment: ui.AlignmentCenter,
		},
		Children: []*ui.Base{
			{
				Layers: []ui.Layer{{
					Placement: ui.Maximized().Shrink(2).Shift(6, 0),
					Visual:    ui.MustTextToVisual("{w:none}{s:20}{pv:0.5}{k:-1}" + title),
				}},
				LayoutData: ui.LayoutData{
					WidthWeight: 1,
				},
			},
			newWindowHide(frame, barSize),
			newWindowMinimizeMaximize(frame, barSize),
			newWindowClose(frame, barSize),
		},
	}

	frame.Children = append(
		frame.Children, bar,
		newWindowResizer(frame, id.Get("resizer"), ui.Placement{}.RightFixedWidth(0, ResizerSize, barSize, ResizerSize), ui.Bounds{Right: 1}),
		newWindowResizer(frame, id.Get("resizebr"), ui.Placement{}.Attach(1, 1, ResizerSize, ResizerSize), ui.Bounds{Right: 1, Bottom: 1}),
		newWindowResizer(frame, id.Get("resizeb"), ui.Placement{}.BottomFixedHeight(0, ResizerSize, ResizerSize, ResizerSize), ui.Bounds{Bottom: 1}),
		newWindowResizer(frame, id.Get("resizebl"), ui.Placement{}.Attach(0, 1, ResizerSize, ResizerSize), ui.Bounds{Left: 1, Bottom: 1}),
		newWindowResizer(frame, id.Get("resizel"), ui.Placement{}.LeftFixedWidth(0, ResizerSize, barSize, ResizerSize), ui.Bounds{Left: 1}),
		// newWindowResizeRight(frame, barSize),
		// newWindowResizeBottom(frame),
		// newWindowResizeBottomRight(frame),
		// newWindowResizeBottomLeft(frame),
	)

	return frame
}

func newWindowClose(win *ui.Base, barSize float32) *ui.Base {
	return &ui.Base{
		MinSize: ui.NewAmountPoint(barSize, barSize),
		Layers: []ui.Layer{{
			Background: ui.BackgroundColor{Color: color.LightGray.Alpha(0.3)},
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
			Background: ui.BackgroundColor{Color: color.Black},
		}},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("click"),
			ui.CursorEventDown:  id.Get("clicking"),
		}),
		Draggable: true,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventUp {
					win.SetTransparency(1)

					go func() {
						time.Sleep(time.Second * 3)
						win.SetTransparency(0)
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
		MinSize: ui.NewAmountPoint(barSize, barSize),
		Layers: []ui.Layer{{
			Background: ui.BackgroundColor{Color: color.LightGray.Alpha(0.3)},
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
			Background: ui.BackgroundColor{Color: color.Black},
		}},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("click"),
			ui.CursorEventDown:  id.Get("clicking"),
		}),
		Draggable: true,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventUp {
					if maximized {
						win.Play(ui.PlacementAnimation{
							Target:   minimized,
							Duration: 0.3,
						})
					} else {
						minimized = win.Placement
						win.Play(ui.PlacementAnimation{
							Target:   ui.Maximized(),
							Duration: 0.3,
						})
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
	hiding := false
	hidingPlacement := ui.Placement{}

	return &ui.Base{
		MinSize: ui.NewAmountPoint(barSize, barSize),
		Layers: []ui.Layer{{
			Background: ui.BackgroundColor{Color: color.LightGray.Alpha(0.3)},
			Visual:     ui.VisualFilled{Shape: ui.ShapeRectangle{}},
			States:     ui.StateHover.Is,
		}, {
			Placement: ui.Absolute(8, barSize-10, barSize-14, 3),
			Visual: ui.VisualFilled{
				Shape: ui.ShapeRectangle{},
			},
			Background: ui.BackgroundColor{Color: color.Black},
		}},
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: id.Get("click"),
			ui.CursorEventDown:  id.Get("clicking"),
		}),
		Draggable: true,
		Events: ui.Events{
			OnPointer: func(ev *ui.PointerEvent) {
				if !ev.Capture && ev.Type == ui.PointerEventUp {
					if hiding {
						win.Play(ui.PlacementAnimation{
							Target:       hidingPlacement,
							Duration:     0.3,
							ApplyMaxSize: true,
						})
					} else {
						hidingPlacement = win.Placement
						win.Play(ui.PlacementAnimation{
							Target: ui.Placement{
								Top:    ui.Anchor{Delta: 1},
								Bottom: ui.Anchor{Delta: 1},
							},
							Duration:     0.3,
							ApplyMaxSize: true,
						})
					}
					for i := range win.Children {
						if i > 0 {
							win.Children[i].SetVisible(hiding)
						}
					}
					hiding = !hiding
					// win.PlayMaybe("hide")
					// go func() {
					// 	time.Sleep(time.Second * 3)
					// 	win.PlayMaybe("show")
					// }()
				}
			},
			OnDrag: func(ev *ui.DragEvent) {
				ev.Cancel = true
				ev.Stop = true
			},
		},
	}
}

const ResizerSize = 12

func newWindowResizer(win *ui.Base, cursor id.Identifier, placement ui.Placement, dirs ui.Bounds) *ui.Base {
	start := win.Placement
	var resizer *ui.Base

	move := func(base *float32, dir, available, move float32, towardsEdge, onlyIf bool) {
		if onlyIf && dir != 0 && move != 0 {
			if towardsEdge {
				*base += util.MaxMagnitude(move, available)
			} else {
				*base += move
			}
		}
	}

	resizer = &ui.Base{
		Draggable: true,
		Placement: placement,
		Cursors: ui.NewCursors(map[ui.CursorEvent]id.Identifier{
			ui.CursorEventHover: cursor,
		}),
		Events: ui.Events{
			OnDrag: func(ev *ui.DragEvent) {
				if !ev.Capture && !win.Placement.IsMaximized() {
					ev.Stop = true
					switch ev.Type {
					case ui.DragEventStart:
						win.SetTransparency(0.2)
						start = win.Placement
					case ui.DragEventMove:
						current := win.Placement
						winParentBounds := win.Parent().Bounds
						win.Bounds = win.Placement.GetBoundsIn(winParentBounds)
						available := winParentBounds.Sub(win.Bounds)

						move(&current.Left.Base, dirs.Left, -available.Left, ev.DeltaMove.X, ev.DeltaMove.X < 0, (ev.Point.X < resizer.Bounds.Left-ev.DeltaMove.X) || ev.DeltaMove.X > 0)
						move(&current.Top.Base, dirs.Top, -available.Top, ev.DeltaMove.Y, ev.DeltaMove.Y < 0, (ev.Point.Y < resizer.Bounds.Top-ev.DeltaMove.Y) || ev.DeltaMove.Y > 0)
						move(&current.Right.Base, dirs.Right, available.Right, ev.DeltaMove.X, ev.DeltaMove.X > 0, (ev.Point.X > resizer.Bounds.Right-ev.DeltaMove.X) || ev.DeltaMove.X < 0)
						move(&current.Bottom.Base, dirs.Bottom, available.Bottom, ev.DeltaMove.Y, ev.DeltaMove.Y > 0, (ev.Point.Y > resizer.Bounds.Bottom-ev.DeltaMove.Y) || ev.DeltaMove.Y < 0)

						win.SetPlacement(current)
					case ui.DragEventCancel:
						win.SetPlacement(start)
					case ui.DragEventEnd:
						win.SetTransparency(0)
					}
				}
			},
		},
	}

	return resizer
}
