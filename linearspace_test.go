// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import "fmt"

func ExampleLinearSpace() {

	start, end, count := 1.0, 5.0, 5

	v := LinearSpace(start, end, count)

	fmt.Println(v)
	// Output: [1 2 3 4 5]
}
