// x. xxx
// https://leetcode.com/problems/xxx/
package main

import "github.com/k0kubun/pp"

func main() {
	// pp.Println("=========================================")
	// pp.Println(numWays(3, 2))
	// pp.Println(4)
	// pp.Println("=========================================")
	// pp.Println(numWays(2, 4))
	// pp.Println(2)
	// pp.Println("=========================================")
	// pp.Println(numWays(4, 2))
	// pp.Println(8)
	// pp.Println("=========================================")
	// pp.Println(numWays(6, 13))
	// pp.Println(51)
	pp.Println("=========================================")
	pp.Println(numWays(27, 7))
	// pp.Println()
	pp.Println("=========================================")
}

// func calcWays(pointer, step, limit int) int {
// 	if pointer > step {
// 		return 0
// 	}

// 	if step == 0 {
// 		if pointer == 0 {
// 			return 1
// 		} else {
// 			return 0
// 		}
// 	}

// 	stay := calcWays(pointer, step-1, limit)
// 	var left, right int
// 	if pointer == 0 {
// 		right = calcWays(pointer+1, step-1, limit)
// 	} else if pointer >= limit {
// 		left = calcWays(pointer-1, step-1, limit)
// 	} else {
// 		right = calcWays(pointer+1, step-1, limit)
// 		left = calcWays(pointer-1, step-1, limit)
// 	}

// 	pp.Println(stay, left, right)

// 	return (stay + left + right) % 1000000007 // 10 ^ 9 + 7
// }

var cache map[int]int = map[int]int{}

func calcWays(pointer, step, leftLimit, rightLimit, limit int) int {
	if leftLimit == 0 && rightLimit == 0 {
		return 1
	}

	if pointer == 0 {
		if val, found := cache[step]; found {
			return val
		}
	}

	// stay := calcWays(pointer, step-1, limit)
	var left, right int
	if pointer == 0 || leftLimit == 0 {
		right = calcWays(pointer+1, step-1, leftLimit, rightLimit-1, limit)
	} else if pointer >= limit || rightLimit == 0 {
		left = calcWays(pointer-1, step-1, leftLimit-1, rightLimit, limit)
	} else {
		right = calcWays(pointer+1, step-1, leftLimit, rightLimit-1, limit)
		left = calcWays(pointer-1, step-1, leftLimit-1, rightLimit, limit)
	}

	return (left + right) % 1000000007 // 10 ^ 9 + 7
}

// 1 <= steps <= 500
// 1 <= arrLen <= 10^6
func numWays(steps int, arrLen int) int {
	for i := 2; i <= 500; i += 2 {
		result := calcWays(0, i, i/2, i/2, arrLen-1)
		cache[i] = result
		pp.Println(i, result)
	}

	return calcWays(0, steps, steps/2, steps/2, arrLen-1)
	// Since the answer may be too large, return it modulo 10^9 + 7.
	// return 0
}

// ====================================

// 2 ~ 250stepまでの辞書を作る？
// for steps := 2; steps <= 250; steps += 2 {
// 	var pointer int
// 	for i := 0; i < steps; i++ {
// 	}
// 	// pp.Println(i)
// }

// ====================================

// 全部留まる
// 2回で戻ってくる ~ step / 2回で戻ってくる
// 	→ 何回？
// 	→ いろいろなところでととどまれる可能性を考慮

// 2回で戻ってくる
// 	rl

// 4回で戻ってくる
// 	rrll
// 	rlrl

// 6回で戻ってくる
// 	rrrlll
// 	rrlrll
// 	rrllrl
// 	rlrlrl

// 8回で戻ってくる
// 	rrrrllll

// 	rrrlrlll
// 	rrrllrll
// 	rrrlllrl

// 	rrlrrlll
// 	rrlrlrll
// 	rrlrllrl
// 	rrlrllrl
// 	rrllrrll
// 	rrllrlrl
// 	rlrlrlrl
