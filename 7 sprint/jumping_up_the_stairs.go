package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

func readInputData() (n int, k int) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	_n, _ := strconv.Atoi(scanner.Text())
	n = _n

	scanner.Scan()
	_k, _ := strconv.Atoi(scanner.Text())
	k = _k
	return n, k
}

func solution(k, n int) (count int) {
	dp := make([]int, n)
	dp[0] = 1
	for i:=1; i<n; i++ {
		dp[i] = 0
		for j:=int(math.Max(0, float64(i-k))); j<i; j++ {
			dp[i] += dp[j]
			dp[i] %= 1000000007
		}
	}
	return dp[n - 1]
}

func main() {
	n, k := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	count := solution(k, n)
	_, _ = writer.WriteString(strconv.Itoa(count) + "\n")
	_ = writer.Flush()
}
