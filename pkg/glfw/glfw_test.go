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

	logInput := false
	itb := axe.InputTriggerBuilder{}

	game := axe.NewGame(axe.GameSettings{
		EnableDebug:        true,
		Name:               "Test GLFW",
		FixedDrawFrequency: time.Second / 60,
		FirstStage:         "cube",
		WorldSettings: axe.WorldSettings{
			EntityCapacity:            2048,
			EntityStageCapacity:       128,
			AverageComponentPerEntity: 4,
			DeleteOnDestroy:           true,
		},
		Windows: []axe.StageWindow{{
			Title:     "Test GLFW Main Window",
			Placement: ui.Centered(720, 480),
		}},
		Stages: []axe.Stage{{
			Name: "cube",
			Assets: []axe.AssetRef{
				{Name: "cube model", URI: "cube.obj"},
			},
			Actions: axe.CreateInputActionSets(map[string]map[string]axe.InputTrigger{
				"main": {
					"close":    itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyEscape}),
					"down":     itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyZ}),
					"undo":     itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyZ, CmdCtrl: true}),
					"pasteUp":  itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyV, Ctrl: true, UpOnly: true}),
					"pressA":   itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyA, PressInterval: time.Second / 4, FirstPressDelay: time.Second}),
					"logInput": itb.Key(axe.InputKeyTrigger{Key: axe.InputKeyC}),
				},
			}),
			Views3: []axe.View3f{},
			Scenes3: []axe.Scene3f{{
				Load: func(scene *axe.Scene3f, game *axe.Game) {

					scene.World.Enable(
						// Component data settings
						axe.EntityDataSettings{Capacity: 2048, StageCapacity: 128},
						// Components
						axe.TAG, axe.MESH, axe.TRANSFORM3, axe.AUDIO, axe.INPUTACTION, axe.LIGHT, axe.LOGIC,
					)

					// Systems
					axe.TRANSFORM3.AddSystem(axe.NewTransformSystem[axe.Vec4f]())
					axe.INPUTACTION.AddSystem(axe.NewInputActionSystem(nil))
					axe.LOGIC.AddSystem(axe.NewLogicSystem())

					// Entities
					e := axe.NewEntity()
					axe.TAG.Set(e, axe.Tag("cube"))
					axe.MESH.Set(e, axe.Mesh{Ref: axe.AssetRef{Name: "cube model"}})
					axe.TRANSFORM3.Set(e, axe.NewTransform4(axe.TransformCreate4f{
						Position: axe.Vec4f{X: 0, Y: 0, Z: -3, W: 0},
						Scale:    axe.Vec4f{X: 1, Y: 1, Z: 1, W: 0},
					}))
					axe.LOGIC.Set(e, func(e *axe.Entity, ctx axe.EntityContext) {
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
					axe.INPUTACTION.Set(e, axe.InputActionListener{
						Handler: func(action *axe.InputAction) bool {
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
							return true
						},
					})
				},
			}},
		}},
	})
	Setup(game)

	// WIP below

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
