package input

import (
	"time"

	"github.com/axe/axe-go/pkg/core"
)

type EmptySystem struct{}

var _ InputSystem = &EmptySystem{}

func (in *EmptySystem) Devices() []*Device                    { return nil }
func (in *EmptySystem) Inputs() []*Input                      { return nil }
func (in *EmptySystem) InputTime() time.Time                  { return time.Time{} }
func (in *EmptySystem) Get(name string) *Input                { return nil }
func (in *EmptySystem) Points() []*Point                      { return nil }
func (in *EmptySystem) Events() *core.Listeners[SystemEvents] { return nil }
