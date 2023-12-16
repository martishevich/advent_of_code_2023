package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
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

	for scanner.Scan() {
		var firstFound = false
		var firstInt int

		var lastFound = false
		var lastInt int

		for _, c := range scanner.Text() {
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
