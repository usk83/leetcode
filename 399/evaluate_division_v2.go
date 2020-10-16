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
 * Sort
 */
func calcEquation(pairs [][]string, answers []float64, queries [][]string) []float64 {
	equations := make([]struct {
		left   string
		right  string
		answer float64
	}, len(pairs))

	for i, pair := range pairs {
		if pair[0] <= pair[1] {
			equations[i].left = pair[0]
			equations[i].right = pair[1]
			equations[i].answer = answers[i]
		} else {
			equations[i].left = pair[1]
			equations[i].right = pair[0]
			equations[i].answer = 1.0 / answers[i]
		}
	}

	sort.Slice(equations, func(i, j int) bool {
		return equations[i].left < equations[j].left
	})

	type val struct {
		value float64
		group int
	}
	values := map[string]val{}
	var groupCount int
	for _, equation := range equations {
		leftValue, leftOK := values[equation.left]
		rightValue, rightOK := values[equation.right]
		switch {
		case leftOK && rightOK:
		case leftOK:
			rightValue.value = leftValue.value / equation.answer
			rightValue.group = leftValue.group
			values[equation.right] = rightValue
		case rightOK:
			leftValue.value = rightValue.value * equation.answer
			leftValue.group = rightValue.group
			values[equation.left] = leftValue
		default:
			leftValue.value = equation.answer
			leftValue.group = groupCount
			rightValue.value = 1
			rightValue.group = groupCount
			values[equation.left] = leftValue
			values[equation.right] = rightValue
			groupCount++
		}
	}

	result := make([]float64, len(queries))
	for i, query := range queries {
		leftValue, leftOK := values[query[0]]
		rightValue, rightOK := values[query[1]]
		if !leftOK || !rightOK || leftValue.group != rightValue.group {
			result[i] = -1.0
		} else {
			result[i] = leftValue.value / rightValue.value
		}
	}
	return result
}
