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
type Parents []*Node
type Color map[*Node]int
type Queue []*Node
type Nodes map[*Node]Empty
type Node struct {
	value    uint32
	children []*Node
}

func NewNode(value int) Node {
	return Node{
		value:    uint32(value),
		children: []*Node{},
	}
}

// Добавляем узел в граф
func (g *Nodes) add(parent *Node, child *Node) {
	if _, isVisited := (*g)[parent]; !isVisited {
		(*g)[parent] = Empty{}
	}
	parent.children = append(parent.children, child)
}

func readInputData() (*Node, *Node, Nodes) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	count, _ := strconv.Atoi( scanner.Text())
	scanner.Scan()
	rib, _ := strconv.Atoi( scanner.Text())
	var nodes []*Node
	for i := 0; i < count; i++ {
		node := NewNode(i)
		nodes = append(nodes, &node)
	}
	graph := make(Nodes)

	for i := 0; i < rib; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi( scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi( scanner.Text())
		graph.add(nodes[a - 1], nodes[b - 1])
		graph.add(nodes[b - 1], nodes[a - 1])
	}
	scanner.Scan()
	begin, _ := strconv.Atoi( scanner.Text())
	scanner.Scan()
	end, _ := strconv.Atoi( scanner.Text())
	return nodes[begin-1], nodes[end-1], graph
}

func (g *Nodes) bfs(node *Node, end *Node, visited Color, queue Queue) int {
	count := 0
	visited[node] = count
	queue = append(queue, node)

	for len(queue) > 0 {
		value := queue[0]
		queue = queue[1:]

		if value == end {
			return visited[value]
		}

		for _, c:=range value.children {
			if _, isFound := visited[c]; !isFound {
				visited[c] = visited[value] + 1
				queue = append(queue, c)
			}
		}
	}
	return -1
}

func main() {
	begin, end, graph := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)

	visited, queue := Color{}, Queue{}
	_, _ = writer.WriteString(strconv.Itoa(graph.bfs(begin, end, visited, queue)) + " ")
	_ = writer.Flush()
}

