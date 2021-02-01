package sllist

import (
	"testing"
)

func TestList(t *testing.T) {
	l := New()

	n1 := l.Push(1)
	if l.head != n1 {
		t.Errorf("Head of list should be the first node, but is %v", n1)
	}
	l.Push(2)
	n3 := l.Push(3)
	l.Push(4)
	l.Push(5)
	l.Push(6)
	if l.len != 6 {
		t.Errorf("Length of list is %v, but should be 6", l.len)
	}
	l.Append(7)
	l.Append(8)
	if l.len != 8 {
		t.Errorf("Length of list is %v, but should be 8", l.len)
	}
	stringNode := l.InsertAfter(n3, "StringNode")
	if l.len != 9 {
		t.Errorf("Length of list is %v, but should be 8", l.len)
	}
	if n3.next != stringNode {
		t.Errorf("Node after reference Node should be the newly inserted node, but is %v", n3.next)
	}
	prevStringNode := l.previousNode(stringNode)
	if prevStringNode != n3 {
		t.Errorf("Previous Node of the newly inserted stringNode should be the node after it's been inserted, but is %v", prevStringNode)
	}
	l.RemoveNode(n3)
	if l.len != 8 {
		t.Errorf("Length of list is %v, but should be 8", l.len)
	}
	l.Print()
	l.Clear()
	if l.len != 0 {
		t.Errorf("List should be empty, but has length %v", l.len)
	}

	emptyNode := l.previousNode(n3)
	if emptyNode != nil {
		t.Error("nil expected for a previous node in an empty list")
	}
	emptyRefNode := l.InsertAfter(nil, "Will not be in list")
	if emptyRefNode != nil {
		t.Error("nil expected for insert after nil node")
	}

	l.Append("NewItem")
	if l.len != 1 {
		t.Errorf("Length of list is %v, but should be 1", l.len)
	}

}
