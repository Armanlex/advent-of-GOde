package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	fmt.Println("Part 1 answer:", part1(input))
	fmt.Println("Part 2 answer:", part2(input))
}

func part1(input string) string {
	inputLines := strings.Split(input, "\n")
	totalPaper := 0
	for i := 0; i < len(inputLines); i++ {
		if inputLines[i] == "" {
			continue
		}

		parts := strings.Split(inputLines[i], "x")
		length, err1 := strconv.Atoi(parts[0])
		width, err2 := strconv.Atoi(parts[1])
		height, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			panic("not a number in atoi!")
		}

		totalPaper += 2*length*width + 2*width*height + 2*height*length
		smallest := min(length*width, width*height, height*length)
		totalPaper += smallest

	}

	return fmt.Sprint(totalPaper)

}

func part2(input string) string {
	inputLines := strings.Split(input, "\n")
	totalRibbon := 0
	for i := 0; i < len(inputLines); i++ {
		if inputLines[i] == "" {
			continue
		}

		parts := strings.Split(inputLines[i], "x")
		length, err1 := strconv.Atoi(parts[0])
		width, err2 := strconv.Atoi(parts[1])
		height, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			panic("not a number in atoi!")
		}

		totalRibbon += min(length+width, width+height, height+length) * 2
		totalRibbon += length * width * height
	}

	return fmt.Sprint(totalRibbon)
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
