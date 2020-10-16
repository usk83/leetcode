package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	// fmt.Println(true)
	// pp.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
	// fmt.Println(false)
	pp.Println(checkStraightLine([][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}}))
	fmt.Println(true)
}

func checkStraightLine(coordinates [][]int) bool {
	xDiff := coordinates[1][0] - coordinates[0][0]
	yDiff := coordinates[1][1] - coordinates[0][1]

	// pp.Println(xDiff, yDiff)

	ratio := float64(0)
	if xDiff == 0 && yDiff == 0 {
		ratio = 0
	} else if xDiff == 0 {
		ratio = 1
	} else if yDiff == 0 {
		ratio = 0
	} else {
		ratio = float64(xDiff) / float64(yDiff)
	}

	// pp.Println(ratio)

	for i := 1; i < len(coordinates)-1; i++ {
		inXDiff := coordinates[i+1][0] - coordinates[i][0]
		inYDiff := coordinates[i+1][1] - coordinates[i][1]

		// pp.Println(inXDiff, inYDiff)

		inRatio := float64(0)
		if inXDiff == 0 && inYDiff == 0 {
			inRatio = 0
		} else if inXDiff == 0 {
			inRatio = 1
		} else if inYDiff == 0 {
			inRatio = 0
		} else {
			inRatio = float64(inXDiff) / float64(inYDiff)
		}

		// pp.Println(inRatio)

		if ratio != inRatio {
			return false
		}
	}

	return true
}
