package axe

import (
	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/ecs"
)

var MESH = ecs.DefineComponent("mesh", Mesh{})

type Mesh struct {
	Ref asset.Ref

	data *MeshData
}
