// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(fib(1))
	pp.Println("=========================================")
	pp.Println(fib(2))
	pp.Println("=========================================")
	pp.Println(fib(3))
	pp.Println("=========================================")
	pp.Println(fib(4))
	pp.Println("=========================================")
	pp.Println(fib(5))
	pp.Println("=========================================")

	// pp.Println("=========================================")
	// pp.Println(fib(10))
	// pp.Println(55)
	// pp.Println("=========================================")
	// pp.Println(fib(15))
	// pp.Println(610)
	// pp.Println("=========================================")
	// pp.Println(fib(22))
	// pp.Println(17711)
	// pp.Println("=========================================")
	// pp.Println(fib(30))
	// pp.Println(832040)
	// pp.Println("=========================================")
}

func fib(N int) int {
	if N <= 1 {
		return N
	}
	var A = [2][2]int{
		{1, 1},
		{1, 0},
	}
	A = matrixPower(A, N-1)
	return A[0][0]
}

func matrixPower(A [2][2]int, N int) [2][2]int {
	if N <= 1 {
		return A
	}
	A = matrixPower(A, N/2)
	A = multiply(A, A)

	var B = [2][2]int{
		{1, 1},
		{1, 0},
	}
	if N%2 != 0 {
		A = multiply(A, B)
	}

	pp.Println("N:", N)
	fmt.Println(A)

	return A
}

func multiply(A [2][2]int, B [2][2]int) [2][2]int {
	x := A[0][0]*B[0][0] + A[0][1]*B[1][0]
	y := A[0][0]*B[0][1] + A[0][1]*B[1][1]
	z := A[1][0]*B[0][0] + A[1][1]*B[1][0]
	w := A[1][0]*B[0][1] + A[1][1]*B[1][1]

	A[0][0] = x
	A[0][1] = y
	A[1][0] = z
	A[1][1] = w

	return A
}

/*
 * DP solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms, faster than 100.00%
 * > Memory Usage: 1.9 MB, less than 100.00%
 */
// func fib(n int) int {
// 	fibs := [3]int{0, 1, 1}
// 	for i := 3; i <= n; i++ {
// 		fibs[i%3] = fibs[(i-1)%3] + fibs[(i-2)%3]
// 	}
// 	return fibs[n%3]
// }
