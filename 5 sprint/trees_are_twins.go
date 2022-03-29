package main

//type Node struct {
//	value  int
//	left   *Node
//	right  *Node
//}

func Solution(root1 *Node, root2 *Node) bool {
	if root1 == nil || root2 == nil {
		if root1 == nil && root2 == nil {
			return true
		}
		return false
	}

	isOk := root1.value == root2.value
	if !isOk {
		return false
	}
	isLeftMirror := Solution(root1.left, root2.left)
	isRightMirror := Solution(root1.right, root2.right)
	return isLeftMirror && isRightMirror && isOk
}

//func main() {
//	node2 := Node{3, nil, nil}
//	node1 := Node{1, &node2, nil}
//	node0 := Node{0, nil, &node1}
//	node20 := Node{3, nil, nil}
//	node10 := Node{1, &node20, nil}
//	node00 := Node{0, &node10, nil}
//
//	print(Solution(&node0, &node00))
//}