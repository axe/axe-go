package glfw

import (
	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/geom"
	"github.com/go-gl/gl/v2.1/gl"
)

func NewGraphicsSystem() axe.GraphicsSystem {
	return &graphicsSystem{
		offs: make(axe.ListenerOffs, 0),
	}
}

type graphicsSystem struct {
	texture   *texture
	rotationX float32
	rotationY float32
	window    *window
	offs      axe.ListenerOffs
}

var _ axe.GraphicsSystem = &graphicsSystem{}

func (gr *graphicsSystem) Init(game *axe.Game) error {
	err := gl.Init()
	if err != nil {
		return err
	}

	gr.window = game.Windows.MainWindow().(*window)

	off := game.Windows.Events().On(axe.WindowSystemEvents{
		WindowResize: func(window axe.Window, oldSize, newSize geom.Vec2i) {
			gr.resize(newSize.X, newSize.Y)
		},
	})
	gr.offs.Add(off)

	gl.Enable(gl.DEPTH_TEST) // view dependent
	gl.Enable(gl.LIGHTING)   // view dependent

	gl.ClearColor(0.5, 0.5, 0.5, 0.0) // system dependent
	gl.ClearDepth(1)                  // system dependent
	gl.DepthFunc(gl.LEQUAL)           // view dependent

	ambient := []float32{0.5, 0.5, 0.5, 1}
	diffuse := []float32{1, 1, 1, 1}
	lightPosition := []float32{-5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])        // view dependent
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])        // view dependent
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0]) // view dependent
	gl.Enable(gl.LIGHT0)                                  // view dependent

	gr.resize(gr.window.size.X, gr.window.size.Y) // view dependent
	gl.MatrixMode(gl.MODELVIEW)                   // view dependent
	gl.LoadIdentity()                             // view dependent

	asset := game.Assets.Add(axe.AssetRef{
		Name: "square",
		URI:  "./square.png",
		Type: axe.AssetTypeTexture,
	})
	err = asset.Activate()
	if err != nil {
		return err
	}
	gr.texture = asset.Data.(*texture)

	return nil
}
func (gr *graphicsSystem) resize(width int, height int) {
	// gl.Viewport(0, 0, int32(width), int32(height))

	// 3d
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	f := float64(width)/float64(height) - 1
	gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 10.0)

	// 2d
	// gl.MatrixMode(gl.PROJECTION)
	// gl.LoadIdentity()
	// gl.Ortho(0, float64(w), float64(h), 0, 0, 1)
}
func (gr *graphicsSystem) Update(game *axe.Game) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) // system dependent (potentially)

	/**
	 * get current stage
	 * for each view in current stage
	 *    set render area based on placement
	 *    update camera
	 *    update view matrix from camera
	 *    update combined matrix
	 *    transform world
	 *    get all renderables
	 *      if in space, skip if not in camera view
	 *      render
	 */

	gl.MatrixMode(gl.MODELVIEW) // view dependent
	gl.LoadIdentity()           // view dependent

	// done by renderer
	gl.Translatef(0, 0, -3.0)
	gl.Rotatef(gr.rotationX, 1, 0, 0)
	gl.Rotatef(gr.rotationY, 0, 1, 0)

	dt := game.State.UpdateTimer.Elapsed.Seconds()

	gr.rotationX += float32(dt * 6)
	gr.rotationY += float32(dt * 4)

	gl.BindTexture(gl.TEXTURE_2D, gr.texture.id)

	gl.Color4f(1, 1, 1, 1)

	gl.Begin(gl.QUADS)

	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, 1)

	gl.Normal3f(0, 0, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, 1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, -1)

	gl.Normal3f(0, 1, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, 1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, -1)

	gl.Normal3f(0, -1, 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, -1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, 1)

	gl.Normal3f(1, 0, 0)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, -1, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, 1)

	gl.Normal3f(-1, 0, 0)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, 1, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, -1)

	gl.End()

	gr.window.window.SwapBuffers() // system dependent
}

func (gr *graphicsSystem) Destroy() {
	gr.offs.Off()
}
