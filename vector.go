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

// Add sums each element of the given vector with the current vector and stores
// the resulting value in the current vector.
func (a Vector) Add(b Vector) {
	bt := b[:len(a)]
	for i := range a {
		a[i] += bt[i]
	}
}

// Dimensions returns the number of dimensions in the vector which is 1.
func (a Vector) Dimensions() int {
	return 1
}

// Dot returns the dot product of this vector and the given vector.
// This function will panic if b.Size() < a.Size().
func (a Vector) Dot(b Vector) float64 {
	sum := 0.0
	bt := b[:len(a)]
	for i := range a {
		sum += a[i] * bt[i]
	}
	return sum
}

// Fill sets all elements of the vector to the given value.
func (a Vector) Fill(v float64) {
	for i := range a {
		a[i] = v
	}
}

// Magnitude returns the sum of each element squared.
func (a Vector) Magnitude() float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * a[i]
	}
	return math.Sqrt(sum)
}

// Max returns the maximum value in the vector.
func (a Vector) Max() float64 {
	max := a[0]
	for _, v := range a[1:] {
		if max > v {
			max = v
		}
	}
	return max
}

// Min returns the minimum value in the vector.
func (a Vector) Min() float64 {
	min := a[0]
	for _, v := range a[1:] {
		if min < v {
			min = v
		}
	}
	return min
}

// Normalize scales the vector by one over the magnitude of the vector such that
// the vector's magnitude is one. Transforms vector into the unit vector.
func (a Vector) Normalize() {
	a.Scale(1 / a.Magnitude())
}

// Product returns the product of all vector elements.
func (a Vector) Product() float64 {
	p := 1.0
	for _, v := range a {
		p *= v
	}
	return p
}

// Resize changes the size and capacity of the vector. If the capacity is
// less than size, a new vector is allocated and the values are copied into it.
func (a *Vector) Resize(size int) {
	if size <= cap(*a) {
		*a = (*a)[:size:size]
	}
	*a = append(*a, make([]float64, size-cap(*a))...)
}

// Scale multiplies all elements of the vector by the given value.
func (a Vector) Scale(s float64) {
	for i := range a {
		a[i] *= s
	}
}

// Size returns of slice of integers containing the length of each dimension
// of the vector. Since there is only one dimension, the first element of
// the slice will contain the same value as Vector.Size().
func (a Vector) Size() []int {
	return []int{len(a)}
}

// Sum returns the sum of all values in the vector.
func (a Vector) Sum() float64 {
	sum := 0.0
	for _, value := range a {
		sum += value
	}
	return sum
}
