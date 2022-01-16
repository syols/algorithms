package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFile = "input.txt"
const outputFile = "output.txt"
var exist = struct{}{}

func stringToIntArr(line string) []int {
	split := strings.Split(line, " ")
	result := make([]int, len(split))
	for i, v := range split {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

func input(lineNum map[int]struct{}) []string {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	var result []string
	currentLine := 0
	for scanner.Scan() {
		if _, isFind := lineNum[currentLine]; isFind {
			result = append(result, scanner.Text())
		}
		currentLine++
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
	m := make(map[int]struct{})
	m[0] = exist
	m[1] = exist

	lines := input(m)
	i := stringToIntArr(lines[0])[0]
	j := stringToIntArr(lines[1])[0]

	m = make(map[int]struct{})
	arrayStarts := 2
	arrayEnds := arrayStarts + i
	m[arrayEnds] = exist
	m[arrayEnds + 1] = exist
	lines = input(m)
	x := stringToIntArr(lines[0])[0]
	y := stringToIntArr(lines[1])[0]

	findLine := arrayStarts + x
	pos := 0
	m = make(map[int]struct{})
	if (findLine - 1) >= arrayStarts {
		m[findLine - 1] = exist
		pos = 1
	}
	m[findLine] = exist
	if (findLine + 1) < arrayEnds {
		m[findLine + 1] = exist
	}

	lines = input(m)
	var result []int
	for index, line := range lines {
		values := stringToIntArr(line)
		if index != pos {
			result = append(result, values[y])
			continue
		}

		if left := y - 1; left >= 0 {
			result = append(result, values[left])
		}
		if right := y + 1; right < j {
			result = append(result, values[right])
		}

	}
	sort.Ints(result)
	output(result)
}

