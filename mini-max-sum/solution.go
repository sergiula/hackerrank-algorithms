package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int64) {
	var sum int64
	var max int64
	var min int64

	for _, n := range arr {
		if max == 0 {
			max = n
		}
		if min == 0 {
			min = n
		}

		if n < min {
			min = n
		}

		if max < n {
			max = n
		}
		sum += n
	}
	// min sum is left size digits

	// max sum - right size digits

	fmt.Printf("%d %d\n", sum-max, sum-min)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int64

	for i := 0; i < 5; i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		// arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
		// arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
