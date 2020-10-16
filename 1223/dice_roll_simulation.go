// 1223. Dice Roll Simulation
// https://leetcode.com/problems/dice-roll-simulation/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(dieSimulator(1, []int{1, 1, 2, 2, 2, 3}))
	pp.Println(6)
	pp.Println(dieSimulator(2, []int{1, 1, 2, 2, 2, 3}))
	pp.Println(34)
	pp.Println(dieSimulator(2, []int{1, 1, 1, 1, 1, 1}))
	pp.Println(30)
	pp.Println(dieSimulator(3, []int{1, 1, 1, 2, 2, 3}))
	pp.Println(181)
}

// 1 <= n <= 5000
// rollMax.length == 6
// 1 <= rollMax[i] <= 15
func dieSimulator(n int, rollMax []int) int {
	return 0
}
