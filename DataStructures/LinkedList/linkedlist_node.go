package main

type Node struct {
	data int
	next *Node
}

func NewNode(data int, next *Node) *Node {
	node := new(Node)
	node.data = data
	node.next = next
	return node
}
