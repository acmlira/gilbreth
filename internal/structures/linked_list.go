package structures

type LinkedList interface {
	Push(any)
	Pop() any
	Size() int
}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

type linkedList struct {
	head *node
	tail *node
	size int
}

func (l *linkedList) Push(data any) {
	n := &node{data: data}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

func (l *linkedList) Pop() any {
	data := l.tail.data
	l.tail.next = nil
	return data
}

func (l *linkedList) Size() int {
	return l.size
}
