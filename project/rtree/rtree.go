package rtree

import (
	"math"
	"strconv"
)

const (
	x = 0
	y = 1
)

var count int
var M int
var m int

type Rtree struct {
	m    int
	M    int
	Root *Node
}

func New(nM int) *Rtree {
	tr := Rtree{
		m:    nM / 2,
		M:    nM,
		Root: &Node{isLeaf: true, data: "N0"},
	}
	m = nM / 2
	M = nM
	return &tr
}

type Point [2]float64

type Rec struct {
	Low Point
	Top Point
}

type Node struct {
	parent  *Node
	isLeaf  bool
	bbox    Rec
	data    string
	entries []*Node
}

var visits int

func (t *Rtree) Search(r Rec) string {
	visits = 0
	res := ""
	t.Root.search(r, res)
	println("VISITS: ", visits)
	return res
}

func (n *Node) search(r Rec, res string) {
	for _, e := range n.entries {
		if overlap(r, e.bbox) {
			res = res + e.data
			println(e.data)
			visits++
			e.search(r, res)
		}
	}
}

func overlap(r1, r2 Rec) bool {
	return r1.Low[x] <= r2.Top[x] &&
		r2.Low[x] <= r1.Top[x] &&
		r1.Low[y] <= r2.Top[y] &&
		r2.Low[y] <= r1.Top[y]
}

func (t *Rtree) Insert(bbox Rec, data string) {
	if t.Root == nil {
		newRoot := &Node{nil, true, bbox, "N0", nil}
		newEntry := &Node{newRoot, true, bbox, data, nil}
		t.Root = newRoot
		t.Root.entries = append(t.Root.entries, newEntry)
		t.Root.setBBOX()
		count = 1
		return
	}

	leaf := t.Root.chooseLeaf(bbox)
	newEntry := &Node{leaf.parent, true, bbox, data, nil}
	leaf.Insert(newEntry, t)
}

func (n *Node) Insert(in *Node, t *Rtree) {
	in.parent = n
	n.entries = append(n.entries, in)
	n.setBBOX()
	if len(n.entries) <= M {
		return
	}

	n2 := n.splitNode()

	//Если ROOT
	if n.parent == nil {
		count++
		newRoot := &Node{nil, false, Rec{}, "N" + strconv.Itoa(count), nil}
		//newRoot.entries = append(newRoot.entries, t.Root.entries...)
		//n.data = "N" + strconv.Itoa(count)
		newRoot.Insert(n, t)
		newRoot.Insert(n2, t)
		t.Root = newRoot
		//fmt.Println("new root          :", newRoot)
		//fmt.Println("new root entries 1:", newRoot.entries[0])
		//fmt.Println("new root entries 2: ", newRoot.entries[1])
		newRoot.setBBOX()
		return
	}

	n.setBBOX()

	//fmt.Printf("befor: %v\n", n.parent.String())
	n.parent.Insert(n2, t)
	n.parent.isLeaf = false
	//fmt.Printf("after: %v\n", n.parent.String())
	//fmt.Printf("inserting to parent: %v\n", n2)
	return
}

func (n *Node) splitNode() *Node {
	count++
	d := "N" + strconv.Itoa(count)

	n1, n2 := &Node{data: n.data, parent: n}, &Node{data: d, parent: n}
	//fmt.Println("NODE SPLIT", n)
	peekSeeds(n, n1, n2)

	for {
		n1.setBBOX()
		n2.setBBOX()
		n1.isLeaf = true
		n2.isLeaf = true
		if len(n.entries) == 0 {
			break
		}
		if len(n.entries) == m-len(n2.entries) {
			n2.entries = append(n2.entries, n.entries...)
			n.entries = nil
			break
		}
		if len(n.entries) == m-len(n1.entries) {
			n1.entries = append(n1.entries, n.entries...)
			n.entries = nil
			break
		}

		idx := peekNext(n, n1, n2)
		s1 := bbox(n.entries[idx].bbox, n1.bbox).S()
		s2 := bbox(n.entries[idx].bbox, n2.bbox).S()
		switch {
		case s1 == s2:
			if len(n1.entries) < len(n2.entries) {
				n1.entries = append(n1.entries, n.entries[idx])
			} else {
				n2.entries = append(n2.entries, n.entries[idx])
			}
		case s1 < s2:
			n1.entries = append(n1.entries, n.entries[idx])
		case s2 < s1:
			n2.entries = append(n2.entries, n.entries[idx])
		}

		n.entries = remove(n.entries, idx)

	}

	n2.setBBOX()
	n2.setLeaf()
	n.entries = append(n.entries, n1.entries...)
	n.setBBOX()
	n.setLeaf()

	//fmt.Println("AFTER SPLIT")
	//fmt.Println("node entries 1:", n)
	//fmt.Println("node entries 2:", n2)
	return n2
}

