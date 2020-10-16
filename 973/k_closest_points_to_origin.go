// 973. K Closest Points to Origin
// https://leetcode.com/problems/k-closest-points-to-origin/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// May LeetCoding Challenge Day 30
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/538/week-5-may-29th-may-31st/3345/
func main() {
	// pp.Println("=========================================")
	// fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))
	// fmt.Println([][]int{{-2, 2}})
	// pp.Println("=========================================")
	// fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	// fmt.Println([][]int{{3, 3}, {-2, 4}})
	// pp.Println("=========================================")

	// min 5 つ
	heapify([]int{412, 3, 4389, 4, 837, 61, 28, 546}, 5)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func heapify(nums []int, k int) []int {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
	return []int(*h)
}

// 1 <= K <= points.length <= 10000
// -10000 < points[i][0] < 10000
// -10000 < points[i][1] < 10000

/*
 * 2. Priority Queue solution
 *   - Time Complexity: O(NlogK)
 *   - Space Complexity: O(K)
 */
func kClosest(points [][]int, k int) [][]int {
	return nil
}

/*
 * 1. Sort solution
 *   - Time Complexity: O(NlogN)
 *   - Space Complexity: O(logN)? for sorting
 */
func _kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:k]
}
