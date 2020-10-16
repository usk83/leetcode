// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(filterRestaurants([][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 1, 50, 10))
	fmt.Println([]int{3, 1, 5})
	pp.Println("=========================================")
	fmt.Println(filterRestaurants([][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 0, 50, 10))
	fmt.Println([]int{4, 3, 2, 1, 5})
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	filtered := [][2]int{}
	for _, restaurant := range restaurants {
		if veganFriendly == 1 && restaurant[2] == 0 {
			continue
		}

		if restaurant[3] <= maxPrice && restaurant[4] <= maxDistance {
			filtered = append(filtered, [2]int{restaurant[0], restaurant[1]})
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i][1] > filtered[j][1] || (filtered[i][1] == filtered[j][1] && filtered[i][0] > filtered[j][0])
	})

	ids := make([]int, len(filtered))

	for i, item := range filtered {
		ids[i] = item[0]
	}

	return ids
}
