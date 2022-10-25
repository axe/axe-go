package scripts

import (
	"fmt"
	"reflect"
)

// The basic structure of all supported standard expressions:
// - And, Or, Not, Body, If, Loop, Break, Return, Throw, Try, Assert, Constant, Compare,
//   Binary, Get, Set, Invoke, Define, Template

// An Expr represents a node in an abstract syntax tree. The have names so a compile, marshal,
// or unmarshal can know the node type and from there determine how to proceed with processing.
type Expr interface {
	Name() string
}

// And is an expression that resolves to a bool value, true if all inner conditions are true and
// false if any of the inner conditions are false.
type And struct {
	Conditions []Expr
}

type AndCompiled[CE CompiledExpr] struct {
	Conditions []CE
}

var _ Expr = &And{}

func (e And) Name() string { return "And" }

// Or is an expression that resolves to a bool value, true if any of the inner conditions are true
// and false if all of the inner conditions are false.
type Or struct {
	Conditions []Expr
}

type OrCompiled[CE CompiledExpr] struct {
	Conditions []CE
}

var _ Expr = &Or{}

func (e Or) Name() string { return "Or" }

// Not is an expression that resolves to a bool value, it negates the bool value of the inner
// condition.
type Not struct {
	Condition Expr
}

type NotCompiled[CE CompiledExpr] struct {
	Condition CE
}

var _ Expr = &Not{}

func (e Not) Name() string { return "Not" }

// Body is an expression that can have multiple sequential expressions. A Body is the only valid
// place for Return, Throw, Assert, and Break expressions.
type Body struct {
	Lines []Expr
}

type BodyCompiled[CE CompiledExpr] struct {
	Lines []CE
}

var _ Expr = &Body{}

func (e Body) Name() string  { return "Body" }
func (e Body) IsEmpty() bool { return e.Lines == nil || len(e.Lines) == 0 }

// If is an expression with one or many condition & body pairs and optionally a fall through
// expression. The first case which has a true condition gets its Then processed, if no cases
// are true and an Else expression is given that is processed.
type If struct {
	Cases []IfCase
	Else  Expr
}

type IfCase struct {
	Condition Expr
	Then      Expr
}

type IfCompiled[CE CompiledExpr] struct {
	Cases []IfCaseCompiled[CE]
	Else  CE
}

type IfCaseCompiled[CE CompiledExpr] struct {
	Condition CE
	Then      CE
}

var _ Expr = &If{}

func (e If) Name() string { return "If" }

// Switch is an expression that evaluates a value and compares it against expected values and
// if its equivalent then the Then expression is processed. If value never matches an expected
// value then the Default expression, if given, is processed.
type Switch struct {
	Value   Expr
	Cases   []SwitchCase
	Default Expr
}

type SwitchCase struct {
	Expected []Expr
	Then     Expr
}

type SwitchCompiled[CE CompiledExpr] struct {
	Value   CE
	Cases   []SwitchCaseCompiled[CE]
	Default CE
}

type SwitchCaseCompiled[CE CompiledExpr] struct {
	Expected []CE
	Then     CE
}

var _ Expr = &Switch{}

func (e Switch) Name() string { return "Switch" }

// Loop is an expression which performs a loop. All inner expressions are optional. A loop creates
// a new variable scope. If a Set expression is specified that is processed first at the beginning
// of the loop. If a While expression is specified that is processed and if it returns a false
// value then looping ceases. If a Then expression is specified, that is processed if While returned
// true (or is not present). The Then expression catches Breaks called from within and that stops
// looping immediately as well. Finally if a Next expression is specified it's processed after Then
// would be. This format allows for various types of looping styles.
type Loop struct {
	Start Expr
	While Expr
	Next  Expr
	Then  Expr
}

type LoopCompiled[CE CompiledExpr] struct {
	Start CE
	While CE
	Next  CE
	Then  CE
}

var _ Expr = &Loop{}

func (e Loop) Name() string { return "Loop" }

// Break is an expression which is only valid within a Body inside of a Loop. When detected
// the inner loop expressions after it are skipped and processing is resumed after the loop.
type Break struct{}

var _ Expr = &Break{}

func (e Break) Name() string { return "Break" }

// Return is an expression which is only valid within a Body. It will cease all processing of
// expressions in the tree or function and optionally can return a value.
type Return struct {
	Value Expr
}

type ReturnCompiled[CE CompiledExpr] struct {
	Value CE
}

var _ Expr = &Return{}

