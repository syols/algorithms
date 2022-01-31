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
	next *Node
}

const (
	GET = 1
	PUT = 2
	SIZE = 4
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
	var head *Node = nil
	size := 0
	for _ ,command := range commands {
		if command.action == PUT {
			node := Node{command.value, nil}
			if head == nil {
				head = &node
			}
			if tail != nil  {
				tail.next = &node
			}
			tail = &node
			size++
		}
		if command.action == GET {
			if head == nil {
				if _, err := writer.WriteString("error" + Separator); err != nil {
					handleError(err)
				}
			} else {
				if _, err := writer.WriteString(strconv.Itoa(head.value) + Separator); err != nil {
					handleError(err)
				}
				head = head.next
				size--
			}
		}
		if command.action == SIZE {
			if _, err := writer.WriteString(strconv.Itoa(size) + Separator); err != nil {
				handleError(err)
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
		if text == "put" {
			scanner.Scan()
			value, _ := strconv.Atoi(scanner.Text())
			values[index].action = PUT
			values[index].value = value
		}
		if text == "get" {
			values[index].action = GET
			values[index].value = -1
		}
		if text == "size" {
			values[index].action = SIZE
		}
		index++
	}
	return values
}

func main() {
	commands := readInputData()
	solution(commands)
}

