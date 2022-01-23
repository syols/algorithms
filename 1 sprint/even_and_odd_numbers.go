package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//const maxCapacity = 64
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
	//buffer := make([]byte, 0, maxCapacity)
	//scanner.Buffer(buffer, maxCapacity)

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
		_, _ = writer.WriteString(line + "\n")
	}
	_ = writer.Flush()
}

func main() {
	lines := input()
	values := stringToIntArr(lines[0])
	a := values[0] % 2 == 0
	b := values[1] % 2 == 0
	if b != a {
		output([]string {"FAIL"})
		os.Exit(0)
	}
	c := values[2] % 2 == 0
	if c != a {
		output([]string {"FAIL"})
		os.Exit(0)
	}
	output([]string {"WIN"})
}

