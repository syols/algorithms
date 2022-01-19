package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
	//start := time.Now()

	lines := input()
	first := lines[0]

	r, _ := strconv.Atoi(first)
	result := "True"

	for {
		if r == 1 {
			break
		}
		if r % 4 != 0 {
			result = "False"
			break
		}
		r = r/4
	}
	output([]string{result})
	//
	//elapsed := time.Since(start)
	//log.Printf("Took %s", elapsed)
}
