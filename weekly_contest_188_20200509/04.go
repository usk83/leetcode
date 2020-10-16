// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(ways([]string{"A..", "AAA", "..."}, 3))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(ways([]string{"A..", "AA.", "..."}, 3))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(ways([]string{"A..", "A..", "..."}, 1))
	pp.Println(1)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// 1 <= rows, cols <= 50
// rows == pizza.length
// cols == pizza[i].length
// 1 <= k <= 10
// pizza consists of characters 'A' and '.' only.
const modulo = 1000000007

func ways(pizzaStringArray []string, k int) int {
	pizza := make([][]int, len(pizzaStringArray))
	var apples int
	for i := range pizza {
		pizza[i] = make([]int, len(pizzaStringArray[i]))
		for j, r := range pizzaStringArray[i] {
			if r == 'A' {
				pizza[i][j] = 1
				apples++
			}
		}
	}

	// 各行ごとにいくつのappleがあるかの累積和
	// 各列ごとにいくつのappleがあるかの累積和
	// 以下の逆を求める
	sumAppleRow, sumAppleCol := make([][]int, len(pizza)), make([][]int, len(pizza[0]))
	for i, row := range pizza {
		sumAppleRow[i] = make([]int, len(row))
		sumAppleRow[i][len(row)-1] = row[len(row)-1]
		for j := len(row) - 1 - 1; j >= 0; j-- {
			sumAppleRow[i][j] = sumAppleRow[i][j+1] + row[j]
		}
	}

	for i := range pizza[0] {
		sumAppleCol[i] = make([]int, len(pizza))
		sumAppleCol[i][len(pizza)-1] = pizza[len(pizza)-1][i]
		for j := len(pizza) - 1 - 1; j >= 0; j-- {
			sumAppleCol[i][j] = sumAppleCol[i][j+1] + pizza[j][i]
		}
	}

	// pp.Println("=== DEBUG START ======================================")
	// fmt.Println(pizza)
	// fmt.Println(apples)
	// fmt.Println(sumAppleRow)
	// fmt.Println(sumAppleCol)
	// pp.Println("=== DEBUG END ========================================")

	var count int
	memo := map[[3]int]int{}
	var dfs func(int, int, [][]int, [][]int, []string) int
	dfs = func(apples, k int, sumAppleRow, sumAppleCol [][]int, ops []string) int {

		// バックトラッキング

		if cache, ok := memo[[2]int{len(sumAppleRow), len(sumAppleCol), k}]; ok {
			count = (count + cache) % modulo
			return
		}

		var successCount int

		// pp.Println("=== DEBUG START ======================================")
		// pp.Println("current:", ops)
		// fmt.Println(sumAppleRow)
		// fmt.Println(sumAppleCol)
		// pp.Println("=== DEBUG END ========================================")

		if k == 1 {
			if apples != 0 {
				count = (count + 1) % modulo
				// pp.Println("=== DEBUG START ======================================")
				// pp.Println("count:", count)
				// pp.Println(ops)
				// pp.Println("=== DEBUG END ========================================")
			}
			return
		}

		// 水平方向に好きなだけ切る
		for i := 1; i < len(sumAppleRow); i++ {
			// pp.Printf("try h %s\n", i)
			var consume int
			for index := 0; index < i; index++ {
				consume += sumAppleRow[index][0]
			}
			// pp.Println("consume:", consume, "k:", k, "apples-consume:", apples-consume, "k-1:", k-1)
			if consume == 0 {
				continue
			}
			if apples-consume < k-1 {
				continue
			}
			newSumAppleCol := make([][]int, len(sumAppleCol))
			for j := range sumAppleCol {
				newSumAppleCol[j] = sumAppleCol[j][i:]
			}
			// pp.Println("next")
			successCount += dfs(apples-consume, k-1, sumAppleRow[i:], newSumAppleCol, append(ops, fmt.Sprintf("h%d", i)))
			// pp.Println("return current:", ops)
		}

		// 垂直方向に好きなだけ切る
		for i := 1; i < len(sumAppleCol); i++ {
			// pp.Printf("try v %s\n", i)
			var consume int
			for index := 0; index < i; index++ {
				consume += sumAppleCol[index][0]
			}
			// pp.Println("consume:", consume, "k:", k, "apples-consume:", apples-consume, "k-1:", k-1)
			if consume == 0 {
				continue
			}
			if apples-consume < k-1 {
				continue
			}
			newSumAppleRow := make([][]int, len(sumAppleRow))
			for j := range sumAppleRow {
				newSumAppleRow[j] = sumAppleRow[j][i:]
			}
			// pp.Println("next")
			successCount += dfs(apples-consume, k-1, sumAppleCol[i:], newSumAppleRow, append(ops, fmt.Sprintf("v%d", i)))
			// pp.Println("return current:", ops)
		}

		memo[[2]int{len(sumAppleRow), len(sumAppleCol), k}] = successCount

		// 切ったときに
		// 	Appleが入らないときは無効
		// 	入るとき
		// 		kを1減らす
		// 		applesを使っただけ減らす
		// 		行と列を更新して次に渡す

	}
	dfs(apples, k, sumAppleRow, sumAppleCol, []string{})
	return count
}
