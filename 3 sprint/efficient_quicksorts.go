package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

type Competitors []Competitor

type Competitor struct {
	login     string
	completed int
	penalty   int
}

func (c Competitors) Sort() {
	length := len(c)
	if length > 1 {
		c.Splitting(0, length - 1)
	}
}

func (c Competitors) Splitting(leftIndex int, rightIndex int) {
	pivot := c.pivot(leftIndex, rightIndex)
	left, right := leftIndex, rightIndex
	for left <= right {
		for ; c.Less(&pivot, &c[left]); left++ {}
		for ; c.Less(&c[right], &pivot); right-- {}

		if left <= right {
			c[right], c[left] = c[left], c[right]
			right--
			left++
		}
	}

	if leftIndex < right {
		c.Splitting(leftIndex, right)
	}

	if rightIndex > left {
		c.Splitting(left, rightIndex)
	}

}

func (c Competitors) pivot(leftIndex int, rightIndex int) Competitor {
	return c[rand.Intn(rightIndex-leftIndex) + leftIndex]
}

func (c Competitors) Less(first *Competitor, second *Competitor) bool {
	if first.completed == second.completed {
		if first.penalty == second.penalty {
			return first.login > second.login
		}
		return first.penalty > second.penalty
	}
	return first.completed < second.completed
}

func readInputData() Competitors {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	competitors := make(Competitors, size)

	for index:=0; index < size; index++ {
		scanner.Scan()
		values := strings.Fields(scanner.Text())
		competitors[index].login = values[0]
		competitors[index].completed, _ = strconv.Atoi(values[1])
		competitors[index].penalty, _ = strconv.Atoi(values[2])
	}
	_ = inputFile.Close()
	return competitors
}

func writeData(competitors Competitors) {
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	for _, competitor := range competitors {
		_, _ = writer.WriteString(competitor.login + fmt.Sprintln())
	}
	_ = writer.Flush()
}

func main() {
	competitors := readInputData()
	competitors.Sort()
	writeData(competitors)
}

