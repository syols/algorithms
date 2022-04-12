package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

func readInputData() (capacity int, result []int) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	capacity = c

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	result = make([]int, count)
	for i := 0; i < count; i++ {
		scanner.Scan()
		cost, _ := strconv.Atoi(scanner.Text())
		result[i] = cost
	}
	sort.Sort(sort.IntSlice(result))
	return
}

func solution(capacity int, array []int) (count int) {
	dp := make([]int, capacity+1)
	dp[0] = 1
	for _, coin := range array {
		for m := coin; m < capacity+1; m++ {
			dp[m] += dp[m - coin]
		}
	}
	return dp[capacity]
}

func main() {
	capacity, array := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	count := solution(capacity, array)
	_, _ = writer.WriteString(strconv.Itoa(count) + "\n")
	_ = writer.Flush()
}
