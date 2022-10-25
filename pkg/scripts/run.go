package scripts

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
	"text/template"
)

// The Run compiler executes expressions against user provided state.

// An expression compiled into something that can run on a given program state.
type RunExpr struct {
	// Runs the expression against the state and returns the result and if any errors occurred.
	Run func(state *State) (any, error)

	expr Expr
	err  error
}

var _ CompiledExpr = &RunExpr{}

func (e RunExpr) Expr() Expr   { return e.expr }
func (e RunExpr) Error() error { return e.err }

// Allocates a new compiler for RunExpr that doesn't contain any expression compiler entries.
func NewRunCompiler() *Compiler[RunExpr] {
	return NewCompiler(func(e Expr, err error) RunExpr {
		return RunExpr{expr: e, err: err}
	})
}

// Allocates and allows manipulation of a new compiler for RunExpr before returning it.
func CreateRunCompiler(with func(c *Compiler[RunExpr])) *Compiler[RunExpr] {
	c := NewRunCompiler()
	with(c)
	return c
}

// The standard RunExpr compiler for standard expressions.
var RunC *Compiler[RunExpr] = CreateRunCompiler(func(c *Compiler[RunExpr]) {
	c.Defines(RunCompilerEntries)
})

// Compiles and runs the given expression against the given state and returns the result or the error.
func Run(e Expr, state *State) (any, error) {
	runExpr := RunC.Compile(e)
	if runExpr.err != nil {
		return nil, runExpr.err
	}
	result, err := runExpr.Run(state)
	if hasRootCause(err, RETURN) {
		return result, nil
	}
	return result, err
}

