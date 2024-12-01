package main

import "fmt"

var input int = 0 //YOUR INPUT HERE

var limit int = 2018

var slice []int

func main() {
	slice = make([]int, 1, limit)
	pointer := 0
	for i := 1; i < limit; i++ {
		slice = append(slice[:pointer+1], slice[pointer:]...) //basically this cuts and combines the slice, but leaves an extra slot in the middle
		pointer++
		slice[pointer] = i //so we add our value into the middle
		if i != limit-1 {  //don't want to move the pointer on the very last step, cause then it'd be harder to print the solution
			pointer = (pointer + input) % len(slice)
		}
	}
	fmt.Println("Part 1:", slice[pointer+1])

}
