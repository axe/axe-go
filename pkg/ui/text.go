package ui

import (
	"encoding"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/axe/axe-go/pkg/color"
	"github.com/axe/axe-go/pkg/ease"
	"github.com/axe/axe-go/pkg/id"
	"github.com/axe/axe-go/pkg/util"
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
		multiplier := float64(1)
		if strings.HasSuffix(s, "%") {
			s = strings.TrimSuffix(s, "%")
			multiplier = 0.01
		}
		asFloat, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		*a = Alignment(asFloat * multiplier)
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
	Spacing               Amount       // extra space between glyphs TODO
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
		Spacing:               util.Coalesce(o.Spacing, s.Spacing),
		LineSpacing:           util.Coalesce(o.LineSpacing, s.LineSpacing),
		LineHeight:            util.Coalesce(o.LineHeight, s.LineHeight),
		LineVerticalAlignment: util.Coalesce(o.LineVerticalAlignment, s.LineVerticalAlignment),
		Indent:                util.Coalesce(o.Indent, s.Indent),
		HorizontalAlignment:   util.Coalesce(o.HorizontalAlignment, s.HorizontalAlignment),
		Wrap:                  util.Coalesce(o.Wrap, s.Wrap),
		ParagraphPadding:      util.Coalesce(o.ParagraphPadding, s.ParagraphPadding),
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
		Spacing:               util.Clone(o.Spacing),
		LineSpacing:           util.Clone(o.LineSpacing),
		LineHeight:            util.Clone(o.LineHeight),
		LineVerticalAlignment: util.Clone(o.LineVerticalAlignment),
		Indent:                util.Clone(o.Indent),
		HorizontalAlignment:   util.Clone(o.HorizontalAlignment),
		Wrap:                  util.Clone(o.Wrap),
		ParagraphPadding:      util.Clone(o.ParagraphPadding),
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
		ParagraphSpacing:  util.Coalesce(o.ParagraphSpacing, s.ParagraphSpacing),
		VerticalAlignment: util.Coalesce(o.VerticalAlignment, s.VerticalAlignment),
		ClipShowX:         util.Coalesce(o.ClipShowX, s.ClipShowX),
		ClipShowY:         util.Coalesce(o.ClipShowY, s.ClipShowY),
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
	return o.ParagraphSpacing != nil || o.VerticalAlignment != nil || o.ClipShowX != nil || o.ClipShowY != nil
}

type TextStyles struct {
	ParagraphStyles
	ParagraphsStyles
	Color    color.Able
	Font     id.Identifier
	FontSize Amount
}

func (s *TextStyles) Override(o *TextStylesOverride) *TextStyles {
	if !o.HasOverride() {
		return s
	}
	return &TextStyles{
		ParagraphStyles:  *s.ParagraphStyles.Override(o.ParagraphStylesOverride),
		ParagraphsStyles: *s.ParagraphsStyles.Override(o.ParagraphsStylesOverride),
		Color:            util.Coalesce(o.Color, s.Color),
		Font:             util.Coalesce(o.Font, s.Font),
		FontSize:         util.Coalesce(o.FontSize, s.FontSize),
	}
}

type TextStylesOverride struct {
	ParagraphStylesOverride  *ParagraphStylesOverride
	ParagraphsStylesOverride *ParagraphsStylesOverride
	Color                    *color.Able
	Font                     *id.Identifier
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
	KeepEmpty  bool
}

