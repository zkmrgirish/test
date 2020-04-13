// Package geometry provides elementary geometric entites
// and basic algorithms related to 2d computational geometry
package geometry

// Point2d with x, y coordinates
type Point2d struct {
	x, y float64
}

// NewPoint2d with x,y coordinates
func NewPoint2d(x, y float64) Point2d {
	return Point2d{
		x: x,
		y: y,
	}
}

// X coordinate of the point2d
func (p Point2d) X() float64 {
	return p.x
}

// Y coordinate of the point2d
func (p Point2d) Y() float64 {
	return p.y
}
