// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"container/heap"

	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("=========================================")
	// pp.Println(numSubseq([]int{3, 5, 6, 7}, 9))
	// pp.Println(4)
	pp.Println("=========================================")
	pp.Println(numSubseq([]int{3, 3, 6, 8}, 10))
	pp.Println(6)
	pp.Println("=========================================")
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
	return pq[i].priority > pq[j].priority || (pq[i].priority == pq[j].priority && pq[i].appeared < pq[j].appeared)
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
	return pq[i].priority < pq[j].priority || (pq[i].priority == pq[j].priority && pq[i].appeared < pq[j].appeared)
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
	heap.Push(&minPQ, &Item{
		appeared: 0,
		priority: nums[0],
	})
	heap.Push(&maxPQ, &Item{
		appeared: 0,
		priority: nums[0],
	})

	var count int
	for left < len(nums) && right < len(nums) {
		if right-left < 0 {
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
			_ = heap.Pop(&maxPQ)
		}
		for minPQ[0].appeared < left {
			_ = heap.Pop(&minPQ)
		}

		max, min := maxPQ[0].priority, minPQ[0].priority

		pp.Println("=== DEBUG START ======================================")
		pp.Println("left:", left)
		pp.Println("right:", right)
		// pp.Println("maxPQ:", maxPQ)
		// pp.Println("minPQ:", minPQ)
		pp.Println("max:", max)
		pp.Println("min:", min)

		if max+min <= target {
			pp.Println("加算")

			if nums[right]<<1 <= target {
				count++
			}

			count = (count + (right - left)) % mod

			right++
			if right == len(nums) {
				break
			}
			heap.Push(&minPQ, &Item{
				appeared: right,
				priority: nums[right],
			})
			heap.Push(&maxPQ, &Item{
				appeared: right,
				priority: nums[right],
			})
		} else {
			left++
		}
		pp.Println("=== DEBUG END ========================================")
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
