package geometry

import "testing"

// TestLength func for a Segment2d
func TestLength(t *testing.T) {
	segments := []Segment2d{
		NewSegment2d(NewPoint2d(0, 0), NewPoint2d(3, 4)),
		NewSegment2d(NewPoint2d(0, 0), NewPoint2d(4, 3)),
		NewSegment2d(NewPoint2d(0, 1), NewPoint2d(4, 1)),
		NewSegment2d(NewPoint2d(0, 1), NewPoint2d(8, 7)),
	}

	results := []float64{5.0, 5.0, 4.0, 10.0}

	for i, seg := range segments {
		length := seg.Length()
		if length != results[i] {
			t.Errorf("testcase [%v]: Length failed, expected %v, got %v",
				i+1, results[i], length)
		}
	}

}

// TestMidPoint func for a Segment2d
func TestMidPoint(t *testing.T) {
	segments := []Segment2d{
		NewSegment2d(NewPoint2d(0, 0), NewPoint2d(3, 4)),
		NewSegment2d(NewPoint2d(0, 0), NewPoint2d(4, 3)),
		NewSegment2d(NewPoint2d(0, 1), NewPoint2d(4, 1)),
		NewSegment2d(NewPoint2d(0, 1), NewPoint2d(8, 7)),
	}

	midpoints := []Point2d{
		NewPoint2d(1.5, 2),
		NewPoint2d(2, 1.5),
		NewPoint2d(2, 1),
		NewPoint2d(4, 4),
	}

	for i, seg := range segments {
		midpoint := seg.MidPoint()
		if midpoint != midpoints[i] {
			t.Errorf("testcase [%v]: MidPoint failed, expected %v, got %v",
				i+1, midpoints[i], midpoint)
		}
	}
}

// TestIntersects of two Segment2d
func TestIntersects(t *testing.T) {
	target := NewSegment2d(NewPoint2d(0, 0), NewPoint2d(1, 1))
	segments := []Segment2d{
		NewSegment2d(NewPoint2d(0, 1), NewPoint2d(1, 0)),
		NewSegment2d(NewPoint2d(0.5, 0.5), NewPoint2d(2, 1)),
		NewSegment2d(NewPoint2d(0, 0.75), NewPoint2d(2, 3.75)),
		NewSegment2d(NewPoint2d(1.5, 1.5), NewPoint2d(3, 2)),
		NewSegment2d(NewPoint2d(0.5, 0.5), NewPoint2d(1.5, 1.5)),
		NewSegment2d(NewPoint2d(-1, -1), NewPoint2d(0.5, 0.5)),
		NewSegment2d(NewPoint2d(1.5, 1.5), NewPoint2d(2, 2)),
	}

	results := []bool{true, true, false, false, true, true, false}

	for i, seg := range segments {
		output := target.Intersects(seg)
		result := results[i]
		if output != result {
			t.Errorf("testcase [%v]: Intersects failed, got %v, expected %v",
				i+1, output, result)
		}
	}

}
