package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
	pp.Println(18)
	pp.Println("=========================================")
	pp.Println(maxSumDivThree([]int{4}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(maxSumDivThree([]int{1, 2, 3, 4, 4}))
	pp.Println(12)
	pp.Println("=========================================")
}

func maxSumDivThree(nums []int) int {
	result := 0

	oneMin1 := 0
	oneMin2 := 0
	twoMin1 := 0
	twoMin2 := 0
	one := 0
	two := 0
	for _, num := range nums {
		switch num % 3 {
		case 0:
			result += num
		case 1:
			if oneMin1 == 0 {
				oneMin1 = num
				continue
			}
			if oneMin1 > num {
				oneMin1, num = num, oneMin1
			}
			if oneMin2 == 0 {
				oneMin2 = num
				continue
			}
			if oneMin2 > num {
				oneMin2, num = num, oneMin2
			}
			one++
			result += num
		case 2:
			if twoMin1 == 0 {
				twoMin1 = num
				continue
			}
			if twoMin1 > num {
				twoMin1, num = num, twoMin1
			}
			if twoMin2 == 0 {
				twoMin2 = num
				continue
			}
			if twoMin2 > num {
				twoMin2, num = num, twoMin2
			}
			two++
			result += num
		}
	}

	var diff int
	if one > two {
		diff = one - two
	} else {
		diff = two - one
	}

	diff = diff % 3

	// if diff == 0 {
	// 	result += oneMin1
	// 	result += oneMin2
	// 	result += twoMin1
	// 	result += twoMin2
	// } else if diff == 1 {
	// 	result += oneMin2
	// 	result += twoMin2
	// 	if oneMin1 > twoMin1 {
	// 		result += oneMin1
	// 	} else {
	// 		result += twoMin2
	// 	}
	// } else if diff == 2 {
	// 	if oneMin1+oneMin2 > twoMin1+twoMin2 {
	// 		result += oneMin1
	// 		result += oneMin2
	// 	} else {
	// 		result += twoMin1
	// 		result += twoMin2
	// 	}
	// }

	return result
}

// func maxSumDivThree(nums []int) int {
// 	result := 0

// 	// 111 111 11
// 	// 222 222 22

// 	// 1 1 1 1
// 	// 2 2 2

// 	oneMin1 := 0
// 	oneMin2 := 0
// 	oneMin3 := 0
// 	oneMin4 := 0
// 	oneMin5 := 0
// 	twoMin1 := 0
// 	twoMin2 := 0
// 	twoMin3 := 0
// 	twoMin4 := 0
// 	twoMin5 := 0
// 	one := 0
// 	two := 0
// 	for _, num := range nums {
// 		switch num % 3 {
// 		case 0:
// 			result += num
// 		case 1:
// 			if oneMin1 == 0 {
// 				oneMin1 = num
// 				continue
// 			}
// 			if oneMin1 > num {
// 				oneMin1, num = num, oneMin1
// 			}
// 			if oneMin2 == 0 {
// 				oneMin2 = num
// 				continue
// 			}
// 			if oneMin2 > num {
// 				oneMin2, num = num, oneMin2
// 			}
// 			if oneMin3 == 0 {
// 				oneMin3 = num
// 				continue
// 			}
// 			if oneMin3 > num {
// 				oneMin3, num = num, oneMin3
// 			}
// 			if oneMin4 == 0 {
// 				oneMin4 = num
// 				continue
// 			}
// 			if oneMin4 > num {
// 				oneMin4, num = num, oneMin4
// 			}
// 			if oneMin5 == 0 {
// 				oneMin5 = num
// 				continue
// 			}
// 			if oneMin5 > num {
// 				oneMin5, num = num, oneMin5
// 			}
// 			one++
// 			result += num
// 		case 2:
// 			if twoMin1 == 0 {
// 				twoMin1 = num
// 				continue
// 			}
// 			if twoMin1 > num {
// 				twoMin1, num = num, twoMin1
// 			}
// 			if twoMin2 == 0 {
// 				twoMin2 = num
// 				continue
// 			}
// 			if twoMin2 > num {
// 				twoMin2, num = num, twoMin2
// 			}
// 			if twoMin3 == 0 {
// 				twoMin3 = num
// 				continue
// 			}
// 			if twoMin3 > num {
// 				twoMin3, num = num, twoMin3
// 			}
// 			if twoMin4 == 0 {
// 				twoMin4 = num
// 				continue
// 			}
// 			if twoMin4 > num {
// 				twoMin4, num = num, twoMin4
// 			}
// 			if twoMin5 == 0 {
// 				twoMin5 = num
// 				continue
// 			}
// 			if twoMin5 > num {
// 				twoMin5, num = num, twoMin5
// 			}
// 			two++
// 			result += num
// 		}
// 	}

// 	ones := []int{}
// 	if oneMin1 != 0 {
// 		ones = append([]int{oneMin1}, ones...)
// 	}
// 	if oneMin2 != 0 {
// 		ones = append([]int{oneMin2}, ones...)
// 	}
// 	if oneMin3 != 0 {
// 		ones = append([]int{oneMin3}, ones...)
// 	}
// 	if oneMin4 != 0 {
// 		ones = append([]int{oneMin4}, ones...)
// 	}
// 	if oneMin5 != 0 {
// 		ones = append([]int{oneMin5}, ones...)
// 	}

// 	twos := []int{}
// 	if twoMin1 != 0 {
// 		twos = append([]int{twoMin1}, twos...)
// 	}
// 	if twoMin2 != 0 {
// 		twos = append([]int{twoMin2}, twos...)
// 	}
// 	if twoMin3 != 0 {
// 		twos = append([]int{twoMin3}, twos...)
// 	}
// 	if twoMin4 != 0 {
// 		twos = append([]int{twoMin4}, twos...)
// 	}
// 	if twoMin5 != 0 {
// 		twos = append([]int{twoMin5}, twos...)
// 	}

// 	if one%3 == 1 {
// 		result += ones[0]
// 		result += ones[1]
// 		ones = ones[2:]
// 	}
// 	if one%3 == 2 {
// 		result += ones[0]
// 		ones = ones[1:]
// 	}

// 	if two%3 == 1 {
// 		result += twos[0]
// 		result += twos[1]
// 		twos = twos[2:]
// 	}
// 	if two%3 == 2 {
// 		result += twos[0]
// 		twos = twos[1:]
// 	}

// 	if len(ones) == len(twos) {
// 		for _, n := range ones {
// 			result += n
// 		}
// 		for _, n := range twos {
// 			result += n
// 		}
// 		return result
// 	}

// 	差が3以上ある場合

// 	差をみて
// 	多くある方が残る

// 	もしくは片方から3取って
// 	差をみて
// 	大きいほうが残る

// 	もしくは片方から3取って
// 	差をみて
// 	大きいほうが残る

// 	// switch {
// 	// case len(ones) == 3 && len(twos) == 2:
// 	// 	if ones[2] > twos[0]+twos[1] {
// 	// 		result += ones[0]
// 	// 		result += ones[1]
// 	// 		result += ones[2]
// 	// 	} else {
// 	// 		result += ones[0]
// 	// 		result += ones[1]
// 	// 		result += twos[0]
// 	// 		result += twos[1]
// 	// 	}
// 	// case len(ones) == 2 && len(twos) == 3:
// 	// 	if twos[2] > ones[0]+ones[1] {
// 	// 		result += twos[0]
// 	// 		result += twos[1]
// 	// 		result += twos[2]
// 	// 	} else {
// 	// 		result += twos[0]
// 	// 		result += twos[1]
// 	// 		result += ones[0]
// 	// 		result += ones[1]
// 	// 	}
// 	// case len(ones) == 3 && len(twos) == 1:
// 	// 	if twos[0] > ones[1]+ones[2] {
// 	// 		result += twos[0]
// 	// 		result += ones[0]
// 	// 	} else {
// 	// 		result += ones[0]
// 	// 		result += ones[1]
// 	// 		result += ones[2]
// 	// 	}
// 	// case len(ones) == 1 && len(twos) == 3:
// 	// 	if ones[0] > twos[1]+twos[2] {
// 	// 		result += ones[0]
// 	// 		result += twos[0]
// 	// 	} else {
// 	// 		result += twos[0]
// 	// 		result += twos[1]
// 	// 		result += twos[2]
// 	// 	}
// 	// case len(ones) == 1 && len(twos) == 2:
// 	// 	fallthrough
// 	// case len(ones) == 2 && len(twos) == 1:
// 	// 	result += ones[0]
// 	// 	result += twos[0]
// 	// }

// 	// var oneReft int
// 	// var twoReft int
// 	// if ones%3 == 1 {
// 	// 	oneReft = ones[len(ones)-1]
// 	// } else if ones%3 == 2 {
// 	// 	oneReft = ones[len(ones)-1] + ones[len(ones)-2]
// 	// } else {
// 	// 	oneReft = ones[len(ones)-1] + ones[len(ones)-2] + ones[len(ones)-3]
// 	// }

// 	// if twos%3 == 1 {
// 	// 	twoReft = twos[len(twos)-1]
// 	// } else if twos%3 == 2 {
// 	// 	twoReft = twos[len(twos)-1] + twos[len(twos)-2]
// 	// } else {
// 	// 	twoReft = twos[len(twos)-1] + twos[len(twos)-2] + twos[len(twos)-3]
// 	// }

// 	return result
// }
