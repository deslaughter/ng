// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import "fmt"

func ExampleMatrix_Multiply_first() {

	// Create vectors
	va := Vector{1, 2, 3}
	vb := Vector{4, 5, 6}

	// Create matrix views of vectors
	mA := va.Matrix(len(va), 1) // [3 x 1]
	mB := vb.Matrix(1, len(vb)) // [1 x 3]

	// Multiply matrices and receive vector representing matrix
	// with the given rows and columns
	vc, rows, columns := mA.Multiply(mB)

	// Create matrix view of result vector
	mC := vc.Matrix(rows, columns)

	fmt.Println(mA)
	fmt.Println(mB)
	fmt.Println(mC)
	// Output:
	// [[1] [2] [3]]
	// [[4 5 6]]
	// [[4 5 6] [8 10 12] [12 15 18]]
}

func ExampleMatrix_Multiply_second() {

	va := Vector{1, 2, 3}
	vb := Vector{4, 5, 6}

	mA := va.Matrix(1, len(va))
	mB := vb.Matrix(len(vb), 1)

	vc, rows, columns := mA.Multiply(mB)

	mC := vc.Matrix(rows, columns)

	fmt.Println(mC)
	// Output:
	// [[32]]
}
