// 11. Container With Most Water
// https://leetcode.com/problems/container-with-most-water/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	pp.Println(49)
	pp.Println("=========================================")
	pp.Println(maxArea([]int{1, 2, 1}))
	pp.Println(2)
	pp.Println("=========================================")
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := min(height[left], height[right]) * right
	for left < right {
		diff := height[left] - height[right]
		if diff <= 0 {
			left++
		}
		if diff >= 0 {
			right--
		}
		area = max(area, min(height[left], height[right])*(right-left))
	}
	return area
}

// func maxArea(height []int) int {
// 	left, right := 0, len(height)-1
// 	area := min(height[left], height[right]) * right
// 	for left < right {
// 		switch {
// 		case height[left] < height[right]:
// 			orig := left
// 			for left = left + 1; left < right; left++ {
// 				if height[left] > height[orig] {
// 					break
// 				}
// 			}
// 		case height[left] > height[right]:
// 			orig := right
// 			for right = right - 1; right > left; right-- {
// 				if height[right] > height[orig] {
// 					break
// 				}
// 			}
// 		default: // height[left] == height[right]
// 			left, right = left+1, right-1
// 		}
// 		area = max(area, min(height[left], height[right])*(right-left))
// 	}
// 	return area
// }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func maxArea(height []int) int {
// 	var left, area int
// 	for i := 1; i < len(height); i++ {
// 		if width := i - left; height[i] > height[left] {
// 			area = max(area, height[left]*width)
// 			left = i
// 		} else {
// 			area = max(area, height[i]*width)
// 		}
// 	}
// 	return area
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
