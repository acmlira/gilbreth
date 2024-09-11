package structures

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	L := new(linkedList)
	L.Push("A")
	L.Push("B")
	L.Push("C")

	if L.Pop() != "C" {
		t.Fail()
	}
	if L.Pop() != "B" {
		t.Fail()
	}
	if L.Pop() != "A" {
		t.Fail()
	}
}
