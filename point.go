package geometry

// Point2D is a point in 2-deimension
type Point2D struct {
	x, y float64
}

func NewPoint2D(x, y float64) Point2D {
	return Point2D{
		x: x,
		y: y,
	}
}

func (p Point2D) X() float64 {
	return p.x
}

func (p Point2D) Y() float64 {
	return p.y
}
