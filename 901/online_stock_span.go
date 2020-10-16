// 901. Online Stock Span
// https://leetcode.com/problems/online-stock-span/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 19
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3334/
func main() {
	pp.Println("=========================================")
	stockSpanner := Constructor()
	pp.Println(stockSpanner.Next(100))
	pp.Println(stockSpanner.Next(80))
	pp.Println(stockSpanner.Next(60))
	pp.Println(stockSpanner.Next(70))
	pp.Println(stockSpanner.Next(60))
	pp.Println(stockSpanner.Next(75))
	pp.Println(stockSpanner.Next(85))
	pp.Println("=========================================")
	pp.Println(1)
	pp.Println(1)
	pp.Println(1)
	pp.Println(2)
	pp.Println(1)
	pp.Println(4)
	pp.Println(6)
	pp.Println("=========================================")
}

// 降順配列を保持して
//   二分探索
//   からの累積和

const MaxInt = int(^uint(0) >> 1)

type StockSpanner struct {
	list  []int // 降順リスト
	cusum []int // 累積和
}

func Constructor() StockSpanner {
	return StockSpanner{
		list:  []int{MaxInt}, // dummy
		cusum: []int{0},      // dummy
	}
}

func (s *StockSpanner) Next(price int) int {
	index := sort.Search(len(s.list), func(i int) bool {
		return price >= s.list[i]
	})

	s.list = append(s.list[:index], price)
	s.cusum = append(s.cusum[:index], s.cusum[len(s.cusum)-1]+1)
	return s.cusum[len(s.cusum)-1] - s.cusum[len(s.cusum)-2]
}
