// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import (
	"math"
)

// Vector is the base type of ng
type Vector []float64

func NewVector(size int) Vector {
	return make([]float64, size)
}

func (v Vector) Dimensions() int {
	return 1
}

func (v Vector) Size() []int {
	return []int{len(v)}
}

func (v Vector) Dot(vn Vector) float64 {
	sum := 0.0
	vt := vn[:len(v)]
	for i := range v {
		sum += v[i] * vt[i]
	}
	return sum
}

func (v Vector) Magnitude() float64 {
	sum := 0.0
	for i := range v {
		sum += v[i] * v[i]
	}
	return math.Sqrt(sum)
}

func (v Vector) Scale(s float64) {
	for i := range v {
		v[i] *= s
	}
}

func (v Vector) Add(vn Vector) {
	vt := vn[:len(v)]
	for i := range v {
		v[i] += vt[i]
	}
}

func (v Vector) Normalize() {
	v.Scale(1 / v.Magnitude())
}

func (v Vector) Matrix(rows, columns int) Matrix {
	d := v[:rows*columns]
	m := make([][]float64, rows)
	for i := range m {
		m[i], d = d[:columns:columns], d[columns:]
	}
	return m
}

func (v Vector) Fill(value float64) {
	for i := range v {
		v[i] = value
	}
}

func (v *Vector) Resize(size int) {
	if size <= cap(*v) {
		*v = (*v)[:size:size]
	}
	*v = append(*v, make([]float64, size-cap(*v))...)
}

func (v Vector) Sum() float64 {
	sum := 0.0
	for _, value := range v {
		sum += value
	}
	return sum
}
