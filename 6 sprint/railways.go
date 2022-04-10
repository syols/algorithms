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

func readInputData() (Queue, Graph) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	var nodes Queue
	for i := 0; i < count; i++ {
		node := NewNode(i)
		nodes = append(nodes, &node)
	}
	graph := Graph{}
	for parent := 0; scanner.Scan(); parent++ {
		for offset, char := range scanner.Text() {
			child := parent + offset + 1
			if char == RedNode {
				graph.add(nodes[child], nodes[parent])
			} else {
				graph.add(nodes[parent], nodes[child])
			}
		}
	}
	return nodes, graph
}

func NewNode(value int) Node {
	return Node{
		value: value,
		children: []*Node{},
		incomingEdgeCount: 0,
	}
}

func main() {
	nodes, graph := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)

	if solution(nodes, graph) {
		_, _ = writer.WriteString(YesMessage)
	} else {
		_, _ = writer.WriteString(NoMessage)
	}
	_ = writer.Flush()
}

type Node struct {
	value int
	children Queue
	incomingEdgeCount int
}
type Graph map[int]*Node
type Queue []*Node


// Добавляем узел в граф
func (g *Graph) add(parent *Node, child *Node) {
	parent.children = append(parent.children, child)
	child.incomingEdgeCount++
	if _, isFound := (*g)[parent.value]; !isFound {
		(*g)[parent.value] = parent
	}
}

// Находим вершину с которой начать поиск
func solution(nodes Queue, graph Graph) bool {
	queue := Queue{}
	processed := 0
	for _, node := range nodes{
		if node.incomingEdgeCount == 0 {
			queue = append(queue, node)
			processed++
			break
		}
	}

	// Если нет такой классной вершины, то берем первую попавшуюся
	if len(queue) == 0 {
		queue = append(queue, graph[0])
		processed++
	}
	return noContainsCircle(nodes, queue, processed)
}

// Заменил на BFS
func noContainsCircle(nodes Queue, queue Queue, processed int) bool {
	count := len(nodes)
	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		for _, node := range parent.children {
			if processed > count {
				return false
			}
			node.incomingEdgeCount--
			if node.incomingEdgeCount == 0 {
				queue = append(queue, node)
				processed++
			}
		}
	}
	return processed == count
}
