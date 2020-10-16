// 5353. Bulb Switcher III
// https://leetcode.com/contest/weekly-contest-179/problems/bulb-switcher-iii/
package main

import (
	"github.com/k0kubun/pp"
)

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

func numTimesAllBlue(lights []int) int {
	bulbs := make([]bool, len(lights)+1)
	bulbs[0] = true
	currentPosition := 0
	count := 0
	for i, light := range lights {
		bulbs[light] = true
		if light == currentPosition+1 {
			currentPosition++
			for currentPosition < len(lights) && bulbs[currentPosition+1] {
				currentPosition++
			}
		}
		if currentPosition == i+1 {
			count++
		}
	}
	return count
}
