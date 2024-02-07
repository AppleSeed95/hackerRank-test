package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'ModuloFibonacciSequence' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 */

var ticker = time.NewTicker(10 * time.Millisecond)
var first int = 0
var second int = 0

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	start := 1
	for {
		request := <-requestChan
		if request {
			<-ticker.C
			result := generateFibo(start)
			resultChan <- result
			start++
		}
	}
}
func generateFibo(nth int) int {
	if first == 0 {
		first = 1
		return 1
	}
	if second == 0 {
		second = 2
		return 2
	}
	result := (first + second) % 1000000000
	first, second = second, result
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	skip := int32(skipTemp)

	totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	total := int32(totalTemp)

	resultChan := make(chan int)
	requestChan := make(chan bool)
	go ModuloFibonacciSequence(requestChan, resultChan)
	for i := int32(0); i < skip+total; i++ {
		start := time.Now().UnixNano()
		requestChan <- true
		new := <-resultChan
		if i < skip {
			continue
		}
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000
		if timeDiff < 10 {
			fmt.Println("Rate is too high")
			break
		}
		fmt.Println(new)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return " "
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
