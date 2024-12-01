package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

var input string = `YOUR_INPUT_HERE`

func main() {

	i := 0
	times := 0
	pass := ""
	for times < 8 {
		i++
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if hash[:5] == "00000" {
			pass += string(hash[5])
			times++
			fmt.Println("in progress: ", pass, " < ", hash, " on index: ", i)
		}
	}

	fmt.Println("Part 1: ", pass)

}
