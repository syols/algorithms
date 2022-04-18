package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const Separator = "\n"



func solution(n int) int {
	first, second := 0, 1
	tmp := 0
	for i := 1; i < n + 1; i++ {
		tmp = first + second
		first, second = second % 1000000007, tmp % 1000000007
	}
	return second
}


func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func readInputData() int {
	inputFile, inputFileErr := os.Open(InputFileName)
	if inputFileErr != nil {
		handleError(inputFileErr)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	return count
}

func writeData(value int) {
	outputFile, outputFileErr := os.Create(OutputFileName)
	if outputFileErr != nil {
		handleError(outputFileErr)
	}

	writer := bufio.NewWriter(outputFile)
	if _, err := writer.WriteString(strconv.Itoa(value) + Separator); err != nil {
		handleError(err)
	}

	if err := writer.Flush(); err != nil {
		handleError(err)
	}

	if err := outputFile.Close(); err != nil {
		handleError(err)
	}
}

func main() {
	count := readInputData()
	writeData(solution(count))
}

