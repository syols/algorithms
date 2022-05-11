package main

import (
	"bufio"
	"math"
	"os"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

func readInputData() (first string, second string) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	first = scanner.Text()
	scanner.Scan()
	second = scanner.Text()
	return
}

func solution(pass, base string) int {

	min := int(math.Min(float64(len(base)), float64(len(pass))))
	diff := len(base)-len(pass)
	if int(math.Abs(float64(diff))) > 1 {
		return 2
	}
	dist := 0
	for i, j := 0, 0; i < min; i, j = i + 1, j + 1 {
		a, b := pass[j], base[i]
		if a != b {
			dist++
			if dist > 1 {
				return dist
			}
			if diff < 0 {
				i--
			} else if diff > 0 {
				j--
			}
		}
	}
	return 0
}

func main() {
	first, second := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	distance := solution(first, second)
	if distance > 1 {
		_, _ = writer.WriteString("FAIL")
	} else {
		_, _ = writer.WriteString("OK")
	}

	_ = writer.Flush()
}
