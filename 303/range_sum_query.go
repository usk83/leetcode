// 303. Range Sum Query - Immutable
// https://leetcode.com/problems/range-sum-query-immutable/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	pp.Println("=========================================")
	pp.Println(numArray.SumRange(0, 2))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numArray.SumRange(2, 5))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(numArray.SumRange(0, 5))
	pp.Println(-3)
	pp.Println("=========================================")
}

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{
		sums: make([]int, len(nums)+1),
	}
	for i, num := range nums {
		numArray.sums[i+1] = numArray.sums[i] + num
	}
	return numArray
}

func (na *NumArray) SumRange(i int, j int) int {
	// 実際にはあったほうがいいかも
	// if i < 0 || j < 0 || i >= len(na.sums) || j >= len(na.sums) {
	// 	Error handling
	// }
	return na.sums[j+1] - na.sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
