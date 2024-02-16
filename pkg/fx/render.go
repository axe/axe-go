package fx

import (
	"image/color"

	"github.com/axe/axe-go/pkg/gfx"
)

type Renderer interface {
	CreateBuffer(maxParticles int) *gfx.Buffer
	Render()
}

type Vertex struct {
	X, Y, Z      float32
	TextureCoord gfx.TextureCoord
	Color        color.Color
}

type ParticleRenderer[P any] func(*P)

type Point2 struct {
	X, Y float32
}

type Particle2 struct {
	Position Point2
	Anchor   Point2
	Size     Point2
	Scale    Point2
	Angle    float32
	Tile     gfx.Tile
}

func (p *Particle2) Init() {
	p.Anchor.X = 0.5
	p.Anchor.Y = 0.5
	p.Scale.X = 1
	p.Scale.Y = 1
	p.Angle = 0
}

type Point3 struct {
	X, Y, Z float32
}

type Particle3 struct {
	Position Point3
	Up       Point3
	Right    Point3
	Anchor   Point2
	Size     Point2
	Scale    Point2
	Angle    float32
	Tile     gfx.Tile
}

func (p *Particle3) Init() {
	p.Anchor.X = 0.5
	p.Anchor.Y = 0.5
	p.Scale.X = 1
	p.Scale.Y = 1
	p.Angle = 0
}
