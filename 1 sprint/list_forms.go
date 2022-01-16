package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

const inputFile = "input.txt"
const outputFile = "output.txt"


func input() []int {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	var result []int
	var mas []rune
	for scanner.Scan() {
		value := scanner.Text()
		if value == "\n" {
			v, _ :=strconv.Atoi(string(mas))
			result = append(result, v)
			mas = nil
			continue
		}
		for _, value := range value {
			if unicode.IsNumber(value) {
				mas = append(mas, value)
				continue
			}
		}
	}
	v, _ :=strconv.Atoi(string(mas))
	result = append(result, v)
	return result
}

func output(lines []rune)  {
	file, _ := os.Create(outputFile)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, _ = writer.WriteString(string(line) + " ")
	}
	_ = writer.Flush()
}


func main() {
	intArrays := input()
	result := strconv.Itoa(intArrays[1] + intArrays[2])
	var runes []rune
	for _, rune := range result {
		runes = append(runes, rune)
	}
	output(runes)
}

