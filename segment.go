package test

import "math"

// Segment2d 2-dimensional entity
type Segment2d struct {
	p1 Point2d
	p2 Point2d
}

// New2d gives new 2-dimensional segment
func NewSegment2d(p1, p2 Point2d) Segment2d {
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
func (s Segment2d) MidPoint() Point2d {
	return NewPoint2d(
		(s.p1.X()+s.p2.X())/2,
		(s.p1.Y()+s.p2.Y())/2,
	)
}

// contains checks if point p is on segment s
// given p lies on same line on which segment s lies
func (s Segment2d) contains(p Point2d) bool {
	onSegment := (p.X() <= math.Max(s.p1.X(), s.p2.X())) &&
		(p.X() >= math.Min(s.p1.X(), s.p2.X())) &&
		(p.Y() <= math.Max(s.p1.Y(), s.p2.Y())) &&
		(p.Y() >= math.Min(s.p1.Y(), s.p2.Y()))

	return onSegment
}

// Intersects check if two segments in 2d intersects or not
// ref: https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
func (s Segment2d) Intersects(t Segment2d) bool {
	o1 := Orientation(s.p1, s.p2, t.p1)
	o2 := Orientation(s.p1, s.p2, t.p2)
	o3 := Orientation(t.p1, t.p2, s.p1)
	o4 := Orientation(t.p1, t.p2, s.p2)

	if o1 != o2 && o3 != o4 {
		return true
	}

	if o1 == COLINEAR && s.contains(t.p1) {
		return true
	}

	if o2 == COLINEAR && s.contains(t.p2) {
		return true
	}

	return false
}
