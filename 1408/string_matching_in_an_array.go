// 1408. String Matching in an Array
// https://leetcode.com/problems/string-matching-in-an-array/
package main

import (
	"fmt"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))
	fmt.Println([]string{"as", "hero"})
	pp.Println("=========================================")
	fmt.Println(stringMatching([]string{"leetcode", "et", "code"}))
	fmt.Println([]string{"et", "code"})
	pp.Println("=========================================")
	fmt.Println(stringMatching([]string{"blue", "green", "bu"}))
	fmt.Println([]string{})
	pp.Println("=========================================")
}

/*
 * 1. The first solution (brute force)
 */
func stringMatching(words []string) []string {
	result := []string{}
	for _, word := range words {
		for _, comp := range words {
			if word != comp && strings.Contains(comp, word) {
				result = append(result, word)
				break
			}
		}
	}
	return result
}
