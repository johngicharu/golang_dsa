package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insert(data int) {
	new_node_ptr := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = new_node_ptr

		return
	}

	next_node := list.head

	for next_node.next != nil {
		next_node = next_node.next
	}

	next_node.next = new_node_ptr
}

func (node *Node) String() (p string) {
	if node == nil {
		return "{}"
	}

	return fmt.Sprintf("{data: %v, next: %p}", node.data, node.next)
}

func main() {
	linked_list := LinkedList{}

	linked_list.insert(1)
	linked_list.insert(2)

	fmt.Println(linked_list.String())
}

func (list *LinkedList) String() (p string) {
	p += "["

	if list.head == nil {
		p += "]"

		return
	}

	next_node := *list.head

	for {
		p += fmt.Sprintf("%v, ", &next_node)

		if next_node.next == nil {
			break
		}

		next_node = *next_node.next
	}

	p += "]"

	return
}
