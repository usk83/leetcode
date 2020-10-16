// 5276. Number of Burgers with No Waste of Ingredients
// https://leetcode.com/contest/weekly-contest-165/problems/number-of-burgers-with-no-waste-of-ingredients/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(numOfBurgers(16, 7))
	fmt.Println([]int{1, 6})
	pp.Println("=========================================")
	fmt.Println(numOfBurgers(17, 4))
	fmt.Println([]int{})
	pp.Println("=========================================")
	fmt.Println(numOfBurgers(4, 17))
	fmt.Println([]int{})
	pp.Println("=========================================")
	fmt.Println(numOfBurgers(0, 0))
	fmt.Println([]int{0, 0})
	pp.Println("=========================================")
	fmt.Println(numOfBurgers(2, 1))
	fmt.Println([]int{0, 1})
	pp.Println("=========================================")
	fmt.Println(numOfBurgers(2385088, 164763))
	fmt.Println([]int{})
	pp.Println("=========================================")

}

// Jumbo Burger: 4 tmato slices and 1 cheese slice.
// Small Burger: 2 Tmato slices and 1 cheese slice.

// 0 <= tomatoSlices <= 10^7
// 0 <= cheeseSlices <= 10^7
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	over := tomatoSlices - cheeseSlices*2

	if over < 0 || over%2 == 1 {
		return []int{}
	}

	jumbo := over / 2
	small := cheeseSlices - jumbo

	if small < 0 {
		return []int{}
	}

	return []int{jumbo, small}
}
