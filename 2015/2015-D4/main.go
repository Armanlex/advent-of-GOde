package main

import (
	"crypto/md5"
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
	num := 0
	for {
		if HexStartsWithFiveZeroes(input+fmt.Sprint(num), "00000") {
			return fmt.Sprint(fmt.Sprint(num))
		}
		num++
	}
}

func part2(input string) string {
	num := 0
	for {
		if HexStartsWithFiveZeroes(input+fmt.Sprint(num), "000000") {
			return fmt.Sprint(fmt.Sprint(num))
		}
		num++
	}
}

func HexStartsWithFiveZeroes(str, prefix string) bool {
	hex := md5.Sum([]byte(str))
	return strings.HasPrefix(fmt.Sprintf("%.2x%.2x%.2x", hex[0], hex[1], hex[2]), prefix)
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
