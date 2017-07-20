// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import (
	"fmt"
	"testing"
)

func TestNewVector(t *testing.T) {
	const size = 5
	v := NewVector(size)
	if act, exp := len(v), size; act != exp {
		t.Fatalf("a := NewVector(size); len(a) = %d, expected %d", act, exp)
	}
}

func TestVector_Matrix(t *testing.T) {

	// Number of rows and columns in matrix
	rows, columns := 2, 3

	// Create vector
	v := make(Vector, rows*columns)

	// Fill vector with unique values
	for i := range v {
		v[i] = float64(i)
	}

	// Create matrix view of vector
	m := v.Matrix(rows, columns)

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

func TestVector_Fill(t *testing.T) {
	const value = 1.0
	v := NewVector(5)
	v.Fill(value)
	for i := range v {
		if v[i] != value {
			t.Fatalf("a[i] = %f, expected %f", v[i], value)
		}
	}
}

func TestVector_Resize(t *testing.T) {
	const size = 10
	v := NewVector(5)
	v.Resize(size)
	if act, exp := len(v), size; act != exp {
		t.Fatalf("a.Resize(size); len(a) = %d, expected %d", act, exp)
	}
}

func TestVector_Sum(t *testing.T) {

	v := NewVector(10)
	for i := range v {
		v[i] = 1.0
	}
	if act, exp := v.Sum(), float64(len(v)); act != exp {
		t.Fatalf("v.Resize(size); len(v) = %v, expected %v", act, exp)
	}
}

func ExampleVector_Sum() {
	v := Vector{1, 2, 3, 4, 5}
	fmt.Println(v.Sum())
	// Output: 15
}
