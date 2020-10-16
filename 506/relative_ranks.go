// 506. Relative Ranks
// https://leetcode.com/problems/relative-ranks/
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	for i, arg := range os.Args {
		pp.Println(i, arg)
	}

	// pp.Println("=========================================")
	// fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	// fmt.Println([]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"})
	pp.Println("=========================================")
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
	fmt.Println([]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"})
	pp.Println("=========================================")
}

/*
 * v2.
 * ref: https://leetcode.com/problems/relative-ranks/discuss/98468/Easy-Java-Solution-Sorting.


 *   - Time Complexity: O(NlogN)
 *   - Space Complexity: O(N)
 * > Runtime: 12 ms, faster than 96.30%
 * > Memory Usage: 6 MB, less than 100.00%


 */
func findRelativeRanks(nums []int) []string {
	indices := make([]int, len(nums))
	for i := range indices {
		indices[i] = i
	}

	pp.Println("=== DEBUG START ======================================")

	// fmt.Println(indices)
	fmt.Println(nums)

	sort.SliceStable(indices, func(i, j int) bool {
		pp.Printf("%s: %s, %s: %s\n", indices[i], nums[i], indices[j], nums[j])
		return nums[i] > nums[j]
	})

	fmt.Println(indices)

	pp.Println("=== DEBUG END ========================================")

	ranks := make([]string, len(nums))
	for rank, index := range indices {
		switch rank {
		case 0:
			ranks[index] = "Gold Medal"
		case 1:
			ranks[index] = "Silver Medal"
		case 2:
			ranks[index] = "Bronze Medal"
		default:
			ranks[index] = strconv.Itoa(rank + 1)
		}
	}
	return ranks
}

/*
 * v1. The first solution
 *   - Time Complexity: O(NlogN)
 *   - Space Complexity: O(N)
 * > Runtime: 12 ms, faster than 96.30%
 * > Memory Usage: 6 MB, less than 100.00%
 */
// func findRelativeRanks(nums []int) []string {
// 	scores := make([]int, len(nums))
// 	copy(scores, nums)
// 	sort.Slice(scores, func(i, j int) bool {
// 		return scores[i] > scores[j]
// 	})

// 	rankMap := map[int]int{}
// 	for rank, score := range scores {
// 		rankMap[score] = rank + 1
// 	}

// 	ranks := make([]string, len(nums))
// 	for i, score := range nums {
// 		switch rank := rankMap[score]; rank {
// 		case 1:
// 			ranks[i] = "Gold Medal"
// 		case 2:
// 			ranks[i] = "Silver Medal"
// 		case 3:
// 			ranks[i] = "Bronze Medal"
// 		default:
// 			ranks[i] = strconv.Itoa(rank)
// 		}
// 	}
// 	return ranks
// }
