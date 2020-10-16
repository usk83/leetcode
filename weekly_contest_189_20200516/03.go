// 5414. People Whose List of Favorite Companies Is Not a Subset of Another List
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

func peopleIndexes(favoriteCompanies [][]string) []int {
	// 会社名でそれをリストに入れている人を引けるようにする
	// 全ての会社について、他に入れている人をとる
	// 同じやつが会社の数出てたら出てたらだめ

	// not subsetの人のindexのリストを返す

	companyPeople := map[string][]int{}
	for i := range favoriteCompanies {
		for _, company := range favoriteCompanies[i] {
			companyPeople[company] = append(companyPeople[company], i)
		}
	}

	ret := []int{}
LOOP:
	for i := range favoriteCompanies {
		conf := map[int]int{}
		for _, company := range favoriteCompanies[i] {
			for _, other := range companyPeople[company] {
				if other == i {
					continue
				}
				conf[other]++
				if conf[other] == len(favoriteCompanies[i]) {
					continue LOOP
				}
			}
		}
		ret = append(ret, i)
	}

	return ret
}
