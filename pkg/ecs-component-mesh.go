package axe

var MESH = DefineComponent("mesh", Mesh{})

type Mesh struct {
	Name string

	mesh *MeshData
}
