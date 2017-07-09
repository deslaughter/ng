// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

// Array is the base type of the ng package.
type Array []float64

func NewArray(size int) Array {
	return make([]float64, size)
}

func (a Array) Vector(columns int) Vector {
	return Vector(a[:columns:columns])
}

func (a Array) Matrix(rows, columns int) Matrix {
	d := a[:rows*columns]
	m := make([][]float64, rows)
	for i := range m {
		m[i], d = d[:columns:columns], d[columns:]
	}
	return m
}

func (a Array) Fill(value float64) {
	for i := range a {
		a[i] = value
	}
}

func (a *Array) Resize(size int) {
	if size <= cap(*a) {
		*a = (*a)[:size:size]
	}
	*a = append(*a, make([]float64, size-cap(*a))...)
}

func (a Array) Sum() float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum
}
