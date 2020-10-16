package helper

import (
	"fmt"
	"sort"
)

type IntSet map[int]struct{}

func NewIntSet() IntSet {
	return IntSet{}
}

func NewIntSetFromSlice(ints []int) IntSet {
	s := IntSet{}
	for _, n := range ints {
		s[n] = struct{}{}
	}
	return s
}

func (s IntSet) Add(n int) {
	s[n] = struct{}{}
}

func (s IntSet) AddAll(nums []int) {
	for _, n := range nums {
		s[n] = struct{}{}
	}
}

func (s IntSet) Contains(n int) bool {
	_, ok := s[n]
	return ok
}

func (s IntSet) First() int {
	for v := range s {
		return v
	}
	return 0
}

func (s IntSet) Remove(n int) {
	delete(s, n)
}

func (s IntSet) RemoveAll(nums []int) {
	for _, n := range nums {
		delete(s, n)
	}
}

func (s IntSet) ToSlice() []int {
	slice := make([]int, 0, len(s))
	for v := range s {
		slice = append(slice, v)
	}
	return slice
}

func (s IntSet) String() string {
	slice := s.ToSlice()
	sort.Ints(slice)
	return fmt.Sprintf("%v", slice)
}
