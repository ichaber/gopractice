package main

import (
	"fmt"

	"github.com/ichaber/linkedlist/sllist"
)

func main() {

	var test *sllist.SLinkedList = sllist.New()
	test.Push(2)
	removeNode := test.Push(3)
	test.InsertAfter(removeNode, "Nonono")
	test.Push(4)
	test.Push(5)
	test.Push(6)
	test.Append(7)
	test.Append(8)
	test.RemoveNode(removeNode)
	test.Print()
	fmt.Println(test.Len())

}
