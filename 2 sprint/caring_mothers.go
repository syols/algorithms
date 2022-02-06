package main

//type ListNode struct {
//	data   string
//	next *ListNode
//}

func Solution(head *ListNode, data string) int {
	currentIndex := 0
	current := head
	if current.data == data {
		return 0
	}
	for current.next != nil {
		current = current.next
		currentIndex++
		if current.data == data {
			return currentIndex
		}
	}
	return -1
}


//func main() {
//	node3 := ListNode{"node3", nil}
//	node2 := ListNode{"node2", &node3}
//	node1 := ListNode{"node1", &node2}
//	node0 := ListNode{"node0", &node1}
//	val:=Solution(&node0, "node0")
//	val=Solution(&node0, "node1")
//	val=Solution(&node0, "node3")
//	val=Solution(&node0, "node23")
//	fmt.Print(val)
//}