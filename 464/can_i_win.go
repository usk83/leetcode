// 464. Can I Win
// https://leetcode.com/problems/can-i-win/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(canIWin(10, 11))
	// pp.Println(canIWin(15, 11))
	pp.Println(canIWin(20, 11))
	// pp.Println(canIWin(5, 11))
	// pp.Println(canIWin(1, 1))
	// pp.Println(canIWin(2, 1))
	// pp.Println(canIWin(3, 1))
	// pp.Println(canIWin(4, 1))
	// pp.Println(canIWin(5, 1))

	// canIWin(10, 11)
	// canIWin(15, 11)
	// canIWin(20, 11)
	// canIWin(5, 11)
	// canIWin(1, 1)
	// canIWin(2, 1)
	// canIWin(3, 1)
	// canIWin(4, 1)
	// canIWin(5, 1)
}

// maxChoosableInteger: 20
// desiredTotal: 300
//   → 実質210では？
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	// 後手が勝てる確率の反対
	// 後手が勝てる確率を考える

	// 1..10
	// 11 後手
	// 12 先手
	// 13 先手
	// 14 先手
	// 15 先手
	// 16 先手
	// 17 先手
	// 18 先手
	// 19 先手
	// 20
	// 21 後手？
	// 22
	// 23
	// 24
	// 25
	// 26
	// 27
	// 28
	// 29

	// 超える数字の組み合わせを考える

	// 下からの組みあわせ数と上からの組み合わせ数

	勝ちに達するまでの数字列の長さを奇数に強制できれば勝てる
	勝ちに達するまでの数列を考える（毎回）
	毎回奇数回の組み合わせになる数字からどれかを取る

	相手の番
		最大値 < 残り <= 最小値+最大値

	相手の番
 		最大値 < 残り-(最小値+最大値) or 残り-(最小値+最小値) <= 最小値+最大値



	相手の番に残り数が
		最大値より大きく
		最小値+最大値かより小さくなれば勝てる
			この差は最小値
	相手の番に残り数が最小値+最大値より大きいとき
		一手で上記条件にされると負ける





















	if desiredTotal == 0 {
		return true
	}

	if maxChoosableInteger < desiredTotal {
		return true
	}

	sum := 0
	for i := 1; i <= maxChoosableInteger; i++ {
		sum += i
	}
	// pp.Println(sum)
	if sum < desiredTotal {
		return false
	}

	x := 0
	for sum := 0;; {
		sum += x
		if sum >= desiredTotal {
			break
		}
		x++
	}

	if x == maxChoosableInteger {
		return x % 2 == 1
	}

	nums := make([]int, maxChoosableInteger)
	for i := 0; i < maxChoosableInteger; i++ {
		nums[i] = i + 1
	}

	dp := make([][]int, x)
	for i := 0; i < x; i++ {
		dp[i] = make([]int, maxChoosableInteger)
	}

	fmt.Println(dp)

	// 1, 2, 3, 4, 5, 6, 7
	// 8
	// 9, 10, 11, 12, 13, 14, 15

	// →今の合計（N回目）
	// ↓残りの選択肢

	// Nの最大値は計算できる

	return false
}
