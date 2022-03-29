package main

//type Node struct {
//	value  int
//	left   *Node
//	right  *Node
//}

func Solution(root *Node) bool {
	return isBinarySearch(root, -2147483648, 2147483647)
}

func isBinarySearch(root *Node, lo int, hi int) bool {
	if root == nil {
		return true
	}

	if root.value <= lo || root.value >= hi {
		return false
	}
	return isBinarySearch(root.left, lo, root.value) && isBinarySearch(root.right, root.value, hi)
}

//func main() {
//	node1 := Node{1, nil, nil}
//	node2 := Node{4, nil, nil}
//	node3 := Node{3, &node1, &node2}
//	node4 := Node{8, nil, nil}
//	node5 := Node{5, &node3, &node4}
//	if !Solution(&node5) {
//		panic("WA")
//	}
//	node2.value = 5
//	if Solution(&node5) {
//		panic("WA")
//	}
//}
