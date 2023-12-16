package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type ReplaceStruct struct {
	StringToReplace     string
	IntegerForReplacing string
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	sum := solve(file)

	fmt.Println("sum: ", sum)
}

func solve(reader io.Reader) int {
	replaceArr := []ReplaceStruct{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
	}

	var sum int = 0
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println("line before: ", line)
		for {
			minIndex, indexArr := getMinOccurance(line, replaceArr)

			if minIndex == -1 {
				break
			}

			line = strings.Replace(line, replaceArr[indexArr].StringToReplace, replaceArr[indexArr].IntegerForReplacing, 1)
		}

		fmt.Println("line after: ", line)

		var firstFound = false
		var firstInt int

		var lastFound = false
		var lastInt int

		for _, c := range line {

			isInt := unicode.IsDigit(rune(c))

			if isInt == false {
				continue
			}

			if firstFound == false {
				firstFound = true
				parsedInt, err := strconv.Atoi(string(c))

				if err != nil {
					fmt.Println("Error: ", err)
					os.Exit(1)
				}

				firstInt = parsedInt

				continue
			}

			parsedInt, err := strconv.Atoi(string(c))

			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}

			lastInt = parsedInt
			lastFound = true
		}

		if lastFound == false {
			lastInt = firstInt
		}

		lineResult := firstInt*10 + lastInt
		sum = sum + lineResult
	}

	return sum
}

func getMinOccurance(line string, replaceArr []ReplaceStruct) (int, int) {
	var minIndex int = -1
	var indexInArr = 0

	for i, rs := range replaceArr {
		returnedIndex := strings.Index(line, rs.StringToReplace)

		if returnedIndex != -1 {
			if minIndex == -1 {
				minIndex = returnedIndex
				indexInArr = i
			}

			if returnedIndex < minIndex {
				minIndex = returnedIndex
				indexInArr = i
			}
		}
	}

	return minIndex, indexInArr
}
