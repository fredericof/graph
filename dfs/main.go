package main

import "fmt"

type Stack []*Vertex

type Vertex struct {
	Value    string
	Explored bool
	Length   uint16
	NumCC    uint16
}

type Edge struct {
	Name    string
	VertexA *Vertex
	VertexB *Vertex
}

func main() {
	va := NewVertex("A")
	vb := NewVertex("B")
	vc := NewVertex("C")
	vd := NewVertex("D")
	ve := NewVertex("E")
	vf := NewVertex("F")
	vg := NewVertex("G")
	vh := NewVertex("H")

	e1 := NewEdge("E1", va, vb)
	e2 := NewEdge("E2", vb, vc)
	e3 := NewEdge("E3", vc, vd)
	e4 := NewEdge("E4", vb, vd)
	e5 := NewEdge("E5", va, ve)
	e6 := NewEdge("E6", vf, vd)
	e7 := NewEdge("E7", ve, vf)
	e8 := NewEdge("E8", vg, vh)
	e9 := NewEdge("E9", vh, vf)

	vertexList := []*Vertex{va, vb, vc, vd, ve, vf, vg, vh}
	edgeList := []*Edge{e1, e2, e3, e4, e5, e6, e7, e8, e9}

	iterativeDfs(va, edgeList)
	//recursiveDfs(va, edgeList)

	fmt.Println("----------------------------")

	for _, vertex := range vertexList {
		fmt.Println(fmt.Sprint(vertex.Value, " - Explored: ", vertex.Explored))
	}
}

func recursiveDfs(vertex *Vertex, edgeList []*Edge) {
	vertex.Explored = true
	for _, vertexW := range findEdgesByVertex(vertex, edgeList) {
		if !vertexW.Explored {
			fmt.Println(fmt.Sprint(vertexW.Value))
			recursiveDfs(vertexW, edgeList)
		}
	}
}

// Iterative Depth-first search
func iterativeDfs(vertex *Vertex, edgeList []*Edge) {
	var stack Stack
	stack.Push(vertex)

	for !stack.IsEmpty() {
		vertexV, _ := stack.Pop()

		if !vertexV.Explored {
			vertexV.Explored = true

			for _, vertexW := range findEdgesByVertex(vertexV, edgeList) {
				stack.Push(vertexW)
				//fmt.Println(fmt.Sprint(vertexW.Value, "-", vertexW.Explored))
			}
		}
	}
}

func findEdgesByVertex(vertex *Vertex, edgeList []*Edge) []*Vertex {
	list := []*Vertex{}

	for _, edge := range edgeList {
		if vertex == edge.VertexA {
			list = append(list, edge.VertexB)
		} else if vertex == edge.VertexB {
			list = append(list, edge.VertexA)
		}
	}

	return list
}

func NewEdge(name string, vertexA *Vertex, vertexB *Vertex) *Edge {
	return &Edge{
		Name:    name,
		VertexA: vertexA,
		VertexB: vertexB,
	}
}

func NewVertex(value string) *Vertex {
	return &Vertex{
		Value:    value,
		Explored: false,
		Length:   ^uint16(0),
		NumCC:    0,
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(vertex *Vertex) {
	*s = append(*s, vertex)
}

func (s *Stack) Pop() (*Vertex, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
