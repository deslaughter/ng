// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ng

import (
	"testing"
)

func TestNewArray(t *testing.T) {

	a := NewArray(nil, 2, 3, 4)

	if act, exp := a.strides[0], 12; act != exp {
		t.Fatalf("a.strides[0] = %v, expected %v", act, exp)
	}
	if act, exp := a.strides[1], 4; act != exp {
		t.Fatalf("a.strides[1] = %v, expected %v", act, exp)
	}
	if act, exp := a.strides[2], 1; act != exp {
		t.Fatalf("a.strides[2] = %v, expected %v", act, exp)
	}
}
