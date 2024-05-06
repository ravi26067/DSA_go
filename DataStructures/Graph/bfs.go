package main

import "fmt"

// Node to create a node of graph
type Node struct {
	Val       int
	Neighbors []*Node
}

func bfsTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Println("BFS Traversal")

	visited := make(map[int]bool)

	//maintain queue for bfs and push the first node
	queue := []*Node{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.Val] {
			continue
		}

		visited[current.Val] = true

		fmt.Printf("%d ", current.Val)

		for _, neighbour := range current.Neighbors {
			queue = append(queue, neighbour)
		}
	}
	fmt.Println()

}

func dfsTraversal(node *Node) {

	fmt.Println("DFS Traversal")

	if node == nil {
		return
	}

	visited := make(map[int]bool, 0)
	dfsRec(node, visited)
	fmt.Println()
}

func dfsRec(node *Node, visited map[int]bool) {

	if visited[node.Val] {
		return
	}

	fmt.Printf("%d ", node.Val)
	visited[node.Val] = true

	for _, neighbour := range node.Neighbors {
		dfsRec(neighbour, visited)
	}
}

func main() {
	node0 := &Node{Val: 0}
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}

	node0.Neighbors = []*Node{node1, node2}
	node1.Neighbors = []*Node{node0, node3}
	node2.Neighbors = []*Node{node0, node4}
	node3.Neighbors = []*Node{node1, node4, node5}
	node4.Neighbors = []*Node{node2, node3, node5}
	node5.Neighbors = []*Node{node3, node4, node6}
	node6.Neighbors = []*Node{node5}

	bfsTraversal(node0)
	dfsTraversal(node0)
}
