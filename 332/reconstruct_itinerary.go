// 332. Reconstruct Itinerary
// https://leetcode.com/problems/reconstruct-itinerary/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 28
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/542/week-4-june-22nd-june-28th/3374/
func main() {
	pp.Println("=========================================")
	fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	fmt.Println([]string{"JFK", "MUC", "LHR", "SFO", "SJC"})
	pp.Println("=========================================")
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
	fmt.Println([]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"})
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func findItinerary(tickets [][]string) []string {
	flights := map[string][]string{}
	for _, ticket := range tickets {
		flights[ticket[0]] = append(flights[ticket[0]], ticket[1])
	}
	for _, flight := range flights {
		sort.Strings(flight)
	}
	itinerary := make([]string, len(tickets)+1)
	itinerary[0] = "JFK"
	for i := range itinerary[1:] {
		itinerary[i+1] = flights[itinerary[i]][0]
		flights[itinerary[i]] = flights[itinerary[i]][1:]
	}
	return itinerary
}