type RenderedText struct {
	Bounds     Bounds
	Glyphs     []RenderedGlyph
	Paragraphs Paragraphs
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

func (text RenderedText) CountVisible(threshold *GlyphVisibility) int {
	visible := 0
	for i := range text.Glyphs {
		g := &text.Glyphs[i]
		if g.Empty || (threshold != nil && g.Visibility > *threshold) {
			continue
		}
		visible++
	}
	return visible
}

func (text RenderedText) Closest(x, y float32) int {
	closest := -1
	closestDistanceSq := float32(0)
	for index, glyph := range text.Glyphs {
		cx, cy := glyph.Bounds.Center()
		distanceSq := LengthSq(cx-x, cy-y)
		if closest == -1 || distanceSq < closestDistanceSq {
			closest = index
			closestDistanceSq = distanceSq
		}
	}
	return closest
}
func (text RenderedText) ClosestByLine(x, y float32) int {
	closest := -1
	closestY := float32(math.MaxFloat32)
	closestX := float32(math.MaxFloat32)
	closestLine := -1
	for index, glyph := range text.Glyphs {
		cx, cy := glyph.Bounds.Closest(x, y)
		dy := util.Abs(cy - y)
		if dy <= closestY || closestLine == glyph.Line {
			dx := util.Abs(cx - x)
			if dx < closestX || closestLine != glyph.Line {
				closest = index
				closestLine = glyph.Line
				closestY = dy
				closestX = dx
			}
		}
	}
	return closest
}

type RenderedGlyph struct {
	Tile
	Bounds         Bounds
	Color          color.Color
	Visibility     GlyphVisibility
	Line           int
	Column         int
	Word           int
	Paragraph      int
	ParagraphIndex int
	Index          int
	Empty          bool
}

func (g RenderedGlyph) Write(quad []Vertex) {
	quad[0] = Vertex{X: g.Bounds.Left, Y: g.Bounds.Top, Tex: g.Coord(0, 0), HasCoord: true, Color: g.Color, HasColor: true}
	quad[1] = Vertex{X: g.Bounds.Right, Y: g.Bounds.Top, Tex: g.Coord(1, 0), HasCoord: true, Color: g.Color, HasColor: true}
	quad[2] = Vertex{X: g.Bounds.Right, Y: g.Bounds.Bottom, Tex: g.Coord(1, 1), HasCoord: true, Color: g.Color, HasColor: true}
	quad[3] = Vertex{X: g.Bounds.Left, Y: g.Bounds.Bottom, Tex: g.Coord(0, 1), HasCoord: true, Color: g.Color, HasColor: true}
}

type GlyphVisibility int

const (
	GlyphVisibilityVisible GlyphVisibility = iota
	GlyphVisibilityPartial
	GlyphVisibilityInvisible
)

type Glyph interface {
	GetState(ctx *RenderContext, wrap TextWrap, prev Glyph) GlyphState
	Render(ctx *RenderContext, start Coord, base *Base) RenderedGlyph
	String() string
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

const (
	MeasureWidthRoundingError  = 0.001
	MeasureHeightRoundingError = 0.001
)

func (paragraph Paragraph) MinWidth(ctx *RenderContext, keepEmpty bool) float32 {
	style := ctx.TextStyles.ParagraphStyles.Override(paragraph.Styles)
	states := paragraph.GetStates(ctx, style)

	lineIndent := style.Indent.Get(ctx.AmountContext, true)
	kerning := style.Spacing.Get(ctx.AmountContext, true)
	lineWidth := lineIndent
	maxWidth := float32(0)

	for _, state := range states {
		if state.ShouldBreak || state.CanBreak {
			maxWidth = util.Max(maxWidth, lineWidth-kerning)
			lineWidth = 0

			if state.Empty && !keepEmpty {
				continue
			}
		}
		lineWidth += state.Size.X + kerning
	}
	maxWidth = util.Max(maxWidth, lineWidth-kerning)

	minWidth := maxWidth
	minWidth += style.ParagraphPadding.Left.Get(ctx.AmountContext, true)
	minWidth += style.ParagraphPadding.Right.Get(ctx.AmountContext, true)
	minWidth += MeasureWidthRoundingError

	return minWidth
}

func (paragraph Paragraph) String() string {
	out := strings.Builder{}
	out.Grow(len(paragraph.Glyphs))
	for _, glyph := range paragraph.Glyphs {
		out.WriteString(glyph.String())
	}
	return out.String()
}

type paragraphLine struct {
	start, endExclusive int
	width, height       float32
	indent              float32
}

type paragraphLines struct {
	lines       []paragraphLine
	style       *ParagraphStyles
	states      []GlyphState
	padding     Bounds
	lineSpacing float32
	totalHeight float32
	maxWidth    float32
	kerning     float32
}

func (paragraph Paragraph) getLines(ctx *RenderContext, paragraphs Paragraphs) paragraphLines {
	lines := make([]paragraphLine, 0, 8)

	style := ctx.TextStyles.ParagraphStyles.Override(paragraph.Styles)
	states := paragraph.GetStates(ctx, style)

	lineSpacing := style.LineSpacing.Get(ctx.AmountContext, false)
	lineIndent := style.Indent.Get(ctx.AmountContext, true)
	kerning := style.Spacing.Get(ctx.AmountContext, true)
	padding := style.ParagraphPadding.GetBounds(ctx.AmountContext)
	paddingWidth := padding.Left + padding.Right

	lastBreak := -1
	glyphLast := len(paragraph.Glyphs) - 1

	currentLine := paragraphLine{width: paddingWidth + lineIndent, indent: lineIndent}

	endLine := func(nextLineWidth float32, nextLineStart int) {
		currentLine.endExclusive = nextLineStart
		currentLine.width -= nextLineWidth
		if !paragraphs.KeepEmpty {
			for i := currentLine.endExclusive - 1; i >= currentLine.start; i-- {
				if states[i].Empty {
					currentLine.width -= states[i].Size.X + kerning
					currentLine.endExclusive--
				} else {
					break
				}
			}
		}

		lines = append(lines, currentLine)
		currentLine = paragraphLine{
			start:        nextLineStart,
			endExclusive: nextLineStart,
			width:        paddingWidth + nextLineWidth,
		}
		lastBreak = -1
	}

	for glyphIndex := 0; glyphIndex <= glyphLast; glyphIndex++ {
		state := states[glyphIndex]

		if state.ShouldBreak {
			endLine(0, glyphIndex)

			if state.Empty && !paragraphs.KeepEmpty {
				continue
			}
		} else if paragraphs.Wrap(currentLine.width+state.Size.X) && lastBreak != -1 {
			// Wrap everything after lastBreak
			endLineAt := lastBreak + 1
			// Glyphs moving to next line
			nextLineWidth := float32(0)
			for k := endLineAt; k < glyphIndex; k++ {
				nextLineWidth += states[k].Size.X + kerning
			}

			// End line and start next
			endLine(nextLineWidth, endLineAt)
		}

		currentLine.width += state.Size.X + kerning
		currentLine.endExclusive++
		if state.CanBreak {
			lastBreak = glyphIndex
		}
	}

	if currentLine.start != currentLine.endExclusive {
		endLine(0, glyphLast+1)
	}

	totalHeight := float32(0)
	maxWidth := float32(0)

	for lineIndex := range lines {
		line := &lines[lineIndex]
		actualLineHeight := float32(0)
		for k := line.start; k < line.endExclusive; k++ {
			actualLineHeight = util.Max(actualLineHeight, states[k].Size.Y)
		}
		line.height = paragraph.GetLineHeight(ctx, style, actualLineHeight)
		totalHeight += line.height
		maxWidth = util.Max(maxWidth, line.width)
	}

	totalHeight += lineSpacing * float32(len(lines)-1)
	totalHeight += padding.Top + padding.Bottom

	return paragraphLines{
		lines:       lines,
		style:       style,
		states:      states,
		padding:     padding,
		kerning:     kerning,
		lineSpacing: lineSpacing,
		maxWidth:    maxWidth,
		totalHeight: totalHeight,
	}
}

func (paragraph Paragraph) Measure(ctx *RenderContext, paragraphs Paragraphs) Coord {
	size := Coord{}

	lines := paragraph.getLines(ctx, paragraphs)
	size.X = lines.maxWidth + MeasureWidthRoundingError
	size.Y = lines.totalHeight + MeasureHeightRoundingError

	return size
}

func (paragraph Paragraph) Render(ctx *RenderContext, paragraphs Paragraphs, b *Base) RenderedText {
	lines := paragraph.getLines(ctx, paragraphs)

	bounds := Bounds{
		Left:   0,
		Top:    0,
		Right:  0,
		Bottom: lines.totalHeight,
	}

	rendered := make([]RenderedGlyph, 0, len(paragraph.Glyphs))
	offsetY := lines.padding.Top

	maxWidth := util.Max(0, paragraphs.MaxWidth)

	wordCount := 0
	for lineIndex, line := range lines.lines {
		start := Coord{
			X: lines.style.HorizontalAlignment.Compute(maxWidth - line.width),
			Y: offsetY,
		}
		bounds.Left = util.Min(bounds.Left, start.X)
		bounds.Right = util.Max(bounds.Right, start.X+line.width)

		start.X += line.indent + lines.padding.Left
		for k := line.start; k < line.endExclusive; k++ {
			g := paragraph.Glyphs[k]
			s := lines.states[k]
			start.Y = offsetY + lines.style.LineVerticalAlignment.Compute(line.height-s.Size.Y)
			render := g.Render(ctx, start, b)
			render.Line = lineIndex
			render.Column = k - line.start
			render.Word = wordCount
			render.Index = k
			render.ParagraphIndex = k
			render.Empty = s.Empty || render.Bounds.Width() <= 0
			if !render.Empty || paragraphs.KeepEmpty {
				rendered = append(rendered, render)
			} else if s.CanBreak || s.ShouldBreak {
				wordCount++
			}

			start.X += s.Size.X + lines.kerning
		}

		offsetY += line.height + lines.lineSpacing
	}

	return RenderedText{Bounds: bounds, Glyphs: rendered, Paragraphs: paragraphs}
}

func (paragraph *RenderedText) Translate(x, y float32) {
	paragraph.Bounds.Translate(x, y)
	for i := range paragraph.Glyphs {
		paragraph.Glyphs[i].Bounds.Translate(x, y)
	}
}

func (paragraphs Paragraphs) String() string {
	return paragraphs.ToString("\n")
}

func (paragraphs Paragraphs) ToString(paragraphSeparator string) string {
	paras := make([]string, len(paragraphs.Paragraphs))
	for i, para := range paragraphs.Paragraphs {
		paras[i] = para.String()
	}
	return strings.Join(paras, paragraphSeparator)
}

func (paragraphs Paragraphs) Wrap(lineWidth float32) bool {
	return paragraphs.MaxWidth > 0 && lineWidth > paragraphs.MaxWidth-MeasureWidthRoundingError
}

func (paragraphs Paragraphs) MinWidth(ctx *RenderContext) float32 {
	paragraphCtx := ctx.WithParent(util.Max(0, paragraphs.MaxWidth), paragraphs.MaxHeight)
	minWidth := float32(0)
	for _, para := range paragraphs.Paragraphs {
		minWidth = util.Max(minWidth, para.MinWidth(paragraphCtx, paragraphs.KeepEmpty))
	}
	return minWidth
}

func (paragraphs Paragraphs) Measure(ctx *RenderContext) Coord {
	style := ctx.TextStyles.ParagraphsStyles.Override(paragraphs.Styles)
	size := Coord{}
	paragraphCtx := ctx.WithParent(util.Max(0, paragraphs.MaxWidth), paragraphs.MaxHeight)
	paragraphSpacing := style.ParagraphSpacing.Get(paragraphCtx.AmountContext, false)
	for _, paragraph := range paragraphs.Paragraphs {
		paragraphSize := paragraph.Measure(paragraphCtx, paragraphs)
		size.Y += paragraphSize.Y
		size.X = util.Max(size.X, paragraphSize.X)
	}
	size.Y += paragraphSpacing * float32(len(paragraphs.Paragraphs)-1)
	return size
}

func (paragraphs Paragraphs) Render(ctx *RenderContext, b *Base) RenderedText {
	style := ctx.TextStyles.ParagraphsStyles.Override(paragraphs.Styles)
	rendered := make([]RenderedText, len(paragraphs.Paragraphs))
	totalHeight := float32(0)
	totalGlyphs := 0
	paragraphCtx := ctx.WithParent(util.Max(0, paragraphs.MaxWidth), paragraphs.MaxHeight)
	paragraphSpacing := style.ParagraphSpacing.Get(paragraphCtx.AmountContext, false)
	for i, paragraph := range paragraphs.Paragraphs {
		rendered[i] = paragraph.Render(paragraphCtx, paragraphs, b)
		totalHeight += rendered[i].Bounds.Height()
		totalGlyphs += len(rendered[i].Glyphs)
	}
	totalHeight += paragraphSpacing * float32(len(paragraphs.Paragraphs)-1)
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
				offsetX = -paragraph.Bounds.Left
			}
		case ClipShowRight:
			if paragraphs.MaxWidth > 0 && paragraph.Bounds.Right > paragraphs.MaxWidth {
				offsetX = paragraph.Bounds.Right - paragraphs.MaxWidth
			}
		}

		if offsetY != 0 || offsetX != 0 {
			paragraph.Translate(offsetX, offsetY)
		}

		offsetY += paragraph.Bounds.Height() + paragraphSpacing
		joinedBounds = joinedBounds.Union(paragraph.Bounds)
	}

	joined := RenderedText{
		Bounds:     joinedBounds,
		Glyphs:     make([]RenderedGlyph, 0, totalGlyphs),
		Paragraphs: paragraphs,
	}
	lineOffset := 0
	wordOffset := 0
	indexOffset := 0
	for paragraphIndex, paragraph := range rendered {
		n := len(paragraph.Glyphs) - 1
		if n == -1 {
			continue
		}
		for i := 0; i <= n; i++ {
			g := &paragraph.Glyphs[i]
			g.Line += lineOffset
			g.Word += wordOffset
			g.Index += indexOffset
			g.Paragraph = paragraphIndex
		}
		joined.Glyphs = append(joined.Glyphs, paragraph.Glyphs...)
		last := paragraph.Glyphs[n]
		lineOffset = last.Line + 1
		wordOffset = last.Word + 1
		indexOffset = last.Index + 1
	}

	return joined
}

