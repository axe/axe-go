package scripts

import (
	"fmt"
	"strings"
)

// The common interface that expressions can be compiled to. It contains the expression
// that was compiled and a compilation error if any were found.
type CompiledExpr interface {
	// The compiled expression.
	Expr() Expr
	// The compilation error, if any.
	Error() error
}

// An entry in the compiler which is an expression - compile function pair.
type CompilerEntry[CE CompiledExpr] struct {
	// The expression the compile function can handle.
	Expr Expr
	// A compile function for the expression type which generates a CompiledExpr.
	Compile Compile[CE]
}

// A function which takes an expression and compiler and returns the CompiledExpr.
type Compile[CE CompiledExpr] func(e Expr, c *Compiler[CE]) CE

// A typed compile function for a better developer experience.
type CompileTyped[E Expr, CE CompiledExpr] func(e E, c *Compiler[CE]) CE

// A function which generates an error form of the CompiledExpr type.
type CompilerError[CE CompiledExpr] func(e Expr, err error) CE

// A compiler takes an expression and compiles it down to a target CompiledExpr.
type Compiler[CE CompiledExpr] struct {
	// The map of entries (keyed by Expr name) supported in this compiler.
	Entries map[string]CompilerEntry[CE]
	// A function which returns the error form of the CompiledExpr type.
	CreateError CompilerError[CE]
}

// Allocates a new compiler with the given error creation function.
func NewCompiler[EC CompiledExpr](createError CompilerError[EC]) *Compiler[EC] {
	return &Compiler[EC]{
		Entries:     make(map[string]CompilerEntry[EC]),
		CreateError: createError,
	}
}

// Allocates and allows manipulation of a new compiler for CompiledExpr before returning.
func CreateCompiler[EC CompiledExpr](createError CompilerError[EC], with func(c *Compiler[EC])) *Compiler[EC] {
	c := NewCompiler(createError)
	with(c)
	return c
}

// Adds an entry (Expr, Compile pair) to the compiler for a single expression.
func (c *Compiler[CE]) Define(entry CompilerEntry[CE]) {
	c.Entries[strings.ToLower(entry.Expr.Name())] = entry
}

// Adds entries (Expr, Compile pairs) to the compiler for multiple expressions.
func (c *Compiler[CE]) Defines(entries []CompilerEntry[CE]) {
	for i := range entries {
		c.Define(entries[i])
	}
}

// Compiles the given expression to the CompiledExpr type. If there was an error compiling
// the returned CompiledExpr will have a non-nil error.
func (c *Compiler[EC]) Compile(e Expr) EC {
	entry, exists := c.Entries[strings.ToLower(e.Name())]
	if !exists {
		return c.CreateError(e, fmt.Errorf("No compiler found for %s", e.Name()))
	}
	return entry.Compile(e, c)
}

// Compiles the given expression if it's non-nil. If it is nil then an empty CompiledExpr is
// returned. Users of optional compilation need to check the state of the CompiledExpr before
// attempting to use it.
func (c *Compiler[EC]) CompileMaybe(e Expr) EC {
	if e == nil {
		var empty EC
		return empty
	}
	return c.Compile(e)
}

// Compiles a slice of expressions into a slice of CompiledExpr. If any of the items generated an
// error then compilation stops and the error and the generated slice up until that point is
// returned. This will always return a non-nil slice even if the given slice is nil.
func (c *Compiler[CE]) CompileList(list []Expr) ([]CE, error) {
	compiled := make([]CE, len(list))
	if list != nil {
		for index := range list {
			compiledExpr := c.Compile(list[index])
			if compiledExpr.Error() != nil {
				return compiled, compiledExpr.Error()
			}
			compiled[index] = compiledExpr
		}
	}
	return compiled, nil
}

// Compiles a slice of values which can be raw values or expressions into a slice of CompiledExpr.
// If any of the items generated an error then compilation stops and the error and the generated
// slice up until that point is returned. This will always return a non-nil slice even if the
// given slice is nil.
func (c *Compiler[CE]) CompilePath(path []any) ([]CE, error) {
	compiled := make([]CE, len(path))
	if path != nil {
		for index := range path {
			var itemExpr Expr = Constant{path[index]}
			if expr, ok := path[index].(Expr); ok {
				itemExpr = expr
			}
			compiledExpr := c.Compile(itemExpr)
			if compiledExpr.Error() != nil {
				return compiled, compiledExpr.Error()
			}
			compiled[index] = compiledExpr
		}
	}
	return compiled, nil
}

// Compiles a map of expressions into a map of CompiledExpr. If any of the items generated an
// error then compilation stops and the error and the generated map up until that point is
// returned. This will always return a non-nil map even if the given map is nil.
func (c *Compiler[CE]) CompileMap(m map[string]Expr) (map[string]CE, error) {
	compiled := make(map[string]CE)
	if m != nil {
		for key := range m {
			compiledExpr := c.Compile(m[key])
			if compiledExpr.Error() != nil {
				return compiled, compiledExpr.Error()
			}
			compiled[key] = compiledExpr
		}
	}
	return compiled, nil
}

// Creates a compiler entry for a given expression type and compilation function.
func CreateEntry[E Expr, CE CompiledExpr](compile CompileTyped[E, CE]) CompilerEntry[CE] {
	var inst E
	return CompilerEntry[CE]{
		Expr: inst,
		Compile: func(e Expr, c *Compiler[CE]) CE {
			if typed, ok := e.(E); ok {
				return compile(typed, c)
			}
			panic("Corrupt typed compilation.")
		},
	}
}
