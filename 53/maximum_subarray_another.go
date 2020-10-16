// 53. Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{-1}))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{1, 2}))
	pp.Println(3)
	pp.Println("=========================================")
}

// another solution using the divide and conquer approach
// https://leetcode.com/problems/maximum-subarray/discuss/20200/Share-my-solutions-both-greedy-and-divide-and-conquer
type part struct {
	left  int
	max   int
	right int
	whole int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArrayOfPart(nums []int) *part {
	n := len(nums)
	if n == 1 {
		return &part{
			left:  nums[0],
			max:   nums[0],
			right: nums[0],
			whole: nums[0],
		}
	}
	half := n / 2
	left := maxSubArrayOfPart(nums[:half])
	right := maxSubArrayOfPart(nums[half:])
	return &part{
		left:  max(left.left, left.whole+right.left),
		max:   max(left.right+right.left, max(left.max, right.max)),
		right: max(right.right, left.right+right.whole),
		whole: left.whole + right.whole,
	}
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := maxSubArrayOfPart(nums)
	return result.max
}
