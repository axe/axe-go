package axe

import "github.com/axe/axe-go/pkg/ecs"

type Tag string

var TAG = ecs.DefineComponent("tag", Tag(""))
