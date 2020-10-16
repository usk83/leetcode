// 292. Nim Game
// https://leetcode.com/problems/nim-game/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(canWinNim(4))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(canWinNim(1))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canWinNim(5))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canWinNim(123))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canWinNim(1023))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(canWinNim(1024))
	pp.Println(false)
	pp.Println("=========================================")
}

func canWinNim(n int) bool {
	return n&3 != 0
}
