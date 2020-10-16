// 399. Evaluate Division
// https://leetcode.com/problems/evaluate-division/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("=========================================")
	// fmt.Println(calcEquation(
	// 	[][]string{{"a", "b"}, {"b", "c"}},
	// 	[]float64{2.0, 3.0},
	// 	[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
	// ))
	// fmt.Println([]float64{6.0, 0.5, -1.0, 1.0, -1.0})
	pp.Println("=========================================")
	fmt.Println(calcEquation(
		[][]string{{"a", "e"}, {"b", "e"}},
		[]float64{4.0, 3.0},
		[][]string{{"a", "b"}, {"e", "e"}, {"x", "x"}},
	))
	fmt.Println([]float64{1.33333, 1.0, -1.0})
	pp.Println("=========================================")
	// fmt.Println(calcEquation(
	// 	[][]string{{"a", "b"}, {"e", "f"}, {"b", "e"}},
	// 	[]float64{3.4, 1.4, 2.3},
	// 	[][]string{{"b", "a"}, {"a", "f"}, {"f", "f"}, {"e", "e"}, {"c", "c"}, {"a", "c"}, {"f", "e"}},
	// ))
	// fmt.Println([]float64{0.29412, 10.948, 1.0, 1.0, -1.0, -1.0, 0.71429})
	// pp.Println("=========================================")
	// fmt.Println(calcEquation(
	// 	[][]string{{"a", "b"}, {"b", "c"}},
	// 	[]float64{2.0, 3.0},
	// 	[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
	// ))
	// fmt.Println([]float64{6.0, 0.5, -1.0, 1.0, -1.0})
	// pp.Println("=========================================")
}

/*
 * Floyd–Warshall algorithm
 */
func calcEquation(equations [][]string, answers []float64, queries [][]string) []float64 {
	fw := map[string]map[string]float64{}
	for i, equation := range equations {
		for j, answer := range []float64{answers[i], 1.0 / answers[i]} {
			if _, ok := fw[equation[j]]; !ok {
				fw[equation[j]] = map[string]float64{
					equation[j]: 1.0,
				}
			}
			fw[equation[j]][equation[j^1]] = answer
		}
		// fmt.Println(equation)
		for j, left := range equation {
			right := equation[j^1]
			// pp.Println("current:", left)
			for target := range fw[left] {
				for dist := range fw[right] {
					pp.Println("left:", left, "right:", right)
					if _, ok := fw[target][dist]; ok {
						continue
					}
					pp.Println("target:", target, "left:", left, "right:", right, "dist:", dist)
					pp.Println("fw[target][left]: ", fw[target][left])
					pp.Println("fw[left][right]: ", fw[left][right])
					pp.Println("fw[right][dist]: ", fw[right][dist])
					fw[target][dist] = fw[target][left] * fw[left][right] * fw[right][dist]
				}
				if _, ok := fw[target][right]; ok {
					continue
				}
				// fmt.Println("from:", target, "to:", right, "left:", left)
				fw[target][right] = fw[target][left] * fw[left][right]
				// fmt.Println("from:", right, "to:", target, "left:", left)
				fw[right][target] = fw[right][left] * fw[left][target]

			}
		}
	}
	result := make([]float64, len(queries))
	for i, query := range queries {
		if from, ok := fw[query[0]]; !ok {
			result[i] = -1.0
		} else if to, ok := from[query[1]]; !ok {
			result[i] = -1.0
		} else {
			result[i] = to
		}
	}
	return result
}
