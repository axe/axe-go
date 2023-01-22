package axe

import (
	"github.com/axe/axe-go/pkg/input"
)

type InputSystem interface {
	GameSystem
	input.InputSystem
}
