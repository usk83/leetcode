// 5416. Check If a Word Occurs As a Prefix of Any Word in a Sentence
// https://leetcode.com/contest/weekly-contest-190/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isPrefixOfWord("i love eating burger", "burg"))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(isPrefixOfWord("this problem is an easy problem", "pro"))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(isPrefixOfWord("i am tired", "you"))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(isPrefixOfWord("i use triple pillow", "pill"))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(isPrefixOfWord("hello from the other side", "they"))
	pp.Println(-1)
	pp.Println("=========================================")
}

// 1 <= sentence.length <= 100
// 1 <= searchWord.length <= 10
// sentence consists of lowercase English letters and spaces.
// searchWord consists of lowercase English letters.
func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	index := -1
	for i := range words {
		if strings.HasPrefix(words[i], searchWord) {
			index = i + 1
			break
		}
	}
	return index
}
