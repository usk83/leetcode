// 5392. Maximum Score After Splitting a String
// https://leetcode.com/contest/weekly-contest-186/problems/maximum-score-after-splitting-a-string/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxScore("011101"))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(maxScore("00111"))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(maxScore("1111"))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(maxScore("00110001011111100111101010100010001101001110"))
	pp.Println(26)
	pp.Println("=========================================")
}

func maxScore(s string) int {
	var score, zero, one int
	left, right := 0, len(s)-1
	if s[left] == '0' {
		score++
	}
	if s[right] == '1' {
		score++
	}
	for i := left + 1; i < right; i++ {
		switch s[i] {
		case '0':
			zero++
		case '1':
			one++
		}
		if zero > one {
			score += zero
			zero, one = 0, 0
		}
	}
	return score + one
}

// func maxScore(s string) int {
// 	left := 0
// 	right := len(s) - 1
// 	score := 0
// 	if s[left] == '0' {
// 		score++
// 	}
// 	if s[right] == '1' {
// 		score++
// 	}
// 	for left < right-1 {
// 		if s[left+1] == '0' {
// 			left++
// 			score++
// 		} else if s[right-1] == '1' {
// 			right--
// 			score++
// 		} else {
// 			break
// 			// left++
// 			// right--
// 		}
// 	}

// 	zero, one := 0, 0
// 	for i := left + 1; i < right; i++ {
// 		if s[i] == '0' {
// 			zero++
// 		} else {
// 			one++
// 		}
// 	}

// 	max := zero
// 	if max < one {
// 		max = one
// 	}

// 	return score + max

// 	// count := 0
// 	// for i := range s {
// 	// 	if s[i] == '1' {
// 	// 		count++
// 	// 	}
// 	// }
// 	// if s[0] == '1' && s[len(s)-1] == '1' {
// 	// 	count--
// 	// }
// 	// return count
// }
