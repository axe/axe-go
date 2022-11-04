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
		{
			name: "minor 3x3",
			getActual: func() any {
				m := InitMatrix(Matrix3f{})
				m.SetValues([][]float32{
					{2, -1, 3},
					{0, 5, 2},
					{1, -1, -2},
				})

				actual := InitMatrix(Matrix3f{})
				actual.Minor(m)
				return actual.GetValues()
			},
			expected: [][]float32{
				{-8, -2, -5},
				{5, -7, -1},
				{-17, 4, 10},
			},
		},
		{
			name: "cofactor 3x3",
			getActual: func() any {
				m := InitMatrix(Matrix3f{})
				m.SetValues([][]float32{
					{2, -1, 3},
					{0, 5, 2},
					{1, -1, -2},
				})

				actual := InitMatrix(Matrix3f{})
				actual.Cofactor(m)
				return actual.GetValues()
			},
			expected: [][]float32{
				{-8, 2, -5},
				{-5, -7, 1},
				{-17, -4, 10},
			},
		},
		{
			name: "cofactor 3x3 #2",
			getActual: func() any {
				m := InitMatrix(Matrix3f{})
				m.SetValues([][]float32{
					{1, 2, -1},
					{2, 1, 2},
					{-1, 2, 1},
				})

				actual := InitMatrix(Matrix3f{})
				actual.Cofactor(m)
				return actual.GetValues()
			},
			expected: [][]float32{
				{-3, -4, 5},
				{-4, 0, -4},
				{5, -4, -3},
			},
		},
		{
			name: "transpose 3x3",
			getActual: func() any {
				m := InitMatrix(Matrix3f{})
				m.SetValues([][]float32{
					{2, -1, 3},
					{0, 5, 2},
					{1, -1, -2},
				})

				actual := InitMatrix(Matrix3f{})
				actual.Transpose(m)
				return actual.GetValues()
			},
			expected: [][]float32{
				{2, 0, 1},
				{-1, 5, -1},
				{3, 2, -2},
			},
		},
		{
			name: "adjoint 2x2",
			getActual: func() any {
				m := InitMatrix(Matrix2f{})
				m.SetValues([][]float32{
					{3, 6},
					{-4, 8},
				})
				// minor = 8, -4, 6, 3
				//
				actual := InitMatrix(Matrix2f{})
				actual.Adjoint(m)
				return actual.GetValues()
			},
			expected: [][]float32{
				{8, -6},
				{4, 3},
			},
		},
		{
			name: "adjoint 3x3",
			getActual: func() any {
				m := InitMatrix(Matrix3f{})
				m.SetValues([][]float32{
					{2, -1, 3},
					{0, 5, 2},
					{1, -1, -2},
				})

				actual := InitMatrix(Matrix3f{})
				actual.Adjoint(m)
				return actual.GetValues()
			},
			expected: [][]float32{
				{-8, -5, -17},
				{2, -7, -4},
				{-5, 1, 10},
			},
		},
		{
			name: "inverse 2x2",
			getActual: func() any {
				m := InitMatrix(Matrix2f{})
				m.SetValues([][]float32{
					{3, 6},
					{-4, 8},
				})

				actual := InitMatrix(Matrix2f{})
				actual.Invert(m)
				return actual.GetValues()
			},
			expected: [][]float32{
				{1.0 / 6, -1.0 / 8},
				{1.0 / 12, 1.0 / 16},
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
