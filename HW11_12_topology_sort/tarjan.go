package HW11_12_topology_sort

type color int

const (
	white color = iota
	grey
	black
)

type Vertex struct {
	N     int
	Color color
	List  []*Vertex
}

type Graph []*Vertex

func Tarjan(vec [][]int) []int {
	var sorted []int
	graph := vec2Graph(vec)

	for {
		v, ok := graph.peekWhite()
		if !ok {
			break
		}
		graph.DFS(v, &sorted)
	}

	for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
		sorted[i], sorted[j] = sorted[j], sorted[i]
	}

	return sorted
}

func (g Graph) DFS(n int, res *[]int) {
	v := g[n]
	v.Color = grey

	found := false
	for _, vl := range v.List {
		if vl.Color == white || vl.Color == grey {
			found = true
			break
		}
	}

	if !found {
		v.Color = black
	}

	for _, vl := range v.List {
		if vl.Color == white {
			g.DFS(vl.N, res)
		}
	}

	*res = append(*res, v.N)
}

func vec2Graph(vec [][]int) (graph Graph) {

	for k, _ := range vec {
		graph = append(graph, &Vertex{k, 0, nil})
	}

	for i, list := range vec {
		for _, v := range list {
			graph[i].List = append(graph[i].List, graph[v])
		}
	}
	return
}

func (g Graph) peekWhite() (int, bool) {
	for k, v := range g {
		if v.Color == white {
			return k, true
		}
	}
	return 0, false
}
