// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"container/heap"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numSubseq([]int{6, 5, 6, 7}, 9))
	pp.Println(4)
	pp.Println("=========================================")
	// pp.Println(numSubseq([]int{3, 3, 6, 8}, 10))
	// pp.Println(6)
	// pp.Println("=========================================")
	// pp.Println(numSubseq([]int{2, 3, 3, 4, 6, 7}, 12))
	// pp.Println(61)
	// pp.Println("=========================================")
	// pp.Println(numSubseq([]int{5, 2, 4, 1, 7, 6, 8}, 16))
	// pp.Println(127)
	// pp.Println("=========================================")
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
	heap.Fix(pq, item.index)
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

		pp.Println("=== DEBUG START ======================================")
		pp.Println(left)
		pp.Println(right)
		pp.Println(len(pp.Println(left)))
		pp.Println("=== DEBUG END ========================================")

		for maxPQ[0].appeared < left {
			_ = maxPQ.Pop()
		}
		for minPQ[0].appeared < left {
			_ = minPQ.Pop()
		}

		max, min := maxPQ[0].priority, minPQ[0].priority
		if max+min <= target {
			pp.Println("=== DEBUG START ======================================")
			pp.Println("¼ÓËã")
			pp.Println("left:", left)
			pp.Println("right:", right)
			pp.Println("max:", max)
			pp.Println("min:", min)
			pp.Println("=== DEBUG END ========================================")

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
