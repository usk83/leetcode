// x. xxx
// https://leetcode.com/problems/xxx/
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

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(^uint(0) >> 1)
	MinInt  = -MaxInt - 1
)

// ��`�ƥ��󥰴󤭤�혤�С�����
// �Υ��`�ࣿ
func numTeams(rating []int) int {
	var count int

	// �Է֤ΤȤ���ޤ�
	// 	min
	// 		������� 1 �����ģ�
	// 		������� 2 �����ģ�
	// 	max
	// 		������� 1 �����ģ�
	// 		������� 2 �����ģ�

	status := struct {
		max [2]int
		min [2]int
	}{
		max: [2]int{rating[0], MinInt},
		min: [2]int{rating[0], MaxInt},
	}

	for _, soldier := range rating[1:] {

	}

	return count
}

func max(x, y) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y) int {
	if x < y {
		return x
	}
	return y
}
