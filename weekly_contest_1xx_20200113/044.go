// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(minTaps(5, []int{1, 1, 1, 1, 1, 1}))
	// pp.Println("=========================================")
	// pp.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	// // pp.Println(1)
	// pp.Println("=========================================")
	// pp.Println(minTaps(3, []int{0, 0, 0, 0}))
	// // pp.Println(-1)
	// pp.Println("=========================================")
	// pp.Println(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))
	// // pp.Println(3)
	// pp.Println("=========================================")
	// pp.Println(minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4}))
	// // pp.Println(2)
	// pp.Println("=========================================")
	// pp.Println(minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}))
	// // pp.Println(1)
	// pp.Println("=========================================")
}

// 1 <= n <= 10^4
// ranges.length == n + 1
// 0 <= ranges[i] <= 100
func minTaps(n int, ranges []int) int {

	// �Է֤Ǥ������_�Ǥ��ʤ���Ĥϥޥ���
	// �Է֤��I��򥫥Щ`���륿�åפΥꥹ�Ȥ����ɤ���

	// �ޥ��ȤΤ�Ĥ򿼑]�����Ȥ��ˤޤ����Щ`����Ƥ��ʤ��ꥹ�Ȥ򿼤���

	// ��λ���Q������

	// �ޥ��ȤΤ�Ĥ�����Ȥ��Ϥ�������Τ�Ĥ�����������
	// �����ФǤɤ��ȡ��Τ�һ�����ʤ�����

	type data struct {
		taps map[int]bool
	}

	garden := map[int]data{}
	for i := n - 1; i >= 0; i-- {
		garden[i] = data{
			taps: map[int]bool{},
		}
	}
	area := make([][2]int, len(ranges))

	for i, covers := range ranges {
		if covers == 0 {
			continue
		}

		start, end := i-covers, i+covers
		if start < 0 {
			start = 0
		}
		if end > n {
			end = n
		}

		area[i][0] = start
		area[i][1] = end
		for j := start; j < end; j++ {
			garden[j].taps[i] = true
		}
	}

	// fmt.Println(garden)

	// ��λ���Q������
	unnes := []int{}
	for i, cover := range area {
		if cover[0] == 0 && cover[1] == 0 {
			continue
		}
		// �Է֤Υ��Щ`�I����Է֤���range���ڤ���Ĥ������unnes
		starts := garden[cover[0]].taps
		ends := garden[cover[1]-1].taps

		for s := range starts {
			if ends[s] && s != i {
				unnes = append(unnes, i)
				break
			}
		}
	}

	// fmt.Println(unnes)

	var opened int

	// nes := []int{}
	nes := map[int]bool{}
	for _, g := range garden {
		for _, un := range unnes {
			delete(g.taps, un)
		}

		if l := len(g.taps); l == 0 {
			return -1
		} else if l == 1 {
			var tap int
			for t := range g.taps {
				tap = t
			}
			// nes = append(nes, tap)
			nes[tap] = true
		}
	}

	// fmt.Println(nes)

	fmt.Println(garden)

	for n := range nes {
		opened++
		pp.Println(area[n][0], area[n][1])
		for i := area[n][0]; i < area[n][1]; i++ {
			delete(garden, i)
		}
	}

	fmt.Println(garden)

	return opened
}
