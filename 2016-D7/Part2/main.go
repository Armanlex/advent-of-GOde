package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

var demo = `aba[bab]xyz
xyx[xyx]xyx
aaa[kek]eke
zazbz[bzb]cdb`

func main() {

	counter := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		line += "]"
		mymap := make(map[string]int)

		for x := 0; x < 2; x++ {
			hyper := false
			for i := 0; i < len(line)-2; i++ {
				if line[i] == '[' {
					hyper = true
				} else if line[i] == ']' {
					hyper = false
				}
				if hyper && x == 1 {
					if line[i] == line[i+2] && line[i] != line[i+1] && line[i+1] != '[' && line[i+1] != ']' {
						_, ok := mymap[line[i:i+3]]

						if ok {
							counter++
							break
						}
					}
				} else if x == 0 && hyper == false {
					if line[i] == line[i+2] && line[i] != line[i+1] && line[i+1] != '[' && line[i+1] != ']' {
						str := string(line[i+1]) + string(line[i]) + string(line[i+1])
						mymap[str] = 1

					}
				}
			}
		}
	}
	fmt.Println("Part 2: ", counter)

}
