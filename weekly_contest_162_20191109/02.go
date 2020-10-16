package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(reconstructMatrix(2, 1, []int{1, 1, 1}))
	fmt.Println([][]int{{1, 1, 0}, {0, 0, 1}})
	pp.Println("=========================================")
	fmt.Println(reconstructMatrix(2, 3, []int{2, 2, 1, 1}))
	fmt.Println([][]int{})
	pp.Println("=========================================")
	fmt.Println(reconstructMatrix(5, 5, []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}))
	fmt.Println([][]int{{1, 1, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 0, 0, 1, 1, 0, 1}})
	pp.Println("=========================================")
	fmt.Println(reconstructMatrix(9, 2, []int{0, 1, 2, 0, 0, 0, 0, 0, 2, 1, 2, 1, 2}))
	fmt.Println([][]int{})
	pp.Println("=========================================")
}

// 1 <= colsum.length <= 10^5
// 0 <= upper, lower <= colsum.length
// 0 <= colsum[i] <= 2
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var sum, count int
	for _, c := range colsum {
		sum += c
		if c == 2 {
			count++
		}
	}

	// 合計が合わなかったら[]
	if sum != upper+lower {
		return [][]int{}
	}

	// 2の数を数えて
	// 上か下にはいる数を数える
	upper -= count
	lower -= count

	if upper < 0 || lower < 0 {
		return [][]int{}
	}

	colLen := len(colsum)
	res := make([][]int, 2)
	res[0] = make([]int, colLen)
	res[1] = make([]int, colLen)

	for i := 0; i < colLen; i++ {
		if colsum[i] == 0 {
			continue
		}

		if colsum[i] == 2 {
			res[0][i] = 1
			res[1][i] = 1
			continue
		}

		if upper > 0 {
			res[0][i] = 1
			upper--
			continue
		}

		if lower > 0 {
			res[1][i] = 1
			lower--
			continue
		}
	}

	// 2のところは両方
	// 1のとき
	// uppperがなくなるまで上に入れる
	// なくなったらlowerがなくなるまでいれる

	// column から 1 と 2 の数を数えて

	return res
}
