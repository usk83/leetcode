// 5357. Design a Stack With Increment Operation
// https://leetcode.com/contest/weekly-contest-180/problems/design-a-stack-with-increment-operation/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

type CustomStack struct {
	maxSize int
	data    []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		maxSize: maxSize,
		data:    make([]int, 0, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.data) != this.maxSize {
		this.data = append(this.data, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.data) != 0 {
		num := this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
		return num
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(this.data); i++ {
		this.data[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
