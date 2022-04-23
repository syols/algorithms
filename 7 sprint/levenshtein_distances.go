package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

func readInputData() (first string, second string) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	first = scanner.Text()
	scanner.Scan()
	second = scanner.Text()
	return
}

func solution(first, second string) int {
	lengthFirst, lengthSecond := len(first), len(second)
	if lengthFirst > lengthSecond {
		return distance(lengthSecond, second, first)
	}
	return distance(lengthFirst, first, second)
}

func distance(length int, f, s string) int {
	current, previous := fillLines(length)
	for i := 1; i <= len(s); i++ {
		for k := range current {
			previous[k] = current[k]
			if k == 0 {
				current[k] = i
				continue
			}

			if f[k-1] == s[i-1] {
				current[k] = min(previous[k]+1, current[k-1]+1, previous[k-1])
			} else {
				current[k] = min(previous[k]+1, current[k-1]+1, previous[k-1]+1)
			}

		}
	}
	return current[length]
}

func fillLines(length int) ([]int, []int) {
	current := make([]int, length+1)
	for i := range current {
		current[i] = i
	}
	return current, make([]int, length+1)
}

func min(first, second, third int) int {
	min := math.Min(float64(first), float64(second))
	return int(math.Min(min, float64(third)))
}

func main() {
	first, second := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	_, _ = writer.WriteString(strconv.Itoa(solution(first, second)))
	_ = writer.Flush()
}
