// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
	fmt.Println([][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}})
	pp.Println("=========================================")
	fmt.Println(suggestedProducts([]string{"havana"}, "havana"))
	fmt.Println([][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}})
	pp.Println("=========================================")
	fmt.Println(suggestedProducts([]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"))
	fmt.Println([][]string{{"baggage", "bags", "banner"}, {"baggage", "bags", "banner"}, {"baggage", "bags"}, {"bags"}})
	pp.Println("=========================================")
	fmt.Println(suggestedProducts([]string{"havana"}, "tatiana"))
	fmt.Println([][]string{{}, {}, {}, {}, {}, {}, {}})
	pp.Println("=========================================")
}

// 1 <= products.length <= 1000
// 1 <= products[i].length <= 2 * 10^4
// All characters of products[i] are lower-case English letters.
// 1 <= searchWord.length <= 1000
// All characters of searchWord are lower-case English letters.
func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	result := make([][]string, len(searchWord))
	for i := 0; i < len(searchWord); i++ {
		result[i] = make([]string, 0, len(searchWord))
	}

	activeStart, activeEnd := 0, len(products)

	for i := len(searchWord) - 1; i >= 0; i-- {
		word := searchWord[0 : i+1]
		for pointer := activeStart; pointer < activeEnd; pointer++ {
			if strings.HasPrefix(products[pointer], word) {
				activeStart = pointer
				break
			}
		}

		var count int
		for pointer := activeStart; pointer < activeEnd; pointer++ {
			if count == 3 || !strings.HasPrefix(products[pointer], word) {
				activeEnd = pointer
				break
			}
			count++
		}

		pp.Println(3 - activeEnd - activeStart)
		// for x := 1; x <= 3-activeEnd-activeStart; x++ {

		// }

		result[i] = products[activeStart:activeEnd]
	}

	return result
}
