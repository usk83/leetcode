// 468. Validate IP Address
// https://leetcode.com/problems/validate-ip-address/
package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 15
// https://leetcode.com/explore/featured/card/june-leetcoding-challenge/541/week-3-june-15th-june-21st/3362/
func main() {
	pp.Println("=========================================")
	pp.Println(validIPAddress("172.16.254.1"))
	pp.Println("IPv4")
	pp.Println("=========================================")
	pp.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	pp.Println("IPv6")
	pp.Println("=========================================")
	pp.Println(validIPAddress("256.256.256.256"))
	pp.Println("Neither")
	pp.Println("=========================================")
	pp.Println(validIPAddress("172.16.254.01"))
	pp.Println("Neither")
	pp.Println("=========================================")
	pp.Println(validIPAddress("2001::85a3:0:0:8A2E:0370:7334"))
	pp.Println("Neither")
	pp.Println("=========================================")
	pp.Println(validIPAddress("1081:db8:85a3:01:-0:8A2E:0370:7334"))
	pp.Println("Neither")
	pp.Println("=========================================")
	pp.Println(validIPAddress("192.0.0.1"))
	pp.Println("IPv4")
	pp.Println("=========================================")
}

func validIPAddress(ip string) string {
	if parts := strings.Split(ip, "."); len(parts) == 4 {
		// IPv4
		// xxx.xxx.xxx.xxx
		// xxx: 0~255
		// no loading 0
		for _, part := range parts {
			if num, err := strconv.ParseInt(part, 10, 64); err != nil ||
				num < 0 ||
				num > 255 ||
				(part != "0" && int(math.Log10(float64(num)))+1 != len(part)) {
				return "Neither"
			}
		}
		return "IPv4"
	} else if parts = strings.Split(ip, ":"); len(parts) == 8 {
		// IPv6
		// xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx
		// hex
		// leading 0 OK
		// no empty
		// no extra 0
		for _, part := range parts {
			if len(part) < 1 || 4 < len(part) {
				return "Neither"
			}
			for _, r := range part {
				if !(('0' <= r && r <= '9') || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F')) {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
	return "Neither"
}
