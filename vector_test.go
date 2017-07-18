// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import "testing"

func TestNewVector(t *testing.T) {
	const size = 5
	v := NewVector(size)
	if act, exp := len(v), size; act != exp {
		t.Fatalf("a := NewVector(size); len(a) = %d, expected %d", act, exp)
	}
}

func TestFill(t *testing.T) {
	const value = 1.0
	v := NewVector(5)
	v.Fill(value)
	for i := range v {
		if v[i] != value {
			t.Fatalf("a[i] = %f, expected %f", v[i], value)
		}
	}
}

func TestResize(t *testing.T) {
	const size = 10
	v := NewVector(5)
	v.Resize(size)
	if act, exp := len(v), size; act != exp {
		t.Fatalf("a.Resize(size); len(a) = %d, expected %d", act, exp)
	}
}

func TestSum(t *testing.T) {

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
