package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const BritneyMessage = "Oops! I did it again"
const StartNode = Node(1)
const NoWeight = -1

func readInputData() (int, Graph) {
	graph := make(Graph)
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	vertices, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	edges, _ := strconv.Atoi(scanner.Text())

	for index := 0; index < edges; index++ {
		scanner.Scan()
		value, _ := strconv.Atoi(scanner.Text())
		first := Node(value)
		scanner.Scan()
		value, _ = strconv.Atoi(scanner.Text())
		second := Node(value)
		scanner.Scan()
		value, _ = strconv.Atoi(scanner.Text())
		graph.add(first, second, value)
	}
	return vertices, graph
}

func solution(vertices int, graph Graph) (result int, err error) {
	if len(graph) == 0 && vertices > int(StartNode) {
		return NoWeight, errors.New(BritneyMessage)
	}

	nodes := NewProcessedNodes(graph)
	nodes[StartNode] = ProcessedNode{}
	var nextErr error
	for nextNode := StartNode; nextErr == nil; nextNode, nextErr = nodes.next(){
		if edges, isFound := graph[nextNode]; isFound {
			for _, edge := range edges {
				if edge.node == nextNode {
					continue
				}
				if pNode := nodes[edge.node]; pNode.weight < edge.weight && !pNode.isProcessedNode {
					nodes[edge.node] = NewProcessedNode(edge.weight, false)
				}
			}
		}
		nodes[nextNode] = NewProcessedNode(nodes[nextNode].weight, true)
	}

	for _, v := range nodes {
		if v.weight == NoWeight {
			return NoWeight, errors.New(BritneyMessage)
		}
		result += v.weight
	}
	return
}

func main() {
	vertices, graph := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	if output, err := solution(vertices, graph); err == nil {
		_, _ = writer.WriteString(strconv.Itoa(output))
	} else {
		_, _ = writer.WriteString(err.Error())
	}
	_ = writer.Flush()
}


type Node int
type Graph map[Node][]Edge
type Edge struct {
	node   Node
	weight int
}
type ProcessedNodes map[Node]ProcessedNode
type ProcessedNode struct {
	weight          int
	isProcessedNode bool
}


func (h *Graph) append(node Node, edge Edge) {
	if value, isFound := (*h)[node]; isFound {
		(*h)[node] = append(value, edge)
	} else {
		(*h)[node] = []Edge{edge}
	}
}

func (h *Graph) add(first Node, second Node, weight int) {
	h.append(second, Edge{node: first, weight: weight})
	h.append(first, Edge{node: second, weight: weight})
}

func NewProcessedNode(weight int, isProcessedNode bool) ProcessedNode {
	return ProcessedNode{
		weight:          weight,
		isProcessedNode: isProcessedNode,
	}
}

func NewProcessedNodes(graph Graph) (processedNodes ProcessedNodes) {
	processedNodes = make(ProcessedNodes)
	for node, _ := range graph {
		processedNodes[node] = NewProcessedNode(NoWeight, false)
	}
	return
}

func (p *ProcessedNodes) next() (Node, error) {
	// Берем следующую необработанную вершину (с максимальным ребром)
	maxNode := StartNode
	maxWeight := NoWeight
	for node, value := range *p {
		if weight := value.weight; weight > maxWeight && !value.isProcessedNode {
			maxNode, maxWeight = node, weight
		}
	}
	if maxWeight == NoWeight {
		return StartNode, errors.New(BritneyMessage)
	}
	return maxNode, nil
}
