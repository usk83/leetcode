// 1232. Check If It Is a Straight Line
// https://leetcode.com/problems/check-if-it-is-a-straight-line/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(checkStraightLine([][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}}))
	pp.Println(true)
	pp.Println("=========================================")
}

// 2 <= coordinates.length <= 1000
// coordinates[i].length == 2
// -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
// coordinates contains no duplicate point.
func checkStraightLine(coordinates [][]int) bool {
	baseXDiff := coordinates[1][0] - coordinates[0][0]
	baseYDiff := coordinates[1][1] - coordinates[0][1]
	for i := 1; i < len(coordinates)-1; i++ {
		curXDiff := coordinates[i+1][0] - coordinates[i][0]
		curYDiff := coordinates[i+1][1] - coordinates[i][1]
		// if baseYDiff:baseXDiff is not equal to curYDiff:curXDiff, return false
		if baseYDiff*curXDiff != baseXDiff*curYDiff {
			return false
		}
	}
	return true
}
