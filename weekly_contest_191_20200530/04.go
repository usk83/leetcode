// 5427. Probability of a Two Boxes Having The Same Number of Distinct Balls
// https://leetcode.com/contest/weekly-contest-191/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(getProbability([]int{1, 1}))
	pp.Println(1.00000)
	pp.Println("=========================================")
	pp.Println(getProbability([]int{2, 1, 1}))
	pp.Println(0.66667)
	pp.Println("=========================================")
	pp.Println(getProbability([]int{1, 2, 1, 2}))
	pp.Println(0.60000)
	pp.Println("=========================================")
	pp.Println(getProbability([]int{3, 2, 1}))
	pp.Println(0.30000)
	pp.Println("=========================================")
	pp.Println(getProbability([]int{6, 6, 6, 6, 6, 6}))
	pp.Println(0.90327)
	pp.Println("=========================================")
}

// 1 <= balls.length <= 8
// 1 <= balls[i] <= 6
// sum(balls) is even.
// Answers within 10^-5 of the actual value will be accepted as correct.

/*
 * 小数点以下の計算の精度が足りない
 *
 * 組み合わせををいくつ作れるか
 * それぞれ何パターンになるか
 *
 * 8桁DPできそう
 *
 * BFS
 * 同じになったら合流させる
 *
 * 3ビット*8色 = 24bits あれば足りる
 * 自分の方に何色あるかだけ確認
 *
 * キュー: map[現在の状態: int]struct{残り: int, 確率: float64}
 */

type data struct {
	remain      int
	probability float64
}

func getProbability(balls []int) float64 {
	var wholeCount, remain int
	for i, count := range balls {
		remain |= count << (3 * uint(i))
		wholeCount += count
	}

	n := wholeCount >> 1
	// k := len(balls)

	queue := map[int]data{
		0: {remain, 1.0},
	}

	for loop := 0; loop < n; loop++ {
		nextQueue := map[int]data{}
		for current, datum := range queue {
			for i := range balls {
				colorCount := (datum.remain >> (3 * uint(i))) & 7
				if colorCount == 0 {
					continue
				}

				nextState := current + 1<<(3*uint(i))
				nextPro := datum.probability * float64(colorCount) / float64(wholeCount-loop)

				if _, ok := nextQueue[nextState]; ok {
					nextQueue[nextState] = data{
						remain:      nextQueue[nextState].remain,
						probability: nextQueue[nextState].probability + nextPro,
					}
				} else {
					nextQueue[nextState] = data{
						remain:      datum.remain - 1<<(3*uint(i)),
						probability: nextPro,
					}
				}
			}
		}
		queue = nextQueue
	}

	var probability float64
	for current, datum := range queue {
		var cc int
		for i := range balls {
			if (current>>(3*uint(i)))&7 != 0 {
				cc++
			}
		}

		var cc2 int
		for i := range balls {
			if (datum.remain>>(3*uint(i)))&7 != 0 {
				cc2++
			}
		}

		if cc == cc2 {
			probability += datum.probability
		}
	}
	return probability
}

/*
 * 1. brute force (Time Limit Exceeded)
 */
// func getProbability(balls []int) float64 {
// 	var wholeCount int
// 	for _, count := range balls {
// 		wholeCount += count
// 	}
//
// 	box := make([]int, 8)
// 	var probability float64
// 	var bk func(remainCount int, pro float64)
// 	bk = func(remainCount int, pro float64) {
// 		var colorsA, colorsB int
// 		for _, count := range box {
// 			if count != 0 {
// 				colorsA++
// 			}
// 		}
// 		for _, count := range balls {
// 			if count != 0 {
// 				colorsB++
// 			}
// 		}
// 		if colorsA > colorsB {
// 			return
// 		}
//
// 		if remainCount == wholeCount>>1 {
// 			if colorsA == colorsB {
// 				probability += pro
// 			}
// 			return
// 		}
//
// 		remainCountf := float64(remainCount)
// 		for i, count := range balls {
// 			if count == 0 {
// 				continue
// 			}
// 			balls[i]--
// 			box[i]++
// 			bk(remainCount-1, pro*float64(count)/float64(remainCountf))
// 			balls[i]++
// 			box[i]--
// 		}
// 	}
// 	bk(wholeCount, 1)
//
// 	return probability
// }
