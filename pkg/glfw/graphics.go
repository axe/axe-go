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
	// texture   *texture
	// rotationX float32
	// rotationY float32
	window *window
	offs   axe.ListenerOffs
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

	// ambient := []float32{0.5, 0.5, 0.5, 1}
	// diffuse := []float32{1, 1, 1, 1}
	// lightPosition := []float32{-5, 5, 10, 0}
	// gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])        // view dependent
	// gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])        // view dependent
	// gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0]) // view dependent
	// gl.Enable(gl.LIGHT0)                                  // view dependent

	gr.resize(gr.window.size.X, gr.window.size.Y) // view dependent
	gl.MatrixMode(gl.MODELVIEW)                   // view dependent
	gl.LoadIdentity()                             // view dependent

	// asset := game.Assets.Add(axe.AssetRef{
	// 	URI: "./square.png",
	// })
	// err = asset.Activate()
	// if err != nil {
	// 	return err
	// }
	// gr.texture = asset.Data.(*texture)

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
	 *      is renderable not in vertex buffer yet?
	 * 			  add to vertex buffer
	 *      else if dirty
	 *        update vertex buffer
	 * for each renderable in a vertex buffer that hasn't been used in X frames, offload it
	 */

	if !axe.HasActiveWorld() {
		// fmt.Printf("no active world: %v\n", time.Now().Sub(game.State.StartTime))
		gr.window.window.SwapBuffers()
		return
	}

	gl.MatrixMode(gl.MODELVIEW) // view dependent
	gl.LoadIdentity()           // view dependent

	lights := axe.LIGHT.Iterable().Iterator()
	lightIndex := uint32(0)
	for lights.HasNext() {
		if lightIndex == 0 {
			gl.Enable(gl.LIGHTING)
		}
		light := lights.Next()
		gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &light.Data.Ambient.R)   // view dependent
		gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &light.Data.Diffuse.R)   // view dependent
		gl.Lightfv(gl.LIGHT0, gl.POSITION, &light.Data.Position.X) // view dependent
		gl.Enable(gl.LIGHT0 + lightIndex)                          // view dependent
		lightIndex++
		if lightIndex >= 8 {
			break
		}
	}
	if lightIndex == 0 {
		gl.Disable(gl.LIGHTING)
	}

	meshes := axe.MESH.Iterable().Iterator()
	for meshes.HasNext() {
		entityMesh := meshes.Next()

		meshAsset := game.Assets.GetRef(entityMesh.Data.Ref)
		if meshAsset == nil {
			// fmt.Println("no mesh asset")
			continue
		}
		meshData := meshAsset.Data.(axe.MeshData)
		meshMaterialsAsset := game.Assets.GetEither(meshData.Materials)
		if meshMaterialsAsset == nil {
			// fmt.Printf("no mesh materials asset %s\n", meshData.Materials)
			continue
		}
		meshMaterials := meshMaterialsAsset.Data.(axe.Materials)

		transform := axe.TRANSFORM3.Get(entityMesh.Entity)
		if transform != nil {
			pos := transform.GetPosition()
			rot := transform.GetRotation()
			scl := transform.GetScale()

			gl.LoadIdentity()
			gl.Translatef(pos.X, pos.Y, pos.Z)
			gl.Rotatef(rot.X, 1, 0, 0)
			gl.Rotatef(rot.Y, 0, 1, 0)
			gl.Rotatef(rot.Z, 0, 0, 1)
			gl.Scalef(scl.X, scl.Y, scl.Z)
		}

		for _, group := range meshData.Groups {
			if material, ok := meshMaterials[group.Material]; ok {
				textureAsset := game.Assets.GetEither(material.Diffuse.Texture)
				if textureAsset == nil {
					// fmt.Println("no texture asset")
					continue
				}
				texture := textureAsset.Data.(*texture)

				gl.BindTexture(gl.TEXTURE_2D, texture.id)
				gl.Color4f(1, 1, 1, 1)

				gl.Begin(gl.TRIANGLES)
				for _, face := range group.Faces {
					for i := 0; i < 3; i++ {
						if face.Normals != nil {
							gl.Normal3fv(&meshData.Normals[face.Normals[i]][0])
						}
						if face.Uvs != nil {
							gl.TexCoord3fv(&meshData.Uvs[face.Uvs[i]][0])
						}
						gl.Vertex3fv(&meshData.Vertices[face.Vertices[i]][0])
					}
				}
				gl.End()
			}

		}
	}

	gr.window.window.SwapBuffers() // system dependent
}

func (gr *graphicsSystem) Destroy() {
	gr.offs.Off()
}

func (gr *graphicsSystem) renderView2(view axe.View2f, game *axe.Game) {

}

func (gr *graphicsSystem) renderView3(view axe.View2f, game *axe.Game) {

}
