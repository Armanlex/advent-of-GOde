package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPUT_HERE!`

func main() {
	counter := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		line += "]"
		ok := false
		hyper := false
		for i := 0; i < len(line)-3; i++ {
			if line[i] == '[' {
				hyper = true
			} else if line[i] == ']' {
				hyper = false
			}
			if line[i] == line[i+3] && line[i+1] == line[i+2] && line[i] != line[i+1] {
				if hyper == false {
					ok = true
				} else {
					ok = false

					break
				}
			}
		}
		if ok {
			counter++
		}
	}
	fmt.Println("Part 1: ", counter)

}
