// Package polygon provides polygon geometric entity implementation
package polygon

import (
	"github.com/zkmrgirish/geometry"
	"github.com/zkmrgirish/geometry/segment"
)

// Polygon in 2-dimensions
type Polygon struct {
	vertexes []geometry.Point2d
}

// New polygon
func New(vs []geometry.Point2d) Polygon {
	return Polygon{
		vertexes: vs,
	}
}

// Vertexes of the polygon
func (p Polygon) Vertexes() []geometry.Point2d {
	return p.vertexes
}

// Edges of the polygon in order
func (p Polygon) Edges() []segment.Segment2d {
	vertexes := p.Vertexes()
	numVertexes := len(vertexes)

	edges := make([]segment.Segment2d, numVertexes)
	for i := 0; i < numVertexes; i++ {
		edges[i] = segment.New2d(vertexes[i], vertexes[(i+1)%numVertexes])
	}
	return edges
}

// Intersects check if segment s intersects the polygon p
func (p Polygon) Intersects(s segment.Segment2d) bool {
	for _, edge := range p.Edges() {
		if s.Intersects(edge) {
			return true
		}
	}
	return false
}
