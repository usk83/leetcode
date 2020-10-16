// 219. Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	pp.Println(false)
	pp.Println("=========================================")
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}

	pairs := make([][2]int, len(nums))
	for i, num := range nums {
		pairs[i] = [2]int{num, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0] || (pairs[i][0] == pairs[j][0] && pairs[i][1] < pairs[j][1])
	})

	for i, pair := range pairs[1:] {
		if pairs[i][0] == pair[0] && pair[1]-pairs[i][1] <= k {
			return true
		}
	}

	return false
}

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	m := map[int]int{}
// 	for i, num := range nums {
// 		if prev, ok := m[num]; ok {
// 			if i-prev <= k {
// 				return true
// 			}
// 		}
// 		m[num] = i
// 	}
// 	return false
// }
