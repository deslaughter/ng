// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

// Array is the base type of ng
type Array struct {
	data    []float64
	size    []int
	strides []int
	V       Vector
	M       Matrix
}

// NewArray returns an with the given number of elements.
func NewArray(data []float64, sizes ...int) Array {

	strides := make([]int, len(sizes))
	for i := range strides {
		strides[i] = 1
	}

	nelem := 1
	for i, s := range sizes {
		nelem *= s
		for _, v := range sizes[i+1:] {
			strides[i] *= v
		}
	}

	if data != nil {
		data = data[:nelem:nelem]
	} else {
		data = make([]float64, nelem)
	}

	a := Array{
		data:    data,
		size:    sizes,
		strides: strides,
	}

	switch len(sizes) {

	// 1-D Vector
	case 1:
		a.V = a.data

	// 2-D Matrix
	case 2:
		d := a.data
		m := make(Matrix, sizes[0])
		for i := range m {
			m[i], d = d[:sizes[1]:sizes[1]], d[sizes[1]:]
		}
	}

	return a
}

// Order returns the number of dimensions of the array.
func (a Array) Order() int {
	return len(a.size)
}

// At returns the array value at the corresponding indices. If the number of
// indices does not match the array order or they are outside the bounds,
// the function will panic.
func (a Array) At(indices ...int) float64 {
	index := 0
	for j, i := range indices {
		index += a.strides[j] * i
	}
	return a.data[index]
}

// Transpose performs an in-place matrix transpose of the elements in the
// Array. The resulting matrix will have the dimensions [columns x rows].
func (a Array) Transpose() {

	if len(a.size) != 2 {
		panic("Array must be 2-D")
	}

	rows, columns := a.size[0], a.size[1]

	q := rows*columns - 1

	for start := 1; start < q; start++ {

		next := start
		i := 0

		for {
			i++
			next = (next%rows)*columns + next/rows
			if next <= start {
				break
			}
		}

		if next >= start && i != 1 {
			t := a.data[start]
			next = start
			for {
				i = (next%rows)*columns + next/rows
				if i != start {
					a.data[next] = a.data[i]
				} else {
					a.data[next] = t
				}
				next = i
				if next <= start {
					break
				}
			}
		}
	}
}
