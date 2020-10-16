// 986. Interval List Intersections
// https://leetcode.com/problems/interval-list-intersections/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 23
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3338/
func main() {
	pp.Println("=========================================")
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	fmt.Println([][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}})
	pp.Println("=========================================")
}

// 次に行くときの条件
// 	片方からみて
// 		自分の end が他方の start よりも前のとき
// 			自分を進める
// 		自分の start が他方の end よりも後ろのとき
// 			他方を進める
//
// 	両方を満たす -> intersection がある
// 		大きい start から 小さい end までがintersection
//
// 	小さい end の方を次に進める
//
// 	片方のリストを使い切ったら終了

// 0 <= A.length < 1000
// 0 <= B.length < 1000
// 0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9
func intervalIntersection(A [][]int, B [][]int) [][]int {
	intersections := [][]int{}

	var pointerA, pointerB int
	for pointerA != len(A) && pointerB != len(B) {
		if A[pointerA][1] < B[pointerB][0] {
			pointerA++
			continue
		}

		if B[pointerB][1] < A[pointerA][0] {
			pointerB++
			continue
		}

		intersections = append(intersections, []int{max(A[pointerA][0], B[pointerB][0]), min(A[pointerA][1], B[pointerB][1])})

		if A[pointerA][1] <= B[pointerB][1] {
			pointerA++
		} else {
			pointerB++
		}
	}

	return intersections
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
