package opengl

import axe "github.com/axe/axe-go/pkg"

func Setup(game *axe.Game, settings Settings) {
	if !settings.HasVersion() {
		settings.Major = 2
		settings.Minor = 1
	}

	game.Input = NewInputSystem()
	game.Windows = NewWindowSystem(settings)
	game.Graphics = NewGraphicsSystem()

	game.Assets.AddFormat(&TextureFormat{})
}
