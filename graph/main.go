package main

import "fmt"

// a graph doesn't have a specific root, it has vertices and edges that connect the vertives

// type Graph - adjacency list
type Graph struct {
	vertices []*Vertex
}

// type Vertex
type Vertex struct {
	value         int
	adjacencyList []*Vertex
}

// Add Vertex
func (g *Graph) AddVertex(val int) {
	if containVertex(g.vertices, val) {
		fmt.Printf("vertex %v alreay exists, not added\n", val)
	} else {
		g.vertices = append(g.vertices, &Vertex{value: val})
	}
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
	if g.edgeExists(from, to) {
		fromVertexPtr := g.getVertex(from)
		toVertexPtr := g.getVertex(to)
		fromVertexPtr.adjacencyList = append(fromVertexPtr.adjacencyList, toVertexPtr)
	}
}

// Check if edge exists
func (g *Graph) edgeExists(from, to int) bool {
	fromVertexPtr := g.getVertex(from)
	toVertexPtr := g.getVertex(to)

	if fromVertexPtr != nil || toVertexPtr != nil {
		if containVertex(fromVertexPtr.adjacencyList, to) {
			fmt.Printf("edge %v -> %v already exists\n", from, to)
		} else {
			return true
		}
	} else {
		fmt.Printf("Invalid edge %v -> %v\n", from, to)
	}
	return false
}

// getVertex returns a pointer to the vertex with the value to be searched
func (g *Graph) getVertex(val int) *Vertex {
	for i, v := range g.vertices {
		if val == v.value {
			return g.vertices[i]
		}
	}
	return nil
}

// contain vertex
func containVertex(vertexList []*Vertex, toCheck int) bool {
	for _, v := range vertexList {
		if toCheck == v.value {
			return true
		}
	}
	return false
}

// Delete Vertex
// Delete Edge

// Print Graph
func (g *Graph) printGraph() {
	for _, v := range g.vertices {
		fmt.Printf("vertex %v:", v.value)
		for _, v := range v.adjacencyList {
			fmt.Printf(" %v ", v.value)
		}
		fmt.Println()
	}
}

// main to test
func main() {
	testGraph := &Graph{}

	for i := 0; i < 5; i++ {
		testGraph.AddVertex(i)
	}

	// test adding duplicate vertex
	testGraph.AddVertex(5)
	testGraph.AddVertex(5)

	// test add edge
	testGraph.AddEdge(1, 2)
	testGraph.printGraph()

	// test add invalid edge
	testGraph.AddEdge(6, 8)
	testGraph.AddEdge(1, 2)
}
