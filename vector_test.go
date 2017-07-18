// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import "testing"

func TestNewVector(t *testing.T) {

	const size = 5

	a := NewVector(size)
	if act, exp := len(a), size; act != exp {
		t.Fatalf("a := NewVector(size); len(a) = %d, expected %d", act, exp)
	}
}

func TestVector(t *testing.T) {

}

func TestMatrix(t *testing.T) {

}

func TestFill(t *testing.T) {
	const value = 1.0
	a := NewVector(5)
	a.Fill(value)
	for i := range a {
		if a[i] != value {
			t.Fatalf("a[i] = %f, expected %f", a[i], value)
		}
	}
}

func TestResize(t *testing.T) {
	const size = 10
	a := NewVector(5)
	a.Resize(size)
	if act, exp := len(a), size; act != exp {
		t.Fatalf("a.Resize(size); len(a) = %d, expected %d", act, exp)
	}
}

func TestSum(t *testing.T) {

	v := NewVector(10)
	for i := range v {
		v[i] = 1.0
	}
	if act, exp := v.Sum(), 10.0; act != exp {
		t.Fatalf("v.Resize(size); len(v) = %v, expected %v", act, exp)
	}
}
