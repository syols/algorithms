package main

//type Node struct {
//	value  int
//	left   *Node
//	right  *Node
//}

func Solution(root *Node) bool {
	if root.left != nil && root.right != nil {
		var l []int
		var r []int
		lmr(root.left, &l)
		lmr(root.right, &r)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return equal(l, r)
	}
	return  root.left == nil && root.right == nil
}

func lmr(root *Node, arr *[]int) {
	if root == nil {
		*arr = append(*arr, -1)
		return
	}
	lmr(root.left, arr)
	*arr = append(*arr, root.value)
	lmr(root.right, arr)
}

func lml(root *Node, arr *[]int) {
	if root.right != nil {
		lmr(root.right, arr)
	}
	*arr = append(*arr, root.value)
	if root.left != nil {
		lmr(root.left, arr)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}


//func main() {
//	node8 := Node{0, nil, nil}
//	node7 := Node{1, nil, nil}
//	node6 := Node{1, nil, nil}
//	node5 := Node{0, nil, nil}
//	node4 := Node{3, &node7, &node8}
//	node3 := Node{3, &node5, &node6}
//	node2 := Node{2, nil, &node4}
//	node1 := Node{2, &node3, nil}
//	node0 := Node{0, &node1, &node2}
//
//	print(Solution(&node0))
//}