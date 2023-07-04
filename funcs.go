package matrigo

import "math"

// Map applies f to every element of the matrix and returns the result.
func Map(m Matrix, f Mapper) Matrix {
	n := New(m.Rows, m.Columns, nil)

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			val := m.Data[i][j]
			n.Data[i][j] = f(val, i, j)
		}
	}

	return n
}

// Fold accumulates the values in a matrix according to a Folder function.
func Fold(m Matrix, f Folder, accumulator float64) float64 {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			accumulator = f(accumulator, m.Data[i][j], i, j)
		}
	}

	return accumulator
}

// Transpose returns the transposed version of the matrix.
func Transpose(m Matrix) Matrix {
	return Map(New(m.Columns, m.Rows, nil),
		func(val float64, x, y int) float64 {
			return m.Data[y][x]
		})
}

// Scale does scalar multiplication.
func Scale(m Matrix, a float64) Matrix {
	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] * a
	})
}

// Divide does scalar division.
func Divide(m Matrix, a float64) Matrix {
	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] / a
	})
}

// Sum gives the sum of the elements in the matrix.
func Sum(m Matrix) float64 {
	return m.Fold(func(accumulator, val float64, x, y int) float64 {
		return accumulator + val
	}, 0)
}

// Det computes the determinant.
func Det(m Matrix) float64 {
	// Base case -> det([[x]]) = x
	if m.Rows == 1 && m.Columns == 1 {
		return m.Data[0][0]
	}

	// Remove 1st Row and n-th column
	f := func(m Matrix, n int) Matrix {
		data := [][]float64{}

		for i, row := range m.Data {
			// Skip first row
			if i == 0 {
				continue
			}
			current := []float64{}
			for j, col := range row {
				// Skip n-th column
				if j == n {
					continue
				}
				current = append(current, col)
			}
			data = append(data, current)
		}

		return New(m.Rows-1, m.Columns-1, data)
	}

	det := 0.0
	for n, v := range m.Data[0] {
		det += math.Pow(-1, float64(n)) * v * Det(f(m, n))
	}
	return det
}

// AddMatrix adds 2 matrices together.
func AddMatrix(m, n Matrix) Matrix {
	if m.Rows != n.Rows || m.Columns != n.Columns {
		panic("matrix: can't add different sized matricies")
	}

	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] + n.Data[x][y]
	})
}

// Add does scalar addition.
func Add(m Matrix, n float64) Matrix {
	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] + n
	})
}

// SubtractMatrix subtracts 2 matrices.
func SubtractMatrix(m, n Matrix) Matrix {
	if m.Rows != n.Rows || m.Columns != n.Columns {
		panic("matrix: can't subtract different sized matricies")
	}

	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] - n.Data[x][y]
	})
}

// Subtract does scalar subtraction.
func Subtract(m Matrix, n float64) Matrix {
	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] - n
	})
}

// HadamardProduct does Hadamard Product (entrywise).
func HadamardProduct(m Matrix, n Matrix) Matrix {
	if m.Columns != n.Columns || m.Rows != n.Rows {
		panic("matrix: matricies must have the same shape")
	}

	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] * n.Data[x][y]
	})
}

// Multiply does matrix product.
func Multiply(m, n Matrix) Matrix {
	if m.Columns != n.Rows {
		panic("matrix: rows must match with columns of matricies")
	}

	return Map(New(m.Rows, n.Columns, nil), func(_ float64, x, y int) float64 {
		sum := 0.0

		for i := 0; i < n.Rows; i++ {
			sum += m.Data[x][i] * n.Data[i][y]
		}

		return sum
	})
}
