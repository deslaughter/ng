// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

type Matrix [][]float64

// Dimensions returns the number of dimensions in the matrix which is 2.
func (m Matrix) Dimensions() int {
	return 2
}

// Size returns of slice of integers containing the length of each dimension
// of the matrix. Since there are two dimensions, the slice contains two
// elements: the number of rows and the number of columns.
func (m Matrix) Size() []int {
	return []int{len(m), len(m[0])}
}

// Multiply performs matrix multiplication AB = C and returns a vector
// representing matrix C with the corresponding number of rows and columns.
func (A Matrix) Multiply(B Matrix) (v Vector, rows int, columns int) {
	rows, columns = len(A), len(B[0])
	v = make(Vector, rows*columns)
	m := v.Matrix(rows, columns)
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
