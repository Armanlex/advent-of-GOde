package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

func main() {
	val := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		sent := ""

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

				sent += thing + " "
			}
		}
		rSlice := []rune(sent)
		for i := 0; i < len(rSlice); i++ {
			if rSlice[i] != ' ' {
				rSlice[i] = rune((((int(rSlice[i]) - 'a') + val) % 26) + 'a')
			}
		}
		sent = string(rSlice)

		if sent[:3] == "nor" {
			fmt.Println("Part 2: ", val)
			return
		}

	}

}
