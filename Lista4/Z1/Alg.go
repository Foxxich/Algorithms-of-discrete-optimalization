package main

import (
	"fmt"
	"math"
)

type AlgHolder struct {
	graph *Graph
}

func newAlg(graph *Graph) *AlgHolder {
	return &AlgHolder{graph: graph}
}

func (ah *AlgHolder) qBFS(graph *Graph, s int, t int, parent *[]int) bool {
	N := graph.N
	visited := make([]bool, N)

	for i := range visited {
		visited[i] = false
	}

	queue := make([]int, 0)

	queue = append(queue, s)
	visited[s] = true
	(*parent)[s] = -1

	for len(queue) != 0 {
		vI := queue[0]
		queue = queue[1:]

		V := graph.nodes[vI]

		for _, u := range V.neighbours {
			if !visited[u] && V.Node_findNeighCost(u) > 0 {
				if u == t {
					(*parent)[u] = vI
					return true
				}
				queue = append(queue, u)
				(*parent)[u] = vI
				visited[u] = true
			}
		}

	}

	return false
}

func (ah *AlgHolder) Alg_alg(s, t int) int {
	var v, u int
	graph := ah.graph
	parent := make([]int, graph.N)
	max_flow := 0
	for ah.qBFS(graph, s, t, &parent) {
		path_flow := math.MaxInt

		for u = t; u != s; u = parent[u] {
			v = parent[u]
			path_flow = min(path_flow, graph.nodes[v].Node_findNeighCost(u))
		}

		for u = t; u != s; u = parent[u] {
			v = parent[u]
			graph.nodes[v].Node_setNeighCost(u, -path_flow)
			graph.nodes[u].Node_setNeighCost(v, path_flow)
			fmt.Println("U: ", u)
			fmt.Println("V: ", v)
			fmt.Println("Flow: ", path_flow)
			fmt.Println("################################")
		}

		max_flow += path_flow
	}

	return max_flow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
