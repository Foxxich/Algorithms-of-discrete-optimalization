package main

import "math"

type Dynic struct {
	graph *Graph
}

func Dynic_newDynic(graph *Graph) *Dynic {
	return &Dynic{graph: graph}
}

func (dynic *Dynic) Dynic_BFS(s, t int) bool {
	graph := dynic.graph
	for i := range graph.levels {
		graph.levels[i] = -1
	}

	graph.levels[s] = 0
	queue := make([]int, 0)

	queue = append(queue, s)

	for len(queue) != 0 {
		vI := queue[0]
		queue = queue[1:]
		V := graph.nodes[vI]

		for i, u := range V.neighbours {
			if graph.levels[u] < 0 && V.flow[i] < V.neighboursCap[i] {
				graph.levels[u] = graph.levels[vI] + 1
				queue = append(queue, u)
			}
		}
	}

	return graph.levels[t] >= 0
}

func (dynic *Dynic) Dynic_sendFlow(v, flow, t int, start *[]int) int {
	graph := dynic.graph
	if v == t {
		return flow
	}

	for ; (*start)[v] < len(graph.nodes[v].neighbours); (*start)[v]++ {

		node := graph.nodes[v]
		forSomeREASON := (*start)[v]

		if (graph.levels[node.neighbours[forSomeREASON]] == (graph.levels[v] + 1)) && node.flow[forSomeREASON] < node.neighboursCap[forSomeREASON] {
			curr_flow := min(flow, node.neighboursCap[forSomeREASON]-node.flow[forSomeREASON])
			temp_flow := dynic.Dynic_sendFlow(node.neighbours[forSomeREASON], curr_flow, t, start)

			if temp_flow > 0 {
				node.flow[forSomeREASON] += temp_flow

				graph.nodes[graph.nodes[v].neighbours[forSomeREASON]].Node_setNeighFlow(v, -temp_flow)
				return temp_flow
			}
		}
	}

	return 0
}

func (dynic *Dynic) Dynic_maxFlow(s, t int) int {

	if s == t {
		return -1
	}

	total := 0

	for dynic.Dynic_BFS(s, t) {

		start := make([]int, dynic.graph.N+1)

		for {
			flow := dynic.Dynic_sendFlow(s, math.MaxInt, t, &start)

			if flow == 0 {
				break
			}

			total += flow
		}
	}

	return total
}
