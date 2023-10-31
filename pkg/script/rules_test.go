package script

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var CodeRules = map[string]RulePattern{
	"name":         `/[a-zA-Z_]\w*/g`,
	"string":       `/"(\\"|[^"])*"/g`,
	"integer":      `/[-+]?(0|[1-9]\d*)/g`,
	"decimal":      `/<integer>?\.\d+/g`,
	"bool":         `(true|false)`,
	"null":         "null",
	"ws":           `/(\s+|<blockComment>)+/n`,
	"args":         `<arg:value>(<ws>?,<ws>?<arg:value>)*`,
	"argsNamed":    `<param:name><ws>?=<ws>?<arg:value>(<ws>?,<ws>?<param:name><ws>?=<ws>?<arg:value>)*`,
	"param":        `<name><ws><type:name>`,
	"params":       `<param>(<ws>?,<ws>?<param>)*`,
	"func":         `\(<ws>?<params>?<ws>?)<ws>?-><ws>?\{<body>\}`,
	"lvalue":       `<variable:name>(\.<variable:name>)*`,
	"call":         `(<ws>?\(<ws>?(<argsNamed>|<args>)?<ws>?\)<ws>?)`,
	"expr":         `<variable:name><call>?`,
	"invoke":       `<variable:name>(<call>?\.<variable:name>)*<call>`,
	"exprs":        `<expr>(<ws>?\.<ws>?<expr>)*`,
	"value":        `(<null>|<bool>|<string>|<decimal>|<integer>|<exprs>|<func>)`,
	"define":       `var<ws><name>(<ws><type:name>)?<ws>?(=<ws>?<value>)?`,
	"if":           `if<ws><condition:value><ws>then<body>end`,
	"while":        `while<ws><condition:value><ws>then<body>end`,
	"each":         `each<ws><item:name><ws>in<ws><collection:exprs><ws>then<body>end`,
	"set":          `<settable:lvalue><ws>?=<ws>?<value>`,
	"return":       `return(<ws><value>)?`,
	"comment":      `///([^\n]*)\n/g`,
	"blockComment": `//\*(\*[^/]|[^*])*\*//g`,
	"line":         `(<define>|<if>|<while>|<each>|<return>|<comment>|<set>|<invoke>)`,
	"body":         `<ws>?(<line><ws>?)*`,
	"code":         `<body>`,
}

var JsonRules = map[string]RulePattern{
	"string":  `/"(\\"|[^"])*"/g`,
	"null":    `null`,
	"boolean": `(true|false)`,
	"number":  `/-?(0|[1-9]\d*)(\.\d+)?([eE][+-]?\d+)?/g`,
	"member":  `\s*<property:string>\s*:<value>`,
	"object":  `\{(<member>(,<member>)*|\s*)\}`,
	"array":   `\[(<value>(,<value>)*|\s*)\]`,
	"value":   `\s*(<null>|<boolean>|<number>|<string>|<array>|<object>)\s*`,
}

func TestCases(t *testing.T) {
	codeSet, err := NewStdRuleSet(CodeRules)
	if err != nil {
		t.Fatal(err)
	}

	jsonSet, err := NewStdRuleSet(JsonRules)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		rules         RuleSet
		name          string
		rule          string
		input         string
		expected      []string
		expectedError string
		test          bool
	}{{
		rules: jsonSet,
		name:  "json many",
		rule:  "value",
		input: `{"x":[1.2, true, null, "hi"]}`,
		expected: []string{
			`object.member.property: "x"`,
			`object.member.value.array.value.number: 1.2`,
			`object.member.value.array.value.boolean: true`,
			`object.member.value.array.value.null: null`,
			`object.member.value.array.value.string: "hi"`,
		},
	}, {
		rules: codeSet,
		name:  "named args",
		rule:  "code",
		input: `point(x=34, y=true, z="hello")`,
		expected: []string{
			`body.line.invoke.variable: point`,
			`body.line.invoke.call.argsNamed.param: x`,
			`body.line.invoke.call.argsNamed.arg.integer: 34`,
			`body.line.invoke.call.argsNamed.param: y`,
			`body.line.invoke.call.argsNamed.arg.bool: true`,
			`body.line.invoke.call.argsNamed.param: z`,
			`body.line.invoke.call.argsNamed.arg.string: "hello"`,
		},
	}, {
		rules: codeSet,
		name:  "define with type no value",
		rule:  "code",
		input: `var pt point`,
		expected: []string{
			`body.line.define.name: pt`,
			`body.line.define.type: point`,
		},
	}, {
		rules: codeSet,
		name:  "define no type with value",
		rule:  "code",
		input: `var pt = 45`,
		expected: []string{
			`body.line.define.name: pt`,
			`body.line.define.value.integer: 45`,
		},
	}, {
		rules: codeSet,
		name:  "define with type with value",
		rule:  "code",
		input: `var pt int = 45`,
		expected: []string{
			`body.line.define.name: pt`,
			`body.line.define.type: int`,
			`body.line.define.value.integer: 45`,
		},
	}, {
		rules: codeSet,
		name:  "basic if",
		rule:  "code",
		input: `
		if true then
			a=5
		end`,
		expected: []string{
			`body.line.if.condition.bool: true`,
			`body.line.if.body.line.set.settable.variable: a`,
			`body.line.if.body.line.set.value.integer: 5`,
		},
	}, {
		rules: codeSet,
		name:  "complex",
		rule:  "code",
		input: `
		if employee.name.eq("Phil") then
			while employee.age.lt(34) then
				employee.age = employee.age.add(1)
			end
			if employee.age.eq(50) then
				return 50
			end
			employee.save()
			return 34
		end
		return -1
		`,
		expected: []string{},
	}}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			tc.rules.Logger = NewFuncLogger(t.Logf)
			// tc.rules.Logger = NewFuncLogger(func(s string, a ...any) {
			// 	// x := fmt.Sprintf(s, a...)
			// 	// if x == "<variable:name> matched in code.code.body.line.expr at (i=5)\n" {
			// 	// 	x = strings.ToLower(x)
			// 	// }
			// 	// t.Log(x)
			// 	// t.Logf(s, a...)
			// 	fmt.Printf(s, a...)
			// })
			// tc.rules.Logger.Quiet(false, true, false, false)

			p, err := tc.rules.ParseRule(tc.rule, tc.input)
			if (err == nil) != (tc.expectedError == "") {
				if err.Error() != tc.expectedError {
					t.Errorf("Error in %s, expected %s but got %v", t.Name(), tc.expectedError, err)
				}
			} else {
				paths := getBottomValues(p, tc.input)
				if d := cmp.Diff(paths, tc.expected); d != "" {
					t.Error(d)
				}
			}
		})
	}
}

func getBottomValues(p *RuleMatch, input string) []string {
	if len(p.Children) == 0 && p.Start != p.End {
		return []string{fmt.Sprintf("%s: %s", p.Path(), p.Range.String(input))}
	}
	var inner []string
	for _, child := range p.Children {
		if v := getBottomValues(child, input); len(v) > 0 {
			inner = append(inner, v...)
		}
	}
	return inner
}
