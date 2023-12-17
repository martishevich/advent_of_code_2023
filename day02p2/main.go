package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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
	cubeMax := make(map[string]int)

	var sum int = 0
	scanner := bufio.NewScanner(reader)

	reGameNumber := regexp.MustCompile(`Game\s+(\d+):`)

	for scanner.Scan() {
		cubeMax["red"] = 1
		cubeMax["green"] = 1
		cubeMax["blue"] = 1

		line := scanner.Text()

		match := reGameNumber.FindStringSubmatch(line)

		if len(match) == 0 {
			fmt.Println("No match found.")
		}

		count := len(string(match[1]))

		lineWithOutNumber := line[count+7:]

		sets := strings.Split(lineWithOutNumber, "; ")

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {
				cubeData := strings.Split(cube, " ")

				if len(cubeData) != 2 {
					fmt.Println("Error cube data. ", "len: ", len(cubeData), cube)
				}

				cubeCount := cubeData[0]

				parsedCount, err := strconv.Atoi(string(cubeCount))

				if err != nil {
					fmt.Println("Error: ", err)
					os.Exit(1)
				}

				cubeColor := cubeData[1]

				if cubeMax[cubeColor] < parsedCount {
					cubeMax[cubeColor] = parsedCount
				}
			}
		}

		sum = sum + cubeMax["red"]*cubeMax["green"]*cubeMax["blue"]
	}

	return sum
}
