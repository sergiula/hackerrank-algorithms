package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	// посчитаем длину массива
	arrLenght := float32(len(arr))
	// так как мы не знаем отсортирован ли он - то придется перебрать
	// в цикле
	var max, min, equal float32

	for _, v := range arr {
		if v > 0 {
			max++
		}
		// посчитаем количество элементов более 0
		// посчитаем количество элементов менее 09
		if v < 0 {
			min++
		}
		// посчитаем количество элементов равное 0
		if v == 0 {
			equal++
		}
	}

	// получить соотношения (6 знаков после запятой)
	// вывести соотношение количества позитивных кейсов
	fmt.Println(max / arrLenght)
	// вывести соотношение негативных кейсов
	fmt.Println(min / arrLenght)
	// вывести количество кейсов с 0
	fmt.Println(equal / arrLenght)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
