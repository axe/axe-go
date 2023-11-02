package ui

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
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

type TextWrap string

const (
	TextWrapNone TextWrap = "none"
	TextWrapWord TextWrap = "word"
	TextWrapChar TextWrap = "char"
)

func (w TextWrap) MarshalText() ([]byte, error) {
	return []byte(w), nil
}

func (w *TextWrap) UnmarshalText(text []byte) error {
	s := strings.ToLower(string(text))
	switch s {
	case "none", "n", "":
		*w = TextWrapNone
	case "word", "w":
		*w = TextWrapWord
	case "char", "c", "letter":
		*w = TextWrapChar
	default:
		return fmt.Errorf("invalid text wrap: " + s)
	}
	return nil
}

type Glyph interface {
	GetState(theme *Theme, ctx AmountContext, wrap TextWrap, prev Glyph) GlyphState
	Render(theme *Theme, ctx AmountContext, start Coord) RenderedGlyph
}

type GlyphState struct {
	Size        Coord
	CanBreak    bool
	ShouldBreak bool
	Empty       bool
}

type GlyphBlock struct {
	Spacing             Amount  // extra space between glyphs
	LineSpacing         Amount  // extra space between lines
	LineHeight          Amount  // 0=calculated per line
	VerticalAlignment   float32 // 0=top, 0.5=middle, 1=bottom
	HorizontalAlignment float32 // 0=left, 0.5=middle, 1=right
	Wrap                TextWrap
	Padding             AmountBounds
	Glyphs              []Glyph
}

