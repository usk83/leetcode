// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"fmt"
	"strconv"

	"github.com/k0kubun/pp"
)

// 10
// 1000000000
// 100
// 300
// 1000
// 13000

func main() {
	// pp.Println("=========================================")
	// fmt.Println(sequentialDigits(100, 300))
	// fmt.Println([]int{123, 234})
	// pp.Println("=========================================")
	// fmt.Println(sequentialDigits(1000, 13000))
	// fmt.Println([]int{1234, 2345, 3456, 4567, 5678, 6789, 12345})
	pp.Println("=========================================")
	fmt.Println(sequentialDigits(10, 1000000000))
	// fmt.Println([]int{})
	pp.Println("=========================================")
}

// Constraints:
//     10 <= low <= high <= 10^9
func sequentialDigits(low int, high int) []int {
	result := []int{}

	lowStr := strconv.Itoa(low)
	head, _ := strconv.Atoi(string(lowStr[0]))
	count := len(lowStr)

	for {
		if count == 10 {
			break
		}
		num := genNum(count, head)
		head++
		if head == 10 || !isValid(num) {
			count++
			head = 1
			continue
		}
		if num < low {
			continue
		}
		if num > high {
			break
		}
		result = append(result, num)
	}
	return result
}

func genNum(count, start int) int {
	str := ""
	for i := 0; i < count; i++ {
		str += strconv.Itoa(start)
		start = (start + 1) % 10
	}
	num, _ := strconv.Atoi(str)
	return num
}

func isValid(num int) bool {
	str := strconv.Itoa(num)
	prev, _ := strconv.Atoi(string(str[0]))
	for i := 1; i < len(str); i++ {
		cur, _ := strconv.Atoi(string(str[i]))
		if cur != prev+1 {
			return false
		}
		prev = cur
	}
	return true
}
