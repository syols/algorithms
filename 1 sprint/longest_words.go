package main

import (
	"bufio"
	"os"
	"strconv"
)

const inputFile = "input.txt"
const outputFile = "output.txt"

func input() []string {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var maxWord string
	var maxCount int
	step := 0
	for scanner.Scan() {
		step++
		if step <= 1 {
			continue
		}
		word := scanner.Text()
		count := len(word)
		if maxCount < count {
			maxCount = count
			maxWord = word
		}
	}
	return []string{maxWord, strconv.Itoa(maxCount)}
}

func output(values []string)  {
	file, _ := os.Create(outputFile)
	writer := bufio.NewWriter(file)
	for _, value := range values {
		_, _ = writer.WriteString(value + "\n")
	}
	_ = writer.Flush()
}



func main() {
	output(input())
}
