package ui

import (
	"unicode"
)

type FontRune struct {
	// The extends around the baseline point at the start of the font rune
	ExtentTile
	// The advance width of the rune
	Width float32
}

type Font struct {
	// The name of the font
	Name string
	// The texture of the font
	Texture string
	// The lineheight relative to the font size. 1.0=matches font-size
	LineHeight float32
	// Within a line how far down is the baseline from the top relative to the font size.
	Baseline float32
	// Data about each rune in the font
	Runes map[rune]FontRune
	// Kerning data when run B follows rune A
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
	BlockSpacing      float32
	ClampLeft         bool
	ClampTop          bool
}

type RenderedGlyphBlock struct {
	Height float32
	Glyphs []RenderedGlyph
}

type RenderedGlyph struct {
	Tile
	Bounds Bounds
	Color  Color
}

func (block GlyphBlock) GetLineHeight(actualLineHeight float32) float32 {
	if block.LineHeight == 0 {
		return actualLineHeight
	} else {
		return block.LineHeight
	}
}

func (block GlyphBlock) GetStates(theme *Theme) []GlyphState {
	states := make([]GlyphState, len(block.Glyphs))
	var prev Glyph
	for i, g := range block.Glyphs {
		states[i] = g.GetState(theme, block.Wrap, prev)
		prev = block.Glyphs[i]
	}
	return states
}

func (block GlyphBlock) UnwrappedSize(theme *Theme, scale Coord, blocks GlyphBlocks) Coord {
	states := block.GetStates(theme)
	lineWidth := float32(0)
	lineHeight := float32(0)
	lineCount := 0
	size := Coord{}

	for glyphIndex := range block.Glyphs {
		state := states[glyphIndex]
		if state.ShouldBreak {
			if lineWidth > size.X {
				size.X = lineWidth
			}
			if lineCount > 0 {
				size.Y += block.LineSpacing
			}
			size.Y += block.GetLineHeight(lineHeight) * scale.Y
			lineCount++
			lineWidth = 0
			lineHeight = 0
		} else {
			lineWidth += state.Size.X * scale.X
			if state.Size.Y > lineHeight {
				lineHeight = state.Size.Y
			}
		}
	}

	return size
}

func (block GlyphBlock) Render(theme *Theme, blocks GlyphBlocks) RenderedGlyphBlock {
	states := block.GetStates(theme)

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

		if state.ShouldBreak {
			nextLine := line{glyphs: []int{}}
			lines = append(lines, currentLine)
			currentLine = nextLine

			if state.Empty {
				continue
			}
		} else if blocks.Wrap(nextWidth) {
			wrap := len(currentLine.glyphs) - 1
			for wrap > 0 && !states[currentLine.glyphs[wrap]].CanBreak {
				wrap--
			}
			if wrap < 0 {
				wrap = 0
			}

			nextLine := line{glyphs: currentLine.glyphs[wrap:]}
			for _, k := range nextLine.glyphs {
				ks := states[k]
				nextLine.width += ks.Size.X
				currentLine.width -= ks.Size.X
				if ks.Size.Y > nextLine.height {
					nextLine.height = ks.Size.Y
				}
			}

			currentLine.glyphs = currentLine.glyphs[:wrap]
			lines = append(lines, currentLine)
			currentLine = nextLine
		}

		currentLine.width += state.Size.X
		if state.Size.Y > currentLine.height {
			currentLine.height = state.Size.Y
		}
		currentLine.glyphs = append(currentLine.glyphs, glyphIndex)
	}

	lines = append(lines, currentLine)

	totalHeight := float32(0)
	for lineIndex, line := range lines {
		if lineIndex > 0 {
			totalHeight += block.LineSpacing
		}
		totalHeight += block.GetLineHeight(line.height)
	}

	rendered := make([]RenderedGlyph, 0, len(block.Glyphs))
	offsetY := float32(0)

	for _, line := range lines {
		start := Coord{
			X: block.HorizontalAlignment * (blocks.MaxWidth - line.width),
			Y: offsetY,
		}
		if blocks.ClampLeft && start.X < 0 {
			start.X = 0
		}
		for _, glyphIndex := range line.glyphs {
			g := block.Glyphs[glyphIndex]
			s := states[glyphIndex]
			currentY := start.Y
			start.Y += block.VerticalAlignment * (line.height - s.Size.Y)
			render := g.Render(theme, start)
			if render.Bounds.Width() > 0 {
				rendered = append(rendered, render)
			}
			start.X += s.Size.X
			start.Y = currentY
		}
		offsetY += block.GetLineHeight(line.height) + block.LineSpacing
	}

	return RenderedGlyphBlock{Height: totalHeight, Glyphs: rendered}
}

func (block RenderedGlyphBlock) Translate(x, y float32) {
	for i := range block.Glyphs {
		block.Glyphs[i].Bounds.Translate(x, y)
	}
}

func (blocks GlyphBlocks) Wrap(lineWidth float32) bool {
	return blocks.MaxWidth > 0 && lineWidth > blocks.MaxWidth
}

