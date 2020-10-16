// 399. Evaluate Division
// https://leetcode.com/problems/evaluate-division/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(calcEquation(
		[][]string{{"a", "b"}, {"b", "c"}},
		[]float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
	))
	fmt.Println([]float64{6.0, 0.5, -1.0, 1.0, -1.0})
	pp.Println("=========================================")
	fmt.Println(calcEquation(
		[][]string{{"a", "e"}, {"b", "e"}},
		[]float64{4.0, 3.0},
		[][]string{{"a", "b"}, {"e", "e"}, {"x", "x"}},
	))
	fmt.Println([]float64{1.33333, 1.0, -1.0})
	pp.Println("=========================================")
	fmt.Println(calcEquation(
		[][]string{{"a", "b"}, {"e", "f"}, {"b", "e"}},
		[]float64{3.4, 1.4, 2.3},
		[][]string{{"b", "a"}, {"a", "f"}, {"f", "f"}, {"e", "e"}, {"c", "c"}, {"a", "c"}, {"f", "e"}},
	))
	fmt.Println([]float64{0.29412, 10.948, 1.0, 1.0, -1.0, -1.0, 0.71429})
	pp.Println("=========================================")
}

/*
 * Using UnionFind
 */
func calcEquation(equations [][]string, answers []float64, queries [][]string) []float64 {
	uf := NewStringUnionFind()
	values := map[string]float64{}
	bases := map[string]int{}
	baseValues := []float64{}
	for i, equation := range equations {
		left, right := equation[0], equation[1]
		if uf.Connected(left, right) {
			continue
		}
		uf.Union(left, right)

		leftValue, leftOK := values[left]
		rightValue, rightOK := values[right]
		switch {
		case leftOK && rightOK:
			baseValues[bases[left]] = answers[i] * rightValue * baseValues[bases[right]] / leftValue
		case leftOK:
			values[right] = leftValue / answers[i]
			bases[right] = bases[left]
		case rightOK:
			values[left] = rightValue * answers[i]
			bases[left] = bases[right]
		default:
			values[left] = answers[i]
			bases[left] = len(baseValues)
			values[right] = 1.0
			bases[right] = bases[left]
			baseValues = append(baseValues, 1.0)
		}
	}

	result := make([]float64, len(queries))
	for i, query := range queries {
		if uf.Connected(query[0], query[1]) {
			result[i] = (values[query[0]] * baseValues[bases[query[0]]]) / (values[query[1]] * baseValues[bases[query[1]]])
		} else {
			result[i] = -1.0
		}
	}
	return result
}

type ufNodeInfo struct {
	parent string
	size   int
}

type StringUnionFind struct {
	nodes map[string]ufNodeInfo
	Count int
}

func NewStringUnionFind() StringUnionFind {
	return StringUnionFind{
		nodes: map[string]ufNodeInfo{},
	}
}

func NewStringUnionFindFromSlice(strs []string) StringUnionFind {
	uf := NewStringUnionFind()
	for _, str := range strs {
		uf.nodes[str] = ufNodeInfo{
			parent: str,
			size:   1,
		}
	}
	uf.Count = len(uf.nodes)
	return uf
}

func (uf *StringUnionFind) Add(str string) {
	if _, ok := uf.nodes[str]; ok {
		return
	}
	uf.nodes[str] = ufNodeInfo{
		parent: str,
		size:   1,
	}
	uf.Count++
}

func (uf *StringUnionFind) Find(str string) string {
	node := uf.nodes[str]
	root := node.parent
	for root != uf.nodes[root].parent {
		root = uf.nodes[root].parent
	}
	node.parent = root
	uf.nodes[str] = node
	return root
}

func (uf *StringUnionFind) Union(a, b string) {
	if _, ok := uf.nodes[a]; ok {
		a = uf.Find(a)
	} else {
		uf.Add(a)
	}
	if _, ok := uf.nodes[b]; ok {
		b = uf.Find(b)
	} else {
		uf.Add(b)
	}
	if a == b {
		return
	}
	aNode, bNode := uf.nodes[a], uf.nodes[b]
	if aNode.size >= bNode.size {
		bNode.parent = a
		aNode.size += bNode.size
	} else {
		aNode.parent = b
		bNode.size += aNode.size
	}
	uf.nodes[a], uf.nodes[b] = aNode, bNode
	uf.Count--
}

func (uf *StringUnionFind) Connected(a, b string) bool {
	if _, ok := uf.nodes[a]; !ok {
		return false
	}
	if _, ok := uf.nodes[b]; !ok {
		return false
	}
	return uf.Find(a) == uf.Find(b)
}

func (uf StringUnionFind) String() string {
	groups := map[string][]string{}
	for node := range uf.nodes {
		groups[uf.Find(node)] = append(groups[uf.Find(node)], node)
	}
	keys := make([]string, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	orderedGroups := make([][]string, len(keys))
	for i, key := range keys {
		sort.Strings(groups[key])
		orderedGroups[i] = groups[key]
	}
	return fmt.Sprint(orderedGroups)
}
