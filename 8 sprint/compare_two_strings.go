package main

import (
	"bufio"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

func readInputData() (first string, second string) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	first = scanner.Text()
	scanner.Scan()
	second = scanner.Text()
	return
}

func solution(pass, base string) int {
	for i, j := 0, 0; i < len(pass) || j < len(base); {
		first := rune(0)
		if i < len(pass) {
			first = rune(pass[i])
			if (first - 'a') % 2 == 0 {
				i++
				continue
			}
		}

		second := rune(0)
		if j < len(base) {
			second = rune(base[j])
			if (second-'a')%2 == 0 {
				j++
				continue
			}
		}

		if first - second > 0 {
			return 1
		} else if first - second < 0{
			return -1
		}
		i++
		j++
	}
	return 0
}

func main() {
	first, second := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	distance := solution(first, second)
	_, _ = writer.WriteString(strconv.Itoa(distance))
	_ = writer.Flush()
}
