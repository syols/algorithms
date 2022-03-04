package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const ErrorMessage = "None"

const Put = "put"
const Get = "get"
const Delete = "delete"

const NotFoundError = "not found"
const MaxBasketSize = 10000

type Map []*Node
type Node struct {
	key   int
	value int
	next  *Node
}

func position(key int) int {
	return key % MaxBasketSize
}

func NewMap() Map {
	return make(Map, MaxBasketSize)
}

func (m *Map) Put(key, value int) (err error) {
	index := position(key)
	begin := (*m)[index]
	if node, _, _ := m.FindNode(key, begin); node != nil {
		node.value = value
		return nil
	}

	(*m)[index] = &Node{
		key:   key,
		value: value,
		next:  begin,
	}
	return
}

func (m *Map) Get(key int) (int, error) {
	if node, _, err := m.FindNode(key, (*m)[position(key)]); err != nil {
		return 0, err
	} else {
		return node.value, err
	}
}

func (m *Map) Delete(key int) (int, error) {
	index := position(key)
	begin := (*m)[index]

	if curr, prev, err := m.FindNode(key, begin); err == nil {
		if curr == begin {
			(*m)[index] = curr.next
		} else {
			prev.next = curr.next
		}
		return curr.value, nil
	}
	return 0, errors.New(NotFoundError)
}

func (m *Map) FindNode(key int, begin *Node) (*Node, *Node, error) {
	// NOTE Были мысли вернуть предыдущий, но сделал как проще. Замечание понял, учту)
	var prev *Node
	for curr := begin; curr != nil; curr = curr.next {
		if curr.key == key {
			return curr, prev, nil
		}
		prev = curr
	}
	return nil, nil, errors.New(NotFoundError)
}


type ActionMap = map[string]interface{}
type Command struct {
	action string
	key    *int
	value  *int
}

func createCommand(text string) Command {
	var command Command
	values := strings.Fields(text)
	command.action = values[0]
	key, _ := strconv.Atoi(values[1])
	command.key = &key
	if len(values) > 2 {
		value, _ := strconv.Atoi(values[2])
		command.value = &value
	}
	return command
}

func solution(command Command, actions ActionMap) (*int, error) {
	action := actions[command.action]
	if command.value != nil {
		return nil, action.(func(int, int) error)(*command.key, *command.value)
	}
	value, err := action.(func(int) (int, error))(*command.key)
	return &value, err
}

func main() {
	customMap := NewMap()
	var actions = ActionMap{
		Get:    customMap.Get,
		Put:    customMap.Put,
		Delete: customMap.Delete,
	}

	inputFile, _ := os.Open(InputFileName)
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	for index := 0; index < count; index++ {
		scanner.Scan()
		if result, err := solution(createCommand(scanner.Text()), actions); err != nil {
			_, _ = writer.WriteString(ErrorMessage + fmt.Sprintln())
		} else if result != nil {
			_, _ = writer.WriteString(strconv.Itoa(*result) + fmt.Sprintln())
		}
	}

	_ = inputFile.Close()
	_ = writer.Flush()
}
