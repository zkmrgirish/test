// Package segment provide segment entity
package segment

import (
	"math"

	"github.com/zkmrgirish/geometry"
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

// Intersects check if two segments in 2d intersects or not
func (s Segment2d) Intersects(t Segment2d) bool {
	// TODO: find intersection of segment s and t
	return false
}
