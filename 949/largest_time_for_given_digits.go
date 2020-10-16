// 949. Largest Time for Given Digits
// https://leetcode.com/problems/largest-time-for-given-digits/
package main

import (
	"math"
	"strconv"

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
	max := 0
	flag := false
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				for l := 0; l < 4; l++ {
					if i == l || j == l || k == l {
						continue
					}
					sum := 0
					sum += A[0] * int(math.Pow10(i))
					sum += A[1] * int(math.Pow10(j))
					sum += A[2] * int(math.Pow10(k))
					sum += A[3] * int(math.Pow10(l))

					if sum > 2359 {
						continue
					}

					if sum%100 > 59 {
						continue
					}

					if sum <= 2359 && sum >= max {
						flag = true
						max = sum
					}
				}
			}
		}
	}

	// pp.Println(max)

	// if max == 0 {
	if !flag {
		return ""
	}

	result := strconv.Itoa(max)

	for len(result) < 4 {
		result = "0" + result
	}

	return string(result[:2]) + ":" + string(result[2:])

	// if A.contains(2) {
	// 	for i := 3; i
	// }

	// reft := []int{}
	// var index int
	// var max int
	// if A.contains(2) {
	// 	index = A.indexOf(2)
	// 	result += "2"
	// 	max =
	// } else if A.contains(1) {
	// 	index = A.indexOf(1)
	// 	result += "1"
	// } else if A.contains(0) {
	// 	index = A.indexOf(0)
	// 	result += "0"
	// } else {
	// 	return ""
	// }

	// for i, n := range A {
	// 	if i == index {
	// 		continue
	// 	}
	// 	reft = append(reft, n)
	// }

	// return ""
}
