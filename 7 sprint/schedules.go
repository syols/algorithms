package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"

type Task struct {
	begin int
	end int
}

func (a *Task) less(b *Task) bool{
	if b.end == a.end {
		return a.begin < b.begin
	}
	return a.end < b.end
}

func (a *Task) ToString() string{
	return ToString(a.begin) + " " + ToString(a.end) + "\n"
}

func readInputData() (result []Task) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	result = make([]Task, count)
	for i := 0; i < count; i++ {
		scanner.Scan()
		begin := ToMinutes(scanner.Text())
		scanner.Scan()
		end := ToMinutes(scanner.Text())
		result[i] = Task{ begin: begin, end: end}
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].less(&result[j])
	})
	return
}

func ToMinutes(value string) int {
	values := strings.Split(value, ".")
	result, _ := strconv.Atoi(values[0])
	result *= 60
	if len(values) > 1 {
		minutes, _ := strconv.Atoi(values[1])
		return result + minutes
	}
	return result
}

func ToString(value int) string{
	result := strconv.Itoa(value / 60)
	if value % 60 > 0 {
		result += "." + strconv.Itoa(value % 60)
	}
	return result
}

func solution(array []Task) (count int, tasks []Task) {
	currentTime := 0
	for _, task := range array{
		if task.begin >= currentTime {
			tasks = append(tasks, task)
			currentTime = task.end
			count++
		}
	}
	return
}

func main() {
	array := readInputData()
	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	count, tasks := solution(array)
	_, _ = writer.WriteString(strconv.Itoa(count) + "\n")
	for _, task := range tasks{
		_, _ = writer.WriteString(task.ToString())
	}
	_ = writer.Flush()
}
