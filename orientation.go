package test

type orientation int

const (
	// orientations of three points in 2d
	CLOCKWISE orientation = iota
	COUNTERCLOCKWISE
	COLINEAR
)

// Orientation of three points in 2d plain
func Orientation(p, q, r Point2d) orientation {
	orient := (q.Y()-p.Y())*(r.X()-q.X()) - (q.X()-p.X())*(r.Y()-q.Y())

	switch {
	case orient == 0:
		return COLINEAR
	case orient > 0:
		return CLOCKWISE
	default:
		return COUNTERCLOCKWISE
	}
}
