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
	cubeLimitMap := make(map[string]int)
	cubeLimitMap["red"] = 12
	cubeLimitMap["green"] = 13
	cubeLimitMap["blue"] = 14

	var sum int = 0
	scanner := bufio.NewScanner(reader)

	reGameNumber := regexp.MustCompile(`Game\s+(\d+):`)

	for scanner.Scan() {
		var isValidGame bool = true

		line := scanner.Text()

		match := reGameNumber.FindStringSubmatch(line)

		if len(match) == 0 {
			fmt.Println("No match found.")
		}

		gameNumber, err := strconv.Atoi(string(match[1]))
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		count := len(string(match[1]))

		lineWithOutNumber := line[count+7:]

		sets := strings.Split(lineWithOutNumber, "; ")

		for _, set := range sets {
			if isValidGame == false {
				break
			}

			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {
				cubeData := strings.Split(cube, " ")

				if len(cubeData) != 2 {
					fmt.Println("Error cube data.", gameNumber, "len: ", len(cubeData), cube)
				}

				cubeCount := cubeData[0]

				parsedCount, err := strconv.Atoi(string(cubeCount))

				if err != nil {
					fmt.Println("Error: ", err)
					os.Exit(1)
				}

				cubeColor := cubeData[1]

				cubeLimit, isColorExists := cubeLimitMap[cubeColor]

				if isColorExists == false {
					fmt.Println("Color doesnt exist")
					os.Exit(1)
				}

				if parsedCount > cubeLimit {
					isValidGame = false
					break
				}
			}
		}

		if isValidGame {
			sum = sum + gameNumber
		}
	}

	return sum
}
