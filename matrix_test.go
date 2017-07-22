// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import (
	"fmt"
	"testing"
)

func TestNewMatrix(t *testing.T) {

	// Number of rows and columns in matrix
	rows, columns := 2, 3

	// Create vector
	v := make(Vector, rows*columns)

	// Fill vector with unique values
	for i := range v {
		v[i] = float64(i)
	}

	// Create matrix view of vector
	m := NewMatrix(v, rows, columns)

	// If number of rows of matrix is not correct
	if act, exp := len(m), rows; act != exp {
		t.Fatalf("len(m) = %v, expected %v", act, exp)
	}

	// If number of columns of matrix is not correct
	if act, exp := len(m[0]), columns; act != exp {
		t.Fatalf("len(m[0]) = %v, expected %v", act, exp)
	}

	// Loop through matrix in row-major order
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {

			// Verify that matrix values match corresponding vector values
			if act, exp := m[i][j], v[i*columns+j]; act != exp {
				t.Fatalf("m[i][j] = %v, expected %v", act, exp)
			}
		}
	}
}

func ExampleNewMatrix() {

	v := Vector{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := NewMatrix(v, 3, 3)

	fmt.Println(m)
	// Output:
	// [[1 2 3] [4 5 6] [7 8 9]]
}

func ExampleMatrix_Dimensions() {

	v := Vector{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := NewMatrix(v, 3, 3)

	fmt.Println(m.Dimensions())
	// Output:
	// 2
}

func TestMatrix_Dimensions(t *testing.T) {
	m := Matrix{}
	if act, exp := m.Dimensions(), 2; act != exp {
		t.Fatalf("m.Dimensions() = %v, expected %v", act, exp)
	}
}

func ExampleMatrix_Multiply_first() {

	// Create vectors
	va := Vector{1, 2, 3}
	vb := Vector{4, 5, 6}

	// Create matrix views of vectors
	mA := NewMatrix(va, len(va), 1) // [3 x 1]
	mB := NewMatrix(vb, 1, len(vb)) // [1 x 3]

	// Multiply matrices and receive vector representing matrix
	// with the given rows and columns
	vc, rows, columns := mA.Multiply(mB)

	// Create matrix view of result vector
	mC := NewMatrix(vc, rows, columns)

	fmt.Println(mA)
	fmt.Println(mB)
	fmt.Println(mC)
	// Output:
	// [[1] [2] [3]]
	// [[4 5 6]]
	// [[4 5 6] [8 10 12] [12 15 18]]
}

func ExampleMatrix_Multiply_second() {

	va := Vector{1, 2, 3}
	vb := Vector{4, 5, 6}

	mA := NewMatrix(va, 1, len(va))
	mB := NewMatrix(vb, len(vb), 1)

	vc, rows, columns := mA.Multiply(mB)

	mC := NewMatrix(vc, rows, columns)

	fmt.Println(mC)
	// Output:
	// [[32]]
}

func ExampleMatrix_Size() {

	v := Vector{1, 2, 3, 4}
	m := NewMatrix(v, 2, 2)

	fmt.Println(m)
	fmt.Println(m.Size())
	// Output:
	// [[1 2] [3 4]]
	// [2 2]
}

func ExampleMatrix_Sum() {

	v := Vector{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := NewMatrix(v, 3, 3)

	fmt.Println(m)
	fmt.Println(m.Sum())
	// Output:
	// [[1 2 3] [4 5 6] [7 8 9]]
	// 45
}

func ExampleMatrix_SumRows() {

	v := Vector{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := NewMatrix(v, 3, 3)

	fmt.Println(m)
	fmt.Println(m.SumRows())
	// Output:
	// [[1 2 3] [4 5 6] [7 8 9]]
	// [6 15 24]
}

func ExampleMatrix_SumColumns() {

	v := Vector{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := NewMatrix(v, 3, 3)

	fmt.Println(m)
	fmt.Println(m.SumColumns())
	// Output:
	// [[1 2 3] [4 5 6] [7 8 9]]
	// [12 15 18]
}
