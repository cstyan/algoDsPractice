package main

// reverse a linked list

import (
	"fmt"
	// "bytes"
)

type node struct {
	value int
	next  *node
}

func (n linkedList) String() string {
	var listString string
	current := n.head
	for {
		listString = fmt.Sprintf("%s %d", listString, current.value)
		if current.next == nil {
			return listString
		}
		listString = fmt.Sprintf("%s %s", listString, "->")
		current = current.next
	}
}

type linkedList struct {
	head *node
}

// func (n *node) String() string {
// 	var listString bytes.Buffer
// 	for {
// 		fmt.Println("asdf")
// 		listString.WriteString(string(n.value))
// 		if n.next == nil {
// 			return listString.String()
// 		}
// 		listString.WriteString(" -> ")
// 		n = n.next
// 	}
// }

// type linkedList interface {
// 	insert
// }

func reverseList(list *linkedList) {
	var prev *node
	current := list.head
	next := current.next
	for {
		current.next = prev
		prev = current
		current = next
		next = next.next

		if current.next == nil {
			current.next = prev
			return
		}
	}
}

func main() {
	last := node{
		value: 10,
		next:  nil,
	}
	middle := node{
		value: 11,
		next:  &last,
	}
	start := node{
		value: 7,
		next:  &middle,
	}

	fmt.Println(linkedList{head: &start}.String())
	reverseList(&linkedList{head: &start})
	fmt.Println(linkedList{head: &last})
}
