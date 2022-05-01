package main

import (
	"bufio"
	"os"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const YesMessage = "YES"
const NoMessage = "NO"

type Trie struct {
	value rune
	children map[rune]*Trie
	isTerminal bool
}

func (h *Trie) add(word string) {
	head := h
	for _, r := range word {
		if _, isFound := head.children[r]; !isFound {
			trie := NewTrie(r)
			head.children[r] = &trie
		}
		head = &*head.children[r]
	}
	head.isTerminal = true
}

func NewTrie(value rune) Trie {
	return Trie{
		value:    value,
		children: map[rune]*Trie{},
		isTerminal: false,
	}
}

func readInputData() (text string, trie Trie) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	// Выбрал язык на котором не работал и набил шишек (:
	// Go крут, хоть я ничего и не применял из языка толком
	// Заранее спасибо за ревью, все было круто!

	// Для меня задача была не про алгоритмы, она была про этот бу
	// File is too long to be displayed fully
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	scanner.Scan()
	text = scanner.Text()

	head := NewTrie(-1)
	scanner.Scan()
	for scanner.Scan() {
		word := scanner.Text()
		head.add(word)
	}

	return text, head
}

func (h *Trie) solution(text string) bool {
	length := len(text)
	dp := make([]bool, length+1)
	dp[0] = true // В слаке немного "проспойлерили", что сюда надо добавить динамику, после этого все стало понятно, сама динамика очень легкая
	for pos:=0; pos <= length; pos++{
		if !dp[pos] { // Пропускаем "нетерминальные" позиции
			continue
		}

		current := h
		for offset := pos; offset < length; offset++{ // Тут бежим по дереву и проставляем все терминальные символы в dp == true, с них мы будем далее искать
			r := rune(text[offset])
			if _, isFound := current.children[r]; isFound {
				current = current.children[r]
				if current.isTerminal {
					dp[offset + 1] = true
				}
			} else {
				break
			}
		}
	}
	return dp[length] // Ну и по классике 7 спринта, мы должны дойти через чреду true до последнего терминального символа.
	// Кажется тут можно сделать без массива dp, а просто хранить массив терминальных позиций в тексте и последний должен == length
}

func main() {
	text, trie := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	if trie.solution(text) {
		_, _ = writer.WriteString(YesMessage)
	} else {
		_, _ = writer.WriteString(NoMessage)
	}
	_ = writer.Flush()
}
