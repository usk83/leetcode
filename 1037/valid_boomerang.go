// 1037. Valid Boomerang
// https://leetcode.com/problems/valid-boomerang/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isBoomerang([][]int{{0, 0}, {0, 2}, {2, 1}}))
	pp.Println(true)
	pp.Println("=========================================")
}

// points.length == 3
// points[i].length == 2
// 0 <= points[i][j] <= 100
func isBoomerang(points [][]int) bool {
	return (points[1][0]-points[0][0])*(points[2][1]-points[1][1]) != (points[1][1]-points[0][1])*(points[2][0]-points[1][0])
}

// func isBoomerang(points [][]int) bool {
// 	x1, y1, x2, y2, x3, y3 := points[0][0], points[0][1], points[1][0], points[1][1], points[2][0], points[2][1]
// 	if (x1 == x2 && y1 == y2) ||
// 		(x1 == x3 && y1 == x3) ||
// 		(x2 == x3 && y2 == x3) {
// 		return false
// 	}
// 	// a : b = c : d
// 	// a * d == b *c
// 	return (x2-x1)*(y3-y2) != (y2-y1)*(x2-x2)
// }

// The first solution
// func isBoomerang(points [][]int) bool {
// 	test := map[[2]int]bool{}
// 	for i := 0; i < 3; i++ {
// 		_, found := test[[2]int{points[i][0], points[i][1]}]
// 		if found {
// 			return false
// 		}
// 		test[[2]int{points[i][0], points[i][1]}] = true
// 	}

// 	// a : b = c : d
// 	// a * d == b *c
// 	a := points[1][0] - points[0][0]
// 	b := points[1][1] - points[0][1]

// 	c := points[2][0] - points[1][0]
// 	d := points[2][1] - points[1][1]

// 	return a*d != b*c
// }
