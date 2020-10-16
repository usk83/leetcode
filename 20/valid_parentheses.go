// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
package main

import "fmt"

func main() {
	fmt.Println("isValid(\"()\")\n->", isValid("()") == true)
	fmt.Println("isValid(\"()[]{}\")\n->", isValid("()[]{}") == true)
	fmt.Println("isValid(\"(]\")\n->", isValid("(]") == false)
	fmt.Println("isValid(\"([)]\")\n->", isValid("([)]") == false)
	fmt.Println("isValid(\"{[]}\")\n->", isValid("{[]}") == true)
	fmt.Println("isValid(\"{{{{{\")\n->", isValid("{{{{{") == false)
	fmt.Println("isValid(\")(][}{\")\n->", isValid(")(][}{") == false)
	fmt.Println("isValid(\"}}}}}\")\n->", isValid("}}}}}") == false)
}

func isValid(s string) bool {
	length := len(s)
	stack := make([]byte, length>>1)
	parentheses := map[byte]byte{
		')': 0,
		'}': 0,
		']': 0,
		'(': ')',
		'{': '}',
		'[': ']',
	}

	index := 0
	for i := length - 1; i >= 0; i-- {
		c := s[i]
		if p, found := parentheses[c]; found {
			if p == 0 {
				index++
				if index > i {
					return false
				}
				stack[index-1] = c
			} else {
				index--
				if index < 0 || stack[index] != p {
					return false
				}
			}
		}
	}

	return index == 0
}

func isValidV1(s string) bool {
	halfLength := len(s) >> 1
	stack := make([]rune, halfLength)
	parentheses := map[rune]rune{
		'(': 0,
		'{': 0,
		'[': 0,
		')': '(',
		'}': '{',
		']': '[',
	}

	index := 0
	for _, c := range s {
		parenthesis, found := parentheses[c]
		if !found {
			continue
		}
		if parenthesis == 0 {
			index++
			if index > halfLength {
				return false
			}
			stack[index-1] = c
		} else {
			index--
			if index < 0 || stack[index] != parenthesis {
				return false
			}
		}
	}

	return index == 0
}
