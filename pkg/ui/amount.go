package ui

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Unit int

const (
	UnitConstant     Unit = iota // "" px
	UnitFont                     // %f f rem %rem em %em
	UnitParent                   // % p %p
	UnitParentWidth              // %w w %pw pw
	UnitParentHeight             // %h h %ph ph
	UnitView                     // %v v
	UnitViewWidth                // %vw vw
	UnitViewHeight               // %vh vh
	UnitWindow                   // %n n
	UnitWindowWidth              // %nw nw
	UnitWindowHeight             // %nh nh
	UnitScreen                   // %s s
	UnitScreenWidth              // %sw sw
	UnitScreenHeight             // %sh sh
)

func (a Unit) Get(value float32, ctx *AmountContext, widthRelative bool) float32 {
	return ctx.GetScalar(a, widthRelative) * value
}

func (a Unit) SupportsPercent() bool {
	return a != UnitConstant
}

var (
	suffixToUnit = map[string]Unit{}
	unitToSuffix = map[Unit]string{}
)

func AddSuffixes(unit Unit, suffixes ...string) {
	unitToSuffix[unit] = suffixes[0]
	for _, s := range suffixes {
		suffixToUnit[strings.ToLower(s)] = unit
	}
}

func init() {
	AddSuffixes(UnitConstant, "", "px")
	AddSuffixes(UnitFont, "%f", "f", "rem", "%rem", "em", "%em")
	AddSuffixes(UnitParent, "%", "%p", "p")
	AddSuffixes(UnitParentWidth, "%w", "w", "%pw", "pw")
	AddSuffixes(UnitParentHeight, "%h", "h", "%ph", "ph")
	AddSuffixes(UnitView, "%v", "v")
	AddSuffixes(UnitViewWidth, "%vw", "vw")
	AddSuffixes(UnitViewHeight, "%vh", "vh")
	AddSuffixes(UnitWindow, "%n", "n")
	AddSuffixes(UnitWindowWidth, "%nw", "nw")
	AddSuffixes(UnitWindowHeight, "%nh", "nh")
	AddSuffixes(UnitScreen, "%s", "s")
	AddSuffixes(UnitScreenWidth, "%sw", "sw")
	AddSuffixes(UnitScreenHeight, "%sh", "sh")
}

// A unit amount
type Amount struct {
	Value float32
	Unit  Unit
}

func (a Amount) IsZero() bool {
	return a.Value == 0 && a.Unit == UnitConstant
}

func (a Amount) Get(ctx *AmountContext, widthRelative bool) float32 {
	return a.Unit.Get(a.Value, ctx, widthRelative)
}

func (a *Amount) Set(value float32, unit Unit) {
	a.Value = value
	a.Unit = unit
}

func (a Amount) MarshalText() ([]byte, error) {
	preferredSuffix := unitToSuffix[a.Unit]
	if a.Unit.SupportsPercent() {
		return []byte(fmt.Sprintf("%f%%%s", a.Value*100, preferredSuffix)), nil
	} else {
		return []byte(fmt.Sprintf("%f%s", a.Value, preferredSuffix)), nil
	}
}

var AmountFormatRegex = regexp.MustCompile(`^([+-]?\d*(\.\d+)?)((%)?[a-zA-Z]*)$`)

func (a *Amount) UnmarshalText(text []byte) error {
	s := string(text)
	parsed := AmountFormatRegex.FindStringSubmatch(s)
	if len(parsed) != 5 {
		return fmt.Errorf("Amount %s is not in a valid format", s)
	}
	value := parsed[1]
	percent := parsed[4]
	suffix := parsed[3]

	parsedValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}
	if percent == "%" {
		parsedValue *= 0.01
	}
	unit, exists := suffixToUnit[strings.ToLower(suffix)]
	if !exists {
		return fmt.Errorf("Amount %s has unknown unit: %s", s, suffix)
	}
	*a = Amount{Value: float32(parsedValue), Unit: unit}
	return nil
}

type UnitContext struct {
	Width  float32
	Height float32
}

func (c UnitContext) Get(widthRelative bool) float32 {
	if widthRelative {
		return c.Width
	}
	return c.Height
}

type AmountContext struct {
	Parent   UnitContext
	FontSize float32
	View     UnitContext
	Window   UnitContext
	Screen   UnitContext
}

func (c *AmountContext) Resize(width, height float32, fontSize Amount) *AmountContext {
	if c.IsSameSize(width, height, fontSize) {
		return c
	}
	copy := *c
	copy.Parent.Width = width
	copy.Parent.Height = height
	copy.FontSize = fontSize.Get(&copy, true)
	return &copy
}

