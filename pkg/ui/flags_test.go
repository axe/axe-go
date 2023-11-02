package ui_test

import (
	"fmt"
	"testing"

	"github.com/axe/axe-go/pkg/ui"
)

func TestFlagsTake(t *testing.T) {
	var flags ui.Flags = 2 | 4 | 16 | 128
	fmt.Printf("%d\n", flags.Take())
	fmt.Printf("%d\n", flags.Take())
	fmt.Printf("%d\n", flags.Take())
	fmt.Printf("%d\n", flags.Take())
	fmt.Printf("%d\n", flags.Take())
}
