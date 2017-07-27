// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

// LinearSpace returns a vector of length count with values evenly distributed
// from start to end.
func LinearSpace(start, end float64, count int) Vector {

	// Create vector
	v := make(Vector, count)

	// Calculate range
	r := end - start

	// Set values in range
	for i := range v {
		v[i] = r*float64(i)/float64(count-1) + start
	}

	return v
}
