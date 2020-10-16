// 1105. Filling Bookcase Shelves
// https://leetcode.com/problems/filling-bookcase-shelves/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))
	pp.Println(minHeightShelves([][]int{{11, 83}, {170, 4}, {93, 80}, {155, 163}, {134, 118}, {75, 14}, {122, 192}, {123, 154}, {187, 29}, {160, 64}, {170, 152}, {113, 179}, {60, 102}, {28, 187}, {59, 95}, {187, 97}, {49, 193}, {67, 126}, {75, 45}, {130, 160}, {4, 102}, {116, 171}, {43, 170}, {96, 188}, {54, 15}, {167, 183}, {58, 158}, {59, 55}, {148, 183}, {89, 95}, {90, 113}, {51, 49}, {91, 28}, {172, 103}, {173, 3}, {131, 78}, {11, 199}, {77, 200}, {58, 65}, {77, 30}, {157, 58}, {18, 194}, {101, 148}, {22, 197}, {76, 181}, {21, 176}, {50, 45}, {80, 174}, {116, 198}, {138, 9}, {58, 125}, {163, 102}, {133, 175}, {21, 39}, {141, 156}, {34, 185}, {14, 113}, {11, 34}, {35, 184}, {16, 132}, {78, 147}, {85, 170}, {32, 149}, {46, 94}, {196, 3}, {155, 90}, {9, 114}, {117, 119}, {17, 157}, {94, 178}, {53, 55}, {103, 142}, {70, 121}, {9, 141}, {16, 170}, {92, 137}, {157, 30}, {94, 82}, {144, 149}, {128, 160}, {8, 147}, {153, 198}, {12, 22}, {140, 68}, {64, 172}, {86, 63}, {66, 158}, {23, 15}, {120, 99}, {27, 165}, {79, 174}, {46, 19}, {60, 98}, {160, 172}, {128, 184}, {63, 172}, {135, 54}, {40, 4}, {102, 171}, {29, 125}, {81, 9}, {111, 197}, {16, 90}, {22, 150}, {168, 126}, {187, 61}, {47, 190}, {54, 110}, {106, 102}, {55, 47}, {117, 134}, {33, 107}, {2, 10}, {18, 62}, {109, 188}, {113, 37}, {59, 159}, {120, 175}, {17, 147}, {112, 195}, {177, 53}, {148, 173}, {29, 105}, {196, 32}, {123, 51}, {29, 19}, {161, 178}, {148, 2}, {70, 124}, {126, 9}, {105, 87}, {41, 121}, {147, 10}, {78, 167}, {91, 197}, {22, 98}, {73, 33}, {148, 194}, {166, 64}, {33, 138}, {139, 158}, {160, 19}, {140, 27}, {103, 109}, {88, 16}, {99, 181}, {2, 140}, {50, 188}, {200, 77}, {73, 84}, {159, 130}, {115, 199}, {152, 79}, {1, 172}, {124, 136}, {117, 138}, {158, 86}, {193, 150}, {56, 57}, {150, 133}, {52, 186}, {21, 145}, {127, 97}, {108, 110}, {174, 44}, {199, 169}, {139, 200}, {66, 48}, {52, 190}, {27, 86}, {142, 191}, {191, 79}, {126, 114}, {125, 100}, {176, 95}, {104, 79}, {146, 189}, {144, 78}, {52, 106}, {74, 74}, {163, 128}, {34, 181}, {20, 178}, {15, 107}, {105, 8}, {66, 142}, {39, 126}, {95, 59}, {164, 69}, {138, 18}, {110, 145}, {128, 200}, {149, 150}, {149, 93}, {145, 140}, {90, 170}, {81, 127}, {57, 151}, {167, 127}, {95, 89}}, 200))
	pp.Println(minHeightShelves([][]int{{98, 161}, {194, 25}, {34, 106}, {160, 10}, {108, 145}, {80, 122}, {136, 6}, {156, 126}, {6, 18}, {147, 108}, {36, 19}, {174, 123}, {6, 106}, {54, 78}, {3, 9}, {133, 114}, {31, 128}, {189, 139}, {31, 82}, {114, 169}, {73, 93}, {115, 145}, {58, 179}, {142, 42}, {96, 95}, {112, 50}, {199, 60}, {121, 24}, {54, 79}, {168, 81}, {59, 10}, {52, 61}, {128, 30}, {142, 135}, {144, 122}, {95, 56}, {192, 82}, {188, 121}, {107, 177}, {94, 106}, {155, 92}, {136, 58}, {4, 99}, {99, 130}, {92, 83}, {53, 73}, {7, 3}, {75, 12}, {168, 169}, {38, 46}, {77, 162}, {29, 25}, {82, 134}, {140, 76}, {164, 104}, {126, 105}, {105, 6}, {172, 13}, {105, 48}, {109, 200}, {188, 82}, {11, 164}, {150, 53}, {12, 26}, {104, 81}, {108, 86}, {63, 79}, {138, 145}, {139, 111}, {143, 166}, {119, 96}, {83, 137}, {130, 93}, {144, 47}, {196, 157}, {194, 24}, {158, 60}, {85, 33}, {133, 57}, {170, 81}, {24, 128}, {108, 159}, {133, 130}, {45, 138}, {72, 160}, {69, 184}, {34, 161}, {160, 117}, {73, 58}, {54, 97}, {31, 44}, {91, 101}, {37, 141}, {92, 72}, {42, 182}, {171, 86}, {181, 16}, {3, 76}, {128, 18}, {193, 176}, {142, 75}, {177, 18}, {1, 53}, {67, 37}, {3, 68}, {157, 81}, {186, 69}, {57, 36}, {55, 182}, {97, 53}, {15, 181}, {189, 92}, {188, 136}, {167, 41}, {55, 74}, {16, 116}, {85, 155}, {189, 27}, {36, 104}, {77, 21}, {198, 111}, {103, 145}, {66, 156}, {8, 82}, {159, 158}, {7, 114}, {93, 165}, {107, 196}, {141, 37}, {13, 29}, {120, 149}, {99, 55}, {154, 16}, {15, 126}, {9, 82}, {162, 70}, {176, 99}, {37, 16}, {41, 18}, {187, 39}, {40, 107}, {158, 198}, {17, 51}, {150, 10}, {2, 60}, {29, 26}, {6, 19}, {178, 167}, {182, 81}, {135, 133}, {10, 129}, {182, 138}, {170, 168}, {14, 84}, {105, 163}, {86, 69}, {150, 82}, {114, 33}, {152, 135}, {17, 89}, {149, 37}, {31, 145}, {112, 115}, {27, 41}, {19, 193}, {3, 172}, {140, 143}, {149, 109}, {165, 19}, {174, 30}, {99, 188}, {125, 150}, {141, 81}, {175, 29}, {1, 82}, {134, 183}, {100, 147}, {87, 121}, {200, 89}, {199, 123}, {109, 177}, {146, 199}, {149, 82}, {153, 163}, {79, 193}, {96, 130}, {66, 4}, {94, 120}, {148, 192}, {25, 146}, {101, 67}, {170, 159}, {103, 111}, {158, 35}, {42, 149}, {12, 180}, {36, 91}, {49, 128}, {123, 3}, {115, 181}, {160, 154}, {104, 137}, {97, 84}, {24, 144}, {35, 21}, {167, 36}, {92, 24}, {16, 74}, {173, 174}, {130, 148}, {84, 185}, {129, 90}, {34, 111}, {75, 7}, {113, 113}, {20, 144}, {75, 145}, {170, 137}, {160, 183}, {164, 4}, {56, 20}, {174, 47}, {191, 67}, {74, 12}, {16, 108}, {35, 167}, {180, 93}, {191, 90}, {61, 172}, {29, 74}}, 200))
}

