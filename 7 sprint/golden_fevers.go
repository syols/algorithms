package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

type Task struct {
	cost int
	weight int
}

func (a *Task) less(b *Task) bool{
	return a.cost > b.cost
}

func readInputData() (capacity int, result []Task) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	capacity = c

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	result = make([]Task, count)
	for i := 0; i < count; i++ {
		scanner.Scan()
		cost, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		weight, _ := strconv.Atoi(scanner.Text())
		result[i] = Task{ cost: cost, weight: weight}
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].less(&result[j])
	})
	return
}

func solution(capacity int, array []Task) (count int){
	for _, task := range array{
		if capacity >= task.weight {
			capacity -= task.weight
			count += task.cost * task.weight
		} else {
			count += task.cost * capacity
			return count
		}
	}
	return
}

func main() {
	capacity, array := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	count := solution(capacity, array)
	_, _ = writer.WriteString(strconv.Itoa(count) + "\n")
	_ = writer.Flush()
}
