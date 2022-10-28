package glfw

import (
	"fmt"
	"testing"
	"time"

	axe "github.com/axe/axe-go/pkg"
)

type emptyAudioSystem struct {
}

var _ axe.AudioSystem = &emptyAudioSystem{}

func (audio *emptyAudioSystem) Init(game *axe.Game) error { return nil }
func (audio *emptyAudioSystem) Update(game *axe.Game)     {}
func (audio *emptyAudioSystem) Destroy()                  {}

func TestGame(t *testing.T) {
	game := &axe.Game{
		Graphics: NewGraphicsSystem(),
		Input:    NewInputSystem(),
		Windows:  NewWindowSystem(),
		Audio:    &emptyAudioSystem{},
	}

	game.State.UpdateTimer.Frequency = time.Millisecond * 20
	game.State.DrawTimer.Frequency = time.Millisecond * 10

	game.Input.Events().On(axe.InputSystemEvents{
		InputChange: func(input axe.Input) {
			data := input.Data()
			fmt.Printf("%s changed to %v\n", data.Name, data.Value)
		},
		InputChangeMap: map[string]func(input axe.Input){
			"escape": func(input axe.Input) {
				game.Running = false
			},
		},
	})

	game.Run()
}
