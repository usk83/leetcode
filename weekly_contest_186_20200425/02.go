// 5393. Maximum Points You Can Obtain from Cards
// https://leetcode.com/contest/weekly-contest-186/problems/maximum-points-you-can-obtain-from-cards/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	pp.Println(12)
	pp.Println("=========================================")
	pp.Println(maxScore([]int{2, 2, 2}, 2))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7))
	pp.Println(55)
	pp.Println("=========================================")
	pp.Println(maxScore([]int{1, 1000, 1}, 1))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3))
	pp.Println(202)
	pp.Println("=========================================")
}

func maxScore(cardPoints []int, k int) int {
	// 全体の合計から
	// 取らないカードの合計をひく
	// 取らないカードの合計が最小になる箇所を求める
	cache := make([]int, len(cardPoints))
	cache[0] = cardPoints[0]
	for i := range cardPoints[1:] {
		cache[i+1] = cache[i] + cardPoints[i+1]
	}
	notTaken := len(cardPoints) - k
	if notTaken == 0 {
		return cache[len(cache)-1]
	}
	minSub := cache[notTaken-1]
	start := notTaken
	for i := notTaken; i < len(cardPoints); i++ {
		minSub = min(minSub, cache[i]-cache[i-start])
	}
	return cache[len(cache)-1] - minSub
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func maxScore(cardPoints []int, k int) int {
// 	cache := make([]int, len(cardPoints))
// 	cache[0] = cardPoints[0]
// 	for i := range cardPoints[1:] {
// 		cache[i+1] = cache[i] + cardPoints[i+1]
// 	}

// 	checked := map[[2]int]bool{}
// 	var maxScore int
// 	var bf func(left, right, rem, score int)
// 	bf = func(left, right, rem, score int) {
// 		key := [2]int{left, right}
// 		if checked[key] {
// 			return
// 		}
// 		checked[key] = true
// 		if rem == 0 {
// 			maxScore = max(maxScore, score)
// 			return
// 		}

// 		if maxScore != 0 {
// 			allLeft := cache[left+rem]
// 			if left != 0 {
// 				allLeft -= allLeft
// 			}
// 			allRight := cache[right]
// 			if start := right - rem - 1; start >= 0 {
// 				allRight -= cache[start]
// 			}
// 			if score+allLeft+allRight <= maxScore {
// 				return
// 			}
// 		}
// 		bf(left+1, right, rem-1, score+cardPoints[left])
// 		bf(left, right-1, rem-1, score+cardPoints[right])
// 	}
// 	bf(0, len(cardPoints)-1, k, 0)
// 	return maxScore
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
