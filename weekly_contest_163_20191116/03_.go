func maxSumDivThree(nums []int) int {
	result := 0

	// 3の倍数は絶対割れる

	// あまりのペアを作成
	// 大きい順にソート

	// one := make([]int, 0, len(nums))
	// two := make([]int, 0, len(nums))
	both := make([]int, 0, len(nums))

	oneMin1 := 0
	oneMin2 := 0
	twoMin1 := 0
	twoMin2 := 0

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
			result += num
			// one = append(one, num)
			// both = append(both, num)
			fallthrough
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
			result += num
			// two = append(two, num)
			// both = append(both, num)
		}
	}

	// // ソート

	// 1 1 1
	// 2 2 2
	// 1 2

	// 最後2つ
	// 1 2 のみ
	// 最後3つ

	// 最大で余る数
	// // 1 1
	// // 2 2
	// 1 1 2 2
	// 2こ？

	// 11
	// 22

	// それぞれ2つ残せば良い

	// for i := 0; i < len(both)-2; i += 3 {

	// 	a := both[i]
	// 	b := both[i+1]
	// 	c := both[i+2]

	// 	if a%3 == 1 {
	// 		one = one[1:]
	// 	} else {
	// 		two = two[1:]
	// 	}

	// 	if b%3 == 1 {
	// 		one = one[1:]
	// 	} else {
	// 		two = two[1:]
	// 	}

	// 	if c%3 == 1 {
	// 		one = one[1:]
	// 	} else {
	// 		two = two[1:]
	// 	}

	// 	result += a
	// 	result += b
	// 	result += c
	// }

	// count := len(one)
	// if len(two) > count {
	// 	count = len(two)
	// }

	// for i := 0; i < count; i++ {
	// 	onw[i]
	// 	two[i]
	// }

	return result
}
