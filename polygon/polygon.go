package polygon

import (
	"github.com/zkmrgirish/geometry"
	"github.com/zkmrgirish/geometry/segment"
)

// Polygon in 2-dimensions
type Polygon struct {
	vertexes []geometry.Point2D
}

// get all the vertexes of the polygon in order
func (p Polygon) Vertexes() []geometry.Point2D {
	return p.vertexes
}

// get edges of the polygon in order
func (p Polygon) Edges() []segment.Segment2D {
	vertexes := p.vertexes
	no_vertexes := len(vertexes)

	edges := make([]segment.Segment2D, no_vertexes)
	for i := 0; i < no_vertexes; i++ {
		edges[i] = segment.New(vertexes[i], vertexes[(i+1)%no_vertexes])
	}
	return edges
}
