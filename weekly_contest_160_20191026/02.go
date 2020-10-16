package main

import (
	"fmt"
	"math"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(circularPermutation(2, 3))
	// fmt.Println([]int{3, 2, 0, 1})
	// pp.Println("=========================================")
	// fmt.Println(circularPermutation(3, 2))
	// fmt.Println([]int{2, 6, 7, 5, 4, 0, 1, 3})
	pp.Println("=========================================")
}

// 1 <= n <= 16
// 0 <= start < 2 ^ n
func circularPermutation(n int, start int) []int {
	// 16ビットまで
	// startは範囲内

	// startのbinary表現は？
	rep16 := fmt.Sprintf("%016b", start)
	repBase := rep16[16-n:]

	longest := []string{}

	// 横レングス 2 ^ n - 1
	// 値配列
	// 縦ビット
	max := int(math.Pow(2, float64(n)) - 1)
	dp := make([][][]string, max)
	for i := 0; i < max; i++ {
		dp[i] = make([][]string, n)
		for j := 0; j < n; j++ {
			dp[i][j] = []string{repBase}
		}
	}

	for i := 1; i < max; i++ {
		for j := 0; j < n; j++ {
			lastIndex := len(dp[i-1][j]) - 1
			if len(dp[i-1][j][lastIndex]) == 0 {
				break
			}
			flopped := flop(dp[i-1][j][lastIndex], j)
			pp.Println(flopped)
			if contains(dp[i-1][j], flopped) {
				if len(longest) < len(dp[i-1][j]) {
					longest = dp[i-1][j]
				}
				dp[i][j] = []string{}
			} else {
				dp[i][j] = append(dp[i-1][j], flopped)
			}
		}
	}

	for i := 0; i < n; i++ {
		if len(longest) < len(dp[max-1][i]) {
			longest = dp[max-1][i]
		}
	}

	// Nビット
	// 長さNの文字列

	// 数字に直すのは最後

	// // nビットのリストを作成する

	// 一旦Nビットで表現できるすべてを返す

	return []int{}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func flop(s string, i int) string {
	tmp := []rune(s)
	if tmp[i] == '0' {
		tmp[i] = '1'
	} else {
		tmp[i] = '0'
	}
	return string(tmp)
}
