// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import "github.com/deslaughter/ng"
import "math"

type Options struct {
	AbsoluteErrorTolerance float64
	RelativeErrorTolerance float64
	MaxStep                float64
	InitialStep            float64
}

// Func is the ordinary differential equation evaluation function.
type Func func(t float64, y float64) (dydt float64)

// FuncSystem is the ordinary differential equation evaluation function
// for a system of equations
type FuncSystem func(t float64, y ng.Vector) (dydt ng.Vector)

// Dormand-Prince Butcher Tableau data
var (
	dpA = [7][7]float64{
		{},
		{1.},
		{1. / 4, 3. / 4},
		{11. / 9, -14. / 3, 40. / 9},
		{4843. / 1458, -3170. / 243, 8056. / 729, -53. / 162},
		{9017. / 3168, -355. / 33, 46732. / 5247, 49. / 176, -5103. / 18656},
		{35. / 384, 0, 500. / 1113, 125. / 192, -2187. / 6784, 11. / 84},
	}

	dpc = [7]float64{0, 1. / 5, 3. / 10, 4. / 5, 8. / 9, 1, 1}

	dpby = [7]float64{5179. / 57600, 0, 7571. / 16695, 393. / 640,
		-92097. / 339200, 187. / 2100, 1. / 40}

	dpbz = [7]float64{35. / 384, 0, 500. / 1113, 125. / 192, -2187. / 6784,
		11. / 84, 0}
)

// DormandPrince uses the Dormand-Prince algorithm to calculate the value of y
// at the final time tf given the ordinary different equation described by
// ODEFunc, the initial time t0, and the initial state y0.
func (opt Options) DormandPrince(f Func, t0, tf float64, y0 float64) float64 {

	k := [7]float64{}

	t := t0
	y := y0
	h := tf - t0
	eAbs := opt.AbsoluteErrorTolerance

	tmp := 0.0

	for t < tf {

		// Loop over rows in A matrix
		for i, rowA := range dpA {

			// Calculate y for this interval
			tmp = 0.0
			for j := range rowA[:i] {
				tmp += k[j] * rowA[j]
			}

			// Calculate slope for these inputs
			k[i] = f(t+h*dpc[i], y+h*dpc[i]*tmp)
		}

		// Calculate next output
		tmp = 0.0
		for i, v := range dpby {
			tmp += k[i] * v
		}
		yTmp := y + h*tmp

		// Calculate next output
		tmp = 0.0
		for i, v := range dpbz {
			tmp += k[i] * v
		}
		zTmp := y + h*tmp

		s := math.Pow(h*eAbs/(2*(tf-t0)*math.Abs(yTmp-zTmp)), 0.25)

		switch {
		case s < 1:
			h /= 2
		case s >= 2:
			t = t + h
			y = zTmp
			h *= 2
		case t+h > tf:
			t = t + h
			y = zTmp
			h = tf - h
		default:
			t = t + h
			y = zTmp
		}
	}

	return y
}
