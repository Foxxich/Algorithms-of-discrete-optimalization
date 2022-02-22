package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

type Node struct {
	name          int
	Index         int
	neighbours    []int
	neighboursCap []int
	flow          []int
	visited       bool
	pred          int
	pos           int
}

type Graph struct {
	N      int
	M      int
	K      int
	C      int
	nodes  []*Node
	levels []int
}

func newNode(name int) *Node {
	return &Node{name: name, neighbours: make([]int, 0), neighboursCap: make([]int, 0), flow: make([]int, 0), visited: false}
}

func Graph_newGraph(K int) *Graph {
	ncount := int(math.Pow(float64(2), float64(K)))
	mcount := K * ncount / 2

	graph := &Graph{N: ncount, M: mcount, K: K, nodes: make([]*Node, 0), levels: make([]int, ncount)}

	for i := 0; i < ncount; i++ {
		graph.addNode(i)
	}
	return graph
}

func (graph *Graph) addNode(name int) {
	graph.nodes = append(graph.nodes, newNode(name))
}

func (graph *Graph) Graph_addEdge(source int, destination int, cost int) {
	graph.nodes[source].neighbours = append(graph.nodes[source].neighbours, destination)
	graph.nodes[source].neighboursCap = append(graph.nodes[source].neighboursCap, cost)
	graph.nodes[source].flow = append(graph.nodes[source].flow, 0)
}

func (graph *Graph) Graph_generateEdges() {
	N := graph.N
	K := graph.K
	powers := make([]int, 1)
	powers[0] = 1

	for i := 1; i < K; i++ {
		powers = append(powers, powers[i-1]*2)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {

			if i+powers[j] <= N && H(i)+1 == H(i+powers[j]) {
				destination := i + powers[j]
				cost := getCapacity(i, i+powers[j], K)
				graph.Graph_addEdge(i, destination, cost)
				fmt.Println("SRC: ", i)
				fmt.Println("DEST: ", destination)
				fmt.Println("COST: ", cost)
				fmt.Println("################################")
			}
		}
	}
}

func H(i int) int {
	counter := 0
	for i != 0 {
		if i%2 == 1 {
			counter++
		}
		i = i >> 1
	}
	return counter
}

func Z(i int, K int) int {
	return K - H(i)
}

func getCapacity(a, b, K int) int {
	rand.Seed(time.Now().UnixNano())
	limit := int(math.Pow(float64(2), float64(max(max(H(a), Z(a, K)), max(H(b), Z(b, K))))))
	return rand.Intn(limit) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (graph *Graph) Graph_makeCopy() *Graph {
	copy := Graph_newGraph(graph.K)
	copy.nodes = make([]*Node, graph.N)

	for i, n := range graph.nodes {

		copy.nodes[i] = n.copyNode()
	}

	return copy

}

func (node *Node) copyNode() *Node {
	copy := newNode(node.name)

	copy.flow = node.flow
	copy.neighbours = append(copy.neighbours, node.neighbours...)
	copy.neighboursCap = append(copy.neighboursCap, node.neighboursCap...)
	copy.pos = node.pos
	copy.pred = node.pred
	copy.visited = node.visited
	copy.Index = node.Index
	return copy
}

func (node *Node) Node_findNeighCost(u int) int {
	for i, n := range node.neighbours {
		if n == u {
			return node.neighboursCap[i]
		}
	}

	return 0
}

func (node *Node) Node_setNeighCost(u int, diff int) {
	for i, n := range node.neighbours {
		if n == u {
			node.neighboursCap[i] += diff
		}
	}
}

func (node *Node) Node_setNeighFlow(u int, diff int) {
	for i, n := range node.neighbours {
		if n == u {
			node.flow[i] += diff
		}
	}
}

func (graph *Graph) Graph_printGraphGLPK(resultsString string) {

	f, _ := os.OpenFile(resultsString, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString("using JuMP\nusing GLPK\nusing DelimitedFiles\nimport LinearAlgebra\n")
	f.WriteString("G = [\n")
	N := graph.N
	for _, n := range graph.nodes {
		toPrint := make([]int, N)

		for j, k := range n.neighbours {
			toPrint[k] = n.neighboursCap[j]
		}

		fmt.Fprintln(f, "\t", toPrint)
	}
	f.WriteString("]\n")

	f.WriteString("n = size(G)[1]\n")

	f.WriteString("max_flow = Model(GLPK.Optimizer)\n")

	f.WriteString("@variable(max_flow, f[1:n,1:n] >= 0)\n")

	f.WriteString("@constraint(max_flow, [i = 1:n, j = 1:n], f[i, j] <= G[i, j])\n")

	f.WriteString("@constraint(max_flow, [i = 1:n; i != 1 && i != n], sum(f[i, :]) == sum(f[:, i]))\n")

	f.WriteString("@objective(max_flow, Max, sum(f[1, :]))\n")

	f.WriteString("optimize!(max_flow)\n")

	f.WriteString("objective_value(max_flow)\n")

	f.WriteString("value.(f)\n")

	f.Close()
}
