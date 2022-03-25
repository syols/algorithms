package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Heap struct {
	heap []int
	data *InputData
}

const StartIndex = 1

type OutputData []string
type InputData []Competitor
type Competitor struct {
	login     string
	completed int
	penalty   int
}

func NewParticipant(login string, completed int, penalty int) Competitor {
	return Competitor{
		login:     login,
		completed: completed,
		penalty:   penalty,
	}
}

func solution(data InputData) OutputData {
	length := len(data) + StartIndex
	result := make(OutputData, length-StartIndex)
	heap := Heap{
		heap: make([]int, length),
		data: &data,
	}

	for index := StartIndex; index < length; index++ {
		heap.add(index)
	}

	for index := 0; index < length-StartIndex; index++ {
		max := heap.max()
		result[index] = data[max-StartIndex].login
	}
	return result
}

func (h *Heap) add(participantIndex int) {
	h.heap[participantIndex] = participantIndex
	h.siftUp(participantIndex)
}

func (h *Heap) max() (result int) {
	result = h.heap[StartIndex]
	position := len(h.heap) - StartIndex
	h.heap[StartIndex] = h.heap[position]
	h.heap = h.heap[:position]
	h.siftDown(StartIndex)
	return
}

func (h *Heap) siftUp(index int) {
	for ;index != StartIndex; {
		parentIndex := index / 2
		if h.less(h.heap[parentIndex], h.heap[index]) {
			h.swap(parentIndex, index)
		}
		index = parentIndex
	}
}

func (h *Heap) siftDown(index int) {
	for ;;{
		left, right := 2*index, 2*index+1
		size := len(h.heap) - StartIndex
		if size < left {
			return
		}

		largest := left
		if (right <= size) && (h.less(h.heap[left], h.heap[right])) {
			largest = right
		}

		if !h.less(h.heap[index], h.heap[largest]) {
			return
		}

		h.swap(index, largest)
		index = largest
	}
}

func (h *Heap) swap(left int, right int) {
	h.heap[right], h.heap[left] = h.heap[left], h.heap[right]
}

func (h *Heap) less(parentIndex int, childIndex int) bool {
	first := &(*h.data)[parentIndex-StartIndex]
	second := &(*h.data)[childIndex-StartIndex]

	if first.completed == second.completed {
		if first.penalty == second.penalty {
			return first.login > second.login
		}
		return first.penalty > second.penalty
	}
	return first.completed < second.completed
}

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

func readInputData() (inputData InputData) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	inputData = make(InputData, count)
	for index := 0; index < count; index++ {
		scanner.Scan()
		login := scanner.Text()
		scanner.Scan()
		completed, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		penalty, _ := strconv.Atoi(scanner.Text())
		inputData[index] = NewParticipant(login, completed, penalty)
	}
	return
}

func main() {
	data := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	output := solution(data)
	for _, v := range output {
		_, _ = writer.WriteString(v)
		_, _ = writer.WriteString(fmt.Sprintln())
	}
	_ = writer.Flush()
}
