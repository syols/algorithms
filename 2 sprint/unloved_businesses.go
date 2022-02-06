package main

//import "fmt"

//type ListNode struct {
//	data   string
//	next *ListNode
//}

func Solution(head *ListNode, index int) *ListNode {
	currentIndex := 0
	prev := head
	current := head
	if index == 0 {
		return current.next
	}
	for head.next != nil {
		prev = current
		current = current.next
		currentIndex++
		if currentIndex == index {
			prev.next = current.next
			return head
		}


	}
	prev.next = nil
	return head
}


//func main() {
//	node3 := ListNode{"node3", nil}
//	node2 := ListNode{"node2", &node3}
//	node1 := ListNode{"node1", &node2}
//	node0 := ListNode{"node0", &node1}
//	Solution(&node0, 3)
//}