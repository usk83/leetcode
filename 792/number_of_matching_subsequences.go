// 792. Number of Matching Subsequences
// https://leetcode.com/problems/number-of-matching-subsequences/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	pp.Println(3)
	pp.Println("=========================================")
}

// All words in words and S will only consists of lowercase letters.
// The length of S will be in the range of [1, 50000].
// The length of words will be in the range of [1, 5000].
// The length of words[i] will be in the range of [1, 50].

// Trie

/*
 * 3. two pointers + caching solution
 * - Time Complexity: O(M(N+L))
 * - Space Complexity: O(M)
 *   where N: the length of s, M: the number of words, L: the length of each word
 *
 * - Runtime: 32 ms, faster than 100.00%
 * - Memory Usage: 6.5 MB, less than 96.30%
 */
func numMatchingSubseq(s string, words []string) int {
	var count int
	checked := map[string]bool{}
LOOP:
	for _, word := range words {
		if yes, ok := checked[word]; ok {
			if yes {
				count++
			}
			continue
		}
		var index int
		for _, r := range word {
			foundAt := strings.IndexRune(s[index:], r)
			if foundAt == -1 {
				checked[word] = false
				continue LOOP
			}
			index += 1 + foundAt
		}
		checked[word] = true
		count++
	}
	return count
}

/*
 * 2. two pointers solution
 * - Time Complexity: O(M(N+L))
 * - Space Complexity: O(1)
 *   where N: the length of s, M: the number of words, L: the length of each word
 *
 * - Runtime: 112 ms, faster than 66.67%
 * - Memory Usage: 6.5 MB, less than 92.59%
 *
 * - Runtime: 84 ms, faster than 72.73%
 * - Memory Usage: 6.6 MB, less than 70.37%
 */
// func numMatchingSubseq(s string, words []string) int {
// 	var count int
// LOOP:
// 	for _, word := range words {
// 		var index int
// 		for _, r := range word {
// 			foundAt := strings.IndexRune(s[index:], r)
// 			if foundAt == -1 {
// 				continue LOOP
// 			}
// 			index += 1 + foundAt
// 		}
// 		count++
// 	}
// 	return count
// }

/*
 * 1. Binary search solution
 * - Time Complexity: O(N+M*L*logN)
 * - Space Complexity: O(N)
 *   where N: the length of s, M: the number of words, L: the length of each word
 *
 * Runtime: 92 ms, faster than 69.70%
 * Memory Usage: 8.6 MB, less than 18.52%
 */
// func numMatchingSubseq(s string, words []string) int {
// 	var indices [26][]int
// 	for i, c := range s {
// 		indices[c-'a'] = append(indices[c-'a'], i)
// 	}
//
// 	var count int
// LOOP:
// 	for _, word := range words {
// 		pointer := -1
// 		for _, c := range word {
// 			index := sort.Search(len(indices[c-'a']), func(i int) bool {
// 				return indices[c-'a'][i] > pointer
// 			})
// 			if index == len(indices[c-'a']) {
// 				continue LOOP
// 			}
// 			pointer = indices[c-'a'][index]
// 		}
// 		count++
// 	}
// 	return count
// }
