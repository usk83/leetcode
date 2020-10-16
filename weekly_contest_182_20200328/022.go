// 5369. Count Number of Teams
// https://leetcode.com/contest/weekly-contest-182/problems/count-number-of-teams/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numTeams([]int{2, 5, 3, 4, 1}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(numTeams([]int{2, 1, 3}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(numTeams([]int{1, 2, 3, 4}))
	pp.Println(4)
	pp.Println("=========================================")
}

// func numTeams(rating []int) int {
// 	var count int
// 	for i := 0; i < len(rating); i++ {
// 		for j := i + 1; j < len(rating); j++ {
// 			for k := j + 1; k < len(rating); k++ {
// 				if (rating[i] < rating[j] && rating[j] < rating[k]) || (rating[i] > rating[j] && rating[j] > rating[k]) {
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(^uint(0) >> 1)
	MinInt  = -MaxInt - 1
)

func numTeams(rating []int) int {
	// 自分より大きいやつ、小さいやつが右にいくつあるか？
	maxes := make([]int, len(rating))
	mins := make([]int, len(rating))

	type sol struct {
		index  int
		rating int
	}

	soldiers := make([]sol, len(rating))
	for i, rat := range rating {
		soldiers[i] = sol{i, rat}
	}

	sort.Slice(soldiers, func(i, j int) bool {
		return soldiers[i].rating < soldiers[j].rating
	})

	for i := range soldiers {
		var max int
		for j := range soldiers[i+1:] {
			if soldiers[i].index > soldiers[i+1+j].index {
				continue
			}
			if soldiers[i].rating < soldiers[i+1+j].rating {
				max++
			}
		}
		maxes[soldiers[i].index] = max
	}

	for i := range soldiers {
		var min int
		for j := range soldiers[i+1:] {
			if soldiers[len(rating)-1-i].index > soldiers[len(rating)-1-i-1-j].index {
				continue
			}
			if soldiers[len(rating)-1-i].rating > soldiers[len(rating)-1-i-1-j].rating {
				min++
			}
		}
		mins[soldiers[len(rating)-1-i].index] = min
	}

	var count int
	for i := 0; i < len(rating); i++ {
		for j := i + 1; j < len(rating); j++ {
			if rating[i] < rating[j] {
				count += maxes[j]
			} else {
				count += mins[j]
			}
		}
	}
	return count
}

// // レーティング大きい順か小さい順
// // 何チーム？
// func numTeams(rating []int) int {
// 	// 自分のところまで
// 	// 	min
// 	// 		カウント 1 いくつ？
// 	// 		カウント 2 いくつ？
// 	// 	max
// 	// 		カウント 1 いくつ？
// 	// 		カウント 2 いくつ？

// 	status := struct {
// 		max1 int
// 		max2 []int
// 		min1 int
// 		min2 []int
// 	}{
// 		max1: rating[0],
// 		max2: []int{},
// 		min1: rating[0],
// 		min2: []int{},
// 	}

// 	var count int
// 	for _, soldier := range rating[1:] {
// 		if soldier < status.max1 {
// 			status.max1 = soldier
// 		} else if soldier > status.max1 {
// 			status.max2 = append(status.max2, soldier)
// 		}
// 		for _, cmp := range status.max2 {
// 			if soldier > cmp {
// 				fmt.Println([2]int{cmp, soldier})
// 				count++
// 			}
// 		}

// 		if soldier > status.min1 {
// 			status.min1 = soldier
// 		} else if soldier < status.min1 {
// 			status.min2 = append(status.min2, soldier)
// 		}
// 		for _, cmp := range status.min2 {
// 			if soldier < cmp {
// 				fmt.Println([2]int{cmp, soldier})
// 				count++
// 			}
// 		}

// 		// if soldier < status.max1 {
// 		// 	status.max1 = soldier
// 		// } else if soldier > status.max1 && soldier < status.max2 {
// 		// 	status.max2 = soldier
// 		// } else if soldier > status.max2 {
// 		// 	fmt.Println([3]int{status.max1, status.max2, soldier})
// 		// 	count++
// 		// }

// 		// if soldier > status.min1 {
// 		// 	status.min1 = soldier
// 		// } else if soldier < status.min1 && soldier > status.min2 {
// 		// 	status.min2 = soldier
// 		// } else if soldier < status.min2 {
// 		// 	fmt.Println([3]int{status.min1, status.min2, soldier})
// 		// 	count++
// 		// }
// 	}
// 	return count
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
