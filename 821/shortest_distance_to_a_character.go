// 821. Shortest Distance to a Character
// https://leetcode.com/problems/shortest-distance-to-a-character/
package main

import (
	"fmt"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(shortestToChar("loveleetcode", 'e'))
	fmt.Println([]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0})
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// S string length is in [1, 10000].
// C is a single character, and guaranteed to be in string S.
// All letters in S and C are lowercase.
func shortestToChar(S string, C byte) []int {
	result := make([]int, len(S))

	index := strings.IndexByte(S, C)
	for i, j := 0, index; j >= 0; i, j = i+1, j-1 {
		result[i] = j
	}

	// pp.Println("first", index)

	for nextIndex := strings.IndexByte(S[index+1:], C); nextIndex != -1; nextIndex = strings.IndexByte(S[index+1:], C) {
		index++
		nextIndex += index

		// pp.Println("next", nextIndex)

		middle := (index + nextIndex) >> 1
		mod := (index + nextIndex) & 1

		// pp.Println("middle", middle)

		var i int
		for i = 1; index < middle; i, index = i+1, index+1 {
			result[index] = i
		}
		i -= 1 - mod
		// pp.Println("index", index)
		for ; index < nextIndex; i, index = i-1, index+1 {
			result[index] = i
		}

		index = nextIndex
	}

	for i, j := index+1, 1; i < len(S); i, j = i+1, j+1 {
		result[i] = j
	}

	return result
}