func (blocks GlyphBlocks) UnwrappedSize(theme *Theme, scale Coord) Coord {
	size := Coord{}
	for i, block := range blocks.Blocks {
		blockSize := block.UnwrappedSize(theme, scale, blocks)
		if i > 0 {
			size.Y += blocks.BlockSpacing
		}
		size.Y += blockSize.Y
		if blockSize.X > size.X {
			size.X = blockSize.X
		}
	}
	return size
}

func (blocks GlyphBlocks) Render(theme *Theme) RenderedGlyphBlock {
	rendered := make([]RenderedGlyphBlock, len(blocks.Blocks))
	totalHeight := float32(0)
	totalGlyphs := 0
	for i, block := range blocks.Blocks {
		rendered[i] = block.Render(theme, blocks)
		totalHeight += rendered[i].Height
		if i > 0 {
			totalHeight += blocks.BlockSpacing
		}
		totalGlyphs += len(rendered[i].Glyphs)
	}
	offset := blocks.VerticalAlignment * (blocks.MaxHeight - totalHeight)
	for i := range rendered {
		if offset != 0 {
			rendered[i].Translate(0, offset)
		}
		offset += rendered[i].Height + blocks.BlockSpacing
	}
	joined := RenderedGlyphBlock{
		Height: totalHeight,
		Glyphs: make([]RenderedGlyph, 0, totalGlyphs),
	}
	for _, block := range rendered {
		joined.Glyphs = append(joined.Glyphs, block.Glyphs...)
	}
	return joined
}

type TextGlyph struct {
	Text  rune
	Font  string
	Size  Amount
	Color Color

	initialized bool
	font        *Font
	fontRune    *FontRune
}

var _ Glyph = &TextGlyph{}

func (g *TextGlyph) init(theme *Theme) {
	if !g.initialized {
		font := theme.Fonts[g.Font]
		if font == nil {
			font = theme.Fonts[theme.DefaultFont]
		}
		if font != nil {
			fontRune := font.Runes[g.Text]
			if fontRune.Width != 0 {
				g.font = font
				g.fontRune = &fontRune
			}
		}
		g.initialized = true
	}
}

func (g TextGlyph) getColor(theme *Theme) Color {
	if g.Color.IsZero() {
		return theme.DefaultFontColor
	} else {
		return g.Color
	}
}

func (g TextGlyph) getSize(theme *Theme) float32 {
	if g.Size.Value == 0 {
		return theme.DefaultFontSize
	} else {
		return g.Size.Get(theme.DefaultFontSize)
	}
}

func (g *TextGlyph) GetState(theme *Theme, wrap TextWrap, prev Glyph) GlyphState {
	space := unicode.IsSpace(g.Text)
	state := GlyphState{
		CanBreak:    wrap == TextWrapChar || (wrap == TextWrapWord && space),
		ShouldBreak: g.Text == '\r' || g.Text == '\n',
		Empty:       space,
	}

	g.init(theme)
	if g.font == nil {
		state.Empty = true
		return state
	}

	offset := float32(0)
	if prev != nil {
		if prevGlyph, ok := prev.(*TextGlyph); ok {
			offset = g.font.GetKerning(prevGlyph.Text, g.Text)
		}
	}

	size := g.getSize(theme)
	state.Size.X = (size * g.fontRune.Width) + offset*size
	state.Size.Y = size * g.font.LineHeight

	return state
}

func (g *TextGlyph) Render(theme *Theme, topLeft Coord) RenderedGlyph {
	g.init(theme)
	if g.font == nil {
		return RenderedGlyph{}
	}
	color := g.getColor(theme)
	size := g.getSize(theme)
	extents := g.fontRune.Extent
	baselineOffset := g.font.Baseline * size

	return RenderedGlyph{
		Tile:  g.fontRune.Tile,
		Color: color,
		Bounds: Bounds{
			Left:   topLeft.X - extents.Left*size,
			Right:  topLeft.X + extents.Right*size,
			Top:    topLeft.Y - extents.Top*size + baselineOffset,
			Bottom: topLeft.Y + extents.Bottom*size + baselineOffset,
		},
	}
}

// Parses text and converts it to glyphs using a special format.
// In the text you can set:
// - font = {f:name} or {f:} to reset to default
// - color = {c:#RRGGBB} or {c:#RRGGBBAA} or {c:red} or {c:} to reset to default
// - size = {s:12} or {s:50%} or {s:} to reset to default
// - new paragraph = {p}
// - set kerning to paragraph = {k:5}
// - set horizontal alignment in paragraph = {a:0.5}
// - set vertical alignment within a line in paragraph = {va:0}
// - set wrap in paragraph = {w:none} or {w:word} or {w:char}
// - set line spacing in paragraph = {ls:10}
// - set line height in paragraph = {lh:10} or {lh:120%}
// - set spacing between paragraphs = {ps:10}
// - set vertical alignment of paragraphs in block = {pv:0.5}
// - set max height in block = {mh:100}
// - set max width in block = {mw:100}
//
// Ex: "Hello {f:roboto}World{f:} {s:20}HOW{s:} {c:#00F}blue {c:red}red {c:}"
func TextToBlocks(text string) GlyphBlocks {
	blocks := GlyphBlocks{}
	// glyph := TextGlyph{}
	// glyphs := make([]GlyphBlock, 0, len(text))

	return blocks
}
