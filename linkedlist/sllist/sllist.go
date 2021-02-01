package sllist

import (
	"fmt"
)

//New explicitly returns an empty single linked list object of type *SlinkedList
func New() *SLinkedList {
	return &SLinkedList{len: 0, head: nil}
}

//SLinkedList represents the whole list
type SLinkedList struct {
	len  int
	head *Node
}

//Node is a single item of a linked list
type Node struct {
	Value interface{} //Empty interface to allow any type
	next  *Node
}

//Push pushes a node in front of the list
func (l *SLinkedList) Push(value interface{}) *Node {
	node := &Node{Value: value}
	node.next = l.head
	l.head = node
	l.len++
	return node
}

//Append adds a node to the end of the list
func (l *SLinkedList) Append(value interface{}) *Node {
	searchNode := l.head
	node := &Node{Value: value}

	if searchNode == nil {
		l.head = node
	} else {
		for {
			if searchNode.next == nil {
				searchNode.next = node
				break
			}
			searchNode = searchNode.next
		}
	}
	l.len++
	return node

}

//InsertAfter inserts a value after a given Node in the list.
//Returns nil if the reference node is nil
func (l *SLinkedList) InsertAfter(referenceNode *Node, value interface{}) *Node {
	if referenceNode != nil {
		insertNode := &Node{Value: value, next: referenceNode.next}
		referenceNode.next = insertNode
		l.len++
		return insertNode
	}
	return nil
}

//RemoveNode removes a given node from the List
func (l *SLinkedList) RemoveNode(node *Node) {
	previousNode := l.previousNode(node)

	if previousNode != nil {
		previousNode.next = node.next
		l.len--
		node.next = nil
	}
}

//previousNode retunrs an array of a nodes previous and next node or nil if they are edge nodes
func (l *SLinkedList) previousNode(node *Node) *Node {
	var previousNode *Node
	searchNode := l.head

	for {
		if searchNode == node {
			return previousNode
		} else if searchNode == nil {
			return nil
		}
		previousNode = searchNode
		searchNode = searchNode.next
	}
}

//Len returns the length of the list
func (l *SLinkedList) Len() int {
	return l.len
}

//Clear clears sets the lenght to 0 and the head node to nil to reset and empty the link
func (l *SLinkedList) Clear() {
	l.len = 0
	l.head = nil
}

//Print the content of the list
func (l *SLinkedList) Print() {
	node := l.head
	for {
		fmt.Println(node)
		if node == nil || node.next == nil {
			break
		}
		node = node.next
	}
}
