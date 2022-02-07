package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	OpenBraces  int = 0
	CloseBraces     = 1
	OpenParenthesis     = 2
	CloseParenthesis = 3
	OpenBrackets     = 4
	CloseBrackets     = 5
)
var BRACKETS = map[string]int {
	"{": OpenBraces,
	"}": CloseBraces,
	"(": OpenParenthesis,
	")": CloseParenthesis,
	"[": OpenBrackets,
	"]": CloseBrackets,
}
const InputFileName = "input.txt"
const OutputFileName = "output.txt"


type Node struct {
	value    int
	previous *Node
}

func solution(commands []int) bool  {
	var tail *Node = nil
	for _ ,command := range commands {
		if command == OpenBrackets || command == OpenParenthesis || command == OpenBraces {
			node := Node{command, tail}
			tail = &node
			continue
		}
		if tail == nil {
			return false
		}
		result := 
			check(OpenBrackets, CloseBrackets, command, tail.value) &&
			check(OpenParenthesis, CloseParenthesis, command, tail.value) &&
			check(OpenBraces, CloseBraces, command, tail.value)
		
		if result == false {
			return false
		}
		tail = tail.previous
	}
	return tail == nil
}

func check(openCommand int, closeCommand int, command int, value int) bool {
	if command == closeCommand {
		return value == openCommand
	}
	return true
}


func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func readInputData() []int {
	inputFile, inputFileErr := os.Open(InputFileName)
	if inputFileErr != nil {
		handleError(inputFileErr)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanRunes)
	var values []int
	for scanner.Scan() {
		value := scanner.Text()
		for k, v := range BRACKETS {
			if k == value {
				values = append(values, v)
			}
		}
	}
	return values
}

func writeData(data bool) {
	outputFile, outputFileErr := os.Create(OutputFileName)
	if outputFileErr != nil {
		handleError(outputFileErr)
	}

	writer := bufio.NewWriter(outputFile)
		if _, err := writer.WriteString(strings.Title(strconv.FormatBool(data))); err != nil {
			handleError(err)
		}
	if _, err := writer.WriteString("\n"); err != nil {
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
	commands := readInputData()
	writeData(solution(commands))
}

