package main

import "math"

//type Node struct {
//	value  int
//	left   *Node
//	right  *Node
//}

func Solution(root *Node) bool {
	isOK, _ :=  isBalanced(root, -1);
	return isOK
}

func isBalanced(root *Node, depth float64) (bool, float64) {
	if root == nil {
		return true, -1
	}

	isLeftOk, leftDepth := isBalanced(root.left, depth + 1)
	isRightOk, rightDepth := isBalanced(root.right, depth + 1)

	isChildBalanced := math.Abs(leftDepth-rightDepth) <= 1
	height := math.Max(leftDepth, rightDepth) + 1
	return isLeftOk && isRightOk && isChildBalanced, height
}

//func main() {
//	node8 := Node{8, nil, nil}
//	node7 := Node{7, nil, nil}
//	node6 := Node{6, nil, nil}
//	node5 := Node{5, nil, nil}
//	node4 := Node{4, &node7, &node8}
//	node3 := Node{3, &node5, &node6}
//	node2 := Node{2, nil, &node4}
//	node1 := Node{1, &node3, nil}
//	node0 := Node{0, &node1, &node2}
//	r := Solution(&node0)
//	print(r)
//}
