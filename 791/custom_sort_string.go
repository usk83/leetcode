// 791. Custom Sort String
// https://leetcode.com/problems/custom-sort-string/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println("cbad")

	fmt.Println(customSortString("cbafg", "abcd"))
}

// S has length at most 26, and no character is repeated in S.
// T has length at most 200.
// S and T consist of lowercase letters only.
func customSortString(S string, T string) string {
	charIndexMap := make(map[byte]int, 26)
	orderedCharAndCount := make([]struct {
		char  byte
		count uint8
	}, len(S))
	for i := len(S) - 1; i >= 0; i-- {
		charIndexMap[S[i]] = i
		orderedCharAndCount[i].char = S[i]
	}

	resultChars := make([]byte, 0, len(T))
	for i := len(T) - 1; i >= 0; i-- {
		if index, found := charIndexMap[T[i]]; found {
			orderedCharAndCount[index].count++
			continue
		}
		resultChars = append(resultChars, T[i])
	}

	for _, charAndCount := range orderedCharAndCount {
		for i := charAndCount.count; i > 0; i-- {
			resultChars = append(resultChars, charAndCount.char)
		}
	}

	return string(resultChars)
}

func customSortStringV1(S string, T string) string {
	orders := [26]*int{}
	for i := len(S) - 1; i >= 0; i-- {
		index := i
		orders[S[i]-'a'] = &index
	}

	resultChars := make([]byte, 0, len(T))
	ordered := make([]struct {
		char  byte
		count uint8
	}, len(S))
	for i := len(T) - 1; i >= 0; i-- {
		if index := orders[T[i]-'a']; index != nil {
			ordered[*index].char = T[i]
			ordered[*index].count++
			continue
		}
		resultChars = append(resultChars, T[i])
	}

	for _, obj := range ordered {
		for i := obj.count; i > 0; i-- {
			resultChars = append(resultChars, obj.char)
		}
	}

	return string(resultChars)
}
