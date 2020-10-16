// 412. Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/
package main

import (
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

func fizzBuzz(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var s string
		switch {
		case i%3 == 0 && i%5 == 0:
			s = "FizzBuzz"
		case i%3 == 0:
			s = "Fizz"
		case i%5 == 0:
			s = "Buzz"
		default:
			s = strconv.Itoa(i)
		}
		result = append(result, s)
	}
	return result
}
