package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := getInput()
	fmt.Println("Part 1 answer:", part1(input))
	fmt.Println("Part 2 answer:", part2(input))
}

func part1(input string) string {
	floor := 0
	for _, Rune := range input {
		if Rune == ')' {
			floor--
		} else {
			floor++
		}
	}
	return fmt.Sprint(floor)
}

func part2(input string) string {
	floor := 0
	for i, Rune := range input {
		if Rune == ')' {
			floor--
		} else {
			floor++
		}
		if floor == -1 {
			return fmt.Sprint(i + 1)
		}
	}
	return "something went wrong"
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
