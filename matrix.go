// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

type Matrix [][]float64

func (m Matrix) Dimensions() int {
	return 2
}

func (m Matrix) Size() []int {
	return []int{len(m), len(m[0])}
}

func (m Matrix) Multiply(m1 Matrix) {

}
