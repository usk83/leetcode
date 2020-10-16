// 409. Longest Palindrome
// https://leetcode.com/problems/longest-palindrome/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println(7)
	// pp.Println("=========================================")
	// pp.Println(longestPalindrome("ccc"))
	// pp.Println(3)
	// pp.Println("=========================================")
	// pp.Println(longestPalindrome("zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"))
	// pp.Println()
	// pp.Println("=========================================")

	pp.Println("=========================================")
	pp.Println(longestPalindrome("AAA"))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(longestPalindrome("ZZZaaa"))
	pp.Println(5)
	pp.Println("=========================================")
}

func longestPalindrome(s string) int {
	charsMap := [52]int{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > '`' {
			charsMap[s[i]-'a'+26]++
		} else {
			charsMap[s[i]-'@']++
		}
	}
}

// func longestPalindrome(s string) int {
// 	// 0 1 2 3 4 ... 26 27
// 	// A B C         a  b

// 	charsMap := [52]int{}
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] > '`' {
// 			charsMap[s[i]-'a'+26]++
// 		} else {
// 			charsMap[s[i]-'@']++
// 		}
// 	}

// 	// pp.Println("=== DEBUG START ======================================")
// 	// fmt.Println(charsMap)
// 	// pp.Println("=== DEBUG END ========================================")

// 	result := 0
// 	for _, charCount := range charsMap {
// 		if charCount&1 == 0 {
// 			result += charCount
// 		} else {
// 			result += charCount - 1
// 		}
// 	}

// 	if result == len(s) {
// 		return result
// 	}
// 	return result + 1
// }

// []byte()

// 単方向ループ
// 単純回文を探す
// DP? Backtracking?

// var longest int
// pointers := []int{}
// for i := 1; i < len(s); i++ {
// 	for j := 0; j < len(pointers); j++ {
// 		if s[i] == s[pointers[j]] {
// 			pointers[j]--
// 		} else {
// 			if current := i - pointers[j]; current > longest {
// 				longest = current
// 			}
// 			if j == len(pointers)-1 {
// 				pointers = pointers[:j]
// 			} else {
// 				pointers = append(pointers[:j], pointers[j+1:]...)
// 			}
// 			j--
// 		}
// 	}

// 	if s[i] == s[i-1] {
// 		pointers = append(pointers, i-2)
// 	}
// 	if i > 1 && s[i] == s[i-2] {
// 		pointers = append(pointers, i-3)
// 	}
// }

// return longest
// }

// func longestPalindrome(s string) int {

// 	// 単方向ループ
// 	// 単純回文を探す
// 	// DP? Backtracking?

// 	var longest int
// 	pointers := []int{}
// 	for i := 1; i < len(s); i++ {
// 		for j := 0; j < len(pointers); j++ {
// 			if s[i] == s[pointers[j]] {
// 				pointers[j]--
// 			} else {
// 				if current := i - pointers[j]; current > longest {
// 					longest = current
// 				}
// 				if j == len(pointers)-1 {
// 					pointers = pointers[:j]
// 				} else {
// 					pointers = append(pointers[:j], pointers[j+1:]...)
// 				}
// 				j--
// 			}
// 		}

// 		if s[i] == s[i-1] {
// 			pointers = append(pointers, i-2)
// 		}
// 		if i > 1 && s[i] == s[i-2] {
// 			pointers = append(pointers, i-3)
// 		}
// 	}

// 	return longest
// }
