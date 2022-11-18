package glfw

import axe "github.com/axe/axe-go/pkg"

func Setup(game *axe.Game) {
	game.Input = NewInputSystem()
	game.Windows = NewWindowSystem()
	game.Graphics = NewGraphicsSystem()

	game.Assets.AddFormat(&TextureLoader{})
}
