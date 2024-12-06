package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

var aSlice = []int{}
var bSlice = []int{}
var cSlice = []int{}

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
				aSlice = append(aSlice, a)
			} else if b == -1 {
				b, _ = strconv.Atoi(thing)
				bSlice = append(bSlice, b)
			} else if c == -1 {
				c, _ = strconv.Atoi(thing)
				cSlice = append(cSlice, c)
			}

		}
	}

	slices := [][]int{aSlice, bSlice, cSlice}

	for i := 0; i < len(aSlice); i += 3 {
		for si := 0; si < 3; si++ {
			a := slices[si][i]
			b := slices[si][i+1]
			c := slices[si][i+2]
			if a+b > c && a+c > b && b+c > a {
				count++
			}
		}
	}
	fmt.Println("Part 2: ", count)

}
