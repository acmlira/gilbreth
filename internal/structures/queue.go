package structures

type Queue interface {
	Enqueue(any)
	Dequeue() any
	Size() int
	Head() any
	IsEmpty() bool
}

type queue struct {
	head *node
	tail *node
	size int
}

func (q *queue) Enqueue(data any) {
	n := &node{data: data}
	if q.size == 0 {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.size++
}

func (q *queue) Dequeue() any {
	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return data
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) Head() any {
	return q.head.data
}

func (q *queue) IsEmpty() bool {
	return q.size == 0
}

func NewQueue() Queue {
	return &queue{}
}
