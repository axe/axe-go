package ui_test

import (
	"fmt"
	"testing"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/asset"
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/ui"
)

func TestTextMeasure(t *testing.T) {
	ctx := getContextWithFont("roboto")
	para := ui.MustTextToParagraphs("{h:0.5}{pv:0.5}Toggle FullHeight")
	para.MaxWidth = 88
	measure := para.Measure(ctx)
	para.MaxWidth = measure.X
	rendered := para.Render(ctx, nil)

	fmt.Printf("Measure: %+v\n", measure)
	fmt.Printf("Rendered Text Bounds: %+v\n", rendered.Bounds)
	for _, g := range rendered.Glyphs {
		fmt.Printf("%+v\n", g)
	}
}

func getContextWithFont(name string) *ui.RenderContext {
	game := axe.NewGame(axe.GameSettings{})
	game.Assets.Init(game)
	asset := game.Assets.Add(asset.Ref{Name: name, URI: "../../cmd/assets/" + name + ".fnt"})
	asset.Load()
	font := asset.Data.(*ui.Font)
	nameIdentifier := id.Get(name)
	user := ui.NewUI()
	user.Theme.Fonts.Set(nameIdentifier, font)
	user.Theme.TextStyles.Font = nameIdentifier
	user.Root = &ui.Base{}
	user.SetContext(&ui.AmountContext{
		Parent: ui.UnitContext{Width: 720, Height: 480},
	})
	return user.RenderContext()
}
