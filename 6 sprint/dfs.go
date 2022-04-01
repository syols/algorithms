package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const YesMessage = "YES"
const NoMessage = "NO"
const RedNode = 'R'

type Empty struct{}
type Parents []*Node
type Color map[uint32]int
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


func readInputData() (*Node, Nodes) {
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
	pos, _ := strconv.Atoi( scanner.Text())
	for _, node := range nodes {
		sort.Slice(node.children, func(i, j int) bool {
			return node.children[i].value < node.children[j].value
		})
	}
	return nodes[pos-1], graph
}

func (g *Nodes) dfs(node *Node, visited Color, writer *bufio.Writer) {
	_, _ = writer.WriteString(strconv.Itoa(int(node.value + 1)) + " ")
	visited[node.value] = 1
	for _, c:=range node.children {
		if _, isFound := visited[c.value]; !isFound {
			g.dfs(c, visited, writer)
		}
	}
	visited[node.value] = 2
}


func main() {
	node, graph := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)

	graph.dfs(node, make(Color), writer)

	_ = writer.Flush()
}

