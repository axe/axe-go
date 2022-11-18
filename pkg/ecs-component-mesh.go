package axe

var MESH = DefineComponent("mesh", Mesh{})

type Mesh struct {
	Ref AssetRef

	data *MeshData
}
