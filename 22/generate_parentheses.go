// 22. Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("=========================================")
	// pp.Println(generateParenthesis(3))
	pp.Println("=========================================")
	ret := generateParenthesis(14)
	pp.Println(len(ret))
	pp.Println("=========================================")
	// pp.Println(generateParenthesis(15))
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// "((()))",
// "(()())",
// "(())()",
// "()(())",
// "()()()"

// いや通るんかい！！！
// 入力、3, 10, 15 を試していて、15 で *** Limit Exceed が出たので、Time Complexityに問題があると思って改善
// もう一度15で試すも明らかに時間かかりすぎ
// 14までは少なくとも完了する
// よくみたらMemory Limit Exceeded
// 14のときの結果の次点で結果の長さがかなり長いので、もしや15はどうやっても計算量膨大になるかなと思う
// 試しにSubmit
// 100% 100% で通る

func generateParenthesis(n int) []string {
	result := make([][]string, n+1)
	var dfs func(parenthesis []byte, opened int)
	dfs = func(parenthesis []byte, opened int) {
		if len(parenthesis) == cap(parenthesis) {
			result[cap(parenthesis)>>1] = append(result[cap(parenthesis)>>1], string(parenthesis))
			return
		}
		if opened == 0 {
			if len(parenthesis) == 0 {
				dfs(append(parenthesis, '('), opened+1)
				return
			}
			for _, rus := range result[(cap(parenthesis)-len(parenthesis))>>1] {
				result[cap(parenthesis)>>1] = append(result[cap(parenthesis)>>1], string(parenthesis)+rus)
			}
			return
		}
		if cap(parenthesis)-len(parenthesis) == opened {
			dfs(append(parenthesis, ')'), opened-1)
			return
		}
		dfs(append(parenthesis, '('), opened+1)
		dfs(append(parenthesis, ')'), opened-1)
	}
	for i := 1; i <= n; i++ {
		dfs(make([]byte, 0, i<<1), 0)
	}
	return result[n]
}

/*
 * v1. Simply DFS
 */
// func generateParenthesis(n int) []string {
// 	result := []string{}
// 	var dfs func(parenthesis []byte, opened int)
// 	dfs = func(parenthesis []byte, opened int) {
// 		if len(parenthesis) == n<<1 {
// 			result = append(result, string(parenthesis))
// 			return
// 		}
// 		if opened == 0 {
// 			dfs(append(parenthesis, '('), opened+1)
// 			return
// 		}
// 		if n<<1-len(parenthesis) == opened {
// 			dfs(append(parenthesis, ')'), opened-1)
// 			return
// 		}
// 		dfs(append(parenthesis, '('), opened+1)
// 		dfs(append(parenthesis, ')'), opened-1)
// 	}
// 	dfs(make([]byte, 0, n<<1), 0)
// 	return result
// }
