//go:build ignore

package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/axe/axe-go/pkg/ui"
)

//go:embed colors.json
var colorJson string

func main() {
	colors := map[string]string{}
	_ = json.Unmarshal([]byte(colorJson), &colors)

	out := strings.Builder{}
	out.WriteString("package ui\n\nimport \"strings\"\n\nvar (\n")

	for hex, name := range colors {
		color := ui.ColorFromHex(hex)
		varName := strings.ReplaceAll(name, " ", "")
		a := "1"
		if color.A != 1 {
			a = fmt.Sprintf("%.2f", color.A)
		}

		out.WriteString(fmt.Sprintf("\tColor%s = Color{R: %.3f, G: %.3f, B: %.3f, A: %s}\n", varName, color.R, color.G, color.B, a))
	}

	out.WriteString("\tColorTransparent = Color{}\n")
	out.WriteString("\tColorMap = map[string]Color{\n")
	for _, name := range colors {
		varName := strings.ReplaceAll(name, " ", "")
		key := strings.ToLower(varName)
		out.WriteString(fmt.Sprintf("\t\t\"%s\": Color%s,\n", key, varName))
	}
	out.WriteString("\t\t\"transparent\": ColorTransparent,\n")
	out.WriteString("\t}\n")
	out.WriteString(")\n")
	out.WriteString("\n")
	out.WriteString("func ColorNamed(name string) Color {\n")
	out.WriteString("\treturn ColorMap[strings.ToLower(name)]\n")
	out.WriteString("}\n\n")

	os.WriteFile("colors_gen.go", []byte(out.String()), os.ModeType)
}
