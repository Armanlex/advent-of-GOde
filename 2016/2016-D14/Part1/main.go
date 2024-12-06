package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

var input string = `YOUR_INPUT_HERE`
var cacheThrees = []byte{} //holds which letter is the first triplet, if any
var cacheFives = []byte{}  //holds which letter is the first pentuplet, if any
var hashSlice = []string{} //holds all the hashes after they've been generated

func main() {

	var answer = ""
	currhash := ""
	decnum := 0
	var triplet byte = 255
	for {

		//find hash or generate it
		// fmt.Println(len(answer), len(cacheThrees), len(cacheFives), len(hashSlice))
		currhash = getHash(decnum)

		//find triplet of currhash or generate it
		triplet = 255
		if len(cacheThrees) <= decnum {
			calcCombos(currhash)
		}
		triplet = cacheThrees[decnum]

		if triplet != 255 {

			//start 1k procedure
			for i := decnum + 1; i < 1000+decnum+1; i++ {
				if len(cacheFives) <= i {
					calcCombos(getHash(i))
				}

				if cacheFives[i] == triplet {
					answer += string(rune(triplet))
					break
				}

			}
		}

		if len(answer) == 64 {
			fmt.Println("Part 1: ", decnum)
			break
		}
		decnum++
	}
}

// check small cache
func calcCombos(hash string) {
	combo3 := 0
	combo5 := 0
	var combosThree byte = 255
	var combosFive byte = 255

	for i := 0; i < 32-1; i++ {
		if combosThree == 255 {
			if hash[i] == hash[i+1] {
				combo3++
				if combo3 == 2 {
					combosThree = hash[i]
				}
			} else {
				combo3 = 0
			}
		}
		if combosFive == 255 {
			if hash[i] == hash[i+1] {
				combo5++
				if combo5 == 4 {
					combosFive = hash[i]
					break
				}
			} else {
				combo5 = 0
			}
		}
	}

	cacheThrees = append(cacheThrees, combosThree)
	cacheFives = append(cacheFives, combosFive)
}

func getHash(num int) string {

	if len(hashSlice) > num {
		return hashSlice[num]
	} else {
		hashSlice = append(hashSlice, fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(num)))))

		return hashSlice[len(hashSlice)-1]
	}

}
