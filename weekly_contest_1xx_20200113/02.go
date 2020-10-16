// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(printVertically("HOW ARE YOU"))
	pp.Println([]string{"HAY", "ORO", "WEU"})
	pp.Println("=========================================")
	pp.Println(printVertically("TO BE OR NOT TO BE"))
	pp.Println([]string{"TBONTB", "OEROOE", "   T"})
	pp.Println("=========================================")
	pp.Println(printVertically("CONTEST IS COMING"))
	pp.Println([]string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"})
	pp.Println("=========================================")
}

// 1 <= s.length <= 200
// s contains only upper case English letters.
// It's guaranteed that there is only one space between 2 words.
func printVertically(s string) []string {
	words := strings.Split(s, " ")

	length := 0
	for _, word := range words {
		if len(word) > length {
			length = len(word)
		}
	}

	results := make([]string, length)
	for i := 0; i < length; i++ {
		for _, word := range words {
			if len(word) <= i {
				results[i] += " "
			} else {
				results[i] += string(word[i])
			}
		}
	}

	for i := range results {
		results[i] = strings.TrimRight(results[i], " ")
	}
	return results
}
