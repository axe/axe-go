package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	a := Vec2{X: 0, Y: 1}
	b := Vec2{X: 2, Y: 3}
	a.Add(b)
	fmt.Printf("%v\n", a)
}

func proxy(method int) func(args ...any) {
	return func(args ...any) {
		a := make([]any, 0)
		a = append(a, method)
		a = append(a, args...)
		fmt.Printf(strings.Repeat("%v, ", len(args))+"%v", a...)
	}
}

func TestProxy(t *testing.T) {
	sendPoints := proxy(23)

	sendPoints("TEXT", 1, true, "\n")
}
