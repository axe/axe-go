package ui

import (
	"encoding"
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

type Alignment float32

const (
	AlignmentTop    Alignment = 0
	AlignmentLeft   Alignment = 0
	AlignmentMiddle Alignment = 0.5
	AlignmentCenter Alignment = 0.5
	AlignmentRight  Alignment = 1
	AlignmentBottom Alignment = 1
)

var _ encoding.TextMarshaler = Alignment(0)
var _ encoding.TextUnmarshaler = new(Alignment)

func (a Alignment) Compute(span float32) float32 {
	return float32(a) * span
}
func (a Alignment) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%3f", float32(a))), nil
}
func (a *Alignment) UnmarshalText(text []byte) error {
	s := strings.ToLower(string(text))
	switch s {
	case "top", "t", "left", "l":
		*a = 0
	case "bottom", "b", "right", "r":
		*a = 1
	case "middle", "m", "center", "c":
		*a = 0.5
	default:
		asFloat, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		*a = Alignment(asFloat)
	}
	return nil
}

type ClipShow int

const (
	ClipShowNone ClipShow = iota
	ClipShowLeft
	ClipShowRight
	ClipShowTop
	ClipShowBottom
)

func (a ClipShow) MarshalText() ([]byte, error) {
	switch a {
	case ClipShowNone:
		return []byte{}, nil
	case ClipShowLeft:
		return []byte("left"), nil
	case ClipShowTop:
		return []byte("top"), nil
	case ClipShowRight:
		return []byte("right"), nil
	case ClipShowBottom:
		return []byte("bottom"), nil
	}
	return nil, fmt.Errorf("invalid clip show %d", a)
}
func (c *ClipShow) UnmarshalText(text []byte) error {
	s := strings.ToLower(string(text))
	switch s {
	case "none", "":
		*c = ClipShowNone
	case "top", "t":
		*c = ClipShowTop
	case "left", "l":
		*c = ClipShowLeft
	case "bottom", "b":
		*c = ClipShowBottom
	case "right", "r":
		*c = ClipShowRight
	}
	return fmt.Errorf("invalid clip show %s", string(text))
}

type ParagraphStyles struct {
	Spacing               Amount       // extra space between glyphs
	LineSpacing           Amount       // extra space between lines
	LineHeight            Amount       // 0=calculated per line
	LineVerticalAlignment Alignment    // 0=top, 0.5=middle, 1=bottom
	Indent                Amount       // shift first line in a wrappable line
	HorizontalAlignment   Alignment    // 0=left, 0.5=middle, 1=right
	Wrap                  TextWrap     // when we are going beyond the width, how do we decide to wrap
	ParagraphPadding      AmountBounds // padding around the paragraph
}

func (s *ParagraphStyles) Override(o *ParagraphStylesOverride) *ParagraphStyles {
	if !o.HasOverride() {
		return s
	}
	return &ParagraphStyles{
		Spacing:               coalesce(o.Spacing, s.Spacing),
		LineSpacing:           coalesce(o.LineSpacing, s.LineSpacing),
		LineHeight:            coalesce(o.LineHeight, s.LineHeight),
		LineVerticalAlignment: coalesce(o.LineVerticalAlignment, s.LineVerticalAlignment),
		Indent:                coalesce(o.Indent, s.Indent),
		HorizontalAlignment:   coalesce(o.HorizontalAlignment, s.HorizontalAlignment),
		Wrap:                  coalesce(o.Wrap, s.Wrap),
		ParagraphPadding:      coalesce(o.ParagraphPadding, s.ParagraphPadding),
	}
}

type ParagraphStylesOverride struct {
	Spacing               *Amount       // extra space between glyphs
	LineSpacing           *Amount       // extra space between lines
	LineHeight            *Amount       // 0=calculated per line
	LineVerticalAlignment *Alignment    // 0=top, 0.5=middle, 1=bottom
	Indent                *Amount       // shift first line in a wrappable line
	HorizontalAlignment   *Alignment    // 0=left, 0.5=middle, 1=right
	Wrap                  *TextWrap     // when we are going beyond the width, how do we decide to wrap
	ParagraphPadding      *AmountBounds // padding around the paragraph
}

