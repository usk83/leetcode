// 5441. Making File Names Unique
// https://leetcode.com/contest/weekly-contest-194/problems/making-file-names-unique/
package main

import (
	"strconv"
	"strings"

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

func getFolderNames(names []string) []string {
	result := make([]string, 0, len(names))
	used := map[string]map[string]bool{}
	lastCount := map[string]int{}
	for _, name := range names {
		/*
			使用済みでない
				そのまま
				使用済みにする
					数字削ったものも
			使用済み
				最後のカウントを探す
				カウントを追加して
				使用済みにする
		*/

		if used[name] == nil {
			result = append(result, name)
			used[name] = map[string]bool{
				"": true,
			}
			if str, num, ok := separate(name); ok {
				if d := used[str]; d == nil {
					used[str] = map[string]bool{
						strconv.Itoa(num): true,
					}
				} else {
					d[strconv.Itoa(num)] = true
					used[str] = d
				}
			}
		} else {
			num := 1
			if n, ok := lastCount[name]; ok {
				num = n
			}
			for used[name][strconv.Itoa(n)] {
				n++
			}
			rename := name + "(" + strconv.Itoa(n) + ")"
			result = append(result, rename)
			d := used[name]
			d[strconv.Itoa(n)] = true
			used[name] = d
			used[str] = n
		}

	}
	return result
}

func separate(str string) (string, int, bool) {
	if strings.HasSuffix(str, ")") {
		for i := len(str) - 2; i >= 0; i-- {
			if str[i] == '(' {
				n, err := strconv.Atoi(str[i+1 : len(str)-1])
				if err != nil {
					return str, 0, false
				}
				return str[:i], n, true
			}
		}
	}
	return str, 0, false
}
