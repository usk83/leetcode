// 1105. Filling Bookcase Shelves
// https://leetcode.com/problems/filling-bookcase-shelves/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))
}

func minHeightShelves(books [][]int, shelf_width int) int {
	booksLength := len(books)
	if booksLength == 0 {
		return 0
	}

	dp := make([]map[int][2]int, booksLength)
	for i := 0; i < booksLength; i++ {
		dp[i] = make(map[int][2]int, 1)
	}

	dp[0][books[0][0]] = [2]int{0, books[0][1]}

	for i, book := range books[1:] {
		for currentWidth, prev := range dp[i] {
			index := i + 1
			// 横に置けるなら置く
			if newWidth := currentWidth + book[0]; currentWidth+book[0] <= shelf_width {
				if prev[1] < book[1] {
					dp[index][newWidth] = [2]int{prev[0], book[1]}
				} else {
					dp[index][newWidth] = [2]int{prev[0], prev[1]}
					continue
				}
			}

			// 下に置く
			confirmedHeight := prev[0] + prev[1]
			if foundDP, found := dp[index][book[0]]; found {
				if foundDP[0] <= confirmedHeight {
					continue
				}
			}

			dp[index][book[0]] = [2]int{confirmedHeight, book[1]}
		}
	}

	minHeight := int(^uint(0) >> 1)
	for _, item := range dp[booksLength-1] {
		cur := item[0] + item[1]
		if cur < minHeight {
			minHeight = cur
		}
	}

	return minHeight
}
