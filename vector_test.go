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

func ExampleVector_Add() {
	va := Vector{1, 2, 3}
	vb := Vector{4, 5, 6}
	va.Add(vb)
	fmt.Println(va)
	// Output: [5 7 9]
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

func TestVector_Dot(t *testing.T) {
	va := Vector{1, 2, 3}
	vb := Vector{4, 5, 6}
	if act, exp := va.Dot(vb), 32.0; act != exp {
		t.Fatalf("va.Dot(vb) = %v, expected %v", act, exp)
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

func ExampleVector_Fill() {
	v := NewVector(5)
	v.Fill(1.0)
	fmt.Println(v)
	// Output: [1 1 1 1 1]
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

func ExampleVector_Normalize() {
	v := Vector{6, 8}
	v.Normalize()
	fmt.Println(v)
	// Output: [0.6000000000000001 0.8]
}

func ExampleVector_Product() {
	v := Vector{1, 2, 3, 4, 5}
	fmt.Println(v.Product())
	// Output: 120
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

func ExampleVector_Scale() {
	v := Vector{2, 3, 4}
	v.Scale(5)
	fmt.Println(v)
	// Output: [10 15 20]
}

func ExampleVector_Size() {
	v := Vector{1, 2, 3, 4, 5}
	fmt.Println(v.Size())
	// Output: [5]
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

func ExampleVector_Transpose() {

	rows, columns := 3, 5
	v := Vector{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		10, 11, 12, 13, 14,
	}
	A := NewMatrix(v, rows, columns)
	fmt.Println(A)
	B := NewMatrix(v, columns, rows)
	v.Transpose(rows, columns)
	fmt.Println(B)
	// Output:
	// [[0 1 2 3 4] [5 6 7 8 9] [10 11 12 13 14]]
	// [[0 5 10] [1 6 11] [2 7 12] [3 8 13] [4 9 14]]
}