func (o *ParagraphStylesOverride) HasOverride() bool {
	if o == nil {
		return false
	}
	return o.Spacing != nil || o.LineSpacing != nil || o.LineHeight != nil || o.LineVerticalAlignment != nil ||
		o.Indent != nil || o.HorizontalAlignment != nil || o.Wrap != nil || o.ParagraphPadding != nil
}

func (o *ParagraphStylesOverride) Clone() *ParagraphStylesOverride {
	if o == nil {
		return nil
	}
	return &ParagraphStylesOverride{
		Spacing:               clone(o.Spacing),
		LineSpacing:           clone(o.LineSpacing),
		LineHeight:            clone(o.LineHeight),
		LineVerticalAlignment: clone(o.LineVerticalAlignment),
		Indent:                clone(o.Indent),
		HorizontalAlignment:   clone(o.HorizontalAlignment),
		Wrap:                  clone(o.Wrap),
		ParagraphPadding:      clone(o.ParagraphPadding),
	}
}

type ParagraphsStyles struct {
	ParagraphSpacing  Amount    // how much space between paragraphs
	VerticalAlignment Alignment // how to align all paragraphs in an area
	ClipShowX         ClipShow  // when text cannot fit in the area width, which side should we prefer to show (left, right, none)
	ClipShowY         ClipShow  // when text cannot fit in the area height, which side should we prefer to show (top, bottom, none)
}

func (s *ParagraphsStyles) Override(o *ParagraphsStylesOverride) *ParagraphsStyles {
	if !o.HasOverride() {
		return s
	}
	return &ParagraphsStyles{
		ParagraphSpacing:  coalesce(o.ParagraphSpacing, s.ParagraphSpacing),
		VerticalAlignment: coalesce(o.VerticalAlignment, s.VerticalAlignment),
		ClipShowX:         coalesce(o.ClipShowX, s.ClipShowX),
		ClipShowY:         coalesce(o.ClipShowY, s.ClipShowY),
	}
}

type ParagraphsStylesOverride struct {
	ParagraphSpacing  *Amount    // how much space between paragraphs
	VerticalAlignment *Alignment // how to align all paragraphs in an area
	ClipShowX         *ClipShow  // when text cannot fit in the area width, which side should we prefer to show (left, right, none)
	ClipShowY         *ClipShow  // when text cannot fit in the area height, which side should we prefer to show (top, bottom, none)
}

func (o *ParagraphsStylesOverride) HasOverride() bool {
	if o == nil {
		return false
	}
	return o != nil || o.ParagraphSpacing != nil || o.VerticalAlignment != nil || o.ClipShowX != nil || o.ClipShowY != nil
}

type TextStyles struct {
	ParagraphStyles
	ParagraphsStyles
	Color    Color
	Font     string
	FontSize Amount
}

func (s *TextStyles) Override(o *TextStylesOverride) *TextStyles {
	if !o.HasOverride() {
		return s
	}
	return &TextStyles{
		ParagraphStyles:  *s.ParagraphStyles.Override(o.ParagraphStylesOverride),
		ParagraphsStyles: *s.ParagraphsStyles.Override(o.ParagraphsStylesOverride),
		Color:            coalesce(o.Color, s.Color),
		Font:             coalesce(o.Font, s.Font),
		FontSize:         coalesce(o.FontSize, s.FontSize),
	}
}

type TextStylesOverride struct {
	ParagraphStylesOverride  *ParagraphStylesOverride
	ParagraphsStylesOverride *ParagraphsStylesOverride
	Color                    *Color
	Font                     *string
	FontSize                 *Amount
}

func (o *TextStylesOverride) HasOverride() bool {
	if o == nil {
		return false
	}
	return o.ParagraphStylesOverride.HasOverride() || o.ParagraphStylesOverride.HasOverride() ||
		o.Color != nil || o.Font != nil || o.FontSize != nil
}

type Paragraph struct {
	Styles *ParagraphStylesOverride
	Glyphs []Glyph
}

type Paragraphs struct {
	Styles     *ParagraphsStylesOverride
	Paragraphs []Paragraph
	MaxWidth   float32
	MaxHeight  float32
}

type RenderedText struct {
	Bounds Bounds
	Glyphs []RenderedGlyph
}

func (text RenderedText) UpdateVisibility(visibleBounds Bounds) {
	for i := range text.Glyphs {
		g := &text.Glyphs[i]
		if visibleBounds.Contains(g.Bounds) {
			g.Visibility = GlyphVisibilityVisible
		} else if visibleBounds.Intersects(g.Bounds) {
			g.Visibility = GlyphVisibilityPartial
		} else {
			g.Visibility = GlyphVisibilityInvisible
		}
	}
}

