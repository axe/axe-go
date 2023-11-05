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
		Name:               "Test Text",
		FixedDrawFrequency: time.Second / 60,
		FirstStage:         "text",
		WorldSettings: ecs.WorldSettings{
			EntityCapacity:            2048,
			EntityStageCapacity:       128,
			AverageComponentPerEntity: 4,
			DeleteOnDestroy:           true,
		},
		Windows: []axe.StageWindow{{
			Title:     "Test Text",
			Placement: ui.Centered(720, 600),
		}},
		Stages: []axe.Stage{{
			Name: "text",
			Assets: []asset.Ref{
				{Name: "sans-serif", URI: "../assets/sans-serif.fnt"},
				{Name: "warrior", URI: "../assets/warrior.fnt"},
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

					lines := []string{
						"Dear Reader,",
						"{p}{h:0.5}This is centered.",
						"{p}{v:0.5}And {s:2f}{f:warrior}THIS{s}{f} is big!",
						"{p}{v:1}This is bottom & center {s:2f}aligned?",
						"{p}{h:0}{v:0}Top{s} and left aligned.",
						"{p}{h:0.5}{c:red}And {c:orange}this {c:yellow}line {c:green}is {c:blue}super {c:indigo}duper {c:violet}gay!",
						"{p}{h:1}{c:white}Right aligned!",
						"{p}{h:0.25}25% aligned?",
						"{p}{h}{w:word}{i:2rem}This should wrap at the word and not at the character and should take up at least two lines with an indent. Resize the window!",
						"{p}{i}{pt:20}{h:0.5}{w:char}This should wrap at the character and not at the word and be centered.",
					}

					userInterface := axe.NewUserInterface()
					userInterface.Theme.TextStyles.Font = "roboto"
					userInterface.Root = &ui.Base{
						TextStyles: &ui.TextStylesOverride{
							Color:    &ui.ColorWhite,
							FontSize: &ui.Amount{Value: 24},
							ParagraphStylesOverride: &ui.ParagraphStylesOverride{
								LineSpacing:           &ui.Amount{Value: 1.0, Unit: ui.UnitFont},
								LineVerticalAlignment: ui.Override(ui.AlignmentMiddle),
							},
							ParagraphsStylesOverride: &ui.ParagraphsStylesOverride{
								ParagraphSpacing: &ui.Amount{Value: 1.0, Unit: ui.UnitFont},
							},
						},
						Layers: []ui.Layer{{
							Placement: ui.Maximized().Shrink(10),
							Visual:    ui.MustTextToVisual(strings.Join(lines, "\n")),
						}},
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
