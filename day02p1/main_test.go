package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_Simple(t *testing.T) {
	file, err := os.Open("input_test.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	sum := solve(file)

	if sum != 8 {
		t.Error("incorrect sum: excpected 8 but got ", sum)
	}
}
