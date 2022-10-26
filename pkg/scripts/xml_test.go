package scripts

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	x := `
	<if>
		<case>
		  <compare op="=">
				<get>x</get>
				<get>y</get>
			</compare>
			<invoke func="add">
        <a><get>x</get></a>
				<b><get>y</get></b>
			</invoke>
		</case>
		<else>
		  <set>
			  <path>x</path>
				<value>
					<invoke func="add">
						<a><get>x</get></a>
						<b><constant type="int">1</constant></b>
					</invoke>
				</value>
			</set>
		</else>
	</if>
	`

	expr, err := FromXmlString(x)

	fmt.Printf("%+v %v\n", expr, err)
}

func TestEncode(t *testing.T) {
	expr := Loop{
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
	}

	xml, err := ToXmlString(expr, "  ")

	fmt.Printf("XML: %s\nError: %v\n", xml, err)
}
