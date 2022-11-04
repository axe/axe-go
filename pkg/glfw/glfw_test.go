package glfw

import (
	"fmt"
	"strings"
	"testing"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/ui"
)

func TestGame(t *testing.T) {
	game := axe.NewGame(axe.GameSettings{
		EnableDebug:        true,
		Name:               "Test GLFW",
		FixedDrawFrequency: time.Second / 60,
		Stages: []axe.Stage{{
			Name:   "first",
			Assets: []axe.AssetRef{{Name: "cube", URI: "./square.png"}},
			Windows: []axe.StageWindow{{
				Name:       "main",
				Title:      "Test GLFW Main Window",
				Placement:  ui.Maximized(),
				Fullscreen: false,
			}},
			// Scenes: []axe.Scene[float32, axe.Vec2[float32]]{{
			// 	Name: "main",
			// 	World: axe.NewWorld(axe.WorldOptions{
			// 		EntityMax:              128,
			// 		StageSize:              16,
			// 		UncommonComponentCount: 4,
			// 	}),
			// 	Load: func(scene *axe.Scene[float32, axe.Vec2[float32]], game *axe.Game) {
			// 		type Sprite struct {
			// 			Texture axe.Texture
			// 		}

			// 		TRANSFORM := axe.DefineComponent("transform", axe.Transform{}, true)
			// 		SPRITE := axe.DefineComponent("sprite", Sprite{}, true)

			// 		transforms := TRANSFORM.Enable(scene.World)
			// 		TRANSFORM.AddSystem(scene.World, axe.NewTransformSystem)

			// 		sprites := SPRITE.Enable(scene.World)
			// 		SPRITE.AddSystem(scene.World, nil)

			// 		e := scene.World.Create()
			// 		transforms.Set(e, axe.Transform{Local: 23})
			// 		sprites.Take(e).Texture = game.Assets.Get("cube").Data.(axe.Texture)
			// 	},
			// }},
		}},
	})
	Setup(game)

	itb := axe.InputTriggerBuilder{}

	actionSet := axe.NewInputActionSet("main")
	actionSet.Add(axe.NewInputAction("close", itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyEscape})))
	actionSet.Add(axe.NewInputAction("down", itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyZ})))
	actionSet.Add(axe.NewInputAction("undo", itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyZ, CmdCtrl: true})))
	actionSet.Add(axe.NewInputAction("pasteUp", itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyV, Ctrl: true, UpOnly: true})))
	actionSet.Add(axe.NewInputAction("pressA", itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyA, PressInterval: time.Second / 4, FirstPressDelay: time.Second})))
	actionSet.Add(axe.NewInputAction("logInput", itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyC})))

	logInput := false

	game.Actions.Add(actionSet)
	game.Actions.Handler = func(action *axe.InputAction) {
		switch action.Name {
		case "close":
			game.Running = false
		case "logInput":
			logInput = !logInput
		default:
			inputNames := []string{}
			if action.Data.Inputs != nil {
				for _, in := range action.Data.Inputs {
					inputNames = append(inputNames, in.Name)
				}
			}
			fmt.Printf("%s %0.1f (priority=%d, inputs=%s)\n", action.Name, action.Data.Value, action.Priority(), strings.Join(inputNames, ","))
		}
	}

	game.Input.Events().On(axe.InputSystemEvents{
		InputChange: func(input axe.Input) {
			if logInput {
				fmt.Printf("%s changed to %v\n", input.Name, input.Value)
			}
		},
	})

	err := game.Run()
	if err != nil {
		panic(err)
	}
}
