// 5345. Rank Teams by Votes
// https://leetcode.com/contest/weekly-contest-178/problems/rank-teams-by-votes/
package main

import (
	"sort"

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
	pp.Println(rankTeams([]string{"YEQG", "QYGE", "QYGE", "GEQY", "QYEG", "EYGQ", "GYQE", "EQGY", "QYGE", "GQYE", "YGEQ", "GQYE", "EGQY", "YEGQ", "GEQY", "YGEQ", "EYQG", "EGYQ", "GYQE", "GEYQ", "EGYQ", "YEGQ", "YEQG", "QGEY", "QYEG", "YQGE", "GQYE", "EQYG", "EYGQ", "EQGY", "QGYE", "QYGE", "GQEY", "YGEQ", "QEGY", "YQEG", "GYQE", "QYGE", "QYGE", "EYGQ", "EGYQ", "QEGY", "YEGQ", "GQEY", "GYEQ", "EQYG", "QGYE", "QEGY", "YQGE", "QEYG", "GQEY", "YGEQ", "GQYE", "YEQG", "QGEY", "GEYQ", "GYQE", "EQYG", "EQYG", "GQEY", "GEQY", "YGEQ", "GYQE", "GQEY", "GQEY", "YEQG", "EYQG", "YGEQ", "YGQE", "YGEQ", "EGQY", "GQYE", "EGYQ", "QEYG", "GYQE", "GYQE", "YQGE", "YEQG", "GQEY", "YGQE", "GQYE", "GYEQ", "QGYE", "EGYQ", "YEGQ", "EQYG", "QEGY", "YEGQ", "YEGQ", "EYQG", "YEQG", "EQYG", "EQYG", "YEQG", "GYEQ", "GQYE", "EQYG", "YQGE", "QGYE", "EYQG", "EQYG", "GQYE", "GYEQ", "EYQG", "GQYE", "EYQG", "EYGQ", "QGEY", "GEYQ", "GYQE", "QYEG", "QEGY", "YEQG", "YEQG", "YEGQ", "GQYE", "GYEQ", "EGQY", "EQGY", "QEGY", "GEYQ", "YQGE", "YQGE", "YEGQ", "GYQE", "GQYE", "QGEY", "EQGY", "QEGY", "GYEQ", "EGQY", "YEQG", "QGYE", "QEYG", "QYGE", "EQYG", "YGEQ", "YGQE", "GYEQ", "YGEQ", "QYEG", "QGYE", "EQYG", "YGQE", "QEYG", "EYGQ", "QGYE", "YQEG", "GQEY", "GQEY", "GQYE", "GYQE", "EGQY", "YGEQ", "YQGE", "GEQY", "QYGE", "YEQG", "QEYG", "EQYG", "GQYE", "EQGY", "GQYE", "QGYE", "GYQE", "QGEY", "QGYE", "EQGY", "QYGE", "GEYQ", "GEYQ", "YQGE", "EGYQ", "GYEQ", "QGEY", "GQYE", "GYQE", "QEGY", "QGYE", "GEQY", "GQYE", "EGYQ", "GYQE", "GQYE", "EQGY", "GQYE", "EQYG", "QYGE", "QGEY", "GYEQ", "EGYQ", "YQEG", "YGEQ", "EQYG", "QGEY", "QYEG", "GQYE", "EGYQ", "GQEY", "EGYQ", "YEQG", "YQEG", "EQYG", "QGYE", "GQEY", "GQEY", "GEQY", "QYEG", "EGQY", "GEQY", "YEGQ", "GQYE", "GYQE", "YEQG", "YGEQ", "YQGE", "QEGY", "YGEQ", "YGEQ", "QEGY", "EQYG", "EGQY", "EYGQ", "YQEG", "YEGQ", "EGQY", "QEYG", "GYQE", "YEQG", "QYEG", "EYGQ", "GEYQ", "YGEQ", "QGEY", "EGYQ", "GEQY", "EGYQ", "GQEY", "YQEG", "YEGQ", "YEGQ", "QEYG", "YEGQ", "EQGY", "EQGY", "QEYG", "EYQG", "EGYQ", "YQGE", "YQEG", "GEQY", "YGEQ", "EYQG", "QEYG", "YEQG", "YQEG", "YGQE", "YEQG", "YGQE", "GYEQ", "QGYE", "GYEQ", "EGQY", "YEQG", "YEQG", "GQEY", "QEGY", "EGQY", "YQEG", "EYGQ", "GEQY", "GQEY", "QGYE", "YQEG", "QEYG", "YEGQ", "EGYQ", "QGYE", "EQGY", "GEQY", "QEYG", "QGYE", "EGYQ", "YEQG", "QYGE", "EGYQ", "EYQG", "QEYG", "YEGQ", "GEQY", "QYGE", "GYEQ", "YGQE", "QGEY", "YEGQ", "QGEY", "QEYG", "YEGQ", "GYEQ", "GQEY", "QGYE", "YEGQ", "EGYQ", "GEQY", "YQGE", "QGEY", "GYEQ", "GQEY", "EGQY", "YEGQ", "QEYG", "EGQY", "GQEY", "YEGQ", "YGQE", "YQEG", "QYEG", "YQEG", "QGEY", "QGYE", "GQEY", "GEQY", "QYGE", "QEYG", "EGYQ", "EQGY", "YEGQ", "EGYQ", "QGEY", "YEGQ", "QGYE", "GYEQ", "QEYG", "YQEG", "GYEQ", "YQGE", "GEQY", "YGQE", "GQYE", "QEYG", "GEYQ", "GYEQ", "YGEQ", "YEQG", "YEQG", "YGEQ", "YEQG", "GYQE", "YGQE", "YQEG", "EGYQ", "GYQE", "GQYE", "EQGY", "YQEG", "EQYG", "YEQG", "EQYG", "EGQY", "YEGQ", "EYQG", "QEGY", "GEQY", "YGQE", "GQEY", "GEYQ", "EQGY", "EYQG", "GYQE", "YEQG", "YEQG", "QEYG", "EGYQ", "GYQE", "EGQY", "YGEQ", "YQEG", "EGYQ", "EYQG", "QGYE", "QGYE", "YEQG", "YEGQ", "YEQG", "GEYQ", "GEQY", "EGQY", "GYEQ", "QEGY", "EGQY", "YGQE", "GEQY", "EQYG", "YQEG", "YEGQ", "EGYQ", "QGEY", "GYQE", "EGQY", "QGYE", "GEQY", "GYEQ", "YEQG", "YQEG", "QYEG", "GEYQ", "YGEQ", "GYQE", "YEGQ", "QEGY", "QGYE", "YQGE", "EGQY", "YQEG", "QYGE", "GYEQ", "EGQY", "EYGQ", "QYEG", "QEGY", "EGQY", "EGYQ", "GQYE", "GQEY", "YQGE", "YGEQ", "EGYQ", "EQGY", "EGQY", "GQYE", "EQGY", "QYGE", "GEYQ", "EQYG", "QEGY", "EGYQ", "EQYG", "GQYE", "QEGY", "QGEY", "QYEG", "QYEG", "GEYQ", "QGYE", "GEQY", "GYEQ", "EYGQ", "QEYG", "EGQY", "EQYG", "EYGQ", "QEGY", "QYGE", "GEQY", "YEQG", "EYGQ", "YEGQ", "YGQE", "YGQE", "EGQY", "GEQY", "GEQY", "EYQG", "YQEG", "QEGY", "GEYQ", "QYEG", "EQGY", "QGYE", "GQEY", "EGQY", "YGEQ", "EYQG", "YQEG", "QEYG", "EGYQ", "YEGQ", "YEQG", "QGYE", "GQYE", "YQGE", "GYQE", "QGYE", "QYGE", "YQEG", "EGQY", "QEYG", "QEGY", "GYEQ", "QEGY", "EYQG", "QEYG", "GEYQ", "GEYQ", "QGYE", "QYEG", "YEQG", "GQYE", "GQEY", "YEGQ", "QEGY", "GQEY", "EQGY", "QGEY", "YGQE", "GQYE", "QEGY", "YGEQ", "QGYE", "QEGY", "YGQE", "EGYQ", "QGYE", "EYGQ", "GQEY", "EQYG", "EGYQ", "EGYQ", "EYGQ", "EYGQ", "GYQE", "QGEY", "YEGQ", "GYQE", "YGQE", "QEYG", "YQGE", "GEYQ", "GQEY", "YQEG", "YQGE", "EYQG", "YGEQ", "GYQE", "YQEG", "GQYE", "QEYG", "EQYG", "EQGY", "GQEY", "EYQG", "YEGQ", "YEGQ", "EQYG", "EYQG", "YQEG", "QYGE", "YEQG", "QYGE", "QGYE", "GYQE", "EQYG", "GEQY", "YQGE", "GYQE", "EYGQ", "EQYG", "GQEY", "GYQE", "QEYG", "YQEG", "EYQG", "QGYE", "YQEG", "GEYQ", "GYEQ", "QEYG", "EQGY", "GYEQ", "GQEY", "EGQY", "YQGE", "GEQY", "GYQE", "QGYE", "QGYE", "GQYE", "EYQG", "QGYE", "YEQG", "EGQY", "GQEY", "QEYG", "YGQE", "EYQG", "YGQE", "YGQE", "GYEQ", "YEQG", "QGYE", "GYQE", "QGYE", "QYEG", "GYEQ", "QYGE", "QEGY", "EYQG", "YGQE", "GYEQ", "GEYQ", "GEQY", "EGQY", "YGQE", "YGEQ", "EQGY", "QYGE", "EQYG", "YEQG", "QYGE", "EQYG", "YQGE", "EQGY", "GEQY", "YGEQ", "EQGY", "YEGQ", "YEGQ", "EGYQ", "GQEY", "QGEY", "QYEG", "QGEY", "EYQG", "EQYG", "GYEQ", "YGQE", "GYQE", "GEYQ", "QGEY", "EYGQ", "QEYG", "YQEG", "EYQG", "QEYG", "QYEG", "EQGY", "QEGY", "YEQG", "QEYG", "GQEY", "EYGQ", "GQEY", "EYGQ", "YEGQ", "QEGY", "GQYE", "GEQY", "YQEG", "EQYG", "EQGY", "EQGY", "QYEG", "EGYQ", "YQEG", "GYQE", "GQEY", "QEYG", "YGEQ", "GYQE", "GYQE", "EYQG", "EYGQ", "EYQG", "YQGE", "YQGE", "EGYQ", "QYEG", "GYQE", "EQYG", "QYGE", "GYQE", "YGEQ", "YQGE", "QEYG", "YEGQ", "QGEY", "QEYG", "YQEG", "GQEY", "GYQE", "YQGE", "QYEG", "QGYE", "QEGY", "YEQG", "YQEG", "GEYQ", "EGQY", "GQYE", "YGQE", "QGEY", "EYQG", "EGQY", "YQGE", "QEYG", "EQGY", "YEQG", "QEGY", "QGYE", "QEYG", "EGQY", "YEGQ", "EGQY", "GQEY", "EGQY", "GYQE", "EYQG", "GEQY", "YGQE", "YGQE", "YGQE", "GEYQ", "QGEY", "YEGQ", "YEQG", "EYQG", "EGQY", "YEGQ", "QEGY", "GEYQ", "EQYG", "QGYE", "YGEQ", "EQGY", "QEYG", "EGQY", "EYGQ", "GYQE", "EGQY", "YQEG", "QYEG", "EQGY", "GQEY", "QGYE", "GQYE", "YEGQ", "GQEY", "QEGY", "GEYQ", "YEGQ", "YGEQ", "YGEQ", "YQEG", "YGEQ", "EYQG", "YQGE", "QGYE", "EGYQ", "EGYQ", "QEYG", "YEGQ", "GEYQ", "QEYG", "YGEQ", "GYQE", "YGQE", "GYEQ", "GQEY", "QGYE", "QYEG", "YEQG", "YEGQ", "EGYQ", "YEQG", "GQEY", "QGYE", "EYGQ", "EYQG", "EYGQ", "YQGE", "EYQG", "GQEY", "GYQE", "YEGQ", "QYGE", "YGEQ", "GYQE", "GQYE", "EGYQ", "YGEQ", "EQGY", "YQEG", "QYGE", "EQGY", "YGQE", "EGQY", "QEGY", "QYGE", "GEQY", "GYQE", "YEQG", "EYGQ", "EQYG", "EYGQ", "GEQY", "YGEQ", "YGQE", "YGEQ", "GEYQ", "EYQG", "YEGQ", "EQYG", "EQGY", "GEYQ", "YEGQ", "GEQY", "QYGE", "EGYQ", "QEGY", "QEYG", "QGYE", "YEQG", "QGEY", "QGYE", "GYQE", "GYQE", "EQYG", "QGEY", "QEYG", "EYGQ", "YQGE", "QYEG", "QYGE", "GEQY", "YQEG", "GYQE", "GYQE", "YQGE", "GYEQ", "EQGY", "EYQG", "EQYG", "GQYE", "EYQG", "YQGE", "QEGY", "QEYG", "EQYG", "YQGE", "EQYG", "EQYG", "GQEY", "GQYE", "EYQG", "QYEG", "GQYE", "EGQY", "EYGQ", "EGQY", "EGYQ", "EYQG", "GQYE", "EQYG", "QGYE", "GEYQ", "EQYG", "YGQE", "YQGE", "EYGQ", "QEYG", "EQYG", "YEGQ", "EYGQ", "GQEY", "EQGY", "GEYQ", "YGEQ", "EQGY", "EQYG", "YEGQ", "GEQY", "EGQY", "EQYG", "EGYQ", "GQYE", "GEYQ", "YQGE", "GQEY", "GQYE", "EQYG", "GEQY", "QYEG", "EGQY", "EYGQ", "EQGY", "GEQY", "YQGE", "GYQE", "EQYG", "EQGY", "YGEQ", "EYQG", "YGEQ", "YGEQ", "GYQE", "QYEG", "GQEY", "EQYG", "YGQE", "QEYG", "EYGQ", "EGYQ", "EQGY", "EYGQ", "EGYQ", "EYQG", "QYEG", "GYEQ", "GQEY", "EQGY", "YQGE", "EYQG", "YQEG", "GQEY", "GEQY", "EYQG", "EQYG", "QYGE", "GYEQ", "QEYG", "EQGY", "QYGE", "GEYQ", "GEYQ", "EYGQ", "QYEG", "QEGY", "GYQE", "YQEG", "GYQE", "QYGE", "YQGE", "EGYQ", "GQEY", "GQEY", "YEQG", "YEGQ", "QGYE", "GYQE", "YEQG", "GYQE", "GQEY", "YQEG", "GYQE", "EQYG", "EGQY", "YGEQ", "GYQE", "EQYG", "GEQY", "GQEY", "EYGQ", "YQEG", "QGYE", "GQYE", "YQEG", "EYGQ", "EQGY", "QYEG", "EYGQ", "YQGE", "YGQE", "EYQG", "YGEQ", "EGQY", "GEYQ", "EGQY", "QYEG", "GEQY", "QEYG", "QYEG", "YQGE", "QEGY", "YQGE", "GQEY", "YEGQ", "YGQE", "QGEY", "YGQE", "GQEY", "EYQG", "GQEY", "YQEG", "YQEG", "EQYG", "EYQG", "EGYQ", "EGQY", "GQEY", "GQYE", "GQYE", "QEGY", "YGQE", "GEQY", "GQEY"}))
	pp.Println("EGYQ")
	pp.Println("=========================================")
}

func rankTeams(votes []string) string {
	m := len(votes[0])

	ranks := [26][]int{}
	for i := range ranks {
		ranks[i] = make([]int, m)
	}

	for _, vote := range votes {
		for i, t := range vote {
			ranks[t-'A'][i]++
		}
	}

	rankString := []string{}
	for ascii, rank := range ranks {
		ok := false
		chars := []byte{}
		for i := range votes[0] {
			if rank[i] != 0 {
				ok = true
			}
			for part := 0; part < 4; part++ {
				if rank[i] >= 255 {
					chars = append(chars, byte(255))
				} else if rank[i] <= 0 {
					chars = append(chars, byte(0))
				} else {
					chars = append(chars, byte(rank[i]))
				}
				rank[i] -= 255
			}
		}
		if ok {
			rankString = append(rankString, string(chars)+string('Z'-ascii)+string(ascii+'A'))
		}
	}

	sort.Strings(rankString)

	ans := ""
	for i := len(rankString) - 1; i >= 0; i-- {
		ans += string(rankString[i][m*4+1])
	}

	return ans
}