func (e Return) Name() string { return "Return" }

// Throw is an expression which is only valid within a Body. It will cease all processing of
// expressions in the entire tree unless its caught.
type Throw struct {
	Error Expr
}

type ThrowCompiled[CE CompiledExpr] struct {
	Error CE
}

var _ Expr = &Throw{}

func (e Throw) Name() string { return "Throw" }

// Try is an expression which is only valid within a Body. It will process the body and if
// that generates an error then Catch will be processed if it is given with the error string placed
// in the error value in a new context. If Finally is given its always processed last, even if Catch throws
// an error.
type Try struct {
	Body    Expr
	Catch   Expr
	Finally Expr
}

type TryCompiled[CE CompiledExpr] struct {
	Body    CE
	Catch   CE
	Finally CE
}

var _ Expr = &Try{}

func (e Try) Name() string { return "Try" }

// Assert is an expression which is only valid within a Body. Based on configuration it will
// either behave like an error and cease all processing and return the assertion error, or it
// may just log the assertion.
type Assert struct {
	Expect Expr
	Error  Expr
}

type AssertCompiled[CE CompiledExpr] struct {
	Expect CE
	Error  CE
}

var _ Expr = &Assert{}

func (e Assert) Name() string { return "Assert" }

// Constant is an expression which returns a constant value.
type Constant struct {
	Value any
}

var _ Expr = &Constant{}

func (e Constant) Name() string { return "Constant" }

// Compare is an expression which compares two expressions with the given operation and returns
// a bool value with the result. Implementations dictate which types are supported and if the
// user can override the default comparison logic.
type Compare struct {
	Left  Expr
	Type  CompareOp
	Right Expr
}

type CompareCompiled[CE CompiledExpr] struct {
	Left  CE
	Type  CompareOp
	Right CE
}

type CompareOp string

const (
	EQ  CompareOp = "="
	NEQ CompareOp = "!="
	LT  CompareOp = "<"
	GT  CompareOp = ">"
	LTE CompareOp = "<="
	GTE CompareOp = ">="
)

var _ Expr = &Compare{}

func (e Compare) Name() string { return "Compare" }

// Binary is an expression which performs a math operation on two expressions and returns a result.
// Implementations dictate which types are supported and if the user can override the default
// binary operation logic. If the type of the left and right expressions are different, the resulting
// type will be one of the types that makes the most sense for the given operation.
type Binary struct {
	Left  Expr
	Type  BinaryOp
	Right Expr
}

type BinaryCompiled[CE CompiledExpr] struct {
	Left  CE
	Type  BinaryOp
	Right CE
}

type BinaryOp string

const (
	ADD    BinaryOp = "+"
	SUB    BinaryOp = "-"
	MUL    BinaryOp = "*"
	DIV    BinaryOp = "/"
	MOD    BinaryOp = "%"
	POW    BinaryOp = "^^"
	OR     BinaryOp = "|"
	AND    BinaryOp = "&"
	XOR    BinaryOp = "^"
	LSHIFT BinaryOp = "<<"
	RSHIFT BinaryOp = ">>"
	MAX    BinaryOp = "max"
	MIN    BinaryOp = "min"
	GCD    BinaryOp = "gcd"
	ATAN2  BinaryOp = "atan2"
)

var _ Expr = &Binary{}

func (e Binary) Name() string { return "Binary" }

// Unary is an expression which performs a math operation on one expression and returns a result.
// Implementations dictate which types are supported and if the user can override the default
// binary operation logic. If the type of the left and right expressions are different, the resulting
// type will be one of the types that makes the most sense for the given operation.
type Unary struct {
	Type  UnaryOp
	Value Expr
}

type UnaryCompiled[CE CompiledExpr] struct {
	Type  UnaryOp
	Value CE
}

type UnaryOp string

const (
	NEG   UnaryOp = "-"
	NOT   UnaryOp = "~"
	ABS   UnaryOp = "||"
	COS   UnaryOp = "cos"
	SIN   UnaryOp = "sin"
	TAN   UnaryOp = "tan"
	ACOS  UnaryOp = "acos"
	ASIN  UnaryOp = "asin"
	ATAN  UnaryOp = "atan"
	FLOOR UnaryOp = "floor"
	CEIL  UnaryOp = "ceil"
	ROUND UnaryOp = "round"
	SQRT  UnaryOp = "sqrt"
	CBRT  UnaryOp = "cbrt"
	SQR   UnaryOp = "sqr"
	LN    UnaryOp = "ln"
	LOG2  UnaryOp = "log2"
	LOG10 UnaryOp = "log10"
	TRUNC UnaryOp = "trunc"
)

