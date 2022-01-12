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
	second := lines[1]

	a := len(first) - 1
	b := len(second) - 1
	var result []string
	c := 0
	for a >= 0 || b >= 0 {
		val := c
		if a >= 0 {
			r, _ := strconv.Atoi(string(first[a]))
			val += r
		}
		a--

		if b >= 0 {
			r, _ := strconv.Atoi(string(second[b]))
			val += r
		}
		b--

		result = append(result, strconv.Itoa(val%2))
		c = val / 2
	}

	if c != 0 {
		result = append(result, strconv.Itoa(c))
	}
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}
	output(result)
	//
	//elapsed := time.Since(start)
	//log.Printf("Took %s", elapsed)
}
