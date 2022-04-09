package main

import (
	"bufio"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const YesMessage = "YES"
const NoMessage = "NO"
const RedNode = 'R'

type Empty struct{}
type Nodes map[*Node]Empty
type Node struct {
	value    uint16
	children map[uint16]Empty
}

func NewNode(value int) Node {
	return Node{
		value:    uint16(value),
		children: make(map[uint16]Empty),
	}
}


// Добавляем узел в граф
func (g *Nodes) add(parent *Node, child *Node) {
	if _, isVisited := (*g)[parent]; !isVisited {
		(*g)[parent] = Empty{}
	}
	parent.children[child.value] = Empty{}
}

func readInputData() []*Node {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	count, _ := strconv.Atoi( scanner.Text())
	scanner.Scan()
	var nodes []*Node
	for i := 0; i < count; i++ {
		node := NewNode(i)
		nodes = append(nodes, &node)
	}
	graph := make(Nodes)

	for scanner.Scan() {
		a, _ := strconv.Atoi( scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi( scanner.Text())
		graph.add(nodes[a - 1], nodes[b - 1])
	}
	return nodes
}

func main() {
	nodes := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		for j := 0; j < len(nodes); j++ {
			if _, isFind := node.children[uint16(j)]; isFind {
				_, _ = writer.WriteString("1 ")
			} else {
				_, _ = writer.WriteString("0 ")
			}
		}
		_, _ = writer.WriteString("\n")
	}
	_ = writer.Flush()
}

