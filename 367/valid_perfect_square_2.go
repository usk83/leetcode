// 367. Valid Perfect Square
// https://leetcode.com/problems/valid-perfect-square/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 9
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3324/
func main() {
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(16))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(14))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(361))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(1235124))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(77936 * 77936))
	pp.Println(true)
	pp.Println("=========================================")
}

/*
 * ref: https://leetcode.com/problems/valid-perfect-square/discuss/130010/Python-4-Methods-with-time-testing
 * 1. Solving with Bitwise trick.
 * NOTE: 各ビットのON/OFFを調べていく
 */
// func isPerfectSquare(num int) bool {
// 	const OS_BITS = 64 // or 32
// 	var sqrt int
// 	for bit := 1 << (ENV_BITS>>1 - 1); bit > 0; bit >>= 1 {
// 		sqrt |= bit
// 		if sqrt > num/sqrt {
// 			sqrt ^= bit
// 		}
// 	}
// 	return sqrt*sqrt == num
// }

/*
 * ref: https://leetcode.com/problems/valid-perfect-square/discuss/130010/Python-4-Methods-with-time-testing
 * 2. Using Newton's Method
 * NOTE: ニュートン法
 */
// func isPerfectSquare(num int) bool {
// 	xn := num
// 	for xn > num/xn {
// 		xn = (xn + num/xn) >> 1
// 	}
// 	return xn*xn == num
// }

/*
 * ref: https://leetcode.com/problems/valid-perfect-square/discuss/130010/Python-4-Methods-with-time-testing
 * 3. Math Trick for Square number is 1+3+5+ ... +(2n-1)
 * NOTE: 第二階差数列が等差数列になる
 * 0 1 4 9 16 25 36 49 64 81
 *  1 3 5 7  9  11 13 15
 *   2 2 2  2  2  2  2
 */
// func isPerfectSquare(num int) bool {
// 	for i := 1; num > 0; i += 2 {
// 		num -= i
// 	}
// 	return num == 0
// }

/*
 * Big number awareness version
 */
func isPerfectSquare(num int) bool {
	sqrt := sort.Search(num, func(i int) bool {
		return i != 0 && i >= num/i
	})
	return sqrt*sqrt == num
}

/*
 The first solution
*/
// func isPerfectSquare(num int) bool {
// 	sqrt := sort.Search(num, func(i int) bool {
// 		return i*i >= num
// 	})
// 	return sqrt*sqrt == num
// }
