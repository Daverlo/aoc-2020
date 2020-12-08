package graphs

import (
	"container/list"
	"math"
)

// Compute the max flow from the source to the sink
// on the given graph using the EdmondKarp algorithm
func MaxFlow(g *Graph, source int, sink int) int {
	flow := make([][]int, g.Nodes)
	for i := range flow {
		flow[i] = make([]int, g.Nodes)
	}

	maxFlow := 0
	for {
		parents := make([]int, g.Nodes)
		m := bfsEdmondKarp(g, source, sink, flow, parents)
		if m == 0 {
			break
		}

		maxFlow += m
		v := sink
		for v != source {
			u := parents[v]
			flow[u][v] += m
			flow[v][u] -= m
			v = u
		}
	}

	return maxFlow
}

func bfsEdmondKarp(g *Graph, source int, sink int, flow [][]int, parents []int) int {
	minCap := make([]int, g.Nodes)
	minCap[source] = math.MaxInt32

	for i := range parents {
		parents[i] = -1
	}
	parents[source] = -2

	q := list.New()
	q.PushBack(source)
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		current := e.Value.(int)

		for dst, cost := range g.AdjacencyList[current] {
			if cost-flow[current][dst] > 0 && parents[dst] == -1 {
				parents[dst] = current
				minCap[dst] = min(minCap[current], cost-flow[current][dst])
				if dst == sink {
					return minCap[sink]
				}
				q.PushBack(dst)
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