type RenderedGlyph struct {
	Tile
	Bounds     Bounds
	Color      Color
	Visibility GlyphVisibility
}

type GlyphVisibility int

const (
	GlyphVisibilityVisible GlyphVisibility = iota
	GlyphVisibilityPartial
	GlyphVisibilityInvisible
)

type Glyph interface {
	GetState(ctx *RenderContext, wrap TextWrap, prev Glyph) GlyphState
	Render(ctx *RenderContext, start Coord) RenderedGlyph
}

type GlyphState struct {
	Size        Coord
	CanBreak    bool
	ShouldBreak bool
	Empty       bool
}

func (paragraph Paragraph) GetLineHeight(ctx *RenderContext, style *ParagraphStyles, actualLineHeight float32) float32 {
	if style.LineHeight.IsZero() {
		return actualLineHeight
	} else {
		return style.LineHeight.Get(ctx.AmountContext, false)
	}
}

func (paragraph Paragraph) GetStates(ctx *RenderContext, style *ParagraphStyles) []GlyphState {
	states := make([]GlyphState, len(paragraph.Glyphs))
	var prev Glyph
	for i, g := range paragraph.Glyphs {
		states[i] = g.GetState(ctx, style.Wrap, prev)
		prev = paragraph.Glyphs[i]
	}
	return states
}

