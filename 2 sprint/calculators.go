package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

const Addition = "+"
const Subtraction = "-"
const Division = "/"
const Multiplication = "*"

func add(a int, b int) int{
	return a + b;
}

func multiply(a int, b int) int{
	return a * b;
}

func divide(a int, b int) int{
	return int(math.Floor(float64(a) / float64(b)));
}

func substract(a int, b int) int{
	return a - b;
}


func solution(commands []string) int {
	var stack []int
	var actions = map[string]interface{} {
		Addition: add,
		Subtraction: substract,
		Division: divide,
		Multiplication: multiply,
	}

	for _, command := range commands {
		tail := len(stack) - 1
		if value, err := strconv.Atoi(command); err != nil {
			action := actions[command]
			stack[tail - 1] = action.(func(int, int) int)(stack[tail - 1], stack[tail])
			stack = stack[:tail]
		} else {
			stack = append(stack, value)
		}
	}
	return stack[len(stack)-1]
}

func readInputData() []string {
	file, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	_ = file.Close()
	return result
}


func main() {
	file, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(file)
	result := solution(readInputData())
	_, _ = writer.WriteString(strconv.Itoa(result))
	_ = writer.Flush()
}

