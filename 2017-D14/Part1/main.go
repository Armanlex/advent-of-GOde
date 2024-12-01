package main

import (
	"fmt"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

// var disk [128][128]int

func main() {

	p1answer := 0

	for i := 0; i < 128; i++ {
		newInput := input + "-" + fmt.Sprint(i)
		hash := KnotHash(newInput)
		p1answer += strings.Count(hash, "1")
	}

	fmt.Println("Part 1:", p1answer)

}

var arr [256]int

func KnotHash(strInput string) string { //copy pasta from day 10, only changes at line ~60, so that it immediately outputs binary instead of hex
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
		b := fmt.Sprintf("%b", res) //Changes here!, instead of outputting in hex, changed to binary
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
