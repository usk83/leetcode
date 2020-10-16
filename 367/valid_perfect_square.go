// 367. Valid Perfect Square
// https://leetcode.com/problems/valid-perfect-square/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(1))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(4))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(9))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(16))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(14))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(2147483647))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(1<<63 - 1))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPerfectSquare(3037000499 * 3037000499))
	pp.Println(true)
	pp.Println("=========================================")
}

// Big number awareness binary search solution
// (Binary search solution considering a big number)
func isPerfectSquare(num int) bool {
	for small, big := 1, num; small <= big; {
		cur := small + (big-small)>>1
		div := num / cur
		if div == cur && num%cur == 0 {
			return true
		}
		if div > cur {
			small = cur + 1
		} else {
			big = cur - 1
		}
	}
	return false
}

// Binary search solution
// func isPerfectSquare(num int) bool {
// 	for small, big := 1, num; small <= big; {
// 		cur := (small + big) >> 1
// 		square := cur * cur
// 		if square == num {
// 			return true
// 		}
// 		if square < num {
// 			small = cur + 1
// 		} else {
// 			big = cur - 1
// 		}
// 	}
// 	return false
// }

// Shorten math solution
// func isPerfectSquare(num int) bool {
//  for i := 1; num > 0; num, i = num-i, i+2 {
//  }
//  return num == 0
// }

// Math solution
// func isPerfectSquare(num int) bool {
//  if num == 0 {
//    return false
//  }
//  i := 1
//  for num > 0 {
//    num -= i
//    i += 2
//  }
//  return num == 0
// }

// Original solution
// func isPerfectSquare(num int) bool {
//  // 半分以下でbinar search
//  if num == 1 {
//    return true
//  }
//  // small, big := 0, num/2
//  small, big := 0, num
//  for {
//    cur := (small + big) / 2
//    square := cur * cur
//    // pp.Println("=== DEBUG START ======================================")
//    // pp.Println(small, big, cur, square)
//    // pp.Println("=== DEBUG END ========================================")
//    if square == num {
//      return true
//    }
//    if square < num {
//      if small == cur {
//        small++
//      } else {
//        small = cur
//      }
//    } else {
//      if big == cur {
//        big--
//      } else {
//        big = cur
//      }
//    }
//    if small == big {
//      break
//    }
//  }
//  return false
// }