// 1 <= books.length <= 1000
// 1 <= books[i][0] <= shelf_width <= 1000
// 1 <= books[i][1] <= 1000
func minHeightShelves(books [][]int, shelf_width int) int {
	var putBook func(books [][]int, restOfWidth, fullHeight, currentHeight int) int

	putNext := func(bs [][]int, r, f, c int) int {
		r = r - bs[0][0]
		if bs[0][1] > c {
			f += (bs[0][1] - c)
			c = bs[0][1]
		}
		return putBook(bs[1:], r, f, c)
	}

	putBelow := func(bs [][]int, r, f, c int) int {
		r = shelf_width - bs[0][0]
		c = bs[0][1]
		f += c
		return putBook(bs[1:], r, f, c)
	}

	putBook = func(books [][]int, restOfWidth, fullHeight, currentHeight int) int {
		// pp.Printf("残り %s 冊\n", len(books))

		if len(books) == 0 {
			// pp.Println("結果: %s", fullHeight)
			return fullHeight
		}

		if restOfWidth < books[0][0] {
			// pp.Println("下に置く")
			return putBelow(books, restOfWidth, fullHeight, currentHeight)
		}

		if currentHeight >= books[0][1] {
			// pp.Println("横に置く")
			return putNext(books, restOfWidth, fullHeight, currentHeight)
		}

		if len(books) == 1 {
			return putNext(books, restOfWidth, fullHeight, currentHeight)
		}

		if books[0][0]+books[1][0] > shelf_width {
			return putNext(books, restOfWidth, fullHeight, currentHeight)
		}

		// // 入れたときの差分
		// diffA := books[0][1] - currentHeight

		// // 入れなかったときの下の高さ
		// tempWidth := shelf_width
		// tempHeightA := 0
		// for i, leng := 0, len(books); i < leng; i++ {
		// 	if tempWidth -= books[i][0]; tempWidth >= 0 {
		// 		if tempHeightA < books[i][1] {
		// 			tempHeightA += books[i][1]
		// 		}
		// 		continue
		// 	}
		// 	break
		// }

		// // 入れたときの下の高さ
		// tempWidth = shelf_width
		// tempHeightB := 0
		// for i, leng := 0, len(books); i < leng; i++ {
		// 	if tempWidth -= books[i][0]; tempWidth >= 0 {
		// 		if tempHeightB < books[i][1] {
		// 			tempHeightB += books[i][1]
		// 		}
		// 		continue
		// 	}
		// 	break
		// }

		// // その差分
		// diffB := tempHeightA - tempHeightB

		// if diffA <= diffB {
		// 	return putNext(books, restOfWidth, fullHeight, currentHeight)
		// }

		// return putBelow(books, restOfWidth, fullHeight, currentHeight)

		// pp.Println("両方試す")

		// c := make(chan int)

		// var wg sync.WaitGroup
		// wg.Add(2)
		// var patternBelow, patternNext int
		// go func() {
		// 	// pp.Println("まず下に置く")
		// 	defer wg.Done()
		// 	patternBelow = putBelow(books, restOfWidth, fullHeight, currentHeight)
		// 	// c <- 1
		// }()
		// go func() {
		// 	// pp.Println("次の横に置く")
		// 	defer wg.Done()
		// 	patternNext = putNext(books, restOfWidth, fullHeight, currentHeight)
		// 	// c <- 1
		// }()

		// wg.Wait()

		// pp.Printf("下においたパターンは%s, 横においたパターンは%s\n", patternBelow, patternNext)
		// <-c
		// <-c

		patternBelow := putBelow(books, restOfWidth, fullHeight, currentHeight)
		patternNext := putNext(books, restOfWidth, fullHeight, currentHeight)

		if patternBelow < patternNext {
			return patternBelow
		}
		return patternNext
	}

	return putBook(books, 0, 0, 0)
}