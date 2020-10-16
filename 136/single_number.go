// 136. Single Number
// https://leetcode.com/problems/single-number/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(singleNumber([]int{2, 2, 1}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	pp.Println(4)
	pp.Println("=========================================")
}

/*
 * 2020-04-01
 * https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3283/
 *
 * XOR extraction
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 */
func singleNumber(nums []int) int {
	var single int
	for _, num := range nums {
		single ^= num
	}
	return single
}

/*
 * Sort
 * - Time Complexity: O(NlogN)
 * - Space Complexity: O(logN)
 */
// func singleNumber(nums []int) int {
// 	sort.Ints(nums)
// 	lastIndex := len(nums) - 1
// 	for i := 0; i < lastIndex; i += 2 {
// 		if nums[i] != nums[i+1] {
// 			return nums[i]
// 		}
// 	}
// 	return nums[lastIndex]
// }

/*
 * Calcurate sum of unique numbers
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func singleNumber(nums []int) int {
// 	m := map[int]struct{}{}
// 	var wholeSum int
// 	for _, num := range nums {
// 		m[num] = struct{}{}
// 		wholeSum += num
// 	}
// 	var uniqueSum int
// 	for num := range m {
// 		uniqueSum += num
// 	}
// 	return uniqueSum*2 - wholeSum
// }

/*
 * Using a Set to filter the single number
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func singleNumber(nums []int) int {
// 	m := map[int]struct{}{}
// 	for _, num := range nums {
// 		if _, found := m[num]; found {
// 			delete(m, num)
// 		} else {
// 			m[num] = struct{}{}
// 		}
// 	}
// 	for k := range m {
// 		return k
// 	}
// 	return -1
// }
