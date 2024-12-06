package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type page struct {
	left  []int
	right []int
}

func main() {
	input := getInput()
	fmt.Println("Part 1 answer:", part1(input))
	fmt.Println("Part 2 answer:", part2(input))
}

var rulesMap = make(map[int]page)

func part1(input string) string {
	parts := strings.Split(input, "\n\n")

	books := strings.Split(parts[1], "\n")
	orderRules := strings.Split(parts[0], "\n")

	initMap(orderRules)

	total := 0
	for _, pages := range books {
		slice, ok := pagesAreSorted(pages)
		if ok {
			total += slice[len(slice)/2]
		}
	}
	return fmt.Sprint(total)
}

func part2(input string) string {
	parts := strings.Split(input, "\n\n")

	books := strings.Split(parts[1], "\n")

	total := 0
	for _, pages := range books {
		slice, ok := pagesAreSorted(pages)
		if !ok {
			slice = sortSlice(slice)
			total += slice[len(slice)/2]
		}
	}
	return fmt.Sprint(total)
}

// set up map that tells the order of numbers
func initMap(orderRules []string) {
	for _, rule := range orderRules {
		parts := strings.Split(rule, "|")

		A := atoi(parts[0])
		B := atoi(parts[1])

		val, ok := rulesMap[A]
		if ok {
			val.right = append(val.right, B)
			rulesMap[A] = val
		} else {
			rulesMap[A] = page{right: []int{B}}
		}

		val, ok = rulesMap[B]
		if ok {
			val.left = append(val.left, A)
			rulesMap[B] = val
		} else {
			rulesMap[B] = page{left: []int{A}}
		}
	}
}

// ascii to int
func atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("yo, bad numbers", str, err)
		os.Exit(1)
	}
	return val
}

func pagesAreSorted(pages string) ([]int, bool) {
	pageNums := []int{}
	for _, pageStr := range strings.Split(pages, ",") {
		pageNums = append(pageNums, atoi(pageStr))
	}
	//nested loop to check if a number is found that is on the right of another number when it shouldn't be
	for start, num := range pageNums {
		if start == len(pageNums)-1 {
			continue
		}
		Lpage := rulesMap[num]
		for i2 := start + 1; i2 < len(pageNums); i2++ {
			if slices.Contains(Lpage.left, pageNums[i2]) {
				return pageNums, false
			}
		}
	}
	return pageNums, true
}

func sortSlice(slice []int) []int {
	newSlice := []int{slice[0]}
	slice = slice[1:]
	//well try to insert each number into a new slice in the correct order
	for _, addMe := range slice {
		stillNeedToAdd := true
		//from the left, insert next to the first number that has to be to your right
		for i := 0; i < len(newSlice); i++ {
			if slices.Contains(rulesMap[addMe].right, newSlice[i]) {
				newSlice = slices.Insert(newSlice, i, addMe)
				stillNeedToAdd = false
				break
			}
		}

		//from the right, insert next to the first number that has to be on your left
		if stillNeedToAdd {
			for i := len(newSlice) - 1; i >= 0; i-- {
				if slices.Contains(rulesMap[addMe].left, newSlice[i]) {
					newSlice = slices.Insert(newSlice, i+1, addMe)
					stillNeedToAdd = false
					break
				}
			}
		}

		//else just append it to the end, as apparently it doens't matter
		if stillNeedToAdd {
			newSlice = append(newSlice, addMe)
		}
	}
	return newSlice
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
	input := strings.ReplaceAll(string(data), "\r\n", "\n") // doing this replace so it can handle both linux and window text format
	return strings.TrimSpace(input)                         // doing this cause usually there's an extra new line at the bottom of the input
}
