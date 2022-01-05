package HW13_min_tree

import (
	"sort"
)

type Vertex struct {
	Pow      int
	Vertices []int
}

type Graph struct {
	Edges    [][3]int
	Vertices []Vertex
}

func (g *Graph) KruskalDFS() [][3]int {
	sort.Stable(g)
	g.buildVertices()

	var tree [][3]int

	for k, e := range g.Edges {
		if len(tree) == len(g.Vertices)-1 {
			return tree
		}

		if k == 0 {
			tree = append(tree, g.Edges[0])
			g.addEdge(e[1], e[2])
			continue
		}

		switch p1, p2 := g.Vertices[e[1]].Pow, g.Vertices[e[2]].Pow; {
		case p1 == 0 || p2 == 0:
			tree = append(tree, e)
			g.addEdge(e[1], e[2])
		default:
			var found bool
			used := make([]bool, len(g.Vertices))
			g.dfs(e[1], e[2], used, &found)
			if !found {
				tree = append(tree, e)
				g.addEdge(e[1], e[2])
			}
		}

	}

	return tree
}
func (g *Graph) addEdge(v1, v2 int) {
	g.Vertices[v1].Pow = 1
	g.Vertices[v2].Pow = 1
	g.Vertices[v1].Vertices = append(g.Vertices[v1].Vertices, v2)
	g.Vertices[v2].Vertices = append(g.Vertices[v2].Vertices, v1)
}

func (g *Graph) dfs(v1, v2 int, used []bool, found *bool) {
	vert := g.Vertices[v1]
	if used[v1] == true {
		return
	}
	used[v1] = true

	for _, v := range vert.Vertices {
		if v == v2 {
			*found = true
			return
		}
		g.dfs(v, v2, used, found)
	}
}

func (g *Graph) buildVertices() {
	uniques := make(map[int]struct{})
	for _, e := range g.Edges {
		if _, ok := uniques[e[1]]; !ok {
			uniques[e[1]] = struct{}{}
		}
		if _, ok := uniques[e[2]]; !ok {
			uniques[e[2]] = struct{}{}
		}
	}

	g.Vertices = make([]Vertex, len(uniques))
}

func (g Graph) Len() int {
	return len(g.Edges)
}

func (g Graph) Less(i, j int) bool {
	if g.Edges[i][0] < g.Edges[j][0] {
		return true
	}
	return false
}

func (g Graph) Swap(i, j int) {
	g.Edges[i], g.Edges[j] = g.Edges[j], g.Edges[i]
}
