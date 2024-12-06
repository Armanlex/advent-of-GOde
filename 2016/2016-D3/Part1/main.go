package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

func main() {

	count := 0

	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		a := -1
		b := -1
		c := -1

		for _, thing := range strings.Split(line, " ") {
			if thing == "" {
				continue
			}

			if a == -1 {
				a, _ = strconv.Atoi(thing)
			} else if b == -1 {
				b, _ = strconv.Atoi(thing)
			} else if c == -1 {
				c, _ = strconv.Atoi(thing)
			}

		}

		if a+b > c && a+c > b && b+c > a {
			count++
		}

	}
	fmt.Println("Part 1: ", count)

}
