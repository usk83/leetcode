package main

import (
	"math/rand"
	"strconv"
)

func main() {
	num := "1"
	for i := 1; i < 500; i++ {
		num += strconv.Itoa(rand.Int() & 1)
	}
	print(num)
}
