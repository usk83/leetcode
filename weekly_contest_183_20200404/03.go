// 5195. Longest Happy String
// https://leetcode.com/problems/xxx/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(longestDiverseString(1, 1, 7))
	pp.Println("ccaccbcc")
	pp.Println("=========================================")
	pp.Println(longestDiverseString(2, 2, 1))
	pp.Println("aabbc")
	pp.Println("=========================================")
	pp.Println(longestDiverseString(7, 1, 0))
	pp.Println("aabaa")
	pp.Println("=========================================")
	pp.Println(longestDiverseString(1, 4, 5))
	pp.Println("ccbbccbbac")
	pp.Println("=========================================")
	pp.Println(longestDiverseString(0, 8, 11))
	pp.Println("ccbccbbccbbccbbccbc")
	pp.Println("=========================================")
	pp.Println(longestDiverseString(33, 60, 11))
	pp.Println("bbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbabbcbbaabbccaabbccaabbccaabbccaabbcc")
	pp.Println("=========================================")
}

func longestDiverseString(a int, b int, c int) string {
	var sb strings.Builder
	sb.Grow(a + b + c)
	for {
		if a <= 0 && b <= 0 {
			for i := min(c, 2); i > 0; i-- {
				sb.WriteByte('c')
			}
			break
		} else if a <= 0 && c <= 0 {
			for i := min(b, 2); i > 0; i-- {
				sb.WriteByte('b')
			}
			break
		} else if b <= 0 && c <= 0 {
			for i := min(a, 2); i > 0; i-- {
				sb.WriteByte('a')
			}
			break
		}

		if a >= b && b >= c {
			for i := min(a, 2); i > 0; i-- {
				sb.WriteByte('a')
				a--
			}
			if a < b {
				for i := min(2, b); i > 0; i-- {
					sb.WriteByte('b')
					b--
				}
			} else {
				for i := min(1, b); i > 0; i-- {
					sb.WriteByte('b')
					b--
				}
			}
			if a < c {
				for i := min(2, c); i > 0; i-- {
					sb.WriteByte('c')
					c--
				}
				// } else {
				// 	for i := min(1, c); i > 0; i-- {
				// 		sb.WriteByte('c')
				// 		c--
				// 	}
			}

			// for i := max(c-a, 0); i > 0; i-- {
			// 	sb.WriteByte('c')
			// 	c--
			// }
		} else if a >= c && c >= b {
			for i := min(a, 2); i > 0; i-- {
				sb.WriteByte('a')
				a--
			}
			if a < c {
				for i := min(2, c); i > 0; i-- {
					sb.WriteByte('c')
					c--
				}
			} else {
				for i := min(1, c); i > 0; i-- {
					sb.WriteByte('c')
					c--
				}
			}
			if a < b {
				for i := min(2, b); i > 0; i-- {
					sb.WriteByte('b')
					b--
				}
				// } else {
				// 	for i := min(1, b); i > 0; i-- {
				// 		sb.WriteByte('b')
				// 		b--
				// 	}
			}

			// for i := max(b-a, 0); i > 0; i-- {
			// 	sb.WriteByte('b')
			// 	b--
			// }
		} else if b >= a && a >= c {
			for i := min(b, 2); i > 0; i-- {
				sb.WriteByte('b')
				b--
			}
			if b < a {
				for i := min(2, a); i > 0; i-- {
					sb.WriteByte('a')
					a--
				}
			} else {
				for i := min(1, a); i > 0; i-- {
					sb.WriteByte('a')
					a--
				}
			}
			if b < c {
				for i := min(2, c); i > 0; i-- {
					sb.WriteByte('c')
					c--
				}
				// } else {
				// 	for i := min(1, c); i > 0; i-- {
				// 		sb.WriteByte('c')
				// 		c--
				// 	}
			}

			// for i := max(c-b, 0); i > 0; i-- {
			// 	sb.WriteByte('c')
			// 	c--
			// }
		} else if b >= c && c >= a {
			for i := min(b, 2); i > 0; i-- {
				sb.WriteByte('b')
				b--
			}
			if b < c {
				for i := min(2, c); i > 0; i-- {
					sb.WriteByte('c')
					c--
				}
			} else {
				for i := min(1, c); i > 0; i-- {
					sb.WriteByte('c')
					c--
				}
			}
			if b < a {
				for i := min(2, a); i > 0; i-- {
					sb.WriteByte('a')
					a--
				}
				// } else {
				// 	for i := min(1, a); i > 0; i-- {
				// 		sb.WriteByte('a')
				// 		a--
				// 	}
			}

			// for i := max(a-b, 0); i > 0; i-- {
			// 	sb.WriteByte('a')
			// 	a--
			// }
		} else if c >= a && a >= b {
			for i := min(c, 2); i > 0; i-- {
				sb.WriteByte('c')
				c--
			}
			if c < a {
				for i := min(2, a); i > 0; i-- {
					sb.WriteByte('a')
					a--
				}
			} else {
				for i := min(1, a); i > 0; i-- {
					sb.WriteByte('a')
					a--
				}
			}
			if c < b {
				for i := min(2, b); i > 0; i-- {
					sb.WriteByte('b')
					b--
				}
				// } else {
				// 	for i := min(1, b); i > 0; i-- {
				// 		sb.WriteByte('b')
				// 		b--
				// 	}
			}

			// for i := max(b-c, 0); i > 0; i-- {
			// 	sb.WriteByte('b')
			// 	b--
			// }
		} else if c >= b && b >= a {
			for i := min(c, 2); i > 0; i-- {
				sb.WriteByte('c')
				c--
			}
			if c < b {
				for i := min(2, b); i > 0; i-- {
					sb.WriteByte('b')
					b--
				}
			} else {
				for i := min(1, b); i > 0; i-- {
					sb.WriteByte('b')
					b--
				}
			}
			if c < a {
				for i := min(2, a); i > 0; i-- {
					sb.WriteByte('a')
					a--
				}
				// } else {
				// 	for i := min(1, a); i > 0; i-- {
				// 		sb.WriteByte('a')
				// 		a--
				// 	}
			}

			// for i := max(a-c, 0); i > 0; i-- {
			// 	sb.WriteByte('a')
			// 	a--
			// }
		}

	}
	return sb.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func longestDiverseString(a int, b int, c int) string {
// 	var sb strings.Builder
// 	sb.Grow(a + b + c)
// 	for {
// 		if a <= 0 && b <= 0 {
// 			for i := min(c, 2); i > 0; i-- {
// 				sb.WriteByte('c')
// 			}
// 			break
// 		} else if a <= 0 && c <= 0 {
// 			for i := min(b, 2); i > 0; i-- {
// 				sb.WriteByte('b')
// 			}
// 			break
// 		} else if b <= 0 && c <= 0 {
// 			for i := min(a, 2); i > 0; i-- {
// 				sb.WriteByte('a')
// 			}
// 			break
// 		}

// 		if a >= b && b >= c {
// 			for i := min(a, 2); i > 0; i-- {
// 				sb.WriteByte('a')
// 			}
// 			for i := min(b, 2); i > 0; i-- {
// 				sb.WriteByte('b')
// 			}
// 			for i := min(c, 2); i > 0; i-- {
// 				sb.WriteByte('c')
// 			}
// 		} else if a >= c && c >= b {
// 			for i := min(a, 2); i > 0; i-- {
// 				sb.WriteByte('a')
// 			}
// 			for i := min(c, 2); i > 0; i-- {
// 				sb.WriteByte('c')
// 			}
// 			for i := min(b, 2); i > 0; i-- {
// 				sb.WriteByte('b')
// 			}
// 		} else if b >= a && a >= c {
// 			for i := min(b, 2); i > 0; i-- {
// 				sb.WriteByte('b')
// 			}
// 			for i := min(a, 2); i > 0; i-- {
// 				sb.WriteByte('a')
// 			}
// 			for i := min(c, 2); i > 0; i-- {
// 				sb.WriteByte('c')
// 			}
// 		} else if b >= c && c >= a {
// 			for i := min(b, 2); i > 0; i-- {
// 				sb.WriteByte('b')
// 			}
// 			for i := min(c, 2); i > 0; i-- {
// 				sb.WriteByte('c')
// 			}
// 			for i := min(a, 2); i > 0; i-- {
// 				sb.WriteByte('a')
// 			}
// 		} else if c >= a && a >= b {
// 			for i := min(c, 2); i > 0; i-- {
// 				sb.WriteByte('c')
// 			}
// 			for i := min(a, 2); i > 0; i-- {
// 				sb.WriteByte('a')
// 			}
// 			for i := min(b, 2); i > 0; i-- {
// 				sb.WriteByte('b')
// 			}
// 		} else if c >= b && b >= a {
// 			for i := min(c, 2); i > 0; i-- {
// 				sb.WriteByte('c')
// 			}
// 			for i := min(b, 2); i > 0; i-- {
// 				sb.WriteByte('b')
// 			}
// 			for i := min(a, 2); i > 0; i-- {
// 				sb.WriteByte('a')
// 			}
// 		}
// 		a -= 2
// 		b -= 2
// 		c -= 2
// 	}

// 	return sb.String()
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
