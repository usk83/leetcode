// 949. Largest Time for Given Digits
// https://leetcode.com/problems/largest-time-for-given-digits/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
	pp.Println("23:41")
	pp.Println("=========================================")
	pp.Println(largestTimeFromDigits([]int{5, 5, 5, 5}))
	pp.Println("")
	pp.Println("=========================================")
	pp.Println(largestTimeFromDigits([]int{0, 0, 0, 0}))
	pp.Println("00:00")
	pp.Println("=========================================")
	pp.Println(largestTimeFromDigits([]int{1, 9, 6, 0}))
	pp.Println("19:06")
	pp.Println("=========================================")
	pp.Println(largestTimeFromDigits([]int{2, 0, 6, 6}))
	pp.Println("06:26")
	pp.Println("=========================================")
}

// A.length == 4
// 0 <= A[i] <= 9
func largestTimeFromDigits(A []int) string {
	digits := [10]int{}
	for _, a := range A {
		digits[a]++
	}

	// ab:cd
	// a: <2
	// b: if a == 2 <6 else <=9
	// c: <6
	// d: <=9

	if digits.contains(2) {

	}

	return ""
}
