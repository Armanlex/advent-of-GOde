package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()

	inputLines := strings.Split(input, "\n")

	leftSlice := make([]int, len(inputLines))
	rightSlice := make([]int, len(inputLines))

	for lineIndex, line := range inputLines {

		if line == "" {
			continue
		}

		parts := strings.Split(line, "   ")

		left, err := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])
		if err != nil || err2 != nil {
			fmt.Println("Non number found in input: ", left, right)
			return
		}

		leftSlice[lineIndex] = left
		rightSlice[lineIndex] = right
	}

	rightCounter := [100000]int{}
	for i := 0; i < len(rightSlice); i++ {
		rightCounter[rightSlice[i]]++
	}

	part1Answer := 0
	for i := 0; i < len(leftSlice); i++ {
		part1Answer += leftSlice[i] * rightCounter[leftSlice[i]]
	}

	fmt.Println("Part 2 answer:", part1Answer)
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
