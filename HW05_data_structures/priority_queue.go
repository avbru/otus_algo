package HW05_data_structures

import (
	"github.com/avbru/algo/containers"
)

type PriorityQueue interface {
	Enqueue(val interface{}, priority int)
	Dequeue() interface{}
}

type PQueue struct {
	priorities []containers.List
}

func NewPQueue(max int) *PQueue {
	var p []containers.List
	for i := 0; i <= max-1; i++ {
		p = append(p, containers.NewList())
	}
	return &PQueue{priorities: p}
}

func (q *PQueue) Enqueue(val interface{}, priority int) {
	list := q.priorities[priority]
	list.PushBack(val)
}

func (q *PQueue) Dequeue() interface{} {
	for i := len(q.priorities) - 1; i >= 0; i-- {
		list := q.priorities[i]
		if list.Len() == 0 {
			continue
		}
		el := list.Front()
		list.Remove(el)
		return el.Value
	}

	return nil
}
