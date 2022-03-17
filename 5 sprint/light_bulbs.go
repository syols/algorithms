package main

//
//type Node struct {
//	value  int
//	left   *Node
//	right  *Node
//}
//

func Solution(root *Node) int {
	maxValue := root.value
	if root.left != nil {
		value := Solution(root.left)
		if maxValue < value {
			maxValue = value
		}
	}

	if root.right != nil {
		value := Solution(root.right)
		if maxValue < value {
			maxValue = value
		}
	}
	return maxValue
}

//func main() {
//	node1 := Node{1, nil, nil}
//	node2 := Node{-5, nil, nil}
//	node3 := Node{3, &node1, &node2}
//	node4 := Node{2, &node3, nil}
//	if Solution(&node4) != 3 {
//		panic("WA")
//	}
//}
