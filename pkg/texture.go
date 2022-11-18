package axe

import (
	"fmt"

	"github.com/axe/axe-go/pkg/geom"
)

type Texture interface {
	Asset() *Asset
	Width() int
	Height() int
}

type TextureCoord struct {
	U float32
	V float32
}

type Tile struct {
	Texture Texture
	Name    string
	Coord0  TextureCoord
	Coord1  TextureCoord
}

func NewTile(tex Texture) Tile {
	name := ""
	if tex.Asset() != nil {
		name = tex.Asset().Ref.UniqueName()
	}

	return Tile{
		Texture: tex,
		Name:    name,
		Coord0:  TextureCoord{0, 0},
		Coord1:  TextureCoord{1, 1},
	}
}

func (t *Tile) Rect(x, y, w, h int) Tile {
	tw := float32(t.Texture.Width())
	th := float32(t.Texture.Height())

	return Tile{
		Texture: t.Texture,
		Name:    fmt.Sprintf("%s (x=%d, y=%d, w=%d, h=%d)", t.Name, x, y, w, h),
		Coord0:  TextureCoord{float32(x) / tw, float32(y) / th},
		Coord1:  TextureCoord{float32(x+w) / tw, float32(y+h) / th},
	}
}

func (t *Tile) Rects(w, h int, pos []geom.Vec2i) Tiles {
	rects := make(Tiles, len(pos))
	for i, p := range pos {
		rects[i] = t.Rect(p.X, p.Y, w, h)
	}
	return rects
}

func (t *Tile) RectsGrid(x, y, w, h, columns, rows int) Tiles {
	grid := make(Tiles, 0, columns*rows)
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			grid = append(grid, t.Rect(x+c*w, y+r*h, w, h))
		}
	}
	return grid
}

type Tiles []Tile

func (tiles Tiles) Textures() []Texture {
	unique := make(map[*Asset]Texture)
	for _, tile := range tiles {
		unique[tile.Texture.Asset()] = tile.Texture
	}
	textures := make([]Texture, 0, len(unique))
	for _, tex := range unique {
		textures = append(textures, tex)
	}
	return textures
}

func (t *Tiles) Add(tile Tile) {
	*t = append(*t, tile)
}

func (t *Tiles) AddTiles(tiles Tiles) {
	*t = append(*t, tiles...)
}
