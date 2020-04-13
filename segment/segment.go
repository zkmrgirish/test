// Package segment provide segment entity
package segment

import (
	"math"

	"github.com/zkmrgirish/geometry"
	"github.com/zkmrgirish/geometry/util"
)

// Segment2d 2-dimensional entity
type Segment2d struct {
	p1 geometry.Point2d
	p2 geometry.Point2d
}

// New2d gives new 2-dimensional segment
func New2d(p1, p2 geometry.Point2d) Segment2d {
	return Segment2d{
		p1: p1,
		p2: p2,
	}
}

// Length calculates eucledian distance
func (s Segment2d) Length() float64 {
	dx := s.p1.X() - s.p2.X()
	dy := s.p1.Y() - s.p2.Y()
	return math.Sqrt(dx*dx + dy*dy)
}

// MidPoint returns mid point of the segment
func (s Segment2d) MidPoint() geometry.Point2d {
	return geometry.NewPoint2d(
		(s.p1.X()+s.p2.X())/2,
		(s.p1.Y()+s.p2.Y())/2,
	)
}

// Contains checks if point p is on segment s
func (s Segment2d) Contains(p geometry.Point2d) bool {
	onSegment := (p.X() <= math.Max(s.p1.X(), s.p2.X())) &&
		(p.X() >= math.Min(s.p1.X(), s.p2.X())) &&
		(p.Y() <= math.Max(s.p1.Y(), s.p2.Y())) &&
		(p.Y() >= math.Min(s.p1.Y(), s.p2.Y()))

	return onSegment
}

// Intersects check if two segments in 2d intersects or not
// ref: https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
func (s Segment2d) Intersects(t Segment2d) bool {
	o1 := util.Orientation(s.p1, s.p2, t.p1)
	o2 := util.Orientation(s.p1, s.p2, t.p2)
	o3 := util.Orientation(t.p1, t.p2, s.p1)
	o4 := util.Orientation(t.p1, t.p2, s.p2)

	if o1 != o2 && o3 != o4 {
		return true
	}

	if o1 == util.COLINEAR && s.Contains(t.p1) {
		return true
	}

	if o2 == util.COLINEAR && s.Contains(t.p2) {
		return true
	}

	return false
}
