package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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
	var sum int = 0
	scanner := bufio.NewScanner(reader)

	var dynamicArray [][]string

	for scanner.Scan() {
		line := scanner.Text()

		var lineArray []string
		for _, char := range line {
			lineArray = append(lineArray, string(char))
		}

		dynamicArray = append(dynamicArray, lineArray)
	}

	for lineIndex, line := range dynamicArray {
		lineLen := len(line) - 1

		isNumberStarted := false
		elementStartIndex := 0
		elementEndIndex := 0

		for elementIndex, element := range line {
			fmt.Println("element: ", element)
			_, err := strconv.Atoi(element)
			if err == nil {
				if isNumberStarted && elementIndex != lineLen {
					elementEndIndex = elementIndex
					continue
				}

				if isNumberStarted && elementIndex == lineLen {
					elementEndIndex = elementIndex
					isAround := checkAround(dynamicArray, lineIndex, elementStartIndex, elementEndIndex)
					if isAround {
						number := getNumber(dynamicArray[lineIndex], elementStartIndex, elementEndIndex)

						sum = sum + number
					}
					continue
				}

				if !isNumberStarted {
					isNumberStarted = true
					elementStartIndex = elementIndex
					elementEndIndex = elementIndex
					continue
				}
			}

			if isNumberStarted {
				isAround := checkAround(dynamicArray, lineIndex, elementStartIndex, elementEndIndex)
				if isAround {
					number := getNumber(dynamicArray[lineIndex], elementStartIndex, elementEndIndex)

					sum = sum + number
				}

				isNumberStarted = false
			}
		}
	}

	return sum
}

func getNumber(line []string, elementStartIndex int, elementEndIndex int) int {
	var builder strings.Builder

	for i := elementStartIndex; i <= elementEndIndex; i++ {
		builder.WriteString(line[i])
	}
	newString := builder.String()

	parsedInt, err := strconv.Atoi(newString)

	if err != nil {
		fmt.Println(err)
	}

	return parsedInt
}

func checkAround(dynamicArray [][]string, lineIndex int, elementStartIndex int, elementEndIndex int) bool {
	lineLen := len(dynamicArray[lineIndex])

	if lineIndex != 0 {
		if elementStartIndex != 0 {
			if checkElement(dynamicArray[lineIndex-1][elementStartIndex-1]) {
				return true
			}
		}

		if elementEndIndex != lineLen-1 {
			if checkElement(dynamicArray[lineIndex-1][elementEndIndex+1]) {
				return true
			}
		}

		for i := elementStartIndex; i <= elementEndIndex; i++ {
			if checkElement(dynamicArray[lineIndex-1][i]) {
				return true
			}
		}
	}

	if lineIndex != len(dynamicArray)-1 {
		if elementStartIndex != 0 {
			if checkElement(dynamicArray[lineIndex+1][elementStartIndex-1]) {
				return true
			}
		}

		if elementEndIndex != lineLen-1 {
			if checkElement(dynamicArray[lineIndex+1][elementEndIndex+1]) {
				return true
			}
		}

		for i := elementStartIndex; i <= elementEndIndex; i++ {
			if checkElement(dynamicArray[lineIndex+1][i]) {
				return true
			}
		}
	}

	if elementStartIndex != 0 {
		if checkElement(dynamicArray[lineIndex][elementStartIndex-1]) {
			return true
		}
	}

	if elementEndIndex != lineLen-1 {
		if checkElement(dynamicArray[lineIndex][elementEndIndex+1]) {
			return true
		}
	}

	return false
}

func checkElement(element string) bool {
	_, err := strconv.Atoi(element)

	if err != nil && element != "." {
		return true
	}

	return false
}
