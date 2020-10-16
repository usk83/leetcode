// 1071. Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(gcdOfStrings("ABCABC", "ABC"))
	pp.Println("ABC")
	pp.Println("=========================================")
	pp.Println(gcdOfStrings("ABABAB", "ABAB"))
	pp.Println("AB")
	pp.Println("=========================================")
	pp.Println(gcdOfStrings("LEET", "CODE"))
	pp.Println("")
	pp.Println("=========================================")
	pp.Println(gcdOfStrings("ABCDEF", "ABC"))
	pp.Println("")
	pp.Println("=========================================")
	pp.Println(gcdOfStrings("ABAB", "ABABAB"))
	pp.Println("AB")
	pp.Println("=========================================")
}

// 1 <= str1.length <= 1000
// 1 <= str2.length <= 1000
// str1[i] and str2[i] are English uppercase letters.
/* Euclidean Algorithm */
func gcdOfStrings(str1, str2 string) string {
	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}
	for str2 != "" {
		if div, mod := divModString(str1, str2); div == 0 {
			return ""
		} else {
			str1, str2 = str2, mod
		}
	}
	return str1
}

func divModString(a, b string) (int, string) {
	count, max := 0, len(a)/len(b)
	for count < max {
		if strings.HasPrefix(a[len(b)*count:], b) != true {
			break
		}
		count++
	}
	return count, a[len(b)*count:]
}

/* Without Euclidean Algorithm */
// func gcdOfStrings(str1, str2 string) string {
// 	if len(str2) > len(str1) {
// 		str1, str2 = str2, str1
// 	}
//
// 	deviders := calcDevidersDesc(len(str2))
// 	for i := 0; i < len(deviders); i++ {
// 		if deviders[i] == 0 {
// 			break
// 		}
// 		if !testIfDivisible(str2[deviders[i]:], str2[:deviders[i]]) {
// 			removeCommonNumber(deviders[i+1:], calcDevidersDesc(i)[1:])
// 			continue
// 		}
// 		if !testIfDivisible(str1, str2[:deviders[i]]) {
// 			continue
// 		}
// 		return str2[:deviders[i]]
// 	}
// 	return ""
// }
//
// func calcDevidersDesc(num int) []int {
// 	devidersDesc := []int{}
// 	devidersAsc := []int{}
// 	for i := 1; i*i <= num; i++ {
// 		if num%i == 0 {
// 			devidersDesc = append(devidersDesc, num/i)
// 			if i*i != num {
// 				devidersAsc = append(devidersAsc, i)
// 			}

// 		}
// 	}
// 	deviders := make([]int, 0, len(devidersDesc)+len(devidersAsc))
// 	for _, devider := range devidersDesc {
// 		deviders = append(deviders, devider)
// 	}
// 	for i := len(devidersAsc) - 1; i >= 0; i-- {
// 		deviders = append(deviders, devidersAsc[i])
// 	}
// 	return deviders
// }
//
// func testIfDivisible(target, devider string) bool {
// 	count, max := 0, len(target)/len(devider)
// 	for count < max {
// 		if strings.HasPrefix(target[len(devider)*count:], devider) != true {
// 			break
// 		}
// 		count++
// 	}
// 	return len(target) == len(devider)*count
// }
//
// func removeCommonNumber(nums1, nums2 []int) {
// 	var i2 int
// 	for i1 := range nums1 {
// 		if nums1[i1] == 0 {
// 			break
// 		}
// 		if i2 < len(nums2) && nums1[i1] == nums2[i2] {
// 			nums1[i1] = 0
// 			i2++
// 		} else if i2 != 0 {
// 			nums1[i1-i2], nums1[i1] = nums1[i1], 0
// 		}
// 	}
// }

// Hint:
func gcdOfNumberRecursively(x, y int) int {
	if y == 0 {
		return x
	}
	return gcdOfNumberRecursively(y, x%y)
}

func gcdOfNumberIteratively(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
