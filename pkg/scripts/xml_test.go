package scripts

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	x := `
	<if>
		<and></and>
		<body>
			<and></and>
			<and></and>
		</body>
	</if>

	`

	decoder := xml.NewDecoder(strings.NewReader(x))

	e, err := UnmarshalXML(decoder)

	fmt.Printf("%+v %v\n", e, err)
}
