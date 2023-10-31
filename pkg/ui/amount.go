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

func (a Unit) Get(value float32, ctx AmountContext) float32 {
	return ctx.GetScalar(a) * value
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

func (a Amount) Get(ctx AmountContext) float32 {
	return a.Unit.Get(a.Value, ctx)
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
	Value  float32
	Width  float32
	Height float32
}

type AmountContext struct {
	Parent   UnitContext
	FontSize float32
	View     UnitContext
	Window   UnitContext
	Screen   UnitContext
}

func (c AmountContext) WithParent(width, height float32) AmountContext {
	c.Parent.Value = width
	c.Parent.Width = width
	c.Parent.Height = height
	return c
}

func (c AmountContext) ForWidth() AmountContext {
	c.Parent.Value = c.Parent.Width
	c.View.Value = c.View.Width
	c.Window.Value = c.Window.Width
	c.Screen.Value = c.Screen.Width
	return c
}

func (c AmountContext) ForHeight() AmountContext {
	c.Parent.Value = c.Parent.Height
	c.View.Value = c.View.Height
	c.Window.Value = c.Window.Height
	c.Screen.Value = c.Screen.Height
	return c
}

func (c AmountContext) GetScalar(unit Unit) float32 {
	switch unit {
	case UnitConstant:
		return 1
	case UnitParent:
		return c.Parent.Value
	case UnitParentWidth:
		return c.Parent.Width
	case UnitParentHeight:
		return c.Parent.Height
	case UnitView:
		return c.View.Value
	case UnitViewWidth:
		return c.View.Width
	case UnitViewHeight:
		return c.View.Height
	case UnitWindow:
		return c.Window.Value
	case UnitWindowWidth:
		return c.Window.Width
	case UnitWindowHeight:
		return c.Window.Height
	case UnitScreen:
		return c.Screen.Value
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

func (a AmountBounds) GetBounds(ctx AmountContext) Bounds {
	widthCtx := ctx.ForWidth()
	heightCtx := ctx.ForHeight()

	return Bounds{
		Left:   a.Left.Get(widthCtx),
		Top:    a.Top.Get(heightCtx),
		Right:  a.Right.Get(widthCtx),
		Bottom: a.Bottom.Get(heightCtx),
	}
}

func (a *AmountBounds) SetAmount(amount Amount) {
	a.Left = amount
	a.Top = amount
	a.Right = amount
	a.Bottom = amount
}

func (a *AmountBounds) Set(value float32, unit Unit) {
	a.Left.Set(value, unit)
	a.Top.Set(value, unit)
	a.Right.Set(value, unit)
	a.Bottom.Set(value, unit)
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

func (a AmountPoint) Get(ctx AmountContext) (float32, float32) {
	widthCtx := ctx.ForWidth()
	heightCtx := ctx.ForHeight()

	return a.X.Get(widthCtx), a.Y.Get(heightCtx)
}

func (a *AmountPoint) Set(value float32, unit Unit) {
	a.X.Set(value, unit)
	a.Y.Set(value, unit)
}
