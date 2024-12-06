package main

import (
	"fmt"
)

var input string = `YOUR_INPUT_HERE`

var disk [128][128]int

var seenCells map[string]int //used to see which cells have been visited during the flooding process

var p2answer int = 0

func main() {

	seenCells = make(map[string]int)

	for y := 0; y < 128; y++ {
		newInput := input + "-" + fmt.Sprint(y)
		hash := KnotHash(newInput)
		for x, ch := range hash {
			disk[y][x] = int(ch - '0') //filling in the disk 2d array with all the bits
		}
	}

	for y := 0; y < len(disk); y++ {
		for x := 0; x < len(disk[0]); x++ {
			if disk[y][x] == 1 { //if cell has a 1, try to flood it
				flood(x, y, true)
			}
		}
	}

	fmt.Println("Part 2:", p2answer)

}

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// this basically tries to flood with "water" the cells it's give, the water "spreads" the all adjacent cells that are also 1's.
// We keep track of which cells have been touched by the water so that we don't run in circles.
// we have a flag to differentiate when this function is called recursively, and when not.
// Because the original call implies that we've encountered a new "island", a new group
func flood(x, y int, isroot bool) {
	// fmt.Println("flood on ", x, y, " with ", isroot)
	str := fmt.Sprint(x, ",", y)
	_, ok := seenCells[str]
	if ok {
		return
	}
	seenCells[str]++
	if isroot {
		p2answer++
	}
	//call flood on children, with root false

	for _, dir := range dirs {
		if is_valid(dir[0]+x, dir[1]+y) {
			if disk[dir[1]+y][dir[0]+x] == 1 {
				flood(dir[0]+x, dir[1]+y, false)
			}
		}
	}

}

func is_valid(x, y int) bool { //to avoid accessing outside the bounds of the array
	return x < len(disk[0]) && x >= 0 && y < len(disk) && y >= 0
}

var arr [256]int

func KnotHash(strInput string) string { //copy pasta from day 10, only changes at line 70, so that it immediately outputs binary instead of hex
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	suffix := []int{17, 31, 73, 47, 23}
	nums := []int{}

	for _, numRune := range strInput {
		nums = append(nums, int(numRune))
	}

	nums = append(nums, suffix...)

	pointer := 0
	skipSize := 0

	for round := 0; round < 64; round++ {

		for _, intVal := range nums {
			reverse(&arr, pointer, intVal+pointer-1)
			pointer += intVal + skipSize
			skipSize++
		}

	}

	word := ""

	for x := 0; x < 16; x++ {
		i := (16 * x)
		res := arr[0+i] ^ arr[1+i] ^ arr[2+i] ^ arr[3+i] ^ arr[4+i] ^ arr[5+i] ^ arr[6+i] ^ arr[7+i] ^ arr[8+i] ^ arr[9+i] ^ arr[10+i] ^ arr[11+i] ^ arr[12+i] ^ arr[13+i] ^ arr[14+i] ^ arr[15+i]
		b := fmt.Sprintf("%b", res)
		for len(b) < 8 {
			b = "0" + b
		}
		word += b
	}

	return word

}

func reverse(arr *[256]int, start int, end int) {
	a := 0
	b := 0
	for i := 0; i <= (end-start)/2; i++ {
		a = (start + i) % len(arr)
		b = (end - i) % len(arr)

		(*arr)[a], (*arr)[b] = (*arr)[b], (*arr)[a]
	}

}
