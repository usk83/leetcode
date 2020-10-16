// 5355. Frog Position After T Seconds
// https://leetcode.com/contest/weekly-contest-179/problems/frog-position-after-t-seconds/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 2, 4))
	pp.Println(0.16666666666666666)
	pp.Println("=========================================")
	pp.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 1, 7))
	pp.Println(0.3333333333333333)
	pp.Println("=========================================")
	pp.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 20, 6))
	pp.Println(0.16666666666666666)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// vertex: 1-n
// start: vertex 1
// ���L������
// ������ʴ_�ʤǤɤ�����
// �Ƅ��Ȥ��ʤ��Ȥ���ֹ�ޤ�
// t�����target�����_��
func frogPosition(n int, edges [][]int, t int, target int) float64 {
	// Graph�򘋺B����

	graph := make([][]int, n)
	for i := range graph {
		graph[i] = []int{}
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	current := map[int][2]int{
		1, 1,
	}

	// �F�ڤ�픵�픵�һ�E
	// �����ޤǤδ_�ʤ���
	for sec := 0; sec < t; sec++ {

	}

	// ��K�Ĥ�target�ˤĤ��Ƥ����Ĥ�
	if pos, ok := current[target]; ok {
		return pos[1]
	}

	return 0
}
