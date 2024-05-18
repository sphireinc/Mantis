package datastructures

import (
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := &Graph{}

	// Test adding a vertex
	v := g.AddVertex(1)
	if len(g.vertices) != 1 || g.vertices[0] != v {
		t.Errorf("AddVertex failed, expected 1 vertex, got %d", len(g.vertices))
	}
	if v.value != 1 {
		t.Errorf("AddVertex failed, expected vertex value 1, got %d", v.value)
	}
}

func TestAddEdge(t *testing.T) {
	g := &Graph{}
	v1 := g.AddVertex(1)
	v2 := g.AddVertex(2)

	// Test adding an edge
	g.AddEdge(v1, v2)
	if len(v1.adjacent) != 1 || v1.adjacent[0] != v2 {
		t.Errorf("AddEdge failed, v1 should have 1 adjacent, got %d", len(v1.adjacent))
	}
	if len(v2.adjacent) != 1 || v2.adjacent[0] != v1 {
		t.Errorf("AddEdge failed, v2 should have 1 adjacent, got %d", len(v2.adjacent))
	}
}

func TestGraphString(t *testing.T) {
	g := &Graph{}
	v1 := g.AddVertex(1)
	v2 := g.AddVertex(2)
	v3 := g.AddVertex(3)
	g.AddEdge(v1, v2)
	g.AddEdge(v1, v3)

	// Expected format: "1 -> 2, 3, \n2 -> 1, \n3 -> 1, \n"
	expected := "1 -> 2, 3, \n2 -> 1, \n3 -> 1, \n"
	output := g.String()
	if output != expected {
		t.Errorf("Graph.String() failed, expected %q, got %q", expected, output)
	}
}
