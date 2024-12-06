package main

import "fmt"

//No need for an actual slice here, just keep track of the "slices" length, and only look at what number is inserted when the pointer is at 0

var input int = 0 //YOUR INPUT HERE

var limit int = 50000000

func main() {
	slicelen := 1
	p2answer := 0
	pointer := 0
	for i := 1; i < limit; i++ {
		if pointer == 0 {
			p2answer = i
		}
		pointer++
		slicelen++
		pointer = (pointer + input) % slicelen
	}

	fmt.Println("Part 2:", p2answer)

}
