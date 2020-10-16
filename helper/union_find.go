package helper

import (
	"fmt"
	"sort"
)

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
