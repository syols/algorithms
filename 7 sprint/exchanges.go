package main

import (
	"bufio"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

func readInputData() (result []int) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	result = make([]int, count)
	for i := 0; i < count; i++ {
		scanner.Scan()
		value, _ := strconv.Atoi(scanner.Text())
		result[i] = value
	}
	return
}

func solution(array []int) (count int) {
	current, previous := 0, 0
	for i, value := range array{
		if i == 0 {
			previous = value
			continue
		}
		current = value
		if previous < value {
			count += value - previous
		}
		previous = current
	}
	return
}

func main() {
	array := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	_, _ = writer.WriteString(strconv.Itoa(solution(array)))
	_ = writer.Flush()
}
