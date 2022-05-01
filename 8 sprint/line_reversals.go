package main

import (
	"bufio"
	"os"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const ClosedQuote = ']'
const OpenQuote = '['
const EmptyValue = -1

func readInputData() (result []string) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	result = []string{}
	for ;scanner.Scan(); {
		result = append(result, scanner.Text())
	}
	return
}

func reverse(array []string) (result string) {
	for i := len(array) - 1; i >= 0; i-- {
		value := array[i]
		result += value + " "
	}

	return
}

func main() {
	array := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	result := reverse(array)
	_, _ = writer.WriteString(result)
	_ = writer.Flush()
}
