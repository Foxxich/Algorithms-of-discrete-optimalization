package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	for i := 4; i < 5; i++ {
		K := i
		start := time.Now()
		graph := Graph_newGraph(K)
		graph.Graph_generateEdges()
		graph.Graph_printGraphGLPK("test.jl")

		max_flow := newAlg(graph)
		elapsed := time.Since(start)
		fmt.Println("K ", K)
		fmt.Println("Max ", max_flow.Alg_alg(0, pow(2, K)-1))
		fmt.Println("Used time ", elapsed)
	}
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
