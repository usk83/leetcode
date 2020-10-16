// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxCandies(
		[]int{1, 0, 1, 0},
		[]int{7, 5, 4, 100},
		[][]int{{}, {}, {1}, {}},
		[][]int{{1, 2}, {3}, {}, {}},
		[]int{0}),
	)
	pp.Println(16)
	pp.Println("=========================================")
	pp.Println(maxCandies(
		[]int{1, 0, 0, 0, 0, 0},
		[]int{1, 1, 1, 1, 1, 1},
		[][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
		[][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
		[]int{0}),
	)
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(maxCandies(
		[]int{1, 1, 1},
		[]int{100, 1, 100},
		[][]int{{}, {0, 2}, {}},
		[][]int{{}, {}, {}},
		[]int{1}),
	)
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(maxCandies(
		[]int{1},
		[]int{100},
		[][]int{{}},
		[][]int{{}},
		[]int{}),
	)
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(maxCandies(
		[]int{1, 1, 1},
		[]int{2, 3, 2},
		[][]int{{}, {}, {}},
		[][]int{{}, {}, {}},
		[]int{2, 1, 0}),
	)
	pp.Println(7)
	pp.Println("=========================================")
}

// 1 <= status.length <= 1000
// status.length == candies.length == keys.length == containedBoxes.length == n
// status[i] is 0 or 1.
// 1 <= candies[i] <= 1000
// 0 <= keys[i].length <= status.length
// 0 <= keys[i][j] < status.length
// All values in keys[i] are unique.
// 0 <= containedBoxes[i].length <= status.length
// 0 <= containedBoxes[i][j] < status.length
// All values in containedBoxes[i] are unique.
// Each box is contained in one box at most.
// 0 <= initialBoxes.length <= status.length
// 0 <= initialBoxes[i] < status.length
func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	// status[i]: an integer which is 1 if box[i] is open and 0 if box[i] is closed.
	// candies[i]: an integer representing the number of candies in box[i].
	// keys[i]: an array contains the indices of the boxes you can open with the key in box[i].
	// containedBoxes[i]: an array contains the indices of the boxes found in box[i].

	// まずはこれから検証するboxに追加する
	currentBoxes := initialBoxes // これから検証するbox
	openedBoxes := []int{}       // 開けたbox
	waitingBoxes := []int{}      // まだ開けられていないbox

	for len(currentBoxes) != 0 {
		// fmt.Println("次に", currentBoxes, "を検証する")
		nextBoxes := []int{}
		for _, boxID := range currentBoxes {
			// 検証して開けたboxとまだ開けていないboxに分ける
			if status[boxID] == 1 {
				openedBoxes = append(openedBoxes, boxID)
				// 開けたboxから次のこれから検証するboxに追加していく
				nextBoxes = append(nextBoxes, containedBoxes[boxID]...)
				for _, key := range keys[boxID] {
					status[key] = 1
				}
			} else {
				waitingBoxes = append(waitingBoxes, boxID)
			}
		}
		// pp.Println("=== DEBUG START ======================================")
		// fmt.Println(openedBoxes)
		// fmt.Println(nextBoxes)
		// fmt.Println(waitingBoxes)
		// pp.Println("=== DEBUG END ========================================")
		if len(nextBoxes) != 0 {
			currentBoxes = nextBoxes
			continue
		}
		// これから検証するboxがなくなったとき、
		// まだ開けられていないboxについて鍵がみつかったか捜査する
		// それ終わってもまだ検証するboxがからのとき終了
		nextWaitingBoxes := []int{}
		for _, boxID := range waitingBoxes {
			if status[boxID] == 1 {
				nextBoxes = append(nextBoxes, boxID)
			} else {
				nextWaitingBoxes = append(nextWaitingBoxes, boxID)
			}
		}
		waitingBoxes = nextWaitingBoxes
		currentBoxes = nextBoxes
	}

	// これから検証するboxがなくなったら数える
	count := 0
	for _, box := range openedBoxes {
		count += candies[box]
	}
	return count
}
