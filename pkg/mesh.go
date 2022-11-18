package axe

/*
Formats:
	obj: https://github.com/udhos/gwob/blob/master/obj.go
	stl: https://github.com/neilpa/go-stl/blob/master/binary.go
	fbx: https://github.com/bqqbarbhg/ufbx/blob/master/ufbx.h

*/

type MeshData struct {
	Name      string
	Vertices  [][3]float32
	Normals   [][3]float32
	Uvs       [][2]float32
	Groups    []MeshFaceGroup
	Materials string
}

type MeshFaceGroup struct {
	Material string
	Faces    []Face
}

type Face struct {
	Vertices []int
	Uvs      []int
	Normals  []int
}

type Material struct {
	Name        string       // Material name
	Illum       int          // Illumination model
	Opacity     float32      // Opacity factor
	Refraction  float32      // Refraction factor
	Shininess   float32      // Shininess (specular exponent)
	Ambient     TextureColor // Ambient color reflectivity
	Diffuse     TextureColor // Diffuse color reflectivity
	Specular    TextureColor // Specular color reflectivity
	Emissive    TextureColor // Emissive color
	TextureBump string
}

type TextureColor struct {
	Texture string
	Color   [3]float32
}

type Materials map[string]Material
