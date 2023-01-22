package axe

import (
	"github.com/axe/axe-go/pkg/asset"
)

type AssetSystem struct {
	asset.Assets
}

var _ GameSystem = &AssetSystem{}

func NewAssetSystem() AssetSystem {
	return AssetSystem{
		Assets: asset.NewAssets(),
	}
}

func (assets *AssetSystem) Init(game *Game) error {
	assets.AddDefaults()
	assets.AddFormat(&ObjFormat{})
	assets.AddFormat(&MtlFormat{})

	if len(game.Settings.Assets) > 0 {
		assets.AddMany(game.Settings.Assets)
	}

	return nil
}

func (assets *AssetSystem) Update(game *Game) {

}