// The collection of standard RunExpr compiler entries.
var RunCompilerEntries []CompilerEntry[RunExpr] = []CompilerEntry[RunExpr]{
	// And
	CreateEntry(func(e And, c *Compiler[RunExpr]) RunExpr {
		if e.Conditions == nil || len(e.Conditions) == 0 {
			return c.CreateError(e, errors.New("And must have at least one condition."))
		}
		compiled, err := AndCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				for _, cond := range compiled.Conditions {
					result, err := runBool(cond, state)
					if err != nil {
						return false, ExprError{e, err}
					}
					if !result {
						return false, nil
					}
				}
				return true, nil
			},
		}
	}),
	// Or
	CreateEntry(func(e Or, c *Compiler[RunExpr]) RunExpr {
		if e.Conditions == nil || len(e.Conditions) == 0 {
			return c.CreateError(e, errors.New("Or must have at least one condition."))
		}
		compiled, err := OrCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				for _, cond := range compiled.Conditions {
					result, err := runBool(cond, state)
					if err != nil {
						return false, ExprError{e, err}
					}
					if result {
						return true, nil
					}
				}
				return false, nil
			},
		}
	}),
	// Not
	CreateEntry(func(e Not, c *Compiler[RunExpr]) RunExpr {
		compiled, err := NotCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				result, err := runBool(compiled.Condition, state)
				if err != nil {
					return false, ExprError{e, err}
				}
				return !result, nil
			},
		}
	}),
	// Body
	CreateEntry(func(e Body, c *Compiler[RunExpr]) RunExpr {
		compiled, err := BodyCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				for _, item := range compiled.Lines {
					result, err := item.Run(state)
					if err != nil {
						return result, ExprError{e, err}
					}
				}
				return nil, nil
			},
		}
	}),
	// If
	CreateEntry(func(e If, c *Compiler[RunExpr]) RunExpr {
		if e.Cases == nil || len(e.Cases) == 0 {
			return c.CreateError(e, errors.New("If must have at least one if then statement."))
		}
		compiled, err := IfCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				for _, ifCase := range compiled.Cases {
					result, err := runBool(ifCase.Condition, state)
					if err != nil {
						return nil, nil
					}
					if result {
						thenResult, err := ifCase.Then.Run(state)
						if err != nil {
							return nil, ExprError{e, err}
						}
						return thenResult, nil
					}
				}
				if compiled.Else.Run != nil {
					result, err := compiled.Else.Run(state)
					if err != nil {
						return nil, ExprError{e, err}
					}
					return result, nil
				}
				return nil, nil
			},
		}
	}),
	// Switch
	CreateEntry(func(e Switch, c *Compiler[RunExpr]) RunExpr {
		if e.Cases == nil || len(e.Cases) == 0 {
			return c.CreateError(e, errors.New("Switch must have at least one case statement."))
		}
		compiled, err := SwitchCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				value, err := compiled.Value.Run(state)
				if err != nil {
					return nil, ExprError{e, err}
				}
				for _, switchCase := range compiled.Cases {
					for _, switchExpected := range switchCase.Expected {
						expected, err := switchExpected.Run(state)
						if err != nil {
							return nil, ExprError{e, err}
						}
						if expected == value {
							result, err := switchCase.Then.Run(state)
							if err != nil {
								return nil, ExprError{e, err}
							}
							return result, nil
						}
					}
				}
				if compiled.Default.Run != nil {
					result, err := compiled.Default.Run(state)
					if err != nil {
						return nil, ExprError{e, err}
					}
					return result, nil
				}
				return nil, nil
			},
		}
	}),
	// Loop
	CreateEntry(func(e Loop, c *Compiler[RunExpr]) RunExpr {
		compiled, err := LoopCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				state.EnterScope()
				defer state.ExitScope()
				if compiled.Start.Run != nil {
					_, err := compiled.Start.Run(state)
					if err != nil {
						return nil, ExprError{e, err}
					}
				}
				iterations := 0
				for {
					if compiled.While.Run != nil {
						result, err := runBool(compiled.While, state)
						if err != nil {
							return nil, ExprError{e, err}
						}
						if !result {
							break
						}
					}
					if compiled.Then.Run != nil {
						result, err := compiled.Then.Run(state)
						if hasRootCause(err, BREAK) {
							break
						} else if err != nil {
							return result, ExprError{e, err}
						}
					}
					if compiled.Next.Run != nil {
						_, err := compiled.Next.Run(state)
						if err != nil {
							return nil, ExprError{e, err}
						}
					}
					iterations++
					if state.Runtime.MaxIterations != 0 && iterations >= state.Runtime.MaxIterations {
						return nil, MAX_ITERATIONS
					}
				}
				return nil, nil
			},
		}
	}),
	// Break
	CreateEntry(func(e Break, c *Compiler[RunExpr]) RunExpr {
		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				return nil, BREAK
			},
		}
	}),
	// Return
	CreateEntry(func(e Return, c *Compiler[RunExpr]) RunExpr {
		compiled, err := ReturnCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				if compiled.Value.Run == nil {
					return nil, RETURN
				}
				result, err := compiled.Value.Run(state)
				if err != nil {
					return nil, ExprError{e, err}
				}
				return result, RETURN
			},
		}
	}),
	// Throw
	CreateEntry(func(e Throw, c *Compiler[RunExpr]) RunExpr {
		compiled, err := ThrowCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				if compiled.Error.Run == nil {
					return nil, ERROR
				}
				result, err := runString(compiled.Error, state)
				if err != nil {
					return nil, ExprError{e, err}
				}
				return result, ExprError{e, errors.New(result)}
			},
		}
	}),
	// Try
	CreateEntry(func(e Try, c *Compiler[RunExpr]) RunExpr {
		compiled, err := TryCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				result, err := compiled.Body.Run(state)
				if err != nil {
					if hasRootCause(err, BREAK) || hasRootCause(err, RETURN) {
						return result, err
					}
					if compiled.Catch.Run != nil {
						inner := state.EnterScope()
						inner.Ref("error").Set(err.Error())
						result, err = compiled.Catch.Run(state)
						state.ExitScope()
					}
				}
				if compiled.Finally.Run != nil {
					finallyResult, finallyErr := compiled.Finally.Run(state)
					if hasRootCause(finallyErr, BREAK) || hasRootCause(finallyErr, RETURN) {
						result = finallyResult
					} else if err == nil {
						err = finallyErr
					}
				}
				return result, err
			},
		}
	}),
	// Assert
	CreateEntry(func(e Assert, c *Compiler[RunExpr]) RunExpr {
		compiled, err := AssertCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				result, err := runBool(compiled.Expect, state)
				if err != nil {
					return nil, ExprError{e, err}
				}
				if result {
					return nil, nil
				}
				assertion := ExprError{e, ASSERT_ERROR}
				if compiled.Error.Run != nil {
					message, err := runString(compiled.Error, state)
					if err != nil {
						return nil, ExprError{e, err}
					}
					assertion.inner = errors.New(message)
				}
				return nil, state.Assert(assertion)
			},
		}
	}),
	// Constant
	CreateEntry(func(e Constant, c *Compiler[RunExpr]) RunExpr {
		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				return e.Value, nil
			},
		}
	}),
	// Compare
	CreateEntry(func(e Compare, c *Compiler[RunExpr]) RunExpr {
		compiled, err := CompareCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				leftRaw, err := compiled.Left.Run(state)
				if err != nil {
					return false, ExprError{e, err}
				}
				rightRaw, err := compiled.Right.Run(state)
				if err != nil {
					return false, ExprError{e, err}
				}

				leftBase := concrete(leftRaw)
				rightBase := concrete(rightRaw)

				if !leftBase.IsValid() {
					return false, ExprError{e, fmt.Errorf("Invalid left value in compare")}
				}
				if !rightBase.IsValid() {
					return false, ExprError{e, fmt.Errorf("Invalid right value in compare")}
				}

				leftRaw = leftBase.Interface()
				rightRaw = rightBase.Interface()

				if leftString, ok := leftRaw.(string); ok {
					if rightString, ok := rightRaw.(string); ok {
						if state.Runtime.Insensitive {
							leftString = strings.ToLower(leftString)
							rightString = strings.ToLower(rightString)
						}
						switch e.Type {
						case EQ:
							return leftString == rightString, nil
						case NEQ:
							return leftString != rightString, nil
						case LT:
							return leftString < rightString, nil
						case GT:
							return leftString > rightString, nil
						case LTE:
							return leftString <= rightString, nil
						case GTE:
							return leftString >= rightString, nil
						default:
							return false, INVALID_COMPARISON
						}
					}
				}

				if leftBool, ok := leftRaw.(bool); ok {
					if rightBool, ok := rightRaw.(bool); ok {
						switch e.Type {
						case EQ:
							return leftBool == rightBool, nil
						case NEQ:
							return leftBool != rightBool, nil
						default:
							return false, INVALID_COMPARISON
						}
					}
				}

				left, leftErr := toFloat(leftRaw)
				right, rightErr := toFloat(rightRaw)

				if leftErr == nil && rightErr == nil {
					switch e.Type {
					case EQ:
						return math.Abs(left-right) <= state.Runtime.Epsilon, nil
					case NEQ:
						return math.Abs(left-right) > state.Runtime.Epsilon, nil
					case LT:
						return left < right, nil
					case LTE:
						return left <= right, nil
					case GT:
						return left > right, nil
					case GTE:
						return left >= right, nil
					}
				}

				switch e.Type {
				case EQ:
					return leftRaw == rightRaw, nil
				case NEQ:
					return leftRaw != rightRaw, nil
				}

				return false, INVALID_COMPARISON
			},
		}
	}),
	// Binary
	CreateEntry(func(e Binary, c *Compiler[RunExpr]) RunExpr {
		compiled, err := BinaryCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				leftRaw, err := compiled.Left.Run(state)
				if err != nil {
					return false, ExprError{e, err}
				}
				rightRaw, err := compiled.Right.Run(state)
				if err != nil {
					return false, ExprError{e, err}
				}

				leftBase := concrete(leftRaw)
				rightBase := concrete(rightRaw)

				if !leftBase.IsValid() {
					return false, ExprError{e, fmt.Errorf("Invalid left value in binary operation")}
				}
				if !rightBase.IsValid() {
					return false, ExprError{e, fmt.Errorf("Invalid right value in binary operation")}
				}

				leftRaw = leftBase.Interface()
				rightRaw = rightBase.Interface()

				if leftString, ok := leftRaw.(string); ok {
					if rightString, ok := rightRaw.(string); ok {

						switch e.Type {
						case ADD:
							return leftString + rightString, nil
						default:
							return false, INVALID_BINARY
						}
					}
				}

				if leftBool, ok := leftRaw.(bool); ok {
					if rightBool, ok := rightRaw.(bool); ok {
						switch e.Type {
						case AND:
							return leftBool && rightBool, nil
						case OR:
							return leftBool || rightBool, nil
						case XOR:
							return leftBool != rightBool, nil
						default:
							return false, INVALID_BINARY
						}
					}
				}

				switch e.Type {
				case AND, OR, XOR, LSHIFT, RSHIFT, GCD:
					left, leftErr := toInt(leftRaw)
					if leftErr != nil {
						return false, leftErr
					}
					right, rightErr := toInt(rightRaw)
					if rightErr != nil {
						return false, rightErr
					}
					switch e.Type {
					case AND:
						return left & right, nil
					case OR:
						return left | right, nil
					case XOR:
						return left ^ right, nil
					case LSHIFT:
						return left << right, nil
					case RSHIFT:
						return left >> right, nil
					case GCD:
						for right != 0 {
							t := right
							right = left % right
							left = t
						}
						return left, nil
					}
				case ADD, SUB, MUL, DIV, MOD, POW, MAX, MIN, ATAN2:
					left, leftErr := toFloat(leftRaw)
					if leftErr != nil {
						return false, leftErr
					}
					right, rightErr := toFloat(rightRaw)
					if rightErr != nil {
						return false, rightErr
					}
					switch e.Type {
					case ADD:
						return left + right, nil
					case SUB:
						return left - right, nil
					case MUL:
						return left * right, nil
					case DIV:
						return left / right, nil
					case MOD:
						return math.Mod(left, right), nil
					case POW:
						return math.Pow(left, right), nil
					case MAX:
						return math.Max(left, right), nil
					case MIN:
						return math.Min(left, right), nil
					case ATAN2:
						return math.Atan2(left, right), nil
					}
				}

				return false, INVALID_BINARY
			},
		}
	}),
	// Unary
	CreateEntry(func(e Unary, c *Compiler[RunExpr]) RunExpr {
		compiled, err := UnaryCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				valueRaw, err := compiled.Value.Run(state)
				if err != nil {
					return false, ExprError{e, err}
				}

				valueBase := concrete(valueRaw)

				if !valueBase.IsValid() {
					return false, ExprError{e, fmt.Errorf("Invalid value in unary operation")}
				}

				valueRaw = valueBase.Interface()

				valueFloat, err := toFloat(valueRaw)
				if err != nil {
					return 0, err
				}

				switch e.Type {
				case NEG:
					return -valueFloat, nil
				case NOT:
					return ^int64(valueFloat), nil
				case SQR:
					return valueFloat * valueFloat, nil
				case ABS:
					return math.Abs(valueFloat), nil
				case COS:
					return math.Cos(valueFloat), nil
				case SIN:
					return math.Sin(valueFloat), nil
				case TAN:
					return math.Tan(valueFloat), nil
				case ACOS:
					return math.Acos(valueFloat), nil
				case ASIN:
					return math.Asin(valueFloat), nil
				case ATAN:
					return math.Atan(valueFloat), nil
				case FLOOR:
					return math.Floor(valueFloat), nil
				case CEIL:
					return math.Ceil(valueFloat), nil
				case ROUND:
					return math.Round(valueFloat), nil
				case SQRT:
					return math.Sqrt(valueFloat), nil
				case CBRT:
					return math.Cbrt(valueFloat), nil
				case LN:
					return math.Log(valueFloat), nil
				case LOG2:
					return math.Log2(valueFloat), nil
				case LOG10:
					return math.Log10(valueFloat), nil
				case TRUNC:
					return math.Trunc(valueFloat), nil
				}

				return 0, INVALID_UNARY
			},
		}
	}),
	// Get
	CreateEntry(func(e Get, c *Compiler[RunExpr]) RunExpr {
		compiled, err := GetCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				ref, err := pathRef(compiled.Path, state)
				if err != nil {
					return nil, ExprError{e, err}
				}
				if refError := ref.GetError(); refError != nil {
					return nil, refError
				}
				return ref.Value.Interface(), nil
			},
		}
	}),
	// Set
	CreateEntry(func(e Set, c *Compiler[RunExpr]) RunExpr {
		compiled, err := SetCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				ref, err := pathRef(compiled.Path, state)
				if err != nil {
					return nil, ExprError{e, err}
				}
				if refError := ref.GetError(); refError != nil {
					return nil, refError
				}
				value, err := compiled.Value.Run(state)
				if err != nil {
					return nil, ExprError{e, err}
				}
				err = ref.Set(value)
				if err != nil {
					return nil, ExprError{e, err}
				}
				return nil, nil
			},
		}
	}),
	// Invoke
	CreateEntry(func(e Invoke, c *Compiler[RunExpr]) RunExpr {
		if e.Function == "" {
			return c.CreateError(e, errors.New("Function not specified"))
		}
		compiled, err := InvokeCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				f, fExists := state.Runtime.Funcs[compiled.Function]
				if !fExists {
					return nil, FUNCTION_NOT_FOUND
				}
				params := make([]reflect.Value, 0)
				for _, arg := range f.Arguments {
					paramValue := valueOf(arg.Type)
					if paramExpr, ok := compiled.Params[arg.Name]; ok {
						param, err := paramExpr.Run(state)
						if err != nil {
							return nil, ExprError{e, err}
						}
						paramValue = toType(param, arg.Type)
						if !paramValue.IsValid() {
							return nil, ExprError{e, fmt.Errorf("Invalid argument %s for %s", arg.Name, e.Function)}
						}
					}
					params = append(params, paramValue)
				}
				results := f.Implementation.Call(params)
				result := any(nil)
				err := error(nil)
				if len(results) > 0 {
					first := results[0]
					if firstError, ok := first.Interface().(error); ok {
						err = firstError
					} else {
						result = first.Interface()
						if len(results) > 1 {
							second := results[0]
							if secondError, ok := second.Interface().(error); ok {
								err = secondError
							}
						}
					}
				}
				if err != nil {
					return result, ExprError{e, err}
				}
				return result, nil
			},
		}
	}),
	// Define
	CreateEntry(func(e Define, c *Compiler[RunExpr]) RunExpr {
		if e.Vars == nil || len(e.Vars) == 0 {
			return c.CreateError(e, errors.New("Define requires at least one variable definition."))
		}
		compiled, err := DefineCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				scope := state.EnterScope()
				defer state.ExitScope()
				for _, varPair := range compiled.Vars {
					varRef := scope.Ref(varPair.Name)
					if !varRef.Valid() {
						return nil, varRef.Error
					}
					varValue, err := varPair.Value.Run(state)
					if err != nil {
						return nil, ExprError{e, err}
					}
					err = varRef.Set(varValue)
					if err != nil {
						return nil, ExprError{e, err}
					}
				}
				result, err := compiled.Body.Run(state)
				if err != nil {
					return nil, ExprError{e, err}
				}
				return result, nil
			},
		}
	}),
	// Template
	CreateEntry(func(e Template, c *Compiler[RunExpr]) RunExpr {
		compiled, err := TemplateCompile(c, e)
		if err != nil {
			return c.CreateError(e, err)
		}

		return RunExpr{
			expr: e,
			Run: func(state *State) (any, error) {
				format, err := runString(compiled.Format, state)
				if err != nil {
					return "", ExprError{e, err}
				}
				formatParsed, err := template.New(format).Parse(format)
				if err != nil {
					return "", ExprError{e, err}
				}
				vars := make(map[string]any)
				for varName, varExpr := range compiled.Vars {
					varValue, err := varExpr.Run(state)
					if err != nil {
						return "", ExprError{e, err}
					}
					vars[varName] = concrete(varValue).Interface()
				}
				out := bytes.Buffer{}
				err = formatParsed.Execute(&out, vars)
				if err != nil {
					return "", ExprError{e, err}
				}
				return out.String(), nil
			},
		}
	}),
}

