package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func solution(array []int) bool {
	length := len(array)
	if length == 1 {
		return false
	}

	first, second := array[0:length/2], array[length/2:length]
	sumFirst, sumSecond := sum(first), sum(second)
	difference := sumFirst - sumSecond
	if difference % 2 != 0 {
		return false // Разница между двумя массивами (даже случайно поделены), должна делиться пополам
	}

	expected := (sumFirst + sumSecond) / 2
	previous, current := fillLines(expected, length)
	for m := 1; m <= length; m++ {
		for n := 1; n <= expected; n++ {
			current[n] = previous[n]
			if n >= array[m - 1] {
				current[n] = current[n] || previous[n-array[m-1]]
			}
		}
		previous = current
	}
	return current[expected]
}

func fillLines(m, n int) (previous, current []bool) {
	previous, current = make([]bool, m+1), make([]bool, m+1)
	for i := range previous {
		previous[i] = false
	}
	for i := range current {
		current[i] = true
	}
	return previous, current
}

func sum(array []int) (sum int) {
	for i := range array {
		sum += array[i]
	}
	return
}

func main() {
	array := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	result := solution(array)
	_, _ = writer.WriteString(strings.Title(strconv.FormatBool(result)))
	_ = writer.Flush()
}
