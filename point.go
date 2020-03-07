package geometry

// Point2D is a point in 2-deimension
type Point2D struct {
	x, y int
}

func (p Point2D) X() int {
	return p.x
}

func (p Point2D) Y() int {
	return p.y
}
