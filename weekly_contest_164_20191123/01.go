// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"math"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
	pp.Println(7)
	pp.Println("=========================================")
	pp.Println(minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))
	pp.Println(5)
	pp.Println("=========================================")
}

// points.length == n
// 1 <= n <= 100
// points[i].length == 2
// -1000 <= points[i][0], points[i][1] <= 1000
func minTimeToVisitAllPoints(points [][]int) int {
	// まずななめ
	// そのあとその縦か横
	var time int
	for i, point := range points[1:] {
		x := points[i][0] - point[0]
		y := points[i][1] - point[1]

		x = int(math.Abs(float64(x)))
		y = int(math.Abs(float64(y)))

		if x > y {
			time += y
			time += x - y
		} else {
			time += x
			time += y - x
		}
	}

	return time
}
