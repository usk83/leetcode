// 202. Happy Number
// https://leetcode.com/problems/happy-number/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(isHappy(19))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isHappy(279))
	pp.Println(false)
	pp.Println("=========================================")

}

/*
 * 2020-04-02
 * https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3284/
 */

/*
 * Floyd's Tortoise and Hare(Floyd's cycle-finding algorithm)
 * - Runtime: 0 ms, faster than 100.00%
 * - Memory Usage: 2 MB, less than 100.00%
 */
func sumOfDigitsSquare(n int) int {
	var sum, mod int
	for n != 0 {
		n, mod = n/10, n%10
		sum += mod * mod
	}
	return sum
}

func isHappy(n int) bool {
	tortoise, hare := n, n
	for {
		tortoise, hare = sumOfDigitsSquare(tortoise), sumOfDigitsSquare(sumOfDigitsSquare(hare))
		if tortoise == hare {
			break
		}
	}
	return tortoise == 1
}

/*
 * Loop 7 times
 * One of the longest paths is `9999999999922 -> 899 -> 226 -> 44 -> 32 -> 13 -> 10 -> 1`
 * - Runtime: 0 ms, faster than 100.00%
 * - Memory Usage: 2 MB, less than 100.00%
 */
// func isHappy(n int) bool {
// 	for i := 0; i < 7; i++ {
// 		var sum, mod int
// 		for n != 0 {
// 			n, mod = n/10, n%10
// 			sum += mod * mod
// 		}
// 		n = sum
// 	}
// 	return n == 1
// }

/*
 * Using Set and pre-calculating powers
 * - Runtime: 0 ms, faster than 100.00%
 * - Memory Usage: 2.1 MB, less than 100.00%
 */
// func isHappy(n int) bool {
// 	powers := [10]int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
// 	appeared := map[int]struct{}{}
// 	for n != 1 {
// 		appeared[n] = struct{}{}
// 		var next, mod int
// 		for n != 0 {
// 			n, mod = n/10, n%10
// 			next += powers[mod]
// 		}
// 		if _, ok := appeared[next]; ok {
// 			return false
// 		}
// 		n = next
// 	}
// 	return true
// }

/*
 * Pre-calculating
 * - Runtime: 0 ms, faster than 100.00%
 * - Memory Usage: 2 MB, less than 100.00%
 */
