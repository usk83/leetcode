// 804. Unique Morse Code Words
// https://leetcode.com/problems/unique-morse-code-words/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(uniqueMorseRepresentations([]string{"gi", "ze", "ms"}))
	pp.Println(1)
	pp.Println("=========================================")
}

// The length of words will be at most 100.
// Each words[i] will have length in range [1, 12].
// words[i] will only consist of lowercase letters.

/*
 * v2. Bit manipulation
 */
type morseCode struct {
	length uint
	value  int
}

// For Go 1.13 and later
var morseCodes = [26]morseCode{
	morseCode{2, 0b10},   // 'a': ".-"
	morseCode{4, 0b0111}, // 'b': "-..."
	morseCode{4, 0b0101}, // 'c': "-.-."
	morseCode{3, 0b011},  // 'd': "-.."
	morseCode{1, 0b1},    // 'e': "."
	morseCode{4, 0b1101}, // 'f': "..-."
	morseCode{3, 0b001},  // 'g': "--."
	morseCode{4, 0b1111}, // 'h': "...."
	morseCode{2, 0b11},   // 'i': ".."
	morseCode{4, 0b1000}, // 'j': ".---"
	morseCode{3, 0b010},  // 'k': "-.-"
	morseCode{4, 0b1011}, // 'l': ".-.."
	morseCode{2, 0b00},   // 'm': "--"
	morseCode{2, 0b01},   // 'n': "-."
	morseCode{3, 0b000},  // 'o': "---"
	morseCode{4, 0b1001}, // 'p': ".--."
	morseCode{4, 0b0010}, // 'q': "--.-"
	morseCode{3, 0b101},  // 'r': ".-."
	morseCode{3, 0b111},  // 's': "..."
	morseCode{1, 0b0},    // 't': "-"
	morseCode{3, 0b110},  // 'u': "..-"
	morseCode{4, 0b1110}, // 'v': "...-"
	morseCode{3, 0b100},  // 'w': ".--"
	morseCode{4, 0b0110}, // 'x': "-..-"
	morseCode{4, 0b0100}, // 'y': "-.--"
	morseCode{4, 0b0011}, // 'z': "--.."
}

func uniqueMorseRepresentations(words []string) int {
	set := map[int]struct{}{}
	for _, word := range words {
		var code int
		for _, c := range word {
			morseCode := morseCodes[c-'a']
			code = code<<morseCode.length | morseCode.value
		}
		set[code] = struct{}{}
	}
	return len(set)
}

/*
 * v1. Using `strings.Builder`
 */
// var morseCodes = [26]string{
// 	".-",   // 'a'
// 	"-...", // 'b'
// 	"-.-.", // 'c'
// 	"-..",  // 'd'
// 	".",    // 'e'
// 	"..-.", // 'f'
// 	"--.",  // 'g'
// 	"....", // 'h'
// 	"..",   // 'i'
// 	".---", // 'j'
// 	"-.-",  // 'k'
// 	".-..", // 'l'
// 	"--",   // 'm'
// 	"-.",   // 'n'
// 	"---",  // 'o'
// 	".--.", // 'p'
// 	"--.-", // 'q'
// 	".-.",  // 'r'
// 	"...",  // 's'
// 	"-",    // 't'
// 	"..-",  // 'u'
// 	"...-", // 'v'
// 	".--",  // 'w'
// 	"-..-", // 'x'
// 	"-.--", // 'y'
// 	"--..", // 'z'
// }

// func uniqueMorseRepresentations(words []string) int {
// 	set := map[string]struct{}{}
// 	for _, word := range words {
// 		var builder strings.Builder
// 		builder.Grow(len(word) << 2)
// 		for _, c := range word {
// 			_, _ = builder.WriteString(morseCodes[c-'a'])
// 		}
// 		set[builder.String()] = struct{}{}
// 		pp.Println(word, ":", builder.String())
// 	}
// 	return len(set)
// }
