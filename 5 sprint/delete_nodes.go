package main

//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

func remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}

	if node.left == nil && node.right == nil {
		if node.value == key {
			return nil
		} else {
			return node
		}
	}

	parent, deleted := deletedNode(node, key)
	if n := checkDeleted(node, deleted, parent); n != nil {
		return n
	}

	parentReplacement, replacement := replacementNode(deleted)
	if parent != nil {
		substitution(parent, deleted, replacement)
	} else {
		node = replacement
	}

	rebinding(parentReplacement, deleted, replacement)
	return node
}

func checkDeleted(node *Node, deleted *Node, parent *Node) *Node {
	if deleted == nil {
		return node
	}

	if deleted.left == nil && deleted.right == nil {
		substitution(parent, deleted, nil)
		return node
	}

	if deleted.left == nil {
		if parent == nil {
			return deleted.right
		}
		substitution(parent, deleted, deleted.right)
		return node
	}
	return nil
}

func rebinding(parentReplacement *Node, deleted *Node, replacement *Node) {
	if parentReplacement != deleted {
		if replacement.left == nil && replacement.right == nil {
			parentReplacement.right = nil
		} else {
			parentReplacement.right = replacement.left
		}
	}

	//Привязываем детей удаляемого элемента
	if replacement != deleted.left {
		replacement.left = deleted.left
	}

	if replacement != deleted.right {
		replacement.right = deleted.right
	}
}

func substitution(parent *Node, deleted *Node, child *Node) {
	if parent.left == deleted {
		parent.left = child
	} else {
		parent.right = child
	}
}

func deletedNode(node *Node, key int) (parentNode *Node, childNode *Node) {
	for {
		if node == nil || node.value == key {
			childNode = node
			return
		}
		parentNode = node
		if node.value > key {
			node = node.left
		} else {
			node = node.right
		}
	}
}

func replacementNode(parentNode *Node) (*Node, *Node) {
	childNode := parentNode.left

	for {
		if childNode.right == nil {
			return parentNode, childNode
		}
		parentNode = childNode
		childNode = childNode.right
	}
}

//func main() {
//	node4 := Node{4, nil, nil}
//	node3 := Node{3, nil, &node4}
//	node2 := Node{2, nil, &node3}
//	node1 := Node{1, nil, &node2}
//	newHead := remove(&node1, 3)
//	print(newHead.value)
//}
