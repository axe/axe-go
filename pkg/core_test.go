package axe

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	tests := []struct {
		name      string
		getActual func() any
		expected  any
	}{
		{
			name: "determinant 2x2",
			getActual: func() any {
				m := InitMatrix(Matrix2f{})
				m.SetValues([][]float32{
					{4, 6},
					{3, 8},
				})
				return m.Determinant()
			},
			expected: 14,
		},
		{
			name: "determinant 3x3",
			getActual: func() any {
				m := InitMatrix(Matrix3f{})
				m.SetValues([][]float32{
					{6, 1, 1},
					{4, -2, 5},
					{2, 8, 7},
				})
				return m.Determinant()
			},
			expected: -306,
		},
		{
			name: "minor 2x2",
			getActual: func() any {
				m := InitMatrix(Matrix2f{})
				m.SetValues([][]float32{
					{3, 6},
					{-4, 8},
				})

				minor := InitMatrix(Matrix2f{})
				minor.Minor(m)
				return minor.GetValues()
			},
			expected: [][]float32{
				{8, -4},
				{6, 3},
			},
		},
	}

	for _, test := range tests {
		actual := test.getActual()
		actualText := fmt.Sprintf("%+v", actual)
		expectedText := fmt.Sprintf("%+v", test.expected)
		if actualText != expectedText {
			t.Errorf("%s: expected %s but got %s", test.name, expectedText, actualText)
		}
	}
}
