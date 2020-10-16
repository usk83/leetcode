// 168. Excel Sheet Column Title
// https://leetcode.com/problems/excel-sheet-column-title/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Printf("%s - A\n", convertToTitle(1))
	pp.Printf("%s - Z\n", convertToTitle(26))
	pp.Printf("%s - AA\n", convertToTitle(27))
	pp.Printf("%s - AZ\n", convertToTitle(52))
	pp.Printf("%s - BA\n", convertToTitle(53))
	pp.Printf("%s - BZ\n", convertToTitle(78))
	pp.Printf("%s - AAA\n", convertToTitle(703))
	pp.Println("=========================================")
	pp.Printf("%s - A\n", convertToTitle(1))
	pp.Printf("%s - AA\n", convertToTitle(27))
	pp.Printf("%s - AAA\n", convertToTitle(703))
	pp.Printf("%s - AAAA\n", convertToTitle(18279))
	pp.Printf("%s - AAAAA\n", convertToTitle(475255))
	pp.Printf("%s - BAAAA\n", convertToTitle(932231))
	pp.Printf("%s - ZAAAA\n", convertToTitle(11899655))
	pp.Printf("%s - AAAAAA\n", convertToTitle(12356631))
	pp.Println("=========================================")
}

// A-Z 1-26
// AA-AZ 27-52
// BA-BZ 53-78
//
// A				1					1
// AA				27				1 + 26
// AAA			703				1 + 26 + 26*26
// AAAA			18279			1 + 26 + 26*26 + 26*26*26
// AAAAA		475255		1 + 26 + 26*26 + 26*26*26 + 26*26*26*26
// BAAAA		932231		1 + 26 + 26*26 + 26*26*26 + 26*26*26*26*2
// ZAAAA		11899655	1 + 26 + 26*26 + 26*26*26 + 26*26*26*26*26
// AAAAAA		12356631	1 + 26 + 26*26 + 26*26*26 + 26*26*26*26 + 26*26*26*26*26
func convertToTitle(n int) string {
	chars := []byte{}

	for n > 0 {
		n--
		chars = append(chars, byte('A'+n%26))
		n /= 26
	}

	for h, t := 0, len(chars)-1; h < t; h, t = h+1, t-1 {
		chars[h], chars[t] = chars[t], chars[h]
	}

	return string(chars)
}
