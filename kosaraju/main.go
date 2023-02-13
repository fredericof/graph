package main

import (
	"fmt"
	"sort"
)

var CurLabel int
var NumScc int

type Graph struct {
	V []*Vertex
	E []*Edge
}

type Vertex struct {
	Name     string
	Explored bool
	Position int
	SCC      int
}

type Edge struct {
	Name    string
	VertexA *Vertex
	VertexB *Vertex
}

func main() {

	v1 := NewVertex("1")
	v2 := NewVertex("2")
	v3 := NewVertex("3")
	v4 := NewVertex("4")
	v5 := NewVertex("5")
	v6 := NewVertex("6")
	v7 := NewVertex("7")
	v8 := NewVertex("8")
	v9 := NewVertex("9")
	v10 := NewVertex("10")
	v11 := NewVertex("11")

	// SCC 1
	e1 := NewEdge("E1", v1, v3)
	e2 := NewEdge("E2", v3, v5)
	e3 := NewEdge("E3", v5, v1)
	e4 := NewEdge("E4", v3, v11)
	e5 := NewEdge("E5", v5, v7)
	e6 := NewEdge("E6", v5, v9)

	// SCC2
	e7 := NewEdge("E7", v11, v8)
	e8 := NewEdge("E8", v11, v6)

	// SCC3
	e9 := NewEdge("E9", v9, v8)
	e10 := NewEdge("E10", v9, v2)
	e11 := NewEdge("E11", v9, v4)
	e12 := NewEdge("E12", v7, v9)
	e13 := NewEdge("E13", v4, v7)
	e14 := NewEdge("E14", v2, v4)
	e15 := NewEdge("E15", v2, v10)

	// SCC4
	e16 := NewEdge("E16", v8, v6)
	e17 := NewEdge("E17", v6, v10)
	e18 := NewEdge("E18", v10, v8)

	vertexList := []*Vertex{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11}
	edgeList := []*Edge{e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16, e17, e18}
	graph := NewGraph(vertexList, edgeList)

	kosarajo(graph)

	for _, vertex := range graph.V {
		fmt.Printf("%s - Position: %d - SSC: %d \n", vertex.Name, vertex.Position, vertex.SCC)
	}
}

func NewGraph(vertexList []*Vertex, edgeList []*Edge) *Graph {
	return &Graph{
		V: vertexList,
		E: edgeList,
	}
}

func NewEdge(name string, vertexA *Vertex, vertexB *Vertex) *Edge {
	return &Edge{
		Name:    name,
		VertexA: vertexA,
		VertexB: vertexB,
	}
}

func NewVertex(Name string) *Vertex {
	return &Vertex{
		Name:     Name,
		Explored: false,
		Position: 0,
		SCC:      0,
	}
}

func (g *Graph) makeAllUnexplored() {
	for i := 0; i < len(g.V); i++ {
		g.V[i].Explored = false
	}
}

func kosarajo(graph *Graph) {
	reverseGraph(graph)
	graph.makeAllUnexplored()

	TopoSort(graph)

	graph.makeAllUnexplored()

	NumScc = 0

	// Vertex list Sorted by increasing order
	sort.Slice(graph.V[:], func(i, j int) bool {
		return graph.V[i].Position < graph.V[j].Position
	})

	for _, vertex := range graph.V {
		if !vertex.Explored {
			NumScc = NumScc + 1
			dfsScc(graph, vertex)
		}
	}
}

func dfsScc(graph *Graph, sourceVertex *Vertex) {
	sourceVertex.Explored = true
	sourceVertex.SCC = NumScc

	for _, edge := range incomingEdgesFromVertex(graph.E, sourceVertex) {
		if !edge.VertexA.Explored {
			dfsScc(graph, edge.VertexA)
		}
	}
}

func reverseGraph(graph *Graph) {
	for _, edge := range graph.E {
		edge.VertexA, edge.VertexB = edge.VertexB, edge.VertexA
	}
}

func TopoSort(graph *Graph) {
	graph.makeAllUnexplored()
	CurLabel = len(graph.V)

	for _, vertex := range graph.V {
		if !vertex.Explored {
			dfsTopo(graph, vertex)
		}
	}
}

func dfsTopo(graph *Graph, sourceVertex *Vertex) {
	sourceVertex.Explored = true

	for _, edge := range outgoingEdgesFromVertex(graph.E, sourceVertex) {
		if !edge.VertexB.Explored {
			dfsTopo(graph, edge.VertexB)
		}
	}

	sourceVertex.Position = CurLabel
	CurLabel = CurLabel - 1
}

func outgoingEdgesFromVertex(edgeList []*Edge, sourceVertex *Vertex) []*Edge {
	list := []*Edge{}

	for _, edge := range edgeList {
		if sourceVertex == edge.VertexA {
			list = append(list, edge)
		}
	}

	return list
}

func incomingEdgesFromVertex(edgeList []*Edge, sourceVertex *Vertex) []*Edge {
	list := []*Edge{}

	for _, edge := range edgeList {
		if sourceVertex == edge.VertexB {
			list = append(list, edge)
		}
	}

	return list
}
