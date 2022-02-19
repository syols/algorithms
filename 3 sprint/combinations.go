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

const (
	Two="abc"
	Three="def"
	Four="ghi"
	Five="jkl"
	Six="mno"
	Seven="pqrs"
	Eight="tuv"
	Nine="wxyz"
)

func generate(values []string) []string {
	var result []string
	length := len(values)
	value := values[length - 1]
	if length == 1 {
		for _, v := range value {
			result = append(result, string(v))
		}
	} else {
		for _, r := range generate(values[0:length-1]) {
			for _, v := range value {
				result = append(result, r+string(v))
			}
		}
	}
	return result
}


func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func readInputData() []string  {
	inputFile, inputFileErr := os.Open(InputFileName)
	if inputFileErr != nil {
		handleError(inputFileErr)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanRunes)
	var values []string
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		switch value {
			case 2:
				values = append(values, Two)
			case 3:
				values = append(values, Three)
			case 4:
				values = append(values, Four)
			case 5:
				values = append(values, Five)
			case 6:
				values = append(values, Six)
			case 7:
				values = append(values, Seven)
			case 8:
				values = append(values, Eight)
			case 9:
				values = append(values, Nine)
		}
	}
	return values
}

func writeData(data []string) {
	outputFile, outputFileErr := os.Create(OutputFileName)
	if outputFileErr != nil {
		handleError(outputFileErr)
	}

	writer := bufio.NewWriter(outputFile)
	for _, value := range data {
		if _, err := writer.WriteString(value + Separator); err != nil {
			handleError(err)
		}
	}
	_ = writer.Flush();
	_ = outputFile.Close();
}

func main() {
	values := readInputData()
	writeData(generate(values))
}

