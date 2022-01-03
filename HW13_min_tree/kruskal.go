package HW13_min_tree

import (
	"fmt"
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

func (g *Graph) Kruskal() [][3]int {
	sort.Stable(g)
	g.buildVertices()

	var tree [][3]int

	for k, e := range g.Edges {
		if k == 0 {
			tree = append(tree, g.Edges[0])
			g.Vertices[e[1]].Pow = 1
			g.Vertices[e[2]].Pow = 1
			continue
		}

		switch p1, p2 := g.Vertices[e[1]].Pow, g.Vertices[e[2]].Pow; {
		case p1 == 0 || p2 == 0:
			tree = append(tree, e)
			g.Vertices[e[1]].Pow += 1
			g.Vertices[e[2]].Pow += 1
			continue
		default:
			//TODO DFS()
		}

		if len(tree) == len(g.Vertices)-1 {
			return tree
		}
	}

	return tree
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

	for _, e := range g.Edges {
		g.Vertices[e[1]].Vertices = append(g.Vertices[e[1]].Vertices, e[2])
		g.Vertices[e[2]].Vertices = append(g.Vertices[e[2]].Vertices, e[1])
	}
	fmt.Println(len(g.Vertices))
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
