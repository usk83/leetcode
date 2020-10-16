// 859. Buddy Strings
// https://leetcode.com/problems/buddy-strings/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(buddyStrings("ab", "ba"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(buddyStrings("ab", "ab"))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(buddyStrings("aa", "aa"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(buddyStrings("", "aa"))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(buddyStrings("ab", "ab"))
	pp.Println(false)
	pp.Println("=========================================")
}

// 0 <= A.length <= 20000
// 0 <= B.length <= 20000
// A and B consist only of lowercase letters.
func buddyStrings(A string, B string) bool {
	if len(A) < 2 || len(A) != len(B) {
		return false
	}

	letterCount, indices := [26]int{}, []int{}
	for i := 0; i < len(A); i++ {
		letterCount[A[i]-'a']++
		if A[i] != B[i] {
			indices = append(indices, i)
		}
	}

	if len(indices) == 0 {
		for _, count := range letterCount {
			if count > 1 {
				return true
			}
		}
	}

	return len(indices) == 2 && A[indices[0]] == B[indices[1]] && A[indices[1]] == B[indices[0]]
}

// func buddyStrings(A string, B string) bool {
// 	if len(A) < 2 || len(A) != len(B) {
// 		return false
// 	}
// 	lettersCount := [26]uint16{}
// 	wrongIndices := make([]int, 0, 2)
// 	for i := 0; i < len(A); i++ {
// 		lettersCount[A[i]-'a']++
// 		if A[i] != B[i] {
// 			switch len(wrongIndices) {
// 			case 0:
// 				wrongIndices = append(wrongIndices, i)
// 			case 1:
// 				if A[wrongIndices[0]] != B[i] || A[i] != B[wrongIndices[0]] {
// 					return false
// 				}
// 				wrongIndices = append(wrongIndices, i)
// 			default:
// 				return false
// 			}
// 		}
// 	}
//
// 	switch len(wrongIndices) {
// 	case 0:
// 		for _, count := range lettersCount {
// 			if count > 1 {
// 				return true
// 			}
// 		}
// 	case 2:
// 		return true
// 	}
//
// 	return false
// }

// func buddyStrings(A string, B string) bool {
// 	if len(A) < 2 || len(A) != len(B) {
// 		return false
// 	}
//
// 	chars := [26]int{}
// 	index := []int{}
// 	for i := 0; i < len(A); i++ {
// 		chars[A[i]-'a']++
// 		if A[i] != B[i] {
// 			index = append(index, i)
// 		}
// 	}
//
// 	if len(index) == 0 {
// 		for _, chars := range chars {
// 			if chars > 1 {
// 				return true
// 			}
// 		}
// 		return false
// 	}
//
// 	if len(index) != 2 {
// 		return false
// 	}
//
// 	return A[index[0]] == B[index[1]] && A[index[1]] == B[index[0]]
// }
