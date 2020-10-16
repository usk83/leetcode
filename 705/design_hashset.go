// 705. Design HashSet
// https://leetcode.com/problems/design-hashset/
// Posted this solution to https://leetcode.com/problems/design-hashset/discuss/482798
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	hashSet := Constructor()
	hashSet.Add(1)
	hashSet.Add(2)
	pp.Println(hashSet.Contains(1)) // returns true
	pp.Println(hashSet.Contains(3)) // returns false (not found)
	hashSet.Add(2)
	pp.Println(hashSet.Contains(2)) // returns true
	hashSet.Remove(2)
	pp.Println(hashSet.Contains(2)) // returns false (already removed)
	pp.Println("=========================================")
}

// All values will be in the range of [0, 1000000].
// The number of operations will be in the range of [1, 10000].
// Please do not use the built-in HashSet library.
/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

const ListCapacity = 10
const HashSize = 4

type MyHashSet struct {
	hashSetData
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (m *MyHashSet) Add(key int) {
	_ = m.add(key, 0)
}

func (m *MyHashSet) Remove(key int) {
	_ = m.remove(key, 0)
}

func (m *MyHashSet) Contains(key int) bool {
	return m.contains(key, 0)
}

type hashSetData struct {
	size int
	list []int
	next []MyHashSet
}

func hash(key int, count uint) int {
	return key >> (HashSize * count) & (1<<HashSize - 1)
}

func initList() []int {
	return make([]int, 0, ListCapacity)
}

func initNext() []MyHashSet {
	return make([]MyHashSet, 1<<HashSize)
}

func (h *hashSetData) add(key int, count uint) int {
	var added int
	if h.next != nil {
		added = h.next[hash(key, count)].add(key, count+1)
	} else if len(h.list) != ListCapacity {
		if h.list == nil {
			h.list = initList()
		}
		for _, v := range h.list {
			if v == key {
				return 0
			}
		}
		h.list = append(h.list, key)
		added = 1
	} else {
		h.next = initNext()
		for _, v := range h.list {
			_ = h.next[hash(v, count)].add(v, count+1)
		}
		added = h.next[hash(key, count)].add(key, count+1)
		h.list = nil
	}
	h.size += added
	return added
}

func (h *hashSetData) remove(key int, count uint) int {
	var removed int
	if h.next != nil {
		removed = h.next[hash(key, count)].remove(key, count+1)
	} else {
		for i, v := range h.list {
			if v == key {
				removed = 1
			} else if removed != 0 {
				h.list[i-1] = v
			}
		}
		h.list = h.list[:h.size-removed]
	}
	h.size -= removed
	if h.size == 0 {
		h.list = nil
		h.next = nil
	}
	return removed
}

func (h *hashSetData) contains(key int, count uint) bool {
	if h.next != nil {
		return h.next[hash(key, count)].contains(key, count+1)
	}
	for _, v := range h.list {
		if v == key {
			return true
		}
	}
	return false
}
