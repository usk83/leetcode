// 658. Find K Closest Elements
// https://leetcode.com/problems/find-k-closest-elements/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElements([]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}, 9, 4))
}

func findClosestElements(arr []int, k int, x int) []int {
	ret := make([]int, 0, k) // 正解のスライス

	// xに一番近い値のインデックスを調査する
	cur := xindex(arr, x)
	ret = append(ret, arr[cur])

	head, tail := cur+1, cur-1
	for i := (k - 1); i > 0; i-- {
		if tail < 0 {
			ret = append(ret, arr[head])
			head++
			continue
		} else if head > len(arr)-1 {
			ret = prepend(ret, arr[tail])
			tail--
			continue
		}

		less, more := arr[tail], arr[head]
		diffLess, diffMore := math.Abs(float64(x-less)), math.Abs(float64(x-more))
		if diffLess <= diffMore {
			ret = prepend(ret, less)
			tail--
		} else {
			ret = append(ret, more)
			head++
		}
	}

	return ret
}

func prepend(arr []int, x int) []int {
	arr = append(arr, 0)
	for i := 0; i < len(arr); i++ {
		arr[i], x = x, arr[i]
	}
	return arr
}

func xindex(arr []int, x int) int {
	return finxXIndex(arr, x, 0, len(arr)-1)
}

func finxXIndex(arr []int, x, start, end int) int {
	if start == end {
		return start
	} else if end-start == 1 {
		diffLess, diffMore := math.Abs(float64(x-arr[start])), math.Abs(float64(x-arr[end]))
		if diffLess <= diffMore {
			return start
		} else {
			return end
		}
	}

	index := (start + end) / 2
	mid := arr[index]

	if diff := x - mid; diff == 0 {
		return index
	} else if diff < 0 {
		return finxXIndex(arr, x, start, index)
	} else {
		return finxXIndex(arr, x, index, end)
	}
}
