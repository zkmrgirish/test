package polygon

import (
	"testing"

	"github.com/zkmrgirish/geometry"
	"github.com/zkmrgirish/geometry/segment"
)

// TestIntersects test polygon and segment intersection
func TestIntersects(t *testing.T) {
	vs := []geometry.Point2d{
		geometry.NewPoint2d(0, 0),
		geometry.NewPoint2d(2, 0),
		geometry.NewPoint2d(2, 1),
		geometry.NewPoint2d(1, 2),
		geometry.NewPoint2d(0, 1),
	}
	p := New(vs)
	segments := []segment.Segment2d{
		segment.New2d(geometry.NewPoint2d(-1, 0), geometry.NewPoint2d(-0.5, 0.5)),
		segment.New2d(geometry.NewPoint2d(-1, 0), geometry.NewPoint2d(0.5, 1.5)),
		segment.New2d(geometry.NewPoint2d(-1, 0), geometry.NewPoint2d(2, 3)),
		segment.New2d(geometry.NewPoint2d(-1, 0), geometry.NewPoint2d(3, 1.5)),
		segment.New2d(geometry.NewPoint2d(0.5, 0.4), geometry.NewPoint2d(1.5, 0.6)),
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
