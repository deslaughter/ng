// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

// Matrix represents a 2-D array as a slice of slices. A matrix is created
// as a view of a Vector.
type Matrix [][]float64

// NewMatrix returns a 2-D view of the vector with the given number of rows and
// columns. The vector is interpreted in row-major order. A panic will occur
// if rows * columns > Vector.Size().
func NewMatrix(v Vector, rows, columns int) Matrix {
	d := v[:rows*columns]
	m := make([][]float64, rows)
	for i := range m {
		m[i], d = d[:columns:columns], d[columns:]
	}
	return m
}

// Dimensions returns the number of dimensions in the matrix which is 2.
func (A Matrix) Dimensions() int {
	return 2
}

// Multiply performs matrix multiplication AB = C and returns a vector
// representing matrix C with the corresponding number of rows and columns.
func (A Matrix) Multiply(B Matrix) (v Vector, rows int, columns int) {
	rows, columns = len(A), len(B[0])
	v = make(Vector, rows*columns)
	m := NewMatrix(v, rows, columns)
	for i, aRow := range A {
		mRow := m[i]
		for k, bRow := range B {
			for j := range bRow {
				mRow[j] += aRow[k] * bRow[j]
			}
		}
	}
	return v, rows, columns
}

// Size returns of slice of integers containing the length of each dimension
// of the matrix. Since there are two dimensions, the slice contains two
// elements: the number of rows and the number of columns.
func (A Matrix) Size() []int {
	return []int{len(A), len(A[0])}
}

// Sum returns the sum of all elements in the matrix.
func (A Matrix) Sum() float64 {
	sum := 0.0
	for _, row := range A {
		for _, v := range row {
			sum += v
		}
	}
	return sum
}
