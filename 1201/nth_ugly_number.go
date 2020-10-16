// 1201. Ugly Number III
// https://leetcode.com/problems/ugly-number-iii/
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(nthUglyNumber01(3, 2))
	fmt.Println(nthUglyNumber02(3, 2, 3, 5))
	fmt.Println(nthUglyNumber03(3, 11, 3, 5))
}

func nthUglyNumber01(n int, a int) int {
	return int(math.Pow(float64(a), float64(n)))
}

func nthUglyNumber02(n int, a int, b int, c int) int {
	nn, aa, bb, cc := float64(n), float64(a), float64(b), float64(c)
	min := math.Min(math.Min(aa, bb), cc)
	return int(math.Pow(min, nn))
}

func nthUglyNumber03(n int, a int, b int, c int) int {
	// ソートして再代入する
	vals := []int{a, b, c}
	sort.Ints(vals)
	a, b, c = vals[0], vals[1], vals[2]

	fn, fa, fb, fc := float64(n), float64(a), float64(b), float64(c)
	_ = fb
	_ = fc
	// return int(math.Pow(fa, fn))
	max := int(math.Pow(fa, fn))
	fmt.Println("max:", max)

	// しきい値までに他の2つの数字の倍数がいくつでてくるか
	fmt.Println(max / b)
	fmt.Println(max / c)

	// しかし、公約数を考慮する必要あり（x倍ごとにノーカウント）
	// 3つの値の最小公倍数までを最小値のカウントで考える
	// 最小公倍数は？

	// 積を最大公約数で割る

	fmt.Println(gcd(1071, 1029))

	// 他二つとの最小公倍数までを最小値のカウントで考える

	return 0
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func nthUglyNumber(n int, a int, b int, c int) int {
	return 0
}
