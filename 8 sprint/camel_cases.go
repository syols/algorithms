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
const YesMessage = "YES"
const NoMessage = "NO"

type Words []string
type Trie struct {
	children map[rune]*Trie
	isTerminal bool
	index []int
}

func (h *Trie) add(index int, word string) {
	head := h
	for _, r := range word {
		if unicode.IsUpper(r) != true {
			continue
		}

		if _, isFound := head.children[r]; !isFound {
			trie := NewTrie()
			head.children[r] = &trie
		}
		head = &*head.children[r]
		head.index = append(head.index, index)
	}
	head.isTerminal = true
	return
}

func NewTrie() Trie {
	return Trie{
		children: map[rune]*Trie{},
		isTerminal: false,
		index: []int{},
	}
}

func readInputData() (result Words) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	trie := NewTrie()
	words := make(Words, count)
	for i:=0; i < count; i++ {
		scanner.Scan()
		word := scanner.Text()
		trie.add(i, word)
		words[i] = word
	}

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	for i:=0; i < c; i++ {
		scanner.Scan()
		r := trie.solution(scanner.Text(), count)
		s := []string{}
		for _, v := range r {
				s = append(s, words[v])
		}
		sort.Strings(s)
		result = append(result, s...)
		if len(r) == 0 {
			result = append(result, " ")
		}
	}
	return result
}

func (h *Trie) solution(text string, count int) []int {
	current := h
	if text == "" {
		l := make([]int, count)
		for i:=0; i < count; i++ {
			l[i] = i
		}
		return l
	}

	i:=0
	for i=0; i < len(text); i++ {
		r := rune(text[i])
		if _, isFound := current.children[r]; isFound {
			current = current.children[r]
		} else if len(current.children) == 0 {
			return []int{}
		} else {
			break
		}
	}
	if i == len(text) {
		return current.index
	}
	return []int{}
}

func main() {
	words := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	for _, word := range words {
		_, _ = writer.WriteString(word + "\n")
	}
	_ = writer.Flush()
}