// var happys = map[int]struct{}{
// 	1:    struct{}{},
// 	7:    struct{}{},
// 	10:   struct{}{},
// 	13:   struct{}{},
// 	19:   struct{}{},
// 	23:   struct{}{},
// 	28:   struct{}{},
// 	31:   struct{}{},
// 	32:   struct{}{},
// 	44:   struct{}{},
// 	49:   struct{}{},
// 	68:   struct{}{},
// 	70:   struct{}{},
// 	79:   struct{}{},
// 	82:   struct{}{},
// 	86:   struct{}{},
// 	91:   struct{}{},
// 	94:   struct{}{},
// 	97:   struct{}{},
// 	100:  struct{}{},
// 	103:  struct{}{},
// 	109:  struct{}{},
// 	129:  struct{}{},
// 	130:  struct{}{},
// 	133:  struct{}{},
// 	139:  struct{}{},
// 	167:  struct{}{},
// 	176:  struct{}{},
// 	188:  struct{}{},
// 	190:  struct{}{},
// 	192:  struct{}{},
// 	193:  struct{}{},
// 	203:  struct{}{},
// 	208:  struct{}{},
// 	219:  struct{}{},
// 	226:  struct{}{},
// 	230:  struct{}{},
// 	236:  struct{}{},
// 	239:  struct{}{},
// 	262:  struct{}{},
// 	263:  struct{}{},
// 	280:  struct{}{},
// 	291:  struct{}{},
// 	293:  struct{}{},
// 	301:  struct{}{},
// 	302:  struct{}{},
// 	310:  struct{}{},
// 	313:  struct{}{},
// 	319:  struct{}{},
// 	320:  struct{}{},
// 	326:  struct{}{},
// 	329:  struct{}{},
// 	331:  struct{}{},
// 	338:  struct{}{},
// 	356:  struct{}{},
// 	362:  struct{}{},
// 	365:  struct{}{},
// 	367:  struct{}{},
// 	368:  struct{}{},
// 	376:  struct{}{},
// 	379:  struct{}{},
// 	383:  struct{}{},
// 	386:  struct{}{},
// 	391:  struct{}{},
// 	392:  struct{}{},
// 	397:  struct{}{},
// 	404:  struct{}{},
// 	409:  struct{}{},
// 	440:  struct{}{},
// 	446:  struct{}{},
// 	464:  struct{}{},
// 	469:  struct{}{},
// 	478:  struct{}{},
// 	487:  struct{}{},
// 	490:  struct{}{},
// 	496:  struct{}{},
// 	536:  struct{}{},
// 	556:  struct{}{},
// 	563:  struct{}{},
// 	565:  struct{}{},
// 	566:  struct{}{},
// 	608:  struct{}{},
// 	617:  struct{}{},
// 	622:  struct{}{},
// 	623:  struct{}{},
// 	632:  struct{}{},
// 	635:  struct{}{},
// 	637:  struct{}{},
// 	638:  struct{}{},
// 	644:  struct{}{},
// 	649:  struct{}{},
// 	653:  struct{}{},
// 	655:  struct{}{},
// 	656:  struct{}{},
// 	665:  struct{}{},
// 	671:  struct{}{},
// 	673:  struct{}{},
// 	680:  struct{}{},
// 	683:  struct{}{},
// 	694:  struct{}{},
// 	700:  struct{}{},
// 	709:  struct{}{},
// 	716:  struct{}{},
// 	736:  struct{}{},
// 	739:  struct{}{},
// 	748:  struct{}{},
// 	761:  struct{}{},
// 	763:  struct{}{},
// 	784:  struct{}{},
// 	790:  struct{}{},
// 	793:  struct{}{},
// 	802:  struct{}{},
// 	806:  struct{}{},
// 	818:  struct{}{},
// 	820:  struct{}{},
// 	833:  struct{}{},
// 	836:  struct{}{},
// 	847:  struct{}{},
// 	860:  struct{}{},
// 	863:  struct{}{},
// 	874:  struct{}{},
// 	881:  struct{}{},
// 	888:  struct{}{},
// 	899:  struct{}{},
// 	901:  struct{}{},
// 	904:  struct{}{},
// 	907:  struct{}{},
// 	910:  struct{}{},
// 	912:  struct{}{},
// 	913:  struct{}{},
// 	921:  struct{}{},
// 	923:  struct{}{},
// 	931:  struct{}{},
// 	932:  struct{}{},
// 	937:  struct{}{},
// 	940:  struct{}{},
// 	946:  struct{}{},
// 	964:  struct{}{},
// 	970:  struct{}{},
// 	973:  struct{}{},
// 	989:  struct{}{},
// 	998:  struct{}{},
// 	1000: struct{}{},
// 	1003: struct{}{},
// 	1009: struct{}{},
// 	1029: struct{}{},
// 	1030: struct{}{},
// 	1033: struct{}{},
// 	1039: struct{}{},
// 	1067: struct{}{},
// 	1076: struct{}{},
// 	1088: struct{}{},
// 	1090: struct{}{},
// 	1092: struct{}{},
// 	1093: struct{}{},
// 	1112: struct{}{},
// 	1114: struct{}{},
// 	1115: struct{}{},
// 	1121: struct{}{},
// 	1122: struct{}{},
// 	1125: struct{}{},
// 	1128: struct{}{},
// 	1141: struct{}{},
// 	1148: struct{}{},
// 	1151: struct{}{},
// 	1152: struct{}{},
// 	1158: struct{}{},
// 	1177: struct{}{},
// 	1182: struct{}{},
// 	1184: struct{}{},
// 	1185: struct{}{},
// 	1188: struct{}{},
// 	1209: struct{}{},
// 	1211: struct{}{},
// 	1212: struct{}{},
// 	1215: struct{}{},
// 	1218: struct{}{},
// 	1221: struct{}{},
// 	1222: struct{}{},
// 	1233: struct{}{},
// 	1247: struct{}{},
// 	1251: struct{}{},
// 	1257: struct{}{},
// 	1258: struct{}{},
// 	1274: struct{}{},
// 	1275: struct{}{},
// 	1277: struct{}{},
// 	1281: struct{}{},
// 	1285: struct{}{},
// 	1288: struct{}{},
// 	1290: struct{}{},
// 	1299: struct{}{},
// 	1300: struct{}{},
// 	1303: struct{}{},
// 	1309: struct{}{},
// 	1323: struct{}{},
// 	1330: struct{}{},
// 	1332: struct{}{},
// 	1333: struct{}{},
// 	1335: struct{}{},
// 	1337: struct{}{},
// 	1339: struct{}{},
// 	1353: struct{}{},
// 	1366: struct{}{},
// 	1373: struct{}{},
// 	1390: struct{}{},
// 	1393: struct{}{},
// 	1411: struct{}{},
// 	1418: struct{}{},
// 	1427: struct{}{},
// 	1444: struct{}{},
// 	1447: struct{}{},
// 	1448: struct{}{},
// 	1457: struct{}{},
// 	1472: struct{}{},
// 	1474: struct{}{},
// 	1475: struct{}{},
// 	1478: struct{}{},
// 	1481: struct{}{},
// 	1484: struct{}{},
// 	1487: struct{}{},
// 	1511: struct{}{},
// 	1512: struct{}{},
// 	1518: struct{}{},
// 	1521: struct{}{},
// }

// func isHappy(n int) bool {
// 	var next, mod int
// 	for n != 0 {
// 		n, mod = n/10, n%10
// 		next += mod * mod
// 	}
// 	_, ok := happys[next]
// 	return ok
// }
