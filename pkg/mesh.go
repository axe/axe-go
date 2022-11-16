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
	Faces     []Face
	Materials []Material
}

type Face struct {
	Vertices []int
	Uvs      []int
	Normals  []int
	Material string
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
