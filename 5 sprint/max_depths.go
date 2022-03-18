package main

import "math"

//type Node struct {
//	value  int
//	left   *Node
//	right  *Node
//}

func Solution(root *Node) int {
	return max(root, 0)
}

func max(root *Node, depth int) int {
	if root == nil {
		return 0
	}
	depth = int(math.Max(float64(max(root.right, depth)), float64(max(root.left, depth)))) + 1
	return depth
}

//func main() {
//	node1 := Node{1, nil, nil}
//	node2 := Node{4, nil, nil}
//	node3 := Node{3, &node1, &node2}
//	node4 := Node{8, nil, nil}
//	node5 := Node{5, &node3, &node4}
//	if Solution(&node5) != 3 {
//		panic("WA")
//	}
//}
