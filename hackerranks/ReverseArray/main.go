package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func fakeSort(i sort.Interface) {
	//0:	1 4 3 2
	//F:    2 3 4 1

	// 0 -> len - 1
	// 1 -> len - 2

	length := i.Len()
	for x := 0; x < length; x++ {
		var indexToSwap int
		if x == 0 {
			indexToSwap = i.Len() - 1
		} else {
			indexToSwap = x - 1
		}
		i.Swap(x, indexToSwap)
	}
}

func reverseArray(a []int) []int {
	var reverse []int
	for x := len(a) - 1; x >= 0; x-- {
		reverse = append(reverse, a[x])
	}
	return reverse
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arrCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := reverseArray(arr)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
