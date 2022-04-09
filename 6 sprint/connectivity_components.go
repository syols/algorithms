package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const YesMessage = "YES"
const NoMessage = "NO"
const RedNode = 'R'

type Empty struct{}
type Parents []*Node
type Color map[*Node]int
type RColor map[int][]int
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


func readInputData() ([]*Node, Nodes) {
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
	//scanner.Scan()
	//pos, _ := strconv.Atoi( scanner.Text())
	//for _, node := range nodes {
	//	sort.Slice(node.children, func(i, j int) bool {
	//		return node.children[i].value < node.children[j].value
	//	})
	//}
	return nodes, graph
	//return nodes[pos-1], graph
}

func (g *Nodes) dfs(node *Node, visited Color, count int) {
	visited[node] = count
	for _, c:=range node.children {
		if visited[c] != count {
			g.dfs(c, visited, count)
		}
	}
}


func main() {
	nodes, graph := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)

	visited := make(Color, len(nodes))
	componentCount := 1
	for _, node:= range nodes {
		visited[node] = componentCount
	}

	for node, color := range visited {
		if color == 1 {
			graph.dfs(node, visited, componentCount)
			componentCount++
		}
	}
	rev := reverse(visited)
	_, _ = writer.WriteString(strconv.Itoa(len(rev)) + "\n")

	var result [][]int
	for _, v := range rev {
		sort.Ints(v)
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	
	for _, v := range result {
		_, _ = writer.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v)), " "), "[]") + "\n")
	}
	_ = writer.Flush()
}

func reverse(m Color) RColor {
	n := make(RColor, len(m))
	for k, v := range m {
		if _, isFound := n[v]; isFound {
			n[v] = append(n[v], int(k.value + 1))
		} else {
			n[v] = []int{int(k.value + 1)}
		}

	}
	return n
}

