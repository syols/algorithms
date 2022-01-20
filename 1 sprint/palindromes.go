package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

const inputFile = "input.txt"
const outputFile = "output.txt"

func input() []string {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	var mas []int32
	count := 0
	for scanner.Scan() {
		value := scanner.Text()
		if value == "\n" {
			continue
		}
		value = strings.ToLower(value)
		for _, value := range value {
			if unicode.IsLetter(value) || unicode.IsNumber(value) {
				mas = append(mas, value)
				count++
				continue
			}
		}
	}
	for index, val := range(mas) {
		if index > count / 2 {
			break
		}
		rindex := count - index - 1
		if val != mas[rindex] {
			return []string{"False"}
		}
	}
	_ = file.Close()
	return []string{"True"}
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
