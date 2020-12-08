package graphs

type Adjacent map[int]int

type Graph struct {
	Nodes int
	Edges int

	InDegrees     []int
	AdjacencyList []Adjacent
}

func NewGraph(nodes int) *Graph {
	g := new(Graph)

	g.Nodes = nodes
	g.InDegrees = make([]int, nodes)
	g.AdjacencyList = make([]Adjacent, nodes)
	for i := range g.AdjacencyList {
		g.AdjacencyList[i] = make(Adjacent)
	}

	return g
}

func (g *Graph) AddEdge(src int, dst int, cost int) {
	g.AdjacencyList[src][dst] = cost
	g.InDegrees[dst]++
	g.Edges++
}

func (g *Graph) RemoveEdge(src int, dst int) {
	_, present := g.AdjacencyList[src][dst]
	if present {
		delete(g.AdjacencyList[src], dst)
		g.InDegrees[src]--
		g.Edges--
	}
}

func (g *Graph) GetEdge(src int, dst int) (int, bool) {
	cost, present := g.AdjacencyList[src][dst]
	return cost, present
}
