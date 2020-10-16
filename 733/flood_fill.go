// 733. Flood Fill
// https://leetcode.com/problems/flood-fill/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	fmt.Println([][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}})
	pp.Println("=========================================")
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
	fmt.Println([][]int{{0, 0, 0}, {0, 1, 1}})
	pp.Println("=========================================")
}

// The length of image and image[0] will be in the range [1, 50].
// The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
// The value of each color in image[i][j] and newColor will be an integer in [0, 65535].
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var fill func(int, int)
	fill = func(r, c int) {
		curColor := image[r][c]
		image[r][c] = -1

		if r > 0 && image[r-1][c] == curColor {
			fill(r-1, c)
		}
		if c < len(image[0])-1 && image[r][c+1] == curColor {
			fill(r, c+1)
		}
		if r < len(image)-1 && image[r+1][c] == curColor {
			fill(r+1, c)
		}
		if c > 0 && image[r][c-1] == curColor {
			fill(r, c-1)
		}

		image[r][c] = newColor
	}

	fill(sr, sc)

	return image
}
