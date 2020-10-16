// 5525. Throne Inheritance
// https://leetcode.com/contest/weekly-contest-208/problems/throne-inheritance/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")

	t := Constructor("king")         // order: king
	t.Birth("king", "andy")          // order: king > andy
	t.Birth("king", "bob")           // order: king > andy > bob
	t.Birth("king", "catherine")     // order: king > andy > bob > catherine
	t.Birth("andy", "matthew")       // order: king > andy > matthew > bob > catherine
	t.Birth("bob", "alex")           // order: king > andy > matthew > bob > alex > catherine
	t.Birth("bob", "asha")           // order: king > andy > matthew > bob > alex > asha > catherine
	order := t.GetInheritanceOrder() // return ["king", "andy", "matthew", "bob", "alex", "asha", "catherine"]
	fmt.Println(order)
	t.Death("bob")                  // order: king > andy > matthew > bob > alex > asha > catherine
	order = t.GetInheritanceOrder() // return ["king", "andy", "matthew", "alex", "asha", "catherine"]
	fmt.Println(order)

	pp.Println("=========================================")
}

// 1 <= kingName.length, parentName.length, childName.length, name.length <= 15
// kingName, parentName, childName, and name consist of lowercase English letters only.
// All arguments childName and kingName are distinct.
// All name arguments of death will be passed to either the constructor or as childName to birth first.
// For each call to birth(parentName, childName), it is guaranteed that parentName is alive.
// At most 105 calls will be made to birth and death.
// At most 10 calls will be made to getInheritanceOrder.

// 各人ごとにLinkedListで子供一覧をもつ
type Node struct {
	name          string
	alive         bool
	nextSibling   *Node
	oldestChild   *Node
	youngestChild *Node
}

type ThroneInheritance struct {
	kingName string
	// 各人ごとにmapでそのノードを特定できるようにする
	family map[string]*Node
}

func Constructor(kingName string) ThroneInheritance {
	king := &Node{
		name:          kingName,
		alive:         true,
		nextSibling:   nil,
		oldestChild:   nil,
		youngestChild: nil,
	}
	return ThroneInheritance{
		kingName: kingName,
		family: map[string]*Node{
			kingName: king,
		},
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	parent := this.family[parentName]
	child := &Node{
		name:          childName,
		alive:         true,
		nextSibling:   nil,
		oldestChild:   nil,
		youngestChild: nil,
	}
	this.family[childName] = child

	if parent.oldestChild == nil {
		parent.oldestChild = child
		parent.youngestChild = child
	} else {
		parent.youngestChild.nextSibling = child
		parent.youngestChild = child
	}
}

func (this *ThroneInheritance) Death(name string) {
	node := this.family[name]
	node.alive = false
}

func (this *ThroneInheritance) getChildren(name string) []string {
	// 対象の人から順にnextSiblingをたどる
	parent := this.family[name]
	if parent.oldestChild == nil {
		return nil
	}
	children := []string{}
	childNode := parent.oldestChild
	children = append(children, childNode.name)
	for childNode.nextSibling != nil {
		childNode = childNode.nextSibling
		children = append(children, childNode.name)
	}
	return children
}

func (this *ThroneInheritance) getInheritanceOrder(name string) []string {
	order := []string{}
	parent := this.family[name]
	// 子どもたちを追加
	children := this.getChildren(parent.name)
	for _, child := range children {
		childNode := this.family[child]
		if childNode.alive {
			order = append(order, child)
		}
		// 子供の子供を先頭から順に追加（再帰）
		order = append(order, this.getInheritanceOrder(child)...)
	}
	return order
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	// （常に生きていれば）
	// kingを追加
	// 子どもたちを追加
	// 先頭から順に子どもたちを追加
	// 	再帰
	order := []string{}
	king := this.family[this.kingName]
	if king.alive {
		order = append(order, king.name)
	}
	children := this.getChildren(king.name)
	for _, child := range children {
		childNode := this.family[child]
		if childNode.alive {
			order = append(order, child)
		}
		order = append(order, this.getInheritanceOrder(child)...)
	}
	return order
}
