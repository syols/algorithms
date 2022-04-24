package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

type Pos struct{
	iIndex int
	jIndex int
}


func readInputData() ([]int, []int) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	scanner.Scan()
	first:= scanner.Text()
	f := strings.Split(first, " ")
	scanner.Scan()
	scanner.Scan()
	second := scanner.Text()
	s := strings.Split(second, " ")

	a := make([]int, len(f))
	for i, v := range f {
		a[i], _ = strconv.Atoi(v)
	}

	b := make([]int, len(s))
	for i, v := range s {
		b[i], _ = strconv.Atoi(v)
	}
	return a, b
}

func solution(first, second []int) (int, []Pos) {
	lengthFirst, lenSecond := len(first), len(second)
	result := fillLine(lengthFirst, lenSecond)
	for k := 0; k <= lengthFirst; k++ {
		for m := 0; m <= lenSecond; m++ {
			if k == 0 || m == 0 {
				result[k][m] = 0
				continue
			}

			if first[k-1] == second[m-1] {
				result[k][m] = result[k-1][m-1] + 1
				continue
			}
			
			result[k][m] = int(math.Max(float64(result[k-1][m]), float64(result[k][m-1])))
		}
	}
	count := result[lengthFirst][lenSecond]
	pos := positions(result, lengthFirst-1, lenSecond-1, first, second)
	return count, pos
}

func fillLine(lengthFirst int, lenSecond int) [][]int {
	result := make([][]int, lengthFirst+1)
	for i := range result {
		result[i] = make([]int, lenSecond+1)
	}
	return result
}


func positions(result [][]int, lengthFirst, lengthSecond int, first, second []int) []Pos {
	if lengthFirst < 0 || lengthSecond < 0 {
		return []Pos{}
	}

	if first[lengthFirst] == second[lengthSecond] {
		return append(positions(result, lengthFirst-1, lengthSecond-1,first, second), Pos{iIndex: lengthFirst, jIndex: lengthSecond})
	}

	if result[lengthFirst+1][lengthSecond] <= result[lengthFirst][lengthSecond+1] {
		return positions(result, lengthFirst-1, lengthSecond, first, second)
	}

	return positions(result, lengthFirst, lengthSecond-1, first, second)
}

func main() {
	capacity, array := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	count, b := solution(capacity, array)
	_, _ = writer.WriteString(strconv.Itoa(count) + "\n")
	s, f := "", ""
	for _, e := range b {
		s += strconv.Itoa(e.iIndex+ 1) + " "
		f += strconv.Itoa(e.jIndex+ 1) + " "
	}
	_, _ = writer.WriteString(s + "\n")
	_, _ = writer.WriteString(f + "\n")
	_ = writer.Flush()
}