var _ Expr = &Unary{}

func (e Unary) Name() string { return "Unary" }

// Get is an expression which retrieves a value from the current state. It will start
// at the current scope and make its way to the base data of the state and finally
// to global state until it finds a variable with the same name as the first element
// in the path. The path can be a sequence of int or strings constants or even Exprs
// to be evaluated. Get will walk along the value based on the path to get the final value.
type Get struct {
	Path []any
}

type GetCompiled[CE CompiledExpr] struct {
	Path []CE
}

var _ Expr = &Get{}

func (e Get) Name() string { return "Get" }

// Set is an expression which updates the current state. It will start at the current scope and
// make its way to the base data of the state and finally to global state until it finds a
// variable with the same name as the first element in the path. The path can be a sequence of int
// or strings constants or even Exprs to be evaluated. Set will walk along the value based on the
// path to set the final value. If the path is not compatible with the structure of the current
// state and global data then an error will be returned. Set should perform all necessary
// initialization of values before applying the new value. This may create maps, slices, pointer
// values, or change the size of a slice if a part of the path is an index outside of the slice's
// current length.
type Set struct {
	Path  []any
	Value Expr
}

type SetCompiled[CE CompiledExpr] struct {
	Path  []CE
	Value CE
}

var _ Expr = &Set{}

func (e Set) Name() string { return "Set" }

// Invoke is an expression which finds a defined function with the given name and invokes it
// passing along the parameters specified. If a parameter to the function is not specified the
// default value for that type will be supplied.
type Invoke struct {
	Function string
	Params   map[string]Expr
}

type InvokeCompiled[CE CompiledExpr] struct {
	Function string
	Params   map[string]CE
}

var _ Expr = &Invoke{}

func (e Invoke) Name() string { return "Invoke" }

// Define is an expression which creates a new variable scope and defines new variables that can
// be accessed and updated within the Body of the define.
type Define struct {
	Vars []DefineVar
	Body Expr
}

type DefineVar struct {
	Name  string
	Value Expr
}

type DefineVarCompiled[CE CompiledExpr] struct {
	Name  string
	Value CE
}

type DefineCompiled[CE CompiledExpr] struct {
	Vars []DefineVarCompiled[CE]
	Body CE
}

var _ Expr = &Define{}

func (e Define) Name() string { return "Define" }

// Template is an expression which returns a string with variables injected into it.
type Template struct {
	Format Expr
	Vars   map[string]Expr
}

type TemplateCompiled[CE CompiledExpr] struct {
	Vars   map[string]CE
	Format CE
}

var _ Expr = &Template{}

func (e Template) Name() string { return "Template" }

// An error that retains the expression it happened on and in inner exception.
// This may end up being a linked list of errors which should identify where in the
// tree the error occurred.
type ExprError struct {
	expr  any
	inner error
}

var _ error = &ExprError{}

func (ee ExprError) Error() string {
	return fmt.Sprintf("%v: %v", reflect.TypeOf(ee.expr), ee.inner.Error())
}
func (ee ExprError) Root() error {
	if innerEE, ok := ee.inner.(ExprError); ok {
		return innerEE.Root()
	}
	return ee.inner
}

// Returns the stack trace for the given error, in order from deepest
func Stacktrace(err error) []ExprError {
	stack := make([]ExprError, 0)
	for err != nil {
		if exprErr, ok := err.(ExprError); ok {
			stack = append(stack, exprErr)
			err = exprErr.inner
		} else {
			break
		}
	}
	return stack
}

// Compilation helpers

func AndCompile[CE CompiledExpr](c *Compiler[CE], e And) (AndCompiled[CE], error) {
	compiled := AndCompiled[CE]{}
	conditions, err := c.CompileList(e.Conditions)
	compiled.Conditions = conditions
	if err != nil {
		return compiled, err
	}
	return compiled, nil
}

func OrCompile[CE CompiledExpr](c *Compiler[CE], e Or) (OrCompiled[CE], error) {
	compiled := OrCompiled[CE]{}
	conditions, err := c.CompileList(e.Conditions)
	compiled.Conditions = conditions
	if err != nil {
		return compiled, err
	}
	return compiled, nil
}