func (paragraph Paragraph) UnwrappedSize(ctx *RenderContext, scale Coord, paragraphs Paragraphs) Coord {
	style := ctx.TextStyles.ParagraphStyles.Override(paragraph.Styles)
	states := paragraph.GetStates(ctx, style)

	lineCount := 0
	lineSpacing := style.LineSpacing.Get(ctx.AmountContext, false)
	lineIndent := style.Indent.Get(ctx.AmountContext, true)
	lineWidth := lineIndent
	lineHeight := float32(0)
	size := Coord{}

	for glyphIndex := range paragraph.Glyphs {
		state := states[glyphIndex]
		if state.ShouldBreak {
			if lineWidth > size.X {
				size.X = lineWidth
			}
			if lineCount > 0 {
				size.Y += lineSpacing
			}
			size.Y += paragraph.GetLineHeight(ctx, style, lineHeight) * scale.Y
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

	padding := style.ParagraphPadding.GetBounds(ctx.AmountContext)

	size.Y += padding.Top
	size.Y += padding.Bottom
	size.X += padding.Left
	size.X += padding.Right

	return size
}

func (paragraph Paragraph) Render(ctx *RenderContext, paragraphs Paragraphs) RenderedText {
	style := ctx.TextStyles.ParagraphStyles.Override(paragraph.Styles)
	states := paragraph.GetStates(ctx, style)

	type line struct {
		width  float32
		height float32
		glyphs []int
	}

	lineIndent := style.Indent.Get(ctx.AmountContext, true)
	padding := style.ParagraphPadding.GetBounds(ctx.AmountContext)
	paddingWidth := padding.Left + padding.Right

	lines := make([]line, 0, 8)
	currentLine := line{glyphs: []int{}, width: paddingWidth + lineIndent}

	for glyphIndex := range paragraph.Glyphs {
		state := states[glyphIndex]
		nextWidth := currentLine.width + state.Size.X

		if state.ShouldBreak {
			nextLine := line{glyphs: []int{}, width: paddingWidth}
			lines = append(lines, currentLine)
			currentLine = nextLine

			if state.Empty {
				continue
			}
		} else if paragraphs.Wrap(nextWidth) {
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
	lineSpacing := style.LineSpacing.Get(ctx.AmountContext, false)

	bounds := Bounds{
		Left:   0,
		Top:    0,
		Right:  0,
		Bottom: padding.Top + padding.Bottom,
	}

	for lineIndex := range lines {
		if lineIndex > 0 {
			bounds.Bottom += lineSpacing
		}

		line := &lines[lineIndex]
		actualLineHeight := float32(0)
		for _, glyphIndex := range line.glyphs {
			state := states[glyphIndex]
			if state.Size.Y > actualLineHeight {
				actualLineHeight = state.Size.Y
			}
		}
		line.height = paragraph.GetLineHeight(ctx, style, actualLineHeight)
		bounds.Bottom += line.height
	}

	rendered := make([]RenderedGlyph, 0, len(paragraph.Glyphs))
	offsetY := padding.Top

	for lineIndex, line := range lines {
		start := Coord{
			X: style.HorizontalAlignment.Compute(paragraphs.MaxWidth - line.width),
			Y: offsetY,
		}
		if start.X < bounds.Left {
			bounds.Left = start.X
		}
		right := start.X + line.width
		if right > bounds.Right {
			bounds.Right = right
		}
		if lineIndex == 0 && len(lines) > 1 {
			start.X += lineIndent
		}
		start.X += padding.Left

		for _, glyphIndex := range line.glyphs {
			g := paragraph.Glyphs[glyphIndex]
			s := states[glyphIndex]
			start.Y = offsetY + style.LineVerticalAlignment.Compute(line.height-s.Size.Y)
			render := g.Render(ctx, start)
			if render.Bounds.Width() > 0 {
				rendered = append(rendered, render)
			}
			start.X += s.Size.X
		}

		offsetY += line.height + lineSpacing
	}

	return RenderedText{Bounds: bounds, Glyphs: rendered}
}

func (paragraph RenderedText) Translate(x, y float32) {
	for i := range paragraph.Glyphs {
		paragraph.Glyphs[i].Bounds.Translate(x, y)
	}
}

func (paragraphs Paragraphs) Wrap(lineWidth float32) bool {
	return paragraphs.MaxWidth > 0 && lineWidth > paragraphs.MaxWidth
}

func (paragraphs Paragraphs) UnwrappedSize(ctx RenderContext, scale Coord) Coord {
	style := ctx.TextStyles.ParagraphsStyles.Override(paragraphs.Styles)
	size := Coord{}
	paragraphCtx := ctx.WithParent(paragraphs.MaxWidth, paragraphs.MaxHeight)
	paragraphSpacing := style.ParagraphSpacing.Get(paragraphCtx.AmountContext, false)
	for i, paragraph := range paragraphs.Paragraphs {
		paragraphSize := paragraph.UnwrappedSize(paragraphCtx, scale, paragraphs)
		if i > 0 {
			size.Y += paragraphSpacing
		}
		size.Y += paragraphSize.Y
		if paragraphSize.X > size.X {
			size.X = paragraphSize.X
		}
	}
	return size
}

func (paragraphs Paragraphs) Render(ctx *RenderContext) RenderedText {
	style := ctx.TextStyles.ParagraphsStyles.Override(paragraphs.Styles)
	rendered := make([]RenderedText, len(paragraphs.Paragraphs))
	totalHeight := float32(0)
	totalGlyphs := 0
	paragraphCtx := ctx.WithParent(paragraphs.MaxWidth, paragraphs.MaxHeight)
	paragraphSpacing := style.ParagraphSpacing.Get(paragraphCtx.AmountContext, false)
	for i, paragraph := range paragraphs.Paragraphs {
		rendered[i] = paragraph.Render(paragraphCtx, paragraphs)
		totalHeight += rendered[i].Bounds.Height()
		if i > 0 {
			totalHeight += paragraphSpacing
		}
		totalGlyphs += len(rendered[i].Glyphs)
	}
	offsetY := style.VerticalAlignment.Compute(paragraphs.MaxHeight - totalHeight)

	switch style.ClipShowY {
	case ClipShowTop:
		if offsetY < 0 {
			offsetY = 0
		}
	case ClipShowBottom:
		if totalHeight > paragraphs.MaxHeight {
			offsetY = paragraphs.MaxHeight - totalHeight
		}
	}

	joinedBounds := Bounds{}

	for i := range rendered {
		paragraph := &rendered[i]

		offsetX := float32(0)
		switch style.ClipShowX {
		case ClipShowLeft:
			if paragraph.Bounds.Left < 0 {
				offsetX = paragraph.Bounds.Left
			}
		case ClipShowRight:
			if paragraph.Bounds.Right > paragraphs.MaxWidth {
				offsetX = paragraphs.MaxWidth - paragraph.Bounds.Right
			}
		}

		if offsetY != 0 || offsetX != 0 {
			paragraph.Translate(offsetX, offsetY)
		}

		offsetY += paragraph.Bounds.Height() + paragraphSpacing
		joinedBounds = joinedBounds.Union(paragraph.Bounds)
	}

	joined := RenderedText{
		Bounds: joinedBounds,
		Glyphs: make([]RenderedGlyph, 0, totalGlyphs),
	}
	for _, paragraph := range rendered {
		joined.Glyphs = append(joined.Glyphs, paragraph.Glyphs...)
	}

	return joined
}

type BaseGlyph struct {
	Text  rune
	Font  string
	Size  Amount
	Color Color

	initialized bool
	font        *Font
	fontRune    *FontRune
}

var _ Glyph = &BaseGlyph{}

func (g *BaseGlyph) init(ctx *RenderContext) {
	if !g.initialized {
		font := ctx.Theme.Fonts[g.Font]
		if font == nil {
			font = ctx.Theme.Fonts[ctx.TextStyles.Font]
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

func (g BaseGlyph) getColor(ctx *RenderContext) Color {
	if g.Color.IsZero() {
		return ctx.TextStyles.Color
	} else {
		return g.Color
	}
}

func (g BaseGlyph) getSize(ctx *RenderContext) float32 {
	if g.Size.Value == 0 {
		return ctx.TextStyles.FontSize.Get(ctx.AmountContext, true)
	} else {
		return g.Size.Get(ctx.AmountContext, true)
	}
}

func (g *BaseGlyph) GetState(ctx *RenderContext, wrap TextWrap, prev Glyph) GlyphState {
	space := unicode.IsSpace(g.Text)
	state := GlyphState{
		CanBreak:    wrap == TextWrapChar || (wrap == TextWrapWord && space),
		ShouldBreak: g.Text == '\r' || g.Text == '\n',
		Empty:       space,
	}

	g.init(ctx)
	if g.font == nil {
		state.Empty = true
		return state
	}

	offset := float32(0)
	if prev != nil {
		if prevGlyph, ok := prev.(*BaseGlyph); ok {
			offset = g.font.GetKerning(prevGlyph.Text, g.Text)
		}
	}

	size := g.getSize(ctx)
	state.Size.X = (size * g.fontRune.Width) + offset*size
	state.Size.Y = size * g.font.LineHeight

	return state
}

func (g *BaseGlyph) Render(ctx *RenderContext, topLeft Coord) RenderedGlyph {
	g.init(ctx)
	if g.font == nil {
		return RenderedGlyph{}
	}
	color := g.getColor(ctx)
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
// - new paragraph = {p} or {p:reset} to reset paragraph style
// - set kerning to paragraph = {k:5}
// - set horizontal alignment in paragraph = {h:0.5}
// - set vertical alignment within a line in paragraph = {v:0}
// - set wrap in paragraph = {w:none} or {w:word} or {w:char}
// - set line spacing in paragraph = {ls:10}
// - set line height in paragraph = {lh:10} or {lh:120%}
// - set indent = {i:5em}
// - set padding in paragraph = {pa:10} or {pl:10%} or {pr:0} or {pb} or {pt:}
// - set spacing between paragraphs = {ps:10}
// - set vertical alignment of paragraphs in area = {pv:0.5}
// - set max height in area = {mh:100}
// - set max width in area = {mw:100}
// - set clip show x in area = {cx} aka {cx:} aka {cx:none} or {cx:left} or {cx:right}
// - set clip show y in area = {cy} aka {cy:} aka {cy:none} or {cy:top} or {cy:bottom}
//
// Ex: "Hello {f:roboto}World{f:} {s:20}HOW{s:} {c:#00F}blue {c:red}red {c:} \{"
func TextToParagraphs(text string) (paragraphs Paragraphs, err error) {
	paragraph := Paragraph{}
	glyph := BaseGlyph{}

	readAmount := func(s string, required bool) *Amount {
		if s == "" {
			if required {
				return &Amount{}
			}
			return nil
		}
		amt := Amount{}
		err = amt.UnmarshalText([]byte(s))
		return &amt
	}
	readFloat := func(s string) float32 {
		f := float64(0)
		if s != "" {
			f, err = strconv.ParseFloat(s, 32)
		}
		return float32(f)
	}
	readAlignment := func(s string) *Alignment {
		if s == "" {
			return nil
		}
		align := Alignment(0)
		err = align.UnmarshalText([]byte(s))
		return &align
	}
	readClipShow := func(s string) *ClipShow {
		if s == "" {
			return nil
		}
		show := ClipShow(0)
		err = show.UnmarshalText([]byte(s))
		return &show
	}
	readWrap := func(s string) *TextWrap {
		if s == "" {
			return nil
		}
		wrap := TextWrap("")
		err = wrap.UnmarshalText([]byte(s))
		return &wrap
	}
	readColor := func(s string) Color {
		color := Color{}
		if s != "" {
			err = color.UnmarshalText([]byte(s))
		}
		return color
	}
	getParagraphStyles := func() *ParagraphStylesOverride {
		if paragraph.Styles == nil {
			paragraph.Styles = &ParagraphStylesOverride{}
		}
		return paragraph.Styles
	}
	getParagraphPadding := func() *AmountBounds {
		style := getParagraphStyles()
		if style.ParagraphPadding == nil {
			style.ParagraphPadding = &AmountBounds{}
		}
		return style.ParagraphPadding
	}
	getParagraphsStyles := func() *ParagraphsStylesOverride {
		if paragraphs.Styles == nil {
			paragraphs.Styles = &ParagraphsStylesOverride{}
		}
		return paragraphs.Styles
	}

	pieces := TextFormatRegex.FindAllStringSubmatch(text, -1)
	for _, piece := range pieces {
		runes := piece[0]
		command := piece[1]
		value := piece[2]

		switch command {
		case "": // Add glyph to paragraph
			copy := glyph
			copy.Text, _ = utf8.DecodeRuneInString(runes)
			paragraph.Glyphs = append(paragraph.Glyphs, &copy)
		case "f": // Glyph commands
			glyph.Font = value
		case "s":
			glyph.Size = *readAmount(value, true)
		case "c":
			glyph.Color = readColor(value)
		case "k": // Paragraph commands
			getParagraphStyles().Spacing = readAmount(value, false)
		case "h":
			getParagraphStyles().HorizontalAlignment = readAlignment(value)
		case "v":
			getParagraphStyles().LineVerticalAlignment = readAlignment(value)
		case "w":
			getParagraphStyles().Wrap = readWrap(value)
		case "ls":
			getParagraphStyles().LineSpacing = readAmount(value, false)
		case "lh":
			getParagraphStyles().LineHeight = readAmount(value, false)
		case "pa":
			getParagraphPadding().SetAmount(*readAmount(value, true))
		case "pl":
			getParagraphPadding().Left = *readAmount(value, true)
		case "pt":
			getParagraphPadding().Top = *readAmount(value, true)
		case "pr":
			getParagraphPadding().Right = *readAmount(value, true)
		case "pb":
			getParagraphPadding().Bottom = *readAmount(value, true)
		case "i":
			getParagraphStyles().Indent = readAmount(value, false)
		case "p": // Paragraphs commands
			paragraphs.Paragraphs = append(paragraphs.Paragraphs, paragraph)
			paragraph.Glyphs = make([]Glyph, 0)
			if value != "reset" {
				paragraph.Styles = paragraph.Styles.Clone()
			}
		case "ps":
			getParagraphsStyles().ParagraphSpacing = readAmount(value, false)
		case "pv":
			getParagraphsStyles().VerticalAlignment = readAlignment(value)
		case "mw":
			paragraphs.MaxWidth = readFloat(value)
		case "mh":
			paragraphs.MaxHeight = readFloat(value)
		case "cx":
			getParagraphsStyles().ClipShowX = readClipShow(value)
		case "cy":
			getParagraphsStyles().ClipShowY = readClipShow(value)
		}
		if err != nil {
			return
		}
	}

	if len(paragraph.Glyphs) > 0 {
		paragraphs.Paragraphs = append(paragraphs.Paragraphs, paragraph)
	}

	return
}

func MustTextToParagraphs(text string) Paragraphs {
	paragraphs, err := TextToParagraphs(text)
	if err != nil {
		panic(err)
	}
	return paragraphs
}

func TextToVisual(text string) (*VisualText, error) {
	paragraphs, err := TextToParagraphs(text)
	if err != nil {
		return nil, err
	}
	return &VisualText{Paragraphs: paragraphs}, nil
}

func MustTextToVisual(text string) *VisualText {
	return &VisualText{Paragraphs: MustTextToParagraphs(text)}
}
