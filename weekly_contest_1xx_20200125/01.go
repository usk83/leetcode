// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(removePalindromeSub("ababa"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(removePalindromeSub("abb"))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(removePalindromeSub("baabb"))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(removePalindromeSub(""))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(removePalindromeSub("bbbaaabaababbbbbbaababbaababbbbabaaaabaababbbbaababababababaabaaaabbbaabbbbbbbbaaababaaaabbbababbbbb"))
	pp.Println()
	pp.Println("=========================================")
	pp.Println(removePalindromeSub("abbaaaab"))
	pp.Println(2)
	pp.Println("=========================================")

}

// a,bそれぞれそれぞれ偶数個だったら少なくとも1回で消せる
// 奇数個だったときそれぞれについては2回かかる
// 最悪で4回

func removePalindromeSub(s string) int {
	var count int
	for s != "" {
		count++
		next := ""
		for head, tail := 0, len(s)-1; head <= tail; {
			if s[head] == s[tail] {
				head++
				tail--
				continue
			}
			next = string(s[tail]) + next
			tail--
		}
		s = next
	}

	if count > 1 {
		return 2
	}

	return count
}