func BodyCompile[CE CompiledExpr](c *Compiler[CE], e Body) (BodyCompiled[CE], error) {
	compiled := BodyCompiled[CE]{}
	lines, err := c.CompileList(e.Lines)
	compiled.Lines = lines
	if err != nil {
		return compiled, err
	}
	return compiled, nil
}

func NotCompile[CE CompiledExpr](c *Compiler[CE], e Not) (NotCompiled[CE], error) {
	compiled := NotCompiled[CE]{}
	compiled.Condition = c.Compile(e.Condition)
	if compiled.Condition.Error() != nil {
		return compiled, compiled.Condition.Error()
	}
	return compiled, nil
}

func IfCompile[CE CompiledExpr](c *Compiler[CE], e If) (IfCompiled[CE], error) {
	compiled := IfCompiled[CE]{}

	compiled.Cases = sliceMap(e.Cases, func(source IfCase) IfCaseCompiled[CE] {
		return IfCaseCompiled[CE]{
			Condition: c.Compile(source.Condition),
			Then:      c.Compile(source.Then),
		}
	})

	for _, ifCase := range compiled.Cases {
		if ifCase.Condition.Error() != nil {
			return compiled, ifCase.Condition.Error()
		}
		if ifCase.Then.Error() != nil {
			return compiled, ifCase.Then.Error()
		}
	}

	compiled.Else = c.CompileMaybe(e.Else)
	if compiled.Else.Error() != nil {
		return compiled, compiled.Else.Error()
	}

	return compiled, nil
}

func SwitchCompile[CE CompiledExpr](c *Compiler[CE], e Switch) (SwitchCompiled[CE], error) {
	compiled := SwitchCompiled[CE]{
		Value:   c.Compile(e.Value),
		Default: c.CompileMaybe(e.Default),
	}

	if compiled.Value.Error() != nil {
		return compiled, compiled.Value.Error()
	}
	if compiled.Default.Error() != nil {
		return compiled, compiled.Default.Error()
	}

	compiled.Cases = sliceMap(e.Cases, func(source SwitchCase) SwitchCaseCompiled[CE] {
		expected, _ := c.CompileList(source.Expected)
		return SwitchCaseCompiled[CE]{
			Expected: expected,
			Then:     c.Compile(source.Then),
		}
	})

	for _, switchCase := range compiled.Cases {
		for _, switchExpected := range switchCase.Expected {
			if switchExpected.Error() != nil {
				return compiled, switchExpected.Error()
			}
		}
		if switchCase.Then.Error() != nil {
			return compiled, switchCase.Then.Error()
		}
	}

	return compiled, nil
}

func LoopCompile[CE CompiledExpr](c *Compiler[CE], e Loop) (LoopCompiled[CE], error) {
	compiled := LoopCompiled[CE]{
		Start: c.CompileMaybe(e.Start),
		While: c.CompileMaybe(e.While),
		Next:  c.CompileMaybe(e.Next),
		Then:  c.CompileMaybe(e.Then),
	}
	if compiled.Start.Error() != nil {
		return compiled, compiled.Start.Error()
	}
	if compiled.While.Error() != nil {
		return compiled, compiled.While.Error()
	}
	if compiled.Next.Error() != nil {
		return compiled, compiled.Next.Error()
	}
	if compiled.Then.Error() != nil {
		return compiled, compiled.Then.Error()
	}

	return compiled, nil
}

func SetCompile[CE CompiledExpr](c *Compiler[CE], e Set) (SetCompiled[CE], error) {
	compiled := SetCompiled[CE]{}
	path, err := c.CompilePath(e.Path)
	compiled.Path = path
	if err != nil {
		return compiled, err
	}
	compiled.Value = c.Compile(e.Value)
	if compiled.Value.Error() != nil {
		return compiled, compiled.Value.Error()
	}

	return compiled, nil
}

func DefineCompile[CE CompiledExpr](c *Compiler[CE], e Define) (DefineCompiled[CE], error) {
	compiled := DefineCompiled[CE]{
		Vars: sliceMap(e.Vars, func(source DefineVar) DefineVarCompiled[CE] {
			return DefineVarCompiled[CE]{
				Name:  source.Name,
				Value: c.Compile(source.Value),
			}
		}),
		Body: c.Compile(e.Body),
	}
	for _, defineVar := range compiled.Vars {
		if defineVar.Value.Error() != nil {
			return compiled, defineVar.Value.Error()
		}
	}

	if compiled.Body.Error() != nil {
		return compiled, compiled.Body.Error()
	}

	return compiled, nil
}

