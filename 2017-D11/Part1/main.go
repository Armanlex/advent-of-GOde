package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

//   \ n  /
// nw +--+ ne
//   /    \
// -+      +-
//   \    /
// sw +--+ se
//   / s  \

func main() {
	x := 0
	y := 0

	if strings.Contains(input, "\n") {
		input = strings.Split(input, "\n")[0]
	}

	for _, dir := range strings.Split(input, ",") {

		switch dir {
		case "n":
			y -= 2
		case "s":
			y += 2
		case "ne":
			y--
			x++
		case "nw":
			y--
			x--
		case "se":
			y++
			x++
		case "sw":
			y++
			x--
		default:
			fmt.Println("WTF")
		}

	}

	fmt.Println(x, y)
	if y < 0 {
		y *= -1
	}
	if x < 0 {
		x *= -1
	}

	if x > y {
		fmt.Println("Part 1:", x) //hard to explain, I barely understand it myself
	} else {
		fmt.Println("Part 1:", (y+x)/2)
	}

}
