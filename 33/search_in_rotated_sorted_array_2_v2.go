// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(search([]int{}, 5))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(search([]int{5}, 5))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(search([]int{5, 1}, 5))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(search([]int{1, 5}, 5))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(search([]int{4, 5, 6, 7, -2, -1, 0, 1, 2}, 0))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(search([]int{6, 7, 0, 1, 2, 4, 5}, 3))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(search([]int{4, 5, 6, 7, 0, 1}, 0))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(search([]int{100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 1))
	pp.Println(1)
	pp.Println("=========================================")
}

func search(nums []int, target int) int {
	lastIndex := len(nums) - 1
	if lastIndex < 0 {
		return -1
	}
	isBigTarget := target > nums[lastIndex]
	start, end := 0, lastIndex
	for start <= end {
		middle := start + (end-start)/2 // overflow awareness
		if nums[middle] == target {
			return middle
		}
		isBigMiddle := nums[middle] > nums[lastIndex]
		isGreater := nums[middle] > target

		// 中間の値が小さいとき
		// 	中間の値が小さい方である場合
		// 		ターゲットが小さい方にあるなら右
		// 		ターゲットが大きい方にあるなら左
		// 	中間の値が大きい方である場合
		// 		右
		// 中間の値が大きいとき
		// 	中間の値が小さい方である場合
		// 		左
		// 	中間の値が大きい方である場合
		// 		ターゲットが小さい方にあるなら右
		// 		ターゲットが大きい方にあるなら左

		if !isGreater && !isBigMiddle && !isBigTarget ||
			!isGreater && isBigMiddle && isBigTarget ||
			!isGreater && isBigMiddle && !isBigTarget ||
			isGreater && isBigMiddle && !isBigTarget {
			// 中間の値が小さく、中間の値が小さい方にあり、ターゲットが小さいほうにある
			// 中間の値が小さく、中間の値が大きい方にある
			// 中間の値が大きく、中間の値が大きい方にあり、ターゲットが小さいほうにある
			start = middle + 1
		} else {
			// !isGreater && !isBigMiddle && isBigTarget ||
			// 	isGreater && !isBigMiddle && isBigTarget ||
			// 	isGreater && !isBigMiddle && !isBigTarget ||
			// 	isGreater && isBigMiddle && isBigTarget

			// 中間の値が小さく、中間の値が小さい方にあり、ターゲットが大きいほうにある
			// 中間の値が大きく、中間の値が小さい方にある
			// 中間の値が大きく、中間の値が大きい方にある、ターゲットが大きいほうにある
			end = middle - 1
		}
	}
	return -1
}
