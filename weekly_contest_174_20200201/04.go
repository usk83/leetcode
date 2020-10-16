// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(maxJumps([]int{3, 3, 3, 3, 3}, 3))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(maxJumps([]int{7, 6, 5, 4, 3, 2, 1}, 1))
	pp.Println(7)
	pp.Println("=========================================")
	pp.Println(maxJumps([]int{7, 1, 7, 1, 7, 1}, 2))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(maxJumps([]int{66}, 1))
	pp.Println(1)
	pp.Println("=========================================")
}

// 1 <= arr.length <= 1000
// 1 <= arr[i] <= 10^5
// 1 <= d <= arr.length
func maxJumps(arr []int, d int) int {
	// area
	cache := make(map[int]int, len(arr))
	var jump func(cur int) int
	jump = func(cur int) int {
		if _, ok := cache[cur]; ok {
			return cache[cur]
		}

		start := max(cur-d, 0)
		end := min(cur+d, len(arr)-1)

		leftNext := -1
		rightNext := -1
		leftIndices := []int{}
		rightIndices := []int{}
		for i := cur - 1; i >= start; i-- {
			if arr[i] >= arr[cur] {
				break
			}
			if arr[i] < leftNext {
				continue
			}
			if arr[i] == leftNext {
				leftIndices = append(leftIndices, i)
			} else {
				leftNext = arr[i]
				leftIndices = []int{i}
			}
		}
		for i := cur + 1; i <= end; i++ {
			if arr[i] >= arr[cur] {
				break
			}
			if arr[i] < rightNext {
				continue
			}
			if arr[i] == rightNext {
				rightIndices = append(rightIndices, i)
			} else {
				rightNext = arr[i]
				rightIndices = []int{i}
			}
		}

		// fmt.Println(indices)

		prev := 0
		for _, index := range leftIndices {
			var val int
			if _, ok := cache[index]; ok {
				val = cache[index]
			} else {
				val = jump(index)
				cache[index] = val
				// pp.Println(index, cache[index])
			}
			prev = max(prev, val)
		}

		for _, index := range rightIndices {
			var val int
			if _, ok := cache[index]; ok {
				val = cache[index]
			} else {
				val = jump(index)
				cache[index] = val
				// pp.Println(index, cache[index])
			}
			prev = max(prev, val)
		}

		return prev + 1
	}

	for cur := range arr {
		if _, ok := cache[cur]; !ok {
			cache[cur] = jump(cur)
			// pp.Println(cur, cache[cur])
		}
	}

	biggest := 0
	for _, c := range cache {
		if c > biggest {
			biggest = c
		}
	}

	return biggest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
