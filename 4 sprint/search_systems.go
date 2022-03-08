package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"
const OutputFileName = "output.txt"
const Separator = " "
const PrimeNumber = 1000000007
const MaxValues = 5

type HashSum uint64
type Position uint16
type Count map[Position]uint16

type OutputData struct {
	position int
	count    uint16
}

type InputData struct {
	documents      map[HashSum]Count
	requests       [][]HashSum
	requestCount   int
	documentsCount int
}

func NewInputData(documentsCount, requestCount int) InputData {
	return InputData{
		documents:      make(map[HashSum]Count),
		requests:       make([][]HashSum, requestCount),
		requestCount:   requestCount,
		documentsCount: documentsCount,
	}
}

func hash(word string) HashSum {
	var result uint64
	for _, s := range word {
		// Решил не добавлять % mod т.к. помещаемся в uint64 и не нужно разрешать коллизии
		result = (result + uint64(s)) * PrimeNumber
	}
	return HashSum(result)
}

func (d *InputData) addDocument(index int, data string) {
	position := Position(index)
	words := strings.Fields(data)
	for _, word := range words {
		wordHash := hash(word)
		//  Записываем только те значения, которые нужны для запросов
		if _, ok := d.documents[wordHash]; ok {
			value, _ := d.documents[wordHash][position]
			d.documents[wordHash][position] = value + 1
		}
	}
}

func (d *InputData) addRequest(index int, data string) {
	position := Position(index)
	words := strings.Fields(data)
	request := make(map[HashSum]struct{}, len(words))
	for _, word := range words {
		wordHash := hash(word)
		if _, ok := request[wordHash]; !ok {
			// Заранее указываем нужный хэш для документов
			d.documents[wordHash] = make(Count)
			request[wordHash] = struct{}{}
		}
	}

	keys := make([]HashSum, 0, len(request))
	for k, _ := range request {
		keys = append(keys, k)
	}

	d.requests[position] = keys
}

func (d *InputData) solution(request int) (data []OutputData) {
	positions := make([]OutputData, d.documentsCount)
	for j := 0; j < d.documentsCount; j++ {
		positions[j].position = j + 1
	}

	for _, hashSum := range d.requests[request] {
		if document, ok := d.documents[hashSum]; ok {
			for position, count := range document {
				positions[position].count += count
			}
		}
	}

	sort.Slice(positions, func(first, second int) bool {
		if positions[first].count == positions[second].count {
			return positions[first].position < positions[second].position
		}
		return positions[first].count >= positions[second].count
	})

	for i := 0; i < len(positions) && i < MaxValues; i++ {
		value := positions[i]
		if value.count > 0 {
			data = append(data, value)
		}
	}
	return
}

func readInputData() (inputData InputData) {
	inputFile, _ := os.Open(InputFileName)
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	documentCount, _ := strconv.Atoi(scanner.Text())
	for line := 0; line < documentCount; line++ {
		scanner.Scan()
	}

	scanner.Scan()
	requestsCount, _ := strconv.Atoi(scanner.Text())
	inputData = NewInputData(documentCount, requestsCount)
	for index := 0; index < requestsCount; index++ {
		scanner.Scan()
		inputData.addRequest(index, scanner.Text())
	}
	_ = inputFile.Close()

	inputFile, _ = os.Open(InputFileName)
	scanner = bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	for index := 0; index < documentCount; index++ {
		scanner.Scan()
		inputData.addDocument(index, scanner.Text())
	}
	_ = inputFile.Close()

	return
}

func main() {
	data := readInputData()

	outputFile, _ := os.Create(OutputFileName)
	writer := bufio.NewWriter(outputFile)
	for i := 0; i < len(data.requests); i++ {
		for _, v := range data.solution(i) {
			_, _ = writer.WriteString(strconv.Itoa(v.position) + Separator)
		}
		_, _ = writer.WriteString(fmt.Sprintln())
	}
	_ = writer.Flush()
}
