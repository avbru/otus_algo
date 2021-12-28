package HW08_bst

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (t *BST) Delete(v int) {
	node := t.root.search(v)
	if node == nil {
		return
	}

	//Delete root
	if node.P == nil {
		if node.L == nil && node.R == nil {
			t.root = nil
			return
		}

		if node.R != nil && node.L == nil {
			t.root = node.R
			t.root.L = node.L
			t.root.P = nil
			return
		}

		if node.L != nil && node.R == nil {
			t.root = node.L
			t.root.P = nil
			return
		}
	}

	node.delete()
}

func (n *Node) delete() {
	switch {
	case n.L == nil && n.R == nil:
		n.removeLeaf()
	case n.L != nil && n.R != nil:
		if n.R.L == nil {
			n.Val, n.R.Val = n.R.Val, n.Val
			n.R = n.R.R
			return
		}

		left, val := n.R.L, n.R.Val
		for left.L != nil {
			if left.L.Val < val {
				val = left.L.Val
				left = left.L
			}
		}

		left.Val, n.Val = n.Val, left.Val //Swap values
		left.delete()
	default:
		n.removeSingle()
	}
}

func (n *Node) removeLeaf() {
	switch {
	case n.P.L == n:
		n.P.L = nil
	case n.P.R == n:
		n.P.R = nil
	}
}

func (n *Node) removeSingle() {
	switch {
	case n.L != nil:
		if n.P.L == n {
			n.P.L = n.L
			n.L.P = n.P
		}
		if n.P.R == n {
			n.P.R = n.L
			n.L.P = n.P
		}
	case n.R != nil:
		if n.P.L == n {
			n.P.L = n.R
			n.R.P = n.P
		}
		if n.P.R == n {
			n.P.R = n.R
			n.R.P = n.P
		}
	}
	n.P, n.L, n.R = nil, nil, nil
}

func (t *BST) List() []int {
	var arr []int
	if t.root == nil {
		return nil
	}
	t.root.traverse(&arr)
	return arr
}

func (n *Node) traverse(arr *[]int) {
	if n.L != nil {
		n.L.traverse(arr)
	}

	*arr = append(*arr, n.Val)

	if n.R != nil {
		n.R.traverse(arr)
	}

	return
}

func (t *BST) Draw() {
	var arr []string
	if t.root == nil {
		println("empty tree")
		return
	}

	t.root.draw(&arr)
	for _, raw := range arr {
		println(raw)
	}
}

func (n *Node) draw(arr *[]string) {
	lev := n.level()

	//Ensure level string exists
	if len(*arr) <= lev {
		*arr = append(*arr, strings.Repeat("     ", 100))
	}

	parentVal := 0
	if n.P != nil {
		parentVal = n.P.Val
	}
	yellow := "\033[33m"
	white := "\033[0m"
	value := fmt.Sprintf("%d(%s%d%s) ", n.Val, yellow, parentVal, white)

	if n.P == nil {
		(*arr)[lev] = insert(len((*arr)[lev])/2, (*arr)[lev], value, false)
	} else {
		pos := strings.Index((*arr)[lev-1]+"(", strconv.Itoa(n.P.Val))
		//IS LEFT
		if n.P.L == n {
			(*arr)[lev] = insert(pos-30/lev/lev, (*arr)[lev], value, true)
		}
		//IS RIGHT
		if n.P.R == n {
			(*arr)[lev] = insert(pos+30/lev/lev, (*arr)[lev], value, false)
		}
	}

	if n.R != nil {
		n.R.draw(arr)
	}
	if n.L != nil {
		n.L.draw(arr)
	}

	return
}

func (n *Node) level() int {
	if n == nil {
		return 0
	}
	count := 0
	for n.P != nil {
		count++
		n = n.P
	}

	return count
}

func insert(index int, s, val string, isLeft bool) string {
	if isLeft {
		return s[:index-len(val)] + val + s[index:]
	}
	return s[:index] + val + s[index:]
}
