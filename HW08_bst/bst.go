package HW08_bst

type Node struct {
	L, R, P *Node
	Val     int
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{}
}

func (t *BST) Insert(v int) {
	if t.root == nil {
		t.root = &Node{Val: v}
		return
	}
	t.root.add(v)
}

func (n *Node) add(v int) {
	switch {
	case v == n.Val:
		return
	case v < n.Val:
		if n.L == nil {
			newNode := &Node{Val: v, P: n}
			n.L = newNode
			return
		}
		n.L.add(v)
	case v > n.Val:
		if n.R == nil {
			newNode := &Node{Val: v, P: n}
			n.R = newNode
			return
		}
		n.R.add(v)
	}
}

func (t *BST) Search(v int) *Node {
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

func (t *BST) Delete(v int) {
	node := t.root.search(v)
	if node == nil {
		return
	}
	node.delete()
}

func (n *Node) delete() {
	switch {
	case n.L == nil && n.R == nil:
		if n.P.L == n {
			n.P.L = nil
		}
		if n.P.R == n {
			n.P.R = nil
		}
		n.P = nil
	case n.L != nil:
		
	case n.R != nil:

	}
}

func (t *BST) List() []int {
	var arr []int
	t.root.traverse(&arr)
	return arr
}

func (n *Node) traverse(arr *[]int) {
	if n.L != nil {
		n.L.traverse(arr)
	}

	*arr = append(*arr, n.Val)
	//println("value:", n.Val)

	if n.R != nil {
		n.R.traverse(arr)
	}

	return
}
