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
	inputLines := strings.Split(input, "\n")
	niceCount := 0
	for _, line := range inputLines {
		if isNaughtyP1(line) {
			continue
		}
		niceCount++
	}
	return fmt.Sprint(niceCount)
}

func part2(input string) string {
	inputLines := strings.Split(input, "\n")
	niceCount := 0
	for _, line := range inputLines {
		if isNaughtyP2(line) {
			continue
		}
		niceCount++
	}
	return fmt.Sprint(niceCount)
}

func isNaughtyP1(str string) bool {
	count := 0
	noDupes := true
	runeSlice := []rune(str)
	for i, r := range runeSlice {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
		if i > 0 {
			a := runeSlice[i-1]
			b := runeSlice[i]
			if a == b {
				noDupes = false
			}
			switch string(a) + string(b) {
			case "ab", "cd", "pq", "xy":
				return true
			}
		}
	}
	if count < 3 || noDupes {
		return true
	}
	return false
}

func isNaughtyP2(str string) bool {
	pairMissing := true
	repetitionMissing := true
	runeSlice := []rune(str)
	myMap := make(map[string]int)
	for i := 0; i < len(runeSlice)-1; i++ {
		if pairMissing {
			word := string(runeSlice[i]) + string(runeSlice[i+1])
			val, ok := myMap[word]
			if ok {
				if val != i {
					pairMissing = false
				}
			} else {
				myMap[word] = i + 1
			}
		}
		if i > 0 && repetitionMissing {
			if runeSlice[i-1] == runeSlice[i+1] {
				repetitionMissing = false
			}
		}
		if !repetitionMissing && !pairMissing {
			return false
		}
	}

	if pairMissing || repetitionMissing {
		return true
	}

	return false
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
