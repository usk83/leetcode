// 5178. Four Divisors
// https://leetcode.com/contest/weekly-contest-181/problems/four-divisors/
package main

import (
	"math"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(sumFourDivisors([]int{21, 4, 7}))
	pp.Println(32)
	pp.Println("=========================================")
	pp.Println(sumFourDivisors([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	pp.Println(45)
	pp.Println("=========================================")
}

func sumFourDivisors(nums []int) int {
	var ans int
LOOP:
	for _, num := range nums {
		// pp.Println("num:", num)
		var sum, count int
		root := int(math.Sqrt(float64(num)))
		div := num / root
		mod := num % root
		if div == root && mod == 0 {
			// pp.Println("NEVER:", root)
			continue
		}
		for d := 1; d <= root; d++ {
			if mod := num % d; mod == 0 {
				count += 2
				if count > 4 {
					continue LOOP
				}
				// pp.Println("d:", d, "num/d:", num/d)
				sum += d + num/d
			}
		}
		if count == 4 {
			// pp.Println("OK")
			ans += sum
		}
	}
	return ans
}
