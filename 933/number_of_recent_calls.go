// 933. Number of Recent Calls
// https://leetcode.com/problems/number-of-recent-calls/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: []int{},
	}
}

/*
 * v2. Binary search solution
 */
func (rc *RecentCounter) Ping(t int) int {
	lowerBound := t - 3000
	i := sort.Search(len(rc.pings), func(i int) bool {
		return rc.pings[i] >= lowerBound
	})
	rc.pings = append(rc.pings[i:], t)
	return len(rc.pings)
}

/*
 * v1. Linear search solution
 */
func (rc *RecentCounter) Ping(t int) int {
	i, lowerBound := 0, t-3000
	for i < len(rc.pings) && rc.pings[i] < lowerBound {
		i++
	}
	rc.pings = append(rc.pings[i:], t)
	return len(rc.pings)
}

/*
 * The first solution
 */
// func (this *RecentCounter) Ping(t int) int {
// 	newPings := []int{}
// 	var count int
// 	for _, ping := range this.pings {
// 		if ping >= t-3000 {
// 			count++
// 			newPings = append(newPings, ping)
// 		}
// 	}
// 	newPings = append(newPings, t)
// 	count++
// 	this.pings = newPings
// 	return count
// }
