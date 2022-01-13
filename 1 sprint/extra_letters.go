package main

import (
	"bufio"
	"os"
	"unicode"
)

const inputFile = "input.txt"
const outputFile = "output.txt"



func output(lines []rune)  {
	file, _ := os.Create(outputFile)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, _ = writer.WriteString(string(line) + " ")
	}
	_ = writer.Flush()
}


func main() {
	first := map[rune]int{}
	second := map[rune]int{}
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	pos := 0
	for scanner.Scan() {
		value := scanner.Text()
		if value == "\n" {
			pos++
			continue
		}
		if pos >= 2 {
			break
		}
		for _, value := range value {
			if unicode.IsLetter(value) {

				if pos == 0 {
					count, _ := first[value]
					first[value] = count + 1
				}
				if pos == 1 {
					count, _ := second[value]
					second[value] = count + 1
				}
				continue
			}
		}
	}

	for k, v := range second {
		v1, _ := first[k]
		if v != v1 {
			output([]rune(string(k)))
			break
		}
	}
}

