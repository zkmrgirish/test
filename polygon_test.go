package geometry

import "testing"

// TestPolygonIntersects test polygon and segment intersection
func TestPolygonIntersects(t *testing.T) {
	vs := []Point2d{
		NewPoint2d(0, 0),
		NewPoint2d(2, 0),
		NewPoint2d(2, 1),
		NewPoint2d(1, 2),
		NewPoint2d(0, 1),
	}
	p := New(vs)
	segments := []Segment2d{
		NewSegment2d(NewPoint2d(-1, 0), NewPoint2d(-0.5, 0.5)),
		NewSegment2d(NewPoint2d(-1, 0), NewPoint2d(0.5, 1.5)),
		NewSegment2d(NewPoint2d(-1, 0), NewPoint2d(2, 3)),
		NewSegment2d(NewPoint2d(-1, 0), NewPoint2d(3, 1.5)),
		NewSegment2d(NewPoint2d(0.5, 0.4), NewPoint2d(1.5, 0.6)),
	}

	results := []bool{false, true, true, true, false}
	for i, seg := range segments {
		output := p.Intersects(seg)
		result := results[i]
		if output != result {
			t.Errorf("testcase [%v]: Intersects failed, got %v, expected %v",
				i+1, output, result)
		}
	}
}
