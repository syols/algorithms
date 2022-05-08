package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"unicode"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const ClosedQuote = ']'
const Zero = '0'

type Nested struct {
	multiplier int
	text       []rune
}

func NewNested(multiple int) Nested {
	return Nested{
		multiplier: multiple,
		text:       []rune{},
	}
}

func (n *Nested) add(value *Nested) {
	(*n).text = append((*n).text, value.multiply()...)
}

func (n *Nested) multiply() (result []rune) {
	for i := 0; i < (*n).multiplier; i++ {
		result = append(result, (*n).text...)
	}
	return
}

func solution(array []string) (prefix []rune) {
	sort.Slice(array, func(f, s int) bool {
		return len(array[f]) < len(array[s])
	})

	prefix = unpack(array[0])
	if len(array) > 0 {
		for _, value := range array[1:] {
			text, prefixLength := unpack(value), len(prefix)
			for k, v := range text {
				if k >= prefixLength || v != prefix[k] {
					prefix = prefix[0:k]
					break
				}
			}
		}
	}
	return
}

func unpack(value string) []rune {
	stack := []Nested{{}}
	for i := 0; i < len(value); i++ {
		r := rune(value[i])
		currentIndex := len(stack) - 1
		if unicode.IsDigit(r) {
			stack = append(stack, NewNested(int(r-Zero)))
			i++
		} else if r == ClosedQuote {
			stack[currentIndex-1].add(&stack[currentIndex])
			stack = stack[0:currentIndex]
		} else {
			stack[currentIndex].text = append(stack[currentIndex].text, r)
		}
	}
	return stack[0].text
}

func readInputData() (result []string) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	result = make([]string, count)
	for i := count - 1; i >= 0; i-- {
		scanner.Scan()
		result[i] = scanner.Text()
	}
	return
}

func main() {
	array := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	result := solution(array)
	_, _ = writer.WriteString(string(result))
	_ = writer.Flush()
}
