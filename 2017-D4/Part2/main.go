package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPU_THERE`

var myMap map[string]int

func main() {

	p2answer := 0

	for _, phrase := range strings.Split(input, "\n") {
		if phrase == "" { //the newline at the end of the input appers as an empty phrase
			continue
		}
		myMap = make(map[string]int)
		allgood := true

		for _, word := range strings.Split(phrase, " ") {
			word = sort(word)
			_, ok := myMap[word]
			myMap[word]++
			if ok {
				allgood = false
				break
			}
		}
		if allgood {
			p2answer++
		}

	}

	fmt.Println("Part 2:", p2answer)

}

var arr [26]int

func sort(str string) string {
	retStr := ""
	for i, _ := range arr {
		arr[i] = 0
	}

	for _, ch := range str {
		arr[int(ch-'a')]++
	}

	for i, _ := range arr {
		for rep := 0; rep < arr[i]; rep++ {
			retStr += string(i + 'a')
		}

	}
	return retStr

}