func hasRootCause(err error, expected error) bool {
	if err == nil {
		return false
	}
	if err == expected {
		return true
	}
	if exprErr, ok := err.(ExprError); ok {
		return exprErr.Root() == expected
	}
	return false
}

func nodeResolve(node any, state *State) (any, error) {
	if runExpr, ok := node.(RunExpr); ok {
		return runExpr.Run(state)
	} else {
		return node, nil
	}
}

func pathRef(e []RunExpr, state *State) (*Ref, error) {
	if e == nil || len(e) == 0 {
		return nil, INVALID_REF
	}
	first, err := nodeResolve(e[0], state)
	if err != nil {
		return nil, err
	}
	value := state.Ref(first)
	if !value.Valid() {
		return value, INVALID_REF
	}
	for i := 1; i < len(e); i++ {
		next, err := nodeResolve(e[i], state)
		if err != nil {
			return nil, err
		}
		value = value.Ref(next)
		if !value.Valid() {
			return value, INVALID_REF
		}
	}
	return value, nil
}

func runBool(e RunExpr, state *State) (bool, error) {
	return runExpr(e, state, toBool)
}

func runString(e RunExpr, state *State) (string, error) {
	return runExpr(e, state, toString)
}

func runFloat(e RunExpr, state *State) (float64, error) {
	return runExpr(e, state, toFloat)
}

func runExpr[T any](e RunExpr, state *State, fn func(value any) (T, error)) (T, error) {
	var empty T
	result, err := e.Run(state)
	if err != nil {
		return empty, ExprError{e, err}
	}
	resultTyped, err := fn(result)
	if err != nil {
		return empty, ExprError{e, err}
	}
	return resultTyped, nil
}
