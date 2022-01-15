package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const maxCapacity = 64
const inputFile = "input.txt"
const outputFile = "output.txt"


func stringToIntArr(line string) []int {
	split := strings.Split(line, " ")
	result := make([]int, len(split))
	for i, v := range split {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

func input() []string {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	buffer := make([]byte, 0, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func output(lines []string)  {
	file, _ := os.Create(outputFile)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, _ = writer.WriteString(line)
	}
	_ = writer.Flush()
}

func main() {
	lines := input()
	value := stringToIntArr(lines[0])[0]
	var result []string
	for {
		result = append([]string{ strconv.Itoa(value % 2) }, result...)
		value = value / 2
		if value == 0 {
			break
		}
	}
	output(result)
}

