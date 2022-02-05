package main

import (
	"bufio"
	"fmt"
	"os"
)

const OutputFileName = "output.txt"
const Separator = "\n"

//type ListNode struct {
//	data   string
//	next *ListNode
//}

func Solution (head *ListNode) {
	var data []string
	for head.next != nil {
		data = append(data, head.data)
		head = head.next
	}
	data = append(data, head.data)
	writeData(data)
}

func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func writeData(data []string) {
	outputFile, outputFileErr := os.Create(OutputFileName)
	if outputFileErr != nil {
		handleError(outputFileErr)
	}

	writer := bufio.NewWriter(outputFile)
	for _, value := range data {
		if _, err := writer.WriteString(value + Separator); err != nil {
			handleError(err)
		}
	}
	if _, err := writer.WriteString("\n"); err != nil {
		handleError(err)
	}

	if err := writer.Flush(); err != nil {
		handleError(err)
	}

	if err := outputFile.Close(); err != nil {
		handleError(err)
	}
}

//func main() {
//	node3 := ListNode{"node3", nil}
//	node2 := ListNode{"node2", &node3}
//	node1 := ListNode{"node1", &node2}
//	node0 := ListNode{"node0", &node1}
//	Solution(&node0)
//}