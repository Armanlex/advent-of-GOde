package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

var line string = ``

func main() {
	for _, line = range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		fmt.Println("Part 2: ", read(0, len(line), 1))
	}
}

func read(start int, end int, mult int) int { //return how many words this part printed

	counter := 0
	in_marker := false
	marker_text := ""

	for i := start; i < end; i++ {
		fmt.Print(string(line[i]), ", ")
		if line[i] == '(' {
			in_marker = true
			marker_text = ""

		} else if line[i] == ')' {
			in_marker = false
			parts := strings.Split(marker_text, "x")
			length, err := strconv.Atoi(parts[0])

			if err != nil {
				fmt.Print("yo there's an error here")
			}

			reps, err2 := strconv.Atoi(parts[1])
			if err2 != nil {
				fmt.Print("yo there's an error here2")
			}

			counter += read(i+1, i+length+1, mult*reps)

			i += length

		} else {

			if in_marker {
				marker_text += string(line[i])
			} else {
				counter += mult
			}
		}
	}

	return counter
}
