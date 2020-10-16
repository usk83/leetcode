// 9. Palindrome Number
// https://leetcode.com/problems/palindrome-number/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isPalindrome(121))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPalindrome(-858))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPalindrome(10))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPalindrome(123454321))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPalindrome(123456789))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isPalindrome(1234554321))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPalindrome(1234564321))
	pp.Println(false)
	pp.Println("=========================================")
}

// 一度配列に変換せずに特定の数字をチェックしていく方法
// - Time Complexity: O(n)
// - Space Complexity: O(1)
// > Runtime: 4 ms, faster than 99.27%
// > Memory Usage: 5 MB, less than 100.00%
// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}
// 	for base := int(math.Pow10(int(math.Log10(float64(x))))); base >= 10; base /= 100 {
// 		if x/base != x%10 {
// 			return false
// 		}
// 		x = x % base / 10
// 	}
// 	return true
// }

// オーバーフローの恐れなし
// - Time Complexity: O(n)
// - Space Complexity: O(1)
// > Runtime: 4 ms, faster than 99.27% （平均はオーバーフローを考慮していないものより遅い）
// > Memory Usage: 5 MB, less than 100.00%
// func isPalindrome(x int) bool {
// 	if x < 0 || (x != 0 && x%10 == 0) {
// 		return false
// 	}
// 	var rev int
// 	for x > rev {
// 		rev = rev*10 + x%10
// 		x = x / 10
// 	}
// 	return x == rev || x == rev/10
// }

// オーバーフローの恐れがあるが問題はない
// https://leetcode.com/problems/palindrome-number/discuss/5127/9-line-accepted-Java-code-without-the-need-of-handling-overflow/5869
// 厳密にはオーバーフロー時の挙動は処理系依存？
// - Time Complexity: O(n)
// - Space Complexity: O(1)
// > Runtime: 4 ms, faster than 99.27%
// > Memory Usage: 5 MB, less than 100.00%
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original, reversed := x, 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}

// - Time Complexity: O(n)
// - Space Complexity: O(n)
// > Runtime: 8 ms, faster than 95.53%
// > Memory Usage: 6.1 MB, less than 25.00%
// func isPalindrome(x int) bool {
//  if x < 0 {
//    return false
//  }
//  digits := []int{}
//  for x != 0 {
//    digits = append(digits, x%10)
//    x /= 10
//  }
//  numberOfDigits := len(digits)
//  lastIndex := numberOfDigits - 1
//  for i := numberOfDigits/2 - 1; i >= 0; i-- {
//    if digits[i] != digits[lastIndex-i] {
//      return false
//    }
//  }
//  return true
// }
