package containers

type Queue []interface{}

func (q *Queue) Enqueue(el interface{}) {
	*q = append(*q, el)
}

func (q *Queue) Dequeue() interface{} {
	el := (*q)[0]
	*q = (*q)[1:]
	return el
}

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) List() []interface{} {
	return q
}
