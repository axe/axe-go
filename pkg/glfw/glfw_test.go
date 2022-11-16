package glfw

import (
	"fmt"
	"strings"
	"testing"
	"time"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/job"
	"github.com/axe/axe-go/pkg/ui"
)

func TestGame(t *testing.T) {
	game := axe.NewGame(axe.GameSettings{
		EnableDebug:        true,
		Name:               "Test GLFW",
		FixedDrawFrequency: time.Second / 60,
		Stages: []axe.Stage{{
			Name: "first",
			Assets: []axe.AssetRef{
				{Name: "cube texture", URI: "./square.png"},
				// {Name: "cube model", URI: "./cube.obj"}, // OBJ not supported yet for MeshData
			},
			Windows: []axe.StageWindow{{
				Name:       "main",
				Title:      "Test GLFW Main Window",
				Placement:  ui.Maximized(),
				Fullscreen: false,
			}},
			Scenes3: []axe.Scene3f{{
				Name: "main",
				World: axe.NewWorld("main", axe.WorldSettings{
					EntityCapacity:            2048,
					EntityStageCapacity:       128,
					AverageComponentPerEntity: 4,
					DeleteOnDestroy:           true,
				}),
				Jobs: job.NewRunner(12, 1000),
				Load: func(scene *axe.Scene3f, game *axe.Game) {
					scene.World.Enable(
						// Component data settings
						axe.EntityDataSettings{Capacity: 2048, StageCapacity: 128},
						// Components
						axe.TAG, axe.MESH, axe.TRANSFORM3, axe.AUDIO,
					)

					// Systems
					axe.TRANSFORM3.AddSystem(axe.NewTransformSystem[axe.Vec4f]())

					// Entities
					e := axe.NewEntity()
					axe.TAG.Set(e, axe.Tag("cube"))
					axe.MESH.Set(e, axe.Mesh{Name: "cube model"})
					axe.TRANSFORM3.Set(e, axe.NewTransform4(axe.TransformCreate4f{
						Position: axe.Vec4f{0, 0, 0, 0},
						Scale:    axe.Vec4f{1, 1, 1, 0},
					}))
				},
			}},
		}},
	})
	Setup(game)

	// WIP below

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
