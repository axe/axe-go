package axe

import (
	"math"
	"math/bits"
)

type Matrix[A Attr[A]] struct {
	columns []A
}

type Matrix2f = Matrix[Vec2f]
type Matrix3f = Matrix[Vec3f]
type Matrix4f = Matrix[Vec4f]

// type Mat2d = Matrix[Vec3f]
// type Mat3d = Matrix[Vec4f]

// type Matrix4f = Matrix[Vec4f, *Vec4f]

func NewMatrix[A Attr[A]]() Matrix[A] {
	return InitMatrix(Matrix[A]{})
}

func InitMatrix[A Attr[A]](m Matrix[A]) Matrix[A] {
	var empty A
	m.columns = make([]A, empty.Components())
	return m
}

func (m *Matrix[A]) SetValues(values [][]float32) {
	n := m.Size()
	for c := 0; c < n; c++ {
		for r := 0; r < n; r++ {
			if len(values) > r && len(values[r]) > c {
				m.columns[c].SetComponent(r, values[r][c], &m.columns[c])
			}
		}
	}
}

func (m Matrix[A]) GetValues() [][]float32 {
	n := m.Size()
	values := make([][]float32, n)
	for r := 0; r < n; r++ {
		values[r] = make([]float32, n)
		for c := 0; c < n; c++ {
			values[r][c] = m.columns[c].GetComponent(r)
		}
	}
	return values
}

func (m *Matrix[A]) Col(index int) A {
	return m.columns[index]
}

func (m *Matrix[A]) SetCol(index int, col A) {
	m.columns[index] = col
}

func (m *Matrix[A]) Row(index int) A {
	var row A
	for i := 0; i < row.Components(); i++ {
		row.SetComponent(i, m.columns[i].GetComponent(index), &row)
	}
	return row
}

func (m *Matrix[A]) SetRow(index int, row A) {
	for i := 0; i < row.Components(); i++ {
		m.columns[i].SetComponent(index, row.GetComponent(i), &m.columns[i])
	}
}

func (m *Matrix[A]) Zero() {
	for _, c := range m.columns {
		c.SetComponents(0, &c)
	}
}

func (m *Matrix[A]) Identity() {
	for i, c := range m.columns {
		c.SetComponents(0, &c)
		c.SetComponent(i, 1, &c)
	}
}

func (m *Matrix[A]) Set(other Matrix[A]) {
	n := other.Size()
	for i := 0; i < n; i++ {
		m.columns[i] = other.columns[i]
	}
}

func (m *Matrix[A]) Mul(base Matrix[A], other Matrix[A]) {
	n := other.Size()
	for c := 0; c < n; c++ {
		otherCol := other.Col(c)
		for r := 0; r < n; r++ {
			baseRow := base.Row(r)
			dot := baseRow.Dot(otherCol)
			baseRow.SetComponent(r, dot, &m.columns[c])
		}
	}
}

func (m *Matrix[A]) Transpose(other Matrix[A]) {
	n := other.Size()
	for c := 0; c < n; c++ {
		m.columns[c] = other.Row(c)
	}
}

func (m *Matrix[A]) Determinant() float32 {
	all := uint((1 << m.Size()) - 1)
	return m.subDeterminant(all, all)
}

func (m *Matrix[A]) subDeterminant(columns uint, rows uint) float32 {
	size := bits.OnesCount(columns)
	if size == 1 {
		c0 := bits.TrailingZeros(columns)
		r0 := bits.TrailingZeros(rows)

		return m.columns[c0].GetComponent(r0)
	} else if size == 2 {
		c0 := bits.TrailingZeros(columns)
		c1 := bits.TrailingZeros(columns ^ (1 << c0))
		r0 := bits.TrailingZeros(rows)
		r1 := bits.TrailingZeros(rows ^ (1 << r0))

		a := m.columns[c0].GetComponent(r0)
		b := m.columns[c1].GetComponent(r0)
		c := m.columns[c0].GetComponent(r1)
		d := m.columns[c1].GetComponent(r1)

		return a*d - b*c
	} else {
		d := float32(0)
		handledColumns := columns
		rowIndex := bits.TrailingZeros(rows)
		rowBit := uint(1) << rowIndex
		otherRows := rows ^ rowBit
		for i := 0; i < size; i++ {
			columnIndex := bits.TrailingZeros(handledColumns)
			columnBit := uint(1) << columnIndex
			otherColumns := columns ^ columnBit
			a := m.columns[columnIndex].GetComponent(rowIndex)
			b := m.subDeterminant(otherColumns, otherRows)
			c := a * b
			if (i & 1) == 1 {
				c = -c
			}
			d += c
			handledColumns ^= columnBit
		}
		return d
	}
}

