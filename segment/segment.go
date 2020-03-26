package segment

import (
	"math"

	"github.com/zkmrgirish/geometry"
)

// Segment2D 2-dimensional entity
type Segment2D struct {
	p1 geometry.Point2D
	p2 geometry.Point2D
}

func New(p1, p2 geometry.Point2D) Segment2D {
	return Segment2D{
		p1: p1,
		p2: p2,
	}
}

// Length calculates eucledian distance
func (s Segment2D) Length() float64 {
	dx := s.p1.X() - s.p2.X()
	dy := s.p1.Y() - s.p2.Y()
	return math.Sqrt(dx*dx + dy*dy)
}

// MidPoint returns mid point of the segment
func (s Segment2D) MidPoint() geometry.Point2D {
	return geometry.NewPoint2D(
		(s.p1.X()+s.p2.X())/2,
		(s.p1.Y()+s.p2.Y())/2,
	)
}
