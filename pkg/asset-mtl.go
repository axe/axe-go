package axe

import (
	"regexp"

	"github.com/axe/axe-go/pkg/asset"
	"github.com/udhos/gwob"
)

var _ asset.Format = &MtlFormat{}
var mtlFormatRegex, _ = regexp.Compile(`\.mtl$`)

type MtlFormat struct {
}

func (format *MtlFormat) Handles(ref asset.Ref) bool {
	return mtlFormatRegex.MatchString(ref.URI)
}

func (format *MtlFormat) Types() []asset.Type {
	return []asset.Type{asset.TypeModel}
}

func (format *MtlFormat) Load(a *asset.Asset) error {
	lib, err := gwob.ReadMaterialLibFromReader(a.SourceReader, &gwob.ObjParserOptions{})

	if err != nil {
		a.LoadStatus.Fail(err)
		return err
	}

	materials := Materials{}
	a.Dependent = make([]asset.Ref, 0)

	for materialName, material := range lib.Lib {
		mat := Material{
			Name:       material.Name,
			Illum:      material.Illum,
			Opacity:    material.D,
			Refraction: material.Ni,
			Shininess:  material.Ns,
			Ambient: TextureColor{
				Texture: material.MapKa,
				Color:   material.Ka,
			},
			Diffuse: TextureColor{
				Texture: material.MapKd,
				Color:   material.Kd,
			},
			Specular: TextureColor{
				Texture: material.MapKs,
				Color:   material.Ks,
			},
			Emissive: TextureColor{
				Texture: material.MapKe,
			},
			TextureBump: material.Bump,
		}

		if mat.Ambient.Texture != "" {
			a.AddNext(mat.Ambient.Texture, true)
		}
		if mat.Specular.Texture != "" {
			a.AddNext(mat.Specular.Texture, true)
		}
		if mat.Diffuse.Texture != "" {
			a.AddNext(mat.Diffuse.Texture, true)
		}
		if mat.Emissive.Texture != "" {
			a.AddNext(mat.Emissive.Texture, true)
		}
		if mat.TextureBump != "" {
			a.AddNext(mat.TextureBump, true)
		}

		materials[materialName] = mat
	}

	a.Data = materials
	a.LoadStatus.Success()

	return nil
}

func (format *MtlFormat) Unload(asset *asset.Asset) error {
	asset.LoadStatus.Reset()
	return nil
}
func (format *MtlFormat) Activate(asset *asset.Asset) error {
	asset.ActivateStatus.Success()
	return nil
}
func (format *MtlFormat) Deactivate(asset *asset.Asset) error {
	asset.ActivateStatus.Reset()
	return nil
}
