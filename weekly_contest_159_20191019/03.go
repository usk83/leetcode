package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("====================")
	// pp.Println(balancedString("QWER"))
	// pp.Println(0)
	// pp.Println("====================")
	// pp.Println(balancedString("QQWE"))
	// pp.Println(1)
	// pp.Println("====================")
	// pp.Println(balancedString("QQQW"))
	// pp.Println(2)
	// pp.Println("====================")
	// pp.Println(balancedString("QQQQ"))
	// pp.Println(3)
	pp.Println("====================")
	pp.Println(balancedString("WWQQRRRRQRQQ"))
	pp.Println(4)
	pp.Println("====================")
}

func balancedString(s string) int {
	type combi struct {
		Line    string
		Barance [4]int
	}

	full := len(s)
	limit := full / 4

	dp := make([][]combi, full)
	for i := 0; i < full; i++ {
		dp[i] = make([]combi, 0, 8)
	}

	count := func(c byte, b *[4]int) bool {
		switch c {
		case 'Q':
			b[0]++
			if b[0] > limit {
				return false
			}
		case 'W':
			b[1]++
			if b[1] > limit {
				return false
			}
		case 'E':
			b[2]++
			if b[2] > limit {
				return false
			}
		case 'R':
			b[3]++
			if b[3] > limit {
				return false
			}
		}
		return true
	}

	combi1 := combi{
		s[1:],
		[4]int{0, 0, 0, 0},
	}
	count(s[0], &combi1.Barance)

	combi2 := combi{
		s[:full-1],
		[4]int{0, 0, 0, 0},
	}
	count(s[full-1], &combi2.Barance)

	dp[0] = append(dp[0], combi1, combi2)

	i := 1
	flag := false
	for i = 1; i < full-1; i++ {
		flag = false
		for _, d := range dp[i-1] {
			combi1 := combi{
				d.Line[1:],
				d.Barance,
			}
			if count(d.Line[1], &d.Barance) {
				dp[i] = append(dp[i], combi1)
				flag = true
			}

			l := len(d.Line)
			combi2 := combi{
				d.Line[:l-1],
				d.Barance,
			}
			if count(d.Line[l-1], &d.Barance) {
				dp[i] = append(dp[i], combi2)
				flag = true
			}
		}
		if !flag {
			pp.Println("debugging...")
			break
		}
	}

	// →全長
	// ↓残り文字列とバランス
	pp.Println(i)

	return 0

}

// 1 <= s.length <= 10^5
// s.length is a multiple of 4
// s contains only 'Q', 'W', 'E' and 'R'.

// func balancedString(s string) int {
// 	// 頭からと後ろから
// 	// もっともバランスドに近い状態を選択する

// 	// 全長から期待される出現数を求める
// 	// 超えられると取り返しがつかない
// 	// 超えられるまでは選択したほうがいい

// 	// 前からと後ろから

// 	full := len(s)
// 	limit := full / 4

// 	max := 0

// 	flag := false
// 	hits := []int{0, 0, 0, 0}
// 	for max = 0; max < full; max++ {
// 		switch s[max] {
// 		case 'Q':
// 			hits[0]++
// 			if hits[0] > limit {
// 				flag = true
// 			}
// 		case 'W':
// 			hits[1]++
// 			if hits[1] > limit {
// 				flag = true
// 			}
// 		case 'E':
// 			hits[2]++
// 			if hits[2] > limit {
// 				flag = true
// 			}
// 		case 'R':
// 			hits[3]++
// 			if hits[3] > limit {
// 				flag = true
// 			}
// 		}
// 		if flag {
// 			break
// 		}
// 	}

// 	rev := 0

// 	flag = false
// 	hits = []int{0, 0, 0, 0}
// 	for rev = full - 1; rev >= 0; rev-- {
// 		switch s[rev] {
// 		case 'Q':
// 			hits[0]++
// 			if hits[0] > limit {
// 				flag = true
// 			}
// 		case 'W':
// 			hits[1]++
// 			if hits[1] > limit {
// 				flag = true
// 			}
// 		case 'E':
// 			hits[2]++
// 			if hits[2] > limit {
// 				flag = true
// 			}
// 		case 'R':
// 			hits[3]++
// 			if hits[3] > limit {
// 				flag = true
// 			}
// 		}
// 		if flag {
// 			rev++
// 			break
// 		}
// 	}

// 	if !flag {
// 		return full - max
// 	}

// 	if max > full-rev {
// 		return full - max
// 	}

// 	return rev
// }
