// 174. Dungeon Game
// https://leetcode.com/problems/dungeon-game/
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 21
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/541/week-3-june-15th-june-21st/3367/
func main() {
	pp.Println("=========================================")
	pp.Println(calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}))
	pp.Println(7)
	pp.Println("=========================================")
	pp.Println(calculateMinimumHP([][]int{
		{-5},
	}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(calculateMinimumHP([][]int{
		{100},
	}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(calculateMinimumHP([][]int{
		{0, 0},
	}))
	pp.Println(1)
	pp.Println("=========================================")
}

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

func calculateMinimumHP(dungeon [][]int) int {
	/*
		BFS
			キュー
			現時点での最大の最低体力をMAXで初期化する
		以下の情報を保持する
			現在x座標
			現在y座標
			現在の体力
			最低地点の体力
		処理するときに
			最大の最低体力より最低地点が悪かったら終了
			同じ地点には
	*/

	// 2次元DP
	// 上か左から来る
	// 現時点での体力が一番多いか最下点が高いもののみ保持

	current := [][][2]int{
		{
			{dungeon[0][0], dungeon[0][0]},
		},
	}
	yLast, xLast := len(dungeon)-1, len(dungeon[0])-1
	for i := 1; i <= yLast+xLast; i++ {
		var length int
		switch {
		case i <= min(yLast, xLast):
			length = i + 1
		case i <= max(yLast, xLast):
			length = min(yLast, xLast)
		default:
			length = yLast + xLast - i + 1
		}

		next := make([][][2]int, length)

		yStart, xStart := min(i, yLast), max(i-yLast, 0)
		// pp.Println("=== DEBUG START ======================================")
		// pp.Println("length:", length)
		for i := 0; i < length; i++ {
			yIndex, xIndex := yStart-i, xStart+i
			// fmt.Printf("dungeon[%d][%d] = %d\n", yIndex, xIndex, dungeon[yIndex][xIndex])
			if yIndex == 0 || xIndex == 0 {
				var currentIndex int
				if yIndex == 0 {
					currentIndex = i - 1
				} else {
					currentIndex = i
				}

				// pp.Println("=== DEBUG START ======================================")
				// fmt.Println(next)
				// pp.Println(i)
				// fmt.Println(current)
				// pp.Println(currentIndex)
				// pp.Println("=== DEBUG END ========================================")

				next[i] = make([][2]int, len(current[currentIndex]))
				for j := range current[currentIndex] {
					next[i][j] = [2]int{
						current[currentIndex][j][0] + dungeon[yIndex][xIndex],
						min(current[currentIndex][j][1], current[currentIndex][j][0]+dungeon[yIndex][xIndex]),
					}
				}
			} else {
				maxHealth, minHealth := [2]int{MinInt, MinInt}, [2]int{MinInt, MinInt}

				var targets [][][2]int
				if len(next) > len(current) {
					targets = [][][2]int{current[i-1], current[i]}
				} else {
					targets = [][][2]int{current[i], current[i+1]}
				}

				for _, cur := range targets {
					for j := range cur {
						if cur[j][0] > maxHealth[0] || (cur[j][0] == maxHealth[0] && cur[j][1] > maxHealth[1]) {
							maxHealth = cur[j]
						}
						if cur[j][1] > minHealth[1] || (cur[j][1] == minHealth[1] && cur[j][0] > minHealth[0]) {
							minHealth = cur[j]
						}
					}
				}
				if maxHealth == minHealth {
					next[i] = [][2]int{
						{
							maxHealth[0] + dungeon[yIndex][xIndex],
							min(maxHealth[1], maxHealth[0]+dungeon[yIndex][xIndex]),
						},
					}
				} else {
					next[i] = [][2]int{
						{
							maxHealth[0] + dungeon[yIndex][xIndex],
							min(maxHealth[1], maxHealth[0]+dungeon[yIndex][xIndex]),
						},
						{
							minHealth[0] + dungeon[yIndex][xIndex],
							min(minHealth[1], minHealth[0]+dungeon[yIndex][xIndex]),
						},
					}
				}
			}
			// fmt.Printf("dungeon[%d][%d] = %d\n", yIndex, xIndex, dungeon[yIndex][xIndex])
			// fmt.Println(next[i])
		}
		// pp.Println("=== DEBUG END ========================================")

		current = next
	}

	minHealth := MinInt
	for _, stat := range current[0] {
		minHealth = max(minHealth, stat[1])
	}

	return max(-minHealth+1, 1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
