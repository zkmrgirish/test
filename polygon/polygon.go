package polygon

import (
	"github.com/zkmrgirish/geometry"
	"github.com/zkmrgirish/geometry/segment"
)

// Polygon2d in 2-dimensions
type Polygon2d struct {
	vertexes []geometry.Point2d
}

// Vertexes of the polygon
func (p Polygon2d) Vertexes() []geometry.Point2d {
	return p.vertexes
}

// Edges of the polygon in order
func (p Polygon2d) Edges() []segment.Segment2d {
	vertexes := p.vertexes
	numVertexes := len(vertexes)

	edges := make([]segment.Segment2d, numVertexes)
	for i := 0; i < numVertexes; i++ {
		edges[i] = segment.New2d(vertexes[i], vertexes[(i+1)%numVertexes])
	}
	return edges
}
