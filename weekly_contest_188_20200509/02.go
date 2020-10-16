// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(countTriplets([]int{2, 3, 1, 6, 7}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(countTriplets([]int{1, 1, 1, 1, 1}))
	pp.Println(10)
	pp.Println("=========================================")
	pp.Println(countTriplets([]int{2, 3}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(countTriplets([]int{1, 3, 5, 7, 9}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(countTriplets([]int{7, 11, 12, 9, 5, 2, 7, 17, 22}))
	pp.Println(8)
	pp.Println("=========================================")
}

func countTriplets(arr []int) int {
	// 追跡和的な考え
	sumXOR := make([]int, len(arr))
	sumXOR[0] = arr[0]
	for i, num := range arr[1:] {
		index := i + 1
		sumXOR[index] = sumXOR[index-1] ^ num
	}

	// pp.Println("=== DEBUG START ======================================")
	// for i, num := range sumXOR {
	// 	fmt.Printf("%b : %b\n", arr[i], num)
	// }
	// pp.Println("=== DEBUG END ========================================")

	var count int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j; k < len(arr); k++ {
				var a, b int
				if i == 0 {
					a = sumXOR[j-1]
				} else {
					a = sumXOR[j-1] ^ sumXOR[i-1]
				}
				b = sumXOR[k] ^ sumXOR[j-1]
				if a == b {
					count++
				}
			}
		}
	}

	return count
}
