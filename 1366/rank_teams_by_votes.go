// 1366. Rank Teams by Votes
// https://leetcode.com/problems/rank-teams-by-votes/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"}))
	pp.Println("ACB")
	pp.Println("=========================================")
	pp.Println(rankTeams([]string{"WXYZ", "XYZW"}))
	pp.Println("XWYZ")
	pp.Println("=========================================")
	pp.Println(rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}))
	pp.Println("ZMNAGUEDSJYLBOPHRQICWFXTVK")
	pp.Println("=========================================")
	pp.Println(rankTeams([]string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}))
	pp.Println("ABC")
	pp.Println("=========================================")
	pp.Println(rankTeams([]string{"M", "M", "M", "M"}))
	pp.Println("M")
	pp.Println("=========================================")
}

func rankTeams(votes []string) string {

}
