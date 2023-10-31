package axe

type Graphics interface {
	Draw()
}

/*
Draw view (camera + placement + scene):
Frame ready?
Build per frame render state:
Get camera state (view, projection, position, frustum)
Get all renderers that should be rendered (in frustom or marked always)
Get all light perspectives (viewProjection, position, shadows?, ambient, diffuse, intensity colors)
Set viewport & scissor (based on placement dimensions)
Bind things to render
For each render pass...
	pass.render(frameData)
		for each technique, bind & record

Render passes are built when view size is given or changes:
- render pass for lit objects
- render pass for each shadow casting point light

*/