type BaseGlyph struct {
	Text  rune
	Font  id.Identifier
	Size  Amount
	Color color.Color

	initialized bool
	font        *Font
	fontRune    *FontRune
}

var _ Glyph = &BaseGlyph{}

func (g *BaseGlyph) init(ctx *RenderContext) {
	if !g.initialized {
		font := ctx.Theme.Fonts.Get(g.Font)
		if font == nil {
			font = ctx.Theme.Fonts.Get(ctx.TextStyles.Font)
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

func (g BaseGlyph) getColor(ctx *RenderContext, b *Base) color.Color {
	if g.Color.IsZero() {
		color, _ := color.Get(ctx.TextStyles.Color, b)
		return color
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

func (g *BaseGlyph) Render(ctx *RenderContext, topLeft Coord, b *Base) RenderedGlyph {
	g.init(ctx)
	if g.font == nil {
		return RenderedGlyph{}
	}
	color := g.getColor(ctx, b)
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

func (g BaseGlyph) String() string {
	return string([]rune{g.Text})
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
	readColor := func(s string) color.Color {
		color := color.Color{}
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
			glyph.Font = id.Get(value)
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

type TextAnimation interface {
	Init(base *Base)
	Update(base *Base, animationTime float32, update Update) Dirty
	IsDone(base *Base, animationTime float32) bool
	Render(base *Base, animationTime float32, bounds Bounds, ctx *RenderContext, out *VertexBuffer)
}

type TextAnimationFactory interface {
	GetAnimation(text *RenderedText) TextAnimation
}

type BasicTextAnimationKind int

const (
	BasicTextAnimationKindChar BasicTextAnimationKind = iota
	BasicTextAnimationKindWord
	BasicTextAnimationKindLine
	BasicTextAnimationKindParagraph
	BasicTextAnimationKindColumn
)

func (k BasicTextAnimationKind) Get(g *RenderedGlyph) int {
	switch k {
	case BasicTextAnimationKindChar:
		return g.Index
	case BasicTextAnimationKindColumn:
		return g.Column
	case BasicTextAnimationKindLine:
		return g.Line
	case BasicTextAnimationKindParagraph:
		return g.Paragraph
	case BasicTextAnimationKindWord:
		return g.Word
	}
	return -1
}

type BasicTextAnimationState struct {
	Start, End int
	Min, Max   int
	Total      int
	Duration   float32
}

type BasicTextAnimationSettings struct {
	// How to break up the text into pieces to animate
	Kind BasicTextAnimationKind
	// The frames to animate over the given duration for a piece
	Frames []BasicAnimationFrame
	// How long to animate a piece
	Duration float32
	// How long between starting each piece animation
	Delay float32
	// Optional easing function for determining the frame
	Easing ease.Easing
	// The index of the glyph to start at
	Start int
}

func (s BasicTextAnimationSettings) GetState(text *RenderedText, endExclusive int) BasicTextAnimationState {
	glyph := &text.Glyphs[s.Start]
	glyphValue := s.Kind.Get(glyph)
	min := glyphValue
	max := glyphValue
	for i := s.Start + 1; i < endExclusive; i++ {
		glyph = &text.Glyphs[i]
		glyphValue := s.Kind.Get(glyph)
		min = util.Min(min, glyphValue)
		max = util.Max(max, glyphValue)
	}
	total := max - min + 1

	return BasicTextAnimationState{
		Start:    s.Start,
		End:      endExclusive,
		Min:      min,
		Max:      max,
		Total:    total,
		Duration: float32(total-1)*s.Delay + s.Duration,
	}
}

func (s BasicTextAnimationSettings) GetFrames(time float32) (first int, delta float32) {
	first = len(s.Frames) - 2
	for first > 0 && s.Frames[first].Time > time {
		first--
	}
	delta = util.Delta(s.Frames[first].Time, s.Frames[first+1].Time, time)
	return
}

type BasicTextAnimation struct {
	Settings []BasicTextAnimationSettings
	Easing   ease.Easing

	previousAnimationTime float32
	text                  *RenderedText
	duration              float32
	states                []BasicTextAnimationState
}

func (a *BasicTextAnimation) Init(base *Base) {
	a.duration = 0
	a.states = a.GetStates(a.text)
	for _, state := range a.states {
		a.duration += state.Duration
	}
}

func (a *BasicTextAnimation) Update(b *Base, animationTime float32, update Update) Dirty {
	prevTime := a.EasedTime(a.previousAnimationTime)
	nextTime := a.EasedTime(animationTime)

	a.previousAnimationTime = animationTime

	prevState, prevValue, prevTimeInValue := a.ValueAt(prevTime)
	nextState, nextValue, _ := a.ValueAt(nextTime)

	if prevState != nextState || prevValue != nextValue || prevTimeInValue <= a.Settings[prevState].Duration {
		return DirtyVisual
	}

	return DirtyNone
}

func (a *BasicTextAnimation) IsDone(base *Base, animationTime float32) bool {
	return animationTime > a.duration
}

func (a *BasicTextAnimation) Render(b *Base, animationTime float32, bounds Bounds, ctx *RenderContext, out *VertexBuffer) {
	animatingState, animatingValue, animatingTime := a.ValueAt(animationTime)

	animating := map[int]*animateGlyphs{}

	for stateIndex, state := range a.states {
		if stateIndex > animatingState {
			break
		}

		out.ReserveQuads(state.End - state.Start)

		if stateIndex < animatingState {
			for i := state.Start; i < state.End; i++ {
				glyph := a.text.Glyphs[i]
				glyph.Write(out.GetReservedQuad())
			}
		} else {
			settings := &a.Settings[stateIndex]
			for i := state.Start; i < state.End; i++ {
				glyph := &a.text.Glyphs[i]
				value := settings.Kind.Get(glyph)
				if value < animatingValue {
					glyph.Write(out.GetReservedQuad())
				} else {
					timeInValue := animatingTime - float32(value-animatingValue)*settings.Delay
					if timeInValue >= 0 {
						existing := animating[value]
						if existing == nil {
							existing = &animateGlyphs{
								delta:    timeInValue / settings.Duration,
								settings: settings,
							}
							animating[value] = existing
						}
						existing.add(glyph)
					}
				}
			}
		}
	}

	for _, animate := range animating {
		animate.update(ctx)

		for _, g := range animate.glyphs {
			quad := out.GetReservedQuad()
			g.Write(quad)
			quad[0].X, quad[0].Y = animate.transform.Transform(quad[0].X, quad[0].Y)
			quad[1].X, quad[1].Y = animate.transform.Transform(quad[1].X, quad[1].Y)
			quad[2].X, quad[2].Y = animate.transform.Transform(quad[2].X, quad[2].Y)
			quad[3].X, quad[3].Y = animate.transform.Transform(quad[3].X, quad[3].Y)

			if animate.color != nil {
				quad[0].Color = animate.color(quad[0].Color)
				quad[0].HasColor = true
				quad[1].Color = animate.color(quad[1].Color)
				quad[1].HasColor = true
				quad[2].Color = animate.color(quad[2].Color)
				quad[2].HasColor = true
				quad[3].Color = animate.color(quad[3].Color)
				quad[3].HasColor = true
			}
			if animate.transparency != 0 {
				alphaMultiplier := 1 - animate.transparency
				quad[0].Color.A *= alphaMultiplier
				quad[0].HasColor = true
				quad[1].Color.A *= alphaMultiplier
				quad[1].HasColor = true
				quad[2].Color.A *= alphaMultiplier
				quad[2].HasColor = true
				quad[3].Color.A *= alphaMultiplier
				quad[3].HasColor = true
			}
		}
	}
}

func (a *BasicTextAnimation) EasedTime(time float32) float32 {
	return ease.Get(time/a.duration, a.Easing) * a.duration
}

func (a *BasicTextAnimation) ValueAt(time float32) (state int, value int, timeInValue float32) {
	for stateIndex, s := range a.states {
		if time <= s.Duration {
			settings := a.Settings[stateIndex]
			stateTime := ease.Get(time/s.Duration, settings.Easing) * s.Duration
			state = stateIndex
			value = util.Max(0, int((stateTime-settings.Duration+settings.Delay)/settings.Delay))
			timeInValue = stateTime - float32(value)*settings.Delay
			return
		} else {
			time -= s.Duration
		}
	}

	lastIndex := len(a.states) - 1
	last := a.states[lastIndex]
	return lastIndex, last.Max, a.Settings[lastIndex].Duration
}

func (a *BasicTextAnimation) GetStates(text *RenderedText) []BasicTextAnimationState {
	settingsCount := len(a.Settings)
	states := make([]BasicTextAnimationState, len(a.Settings))
	for settingIndex, setting := range a.Settings {
		end := len(text.Glyphs)
		next := settingIndex + 1
		if next < settingsCount {
			end = a.Settings[next].Start
		}
		states[settingIndex] = setting.GetState(text, end)
	}
	return states
}

func (a BasicTextAnimation) GetAnimation(text *RenderedText) TextAnimation {
	return &BasicTextAnimation{
		Settings: a.Settings,
		Easing:   a.Easing,
		text:     text,
	}
}

type animateGlyphs struct {
	delta        float32
	bounds       Bounds
	transform    Transform
	color        color.Modify
	transparency float32
	glyphs       []*RenderedGlyph
	settings     *BasicTextAnimationSettings
}

func (ag *animateGlyphs) add(g *RenderedGlyph) {
	ag.bounds = ag.bounds.Union(g.Bounds)
	ag.glyphs = append(ag.glyphs, g)
}

func (ag *animateGlyphs) update(ctx *RenderContext) {
	frameStart, frameDelta := ag.settings.GetFrames(ag.delta)
	start := ag.settings.Frames[frameStart]
	end := ag.settings.Frames[frameStart+1]
	delta := ease.Get(frameDelta, start.Easing)

	animateCtx := ctx.WithBounds(ag.bounds)

	inter := start.Lerp(end, delta, animateCtx.AmountContext, ag.bounds.Left, ag.bounds.Top)
	ag.transform = inter.Transform()
	ag.color = inter.Color
	ag.transparency = inter.Transparency
}
