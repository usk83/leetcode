// 532. K-diff Pairs in an Array
// https://leetcode.com/problems/k-diff-pairs-in-an-array/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(findPairs([]int{1, 2, 3, 4, 5}, -1))
	pp.Println(0)
	pp.Println("=========================================")
}

// The pairs (i, j) and (j, i) count as the same pair.
// The length of the array won't exceed 10,000.
// All the integers in the given input belong to the range: [-1e7, 1e7].
// Two Pointers Solution
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	sort.Ints(nums)
	length, count, small, big := len(nums), 0, 0, 1
LOOP:
	for big < length {
		if target := nums[small] + k; target > nums[big] {
			for cur := nums[big]; nums[big] == cur; big++ {
				if big == length-1 {
					break LOOP
				}
			}
		} else {
			if target == nums[big] {
				count++
			}
			for cur := nums[small]; nums[small] == cur; small++ {
				if small == length-1 {
					break LOOP
				}
			}
			if small >= big {
				big = small + 1
			}
		}
	}
	return count
}

// func findPairs(nums []int, k int) int {
//  if k < 0 {
//    return 0
//  }

//  sort.Ints(nums)

//  memo := make(map[int]int, len(nums))
//  var pairsCount int
//  for i := len(nums) - 1; i >= 0; i-- {
//    memo[nums[i]]++
//    if count := memo[nums[i]]; count == 1 {
//      if k != 0 {
//        if _, found := memo[nums[i]+k]; found {
//          pairsCount++
//        }
//      }
//    } else if count == 2 {
//      if k == 0 {
//        pairsCount++
//      }
//    }
//  }
//  return pairsCount
// }

// func findPairs(nums []int, k int) int {
//  if k < 0 {
//    return 0
//  }

//  numsCountMap := map[int]int{}
//  for _, num := range nums {
//    numsCountMap[num]++
//  }

//  var pairsCount int
//  for num, count := range numsCountMap {
//    if k == 0 {
//      if count >= 2 {
//        pairsCount++
//      }
//    } else {
//      if _, found := numsCountMap[num+k]; found {
//        pairsCount++
//      }
//    }
//  }
//  return pairsCount
// }
