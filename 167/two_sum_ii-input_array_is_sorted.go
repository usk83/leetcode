// 167. Two Sum II - Input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println([]int{1, 2})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{0, 0, 3, 4}, 0))
	fmt.Println([]int{1, 2})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{5, 25, 75}, 100))
	fmt.Println([]int{2, 3})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{3, 24, 50, 79, 88, 150, 345}, 200))
	fmt.Println([]int{3, 6})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997}, 542))
	fmt.Println([]int{24, 32})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 31, 45, 50}, 61))
	fmt.Println([]int{12, 13})
	pp.Println("=========================================")
}

// func twoSum(numbers []int, target int) []int {
// 	index1, index2 := 0, len(numbers)-1
// 	for {
// 		if sum := numbers[index1] + numbers[index2]; sum == target {
// 			break
// 		} else if sum < target {
// 			// if index1 == 10 {
// 			// 	break
// 			// }
// 			// break
// 			index1++
// 		} else {
// 			if index2 == 5 {
// 				break
// 			}
// 			index2--
// 		}
// 	}
// 	pp.Println("index1:", index1, "index2:", index2)
// 	pp.Println("sum:", numbers[index1]+numbers[index2])
// 	return []int{index1 + 1, index2 + 1}
// }

func twoSum(numbers []int, target int) []int {
	index, diff := 0, len(numbers)-1
	for numbers[index]+numbers[index+diff] != target {
		// index += sort.Search(diff-index, func(i int) bool {
		// 	return target <= numbers[index:][i]+numbers[index:][diff-index]
		// })
		// diff = index + sort.Search(diff-index-1, func(i int) bool {
		// 	return target < numbers[index]+numbers[index+2:][i]
		// }) + 1
	}
	return []int{index + 1, index + diff + 1}
}

func twoSum(numbers []int, target int) []int {
	index1, index2 := 0, len(numbers)-1
	for numbers[index1]+numbers[index2] != target {
		index1 += sort.Search(index2-index1, func(i int) bool {
			return target <= numbers[index1:][i]+numbers[index1:][index2-index1]
		})
		index2 = index1 + sort.Search(index2-index1-1, func(i int) bool {
			return target < numbers[index1]+numbers[index1+2:][i]
		}) + 1
	}
	return []int{index1 + 1, index2 + 1}
}

// func twoSum(numbers []int, target int) []int {
// 	index1, index2 := 0, len(numbers)-1
// 	for numbers[index1]+numbers[index2] != target {
// 		index1 += sort.Search(index2-index1, func(i int) bool {
// 			return target <= numbers[index1:][i]+numbers[index1:][index2-index1]
// 		})
// 		// if index1 == 2 {
// 		// 	pp.Println("先頭: ", numbers[index1+2:][0])
// 		// 	pp.Println("現在位置: ", numbers[index2])
// 		// 	pp.Println("長さ: ", index2-index1-1)
// 		// 	// fmt.Println("残り: ", numbers[index1+2:index2-index1-1])
// 		// 	break
// 		// }
// 		index2 = index1 + sort.Search(index2-index1-1, func(i int) bool {
// 			return target < numbers[index1]+numbers[index1+2:][i]
// 		}) + 1
// 	}
// 	// pp.Println("index1:", index1, "index2:", index2)
// 	// pp.Println("sum:", numbers[index1]+numbers[index2])
// 	return []int{index1 + 1, index2 + 1}
// }

// func twoSum(numbers []int, target int) []int {
// 	index1, index2 := 0, len(numbers)-1
// 	for numbers[index1]+numbers[index2] != target {
// 		index1 += sort.Search(index2-index1, func(i int) bool {
// 			return target <= numbers[index1:][i]+numbers[index1:][index2-index1]
// 		})
// 		// if index1 == 2 {
// 		// 	pp.Println("先頭: ", numbers[index1+2:][0])
// 		// 	pp.Println("現在位置: ", numbers[index2])
// 		// 	pp.Println("長さ: ", index2-index1-1)
// 		// 	// fmt.Println("残り: ", numbers[index1+2:index2-index1-1])
// 		// 	break
// 		// }
// 		index2 = index1 + sort.Search(index2-index1-1, func(i int) bool {
// 			return target < numbers[index1]+numbers[index1+2:][i]
// 		}) + 1
// 	}
// 	// pp.Println("index1:", index1, "index2:", index2)
// 	// pp.Println("sum:", numbers[index1]+numbers[index2])
// 	return []int{index1 + 1, index2 + 1}
// }

/*
 * 1. Two pointers
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 */
// func twoSum(numbers []int, target int) []int {
// 	index1, index2 := 0, len(numbers)-1
// 	for {
// 		if sum := numbers[index1] + numbers[index2]; sum == target {
// 			break
// 		} else if sum < target {
// 			index1++
// 		} else {
// 			index2--
// 		}
// 	}
// 	return []int{index1 + 1, index2 + 1}
// }
