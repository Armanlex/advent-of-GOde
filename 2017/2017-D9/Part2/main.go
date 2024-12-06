package main

import "fmt"

var input string = `YOUR_INPUT_HERE`

// vscode seems to freakout when you enter too long of a string above
// and might mess up the tokenization (coloring) the text of this entire file, but it's fine, don't worry about it

func main() {
	pointer := -1
	p2answer := 0
	//counters

	groupDepth := 0
	//flags
	inGarb := false
	for {

		pointer++

		if pointer >= len(input) {
			break
		}

		switch input[pointer] {
		case '{':
			if inGarb {
				p2answer++
				continue
			}
			groupDepth++

		case '}':
			if inGarb {
				p2answer++
				continue
			}
			groupDepth--

		case '<':
			if inGarb {
				p2answer++
				continue
			}
			inGarb = true

		case '>':
			if inGarb {
				inGarb = false
				continue
			}

		case '!':
			pointer++

		case ',':
			if inGarb {
				p2answer++
				continue
			}

		default: //if non of the above are true
			if inGarb {
				p2answer++
				continue
			}

		}

	}

	fmt.Println("Part 2:", p2answer)
}
