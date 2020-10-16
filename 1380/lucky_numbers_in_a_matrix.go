// 1380. Lucky Numbers in a Matrix
// https://leetcode.com/problems/lucky-numbers-in-a-matrix/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
	fmt.Println([]int{15})
	pp.Println("=========================================")
	fmt.Println(luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}))
	fmt.Println([]int{12})
	pp.Println("=========================================")
	fmt.Println(luckyNumbers([][]int{{7, 8}, {1, 2}}))
	fmt.Println([]int{7})
	pp.Println("=========================================")
}

func luckyNumbers(matrix [][]int) []int {
	rowMins := Ints(MaxInt, len(matrix))
	colMaxes := Ints(MinInt, len(matrix[0]))
	for y, row := range matrix {
		for x, num := range row {
			rowMins[y] = Min(rowMins[y], num)
			colMaxes[x] = Max(colMaxes[x], num)
		}
	}

	set := NewIntSetFromSlice(rowMins)

	nums := []int{}
	for _, num := range colMaxes {
		if set.Contains(num) {
			nums = append(nums, num)
		}
	}

	return nums
}

func Ints(val, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = val
	}
	return s
}

type IntSet struct {
	values map[int]struct{}
}

func NewIntSetFromSlice(ints []int) IntSet {
	intSet := IntSet{
		values: map[int]struct{}{},
	}
	for _, n := range ints {
		intSet.values[n] = struct{}{}
	}
	return intSet
}

func (s *IntSet) Contains(n int) bool {
	_, ok := s.values[n]
	return ok
}

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(^uint(0) >> 1)
	MinInt  = -MaxInt - 1
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
