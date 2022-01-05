package HW13_min_tree

import (
	"sort"
)

type CC []int

func (g *Graph) Kruskal() [][3]int {
	sort.Stable(g)
	g.buildVertices()

	cc := make(CC, len(g.Vertices))
	for i := 0; i < len(g.Vertices); i++ {
		cc[i] = i
	}

	var tree [][3]int

	for _, e := range g.Edges {
		if len(tree) == len(g.Vertices)-1 {
			return tree
		}
		if len(tree) == 0 {
			cc.merge(e[1], e[2])
			tree = append(tree, g.Edges[0])

		}

		if cc.find(e[1]) == cc.find(e[2]) {
			continue
		}

		cc.merge(e[1], e[2])
		tree = append(tree, e)
	}

	return tree
}

func (s CC) find(v int) int {
	if s[v] == v {
		return v
	}
	root := s.find(s[v])
	s[s[v]] = root
	return root
}

func (s CC) merge(v1, v2 int) {
	v1Root := s.find(v1)
	v2Root := s.find(v2)
	if v1Root != v2Root {
		s[v1Root] = v2Root
	}
}
