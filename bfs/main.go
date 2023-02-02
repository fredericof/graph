package main

import "fmt"

type Queue struct {
	Elements []*Vertex
	Size     int
}

type Vertex struct {
	Value    string
	Explored bool
	IsRoot   bool
}

type Edge struct {
	Name    string
	VertexA *Vertex
	VertexB *Vertex
}

func main() {
	vaRoot := NewVertex("A", true)
	vb := NewVertex("B", false)
	vc := NewVertex("C", false)
	vd := NewVertex("D", false)
	ve := NewVertex("E", false)
	vf := NewVertex("F", false)

	e1 := NewEdge("E1", vaRoot, vb)
	e2 := NewEdge("E2", vb, vc)
	e3 := NewEdge("E3", vc, vd)
	e4 := NewEdge("E4", vb, vd)
	e5 := NewEdge("E5", vaRoot, ve)
	e6 := NewEdge("E6", ve, vd)
	e7 := NewEdge("E7", ve, vf)

	vertexList := []*Vertex{vaRoot, vb, vc, vd, ve, vf}
	edgeList := []*Edge{e1, e2, e3, e4, e5, e6, e7}

	// Mark root vertex as explored and all others as unexplored
	initOnlyRootVertexAsExplored(vertexList)

	// Initilize Queue with Root Vertex
	queue := Queue{Size: len(vertexList)}
	queue.Enqueue(vaRoot)

	for len(queue.Elements) != 0 {
		for _, vertex := range findEdgesByVertex(queue.Dequeue(), edgeList) {
			if !vertex.Explored {
				vertex.Explored = true
				queue.Enqueue(vertex)
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

func initOnlyRootVertexAsExplored(vertexList []*Vertex) {
	for i := 0; i < len(vertexList); i++ {
		if vertexList[i].IsRoot {
			vertexList[i].Explored = true
			return
		}
	}
}

func NewEdge(name string, vertexA *Vertex, vertexB *Vertex) *Edge {
	return &Edge{
		Name:    name,
		VertexA: vertexA,
		VertexB: vertexB,
	}
}

func NewVertex(value string, isRoot bool) *Vertex {
	return &Vertex{
		Value:  value,
		IsRoot: isRoot,
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
