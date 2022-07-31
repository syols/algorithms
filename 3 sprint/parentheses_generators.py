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

const (
	OpenBrace = '('
	CloseBrace = ')'
)

func generate(length int) []string {
	var result []string
	if length == 1 {
		result = []string{string(OpenBrace), string(CloseBrace)}
	} else {
		for _, r := range generate(length - 1){
			result = append(result, r + string(OpenBrace))
			result = append(result, r + string(CloseBrace))
		}
	}
	return result
}

func solution(n int) []string {
	var verified []string
	for _, value := range generate(n*2) {
		var k string
		isOk := true
		for _, r := range value {
			if r == OpenBrace {
				k = k + string(OpenBrace)
			}
			if r == CloseBrace {
				l:= len(k)
				if l == 0 {
					isOk = false
					break
				}
				k = k[:l-1]
			}
		}
		if len(k) == 0 && isOk {
			verified = append(verified, value)
		}
	}
	return verified
}


func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func readInputData() int  {
	inputFile, inputFileErr := os.Open(InputFileName)
	if inputFileErr != nil {
		handleError(inputFileErr)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())

	return i
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
	n := readInputData()
	writeData(solution(n))
}

