// 72. Edit Distance
// https://leetcode.com/problems/edit-distance/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minDistance("horse", "ros"))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(minDistance("intention", "execution"))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(minDistance("abc", "abc"))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(minDistance("plasma", "altruism"))
	pp.Println(6)
	pp.Println("=========================================")
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	// 消して足すことはない

	// 多い文字数は消す

	// replaceの数を減らしたい

	// 最長で何文字一致させられるか

	// ただし、一致していない残りの文字数が残り文字数を切った場合は無効

	// 最小で len(word2) - len(word1)
	// 最大で len(word1)
	// 一致する文字があった分だけ減らせる

	// 先頭から差のindexまでで一致する文字があるか

	// 対象文字列の今のindex番目と一致文字数を確保する
	// 一致文字数が同じ場合は今のindexが小さい方が優先

	// あと何文字か？とコストを残す

	// 横軸 word1のlength
	// 縦軸 word2のlength

	// 検索を早くできるようにする
	m := map[byte][]int{}
	for i := range word1 {
		c := word1[i]
		m[c] = append(m[c], i)
	}

	// dp := [len(word1)][len(word2)]int{}
	// dp := [][]int{}

	dp := make([][]int, len(word1)+1+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		for j := range dp[i] {
			dp[i][j] = len(word1) + 1
		}
	}

	// 取れる選択肢
	// 任意の場所で任意の数add
	// remove
	// replace
	// 最後足りない場合はadd

	for index2 := range dp[0] {
		dp[0][index2] = index2
	}

	for index1 := range word1 {
		for index2 := range dp[index1] {
			/* たぶん不要
			 * // ソース文字が尽きたときは残りをaddしていく
			 * if index1 == len(word1)+1 {
			 * 	dp[index1+1][len(word2)+1] = min(dp[index1+1][len(word2)+1], dp[index1][index2]+len(word2)-index2)
			 * 	continue
			 * }
			 */

			// ※ 最後までいってたらremoveのみ
			if index2 == len(word2) {
				dp[index1+1][index2] = min(dp[index1+1][index2], dp[index1][index2]+1)
				continue
			}

			// 次の文字にマッチしている場合はコストそのままで次の文字へ
			if word1[index1] == word2[index2] {
				dp[index1+1][index2+1] = min(dp[index1+1][index2+1], dp[index1][index2])
			} else {
				if index2 < len(word2)-1 && word1[index1] == word2[index2+1] {
					// if word1[index1] == word2[index2+1] {
					dp[index1+1][index2+2] = min(dp[index1+1][index2+2], dp[index1][index2]+1)
				}

				// replace
				dp[index1+1][index2] = min(dp[index1+1][index2], dp[index1][index2]+1)

				// マッチしていない場合はマッチするまで任意の数removeする
				// 今のところを更新できるようなら更新する

				if matches, found := m[word1[index1]]; !found {
					dp[index1+1][len(word2)] = min(dp[index1+1][len(word2)], len(word2)-index2)
				} else {
					nextIndex := -1
					for _, match := range matches {
						if match < index2 {
							continue
						}
						nextIndex = match
						break
					}

					if nextIndex < 0 {
						dp[index1+1][len(word2)] = min(dp[index1+1][len(word2)], dp[index1][index2]+len(word2)-index2)
					} else {
						// dp[index1+1][nextIndex+1] = min(dp[index1+1][nextIndex+1], dp[index1][index2]+nextIndex-index2)
						dp[index1+1][index2+1] = min(dp[index1+1][index2+1], dp[index1][index2]+nextIndex-index2+1)
					}
				}
			}
		}
	}

	// 一番最後はまだマッチしていない文字をaddする
	for index := range dp[len(word1)+1][:len(word2)] {
		dp[len(word1)+1][index] = min(dp[len(word1)+1][index], dp[len(word1)][index]+len(word2)-index)
	}
	dp[len(word1)+1][len(word2)] = min(dp[len(word1)+1][len(word2)], dp[len(word1)][len(word2)])

	pp.Println("=== DEBUG START ======================================")
	fmt.Println(dp)
	pp.Println("=== DEBUG END ========================================")

	result := dp[len(word1)+1][0]
	for _, value := range dp[len(word1)+1][1:] {
		result = min(result, value)
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
