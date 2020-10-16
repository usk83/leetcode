// 5413. Rearrange Words in a Sentence
// https://leetcode.com/problems/xxx/
package main

import (
	"sort"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

func arrangeWords(text string) string {
	words := strings.Split(text, " ")
	words[0] = strings.ToLower(words[0])
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	first := []byte(words[0])
	first[0] -= 'a' - 'A'
	words[0] = string(first)
	var ret string
	for i, word := range words {
		if i != 0 {
			ret += " "
		}
		ret += word
	}
	return ret
}
