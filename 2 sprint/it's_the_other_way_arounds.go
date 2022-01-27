package main

//type ListNode struct {
//	data   string
//	next *ListNode
//}

func Solution(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var current = head
	for i:=head.next; i != nil; i = i.next {
		current.next = prev
		prev = current
		if i.next != nil {
			prev = current
			current = i
		}
		current = i
	}
	current.next = prev
	return current
}

//
//func main() {
//	node3 := ListNode{"node3", nil}
//	node2 := ListNode{"node2", &node3}
//	node1 := ListNode{"node1", &node2}
//	node0 := ListNode{"node0", &node1}
//	val:=Solution(&node3)
//	fmt.Print(val)
//}