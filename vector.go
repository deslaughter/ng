// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import (
	"math"
)

// Vector is the base type of ng
type Vector []float64

// NewVector returns a vector with the given number of elements.
func NewVector(size int) Vector {
	return make([]float64, size)
}

// Matrix returns a 2-D view of the vector with the given number of rows and
// columns. The vector is interpreted in row-major order. A panic will occur
// if rows * columns > Vector.Size().
func (v Vector) Matrix(rows, columns int) Matrix {
	d := v[:rows*columns]
	m := make([][]float64, rows)
	for i := range m {
		m[i], d = d[:columns:columns], d[columns:]
	}
	return m
}

// Dimensions returns the number of dimensions in the vector which is 1.
func (v Vector) Dimensions() int {
	return 1
}

// Size returns of slice of integers containing the length of each dimension
// of the vector. Since there is only one dimension, the first element of
// the slice will contain the same value as Vector.Size().
func (v Vector) Size() []int {
	return []int{len(v)}
}

// Dot returns the dot product of this vector and the given vector.
// This function will panic if vn.Size() < v.Size().
func (v Vector) Dot(vn Vector) float64 {
	sum := 0.0
	vt := vn[:len(v)]
	for i := range v {
		sum += v[i] * vt[i]
	}
	return sum
}

// Magnitude returns the sum of each element squared.
func (v Vector) Magnitude() float64 {
	sum := 0.0
	for i := range v {
		sum += v[i] * v[i]
	}
	return math.Sqrt(sum)
}

// Scale multiplies all elements of the vector by the given value.
func (v Vector) Scale(s float64) {
	for i := range v {
		v[i] *= s
	}
}

// Add sums each element of the given vector with the current vector and stores
// the resulting value in the current vector.
func (v Vector) Add(vn Vector) {
	vt := vn[:len(v)]
	for i := range v {
		v[i] += vt[i]
	}
}

// Normalize scales the vector by one over the magnitude of the vector such that
// the vector's magnitude is one. Transforms vector into the unit vector.
func (v Vector) Normalize() {
	v.Scale(1 / v.Magnitude())
}

// Fill sets all elements of the vector to the given value.
func (v Vector) Fill(value float64) {
	for i := range v {
		v[i] = value
	}
}

// Resize changes the size and capacity of the vector. If the capacity is
// less than size, a new vector is allocated and the values are copied into it.
func (v *Vector) Resize(size int) {
	if size <= cap(*v) {
		*v = (*v)[:size:size]
	}
	*v = append(*v, make([]float64, size-cap(*v))...)
}

// Sum returns the sum of all values in the vector.
func (v Vector) Sum() float64 {
	sum := 0.0
	for _, value := range v {
		sum += value
	}
	return sum
}
