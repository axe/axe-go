package glfw

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/geom"
	"github.com/go-gl/gl/v2.1/gl"
)

func NewGraphicsSystem() axe.GraphicsSystem {
	return &graphicsSystem{}
}

type graphicsSystem struct {
	texture   uint32
	rotationX float32
	rotationY float32
	window    *window
}

var _ axe.GraphicsSystem = &graphicsSystem{}

func (gr *graphicsSystem) Init(game *axe.Game) error {
	err := gl.Init()
	if err != nil {
		return err
	}

	gr.window = game.Windows.MainWindow().(*window)

	game.Windows.Events().On(axe.WindowSystemEvents{
		WindowResize: func(window axe.Window, oldSize, newSize geom.Vec2i) {
			gr.resize(newSize.X, newSize.Y)
		},
	})

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.LIGHTING)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)

	ambient := []float32{0.5, 0.5, 0.5, 1}
	diffuse := []float32{1, 1, 1, 1}
	lightPosition := []float32{-5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Enable(gl.LIGHT0)

	gr.resize(gr.window.size.X, gr.window.size.Y)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gr.texture = newTexture("square.png")

	return nil
}
func (gr *graphicsSystem) resize(width int, height int) {
	// gl.Viewport(0, 0, int32(width), int32(height))

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	f := float64(width)/float64(height) - 1
	gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 10.0)
}
func (gr *graphicsSystem) Update(game *axe.Game) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

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

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, -3.0)
	gl.Rotatef(gr.rotationX, 1, 0, 0)
	gl.Rotatef(gr.rotationY, 0, 1, 0)

	dt := game.State.UpdateTimer.Elapsed.Seconds()

	gr.rotationX += float32(dt * 6)
	gr.rotationY += float32(dt * 4)

	gl.BindTexture(gl.TEXTURE_2D, gr.texture)

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

	gr.window.window.SwapBuffers()
}

func (gr *graphicsSystem) Destroy() {
	gl.DeleteTextures(1, &gr.texture)
}

func newTexture(file string) uint32 {
	imgFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("texture %q not found on disk: %v\n", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}
