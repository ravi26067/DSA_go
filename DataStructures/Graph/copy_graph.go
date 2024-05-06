package main

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// map to store cloned nodes
	clonedNodes := make(map[int]*Node)

	// Queue for BFS traversal (put first node to queue)
	queue := []*Node{node}

	//clone first node
	clonedNodes[node.Val] = &Node{Val: node.Val}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		//iterate through neighbour of
		for _, neighbour := range current.Neighbors {
			if _, ok := clonedNodes[neighbour.Val]; !ok {
				clonedNodes[neighbour.Val] = &Node{Val: neighbour.Val}
				queue = append(queue, neighbour)
			}
			clonedNodes[current.Val].Neighbors = append(clonedNodes[current.Val].Neighbors, clonedNodes[neighbour.Val])
		}
	}
	return clonedNodes[node.Val]
}

// func main() {
// 	node1 := &Node{Val: 1}
// 	node2 := &Node{Val: 2}
// 	node3 := &Node{Val: 3}
// 	node4 := &Node{Val: 4}

// 	node1.Neighbors = []*Node{node2, node4}
// 	node2.Neighbors = []*Node{node1, node3}
// 	node3.Neighbors = []*Node{node2, node4}
// 	node4.Neighbors = []*Node{node1, node3}

// 	node := cloneGraph(node1)
// 	fmt.Println(node)
// }
