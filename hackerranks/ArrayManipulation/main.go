package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"math"
)

func arrayManipulation(arraySize int64, queries [][]int64) int64 {
	//array := NewOneIndexedArray(int(n))
	//for _, v := range queries {
	//    array.alter(int(v[0]), int(v[1]), v[2])
	//}
	//
	//return array.maxValue()
	var maxValue int64
	array := make([]int64, arraySize)
	for _,v:= range queries {
		alterRange(v, array)
	}
	var sum int64
	for _, v := range array {
		sum += v
		maxValue = int64(math.Max(float64(maxValue), float64(sum)))
	}
	return maxValue

}

func alterRange(v []int64, array []int64) {
	start := v[0] - 1
	end := v[1]
	value := v[2]
	array[start] += value
	if end < int64(len(array)) {
		array[end] -= value
	}
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int64(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int64(mTemp)

	var queries [][]int64
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int64
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int64(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
