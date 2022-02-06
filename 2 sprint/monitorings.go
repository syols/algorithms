package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const Separator = " "

func solution(row int, column int, arr []int) [][]int {
	var result = make([][]int, column)
	for i := range result {
		result[i] = make([]int, row)
	}
	for j:=0; j < row; j++ {
		for i:=0; i < column; i++ {
			value := arr[(column * j) + i]
			result[i][j] = value
		}
	}
	return result
}


func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func readInputData() (int, int, []int)  {
	inputFile, inputFileErr := os.Open(InputFileName)
	if inputFileErr != nil {
		handleError(inputFileErr)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	j, _ := strconv.Atoi(scanner.Text())

	var values []int
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		values = append(values, value)
	}
	return i, j, values
}

func writeData(data [][]int) {
	outputFile, outputFileErr := os.Create(OutputFileName)
	if outputFileErr != nil {
		handleError(outputFileErr)
	}

	writer := bufio.NewWriter(outputFile)
	for _, row := range data {
		for _, value := range row {
			if _, err := writer.WriteString(strconv.Itoa(value) + Separator); err != nil {
				handleError(err)
			}
		}
		if _, err := writer.WriteString("\n"); err != nil {
			handleError(err)
		}
	}

	if err := writer.Flush(); err != nil {
		handleError(err)
	}

	if err := outputFile.Close(); err != nil {
		handleError(err)
	}
}

func main() {
	i, j, values := readInputData()
	writeData(solution(i, j, values))
}


