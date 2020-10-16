// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findNumbers([]int{555, 901, 482, 1771}))
	pp.Println(1)
	pp.Println("=========================================")
}

// Constraints:
//     1 <= nums.length <= 500
//     1 <= nums[i] <= 10^5
// Static solution
func findNumbers(nums []int) int {
	var count int
	for _, num := range nums {
		if (10 <= num && num < 100) ||
			(1000 <= num && num < 10000) ||
			(100000 <= num && num < 1000000) {
			count++
		}
	}
	return count
}

// Math solution
// func findNumbers(nums []int) int {
// 	var count int
// 	for _, num := range nums {
// 		var numOfDigits int
// 		for num != 0 {
// 			numOfDigits++
// 			num /= 10
// 		}
// 		count += 1 - numOfDigits&1
// 	}
// 	return count
// }

// Original solution (converting to string solution)
// func findNumbers(nums []int) int {
// 	var count int
// 	for _, num := range nums {
// 		if len(strconv.Itoa(num))&1 == 0 {
// 			count++
// 		}
// 	}
// 	return count
// }
