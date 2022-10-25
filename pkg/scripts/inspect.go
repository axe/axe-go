package scripts

import (
	"reflect"
)

// The Inspect compiler provides type information and expression structure for
// consumers of expressions to understand.

// An expression compiled into something that can provide execution metadata.
type InspectExpr struct {
	// Tries to determine the possible return types for this expression and the given state.
	ReturnType func(state *State) Types

	expr Expr
	err  error
}

var _ CompiledExpr = &InspectExpr{}

func (e InspectExpr) Expr() Expr   { return e.expr }
func (e InspectExpr) Error() error { return e.err }

// Allocates a new compiler for InspectExpr that doesn't contain any expression compiler entries.
func NewInspectCompiler() *Compiler[InspectExpr] {
	return NewCompiler(func(e Expr, err error) InspectExpr {
		return InspectExpr{expr: e, err: err}
	})
}

// Allocates and allows manipulation of a new compiler for InspectExpr before returning it.
func CreateInspectCompiler(with func(c *Compiler[InspectExpr])) *Compiler[InspectExpr] {
	c := NewInspectCompiler()
	with(c)
	return c
}

// The standard InspectExpr compiler for standard expressions.
var InspectC *Compiler[InspectExpr] = CreateInspectCompiler(func(c *Compiler[InspectExpr]) {
	c.Defines(InspectCompilerEntries)
})

// Compiles the given expression using the standard compiler.
func Inspect(e Expr) InspectExpr {
	return InspectC.Compile(e)
}

// A set of types.
type Types map[reflect.Type]struct{}

// Adds the type to the set if it doesn't exist already.
func (t Types) Add(newType reflect.Type) {
	t[newType] = struct{}{}
}

// Adds the given set of types to this set of types if they don't exist already.
func (t Types) AddTypes(types Types) {
	if types != nil {
		for newType := range types {
			t[newType] = struct{}{}
		}
	}
}

// Returns a slice of the types found in this set.
func (t Types) ToSlice() []reflect.Type {
	slice := make([]reflect.Type, 0)
	for key := range t {
		slice = append(slice, key)
	}
	return slice
}

// Common constant return types.
var typesBool = Types{reflect.TypeOf(true): struct{}{}}
var typesNumber = Types{reflect.TypeOf(float64(0)): struct{}{}, reflect.TypeOf(int64(0)): struct{}{}}
var typesVoid = Types{reflect.TypeOf(any(nil)): struct{}{}}
var typesString = Types{reflect.TypeOf(""): struct{}{}}
var typesAny = Types{anyType: struct{}{}}

// The collection of standard InspectExpr compiler entries.
var InspectCompilerEntries []CompilerEntry[InspectExpr] = []CompilerEntry[InspectExpr]{
	// And
	CreateEntry(func(e And, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesBool
			},
		}
	}),
	// Or
	CreateEntry(func(e Or, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesBool
			},
		}
	}),
	// Not
	CreateEntry(func(e Not, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesBool
			},
		}
	}),
	// Body
	CreateEntry(func(e Body, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesVoid
			},
		}
	}),
	// If
	CreateEntry(func(e If, c *Compiler[InspectExpr]) InspectExpr {
		compiled, err := IfCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				ret := Types{}
				for _, cond := range compiled.Cases {
					ret.AddTypes(cond.Then.ReturnType(state))
				}
				if compiled.Else.ReturnType == nil {
					ret.AddTypes(typesVoid)
				} else {
					ret.AddTypes(compiled.Else.ReturnType(state))
				}
				return ret
			},
		}
	}),
	// Switch
	CreateEntry(func(e Switch, c *Compiler[InspectExpr]) InspectExpr {
		compiled, err := SwitchCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				ret := Types{}
				for _, cond := range compiled.Cases {
					ret.AddTypes(cond.Then.ReturnType(state))
				}
				if compiled.Default.ReturnType == nil {
					ret.AddTypes(typesVoid)
				} else {
					ret.AddTypes(compiled.Default.ReturnType(state))
				}
				return ret
			},
		}
	}),
	// For
	CreateEntry(func(e Loop, c *Compiler[InspectExpr]) InspectExpr {
		compiled, err := LoopCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				if compiled.Then.ReturnType == nil {
					return typesVoid
				}
				ret := Types{}
				ret.AddTypes(compiled.Then.ReturnType(state))
				ret.AddTypes(typesVoid)
				return ret
			},
		}
	}),
	// Break
	CreateEntry(func(e Break, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return nil
			},
		}
	}),
	// Return
	CreateEntry(func(e Return, c *Compiler[InspectExpr]) InspectExpr {
		compiled, err := ReturnCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				if compiled.Value.ReturnType == nil {
					return typesVoid
				}
				return compiled.Value.ReturnType(state)
			},
		}
	}),
	// Throw
	CreateEntry(func(e Throw, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return nil
			},
		}
	}),
	// Try
	CreateEntry(func(e Try, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return nil
			},
		}
	}),
	// Assert
	CreateEntry(func(e Assert, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return nil
			},
		}
	}),
	// Constant
	CreateEntry(func(e Constant, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return Types{reflect.TypeOf(e.Value): struct{}{}}
			},
		}
	}),
	// Compare
	CreateEntry(func(e Compare, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesBool
			},
		}
	}),
	// Binary
	CreateEntry(func(e Binary, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesNumber
			},
		}
	}),
	// Unary
	CreateEntry(func(e Unary, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesNumber
			},
		}
	}),
	// Get TODO
	CreateEntry(func(e Get, c *Compiler[InspectExpr]) InspectExpr {
		// path, err := c.CompileList(e.Path)
		// if err != nil {
		// 	return c.CreateError(e, err)
		// }

		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				// ref, _ := pathRef(path, state)
				// return Types{ref.Type: struct{}{}}
				return typesAny
			},
		}
	}),
	// Set
	CreateEntry(func(e Set, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesVoid
			},
		}
	}),
	// Invoke
	CreateEntry(func(e Invoke, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				f, fExists := state.Runtime.Funcs[e.Function]
				if !fExists {
					return nil
				}
				return Types{f.ReturnType: struct{}{}}
			},
		}
	}),
	// Define
	CreateEntry(func(e Define, c *Compiler[InspectExpr]) InspectExpr {
		compiled, err := DefineCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return compiled.Body.ReturnType(state)
			},
		}
	}),
	// Template
	CreateEntry(func(e Template, c *Compiler[InspectExpr]) InspectExpr {
		return InspectExpr{
			expr: e,
			ReturnType: func(state *State) Types {
				return typesString
			},
		}
	}),
}
