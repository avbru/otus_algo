package containers

type List interface {
	Len() int
	Front() *Item
	Back() *Item
	PushFront(v interface{}) *Item
	PushBack(v interface{}) *Item
	Remove(i *Item)
	MoveToFront(i *Item)
}

type Item struct {
	Value interface{}
	Next  *Item
	Prev  *Item
}

type list struct {
	length    int
	firstNode *Item
	lastNode  *Item
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *Item {
	return l.firstNode
}

func (l *list) Back() *Item {
	return l.lastNode
}

func (l *list) PushBack(v interface{}) *Item {
	newItem := &Item{Value: v, Next: nil, Prev: nil}

	if l.length == 0 {
		l.firstNode, l.lastNode = newItem, newItem
	} else {
		newItem.Next = l.lastNode
		l.lastNode.Prev = newItem
		l.lastNode = newItem
	}

	l.length++
	return newItem
}

func (l *list) PushFront(v interface{}) *Item {
	newItem := &Item{Value: v, Next: nil, Prev: nil}

	if l.length == 0 {
		l.firstNode, l.lastNode = newItem, newItem
	} else {
		newItem.Prev = l.firstNode
		l.firstNode.Next = newItem
		l.firstNode = newItem
	}

	l.length++
	return newItem
}

func (l *list) Remove(i *Item) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.lastNode = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.firstNode = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *Item) {
	l.PushFront(i.Value)
	l.Remove(i)
}
