package main

import (
	"fmt"
	"strings"
)

var input string = `ADD_YOUR_INPUT_HERE`

var myMap map[string]int

func main() {

	p1answer := 0

	for _, phrase := range strings.Split(input, "\n") {
		myMap = make(map[string]int)
		allgood := true

		for _, word := range strings.Split(phrase, " ") {
			_, ok := myMap[word]
			myMap[word]++
			if ok {
				allgood = false
				break
			}
		}
		if allgood {
			p1answer++
		}

	}

	fmt.Println("Part 1:", p1answer)

}
