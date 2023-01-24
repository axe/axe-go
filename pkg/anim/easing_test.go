package anim

import "testing"

func TestParse(t *testing.T) {
	tests := []string{
		"",
		"linear",
		"linear-yoyo",
		"ease-inout*1.5",
		"ease-inout * 1.5",
		"bezier(0,1,1,0)",
	}

	for _, test := range tests {
		testCallback := func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("An error has occurred parsing %s: %v", test, err)
				}
			}()

			easing := ParseEasing(test)

			if easing == nil {
				t.Errorf("No easing could be parsed for %s", test)
			}
		}

		testCallback()
	}
}
