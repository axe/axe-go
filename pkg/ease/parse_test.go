package ease

import "testing"

func TestParse(t *testing.T) {
	tests := []string{
		"linear",
		"ease",
		"cubic-in",
		"cubic-out",
		"cubic-inout",
		"cubic-yoyo",
		"cubic*2",
		"linear-yoyo",
		"ease-inout*1.5",
		"ease-inout * 1.5",
		"bezier(0,1,1,0)",
		"bezier(0,1,2,3)",
		"bezier(0,.1,2.3,23)",
		"bezier(0,1,2,3)*.3",
	}

	for _, test := range tests {
		MustParse(test)
	}
}
