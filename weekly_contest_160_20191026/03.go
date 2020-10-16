package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxLength([]string{"un", "iq", "ue"}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(maxLength([]string{"cha", "r", "act", "ers"}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
	pp.Println(26)
	pp.Println("=========================================")
	pp.Println(maxLength([]string{"yy", "bkhwmpbiisbldzknpm"}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(maxLength([]string{"nzxijudcseqmybar", "javz", "fqbtasiplekwnx", "duhzjlkwyo", "hsmguvln", "ncaegmvxi", "xrszqhwylmtbocjgv", "ybivuerzwohqfjnlpa", "jnqchxfagkumsiv", "svxknebywoizramc", "xtowlghkfrjan", "ol"}))
	pp.Println(19)
	pp.Println("=========================================")
}

// 1 <= arr.length <= 16
// 1 <= arr[i].length <= 26
// arr[i] contains only lower case English letters.
func maxLength(arr []string) int {
	if len(arr) == 0 {
		return 0
	}

	// 全部のレングスを取って
	// 長さとmapに変換する
	type set struct {
		length int
		chars  map[rune]struct{}
	}

	setArr := make([]set, len(arr))
	for i, a := range arr {
		l := len(a)
		cm := map[rune]struct{}{}
		check := true
		for _, c := range a {
			_, exist := cm[c]
			if exist {
				check = false
				break
			}
			cm[c] = struct{}{}
		}
		if check {
			setArr[i] = set{
				l,
				cm,
			}
		}
	}

	// 長い順に長さでソート
	sort.Slice(setArr, func(i, j int) bool {
		return setArr[i].length > setArr[j].length
	})

	type combi struct {
		length   int
		charsArr []map[rune]struct{}
	}

	longestLength := 0
	curCombi := &combi{
		0,
		[]map[rune]struct{}{},
	}

	// okならつなげる
	// 最終的に一番長いやつの長さを返す
	var internal func(curCombi *combi, curSet []set)
	internal = func(curCombi *combi, curSet []set) {
		if len(curSet) == 0 {
			if longestLength < curCombi.length {
				longestLength = curCombi.length
			}
			return
		}

		// okかチェック
		check := true
		for r := range curSet[0].chars {
			for _, chs := range curCombi.charsArr {
				_, found := chs[r]
				if found {
					check = false
					break
				}
			}
			if !check {
				break
			}
		}

		// ngなら繋げない場合のみ
		internal(curCombi, curSet[1:])

		if check {
			addLength := curSet[0].length
			curCombi.length += addLength
			curCombi.charsArr = append(curCombi.charsArr, curSet[0].chars)
			internal(curCombi, curSet[1:])
			curCombi.length -= addLength
			curCombi.charsArr = curCombi.charsArr[:len(curCombi.charsArr)-1]
		}
	}
	internal(curCombi, setArr)

	return longestLength
}
