// 1375. Bulb Switcher III
// https://leetcode.com/problems/bulb-switcher-iii/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(numTimesAllBlue([]int{2, 1, 3, 5, 4}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(numTimesAllBlue([]int{3, 2, 4, 1, 5}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(numTimesAllBlue([]int{4, 1, 2, 3}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numTimesAllBlue([]int{2, 1, 4, 3, 6, 5}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(numTimesAllBlue([]int{1, 2, 3, 4, 5, 6}))
	pp.Println(6)
	pp.Println("=========================================")
}

// n == light.length
// 1 <= n <= 5 * 10^4
// light is a permutation of  [1, 2, ..., n]

/*
 * v2. Space O(1) solution
 * ref: https://leetcode.com/problems/bulb-switcher-iii/discuss/532538/JavaC%2B%2BPython-Straight-Forward-O(1)-Space
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 56 ms, faster than 100.00%
 * > Memory Usage: 7 MB, less than 100.00%
 */
// func numTimesAllBlue(lights []int) int {
// 	var pos, count int
// 	for i := range lights {
// 		if lights[i] > pos {
// 			pos = lights[i]
// 		}
// 		if pos == i+1 {
// 			count++
// 		}
// 	}
// 	return count
// }

/*
 * v1. My solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: 56 ms, faster than 100.00%
 * > Memory Usage: 7.1 MB, less than 100.00%
 */
func numTimesAllBlue(lights []int) int {
	bulbs := make([]bool, len(lights)+1)
	bulbs[0] = true
	var pos, count int
	for i, light := range lights {
		bulbs[light] = true
		if light == pos+1 {
			pos++
			for pos < len(lights) && bulbs[pos+1] {
				pos++
			}
		}
		if pos == i+1 {
			count++
		}
	}
	return count
}
