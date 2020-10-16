// 5441. Making File Names Unique
// https://leetcode.com/contest/weekly-contest-194/problems/making-file-names-unique/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(getFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
	fmt.Println([]string{"pes", "fifa", "gta", "pes(2019)"})
	pp.Println("=========================================")
	fmt.Println(getFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
	fmt.Println([]string{"gta", "gta(1)", "gta(2)", "avalon"})
	pp.Println("=========================================")
	fmt.Println(getFolderNames([]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}))
	fmt.Println([]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece(4)"})
	pp.Println("=========================================")
	fmt.Println(getFolderNames([]string{"wano", "wano", "wano", "wano"}))
	fmt.Println([]string{"wano", "wano(1)", "wano(2)", "wano(3)"})
	pp.Println("=========================================")
	fmt.Println(getFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)"}))
	fmt.Println([]string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"})
	pp.Println("=========================================")
	fmt.Println(getFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)"}))
	fmt.Println([]string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"})
	pp.Println("=========================================")
	fmt.Println(getFolderNames([]string{"kingston(0)", "kingston", "kingston"}))
	fmt.Println([]string{"kingston(0)", "kingston", "kingston(1)"})
	pp.Println("=========================================")

	input := []string{"l", "p(4)(4)", "v", "c", "i", "k(2)", "y", "a(4)(2)", "t", "a", "e", "e(1)(4)", "p", "p", "s", "z", "s", "q(2)(1)", "r(1)", "m", "b(3)", "u(2)", "j", "p", "t", "s", "g", "o", "o", "b(3)(3)", "k", "w(2)(3)", "q", "a", "z", "t(4)(4)", "s(2)", "c", "w", "u", "h", "g", "b", "r", "c(1)", "v", "n(1)", "r(2)(1)", "u", "t(2)(3)", "p", "m(3)(1)", "h", "o", "y", "n", "s", "z", "k(2)", "x(4)", "l(1)(4)", "g(2)", "u", "t(2)", "m", "c", "t", "g", "c", "a(2)", "r", "d", "y", "b", "p", "m", "w(2)(4)", "v", "t(4)(2)", "x", "e(4)(1)", "h", "f", "z", "e(1)(4)", "t", "b(1)", "x(4)", "m(3)", "j(2)(4)", "s(3)(2)", "z", "l", "p(2)(2)", "g(3)(2)", "q(1)(4)", "h(1)(1)", "h", "o(3)", "h", "f", "n(4)", "s(4)", "g", "s", "r", "n(1)(1)", "x(3)", "r", "f(3)(1)", "e", "j", "s", "g", "d", "l", "g", "o(4)(3)", "x(4)", "u(2)(4)", "x", "t", "f", "j(1)", "v(2)", "w", "v", "t(2)(2)", "l", "o(1)(1)", "a", "y", "q(4)", "m(1)(2)", "i", "u", "l", "c(1)", "g", "l(2)", "p(1)(1)", "k", "d", "l", "o", "m", "i(1)", "j", "i", "f", "y", "b", "k", "z", "n", "t", "c(2)", "y", "q", "b", "t", "m", "g(1)", "r", "m", "l", "s", "n", "j(4)(4)", "m(3)(3)", "n", "n(2)(3)", "s", "t", "l", "e", "p", "q(2)", "v", "v", "b", "w", "m", "p", "g", "h(2)", "n(1)", "q", "x(4)", "q(2)(4)", "s(2)", "w", "k", "f", "v", "n", "q", "w", "y(3)", "a(4)(1)", "r(1)", "r", "f(4)(2)", "l", "f(2)(3)", "p", "o", "h", "t", "i", "r", "k", "p", "l", "o(3)(4)", "f(1)", "j(4)", "f(3)(4)", "o(2)(1)", "j", "k(2)", "k", "x(4)", "s(3)(4)", "p(1)(3)", "y", "r(2)", "i(3)(4)", "j", "p(4)(3)", "e", "j(4)", "g(4)(2)", "x", "l", "x", "g", "w", "o", "t(2)(4)", "s", "f", "o", "h", "h(2)(4)", "z(4)(2)", "w", "l", "n", "q", "w", "l", "a", "o", "v", "h", "b", "v", "b", "k", "x(2)", "r(2)(1)", "g(2)", "c(4)(3)", "w(1)", "g(4)", "z", "i(3)(2)", "r", "e", "p", "z(4)(2)", "s", "c(1)", "h", "l", "j(3)(4)", "x(1)(2)", "a", "b", "t", "c"}

	output := getFolderNames(input)

	expected := []string{"l", "p(4)(4)", "v", "c", "i", "k(2)", "y", "a(4)(2)", "t", "a", "e", "e(1)(4)", "p", "p(1)", "s", "z", "s(1)", "q(2)(1)", "r(1)", "m", "b(3)", "u(2)", "j", "p(2)", "t(1)", "s(2)", "g", "o", "o(1)", "b(3)(3)", "k", "w(2)(3)", "q", "a(1)", "z(1)", "t(4)(4)", "s(2)(1)", "c(1)", "w", "u", "h", "g(1)", "b", "r", "c(1)(1)", "v(1)", "n(1)", "r(2)(1)", "u(1)", "t(2)(3)", "p(3)", "m(3)(1)", "h(1)", "o(2)", "y(1)", "n", "s(3)", "z(2)", "k(2)(1)", "x(4)", "l(1)(4)", "g(2)", "u(3)", "t(2)", "m(1)", "c(2)", "t(3)", "g(3)", "c(3)", "a(2)", "r(2)", "d", "y(2)", "b(1)", "p(4)", "m(2)", "w(2)(4)", "v(2)", "t(4)(2)", "x", "e(4)(1)", "h(2)", "f", "z(3)", "e(1)(4)(1)", "t(4)", "b(1)(1)", "x(4)(1)", "m(3)", "j(2)(4)", "s(3)(2)", "z(4)", "l(1)", "p(2)(2)", "g(3)(2)", "q(1)(4)", "h(1)(1)", "h(3)", "o(3)", "h(4)", "f(1)", "n(4)", "s(4)", "g(4)", "s(5)", "r(3)", "n(1)(1)", "x(3)", "r(4)", "f(3)(1)", "e(1)", "j(1)", "s(6)", "g(5)", "d(1)", "l(2)", "g(6)", "o(4)(3)", "x(4)(2)", "u(2)(4)", "x(1)", "t(5)", "f(2)", "j(1)(1)", "v(2)(1)", "w(1)", "v(3)", "t(2)(2)", "l(3)", "o(1)(1)", "a(3)", "y(3)", "q(4)", "m(1)(2)", "i(1)", "u(4)", "l(4)", "c(1)(2)", "g(7)", "l(2)(1)", "p(1)(1)", "k(1)", "d(2)", "l(5)", "o(4)", "m(4)", "i(1)(1)", "j(2)", "i(2)", "f(3)", "y(4)", "b(2)", "k(3)", "z(5)", "n(2)", "t(6)", "c(2)(1)", "y(5)", "q(1)", "b(4)", "t(7)", "m(5)", "g(1)(1)", "r(5)", "m(6)", "l(6)", "s(7)", "n(3)", "j(4)(4)", "m(3)(3)", "n(5)", "n(2)(3)", "s(8)", "t(8)", "l(7)", "e(2)", "p(5)", "q(2)", "v(4)", "v(5)", "b(5)", "w(2)", "m(7)", "p(6)", "g(8)", "h(2)(1)", "n(1)(2)", "q(3)", "x(4)(3)", "q(2)(4)", "s(2)(2)", "w(3)", "k(4)", "f(4)", "v(6)", "n(6)", "q(5)", "w(4)", "y(3)(1)", "a(4)(1)", "r(1)(1)", "r(6)", "f(4)(2)", "l(8)", "f(2)(3)", "p(7)", "o(5)", "h(5)", "t(9)", "i(3)", "r(7)", "k(5)", "p(8)", "l(9)", "o(3)(4)", "f(1)(1)", "j(4)", "f(3)(4)", "o(2)(1)", "j(3)", "k(2)(2)", "k(6)", "x(4)(4)", "s(3)(4)", "p(1)(3)", "y(6)", "r(2)(2)", "i(3)(4)", "j(5)", "p(4)(3)", "e(3)", "j(4)(1)", "g(4)(2)", "x(2)", "l(10)", "x(5)", "g(9)", "w(5)", "o(6)", "t(2)(4)", "s(9)", "f(5)", "o(7)", "h(6)", "h(2)(4)", "z(4)(2)", "w(6)", "l(11)", "n(7)", "q(6)", "w(7)", "l(12)", "a(4)", "o(8)", "v(7)", "h(7)", "b(6)", "v(8)", "b(7)", "k(7)", "x(2)(1)", "r(2)(1)(1)", "g(2)(1)", "c(4)(3)", "w(1)(1)", "g(4)(1)", "z(6)", "i(3)(2)", "r(8)", "e(4)", "p(9)", "z(4)(2)(1)", "s(10)", "c(1)(3)", "h(8)", "l(13)", "j(3)(4)", "x(1)(2)", "a(5)", "b(8)", "t(10)", "c(4)"}

	for i := range input {
		if output[i] != expected[i] {
			fmt.Printf("[%d]: %s != %s\n", i, output[i], expected[i])
		}
	}
	pp.Println("=========================================")
}

func getFolderNames(names []string) []string {
	result := make([]string, 0, len(names))
	used := map[string]map[string]bool{}
	lastCount := map[string]int{}
	for iii, name := range names {
		/*
			使用済みでない
				そのまま
				使用済みにする
					数字削ったものも
			使用済み
				最後のカウントを探す
				カウントを追加して
				使用済みにする
		*/

		if used[name] == nil {
			result = append(result, name)
			used[name] = map[string]bool{
				"": true,
			}

			if str, num, ok := separate(name); ok {
				if d := used[str]; d == nil {
					used[str] = map[string]bool{
						strconv.Itoa(num): true,
					}
					if iii == 47 {
						pp.Println(name)
						fmt.Println(str)
						fmt.Println(used[str])
					}
				} else {
					d[strconv.Itoa(num)] = true
					used[str] = d
				}
			}
		} else {
			if iii == 226 {
				pp.Println(name)
				// fmt.Println(str)
				fmt.Println(used[name])
			}

			if !used[name][""] {
				result = append(result, name)
				d := used[name]
				d[""] = true
				used[name] = d

				if str, num, ok := separate(name); ok {
					if d := used[str]; d == nil {
						used[str] = map[string]bool{
							strconv.Itoa(num): true,
						}
					} else {
						d[strconv.Itoa(num)] = true
						used[str] = d
					}
				}

				continue
			}

			num := 1
			if n, ok := lastCount[name]; ok {
				num = n
			}
			for used[name][strconv.Itoa(num)] {
				num++
			}
			rename := name + "(" + strconv.Itoa(num) + ")"
			result = append(result, rename)
			d := used[name]
			d[strconv.Itoa(num)] = true
			used[name] = d

			// if name == "t" {
			// 	pp.Printf("%s: %s\n", iii, used[name])
			// }

			used[rename] = map[string]bool{
				"": true,
			}
		}
	}
	return result
}

func separate(str string) (string, int, bool) {
	if strings.HasSuffix(str, ")") {
		for i := len(str) - 2; i >= 0; i-- {
			if str[i] == '(' {
				n, err := strconv.Atoi(str[i+1 : len(str)-1])
				if err != nil {
					return str, 0, false
				}
				return str[:i], n, true
			}
		}
	}
	return str, 0, false
}
