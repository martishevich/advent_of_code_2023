package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_Simple(t *testing.T) {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	sum := solve(file)

	if sum != 142 {
		t.Error("incorrect sym: excpected 142 but got ", sum)
	}
}
