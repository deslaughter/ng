// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ng provides numerical methods for the Go programming language.
package ng

// Array is the base type of the ng package.
type Array []float64

func NewArray(size int) Array {
	return make([]float64, size)
}

func (a Array) Fill(value float64) {
	d := a[:]
	for i := range d {
		d[i] = value
	}
}

func (a Array) Vector(columns int) Vector {
	return Vector(a[:columns:columns])
}

func (a Array) Matrix(rows, columns int) Matrix {
	d := a[:rows*columns]
	m := make([][]float64, rows)
	for i := range m {
		m[i] = d[i*rows : (i+1)*rows : columns]
	}
	return m
}
