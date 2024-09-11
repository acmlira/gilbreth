package structures

import "testing"

func TestQueue(t *testing.T) {
	L := new(queue)
	L.Enqueue("A")
	L.Enqueue("B")
	L.Enqueue("C")

	if L.Dequeue() != "A" {
		t.Fail()
	}
	if L.Dequeue() != "B" {
		t.Fail()
	}
	if L.Dequeue() != "C" {
		t.Fail()
	}
}
