// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"strconv"

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
}

// [1,0,1]
// 5
// [0]
// 0
// [1]
// 1
// [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
// 18880
// [0,0]
// 0

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	bin := ""
	for head != nil {
		bin += strconv.Itoa(head.Val)
		head = head.Next
	}
	num, _ := strconv.ParseInt(bin, 2, 64)
	return int(num)
}
