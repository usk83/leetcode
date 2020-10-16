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
	length := len(nums)
	pivot := length
	for start, end := 1, length-1; start <= end; {
		current := start + (end-start)/2 // overflow awareness
		prev := current - 1
		if nums[prev] > nums[current] {
			pivot = current
			break
		}
		if nums[current] >= nums[start-1] {
			start = current + 1
		} else {
			end = prev
		}
	}

	smallerHalf := length - pivot
	toRotatedIndex := func(index int) int {
		if index < smallerHalf {
			return pivot + index
		}
		return index - smallerHalf
	}

	targetIndex := -1
	for start, end := 0, length-1; start <= end; {
		current := start + (end-start)/2 // overflow awareness
		rotatedIndex := toRotatedIndex(current)
		if currentValue := nums[rotatedIndex]; currentValue == target {
			targetIndex = rotatedIndex
			break
		} else if currentValue < target {
			start = current + 1
		} else {
			end = current - 1
		}
	}
	return targetIndex
}
