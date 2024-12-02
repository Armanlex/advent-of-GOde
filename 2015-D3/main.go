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
	theMap := make(map[string]int, 8000)
	x := 0
	y := 0
	theMap[fmt.Sprint(x, y)]++
	for _, r := range input {
		switch r {
		case '>':
			x++
		case '<':
			x--
		case 'v':
			y++
		case '^':
			y--
		}
		theMap[fmt.Sprint(x, y)]++
	}
	return fmt.Sprint(len(theMap))
}

func part2(input string) string {
	theMap := make(map[string]int, 8000)
	santaX := 0
	santaY := 0
	roboX := 0
	roboY := 0
	theMap[fmt.Sprint(santaX, santaY)]++
	for i, r := range input {
		if i%2 == 1 {
			switch r {
			case '>':
				santaX++
			case '<':
				santaX--
			case 'v':
				santaY++
			case '^':
				santaY--
			}
			theMap[fmt.Sprint(santaX, santaY)]++
		} else {
			switch r {
			case '>':
				roboX++
			case '<':
				roboX--
			case 'v':
				roboY++
			case '^':
				roboY--
			}
			theMap[fmt.Sprint(roboX, roboY)]++
		}

	}

	return fmt.Sprint(len(theMap))

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
