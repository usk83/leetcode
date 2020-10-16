// 528. Random Pick with Weight
// https://leetcode.com/problems/random-pick-with-weight/
package main

import (
	"math/rand"
	"sort"
	"time"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 5
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3351/
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var s Solution
	pp.Println("=========================================")
	s = Constructor([]int{1})
	pp.Println(s.PickIndex())
	pp.Println("=========================================")
	s = Constructor([]int{1, 3})
	pp.Println(s.PickIndex())
	pp.Println(s.PickIndex())
	pp.Println(s.PickIndex())
	pp.Println(s.PickIndex())
	pp.Println(s.PickIndex())
	pp.Println("=========================================")
}

type Solution struct {
	cusum []int
}

func Constructor(w []int) Solution {
	for i := range w[1:] {
		w[i+1] += w[i]
	}
	return Solution{
		cusum: w,
	}
}

func (this *Solution) PickIndex() int {
	num := rand.Intn(this.cusum[len(this.cusum)-1])
	return sort.Search(len(this.cusum), func(i int) bool {
		return this.cusum[i] > num
	})
}
