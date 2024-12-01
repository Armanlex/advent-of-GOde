package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := getInput()
	inputLines := strings.Split(input, "\n")

	part1Answer := ""
	fmt.Println("Part 1 answer:", part1Answer)
}

func getInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println("Input.txt file is empty")
		os.Exit(1)
	}
	input := strings.ReplaceAll(string(data), "\r\n", "\n") //doing this replace so it can handle both linux and window text format

	return input
}
