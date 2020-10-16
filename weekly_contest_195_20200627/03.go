// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"container/heap"

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

type Item struct {
	appeared int
	priority int // The priority of the item in the queue.
	index    int // The index of the item in the heap.
}

type MaxPriorityQueue []*Item

func (pq MaxPriorityQueue) Len() int { return len(pq) }

func (pq MaxPriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq MaxPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MaxPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MaxPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *MaxPriorityQueue) update(item *Item, priority int, appeared int) {
	item.priority = priority
	item.appeared = appeared
	heap.Fix(pq, item.index)
}

type MinPriorityQueue []*Item

func (pq MinPriorityQueue) Len() int { return len(pq) }

func (pq MinPriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq MinPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MinPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *MinPriorityQueue) update(item *Item, priority int, appeared int) {
	item.priority = priority
	item.appeared = appeared
	heap.Fix(pq, item.index)
}

const mod = 1e9 + 7

func numSubseq(nums []int, target int) int {
	var left, right int
	minPQ, maxPQ := MinPriorityQueue{}, MaxPriorityQueue{}
			item := &Item{
				appeared: 0,
				priority: nums[0],
			}
			minPQ.Push(item)
			maxPQ.Push(item)

	var count int
	for left < len(nums)-1 && right < len(nums) {
		if right-left < 1 {
			right++
			if right == len(nums) {
				break
			}
					item := &Item{
				appeared: right,
				priority: nums[right],
			}
			minPQ.Push(item)
			maxPQ.Push(item)
			continue
		}


		for maxPQ[0].appeared < left {
			_ = maxPQ.Pop()
		}
				for minPQ[0].appeared < left {
			_ = minPQ.Pop()
		}

		max, min := maxPQ[0].(*Item).priority, minPQ[0].(*Item).priority
		if max + min <= target {
			count = (count + (right - left)) % mod
			right++
			if right == len(nums) {
				break
			}
			item := &Item{
				appeared: right,
				priority: nums[right],
			}
			minPQ.Push(item)
			maxPQ.Push(item)
		} else {
			left++
		}

		max, min¤òÓ‹Ëã
		´óÕÉ·ò¤À¤Ã¤¿¤é¼ÓËã
			ÓÒ¤ËÉì¤Ð¤»¤ë¤Ê¤éÉì¤Ð¤¹
			Éì¤Ð¤»¤Ê¤¤¤Ê¤é¤ª¤ï¤ê
		´óÕÉ·ò¤¸¤ã¤Ê¤«¤Ã¤¿¤é
			×ó¤òÏ÷¤ë



	}
	return count
}

// var mins, maxes []int

// func numSubseq(nums []int, target int) int {
// 	mins, maxes := make([]int, len(nums)<<1-1), make([]int, len(nums)<<1-1)

// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
