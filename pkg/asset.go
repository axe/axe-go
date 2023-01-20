package axe

import (
	"github.com/axe/axe-go/pkg/asset"
)

type AssetSystem struct {
	Assets asset.Assets
}

var _ GameSystem = &AssetSystem{}

func NewAssetSystem() AssetSystem {
	return AssetSystem{
		Assets: asset.NewAssets(),
	}
}

func (assets *AssetSystem) Init(game *Game) error {
	assets.Assets.AddDefaults()
	assets.Assets.AddFormat(&ObjFormat{})
	assets.Assets.AddFormat(&MtlFormat{})

	if len(game.Settings.Assets) > 0 {
		assets.Assets.AddMany(game.Settings.Assets)
	}

	return nil
}

func (assets *AssetSystem) Update(game *Game) {

}

func (assets *AssetSystem) Destroy() {
	assets.Assets.Destroy()
}
