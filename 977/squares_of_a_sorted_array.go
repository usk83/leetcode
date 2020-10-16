// 977. Squares of a Sorted Array
// https://leetcode.com/problems/squares-of-a-sorted-array/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println([]int{0, 1, 9, 16, 100})
	pp.Println("=========================================")
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
	fmt.Println([]int{4, 9, 9, 49, 121})
	pp.Println("=========================================")
	fmt.Println(sortedSquares([]int{1, 1, 2, 3, 11}))
	fmt.Println([]int{1, 1, 4, 9, 121})
	pp.Println("=========================================")
	fmt.Println(sortedSquares([]int{-5, -4, -3, -2, -1}))
	fmt.Println([]int{1, 4, 9, 16, 25})
	pp.Println("=========================================")
	fmt.Println(sortedSquares([]int{-1, 0, 1, 1}))
	fmt.Println([]int{0, 1, 1, 1})
	pp.Println("=========================================")
}

/*
 * 末尾から詰めていく
 * O(N)
 */
// func sortedSquares(a []int) []int {
// 	squared := make([]int, len(a))
// 	for l, r, i := 0, len(a)-1, len(a)-1; i >= 0; i-- {
// 		if leftVal, rightVal := a[l]*a[l], a[r]*a[r]; leftVal > rightVal {
// 			squared[i] = leftVal
// 			l++
// 		} else {
// 			squared[i] = rightVal
// 			r--
// 		}
// 	}
// 	return squared
// }

/*
 * 先頭から詰めていく
 * O(logN+N)
 */
func findMinPlus(a []int) int {
	start, end := 0, len(a)
	for start < end {
		if mid := (start + end) / 2; a[mid] < 0 {
			start = mid + 1
		} else if mid == 0 || a[mid-1] < 0 {
			return mid
		} else {
			end = mid - 1
		}
	}
	return start
}
func sortedSquares(a []int) []int {
	plus := findMinPlus(a) // O(logN)
	minus := plus - 1
	squares := make([]int, 0, len(a))
	/* O(N) */
	for 0 <= minus && plus < len(a) {
		if -a[minus] < a[plus] {
			squares = append(squares, a[minus]*a[minus])
			minus--
		} else {
			squares = append(squares, a[plus]*a[plus])
			plus++
		}
	}
	for minus >= 0 {
		squares = append(squares, a[minus]*a[minus])
		minus--
	}
	for plus < len(a) {
		squares = append(squares, a[plus]*a[plus])
		plus++
	}
	return squares
}

/*
 * 自乗してソート
 * O(N+NlogN)
 */
// func sortedSquares(a []int) []int {
// 	for index := range a { // O(N)
// 		a[index] *= a[index]
// 	}
// 	sort.Ints(a) // O(NlogN)
// 	return a
// }
