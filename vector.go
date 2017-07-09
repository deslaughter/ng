// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import (
	"fmt"
	"math"
)

type Vector []float64

func (v Vector) Dimensions() int {
	return 1
}

func (v Vector) Size() []int {
	return []int{len(v)}
}

func (v1 Vector) Dot(v2 Vector) float64 {
	sum := 0.0
	for i := range v1 {
		sum += v1[i] * v2[i]
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
	if len(v) != len(vn) {
		panic(fmt.Errorf("Vector lengths are not equal"))
	}
	for i := range v {
		v[i] += vn[i]
	}
}

func (v Vector) Normalize() {
	v.Scale(1 / v.Magnitude())
}
