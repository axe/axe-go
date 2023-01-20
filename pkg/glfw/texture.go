package glfw

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"regexp"
	"strings"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/asset"
	"github.com/go-gl/gl/v2.1/gl"
)

type texture struct {
	asset *asset.Asset
	image *image.RGBA
	id    uint32
}

var _ axe.Texture = &texture{}

func (tex *texture) Asset() *asset.Asset { return tex.asset }
func (tex *texture) Width() int          { return tex.image.Rect.Size().X }
func (tex *texture) Height() int         { return tex.image.Rect.Size().Y }

type TextureFormat struct{}

var _ asset.Format = &TextureFormat{}
var textureLoaderRegex, _ = regexp.Compile(`\.(png|jpg|jpeg)$`)

func (loader *TextureFormat) Handles(ref asset.Ref) bool {
	return textureLoaderRegex.MatchString(strings.ToLower(ref.URI))
}

func (loader *TextureFormat) Types() []asset.Type {
	return []asset.Type{asset.TypeTexture}
}

func (loader *TextureFormat) Load(asset *asset.Asset) error {
	tex := &texture{}

	asset.LoadStatus.Start()

	img, _, err := image.Decode(asset.SourceReader)
	if err != nil {
		asset.LoadStatus.Fail(err)
		return err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		err = errors.New("unsupported texture Stride")
		asset.LoadStatus.Fail(err)
		return err
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	tex.image = rgba

	asset.LoadStatus.Success()
	asset.Data = tex

	return nil
}

func (loader *TextureFormat) Unload(asset *asset.Asset) error {
	asset.LoadStatus.Reset()
	asset.Data = nil
	return nil
}

func (loader *TextureFormat) Activate(asset *asset.Asset) error {
	asset.ActivateStatus.Start()

	tex, isTexture := asset.Data.(*texture)
	if !isTexture {
		err := fmt.Errorf("error activating missing texture")
		asset.ActivateStatus.Fail(err)
		return err
	}

	gl.GetError()

	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &tex.id)
	gl.BindTexture(gl.TEXTURE_2D, tex.id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(tex.image.Rect.Size().X),
		int32(tex.image.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(tex.image.Pix))

	errCode := gl.GetError()

	if errCode != 0 {
		err := fmt.Errorf("error activating texture with code %d", errCode)
		asset.ActivateStatus.Fail(err)
		return err
	}

	asset.ActivateStatus.Success()

	return nil
}

func (loader *TextureFormat) Deactivate(asset *asset.Asset) error {
	asset.ActivateStatus.Reset()
	if tex, ok := asset.Data.(*texture); ok {
		gl.DeleteTextures(1, &tex.id)
		tex.id = 0
	}
	return nil
}
