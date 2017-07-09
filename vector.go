// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import (
	"math"
)

type Vector []float64

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
