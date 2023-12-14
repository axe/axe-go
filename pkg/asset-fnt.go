package axe

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"regexp"
	"strings"

	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/ui"
	"github.com/fzipp/bmfont"
)

type FontBitmapFormat struct{}

var _ asset.Format = &FontBitmapFormat{}
var textureLoaderRegex, _ = regexp.Compile(`\.(fnt)$`)

func (loader *FontBitmapFormat) Handles(ref asset.Ref) bool {
	return textureLoaderRegex.MatchString(strings.ToLower(ref.URI))
}

func (loader *FontBitmapFormat) Types() []asset.Type {
	return []asset.Type{asset.TypeFontBitmap}
}

func (loader *FontBitmapFormat) Load(a *asset.Asset) error {
	f := &ui.Font{
		Name:    a.Ref.Name,
		Runes:   make(map[rune]ui.FontRune),
		Kerning: make(map[rune]map[rune]float32),
	}

	a.LoadStatus.Start()

	desc, err := bmfont.ReadDescriptor(a.SourceReader)
	if err != nil {
		return err
	}

	scaleX := 1.0 / float32(desc.Common.ScaleW)
	scaleY := 1.0 / float32(desc.Common.ScaleH)
	fontScale := 1.0 / float32(desc.Info.Size)

	f.Baseline = float32(desc.Common.Base) * fontScale
	f.LineHeight = float32(desc.Common.LineHeight) * fontScale

	for r, char := range desc.Chars {
		f.Runes[r] = ui.FontRune{
			Width: float32(char.XAdvance) * fontScale,
			ExtentTile: ui.ExtentTile{
				Tile: gfx.Tile{
					Texture: &gfx.Texture{Name: desc.Pages[char.Page].File},
					TopLeft: gfx.TexelUV{
						X: float32(char.X) * scaleX,
						Y: float32(char.Y) * scaleY,
					},
					BottomRight: gfx.TexelUV{
						X: float32(char.X+char.Width) * scaleX,
						Y: float32(char.Y+char.Height) * scaleY,
					},
				},
				Extent: ui.Bounds{
					Left:   float32(-char.XOffset) * fontScale,
					Right:  float32(char.Width+char.XOffset) * fontScale,
					Top:    float32(desc.Common.Base-char.YOffset) * fontScale,
					Bottom: float32(char.Height+char.YOffset-desc.Common.Base) * fontScale,
				},
			},
		}
	}

	for chars, kerning := range desc.Kerning {
		if _, ok := f.Kerning[chars.First]; !ok {
			f.Kerning[chars.First] = make(map[rune]float32)
		}
		second := f.Kerning[chars.First]
		second[chars.Second] = float32(kerning.Amount) * scaleX
		f.Kerning[chars.First] = second
	}

	for pageIndex, page := range desc.Pages {
		if pageIndex > 0 {
			return fmt.Errorf("Only single page fonts are supported.")
		}
		a.Next = append(a.Next, asset.Ref{
			Name: page.File,
			URI:  a.Source.Relative(a.Ref.URI, page.File),
			Type: asset.TypeTexture,
			Options: TextureSettings{
				Min:    TextureFilterLinear,
				MipMap: TextureFilterLinear.Ptr(),
			},
		})
		f.Texture = page.File
	}

	a.LoadStatus.Success()
	a.Data = f

	return nil
}

func (loader *FontBitmapFormat) Unload(a *asset.Asset) error {
	a.LoadStatus.Reset()
	a.Data = nil
	return nil
}

func (loader *FontBitmapFormat) Activate(a *asset.Asset) error {
	a.ActivateStatus.Success()
	return nil
}

func (loader *FontBitmapFormat) Deactivate(a *asset.Asset) error {
	a.ActivateStatus.Reset()
	return nil
}
