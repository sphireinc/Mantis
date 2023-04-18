package datastructures

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	value    int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(val int) *Vertex {
	v := &Vertex{
		value:    val,
		adjacent: make([]*Vertex, 0),
	}
	g.vertices = append(g.vertices, v)
	return v
}

func (g *Graph) AddEdge(v1, v2 *Vertex) {
	v1.adjacent = append(v1.adjacent, v2)
	v2.adjacent = append(v2.adjacent, v1)
}

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
