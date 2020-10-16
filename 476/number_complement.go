// 476. Number Complement
// https://leetcode.com/problems/number-complement/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findComplement(5))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findComplement(1))
	pp.Println(0)
	pp.Println("=========================================")
}

// Go、定数のオーバーフローチェックがあったりするから最上位ビットをいじりづらくてかえって面倒

// 最後にビットマスクを適用する方法
// func findComplement(num int) int {
// 	unum, mask := uint32(num), uint32(math.MaxUint32)
// 	for bit := uint32(1 << 31); bit != 0; bit >>= 1 {
// 		if unum&bit != 0 {
// 			break
// 		}
// 		mask >>= 1
// 	}
// 	return int(^unum & mask)
// }

// leading zero bitを1にフリップしてから全体をフリップさせる方法
// 特定ビットをビットマスクで抽出する方法
// ※ 32bit環境でも動作する（と思われる）
func findComplement(num int) int {
	unum := uint32(num)
	for bit := uint32(1 << 31); bit != 0; bit >>= 1 {
		if unum&bit != 0 {
			break
		}
		unum |= bit
	}
	return int(^unum)
}

// ※ 32bit環境では動作しない（と思われる）
// func findComplement(num int) int {
// 	for bit := 1 << 31; bit != 0; bit >>= 1 {
// 		if num&bit != 0 {
// 			break
// 		}
// 		num |= bit
// 	}
// 	return int(^int32(num))
// }

// Leading zero bitsを1にしてから全体をフリップさせる方法
// 特定ビットをビットシフトで抽出する方法
// func findComplement(num int) int {
// 	n := int32(num)
// 	for c := uint8(31); c < 32; c-- {
// 		if n>>c&1 != 0 {
// 			break
// 		}
// 		n |= 1 << c
// 	}
// 	return int(^n)
// }
