// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import "testing"
import "fmt"

func TestDormandPrince(t *testing.T) {

	t0 := 0.0
	tf := 1.0
	y0 := 0.0

	f := func(t float64, y float64) (dydt float64) {
		return (y - 1) * (y - 1) * (t - 1) * (t - 1)
	}

	opts := Options{
		AbsoluteErrorTolerance: 1e-5,
	}

	yf := opts.DormandPrince(f, t0, tf, y0)

	ye := (tf*tf*tf - 3*tf*tf + 3*tf) / (tf*tf*tf - 3*tf*tf + 3*tf + 3)

	fmt.Println(yf, ye, yf-ye)

}
