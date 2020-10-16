// 406. Queue Reconstruction by Height
// https://leetcode.com/problems/queue-reconstruction-by-height/
package main

import (
	"container/list"
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 6
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/539/week-1-june-1st-june-7th/3352/
func main() {
	pp.Println("=========================================")
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	fmt.Println([][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}})
	pp.Println("=========================================")
}

//
//
//
//
//
//
//
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	fmt.Println(people)

	output := list.New()
	for i, p := range people {
		j := p[1]

		pp.Println("i:", i, "j:", j)

		e := output.PushBack(i)
		if j < i {
			mark := output.Back()
			for k := i; k > j; k-- {
				mark = mark.Prev()
			}
			output.MoveBefore(e, mark)
		}

		// fmt.Print("[ ")
		// for ee := output.Front(); ee != nil; ee = ee.Next() {
		// 	fmt.Printf("%d, ", ee.Value.(int))
		// }
		// fmt.Println("]")
	}

	ans := make([][]int, len(people))
	for i, e := 0, output.Front(); e != nil; i, e = i+1, e.Next() {
		ans[i] = people[e.Value.(int)]
	}
	return ans
}

//
//
//
//
//
//
//
//
//
//

// func reconstructQueue(p [][]int) [][]int {
// 	sort.Slice(p, func(i, j int) bool {
// 		if p[i][0] == p[j][0] {
// 			return p[i][1] > p[j][1]
// 		}
// 		return p[i][0] < p[j][0]
// 	})
//
// 	l := list.New()
// 	for i := len(p) - 1; i >= 0; i-- {
// 		insert(l, p[i], p[i][1])
// 	}
//
// 	e := l.Front()
// 	for i := range p {
// 		p[i] = e.Value.([]int)
// 		e = e.Next()
// 	}
// 	return p
// }
//
// func insert(l *list.List, p []int, idx int) {
// 	if idx == 0 {
// 		l.PushFront(p)
// 		return
// 	}
// 	if idx == l.Len() {
// 		l.PushBack(p)
// 		return
// 	}
//
// 	e := l.Front()
// 	for i := 0; i < idx; i++ {
// 		e = e.Next()
// 	}
// 	l.InsertBefore(p, e)
// }

// func reconstructQueue(people [][]int) [][]int {
// 	sort.Slice(people, func(i, j int) bool {
// 		if people[i][0] == people[j][0] {
// 			return people[i][1] > people[j][1]
// 		}
// 		return people[i][0] < people[j][0]
// 	})
// 	// pp.Println("=== DEBUG START ======================================")
// 	// fmt.Println(people)
// 	indexes := make([]int, len(people))
// 	for i := range indexes {
// 		indexes[i] = i
// 	}
// 	// fmt.Println(indexes)
// 	queued := make([][]int, len(people))
// 	for _, v := range people {
// 		queued[indexes[v[1]]] = v
// 		indexes = append(indexes[:v[1]], indexes[v[1]+1:]...)
// 		// fmt.Println(indexes)
// 	}
// 	// pp.Println("=== DEBUG END ========================================")
// 	return queued
// }

// func reconstructQueue(people [][]int) [][]int {
// 	sort.Slice(people, func(i, j int) bool {
// 		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
// 	})
// 	newPeople := make([][]int, len(people))
// 	for i := range people {
// 		var count int
// 		for j := range newPeople {
// 			if newPeople[j] != nil {
// 				continue
// 			}
// 			if count == people[i][1] {
// 				newPeople[j] = people[i]
// 				break
// 			}
// 			count++
// 		}
// 	}
// 	return newPeople
// }

// func reconstructQueue(people [][]int) [][]int {
// 	sort.Slice(people, func(i, j int) bool {
// 		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
// 	})
// 	newPeople := make([][]int, len(people))
// 	for _, p := range people {
// 		var count int
// 		for i := range newPeople {
// 			if newPeople[i] != nil {
// 				continue
// 			}
// 			if count == p[1] {
// 				newPeople[i] = []int{p[0], p[1]}
// 				break
// 			}
// 			count++
// 		}
// 	}
// 	return newPeople
// }

// func reconstructQueue(people [][]int) [][]int {
// 	sort.Slice(people, func(i, j int) bool {
// 		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
// 	})
//
// 	var queue [][]int
// 	for len(people) != 0 {
// 		var newPeople [][]int
// 		for _, p := range queue {
// 			for len(people) != 0 && people[0][1] == len(newPeople) {
// 				newPeople, people = append(newPeople, people[0]), people[1:]
// 			}
// 			newPeople = append(newPeople, p)
// 		}
// 		for len(people) != 0 && people[0][1] == len(newPeople) {
// 			newPeople, people = append(newPeople, people[0]), people[1:]
// 		}
// 		queue = newPeople
// 	}
// 	return queue
// }
