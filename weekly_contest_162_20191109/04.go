package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxScoreWords(
		[]string{"dog", "cat", "dad", "good"},
		[]byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
		[]int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	))
	pp.Println(23)
	pp.Println("=========================================")
	pp.Println(maxScoreWords(
		[]string{"xxxz", "ax", "bx", "cx"},
		[]byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'},
		[]int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10},
	))
	pp.Println(27)
	pp.Println("=========================================")
	pp.Println(maxScoreWords(
		[]string{"leetcode"},
		[]byte{'l', 'e', 't', 'c', 'o', 'd'},
		[]int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(maxScoreWords(
		[]string{"da", "ac", "aba", "abcc", "cadc"},
		[]byte{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'c', 'd', 'd', 'd'},
		[]int{3, 7, 9, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	))
	pp.Println(49)
	pp.Println("=========================================")
	pp.Println(maxScoreWords(
		[]string{"ac", "abd", "db", "ba", "dddd", "bca"},
		[]byte{'a', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'd', 'd', 'd', 'd'},
		[]int{6, 4, 4, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	))
	pp.Println(62)
	pp.Println("=========================================")
}

// 1 <= words.length <= 14
// 1 <= words[i].length <= 15
// 1 <= letters.length <= 100
// letters[i].length == 1
// score.length == 26
// 0 <= score[i] <= 10
// words[i], letters[i] contains only lower case English letters.
func maxScoreWords(words []string, letters []byte, score []int) int {
	// letters を 変換
	// [26]int にする
	lettersArray := [26]int{}
	for _, c := range letters {
		lettersArray[c-97]++
	}

	cloneLettersArray := func() [26]int {
		clone := [26]int{}
		for i, v := range lettersArray {
			clone[i] = v
		}
		return clone
	}

	// wordsの消費文字とスコアの組み合わせを作成
	// [26]int と スコア
	// map[string(word)]struct{
	// 	letters [26]int
	// 	score int
	// }
	// set := map[string]struct {
	// 	letters [26]int
	// 	score   int
	// }{}

	set := make([]struct {
		letters [26]int
		score   int
	}, 0, len(words))

WORDS_LOOP:
	for _, word := range words {
		// _, found := set[word]
		// 同じwordが複数回ある可能性を検討する

		rest := cloneLettersArray()
		need := [26]int{}
		var s int
		for _, c := range word {
			index := c - 97
			if rest[index] == 0 {
				continue WORDS_LOOP
			}
			need[index]++
			rest[index]--
			s += score[index]
		}
		set = append(set, struct {
			letters [26]int
			score   int
		}{need, s})
	}

	// 総当たり
	// 最大を返却
	var calc func(rest [26]int, total int, index int) int
	calc = func(rest [26]int, total int, index int) int {
		if index >= len(set) {
			return total
		}

		restClone := [26]int{}
		for i, v := range rest {
			restClone[i] = v
		}

		for i, cost := range set[index].letters {
			rest[i] -= cost
			if rest[i] < 0 {
				return calc(restClone, total, index+1)
			}
		}

		a := calc(restClone, total, index+1)
		b := calc(rest, total+set[index].score, index+1)
		if a > b {
			return a
		}
		return b
	}

	var max int
	for i := 0; i < len(set); i++ {
		ret := calc(cloneLettersArray(), 0, i)
		if ret > max {
			max = ret
		}
	}

	return max
}
