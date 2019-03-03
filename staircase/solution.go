package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the staircase function below.
func staircase(n int32) {
	i := int32(0)
	k := int32(0)
	r := n
	c := n
	var buffer bytes.Buffer
	// rows
	for r > 0 {
		r--
		k++
		// line
		i = 0
		for i < c {
			if i >= c-k {
				buffer.WriteString("#")
			} else {
				buffer.WriteString(" ")
			}
			i++
		}
		buffer.WriteString("\n")

	}
	fmt.Println(buffer.String())

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
