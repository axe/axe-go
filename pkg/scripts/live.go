package scripts

import (
	"errors"
	"reflect"
)

var BREAK = errors.New("Break")
var RETURN = errors.New("Return")
var ERROR = errors.New("Error")
var ASSERT_ERROR = errors.New("Assertion error")
var INVALID_BINARY_TYPE = errors.New("Invalid type for binary operation")
var INVALID_COMPARE_TYPE = errors.New("Types incomparable")
var INVALID_BOOL_TYPE = errors.New("Value could not be converted to bool")
var INVALID_COMPARISON = errors.New("Invalid comparison for types")
var INVALID_BINARY = errors.New("Invalid binary operation for types")
var INVALID_UNARY = errors.New("Invalid unary operation for type")
var INVALID_REF = errors.New("Undefined reference")
var FUNCTION_NOT_FOUND = errors.New("Function not found")
var INDEX_OUTSIDE_ARRAY = errors.New("Index outside of array")
var MAX_ITERATIONS = errors.New("Max iterations met")

type FunctionArg struct {
	Name string
	Type reflect.Type
}

type Function struct {
	ReturnType     reflect.Type
	Arguments      []FunctionArg
	Implementation reflect.Value
}

func NewNativeFunction(argNames []string, impl any) Function {
	implementation := reflect.ValueOf(impl)
	if implementation.Kind() != reflect.Func {
		panic("A function was not passed to NewFunction!")
	}
	implementationType := implementation.Type()
	f := Function{
		Implementation: implementation,
		ReturnType:     reflect.TypeOf(nil),
		Arguments:      make([]FunctionArg, 0),
	}
	if implementationType.NumOut() > 0 {
		f.ReturnType = implementationType.Out(0)
	}
	for i := 0; i < implementationType.NumIn(); i++ {
		f.Arguments = append(f.Arguments, FunctionArg{
			Name: argNames[i],
			Type: implementationType.In(i),
		})
	}
	return f
}

func NewExprFunction(run *Runtime, name string, args []FunctionArg, expr Expr) Function {
	stateData := make(map[string]any)
	for _, arg := range args {
		stateData[arg.Name] = valueOf(arg.Type)
	}

	runner := RunC.Compile(expr)
	inspector := InspectC.Compile(expr)

	innerState := run.NewState(stateData)
	returnTypes := inspector.ReturnType(innerState)

	f := Function{}
	f.Arguments = args
	if len(returnTypes) > 1 {
		f.ReturnType = returnTypes.ToSlice()[0]
	} else {
		f.ReturnType = typesVoid.ToSlice()[0]
	}

	f.Implementation = reflect.ValueOf(func(values ...any) (any, error) {
		for i, arg := range args {
			stateData[arg.Name] = values[i]
		}
		return runner.Run(innerState)
	})
	return f
}

type Runtime struct {
	Data              Ref
	Funcs             map[string]Function
	Epsilon           float64
	Insensitive       bool
	InsensitiveFields bool
	MaxIterations     int
	AssertionLimit    int
}

func NewRuntime(globalData any) *Runtime {
	return &Runtime{
		Data:              refOf(globalData),
		Funcs:             make(map[string]Function),
		Epsilon:           0,
		Insensitive:       false,
		InsensitiveFields: false,
		MaxIterations:     1_000_000,
		AssertionLimit:    1,
	}
}

func (run Runtime) Clone() *Runtime {
	copy := run
	copy.Data = run.Data.Clone()
	copy.Funcs = clone(run.Funcs).Interface().(map[string]Function)
	return &copy
}

func (run *Runtime) AddSystemFunc(name string, argNames []string, impl any) Function {
	f := NewNativeFunction(argNames, impl)
	run.Funcs[name] = f
	return f
}

func (run *Runtime) AddUserFunc(name string, args []FunctionArg, expr Expr) Function {
	f := NewExprFunction(run, name, args, expr)
	run.Funcs[name] = f
	return f
}

func (run *Runtime) NewDetachedState(data any) *State {
	return run.Clone().NewState(data)
}

func (run *Runtime) NewState(data any) *State {
	return &State{
		Data:       refOf(data),
		Scopes:     make([]*Ref, 0),
		Assertions: make([]ExprError, 0),
		Runtime:    run,
	}
}

type State struct {
	Data       Ref
	Scopes     []*Ref
	Assertions []ExprError
	Runtime    *Runtime
}

func (state *State) Ref(next any) *Ref {
	scopeLast := len(state.Scopes) - 1
	for i := scopeLast; i >= 0; i-- {
		scope := state.Scopes[i]
		scopeRef := scope.Ref(next)
		if scopeRef.Valid() && scopeRef.Exists {
			return scopeRef
		}
	}
	dataRef := state.Data.Ref(next)
	if dataRef.Valid() && dataRef.Exists {
		return dataRef
	}
	runRef := state.Runtime.Data.Ref(next)
	if runRef.Valid() && runRef.Exists {
		return runRef
	}
	if scopeLast >= 0 {
		return state.Scopes[scopeLast].Ref(next)
	}
	return &Ref{}
}

func (state *State) EnterScope() *Ref {
	entered := refOf(map[string]any{})
	state.Scopes = append(state.Scopes, &entered)
	return &entered
}

func (state *State) ExitScope() {
	state.Scopes = state.Scopes[:len(state.Scopes)-1]
}

func (state *State) Assert(assertion ExprError) error {
	state.Assertions = append(state.Assertions, assertion)
	if state.Runtime.AssertionLimit > 0 && len(state.Assertions) >= state.Runtime.AssertionLimit {
		return assertion
	}
	return nil
}
