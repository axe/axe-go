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
			Placement: ui.Centered(720, 480),
			Mode:      axe.WindowModeFixed,
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

					userInterface := axe.NewUserInterface()
					userInterface.Root = &ui.Base{
						Layers: []ui.Layer{{
							Visual: &ui.VisualText{
								Glyphs: simpleTextGlyphs("Hello world!", "roboto", 32, ui.ColorWhite),
							},
						}},
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

func simpleTextGlyphs(text string, font string, size float32, color ui.Color) ui.GlyphBlocks {
	gs := make([]ui.Glyph, 0, len(text))
	for _, c := range text {
		gs = append(gs, &ui.TextGlyph{
			Text:  c,
			Font:  font,
			Size:  ui.Amount{Value: size},
			Color: color,
		})
	}
	gbs := ui.GlyphBlocks{
		ClampLeft:         true,
		ClampTop:          true,
		VerticalAlignment: 0.5,
		Blocks: []ui.GlyphBlock{{
			HorizontalAlignment: 0.5,
			Wrap:                ui.TextWrapNone,
			Glyphs:              gs,
		}},
	}
	return gbs
}
