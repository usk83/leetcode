// 1005. Maximize Sum Of Array After K Negations
// https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	arr1 := []int{4, 2, 3}
	pp.Println(largestSumAfterKNegations(arr1, 1))
	pp.Println(5)
	fmt.Println(arr1)
	pp.Println("=========================================")
	arr2 := []int{3, -1, 0, 2}
	pp.Println(largestSumAfterKNegations(arr2, 3))
	pp.Println(6)
	fmt.Println(arr2)
	pp.Println("=========================================")
	arr3 := []int{2, -3, -1, 5, -4}
	pp.Println(largestSumAfterKNegations(arr3, 2))
	pp.Println(13)
	fmt.Println(arr3)
	pp.Println("=========================================")
	arr4 := []int{-8, 3, -5, -3, -5, -2}
	pp.Println(largestSumAfterKNegations(arr4, 6))
	pp.Println(22)
	fmt.Println(arr4)
	pp.Println("=========================================")
}

// 1 <= A.length <= 10000
// 1 <= K <= 10000
// -100 <= A[i] <= 100

const MAX_VALUE = 100

func largestSumAfterKNegations(A []int, K int) int {
	sum, minNum, minIndex, minus := 0, MAX_VALUE+1, 0, [][2]int{}
	for i, num := range A {
		sum += num
		if num >= 0 {
			if num < minNum {
				minNum = num
				minIndex = i
			}
		} else {
			minus = append(minus, [2]int{i, num})
		}
	}

	sort.Slice(minus, func(i, j int) bool { return minus[i][1] < minus[j][1] })

	for i := 0; i < K; i++ {
		if i < len(minus) {
			sum -= minus[i][1] * 2
			A[minus[i][0]] *= -1
			continue
		}

		if (K-i)&1 == 0 {
			break
		}

		if len(minus) != 0 {
			if m := minus[len(minus)-1][1] * -1; m < minNum {
				minNum = m
				minIndex = minus[len(minus)-1][0]
			}
		}

		sum -= minNum * 2
		A[minIndex] *= -1
	}

	return sum
}
