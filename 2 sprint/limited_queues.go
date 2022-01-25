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
	PUSH = 0
	POP = 1
	PEEK = 2
	SIZE = 3
)

type Command struct {
	action int
	value int
}

func solution(commands []Command, maxSize int) *Node {
	outputFile, outputFileErr := os.Create(OutputFileName)
	if outputFileErr != nil {
		handleError(outputFileErr)
	}

	writer := bufio.NewWriter(outputFile)
	var tail *Node = nil
	var head *Node = nil
	size := 0
	for _ ,command := range commands {
		if command.action == PUSH {
			if size == maxSize {
				if _, err := writer.WriteString("error" + Separator); err != nil {
					handleError(err)
				}
				continue
			}
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
		if command.action == POP {
			if head == nil {
				if _, err := writer.WriteString("None" + Separator); err != nil {
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
		if command.action == PEEK {
			if head == nil {
				if _, err := writer.WriteString("None" + Separator); err != nil {
					handleError(err)
				}
			} else {
				if _, err := writer.WriteString(strconv.Itoa(head.value) + Separator); err != nil {
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

func readInputData() ([]Command, int) {
	inputFile, inputFileErr := os.Open(InputFileName)
	if inputFileErr != nil {
		handleError(inputFileErr)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	max_size, _ := strconv.Atoi(scanner.Text())
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
		if text == "peek" {
			values[index].action = PEEK
		}
		if text == "size" {
			values[index].action = SIZE
		}
		index++
	}
	return values, max_size
}

func main() {
	commands, maxSize := readInputData()
	solution(commands, maxSize)
}

