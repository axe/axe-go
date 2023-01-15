package ui

import "unicode"

type FontRune struct {
	FocusTile
	Width float32
}

type Font struct {
	Name    string
	Texture string
	Runes   map[rune]FontRune
	Kerning map[rune]map[rune]float32
}

func (f Font) GetKerning(prev rune, next rune) float32 {
	kerningMap := f.Kerning[prev]
	if kerningMap != nil {
		if kerning, ok := kerningMap[next]; ok {
			return kerning
		}
	}
	return 0
}

type TextFormat struct {
	Font     string
	FontSize float32
}

type Glyph interface {
	GetState(theme *Theme, wrap TextWrap, prev Glyph) GlyphState
	Render(theme *Theme, start Coord) RenderedGlyph
}

type GlyphState struct {
	Size        Coord
	CanBreak    bool
	ShouldBreak bool
	Empty       bool
}

type GlyphBlock struct {
	Spacing             float32 // extra space between glyphs
	LineSpacing         float32 // extra space between lines
	LineHeight          float32 // 0=calculated per line
	VerticalAlignment   float32 // 0=top, 0.5=middle, 1=bottom
	HorizontalAlignment float32 // 0=left, 0.5=middle, 1=right
	Wrap                TextWrap
	Glyphs              []Glyph
}

type GlyphBlocks struct {
	Blocks            []GlyphBlock
	MaxWidth          float32
	MaxHeight         float32
	VerticalAlignment float32
}

func (blocks GlyphBlocks) Wrap(lineWidth float32) bool {
	return blocks.MaxWidth > 0 && lineWidth > blocks.MaxWidth
}

func (block GlyphBlock) Render(theme *Theme, blocks GlyphBlocks) []RenderedGlyph {
	states := make([]GlyphState, len(block.Glyphs))
	var prev Glyph
	for i, g := range block.Glyphs {
		states[i] = g.GetState(theme, block.Wrap, prev)
		prev = block.Glyphs[i]
	}

	type line struct {
		width  float32
		height float32
		glyphs []int
	}

	lines := make([]line, 8)
	currentLine := line{glyphs: []int{}}

	for glyphIndex := range block.Glyphs {
		state := states[glyphIndex]
		nextWidth := currentLine.width + state.Size.X

		if blocks.Wrap(nextWidth) {
			wrap := len(currentLine.glyphs) - 1
			for wrap > 0 && !states[currentLine.glyphs[wrap]].CanBreak {
				wrap--
			}

			nextLine := line{glyphs: currentLine.glyphs[wrap:]}
			for _, k := range nextLine.glyphs {
				ks := states[k]
				nextLine.width += ks.Size.X
				if ks.Size.Y > nextLine.height {
					nextLine.height = ks.Size.Y
				}
			}

			currentLine.glyphs = currentLine.glyphs[:wrap]
			lines = append(lines, currentLine)
			currentLine = nextLine
		} else if state.ShouldBreak {
			nextLine := line{glyphs: []int{}}
			lines = append(lines, currentLine)
			currentLine = nextLine

			if state.Empty {
				continue
			}
		}

		currentLine.width += state.Size.X
		if state.Size.Y > currentLine.height {
			currentLine.height = state.Size.Y
		}
		currentLine.glyphs = append(currentLine.glyphs, glyphIndex)
	}

	lines = append(lines, currentLine)

	rendered := make([]RenderedGlyph, 0, len(block.Glyphs))
	offsetY := block.LineHeight

	for _, line := range lines {
		start := Coord{
			X: block.HorizontalAlignment * (blocks.MaxWidth - line.width),
			Y: offsetY,
		}
		if block.LineHeight == 0 {
			start.Y += line.height
		}
		if start.X < 0 {
			start.X = 0
		}
		for _, glyphIndex := range line.glyphs {
			g := block.Glyphs[glyphIndex]
			render := g.Render(theme, start)
			if render.Bounds.Width() > 0 {
				rendered = append(rendered, render)
			}
			start.X += states[glyphIndex].Size.X
		}
		offsetY += block.LineHeight + block.LineSpacing
	}

	return rendered
}

type RenderedGlyph struct {
	Tile
	Bounds Bounds
	Color  Color
}

type TextGlyph struct {
	Text   rune
	Format string
	Color  Color
}

var _ Glyph = TextGlyph{}

func (g TextGlyph) getData(theme *Theme) (*TextFormat, *Font, *FontRune) {
	format := theme.TextFormats[g.Format]
	if format == nil {
		return nil, nil, nil
	}
	font := theme.Fonts[format.Font]
	if font == nil {
		return nil, nil, nil
	}
	fontRune := font.Runes[g.Text]
	if fontRune.Width == 0 {
		return nil, nil, nil
	}
	return format, font, &fontRune
}
func (g TextGlyph) GetState(theme *Theme, wrap TextWrap, prev Glyph) GlyphState {
	format, font, fontRune := g.getData(theme)
	if format == nil {
		return GlyphState{Empty: true}
	}

	offset := float32(0)
	if prev != nil {
		if prevGlyph, ok := prev.(TextGlyph); ok {
			offset = font.GetKerning(prevGlyph.Text, g.Text)
		}
	}

	space := unicode.IsSpace(g.Text)

	return GlyphState{
		CanBreak:    wrap == TextWrapChar || (wrap == TextWrapWord && space),
		ShouldBreak: g.Text == '\r' || g.Text == '\n',
		Empty:       space,
		Size: Coord{
			X: format.FontSize*fontRune.Width + offset,
			Y: format.FontSize,
		},
	}
}
func (g TextGlyph) Render(theme *Theme, start Coord) RenderedGlyph {
	format, _, fontRune := g.getData(theme)
	if format == nil {
		return RenderedGlyph{}
	}
	extents := fontRune.GetExtents()
	return RenderedGlyph{
		Tile:  fontRune.Tile,
		Color: g.Color,
		Bounds: Bounds{
			Left:   start.X - extents.Left*format.FontSize,
			Right:  start.X + extents.Right*format.FontSize,
			Top:    start.Y - extents.Top*format.FontSize,
			Bottom: start.Y + extents.Bottom*format.FontSize,
		},
	}
}
