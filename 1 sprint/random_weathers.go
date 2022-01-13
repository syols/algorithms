package main

import (
	"bufio"
	"os"
	"strconv"
)

const inputFile = "input.txt"
const outputFile = "output.txt"

func input() int {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	result, previous, current, next, step := 0,0,0,0,0
	for scanner.Scan() {
		if step == 0 {
			count, _ := strconv.Atoi(scanner.Text())
			if count == 1 {
				return 1
			}
			step++
			continue
		}
		previous = current
		current = next
		next, _ = strconv.Atoi(scanner.Text())
		if step > 2 {
			if current > next && current > previous {
				result++
			}
		}
		if step == 2 && current > next{
			result++
		}
		step++
	}
	if next > current{
		result++
	}
	return result
}

func output(values []int)  {
	file, _ := os.Create(outputFile)
	writer := bufio.NewWriter(file)
	for _, value := range values {
		_, _ = writer.WriteString(strconv.Itoa(value) + " ")
	}
	_ = writer.Flush()
}


func main() {
	output([]int{input()})
}

