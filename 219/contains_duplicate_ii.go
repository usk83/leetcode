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
	combi := make([][2]int, len(nums))
	for i, num := range nums {
		combi[i] = [2]int{num, i}
	}
	sort.SliceStable(combi, func(i, j int) bool { return combi[i][0] < combi[j][0] })
	for i := range nums[:len(nums)-1] {
		if combi[i][0] == combi[i+1][0] && combi[i+1][1]-combi[i][1] <= k {
			return true
		}
	}
	return false
}

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	m := map[int]int{}
// 	for i, num := range nums {
// 		if index, ok := m[num]; ok {
// 			if i-index <= k {
// 				return true
// 			}
// 		}
// 		m[num] = i
// 	}
// 	return false
// }
