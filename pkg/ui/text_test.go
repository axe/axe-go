package ui

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	output := TextFormatRegex.FindAllStringSubmatch("Hello {f:roboto}World{f:} {s:20}HOW{s:} {c:#00F}blue {c:red}red {c:} \\{escaped}", -1)
	js, _ := json.MarshalIndent(output, "", "  ")
	fmt.Printf("%s\n", string(js))
}
