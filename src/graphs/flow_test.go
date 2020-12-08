package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFlow(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)

	maxFlow, _ := MaxFlow(g, 0, 2)
	assert.Equal(t, 1, maxFlow)
}

func TestMinCut(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 100)
	g.AddEdge(1, 2, 1)

	cut, minCut := MinCut(g, 0, 2)
	assert.Equal(t, 1, cut)
	assert.Equal(t, 1, minCut[0].Src)
	assert.Equal(t, 2, minCut[0].Dst)
	assert.Equal(t, 1, minCut[0].Cost)
}