func ReturnCompile[CE CompiledExpr](c *Compiler[CE], e Return) (ReturnCompiled[CE], error) {
	compiled := ReturnCompiled[CE]{}
	compiled.Value = c.CompileMaybe(e.Value)
	if compiled.Value.Error() != nil {
		return compiled, compiled.Value.Error()
	}
	return compiled, nil
}

func ThrowCompile[CE CompiledExpr](c *Compiler[CE], e Throw) (ThrowCompiled[CE], error) {
	compiled := ThrowCompiled[CE]{}
	compiled.Error = c.CompileMaybe(e.Error)
	if compiled.Error.Error() != nil {
		return compiled, compiled.Error.Error()
	}
	return compiled, nil
}

func TryCompile[CE CompiledExpr](c *Compiler[CE], e Try) (TryCompiled[CE], error) {
	compiled := TryCompiled[CE]{
		Body:    c.Compile(e.Body),
		Catch:   c.CompileMaybe(e.Catch),
		Finally: c.CompileMaybe(e.Finally),
	}
	if compiled.Body.Error() != nil {
		return compiled, compiled.Body.Error()
	}
	if compiled.Catch.Error() != nil {
		return compiled, compiled.Catch.Error()
	}
	if compiled.Finally.Error() != nil {
		return compiled, compiled.Finally.Error()
	}
	return compiled, nil
}

func AssertCompile[CE CompiledExpr](c *Compiler[CE], e Assert) (AssertCompiled[CE], error) {
	compiled := AssertCompiled[CE]{}
	compiled.Expect = c.Compile(e.Expect)
	if compiled.Expect.Error() != nil {
		return compiled, compiled.Expect.Error()
	}
	compiled.Error = c.CompileMaybe(e.Error)
	if compiled.Error.Error() != nil {
		return compiled, compiled.Error.Error()
	}
	return compiled, nil
}

func CompareCompile[CE CompiledExpr](c *Compiler[CE], e Compare) (CompareCompiled[CE], error) {
	compiled := CompareCompiled[CE]{
		Type:  e.Type,
		Left:  c.Compile(e.Left),
		Right: c.Compile(e.Right),
	}
	if compiled.Left.Error() != nil {
		return compiled, compiled.Left.Error()
	}
	if compiled.Right.Error() != nil {
		return compiled, compiled.Right.Error()
	}
	return compiled, nil
}

func BinaryCompile[CE CompiledExpr](c *Compiler[CE], e Binary) (BinaryCompiled[CE], error) {
	compiled := BinaryCompiled[CE]{
		Type:  e.Type,
		Left:  c.Compile(e.Left),
		Right: c.Compile(e.Right),
	}
	if compiled.Left.Error() != nil {
		return compiled, compiled.Left.Error()
	}
	if compiled.Right.Error() != nil {
		return compiled, compiled.Right.Error()
	}
	return compiled, nil
}

func UnaryCompile[CE CompiledExpr](c *Compiler[CE], e Unary) (UnaryCompiled[CE], error) {
	compiled := UnaryCompiled[CE]{
		Type:  e.Type,
		Value: c.Compile(e.Value),
	}
	if compiled.Value.Error() != nil {
		return compiled, compiled.Value.Error()
	}
	return compiled, nil
}

func GetCompile[CE CompiledExpr](c *Compiler[CE], e Get) (GetCompiled[CE], error) {
	compiled := GetCompiled[CE]{}
	path, err := c.CompilePath(e.Path)
	compiled.Path = path
	if err != nil {
		return compiled, err
	}
	return compiled, nil
}

func InvokeCompile[CE CompiledExpr](c *Compiler[CE], e Invoke) (InvokeCompiled[CE], error) {
	compiled := InvokeCompiled[CE]{
		Function: e.Function,
	}
	params, err := c.CompileMap(e.Params)
	compiled.Params = params
	if err != nil {
		return compiled, err
	}
	return compiled, nil
}

func TemplateCompile[CE CompiledExpr](c *Compiler[CE], e Template) (TemplateCompiled[CE], error) {
	vars, err := c.CompileMap(e.Vars)
	if err != nil {
		return TemplateCompiled[CE]{}, err
	}
	compiled := TemplateCompiled[CE]{
		Vars:   vars,
		Format: c.Compile(e.Format),
	}
	if compiled.Format.Error() != nil {
		return compiled, compiled.Format.Error()
	}

	return compiled, nil
}
