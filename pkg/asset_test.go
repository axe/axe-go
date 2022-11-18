package axe

import "testing"

func TestRelativeURI(t *testing.T) {
	tests := []struct {
		uri      string
		relative string
		expected string
	}{
		{
			uri:      "square.png",
			relative: "square.mtl",
			expected: "square.mtl",
		},
		{
			uri:      "/square.png",
			relative: "square.mtl",
			expected: "/square.mtl",
		},
		{
			uri:      "/inner/square.png",
			relative: "square.mtl",
			expected: "/inner/square.mtl",
		},
		{
			uri:      "/inner/square.png",
			relative: "../square.mtl",
			expected: "/square.mtl",
		},
	}

	source := &LocalAssetSource{}

	for _, test := range tests {
		rel := source.Relative(test.uri, test.relative)

		if rel != test.expected {
			t.Errorf("Expected %s but got %s", test.expected, rel)
		}
	}
}
