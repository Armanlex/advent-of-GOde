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
	safeLevels := 0
	for _, line := range strings.Split(input, "\n") {
		numSlice := line2NumSlice(line)
		if levelIsSafe(numSlice) {
			safeLevels++
		}
	}
	return fmt.Sprint(safeLevels)
}

func part2(input string) string {
	safeLevels := 0
	for _, line := range strings.Split(input, "\n") {
		numSlice := line2NumSlice(line)
		if levelIsSafeOneProblemAllowed(numSlice) {
			safeLevels++
		}
	}

	return fmt.Sprint(safeLevels)
}

func levelIsSafe(slice []int) bool {
	aSmallerThanB := slice[0] < slice[1]
	if pairIsBad(slice[0], slice[1], aSmallerThanB) {
		return false
	}
	for i := 1; i < len(slice)-1; i++ {
		if pairIsBad(slice[i], slice[i+1], aSmallerThanB) {
			return false
		}
	}
	return true
}

func levelIsSafeOneProblemAllowed(slice []int) bool {
	aLessThanBCounter := 0
	for i := 0; i < 4; i++ {
		if slice[i] < slice[i+1] {
			aLessThanBCounter++
		} else {
			aLessThanBCounter--
		}
	}
	aSmallerThanB := true
	if aLessThanBCounter < 0 {
		aSmallerThanB = false
	}

	if pairIsBad(slice[0], slice[1], aSmallerThanB) {
		newSliceA := append([]int{}, slice[1:]...)
		newSliceB := append([]int{slice[0]}, slice[2:]...)
		if !levelIsSafe(newSliceB) && !levelIsSafe(newSliceA) {
			return false
		}

	}
	for i := 1; i < len(slice)-1; i++ {
		if pairIsBad(slice[i], slice[i+1], aSmallerThanB) {
			leftA := append([]int{}, slice[:i]...)
			newSliceA := append(leftA, slice[i+1:]...)
			leftB := append([]int{}, slice[:i+1]...)
			newSliceB := append(leftB, slice[i+2:]...)
			if !levelIsSafe(newSliceB) && !levelIsSafe(newSliceA) {
				return false
			}

		}
	}
	return true
}

func pairIsBad(a, b int, aSmallerThanB bool) bool {
	return a < b != aSmallerThanB || abs(a-b) > 3 || a-b == 0
}

func getInput() string {
	fileName := "demo.txt"
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

	return strings.TrimSpace(input)
}

func abs(num int) int {
	return max(num, -num)
}

func line2NumSlice(line string) []int {
	numSlice := []int{}
	for _, strNum := range strings.Split(line, " ") {
		if strNum == "" {
			continue
		}
		val, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("non number found in input!")
			os.Exit(1)
		}
		numSlice = append(numSlice, val)
	}
	return numSlice
}
