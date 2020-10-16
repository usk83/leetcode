// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(findTheCity(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2))
	pp.Println(0)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	paths := map[int]map[int]int{}
	for _, edge := range edges {
		m := paths[edge[0]]
		if m == nil {
			paths[edge[0]] = map[int]int{
				edge[1]: edge[2],
			}
		} else {
			m[edge[1]] = edge[2]
			paths[edge[0]] = m
		}

		m = paths[edge[1]]
		if m == nil {
			paths[edge[1]] = map[int]int{
				edge[0]: edge[2],
			}
		} else {
			m[edge[0]] = edge[2]
			paths[edge[1]] = m
		}
	}

	id, count := -1, -1
	for i := 0; i < n; i++ {
		// pp.Println(i)
		reached := map[int]int{}
		cities := []int{}
		for to, distance := range paths[i] {
			if distance <= distanceThreshold {
				reached[to] = distance
				cities = append(cities, to)
			}
		}

		for len(cities) != 0 {
			next := []int{}
			for _, city := range cities {
				for to, distance := range paths[city] {
					if to == i {
						// fmt.Println("block", to)
						continue
					}
					// fmt.Println("go", to)
					if cur := reached[city] + distance; cur <= distanceThreshold {
						if prev, ok := reached[to]; ok {
							if cur < prev {
								reached[to] = cur
								next = append(next, to)
							}
						} else {
							reached[to] = cur
							next = append(next, to)
						}
					}
				}
			}
			cities = next
		}
		// fmt.Println(reached)
		if count == -1 || len(reached) <= count {
			id = i
			count = len(reached)
		}
	}

	return id
}
