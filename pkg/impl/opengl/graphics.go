package opengl

import (
	axe "github.com/axe/axe-go/pkg"
	"github.com/axe/axe-go/pkg/color"
	"github.com/axe/axe-go/pkg/core"
	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/geom"
	"github.com/axe/axe-go/pkg/gfx"
	"github.com/axe/axe-go/pkg/ui"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewGraphicsSystem() axe.GraphicsSystem {
	return &graphicsSystem{
		offs: make(core.ListenerOffs, 0),
	}
}

type graphicsSystem struct {
	// texture   *texture
	// rotationX float32
	// rotationY float32
	window *window
	offs   core.ListenerOffs
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

	clear := gr.window.clear
	gl.ClearColor(clear.R, clear.G, clear.B, clear.A) // system dependent
	gl.ClearDepth(1)                                  // system dependent
	gl.DepthFunc(gl.LEQUAL)                           // view dependent

	// ambient := []float32{0.5, 0.5, 0.5, 1}
	// diffuse := []float32{1, 1, 1, 1}
	// lightPosition := []float32{-5, 5, 10, 0}
	// gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])        // view dependent
	// gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])        // view dependent
	// gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0]) // view dependent
	// gl.Enable(gl.LIGHT0)                                  // view dependent

	gr.resize(gr.window.size.X, gr.window.size.Y) // view dependent

	return nil
}

func (gr *graphicsSystem) resize(width int32, height int32) {
	// gl.Viewport(0, 0, int32(width), int32(height))

	// 3d
	// gl.MatrixMode(gl.PROJECTION)
	// gl.LoadIdentity()
	// f := float64(width)/float64(height) - 1
	// gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 10.0)

	// 2d
	// gl.MatrixMode(gl.PROJECTION)
	// gl.LoadIdentity()
	// gl.Ortho(0, float64(w), float64(h), 0, 0, 1)
}

func (gr *graphicsSystem) Update(game *axe.Game) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) // system dependent (potentially)

	defer gr.window.window.SwapBuffers()

	current := game.Stages.Current

	if current == nil {
		return
	}

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

	for _, view3 := range current.Views3 {
		gr.renderView3(view3, game)
	}

	for _, view2 := range current.Views2 {
		gr.renderView2(view2, game)
	}
}

func (gr *graphicsSystem) Destroy() {
	gr.offs.Off()
}

func (gr *graphicsSystem) renderView2(view axe.View2f, game *axe.Game) {
	scene := view.Scene()
	if scene == nil {
		return
	}
	scene.World.Activate()

	initView2(view, game)
	renderUserInterfaces(view, game)
}

func (gr *graphicsSystem) renderView3(view axe.View3f, game *axe.Game) {
	scene := view.Scene()
	if scene == nil {
		return
	}
	scene.World.Activate()

	initView3(view, game)
	renderLights()
	renderMeshes(game)
}

func initView3(view axe.View3f, game *axe.Game) {
	winSize := game.Windows.MainWindow().Size()
	bounds := view.Placement.GetBoundsi(float32(winSize.X), float32(winSize.Y))
	width := int32(bounds.Width())
	height := int32(bounds.Height())

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	f := float64(width)/float64(height) - 1
	gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 10.0)

	if view.Placement.IsMaximized() {
		gl.Viewport(0, 0, width, height)
	} else {
		gl.Viewport(int32(bounds.Min.X), int32(winSize.Y-bounds.Max.Y), int32(width), int32(height))
	}

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func renderLights() {
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
}

func applyTransform3(transform *axe.Transform[axe.Vec4[float32]]) {
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
}

func renderMeshes(game *axe.Game) {
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

		applyTransform3(axe.TRANSFORM3.Get(entityMesh.ID.Entity()))

		gl.Enable(gl.TEXTURE_2D)
		gl.Enable(gl.LIGHTING)
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
}

