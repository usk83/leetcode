// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"math"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(getMinDistSum([][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}}))
	pp.Println(4.00000)
	pp.Println("=========================================")
	pp.Println(getMinDistSum([][]int{{1, 1}, {3, 3}}))
	pp.Println(2.82843)
	pp.Println("=========================================")
	pp.Println(getMinDistSum([][]int{{1, 1}}))
	pp.Println(0.00000)
	pp.Println("=========================================")
	pp.Println(getMinDistSum([][]int{{1, 1}, {0, 0}, {2, 0}}))
	pp.Println(2.73205)
	pp.Println("=========================================")
	pp.Println(getMinDistSum([][]int{{0, 1}, {3, 2}, {4, 5}, {7, 6}, {8, 9}, {11, 1}, {2, 12}}))
	pp.Println(32.94036)
	pp.Println("=========================================")
}

func getMinDistSum(positions [][]int) float64 {
	centerX, centerY := 50.0, 50.0
	for i := 1; i <= 1000; i++ {
		grad := 1.0 / float64(i)
		var nextX, nextY float64
		minSum := math.MaxFloat64
		for x := -50; x <= 50; x++ {
			for y := -50; y <= 50; y++ {
				testX, testY := centerX+float64(x)*grad, centerY+float64(y)*grad

				var sum float64
				for _, pos := range positions {
					curX, curY := float64(pos[0]), float64(pos[1])
					diffX := math.Abs(testX - curX)
					diffY := math.Abs(testY - curY)
					sum += math.Sqrt(diffX*diffX + diffY*diffY)
				}
				if sum < minSum {
					minSum = sum
					nextX, nextY = testX, testY
				}
			}
		}

		centerX, centerY = nextX, nextY
	}

	var sum float64
	for _, pos := range positions {
		curX, curY := float64(pos[0]), float64(pos[1])
		diffX := math.Abs(centerX - curX)
		diffY := math.Abs(centerY - curY)
		sum += math.Sqrt(diffX*diffX + diffY*diffY)
	}
	return sum

	return 0
}
