package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

const inputFile = "input.txt"
const outputFile = "output.txt"


func input() []string {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func output(lines []string)  {
	file, _ := os.Create(outputFile)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, _ = writer.WriteString(line + " ")
	}
	_ = writer.Flush()
}


func main() {
	//start := time.Now()

	lines := input()
	value, _ := strconv.Atoi(lines[0])
	if value == 2 {
		output([]string {"2"})
		os.Exit(0)
	}
	isContinue := true
	var result []string

	ceil := int(math.Sqrt(float64(value)))
	matrix := GetEratosthenesBooleans(ceil)

	var eratosthenes []int
	for s := 2; s <= ceil; s++ {
		if matrix[s] {
			eratosthenes = append(eratosthenes, s)
		}
	}

	max := len(eratosthenes) - 1
	for isContinue {
		for index, i:= range eratosthenes {

			if i == value {
				result = append(result, strconv.Itoa(i))
				isContinue = false
				break
			}

			if value % i == 0 {
				result = append(result, strconv.Itoa(i))
				value = value/i
				isContinue = true
				break
			}

			if  index == max {
				result = append(result, strconv.Itoa(value))
				isContinue = false
				break
			}

		}
	}
	output(result)

	//elapsed := time.Since(start)
	//log.Printf("Took %s", elapsed)
}


func GetEratosthenesBooleans(ceil int) []bool {
	result := make([]bool, ceil + 1)

	for i := 2; i < ceil+1; i++ {
		result[i] = true
	}

	for value := 2; value*value <= ceil; value++ {
		for i := value * 2; i <= ceil; i += value {
			result[i] = false
		}
	}
	return result
}

