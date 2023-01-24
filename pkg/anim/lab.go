package anim

import "github.com/axe/axe-go/pkg/core"

type AnimationObjects[T core.Attr[T]] struct {
	GetTween func() Tween[T]
}
