// 402. Remove K Digits
// https://leetcode.com/problems/remove-k-digits/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 13
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3328/
func main() {
	pp.Println("=========================================")
	pp.Println(removeKdigits("1432219", 3))
	pp.Println("1219")
	pp.Println("=========================================")
	pp.Println(removeKdigits("10200", 1))
	pp.Println("200")
	pp.Println("=========================================")
	pp.Println(removeKdigits("10", 2))
	pp.Println("0")
	pp.Println("=========================================")
	pp.Println(removeKdigits("9", 1))
	pp.Println("0")
	pp.Println("=========================================")
	pp.Println(removeKdigits("112", 1))
	pp.Println("11")
	pp.Println("=========================================")
	pp.Println(removeKdigits("1107", 1))
	pp.Println("107")
	pp.Println("=========================================")
}

// ref: https://leetcode.com/problems/remove-k-digits/discuss/629860/Java-or-C++-or-Python3-or-easy-explanation
// 一旦は採用して、あとから修正するタイプのGreedy
// func removeKdigits(num string, k int) string {
// 	pointer, minNum := 0, make([]byte, 0, len(num))
// 	for pointer < len(num) && k != 0 {
// 		if len(minNum) == 0 || minNum[len(minNum)-1] <= num[pointer] {
// 			if len(minNum) != 0 || num[pointer] != '0' {
// 				minNum = append(minNum, num[pointer])
// 			}
// 			pointer++
// 		} else {
// 			minNum = minNum[:len(minNum)-1]
// 			k--
// 		}
// 	}
// 	if len(minNum) == 0 {
// 		for pointer < len(num) && num[pointer] == '0' {
// 			pointer++
// 		}
// 	}
// 	minNum = append(minNum, num[pointer:]...)
// 	if len(minNum)-k <= 0 {
// 		return "0"
// 	}
// 	return string(minNum[:len(minNum)-k])
// }

// 毎回小さい方から順に届くやつを使っていく
// O(N)
func removeKdigits(num string, k int) string {
	indices := [10][]int{}
	for i, n := range num {
		indices[n-'0'] = append(indices[n-'0'], i)
	}
	pointer, minNum := 0, make([]byte, 0, len(num)-k)
LOOP:
	for 0 < k && k < len(num)-pointer { // O(N)
		for digit := range indices { // O(10)
			for len(indices[digit]) != 0 && indices[digit][0] < pointer {
				indices[digit] = indices[digit][1:]
			}
			if len(indices[digit]) != 0 && indices[digit][0]-pointer <= k {
				k -= indices[digit][0] - pointer
				pointer = indices[digit][0] + 1
				indices[digit] = indices[digit][1:]
				if digit != 0 || len(minNum) != 0 {
					minNum = append(minNum, byte(digit+'0'))
				}
				continue LOOP
			}
		}
	}
	if k == 0 {
		minNum = append(minNum, num[pointer:]...)
	}
	if len(minNum) == 0 {
		return "0"
	}
	return string(minNum)
}

// 届く範囲の最小までカットを繰り返す

// func removeKdigits(num string, k int) string {
// 	digitsIndices := [10][]int{}
// 	for i, n := range num {
// 		digitsIndices[n-'0'] = append(digitsIndices[n-'0'], i)
// 	}

// 	pointer, minNum := 0, make([]byte, 0, len(num)-k)
// 	for digit, indices := range digitsIndices {
// 		for _, index := range indices {
// 			diff := index - pointer
// 			pp.Printf("digit: %s, index: %s, diff: %s, k: %s\n", digit, index, diff, k)
// 			pp.Printf("digit: %s, index: %s\n", digit, index)
// 			if diff < 0 || diff > k {
// 				pp.Println("break")
// 				break
// 			}
// 			k -= diff
// 			pointer = index + 1
// 			if digit != 0 {
// 				minNum = append(minNum, byte(digit+'0'))
// 			}
// 			pp.Println(string(minNum))
// 		}
// 	}

// 	// pp.Println("=== DEBUG START ======================================")
// 	// pp.Println(minNum)
// 	// pp.Println(pointer)
// 	// pp.Println("=== DEBUG END ========================================")

// 	minNum = append(minNum, num[pointer:]...)

// 	if len(minNum) == 0 {
// 		return "0"
// 	}
// 	return string(minNum)
// }