func (m *Matrix[A]) Minor(other Matrix[A]) {
	var empty A
	n := other.Size()
	all := uint((1 << n) - 1)

	for c := 0; c < n; c++ {
		columns := all ^ (1 << c)
		for r := 0; r < n; r++ {
			rows := all ^ (1 << r)
			determinant := other.subDeterminant(columns, rows)
			empty.SetComponent(r, determinant, &m.columns[c])
		}
	}
}

func (m *Matrix[A]) Cofactor(other Matrix[A]) {
	m.Minor(other)
	n := other.Size()
	for c := 0; c < n; c++ {
		col := m.columns[c]
		for r := 0; r < n; r++ {
			negative := (r+c)&1 == 1
			if negative {
				col.SetComponent(r, -col.GetComponent(r), &m.columns[c])
			}
		}
	}
}

func (m *Matrix[A]) Adjoint(other Matrix[A]) {
	temp := NewMatrix[A]()
	temp.Cofactor(other)
	m.Transpose(temp)
}

func (m *Matrix[A]) Invert(other Matrix[A]) {
	d := other.Determinant()
	if d == 0 {
		m.Zero()
	} else {
		m.Adjoint(other)

		n := other.Size()
		for c := 0; c < n; c++ {
			col := m.columns[c]
			for r := 0; r < n; r++ {
				col.SetComponent(r, col.GetComponent(r)/d, &m.columns[c])
			}
		}
	}
}

func (m *Matrix[A]) SetAxisRotaton(radians float32, axis int, hasTranslation bool) {
	var empty A
	cos := float32(math.Cos(float64(radians)))
	sin := float32(math.Sin(float64(radians)))
	n := m.Size()
	if hasTranslation {
		n--
	}
	pattern := []float32{cos, sin, -sin}
	patternN := len(pattern)
	for c := 0; c < n; c++ {
		for r := 0; r < n; r++ {
			if c == axis || r == axis {
				if c == r {
					empty.SetComponent(r, 1, &m.columns[c])
				} else {
					empty.SetComponent(r, 0, &m.columns[c])
				}
			} else {
				comp := pattern[(r-c+patternN)%patternN]
				empty.SetComponent(r, comp, &m.columns[c])
			}
		}
	}
}

func (m *Matrix[A]) SetRotaton(radians A, hasTranslation bool) {
	temp1 := NewMatrix[A]()
	temp2 := NewMatrix[A]()
	n := m.Size()
	if hasTranslation {
		n--
	}
	m.Identity()
	for i := 0; i < n; i++ {
		temp1.Set(*m)
		temp2.SetAxisRotaton(radians.GetComponent(i), i, hasTranslation)
		m.Mul(temp1, temp2)
	}
}

func (m *Matrix[A]) Rotate(radians A, hasTranslation bool) {
	temp1 := NewMatrix[A]()
	temp1.Set(*m)
	temp2 := NewMatrix[A]()
	temp2.SetRotaton(radians, hasTranslation)
	m.Mul(temp1, temp2)
}

func (m *Matrix[A]) SetScale(scale A) {
	n := m.Size()
	for c := 0; c < n; c++ {
		for r := 0; r < n; r++ {
			if c == r {
				scale.SetComponent(r, scale.GetComponent(r), &m.columns[c])
			} else {
				scale.SetComponent(r, 0, &m.columns[c])
			}
		}
	}
}

func (m *Matrix[A]) Scale(scale A) {
	temp1 := NewMatrix[A]()
	temp1.Set(*m)
	temp2 := NewMatrix[A]()
	temp2.SetScale(scale)
	m.Mul(temp1, temp2)
}

func (m *Matrix[A]) SetTranslation(translation A) {
	m.Identity()
	m.SetCol(m.Size()-1, translation)
}

func (m *Matrix[A]) Translate(translation A) {
	temp1 := NewMatrix[A]()
	temp1.Set(*m)
	temp2 := NewMatrix[A]()
	temp2.SetTranslation(translation)
	m.Mul(temp1, temp2)
}

func (m *Matrix[A]) PostTranslate(translation A) {
	m.SetRow(m.Size()-1, translation)
}

func (m Matrix[A]) Transform(point A) A {
	var transformed A
	n := point.Components()
	for i := 0; i < n; i++ {
		transformed.SetComponent(i, point.Dot(m.Row(i)), &transformed)
	}
	return transformed
}

func (m Matrix[A]) TransformVector(point A) A {
	var transformed A
	n := point.Components() - 1
	for i := 0; i < n; i++ {
		transformed.SetComponent(i, point.Dot(m.Row(i)), &transformed)
	}
	return transformed
}

func (m Matrix[A]) Size() int {
	return len(m.columns)
}
