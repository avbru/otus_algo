package HW13_min_tree

func (g *Graph) Boruvka() [][3]int {
	g.buildVertices()
	cc := make(CC, len(g.Vertices))
	for i := 0; i < len(g.Vertices); i++ {
		cc[i] = i
	}

	tree := make(map[[3]int]struct{})
	cheapest := make([]int, len(g.Vertices))
	for {
		if len(tree) >= len(g.Vertices)-1 {
			break
		}

		for i, _ := range cheapest {
			cheapest[i] = -1
		}

		for i, e := range g.Edges {
			u := cc.find(e[1])
			v := cc.find(e[2])

			if u == v {
				continue
			}

			cheapU := cheapest[u]
			cheapV := cheapest[v]

			if cheapU == -1 || g.Edges[cheapU][0] > g.Edges[i][0] {
				cheapest[u] = i
			}

			if cheapV == -1 || g.Edges[cheapV][0] > g.Edges[i][0] {
				cheapest[v] = i
			}
		}

		for _, v := range cheapest {
			if v != -1 {
				cc.merge(g.Edges[v][1], g.Edges[v][2])
				tree[g.Edges[v]] = struct{}{}
			}
		}
	}

	var tr [][3]int
	for k, _ := range tree {
		tr = append(tr, k)
	}

	return tr
}