type GlyphBlocks struct {
	Blocks            []GlyphBlock
	MaxWidth          float32
	MaxHeight         float32
	VerticalAlignment float32
	BlockSpacing      Amount
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

func (block GlyphBlock) GetLineHeight(ctx AmountContext, actualLineHeight float32) float32 {
	if block.LineHeight.IsZero() {
		return actualLineHeight
	} else {
		return block.LineHeight.Get(ctx)
	}
}

func (block GlyphBlock) GetStates(theme *Theme, ctx AmountContext) []GlyphState {
	states := make([]GlyphState, len(block.Glyphs))
	var prev Glyph
	for i, g := range block.Glyphs {
		states[i] = g.GetState(theme, ctx, block.Wrap, prev)
		prev = block.Glyphs[i]
	}
	return states
}

func (block GlyphBlock) UnwrappedSize(theme *Theme, ctx AmountContext, scale Coord, blocks GlyphBlocks) Coord {
	states := block.GetStates(theme, ctx)
	lineWidth := float32(0)
	lineHeight := float32(0)
	lineCount := 0
	lineSpacing := block.LineSpacing.Get(ctx)
	size := Coord{}

	for glyphIndex := range block.Glyphs {
		state := states[glyphIndex]
		if state.ShouldBreak {
			if lineWidth > size.X {
				size.X = lineWidth
			}
			if lineCount > 0 {
				size.Y += lineSpacing
			}
			size.Y += block.GetLineHeight(ctx, lineHeight) * scale.Y
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

	padding := block.Padding.GetBounds(ctx)

	size.Y += padding.Top
	size.Y += padding.Bottom
	size.X += padding.Left
	size.X += padding.Right

	return size
}

func (block GlyphBlock) Render(theme *Theme, ctx AmountContext, blocks GlyphBlocks) RenderedGlyphBlock {
	states := block.GetStates(theme, ctx)

	type line struct {
		width  float32
		height float32
		glyphs []int
	}

	padding := block.Padding.GetBounds(ctx)
	paddingWidth := padding.Left + padding.Right

	lines := make([]line, 0, 8)
	currentLine := line{glyphs: []int{}, width: paddingWidth}

	for glyphIndex := range block.Glyphs {
		state := states[glyphIndex]
		nextWidth := currentLine.width + state.Size.X

		if state.ShouldBreak {
			nextLine := line{glyphs: []int{}, width: paddingWidth}
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
				nextLine.width += states[k].Size.X
			}

			currentLine.width -= nextLine.width
			currentLine.glyphs = currentLine.glyphs[:wrap]
			lines = append(lines, currentLine)
			currentLine = nextLine
			currentLine.width += paddingWidth
		}

		currentLine.width += state.Size.X
		currentLine.glyphs = append(currentLine.glyphs, glyphIndex)
	}

	if len(currentLine.glyphs) > 0 {
		lines = append(lines, currentLine)
	}
	lineSpacing := block.LineSpacing.Get(ctx)

	totalHeight := padding.Top + padding.Bottom
	for lineIndex := range lines {
		line := &lines[lineIndex]
		actualLineHeight := float32(0)
		for _, glyphIndex := range line.glyphs {
			state := states[glyphIndex]
			if state.Size.Y > actualLineHeight {
				actualLineHeight = state.Size.Y
			}
		}
		line.height = block.GetLineHeight(ctx, actualLineHeight)
		totalHeight += line.height
		if lineIndex > 0 {
			totalHeight += lineSpacing
		}
	}

	rendered := make([]RenderedGlyph, 0, len(block.Glyphs))
	offsetY := padding.Top

	for _, line := range lines {
		start := Coord{
			X: block.HorizontalAlignment * (blocks.MaxWidth - line.width),
			Y: offsetY,
		}
		if blocks.ClampLeft && start.X < 0 {
			start.X = 0
		}
		start.X += padding.Left

		for _, glyphIndex := range line.glyphs {
			g := block.Glyphs[glyphIndex]
			s := states[glyphIndex]
			start.Y = offsetY + block.VerticalAlignment*(line.height-s.Size.Y)
			render := g.Render(theme, ctx, start)
			if render.Bounds.Width() > 0 {
				rendered = append(rendered, render)
			}
			start.X += s.Size.X
		}

		offsetY += line.height + lineSpacing
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

func (blocks GlyphBlocks) UnwrappedSize(theme *Theme, ctx AmountContext, scale Coord) Coord {
	size := Coord{}
	blockCtx := ctx.WithParent(blocks.MaxWidth, blocks.MaxHeight)
	blockSpacing := blocks.BlockSpacing.Get(blockCtx)
	for i, block := range blocks.Blocks {
		blockSize := block.UnwrappedSize(theme, blockCtx, scale, blocks)
		if i > 0 {
			size.Y += blockSpacing
		}
		size.Y += blockSize.Y
		if blockSize.X > size.X {
			size.X = blockSize.X
		}
	}
	return size
}

func (blocks GlyphBlocks) Render(theme *Theme, ctx AmountContext) RenderedGlyphBlock {
	rendered := make([]RenderedGlyphBlock, len(blocks.Blocks))
	totalHeight := float32(0)
	totalGlyphs := 0
	blockCtx := ctx.WithParent(blocks.MaxWidth, blocks.MaxHeight)
	blockSpacing := blocks.BlockSpacing.Get(blockCtx)
	for i, block := range blocks.Blocks {
		rendered[i] = block.Render(theme, blockCtx, blocks)
		totalHeight += rendered[i].Height
		if i > 0 {
			totalHeight += blockSpacing
		}
		totalGlyphs += len(rendered[i].Glyphs)
	}
	offset := blocks.VerticalAlignment * (blocks.MaxHeight - totalHeight)
	for i := range rendered {
		if offset != 0 {
			rendered[i].Translate(0, offset)
		}
		offset += rendered[i].Height + blockSpacing
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

func (g TextGlyph) getSize(ctx AmountContext) float32 {
	if g.Size.Value == 0 {
		return ctx.FontSize
	} else {
		return g.Size.Get(ctx)
	}
}

func (g *TextGlyph) GetState(theme *Theme, ctx AmountContext, wrap TextWrap, prev Glyph) GlyphState {
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

	size := g.getSize(ctx)
	state.Size.X = (size * g.fontRune.Width) + offset*size
	state.Size.Y = size * g.font.LineHeight

	return state
}

func (g *TextGlyph) Render(theme *Theme, ctx AmountContext, topLeft Coord) RenderedGlyph {
	g.init(theme)
	if g.font == nil {
		return RenderedGlyph{}
	}
	color := g.getColor(theme)
	size := g.getSize(ctx)
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

var TextFormatRegex = regexp.MustCompile(`\\{|\{([^:}]+):?([^}]*)\}|.|\s`)

// Parses text and converts it to glyphs using a special format.
// In the text you can set:
// - font = {f:name} or {f:} to reset to default
// - color = {c:#RRGGBB} or {c:#RRGGBBAA} or {c:red} or {c:} to reset to default
// - size = {s:12} or {s:50%} or {s:} to reset to default
// - new paragraph = {p}
// - set kerning to paragraph = {k:5}
// - set horizontal alignment in paragraph = {h:0.5}
// - set vertical alignment within a line in paragraph = {v:0}
// - set wrap in paragraph = {w:none} or {w:word} or {w:char}
// - set line spacing in paragraph = {ls:10}
// - set line height in paragraph = {lh:10} or {lh:120%}
// - set padding in paragraph = {pa:10} or {pl:10%} or {pr:0} or {pb} or {pt:}
// - set spacing between paragraphs = {ps:10}
// - set vertical alignment of paragraphs in block = {pv:0.5}
// - set max height in block = {mh:100}
// - set max width in block = {mw:100}
// - set clamp left in block = {cl:1}
// - set clamp top in block = {ct:0}
//
// Ex: "Hello {f:roboto}World{f:} {s:20}HOW{s:} {c:#00F}blue {c:red}red {c:} \{"
func TextToBlocks(text string) (blocks GlyphBlocks, err error) {
	block := GlyphBlock{}
	glyph := TextGlyph{}

	readAmount := func(s string) Amount {
		amt := Amount{}
		if s != "" {
			err = amt.UnmarshalText([]byte(s))
		}
		return amt
	}
	readFloat := func(s string) float32 {
		f := float64(0)
		if s != "" {
			f, err = strconv.ParseFloat(s, 32)
		}
		return float32(f)
	}
	readWrap := func(s string) TextWrap {
		wrap := TextWrap("")
		err = wrap.UnmarshalText([]byte(s))
		return wrap
	}
	readColor := func(s string) Color {
		color := Color{}
		if s != "" {
			err = color.UnmarshalText([]byte(s))
		}
		return color
	}
	readBool := func(s string) bool {
		b := false
		if s != "" {
			b, err = strconv.ParseBool(s)
		}
		return b
	}

	pieces := TextFormatRegex.FindAllStringSubmatch(text, -1)
	for _, piece := range pieces {
		runes := piece[0]
		command := piece[1]
		value := piece[2]

		switch command {
		case "": // Add glyph to block
			copy := glyph
			copy.Text, _ = utf8.DecodeRuneInString(runes)
			block.Glyphs = append(block.Glyphs, &copy)
		case "f": // Glyph commands
			glyph.Font = value
		case "s":
			glyph.Size = readAmount(value)
		case "c":
			glyph.Color = readColor(value)
		case "k": // Block commands
			block.Spacing = readAmount(value)
		case "h":
			block.HorizontalAlignment = readFloat(value)
		case "v":
			block.VerticalAlignment = readFloat(value)
		case "w":
			block.Wrap = readWrap(value)
		case "ls":
			block.LineSpacing = readAmount(value)
		case "lh":
			block.LineHeight = readAmount(value)
		case "pa":
			block.Padding.SetAmount(readAmount(value))
		case "pl":
			block.Padding.Left = readAmount(value)
		case "pt":
			block.Padding.Top = readAmount(value)
		case "pr":
			block.Padding.Right = readAmount(value)
		case "pb":
			block.Padding.Bottom = readAmount(value)
		case "p": // Blocks commands
			blocks.Blocks = append(blocks.Blocks, block)
			block.Glyphs = make([]Glyph, 0)
		case "ps":
			blocks.BlockSpacing = readAmount(value)
		case "pv":
			blocks.VerticalAlignment = readFloat(value)
		case "mw":
			blocks.MaxWidth = readFloat(value)
		case "mh":
			blocks.MaxHeight = readFloat(value)
		case "ct":
			blocks.ClampTop = readBool(value)
		case "cl":
			blocks.ClampLeft = readBool(value)
		}
		if err != nil {
			return
		}
	}

	if len(block.Glyphs) > 0 {
		blocks.Blocks = append(blocks.Blocks, block)
	}

	return
}

func MustTextToBlocks(text string) GlyphBlocks {
	blocks, err := TextToBlocks(text)
	if err != nil {
		panic(err)
	}
	return blocks
}

func TextToVisual(text string) (*VisualText, error) {
	blocks, err := TextToBlocks(text)
	if err != nil {
		return nil, err
	}
	return &VisualText{Glyphs: blocks}, nil
}

func MustTextToVisual(text string) *VisualText {
	return &VisualText{Glyphs: MustTextToBlocks(text)}
}
