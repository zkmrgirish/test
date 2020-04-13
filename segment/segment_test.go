package segment

import (
	"testing"

	"github.com/zkmrgirish/geometry"
)

// TestLength func for a Segment2d
func TestLength(t *testing.T) {
	segments := []Segment2d{
		New2d(geometry.NewPoint2d(0, 0), geometry.NewPoint2d(3, 4)),
		New2d(geometry.NewPoint2d(0, 0), geometry.NewPoint2d(4, 3)),
		New2d(geometry.NewPoint2d(0, 1), geometry.NewPoint2d(4, 1)),
		New2d(geometry.NewPoint2d(0, 1), geometry.NewPoint2d(8, 7)),
	}

	results := []float64{5.0, 5.0, 4.0, 10.0}

	for i, seg := range segments {
		length := seg.Length()
		if length != results[i] {
			t.Errorf("testcase [%v]: Length failed, expected %v, got %v",
				i, results[i], length)
		}
	}

}

// TestMidPoint func for a Segment2d
func TestMidPoint(t *testing.T) {
	segments := []Segment2d{
		New2d(geometry.NewPoint2d(0, 0), geometry.NewPoint2d(3, 4)),
		New2d(geometry.NewPoint2d(0, 0), geometry.NewPoint2d(4, 3)),
		New2d(geometry.NewPoint2d(0, 1), geometry.NewPoint2d(4, 1)),
		New2d(geometry.NewPoint2d(0, 1), geometry.NewPoint2d(8, 7)),
	}

	midpoints := []geometry.Point2d{
		geometry.NewPoint2d(1.5, 2),
		geometry.NewPoint2d(2, 1.5),
		geometry.NewPoint2d(2, 1),
		geometry.NewPoint2d(4, 4),
	}

	for i, seg := range segments {
		midpoint := seg.MidPoint()
		if midpoint != midpoints[i] {
			t.Errorf("testcase [%v]: MidPoint failed, expected %v, got %v",
				i, midpoints[i], midpoint)
		}
	}
}

// TestIntersection of two Segment2d
func TestIntersection(t *testing.T) {
	// TODO: test intersection
}
