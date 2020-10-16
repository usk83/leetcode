// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(canArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canArrange([]int{1, 2, 3, 4, 5, 6}, 7))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canArrange([]int{1, 2, 3, 4, 5, 6}, 10))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(canArrange([]int{-10, 10}, 2))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canArrange([]int{-1, 1, -2, 2, -3, 3, -4, 4}, 3))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canArrange([]int{-4, -7, 5, 2, 9, 1, 10, 4, -8, -3}, 3))
	pp.Println(true)
	pp.Println("=========================================")
}

func canArrange(nums []int, k int) bool {
	pos, neg := []int{}, []int{}
	for i := range nums {
		if mod := nums[i] % k; mod >= 0 {
			pos = append(pos, mod)
		} else {
			neg = append(neg, mod)
		}
	}

	sort.Ints(pos)
	sort.Sort(sort.Reverse(sort.IntSlice(neg)))

	// pp.Println("=== DEBUG START ======================================")

	for len(pos) >= 2 && pos[0] == 0 && pos[1] == 0 {
		pos = pos[2:]
	}

	// fmt.Println(pos)
	// fmt.Println(neg)

	rem := []int{}

	for len(neg) != 0 && len(pos) != 0 {
		if neg[0]+pos[0] == 0 {
			neg = neg[1:]
			pos = pos[1:]
		} else {
			rem = append(rem, pos[0])
			pos = pos[1:]
		}
	}

	rem = append(rem, pos...)

	// fmt.Println(rem)
	// pp.Println("=== DEBUG END ========================================")

	for len(rem) != 0 {
		sum := rem[0] + rem[len(rem)-1]
		if sum%k != 0 {
			return false
		}
		rem = rem[1 : len(rem)-1]
	}
	return len(rem) == 0 && len(neg) == 0
}
