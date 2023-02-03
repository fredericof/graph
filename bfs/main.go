package main

import "fmt"

type Queue struct {
	Elements []*Vertex
	Size     int
}

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

	e1 := NewEdge("E1", va, vb)
	e2 := NewEdge("E2", vb, vc)
	e3 := NewEdge("E3", vc, vd)
	e4 := NewEdge("E4", vb, vd)
	e5 := NewEdge("E5", va, ve)
	e6 := NewEdge("E6", vf, vd)
	e7 := NewEdge("E7", ve, vf)

	vg := NewVertex("G")
	vh := NewVertex("H")

	e8 := NewEdge("E8", vg, vh)

	vertexList := []*Vertex{va, vb, vc, vd, ve, vf, vg}
	edgeList := []*Edge{e1, e2, e3, e4, e5, e6, e7, e8}

	// Connected Components
	numCC := uint16(0)

	for _, vertex := range vertexList {
		if !vertex.Explored {
			// Init vertex like a root
			vertex.Explored = true
			vertex.Length = 0

			// Increment Connected Component
			numCC++

			// Call bfs
			bfs(vertexList, edgeList, vertex, numCC)

			fmt.Println("-------------------")
		}
	}
}

func bfs(vertexList []*Vertex, edgeList []*Edge, rootVertex *Vertex, numCC uint16) {
	// Initilize Queue with Root Vertex
	queue := Queue{Size: len(vertexList)}
	queue.Enqueue(rootVertex)

	fmt.Println(fmt.Sprint(rootVertex.Value, " - ", rootVertex.Length, " | numCC: ", rootVertex.NumCC))

	for len(queue.Elements) != 0 {
		vertexV := queue.Dequeue()

		// Edge has two vertices (V,W) same (W,V) for undirected graphs
		for _, vertexW := range findEdgesByVertex(vertexV, edgeList) {
			if !vertexW.Explored {
				vertexW.Explored = true
				vertexW.Length = vertexV.Length + 1
				vertexW.NumCC = numCC
				queue.Enqueue(vertexW)
				fmt.Println(fmt.Sprint(vertexW.Value, " - ", vertexW.Length, " | numCC: ", vertexW.NumCC))
			}
		}
	}
}

func findEdgesByVertex(vertex *Vertex, edgeList []*Edge) []*Vertex {
	list := []*Vertex{}

	for _, edge := range edgeList {
		if vertex == edge.VertexA || vertex == edge.VertexB {
			list = append(list, edge.VertexB)
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

func (q *Queue) Enqueue(elem *Vertex) {
	if len(q.Elements) == q.Size {
		fmt.Println("Overflow")
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() *Vertex {
	if len(q.Elements) == 0 {
		fmt.Println("Underflow")
		return nil
	}
	element := q.Elements[0]
	if len(q.Elements) == 1 {
		q.Elements = []*Vertex{}
		return element
	}
	q.Elements = q.Elements[1:]
	return element
}
