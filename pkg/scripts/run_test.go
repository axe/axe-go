package scripts

import (
	"os"
	"reflect"
	"testing"
)

var run *Runtime = NewRuntime(struct {
	GlobalVar string
}{
	GlobalVar: "Hi!",
})

func TestMain(m *testing.M) {
	run.AddSystemFunc("negate", []string{"value"}, func(value int) int {
		return -value
	})
	run.AddSystemFunc("add", []string{"a", "b"}, func(a int, b int) int {
		return a + b
	})
	run.AddUserFunc("sayHello",
		[]FunctionArg{
			{Name: "name", Type: reflect.TypeOf("")},
		},
		Template{
			Format: Constant{"Hello {{.Name}}, how are you?"},
			Vars: map[string]Expr{
				"Name": Get{[]any{"name"}},
			},
		},
	)
	os.Exit(m.Run())
}

func TestCommon(t *testing.T) {
	cases := []struct {
		name           string
		data           any
		expr           Expr
		expectedResult any
		expectedError  error
		expectedData   any
	}{
		{
			name: "Invoke system",
			data: struct {
				Counter int
			}{
				Counter: 2,
			},
			expr: Invoke{
				Function: "negate",
				Params: map[string]Expr{
					"value": Get{[]any{"Counter"}},
				},
			},
			expectedResult: -2,
		},
		{
			name: "Template",
			data: struct {
				Counter int
			}{
				Counter: 33,
			},
			expr: Template{
				Format: Constant{"I am {{.X}} years old"},
				Vars: map[string]Expr{
					"X": Get{[]any{"Counter"}},
				},
			},
			expectedResult: "I am 33 years old",
		},
		{
			name: "Invoke user",
			data: struct{}{},
			expr: Invoke{
				Function: "sayHello",
				Params: map[string]Expr{
					"name": Constant{"Phil"},
				},
			},
			expectedResult: "Hello Phil, how are you?",
		},
		{
			name: "For",
			data: &struct {
				Counter int
			}{
				Counter: 0,
			},
			expr: Loop{
				Start: Set{[]any{"i"}, Constant{1}},
				While: Compare{Get{[]any{"Counter"}}, LT, Constant{4}},
				Then: Set{
					[]any{"Counter"},
					Invoke{
						Function: "add",
						Params: map[string]Expr{
							"a": Get{[]any{"Counter"}},
							"b": Get{[]any{"i"}},
						},
					},
				},
			},
			expectedData: struct {
				Counter int
			}{
				Counter: 4,
			},
		},
		{
			name: "And true",
			data: struct {
				A bool
				B bool
			}{
				A: true,
				B: true,
			},
			expr:           And{[]Expr{Get{[]any{"A"}}, Get{[]any{"B"}}}},
			expectedResult: true,
		},
		{
			name: "And false",
			data: struct {
				A bool
				B bool
			}{
				A: true,
				B: false,
			},
			expr:           And{[]Expr{Get{[]any{"A"}}, Get{[]any{"B"}}}},
			expectedResult: false,
		},
		{
			name: "Or true",
			data: struct {
				A bool
				B bool
			}{
				A: false,
				B: true,
			},
			expr:           Or{[]Expr{Get{[]any{"A"}}, Get{[]any{"B"}}}},
			expectedResult: true,
		},
		{
			name: "Or false",
			data: struct {
				A bool
				B bool
			}{
				A: false,
				B: false,
			},
			expr:           Or{[]Expr{Get{[]any{"A"}}, Get{[]any{"B"}}}},
			expectedResult: false,
		},
		{
			name: "Not true",
			data: struct {
				A bool
			}{
				A: true,
			},
			expr:           Not{Get{[]any{"A"}}},
			expectedResult: false,
		},
		{
			name: "Not false",
			data: struct {
				A bool
			}{
				A: false,
			},
			expr:           Not{Get{[]any{"A"}}},
			expectedResult: true,
		},
		{
			name: "Body",
			data: &struct {
				A int
			}{
				A: 0,
			},
			expr: Body{[]Expr{
				Set{[]any{"A"}, Constant{1}},
				Define{
					[]DefineVar{
						{Name: "B", Value: Constant{2}},
					},
					If{[]IfCase{{
						Condition: Compare{Get{[]any{"A"}}, EQ, Constant{1}},
						Then:      Set{[]any{"A"}, Constant{2}},
					}}, nil},
				},
			}},
			expectedData: struct {
				A int
			}{
				A: 2,
			},
		},
	}

	for _, test := range cases {
		actualData := test.data
		s := run.NewState(actualData)
		actual, err := Run(test.expr, s)
		if (err == nil) != (test.expectedError == nil) {
			t.Errorf("Test [%s] expected error %v but got %v", test.name, test.expectedError, err)
		} else {
			actualString, _ := toString(actual)
			expectedString, _ := toString(test.expectedResult)
			if actualString != expectedString {
				t.Errorf("Test [%s] expected result %s but got %s", test.name, expectedString, actualString)
			}
		}
		if test.expectedData != nil {
			actualString, _ := toString(actualData)
			expectedString, _ := toString(test.expectedData)
			if actualString != expectedString {
				t.Errorf("Test [%s] expected data %s but got %s", test.name, expectedString, actualString)
			}
		}
	}
}
