// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/
package main

import (
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

// Linked List について
//
// まず特性を知っておいたほうがいい
// 	LinkedListで実装されているデータ構造に対して計算量のかかる処理を頻繁にしてしまう可能性がある
//
// キューやスタックの実装
//
// 可変長配列がどのように実装されているかを知る
// 	メモリが足りなくなったら再アロケートするケースがほとんど
//
// 両方向Linked List
//
// Linked Listは所詮参照で繋がったリストでしかなく、データが参照で繋がるデータ構造はよくある
// 	グラフ（ツリー）など
// なので、参照の連なりを正しく操作できないと、LinkedList以外にも問題が発生する

type ListNode struct {
	Val   int
	Next  *ListNode
	debug struct {
		index int
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	var recurse func(node *ListNode)
	recurse = func(node *ListNode) {
		if node.Next == nil {
			head = node
			return
		}
		recurse(node.Next)
		node.Next.Next = node
	}
	recurse(head)
	tail.Next = nil
	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current, next := head, head.Next
	head.Next = nil
	for next != nil {
		next.Next, current, next = current, next, next.Next
	}
	return current
}
