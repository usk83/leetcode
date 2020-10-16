// 1232. Check If It Is a Straight Line
// https://leetcode.com/problems/check-if-it-is-a-straight-line/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 8
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3323/
func main() {
	pp.Println("=========================================")
	pp.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
	pp.Println(false)
	pp.Println("=========================================")
}

func checkStraightLine(coordinates [][]int) bool {
	x, y := coordinates[0][0]-coordinates[1][0], coordinates[0][1]-coordinates[1][1]
	for _, coordinate := range coordinates[2:] {
		xx, yy := coordinates[0][0]-coordinate[0], coordinates[0][1]-coordinate[1]
		if x*yy != y*xx {
			return false
		}
	}
	return true
}
