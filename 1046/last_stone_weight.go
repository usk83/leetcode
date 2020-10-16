// 1046. Last Stone Weight
// https://leetcode.com/problems/last-stone-weight/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	pp.Println(1)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// func lastStoneWeight(stones []int) int {
// }

// 1 <= stones.length <= 30
// 1 <= stones[i] <= 1000
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		fmt.Println(stones)
		var largest, larger int
		left := []int{}
		for _, stone := range stones {
			if stone > largest {
				largest, stone = stone, largest
			}
			if stone > larger {
				larger, stone = stone, larger
			}
			if stone != 0 {
				left = append(left, stone)
			}
		}
		if largest != larger {
			left = append(left, largest-larger)
		}
		stones = left
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}
