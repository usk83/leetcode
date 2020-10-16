package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(dieSimulator(1, []int{1, 1, 2, 2, 2, 3}))
	pp.Println(6)
	pp.Println(dieSimulator(2, []int{1, 1, 2, 2, 2, 3}))
	pp.Println(34)
	pp.Println(dieSimulator(2, []int{1, 1, 1, 1, 1, 1}))
	pp.Println(30)
	pp.Println(dieSimulator(3, []int{1, 1, 1, 2, 2, 3}))
	pp.Println(181)
}

// 1 <= n <= 5000
// rollMax.length == 6
// 1 <= rollMax[i] <= 15
func dieSimulator(n int, rollMax []int) int {
	// V3 DP
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 6
	}

	records := make([][6][2]int, n)
	for i := range records[0] {
		records[0][i] = [2]int{0, rollMax[i] - 1}
	}

	for i := range records[1:] {
		prevSum := 0
		for _, record := range records[n-1] {
			prevSum += record[0]
		}

		for j := range records[i] {
			基本的には前6つを足したもの
			前のカウントが0だったら自分の分を減らす
			rollMax[j] - 1
			連続のカウントは1減らす
			rollMax[j] - 1

			// nthRecords[j] = []int{0, rollMax[j] - 1}
		}
	}

	sum := 0
	for _, record := range records[n-1] {
		sum += record[0]
	}

	return sum

	/*
		V2
		n回振った方がよさそう
		数字が出るごとにMaxをマイナス1して
		前回の数字を覚えて
			連続したらさらにマイナス1する
			連続しなければもとに戻してマイナス1する
	*/

	/*
		V1
		// 6のn上が通常
		full := int(math.Pow(6, float64(n)))
		comb := full

		// この時点ではn回連続も含まれている
		// しかし連続はrollMax[i]までしか許されていないため
		// rollMax[i]を超えているものを除く必要がある

		// n - rollMax[i] <= 0: 制約なし
		// n - rollMax[i] == 1: 1パターンだめ
		// n - rollMax[i] == 2: 1パターン + 10

		for _, max := range rollMax {
			if n-max <= 0 {
				continue
			}
			if n-max > 0 {
				comb -= 1
			}
			if n-max > 1 {
				comb -= 10
			}
		}

		両端パターンだけ例外

		return comb % (int(math.Pow(10, 9)) + 7)
	*/
}
