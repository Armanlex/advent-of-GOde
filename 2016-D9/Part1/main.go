package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

func main() {

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		final_str := ""
		in_marker := false
		marker_text := ""
		for i := 0; i < len(line); i++ {
			if line[i] == '(' {
				in_marker = true
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

				for rep := 0; rep < reps; rep++ {
					final_str += line[i+1 : i+1+length]
				}

				marker_text = ""
				i += length

			} else {
				if in_marker {
					marker_text += string(line[i])
				} else {
					final_str += string(line[i])
				}
			}

		}
		fmt.Println("Part 1: ", len(final_str))

	}

}
