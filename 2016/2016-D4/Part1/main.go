package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

func main() {
	val := 0

	counterArray := [26][2]int{} //will use this to sort the letters found by the frequency and then alphabetically
	for i := 0; i < 26; i++ {
		counterArray[i][1] = i + 'a'
	}
	counterArrayCopy := counterArray //creating a copy to use base slices of off

	sectorsum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		counterArrayCopy = counterArray //whiping the slate clean between loops
		counterSlice := counterArrayCopy[:]
		parts := strings.Split(line, "[")

		for _, thing := range strings.Split(parts[0], "-") {
			if thing == "" {
				continue
			}
			if thing[0] >= '0' && thing[0] <= '9' {
				//sector ID

				val, _ = strconv.Atoi(thing)
				//   ^ yeah yeah

			} else {
				//word
				for _, r := range thing {
					counterSlice[r-'a'][0]++
				}
			}
		}
		checksum := parts[1][:len(parts[1])-1]

		sort.Slice(counterSlice, func(i, j int) bool { //custom sorting
			if counterSlice[i][0] == counterSlice[j][0] { //if sizes same,
				if counterSlice[i][1] < counterSlice[j][1] { //compare the letter, smaller letter goes left
					return true
				} else {
					return false
				}
			} else { //if sizes different, compare sizes, bigger goes to the left
				return counterSlice[i][0] > counterSlice[j][0]
			}
		})

		left := ""
		left += string(counterSlice[0][1])
		left += string(counterSlice[1][1])
		left += string(counterSlice[2][1])
		left += string(counterSlice[3][1])
		left += string(counterSlice[4][1])
		fmt.Println(left, checksum)
		if left == checksum {
			sectorsum += val
		}

	}

	fmt.Println("Part 1: ", sectorsum)

}
