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

	largest := 0

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
		largest = max(largest, dist(x, y))

	}

	fmt.Println("Part 2:", largest)
}

func dist(x, y int) int {

	y = max(y, -y) //fastest way I know to calculate the absolute value
	x = max(x, -x)
	if x > y {
		return x
	} else {
		return (y + x) / 2
	}
}