func placementWindowBounds(placement ui.Placement, game *axe.Game) ui.Bounds {
	winSize := game.Windows.MainWindow().Size()
	bounds := placement.GetBounds(float32(winSize.X), float32(winSize.Y))
	return bounds
}

func initView2(view axe.View2f, game *axe.Game) {
	winSize := game.Windows.MainWindow().Size()
	bounds := view.Placement.GetBoundsi(float32(winSize.X), float32(winSize.Y))
	width := int32(bounds.Width())
	height := int32(bounds.Height())

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(width), float64(height), 0, 0, 1)

	if view.Placement.IsMaximized() {
		gl.Viewport(0, 0, width, height)
	} else {
		gl.Viewport(int32(bounds.Min.X), int32(winSize.Y-bounds.Max.Y), int32(width), int32(height))
	}
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

var cursorBuffer = ui.BufferPool.Get()
var vertexBufferQueue = ui.BufferQueuePool.Get()

func renderUserInterfaces(view axe.View2f, game *axe.Game) {
	bounds := placementWindowBounds(view.Placement, game)
	width, height := bounds.Dimensions()
	window := game.Windows.MainWindow()
	windowSize := window.Size()
	screenSize := window.Screen().Size()

	ctx := &ui.AmountContext{
		Parent: ui.UnitContext{Width: width, Height: height},
		View:   ui.UnitContext{Width: width, Height: height},
		Window: ui.UnitContext{Width: float32(windowSize.X), Height: float32(windowSize.Y)},
		Screen: ui.UnitContext{Width: float32(screenSize.X), Height: float32(screenSize.Y)},
	}

	hadCursor := !cursorBuffer.Empty()

	cursorBuffer.Clear()

	uis := axe.UI.Iterable().Iterator()
	for uis.HasNext() {
		u := uis.Next().Data

		u.SetContext(ctx)
		u.Place(bounds)

		if u.NeedsRender() {
			vertexBufferQueue.Clear()
			u.Render(vertexBufferQueue)
		}

		if vertexBufferQueue.Len() > 0 {
			renderBuffers(vertexBufferQueue, game, windowSize.Y)
		}

		if cursorBuffer.Empty() {
			cursor := u.GetCursor()
			if cursor != nil {
				buf := cursorBuffer.Buffer()
				quad := buf.AddQuad()
				copy(quad, cursor)
				renderBuffers(cursorBuffer, game, windowSize.Y)
				if !hadCursor {
					glfw.GetCurrentContext().SetInputMode(glfw.CursorMode, glfw.CursorHidden)
				}
			}
		}
	}

	if cursorBuffer.Empty() {
		glfw.GetCurrentContext().SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
}

func renderBuffers(buffers ui.VertexIterable, game *axe.Game, windowHeight int32) {
	gl.Disable(gl.LIGHTING)
	gl.Disable(gl.TEXTURE_2D)

	lastBlend := gfx.BlendNone
	lastPrimitive := gfx.PrimitiveNone
	began := false
	lastTexture := ""
	coloring := false

	gl.Color4f(1, 1, 1, 1)

	for b := 0; b < buffers.Len(); b++ {
		vb := buffers.At(b)

		if vb.Empty() {
			continue
		}

		applyBlend(vb.Blend, &lastBlend)

		span := vb.IndexSpanAt(0)
		indices := span.Len()

		for i := 0; i < indices; i++ {
			v := span.At(i)

			applyTexture(game, v.Tex.Texture, &lastTexture, &began)
			applyPrimitive(vb.Primitive, &lastPrimitive, &began)
			applyColor(v.Color, v.HasColor, &coloring)
			if v.HasCoord {
				u, v := v.Tex.UV()
				gl.TexCoord2f(u, v)
			}
			gl.Vertex2f(v.X, v.Y)
		}
	}

	applyPrimitive(gfx.PrimitiveNone, &lastPrimitive, &began)
	applyBlend(gfx.BlendNone, &lastBlend)
	applyColor(color.White, false, &coloring)
}

var blendSources = ds.NewEnumMap(map[gfx.Blend]uint32{
	gfx.BlendAdd:          gl.ONE,
	gfx.BlendAlphaAdd:     gl.SRC_ALPHA,
	gfx.BlendAlpha:        gl.SRC_ALPHA,
	gfx.BlendColor:        gl.ONE,
	gfx.BlendMinus:        gl.ONE_MINUS_DST_ALPHA,
	gfx.BlendPremultAlpha: gl.ONE,
	gfx.BlendModulate:     gl.DST_COLOR,
	gfx.BlendXor:          gl.ONE_MINUS_DST_COLOR,
	gfx.BlendNone:         gl.ZERO,
})

var blendTargets = ds.NewEnumMap(map[gfx.Blend]uint32{
	gfx.BlendAdd:          gl.ONE,
	gfx.BlendAlphaAdd:     gl.ONE,
	gfx.BlendAlpha:        gl.ONE_MINUS_SRC_ALPHA,
	gfx.BlendColor:        gl.ONE_MINUS_SRC_COLOR,
	gfx.BlendMinus:        gl.DST_ALPHA,
	gfx.BlendPremultAlpha: gl.ONE_MINUS_SRC_ALPHA,
	gfx.BlendModulate:     gl.ZERO,
	gfx.BlendXor:          gl.ZERO,
	gfx.BlendNone:         gl.ONE,
})

func applyBlend(blend gfx.Blend, lastBlend *gfx.Blend) {
	if blend != *lastBlend {
		if blend == gfx.BlendNone {
			gl.Disable(gl.BLEND)
		} else {
			if *lastBlend == gfx.BlendNone {
				gl.Enable(gl.BLEND)
			}
			gl.BlendFunc(blendSources[blend], blendTargets[blend])
		}
		*lastBlend = blend
	}
}

func applyTexture(game *axe.Game, tex *gfx.Texture, lastTexture *string, began *bool) {
	newName := ""
	if tex != nil {
		newName = tex.Name
	}
	if newName != *lastTexture {
		if *began {
			gl.End()
			*began = false
		}
		if newName == "" {
			gl.Disable(gl.TEXTURE_2D)
		} else {
			if textureID, ok := getTextureID(tex, game); ok {
				gl.Enable(gl.TEXTURE_2D)
				gl.BindTexture(gl.TEXTURE_2D, textureID)
			}
		}
		*lastTexture = newName
	}
}

var primitiveMapping = ds.NewEnumMap(map[gfx.Primitive]uint32{
	gfx.PrimitiveTriangle: gl.TRIANGLES,
	gfx.PrimitiveLine:     gl.LINES,
	gfx.PrimitiveQuad:     gl.QUADS,
})

func applyPrimitive(primitive gfx.Primitive, lastPrimitive *gfx.Primitive, began *bool) {
	if primitive != *lastPrimitive || !*began {
		if *began {
			gl.End()
		}
		mapped := primitiveMapping.Get(primitive)
		if mapped != 0 {
			gl.Begin(mapped)
			*began = true
		}
		*lastPrimitive = primitive
	}
}

func applyColor(c color.Color, has bool, coloring *bool) {
	if has {
		gl.Color4f(c.R, c.G, c.B, c.A)
		*coloring = true
	} else if *coloring {
		gl.Color4f(1, 1, 1, 1)
		*coloring = false
	}
}

func getTextureID(tex *gfx.Texture, game *axe.Game) (uint32, bool) {
	if id, ok := tex.Info.Metadata.(uint32); ok {
		return id, true
	} else {
		textureAsset := game.Assets.GetEither(tex.Name)
		if textureAsset == nil {
			return 0, false
		}

		texture := textureAsset.Data.(*texture)
		tex.Info.Metadata = texture.id
		tex.Info.SetDimensions(texture.Width(), texture.Height())

		return texture.id, true
	}
}
