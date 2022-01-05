package HW09_b_trees

type Node struct {
	L, R, P *Node
	Val     int
}

type Random struct {
	len  int
	root *Node
}

func NewRandom() *Random {
	return &Random{}
}

func (t *Random) Insert(v int) {
	if t.root == nil {
		t.root = &Node{Val: v}
		t.len++
		return
	}

	if t.root.add(v) {
		t.len++
	}

}

func (n *Node) add(v int) bool {
	switch {
	case v == n.Val:
		return false
	case v < n.Val:
		if n.L == nil {
			newNode := &Node{Val: v, P: n}
			n.L = newNode
			return true
		}
		return n.L.add(v)
	case v > n.Val:
		if n.R == nil {
			newNode := &Node{Val: v, P: n}
			n.R = newNode
			return true
		}
		return n.R.add(v)
	}
	return false
}

func (t *Random) Len() int {
	return t.len
}

func (t *Random) Search(v int) *Node {
	if t.root == nil {
		return nil
	}
	return t.root.search(v)
}

func (n *Node) search(v int) *Node {
	switch {
	case v == n.Val:
		return n
	case v < n.Val:
		if n.L != nil {
			return n.L.search(v)
		}
		return nil
	case v > n.Val:
		if n.R != nil {
			return n.R.search(v)
		}
		return nil
	}
	return nil
}