func (c *AmountContext) IsSameSize(width, height float32, fontSize Amount) bool {
	return c.Parent.Width == width && c.Parent.Height == height && fontSize.Get(c, true) == c.FontSize
}

func (c *AmountContext) ForFont(font Amount) *AmountContext {
	computed := font.Get(c, true)
	if computed == c.FontSize {
		return c
	}
	copy := *c
	copy.FontSize = computed
	return &copy
}

func (c AmountContext) GetScalar(unit Unit, widthRelative bool) float32 {
	switch unit {
	case UnitConstant:
		return 1
	case UnitParent:
		return c.Parent.Get(widthRelative)
	case UnitParentWidth:
		return c.Parent.Width
	case UnitParentHeight:
		return c.Parent.Height
	case UnitView:
		return c.View.Get(widthRelative)
	case UnitViewWidth:
		return c.View.Width
	case UnitViewHeight:
		return c.View.Height
	case UnitWindow:
		return c.Window.Get(widthRelative)
	case UnitWindowWidth:
		return c.Window.Width
	case UnitWindowHeight:
		return c.Window.Height
	case UnitScreen:
		return c.Screen.Get(widthRelative)
	case UnitScreenWidth:
		return c.Screen.Width
	case UnitScreenHeight:
		return c.Screen.Height
	case UnitFont:
		return c.FontSize
	}
	return 1
}

type AmountBounds struct {
	Left   Amount
	Right  Amount
	Top    Amount
	Bottom Amount
}

func NewAmountBounds(left, top, right, bottom float32) AmountBounds {
	return NewAmountBoundsUnit(left, top, right, bottom, UnitConstant)
}

func NewAmountBoundsUniform(value float32, unit Unit) AmountBounds {
	return NewAmountBoundsUnit(value, value, value, value, unit)
}

func NewAmountBoundsUnit(left, top, right, bottom float32, unit Unit) AmountBounds {
	return AmountBounds{
		Left:   Amount{Value: left, Unit: unit},
		Top:    Amount{Value: top, Unit: unit},
		Right:  Amount{Value: right, Unit: unit},
		Bottom: Amount{Value: bottom, Unit: unit},
	}
}

func (a AmountBounds) GetBounds(ctx *AmountContext) Bounds {
	return Bounds{
		Left:   a.Left.Get(ctx, true),
		Top:    a.Top.Get(ctx, false),
		Right:  a.Right.Get(ctx, true),
		Bottom: a.Bottom.Get(ctx, false),
	}
}

func (a *AmountBounds) SetAmount(amount Amount) {
	a.Left = amount
	a.Top = amount
	a.Right = amount
	a.Bottom = amount
}

func (a AmountBounds) WithAmount(amount Amount) AmountBounds {
	a.Left = amount
	a.Top = amount
	a.Right = amount
	a.Bottom = amount
	return a
}

func (a *AmountBounds) Set(value float32, unit Unit) {
	a.Left.Set(value, unit)
	a.Top.Set(value, unit)
	a.Right.Set(value, unit)
	a.Bottom.Set(value, unit)
}

func (a AmountBounds) IsZero() bool {
	return a.Left.IsZero() && a.Top.IsZero() && a.Right.IsZero() && a.Bottom.IsZero()
}

type AmountCorners struct {
	TopLeft     Amount
	TopRight    Amount
	BottomLeft  Amount
	BottomRight Amount
}

func (a *AmountCorners) Set(value float32, unit Unit) {
	a.TopLeft.Set(value, unit)
	a.TopRight.Set(value, unit)
	a.BottomLeft.Set(value, unit)
	a.BottomRight.Set(value, unit)
}

type AmountPoint struct {
	X Amount
	Y Amount
}

func NewAmountPoint(x, y float32) AmountPoint {
	return NewAmountPointUnit(x, y, UnitConstant)
}

func NewAmountPointUnit(x, y float32, unit Unit) AmountPoint {
	return AmountPoint{
		X: Amount{Value: x, Unit: unit},
		Y: Amount{Value: y, Unit: unit},
	}
}

func (a AmountPoint) Get(ctx *AmountContext) (float32, float32) {
	return a.X.Get(ctx, true), a.Y.Get(ctx, false)
}

func (a *AmountPoint) Set(value float32, unit Unit) {
	a.X.Set(value, unit)
	a.Y.Set(value, unit)
}
