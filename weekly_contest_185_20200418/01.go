// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
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

func reformat(s string) string {
	digits := []byte{}
	letters := []byte{}
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			digits = append(digits, s[i])
		} else {
			letters = append(letters, s[i])
		}
	}

	if diff := len(digits) - len(letters); diff < -1 || 1 < diff {
		return ""
	}

	fs := []byte{}
	if len(digits) > len(letters) {
		for len(digits) != 0 {
			fs = append(fs, digits[0])
			digits = digits[1:]
			if len(letters) != 0 {
				fs = append(fs, letters[0])
				letters = letters[1:]
			}
		}
	} else {
		for len(letters) != 0 {
			fs = append(fs, letters[0])
			letters = letters[1:]
			if len(digits) != 0 {
				fs = append(fs, digits[0])
				digits = digits[1:]
			}
		}
	}
	return string(fs)
}