func peekNext(n, n1, n2 *Node) int {
	cost1 := make([]float64, len(n.entries))
	cost2 := make([]float64, len(n.entries))
	for i, e := range n.entries {
		cost1[i] = bbox(n1.bbox, e.bbox).S() - n1.bbox.S() - e.bbox.S()
		cost2[i] = bbox(n2.bbox, e.bbox).S() - n2.bbox.S() - e.bbox.S()
	}
	max, idx := 0.0, 0
	for i := 0; i < len(cost1); i++ {
		diff := math.Abs(cost1[i] - cost2[i])
		if diff > max {
			max = diff
			idx = i
		}
	}

	return idx
}

func peekSeeds(n, n1, n2 *Node) {
	d, idx, jdx := 0.0, 0, 0
	for i := 0; i < len(n.entries); i++ {
		for j := i + 1; j < len(n.entries); j++ {

			dd := bbox(n.entries[i].bbox, n.entries[j].bbox).S()
			newD := dd - n.entries[i].bbox.S() - n.entries[j].bbox.S()
			if newD > d {
				d = newD
				idx = i
				jdx = j
			}
		}
	}
	//fmt.Printf("peekseeds %d %v %v \n", len(n.entries), n.entries[idx].data, n.entries[jdx].data)
	n1.entries = append(n1.entries, n.entries[idx])
	n2.entries = append(n2.entries, n.entries[jdx])
	n.entries = remove(n.entries, idx)
	n.entries = remove(n.entries, jdx-1)
}

func remove(slice []*Node, s int) []*Node {
	return append(slice[:s], slice[s+1:]...)
}

func (n *Node) chooseLeaf(box Rec) *Node {
	if n.isLeaf {
		return n
	}

	//Choose subtree (minimal enlargement)
	enlargement, idx := math.Inf(0), 0
	for i, r := range n.entries {
		enl := bbox(r.bbox, box).S() - r.bbox.S()
		if enl < enlargement {
			enlargement = enl
			idx = i
		}
	}

	return n.entries[idx].chooseLeaf(box)
}

func bbox(r1, r2 Rec) Rec {
	minX := min(r1.Low[x], r2.Low[x])
	minY := min(r1.Low[y], r2.Low[y])
	maxX := max(r1.Top[x], r2.Top[x])
	maxY := max(r1.Top[y], r2.Top[y])
	return Rec{
		Low: Point{minX, minY},
		Top: Point{maxX, maxY},
	}
}

func (n *Node) setBBOX() {
	n.bbox = n.entries[0].bbox
	for i := 1; i < len(n.entries); i++ {
		n.bbox = bbox(n.bbox, n.entries[i].bbox)
	}
}

func (n *Node) String() string {
	s := n.data + "|"
	for _, e := range n.entries {
		s = s + " " + e.data
	}
	return s
}

func min(i, j float64) float64 {
	if i < j {
		return i
	}
	return j
}

func max(i, j float64) float64 {
	if i > j {
		return i
	}
	return j
}

func (r Rec) S() float64 {
	return math.Abs((r.Low[x] - r.Top[x]) * (r.Low[y] - r.Top[y]))
}

func (n *Node) setLeaf() {
	n.isLeaf = false
	for _, v := range n.entries {
		if v.data[0] == 'R' {
			n.isLeaf = true
		}
	}
}
