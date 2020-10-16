// 1184. Distance Between Bus Stops
// https://leetcode.com/problems/distance-between-bus-stops/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
	fmt.Println("Expected: 3")

	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
	fmt.Println("Expected: 4")
}

// 1 <= n <= 10^4
// distance.length == n
// 0 <= start, destination < n
// 0 <= distance[i] <= 10^4
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	// Done: Clockwiseに計算する
	// Done: Counterclockwiseに計算する
	// Todo: 同時並行にやって確定次第返却する
	n := len(distance)

	counterClockwise := func(i int) int {
		if i = i - 1; i >= 0 {
			return i
		}
		return n - 1
	}

	clockwiseDistance := 0
	for i := start; i != destination; i = (i + 1) % n {
		clockwiseDistance += distance[i]
	}

	counterClockwiseDistance := 0
	i := counterClockwise(start)
	for {
		counterClockwiseDistance += distance[i]
		if i == destination {
			break
		}
		i = counterClockwise(i)
	}

	if clockwiseDistance < counterClockwiseDistance {
		return clockwiseDistance
	}
	return counterClockwiseDistance
}
