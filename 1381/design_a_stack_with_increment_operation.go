// 1381. Design a Stack With Increment Operation
// https://leetcode.com/problems/design-a-stack-with-increment-operation/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	customStack := Constructor(3) // Stack is Empty []
	customStack.Push(1)           // stack becomes [1]
	customStack.Push(2)           // stack becomes [1, 2]
	pp.Println(customStack.Pop()) // return 2 --> Return top of the stack 2, stack becomes [1]
	customStack.Push(2)           // stack becomes [1, 2]
	customStack.Push(3)           // stack becomes [1, 2, 3]
	customStack.Push(4)           // stack still [1, 2, 3], Don't add another elements as size is 4
	customStack.Increment(5, 100) // stack becomes [101, 102, 103]
	customStack.Increment(2, 100) // stack becomes [201, 202, 103]
	pp.Println(customStack.Pop()) // return 103 --> Return top of the stack 103, stack becomes [201, 202]
	pp.Println(customStack.Pop()) // return 202 --> Return top of the stack 102, stack becomes [201]
	pp.Println(customStack.Pop()) // return 201 --> Return top of the stack 101, stack becomes []
	pp.Println(customStack.Pop()) // return -1 --> Stack is empty return -1.
	pp.Println("=========================================")
	pp.Println(2)
	pp.Println(103)
	pp.Println(202)
	pp.Println(201)
	pp.Println(-1)
	pp.Println("=========================================")
}

type CustomStack struct {
	maxSize int
	data    []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		maxSize: maxSize,
		data:    []int{},
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.data) != this.maxSize {
		this.data = append(this.data, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.data) == 0 {
		return -1
	}
	num := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return num
}

func (this *CustomStack) Increment(k int, val int) {
	for i := range this.data[:Min(k, len(this.data))] {
		this.data[i] += val
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// type CustomStack struct {
// 	data []int
// }
//
// func Constructor(maxSize int) CustomStack {
// 	return CustomStack{
// 		data: make([]int, 0, maxSize),
// 	}
// }
//
// func (this *CustomStack) Push(x int) {
// 	if len(this.data) != cap(this.data) {
// 		this.data = append(this.data, x)
// 	}
// }
//
// func (this *CustomStack) Pop() int {
// 	if len(this.data) == 0 {
// 		return -1
// 	}
// 	num := this.data[len(this.data)-1]
// 	this.data = this.data[:len(this.data)-1]
// 	return num
// }
//
// func (this *CustomStack) Increment(k int, val int) {
// 	for i := 0; i < k && i < len(this.data); i++ {
// 		this.data[i] += val
// 	}
// }
