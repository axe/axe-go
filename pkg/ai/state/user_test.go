package state_test

import (
	"fmt"
	"testing"

	"github.com/axe/axe-go/pkg/ai/state"
)

func TestUser(t *testing.T) {
	var (
		OnGround = state.Bool(false)
		Forward  = state.Float(1.4)
	)

	data := state.NewUserData(OnGround, Forward)

	fmt.Printf("OnGround: %+v, Forward: %+v\n", OnGround.Get(data), Forward.Get(data))

	OnGround.Set(data, true)
	Forward.Set(data, 23.4)

	fmt.Printf("OnGround: %+v, Forward: %+v\n", OnGround.Get(data), Forward.Get(data))
}
