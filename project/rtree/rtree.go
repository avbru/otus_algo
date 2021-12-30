package rtree

import "fmt"

type RTree struct {
	maxEntries int
	minEntries int
	root       *Node
}

func New(nMax, nMin int) *RTree {
	return &RTree{maxEntries: nMax, minEntries: nMin}
}

type Node struct {
	count   int
	Entries []*Entry
}

type Entry struct {
	Low   [2]float64
	High  [2]float64
	Node  *Node
	Value string
}

func (t *RTree) Insert(e *Entry) {
	// Empty Root
	if t.root == nil {
		root := &Node{}
		root.Entries = append(root.Entries, e)
		t.root = root
	}
	//Overflow

}

func (t *RTree) String() {
	t.root.ToString()
}

func (n *Node) ToString() {
	for _, v := range n.Entries {
		fmt.Printf("%v", v)
		if v.Node != nil {
			v.Node.ToString()
		}
	}
}
