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
	pass := []rune("--------")
	for times < 8 {
		i++
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if hash[:5] == "00000" {
			index := hash[5]
			if index >= '0' && index <= '7' {

				if pass[index-'0'] == '-' {
					pass[index-'0'] = rune(hash[6])
					times++
					fmt.Println("in progress: ", string(pass), " < ", hash, " on index: ", i)
				}
			}
		}
	}

	fmt.Println("Part 2: ", string(pass))

}
