package main

import (
	"fmt"
	"math"
)

func solution(n int, k int) int {
	if n == 0 {
		return 1
	}
	if n <= 1 {
		return n
	}

	k = int(math.Pow(10, float64(k)))
	prev := 0
	curr := 1
	for i := 0; i < n; i++ {
		prev, curr = curr, (prev+curr) % k
	}
	return curr
}

func readInputData() (int, int) {
	var n int
	fmt.Scan(&n)
	var k int
	fmt.Scan(&k)
	return n, k
}

func writeData(value int) {
	fmt.Println(value)
}

func main() {
	n, k := readInputData()
	writeData(solution(n, k))
}

