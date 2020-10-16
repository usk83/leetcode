// 941. Valid Mountain Array
// https://leetcode.com/problems/valid-mountain-array/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(validMountainArray([]int{2, 1}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(validMountainArray([]int{3, 5, 5}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(validMountainArray([]int{0, 3, 2, 1}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(validMountainArray([]int{1, 1, 2, 3, 4, 5, 4, 3, 2, 1}))
	pp.Println(false)
	pp.Println("=========================================")
}

// 0 <= A.length <= 10000
// 0 <= A[i] <= 10000
// Using left&right pointers
func validMountainArray(A []int) bool {
	var left, right int
	for left = 0; left < len(A)-2 && A[left] < A[left+1]; left++ {
	}
	for right = len(A) - 1; right > 1 && A[right] < A[right-1]; right-- {
	}
	return left == right && left != 0
}

// Checking a trend
// func validMountainArray(A []int) bool {
// 	passedTop := false
// 	for i := len(A) - 1; i >= 2; i-- {
// 		switch {
// 		case A[i] < A[i-1] && A[i-1] < A[i-2]:
// 			if passedTop {
// 				return false
// 			}
// 		case A[i] > A[i-1] && A[i-1] > A[i-2]:
// 			if !passedTop {
// 				return false
// 			}
// 		case A[i] < A[i-1] && A[i-1] > A[i-2] && !passedTop:
// 			passedTop = true
// 		default:
// 			return false
// 		}
// 	}
// 	return passedTop
// }

// Going one direction (more beautiful)
// func validMountainArray(A []int) bool {
// 	var cursor int
// 	for cursor = 0; cursor < len(A)-1 && A[cursor] < A[cursor+1]; cursor++ {
// 	}
// 	if cursor == 0 || cursor == len(A)-1 {
// 		return false
// 	}
// 	for ; cursor < len(A)-1 && A[cursor] > A[cursor+1]; cursor++ {
// 	}
// 	return cursor == len(A)-1
// }

// Going one direction (foolish but straightforward)
// func validMountainArray(A []int) bool {
// 	if len(A) < 3 || A[0] >= A[1] {
// 		return false
// 	}
// 	goingUp := true
// 	for i := 2; i < len(A); i++ {
// 		if A[i] == A[i-1] {
// 			return false
// 		}
// 		if goingUp {
// 			if A[i] < A[i-1] {
// 				goingUp = false
// 			}
// 		} else {
// 			if A[i] > A[i-1] {
// 				return false
// 			}
// 		}
// 	}
// 	return !goingUp
// }

// Passed but incorrect solution
// func validMountainArray(A []int) bool {
//   if len(A) < 3 || A[0] >= A[1] {
//     return false
//   }
//   prev := A[1]
//   goingUp := true
//   for _, cur := range A[2:] {
//     if goingUp {
//       if cur < prev {
//         goingUp = false
//       }
//     } else {
//       if cur > prev || cur == prev {
//         return false
//       }
//     }
//     prev = cur
//   }
//   return !goingUp
// }
