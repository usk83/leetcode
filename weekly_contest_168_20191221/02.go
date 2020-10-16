// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPossibleDivide([]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPossibleDivide([]int{3, 3, 2, 2, 1, 1}, 3))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPossibleDivide([]int{1, 2, 3, 4}, 3))
	pp.Println(false)
	pp.Println("=========================================")
}

// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^9
// 1 <= k <= nums.length
func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	// 何が分かれば解ける？
	// 数字ごとの出現回数
	// 小さい方からkずつのセット

	return true
}

// Original Solution
// func isPossibleDivide(nums []int, k int) bool {
// 	if len(nums)%k != 0 {
// 		return false
// 	}

// 	// O(n) time, O(n) space
// 	m := map[int]int{}
// 	for _, num := range nums {
// 		m[num]++
// 	}

// 	// O(log n) time, O(?) space
// 	keys := []int{}
// 	for key := range m {
// 		keys = append(keys, key)
// 	}

// 	// O(log n) time, O(?) space
// 	sort.Ints(keys)

// 	for _, key := range keys {
// 		count := m[key]
// 		if count == 0 {
// 			continue
// 		}
// 		for i := 1; i < k; i++ {
// 			if m[key+i] < count {
// 				return false
// 			}
// 			m[key+i] -= count
// 		}
// 	}

// 	return true
// }

// Junk
// func isPossibleDivide(nums []int, k int) bool {
// 	if len(nums)%k != 0 {
// 		return false
// 	}

// 	// O(log n) time, O(?) space
// 	sort.Ints(nums)
// 	for i := 0; i < len(nums); i++ {
// 		num, count := nums[i], 1
// 		for i < len(nums) {
// 			i++
// 			if nums[i] != num {
// 				break
// 			}
// 			count++
// 		}
// 		for i := 1; i < k; i++ {
// 			expect, actual := num+i, 0
// 			for i < len(nums) {
// 				if nums[i] != expect {
// 					break
// 				}
// 				actual++
// 				i++
// 			}
// 			if actual != count {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }
