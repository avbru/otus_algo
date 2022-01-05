package HW14_dijkstra

import (
	"math"
)

const (
	W = iota
	V1
	V2
)

type Vertex struct {
	Used  bool
	W     int
	From  int
	Edges []int
}

type Graph struct {
	Edges    [][3]int //W V1 V2
	Vertices []Vertex
}

func (g *Graph) Dijkstra(node int) {
	g.buildVertices()
	g.Vertices[node].W, g.Vertices[node].From = 0, node

	for {
		g.Vertices[node].Used = true
		v1 := g.Vertices[node]

		for _, e := range v1.Edges {
			v2id := g.peekV2(node, e)
			v2 := g.Vertices[v2id]

			if g.Edges[e][W]+v1.W < v2.W {
				v2.W = g.Edges[e][W] + v1.W
				v2.From = node
			}

			g.Vertices[v2id] = v2
		}

		g.Vertices[node] = v1

		//Peek cheapest Next
		cheapest := g.peekCheapest(node)
		if cheapest == -1 {
			break
		}
		node = cheapest
	}
}

func (g *Graph) peekV2(v1, e int) int {
	if g.Edges[e][V1] == v1 {
		return g.Edges[e][V2]
	}
	return g.Edges[e][V1]
}

func (g *Graph) peekCheapest(vId int) int {
	v := g.Vertices[vId]
	w := math.MaxInt
	v2id := -1
	for _, e := range v.Edges {
		v2 := g.peekV2(vId, e)
		if g.Vertices[v2].Used {
			continue
		}
		if g.Edges[e][W] < w {
			w = g.Edges[e][W]
			v2id = v2
		}
	}
	return v2id
}

func (g *Graph) buildVertices() {
	uniques := make(map[int]Vertex)

	for idx, e := range g.Edges {
		v, ok := uniques[e[1]]
		if !ok {
			uniques[e[1]] = Vertex{
				W:    0,
				From: -1,
			}
		}
		v.Edges = append(v.Edges, idx)
		uniques[e[1]] = v

		v, ok = uniques[e[2]]
		if !ok {
			uniques[e[2]] = Vertex{
				W:    0,
				From: -1,
			}
		}
		v.Edges = append(v.Edges, idx)
		uniques[e[2]] = v

	}

	g.Vertices = make([]Vertex, len(uniques))
	for k, v := range uniques {
		v.W = math.MaxInt
		g.Vertices[k] = v
	}
}
