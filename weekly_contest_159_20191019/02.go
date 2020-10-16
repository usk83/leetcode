package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	// fmt.Println(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
	// pp.Println([]string{"/a", "/c/d", "/c/f"})
	// fmt.Println(removeSubfolders([]string{"/a", "/a/b/c", "/a/b/d"}))
	// pp.Println([]string{"/a"})
	// fmt.Println(removeSubfolders([]string{"/a/b/c", "/a/b/ca", "/a/b/d"}))
	// pp.Println([]string{"/a/b/c", "/a/b/ca", "/a/b/d"})

	fmt.Println(removeSubfolders([]string{"/ah/al/am", "/ah/al"}))
	pp.Println([]string{"/ah/al"})

}

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 1 <= folder.length <= 4 * 10^4
// 2 <= folder[i].length <= 100
// folder[i] contains only lowercase letters and '/'
// folder[i] always starts with character '/'
// Each folder name is unique.
func removeSubfolders(folder []string) []string {
	sort.Sort(ByLen(folder))

	result := []string{}
	result = append(result, folder[0])
	flag := true
	for i := 1; i < len(folder); i++ {
		for _, res := range result {
			if flag && isSubFolder(folder[i], res) {
				flag = false
			}
		}
		if flag {
			result = append(result, folder[i])
		}
		flag = true
	}
	return result
}

func isSubFolder(check, sub string) bool {
	for check != "" && sub != "" {
		ret1 := strings.SplitN(check, "/", 2)
		var checkHead string
		if len(ret1) == 1 {
			checkHead = ret1[0]
			check = ""
		} else {
			checkHead = ret1[0]
			check = ret1[1]
		}
		ret2 := strings.SplitN(sub, "/", 2)
		var subHead string
		if len(ret2) == 1 {
			subHead = ret2[0]
			sub = ""
		} else {
			subHead = ret2[0]
			sub = ret2[1]
		}
		if checkHead != subHead {
			return false
		}
	}

	return true
}
