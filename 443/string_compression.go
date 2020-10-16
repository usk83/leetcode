// 443. String Compression
// https://leetcode.com/problems/string-compression/
package main

import (
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(compress([]byte{'a'}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
	pp.Println(4)
	pp.Println("=========================================")
}

// All characters have an ASCII value in [35, 126].
// 1 <= len(chars) <= 1000.
// The final clean code
func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	lastIndex, count := 0, 1

	allocate := func() {
		if count == 1 {
			return
		}
		for _, digit := range []byte(strconv.Itoa(count)) {
			lastIndex++
			chars[lastIndex] = digit
		}
	}

	for i, charsLen := 1, len(chars); i < charsLen; i++ {
		if chars[i] == chars[lastIndex] {
			chars[i] = 0
			count++
		} else {
			allocate()
			lastIndex++
			chars[lastIndex] = chars[i]
			count = 1
		}

	}
	allocate()

	return lastIndex + 1
}

// 100%, 100% answer
// func compress(chars []byte) int {
// 	if len(chars) == 1 {
// 		return 1
// 	}

// 	lastIndex, count := 0, 1

// 	allocate := func() {
// 		if count == 1 {
// 			return
// 		}
// 		nums := []byte{}
// 		for count != 0 {
// 			nums = append(nums, byte(count%10+48)) // 48: '1' - 1
// 			count /= 10
// 		}
// 		for i := len(nums) - 1; i >= 0; i-- {
// 			lastIndex++
// 			chars[lastIndex] = nums[i]
// 		}
// 	}

// 	for i, charsLen := 1, len(chars); i < charsLen; i++ {
// 		if chars[i] == chars[lastIndex] {
// 			chars[i] = 0
// 			count++
// 		} else {
// 			allocate()
// 			lastIndex++
// 			chars[lastIndex] = chars[i]
// 			count = 1
// 		}

// 	}
// 	allocate()

// 	return lastIndex + 1
// }

// The second O(n), O(n) answer
// func compress(chars []byte) int {
// 	if len(chars) == 0 {
// 		return 1
// 	}
// 	prevIndex, count, size := 0, 1, 0
// 	allocate := func() {
// 		chars[size] = chars[prevIndex]
// 		size++
// 		if count == 1 {
// 			return
// 		}
// 		nums := []byte{}
// 		for count != 0 {
// 			nums = append(nums, byte(count%10+48)) // '1' - 1
// 			count /= 10
// 		}
// 		for i := len(nums) - 1; i >= 0; i-- {
// 			chars[size] = nums[i]
// 			size++
// 		}
// 	}

// 	for i := 1; i < len(chars); i++ {
// 		if chars[i] != chars[prevIndex] {
// 			allocate()
// 			prevIndex = i
// 			count = 1
// 		} else {
// 			chars[i] = 0
// 			count++
// 		}
// 	}
// 	allocate()

// 	return size
// }

// The first answer
// func compress(chars []byte) int {
// 	if len(chars) == 0 {
// 		return 1
// 	}

// 	currentChar := chars[0]
// 	count := 1
// 	newChars := make([]byte, 0, len(chars))

// 	extend := func() {
// 		newChars = append(newChars, currentChar)
// 		if count != 1 {
// 			var num string
// 			for {
// 				num = strconv.Itoa(count%10) + num
// 				count /= 10
// 				if count == 0 {
// 					break
// 				}
// 			}
// 			newChars = append(newChars, []byte(num)...)
// 		}
// 	}

// 	for _, char := range chars[1:] {
// 		if currentChar != 0 && char != currentChar {
// 			extend()
// 			count = 1
// 			currentChar = char
// 		} else {
// 			count++
// 		}
// 	}

// 	extend()

// 	for i := 0; i < len(chars); i++ {
// 		if i < len(newChars) {
// 			chars[i] = newChars[i]
// 		} else {
// 			chars[i] = 0
// 		}
// 	}

// 	return len(newChars)
// }
