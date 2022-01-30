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

type Node struct {
	value    int
	previous *Node
	less *Node
}

const (
	PUSH = 0
	POP = 1
	MAX = 2
)

type Command struct {
	action int
	value int
}

func solution(commands []Command) *Node {
	outputFile, outputFileErr := os.Create(OutputFileName)
	if outputFileErr != nil {
		handleError(outputFileErr)
	}

	writer := bufio.NewWriter(outputFile)
	var tail *Node = nil
	var max *Node = nil
	for _ ,command := range commands {
		if command.action == PUSH {
			node := Node{command.value, tail, nil}
			tail = &node
			if max == nil || command.value > max.value {
				node.less = max
				max = &node
			}
		}
		if command.action == POP {
			if tail == nil {
				writer.WriteString("error" + Separator)
			} else {
				if max == tail {
					max = max.less
				}
				tail = tail.previous
			}
		}
		if command.action == MAX {
			if max == nil {
				if _, err := writer.WriteString("None" + Separator); err != nil {
					handleError(err)
				}
			} else {
				if _, err := writer.WriteString(strconv.Itoa(max.value) + Separator); err != nil {
					handleError(err)
				}
			}
		}
	}

	if err := writer.Flush(); err != nil {
		handleError(err)
	}

	if err := outputFile.Close(); err != nil {
		handleError(err)
	}
	return tail
}


func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func readInputData() []Command {
	inputFile, inputFileErr := os.Open(InputFileName)
	if inputFileErr != nil {
		handleError(inputFileErr)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	var values = make([]Command, count)
	index := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "push" {
			scanner.Scan()
			value, _ := strconv.Atoi(scanner.Text())
			values[index].action = PUSH
			values[index].value = value
		}
		if text == "pop" {
			values[index].action = POP
			values[index].value = -1
		}
		if text == "get_max" {
			values[index].action = MAX
			values[index].value = -1
		}
		index++
	}
	return values
}

func writeData(data []int) {
	outputFile, outputFileErr := os.Create(OutputFileName)
	if outputFileErr != nil {
		handleError(outputFileErr)
	}

	writer := bufio.NewWriter(outputFile)
	for _, value := range data {
			if _, err := writer.WriteString(strconv.Itoa(value) + Separator); err != nil {
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
	commands := readInputData()
	solution(commands)
}

