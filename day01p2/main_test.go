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

	if sum != 281 {
		t.Error("incorrect sym: excpected 281 but got ", sum)
	}
}
