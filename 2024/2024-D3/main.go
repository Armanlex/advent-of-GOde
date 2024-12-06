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
	total := 0
	for _, line := range inputLines {
		parts := strings.Split(line, "mul(")
		newParts := []string{}
		for _, p := range parts {
			newParts = append(newParts, strings.Split(p, ")")...)
		}
		for _, p := range newParts {
			sides := strings.Split(p, ",")
			if len(sides) == 2 {
				val1, err1 := strconv.Atoi(sides[0])
				val2, err2 := strconv.Atoi(sides[1])
				if err1 == nil && err2 == nil {
					total += val1 * val2
				}
			}
		}
	}
	return fmt.Sprint(total)
}

func part2(input string) string {
	inputLines := strings.Split(input, "\n")
	total := 0
	do := true
	for _, line := range inputLines {
		parts := strings.Split(line, "mul(")

		newParts := []string{}
		for _, p := range parts {
			temp := strings.Split(p, ")")
			for _, t := range temp {
				newParts = append(newParts, t+")")
			}
		}
		for _, p := range newParts {
			startI := strings.LastIndex(p, "do()")
			stopI := strings.LastIndex(p, "don't()")
			if startI > stopI {
				do = true
			} else if stopI > startI {
				do = false
			}
			sides := strings.Split(p, ",")
			if len(sides) == 2 {
				a := sides[0]
				b := sides[1]
				if b[len(b)-1] != ')' {
					continue
				}
				val1, err1 := strconv.Atoi(a)
				val2, err2 := strconv.Atoi(b[:len(b)-1])
				if do && err1 == nil && err2 == nil {
					total += val1 * val2
				}
			}
		}
	}

	return fmt.Sprint(total)
}

func getInput() string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println(fileName, " file is empty")
		os.Exit(1)
	}
	input := strings.ReplaceAll(string(data), "\r\n", "\n") //doing this replace so it can handle both linux and window text format
	return strings.TrimSpace(input)                         //doing this cause usually there's an extra new line at the bottom of the input
}
