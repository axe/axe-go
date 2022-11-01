package glfw

import (
	"fmt"
	"strings"
	"testing"
	"time"

	axe "github.com/axe/axe-go/pkg"
)

// type emptyAudioSystem struct {
// 	updates int
// 	elapsed time.Duration
// }

// var _ axe.AudioSystem = &emptyAudioSystem{}

// func (audio *emptyAudioSystem) Init(game *axe.Game) error { return nil }
// func (audio *emptyAudioSystem) Update(game *axe.Game) {
// 	audio.updates++
// 	audio.elapsed += game.State.UpdateTimer.Elapsed

// 	if audio.updates >= 1000 {
// 		fps := float64(audio.updates) / audio.elapsed.Seconds()
// 		fmt.Printf("FPS: %.1f\n", fps)
// 		audio.updates = 0
// 		audio.elapsed = 0
// 	}
// }
// func (audio *emptyAudioSystem) Destroy() {}

func TestGame(t *testing.T) {
	game := axe.NewGame(axe.GameSettings{
		EnableDebug:        true,
		Name:               "Test GLFW",
		FixedDrawFrequency: time.Second / 60,
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

	// view := axe.View[float32, axe.Vec3[float32]]{}

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
