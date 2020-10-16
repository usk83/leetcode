// 5278. Palindrome Partitioning III
// https://leetcode.com/contest/weekly-contest-165/problems/palindrome-partitioning-iii/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(palindromePartition("abc", 2))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(palindromePartition("aabbc", 3))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(palindromePartition("leetcode", 8))
	pp.Println(0)
	pp.Println("=========================================")
}

// 1 <= k <= s.length <= 100.
// s only contains lowercase English letters.
func palindromePartition(s string, k int) int {
	最長箇所をマージ

	// if len(s) < k {
	// 	return 0
	// }

	chars := make([]byte, 0, len(s))
	merged := false
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			chars = append(chars, '?')
			i += 2
		} else {
			chars = append(chars, s[i])
		}
	}

	pp.Println(string(chars))

	return 0

	// 3文字の繰り返しをマージ

	// すでに回文になっている数は？

	// // すでに回分になっているところをマージする

	// // 同じ文字を圧縮する
	// chars := make([]byte, 0, len(s))
	// prev := byte(0)
	// for i := 0; i < len(s); i++ {
	// 	if prev != s[i] {
	// 		chars = append(chars, s[i])
	// 	}
	// 	prev = s[i]
	// }
	// str := string(chars)

	// more := len(str) - k

	// if more <= 0 {
	// 	return 0
	// }

	// pp.Println(more)
	// count := 0

	// 並べられる可能性がある
	// 1文字変えると2つ減らせる場合もある
	// 1文字変えると1減る

	// for i := 0; i < len(s); i++ {
	// }

	// return more
}
