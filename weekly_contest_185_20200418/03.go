// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minNumberOfFrogs("croak"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minNumberOfFrogs("croakcroak"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minNumberOfFrogs("crcoakroak"))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minNumberOfFrogs("croakcrook"))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(minNumberOfFrogs("croakcroa"))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(minNumberOfFrogs("ccrrooaakk"))
	pp.Println(2)
	pp.Println("=========================================")
}

func minNumberOfFrogs(croakOfFrogs string) int {
	dp := map[[4]int]int{}
	dp[[4]int{0, 0, 0, 0}] = 0
	for i := range croakOfFrogs {
		// fmt.Println(dp)
		next := map[[4]int]int{}
		for status, count := range dp {
			switch croakOfFrogs[i] {
			case 'c':
				nextStatus := [4]int{status[0] + 1, status[1], status[2], status[3]}
				s := sum(nextStatus)
				if s > len(croakOfFrogs)-i-1 {
					continue
				}
				nextCount := max(s, count)
				if _, ok := next[nextStatus]; ok {
					next[nextStatus] = min(next[nextStatus], nextCount)
				} else {
					next[nextStatus] = nextCount
				}
			case 'r':
				currentSum := status[1] + 1 + status[2] + status[3]
				for i := status[0] - 1; i >= 0; i-- {
					nextStatus := [4]int{i, status[1] + 1, status[2], status[3]}
					s := currentSum + i
					if s > len(croakOfFrogs)-i-1 {
						continue
					}
					nextCount := max(s, count)
					if _, ok := next[nextStatus]; ok {
						next[nextStatus] = min(next[nextStatus], nextCount)
					} else {
						next[nextStatus] = nextCount
					}
				}
			case 'o':
				currentSum := status[0] + status[2] + 1 + status[3]
				for i := status[1] - 1; i >= 0; i-- {
					nextStatus := [4]int{status[0], i, status[2] + 1, status[3]}
					s := currentSum + i
					if s > len(croakOfFrogs)-i-1 {
						continue
					}
					nextCount := max(s, count)
					if _, ok := next[nextStatus]; ok {
						next[nextStatus] = min(next[nextStatus], nextCount)
					} else {
						next[nextStatus] = nextCount
					}
				}
			case 'a':
				currentSum := status[0] + status[1] + status[3] + 1
				for i := status[2] - 1; i >= 0; i-- {
					nextStatus := [4]int{status[0], status[1], i, status[3] + 1}
					s := currentSum + i
					if s > len(croakOfFrogs)-i-1 {
						continue
					}
					nextCount := max(s, count)
					if _, ok := next[nextStatus]; ok {
						next[nextStatus] = min(next[nextStatus], nextCount)
					} else {
						next[nextStatus] = nextCount
					}
				}
			case 'k':
				currentSum := status[0] + status[1] + status[2]
				for i := status[3] - 1; i >= 0; i-- {
					nextStatus := [4]int{status[0], status[1], status[2], i}
					s := currentSum + i
					if s > len(croakOfFrogs)-i-1 {
						continue
					}
					nextCount := max(s, count)
					if _, ok := next[nextStatus]; ok {
						next[nextStatus] = min(next[nextStatus], nextCount)
					} else {
						next[nextStatus] = nextCount
					}
				}
			}
		}
		dp = next
	}
	if _, ok := dp[[4]int{0, 0, 0, 0}]; !ok {
		return -1
	}
	return dp[[4]int{0, 0, 0, 0}]
}

func sum(arr [4]int) int {
	var s int
	for _, e := range arr {
		s += e
	}
	return s
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// ❌有効だったもののうち、かえるの数が一番少なかったもの
// 有効であるかを考えて、有効である場合、逆順に回して最悪匹数を計算する
// 追加された数を数える

// func minNumberOfFrogs(croakOfFrogs string) int {
// 	status := map[byte][2]int{}
// 	for i := range croakOfFrogs {
// 		switch croakOfFrogs[i] {
// 		case 'c':
// 			status['c'][0]++
// 			status['c'][1]++
// 		case 'r':
// 			if prev, ok := status['c']; !ok || prev[1] == 0 {
// 				return -1
// 			}
// 			status['c'][0] = 0
// 			status['c'][1]--
// 			status['r'][0]++
// 			status['r'][1]++
// 		case 'o':
// 			if prev, ok := status['r']; !ok || prev[1] == 0 {
// 				return -1
// 			}
// 			status['r'][0] = 0
// 			status['r'][1]--
// 			status['o'][0]++
// 			status['o'][1]++
// 		case 'a':
// 			if prev, ok := status['o']; !ok || prev[1] == 0 {
// 				return -1
// 			}
// 			status['o'][0] = 0
// 			status['o'][1]--
// 			status['a'][0]++
// 			status['a'][1]++
// 		case 'k':
// 			if prev, ok := status['a']; !ok || prev[1] == 0 {
// 				return -1
// 			}
// 			status['a'][0] = 0
// 			status['a'][1]--
// 		}
// 		// fmt.Println(len(status), status)
// 		// number = max(number, len(status))
// 	}
// 	for _, s := range status {
// 		if s[0] != 0 {
// 			return -1
// 		}
// 	}

// 	for i := len(croakOfFrogs) - 1; i >= 0; i-- {
// 		croakOfFrogs[i]
// 	}

// }

// // func minNumberOfFrogs(croakOfFrogs string) int {
// // 	status := map[byte]bool{}
// // 	number := 0
// // 	for i := range croakOfFrogs {
// // 		switch croakOfFrogs[i] {
// // 		case 'c':
// // 			status['c'] = true
// // 		case 'r':
// // 			if _, ok := status['c']; !ok {
// // 				return -1
// // 			}
// // 			delete(status, 'c')
// // 			status['r'] = true
// // 		case 'o':
// // 			if _, ok := status['r']; !ok {
// // 				return -1
// // 			}
// // 			delete(status, 'r')
// // 			status['o'] = true
// // 		case 'a':
// // 			if _, ok := status['o']; !ok {
// // 				return -1
// // 			}
// // 			delete(status, 'o')
// // 			status['a'] = true
// // 		case 'k':
// // 			if _, ok := status['a']; !ok {
// // 				return -1
// // 			}
// // 			delete(status, 'a')
// // 		}
// // 		// fmt.Println(len(status), status)
// // 		number = max(number, len(status))
// // 	}

// // 	if len(status) != 0 {
// // 		return -1
// // 	}

// // 	return number
// // }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
