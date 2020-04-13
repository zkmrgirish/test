package util

import (
	"testing"

	"github.com/zkmrgirish/geometry"
)

// TestOrientation for different points in 2d
// tests all three cases CLOCKWISE, COUNTERCLOCKWISE, COLINEAR
func TestOrientation(t *testing.T) {
	var p, q, r geometry.Point2d

	// test for COLINEAR
	p, q, r = geometry.NewPoint2d(0, 0), geometry.NewPoint2d(1, 1), geometry.NewPoint2d(2, 2)
	if orient := Orientation(p, q, r); orient != COLINEAR {
		t.Errorf("testcase [colinear]: Orientation failed, got %v", orient)
	}

	// test for CLOCKWISE
	p, q, r = geometry.NewPoint2d(0, 0), geometry.NewPoint2d(1, 1), geometry.NewPoint2d(2, 0)
	if orient := Orientation(p, q, r); orient != CLOCKWISE {
		t.Errorf("testcase [clockwise]: Orientation failed, got %v", orient)
	}

	// test for COUNTERCLOCKWISE
	p, q, r = geometry.NewPoint2d(0, 0), geometry.NewPoint2d(1, 1), geometry.NewPoint2d(0, 2)
	if orient := Orientation(p, q, r); orient != COUNTERCLOCKWISE {
		t.Errorf("testcase [counterclockwise]: Orientation failed, got %v", orient)
	}
}
