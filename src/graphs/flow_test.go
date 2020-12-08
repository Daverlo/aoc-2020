package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFlow(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)

	flow := MaxFlow(g, 0, 2)
	assert.Equal(t, 1, flow)
}
