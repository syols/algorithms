package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxCountQueueMessage = "max count"
const emptyQueueMessage = "empty queue"


type Deque struct {
	data []int
	head  int
	tail  int
	count int
	size  int
}

func NewDeque(size int) *Deque {
	return &Deque{
		data: make([]int, size),
		size: size,
	}
}

func (r *Deque) PushBack(value int) error {
	if r.count == r.size {
		return errors.New(maxCountQueueMessage)
	}

	r.tail = r.next(r.tail)
	r.data[r.tail] = value
	r.count++
	return nil
}

func (r *Deque) PushFront(value int) error {
	if r.count == r.size {
		return errors.New(maxCountQueueMessage)
	}

	r.data[r.head] = value
	r.head = r.prev(r.head)
	r.count++
	return nil
}


func (r *Deque) PopBack() (int, error) {
	if r.count <= 0 {
		return 0, errors.New(emptyQueueMessage)
	}

	result := r.data[r.tail]
	r.tail = r.prev(r.tail)
	r.count--
	return result, nil
}

func (r *Deque) PopFront() (int, error) {
	if r.count <= 0 {
		return 0, errors.New(emptyQueueMessage)
	}

	r.head = r.next(r.head)
	result := r.data[r.head]
	r.count--
	return result, nil
}

func (r *Deque) prev(i int) int {
	if i == 0 {
		return r.size - 1
	}
	return i - 1
}

func (r *Deque) next(i int) int {
	if i == r.size - 1 {
		return 0
	}
	return i + 1
}

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

const PopBack = "pop_back"
const PopFront = "pop_front"
const PushBack = "push_back"
const PushFront = "push_front"
const ErrorMessage = "error"

type Command struct {
	action string
	value *int
}

func readInputData() ([]Command, int) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())

	var commands = make([]Command, count)
	for index:=0; scanner.Scan(); index++ {
		values := strings.Fields(scanner.Text())
		commands[index].action = values[0]
		if len(values) > 1 {
			value, _ := strconv.Atoi(values[1])
			commands[index].value = &value
		}
	}
	_ = inputFile.Close()
	return commands, size
}

func solution(command Command, actions map[string]interface{}) (*int, error) {
	action := actions[command.action]
	if command.value != nil {
		return nil, action.(func(int) error)(*command.value)
	}
	value, err := action.(func()(int, error))()
	return &value, err
}

func main() {
	commands, size := readInputData()
	deque := NewDeque(size)
	var actions = map[string]interface{}{
		PopFront:  deque.PopFront,
		PopBack:   deque.PopBack,
		PushFront: deque.PushFront,
		PushBack:  deque.PushBack,
	}

	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	for _, command := range commands {
		if result, err := solution(command, actions); err != nil {
			_, _ = writer.WriteString(ErrorMessage + fmt.Sprintln())
		} else if result != nil {
			_, _ = writer.WriteString(strconv.Itoa(*result) + fmt.Sprintln())
		}
	}
	_ = writer.Flush()
}

