package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println(balancedStringSplit("RLRRLLRLRL"))
	pp.Println(4)
	pp.Println(balancedStringSplit("RLLLLRRRLR"))
	pp.Println(3)
	pp.Println(balancedStringSplit("LLLLRRRR"))
	pp.Println(1)
}

// 1 <= s.length <= 1000
// s[i] = 'L' or 'R'
func balancedStringSplit(s string) int {
	count := 0
	countR, countL := 0, 0
	for _, c := range s {
		switch c {
		case 'R':
			countR++
		case 'L':
			countL++
		}
		if countR == countL {
			count++
			countR, countL = 0, 0
		}
	}

	return count
}
