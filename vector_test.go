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

func TestVector_Dimensions(t *testing.T) {
	v := Vector{}
	if act, exp := v.Dimensions(), 1; act != exp {
		t.Fatalf("v.Dimensions() = %v, expected %v", act, exp)
	}
}

func ExampleVector_Dimensions() {
	v := Vector{1, 2, 3, 4, 5}
	fmt.Println(v.Dimensions())
	// Output: 1
}

func TestVector_Size(t *testing.T) {
	v := Vector{1, 2, 3, 4, 5}
	if act, exp := len(v.Size()), 1; act != exp {
		t.Fatalf("len(v.Size()) = %v, expected %v", act, exp)
	}
	if act, exp := v.Size()[0], len(v); act != exp {
		t.Fatalf("v.Size()[0] = %v, expected %v", act, exp)
	}
}

func ExampleVector_Size() {
	v := Vector{1, 2, 3, 4, 5}
	fmt.Println(v.Size())
	// Output: [5]
}

func TestVector_Dot(t *testing.T) {
	va := Vector{1, 2, 3}
	vb := Vector{4, 5, 6}
	if act, exp := va.Dot(vb), 32.0; act != exp {
		t.Fatalf("va.Dot(vb) = %v, expected %v", act, exp)
	}
}

func TestVector_Magnitude(t *testing.T) {
	v := Vector{2, 6, 9}
	if act, exp := v.Magnitude(), 11.0; act != exp {
		t.Fatalf("v.Magnitude() = %v, expected %v", act, exp)
	}
}

func ExampleVector_Magnitude() {
	v := Vector{2, 6, 9}
	fmt.Println(v.Magnitude())
	// Output: 11
}

func ExampleVector_Scale() {
	v := Vector{2, 3, 4}
	v.Scale(5)
	fmt.Println(v)
	// Output: [10 15 20]
}

func ExampleVector_Add() {
	va := Vector{1, 2, 3}
	vb := Vector{4, 5, 6}
	va.Add(vb)
	fmt.Println(va)
	// Output: [5 7 9]
}

func ExampleVector_Normalize() {
	v := Vector{6, 8}
	v.Normalize()
	fmt.Println(v)
	// Output: [0.6000000000000001 0.8]
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

func ExampleVector_Fill() {
	v := NewVector(5)
	v.Fill(1.0)
	fmt.Println(v)
	// Output: [1 1 1 1 1]
}

func TestVector_Resize(t *testing.T) {
	const size = 10
	v := NewVector(5)
	v.Resize(size)
	if act, exp := len(v), size; act != exp {
		t.Fatalf("a.Resize(size); len(a) = %d, expected %d", act, exp)
	}
}

func ExampleVector_Resize() {
	v := NewVector(5)
	v.Fill(5)
	fmt.Println(v)
	v.Resize(10)
	fmt.Println(v)
	v.Resize(3)
	fmt.Println(v)
	// Output:
	// [5 5 5 5 5]
	// [5 5 5 5 5 0 0 0 0 0]
	// [5 5 5]
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
