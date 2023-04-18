package datastructures

import (
	"fmt"
)

// Graph is a collection of vertices
type Graph struct {
	vertices []*Vertex
}

// Vertex defines our vertex struct
type Vertex struct {
	value    int
	adjacent []*Vertex
}

// AddVertex adds a new vertex
func (g *Graph) AddVertex(val int) *Vertex {
	v := &Vertex{
		value:    val,
		adjacent: make([]*Vertex, 0),
	}
	g.vertices = append(g.vertices, v)
	return v
}

// AddEdge adds an edge onto a vertex
func (g *Graph) AddEdge(v1, v2 *Vertex) {
	v1.adjacent = append(v1.adjacent, v2)
	v2.adjacent = append(v2.adjacent, v1)
}

// String returns a string representation of our graph
func (g *Graph) String() string {
	str := ""
	for _, v := range g.vertices {
		str += fmt.Sprintf("%d -> ", v.value)
		for _, adj := range v.adjacent {
			str += fmt.Sprintf("%d, ", adj.value)
		}
		str += "\n"
	}
	return str
}
