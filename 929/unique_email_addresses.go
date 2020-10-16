// 929. Unique Email Addresses
// https://leetcode.com/problems/unique-email-addresses/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numUniqueEmails([]string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}))
	pp.Println(2)
	pp.Println("=========================================")
}

// 1 <= emails[i].length <= 100
// 1 <= emails.length <= 100
// Each emails[i] contains exactly one '@' character.
// All local and domain names are non-empty.
// Local names do not start with a '+' character.

/*
 * The optimized solution
 * Iterate over it only once for each email
 * Time Complexity: O(N*M) N: the number of emails, M: the length of each email
 *              OR  O(N+C) N: the number of emails, C: the total content of emails
 * Space Complexity: O(N+C)
 * > Runtime: 0 ms, faster than 100.00%
 * > Memory Usage: 5.9 MB, less than 80.00%
 */
func numUniqueEmails(emails []string) int {
	set := map[string]struct{}{}
	for _, email := range emails {
		length := len(email)
		chars := make([]byte, 0, length)
	LOOP_BY_CHAR:
		for i := 0; i < length; i++ {
			switch email[i] {
			case '.':
				// do nothing
			case '+':
				i += strings.Index(email[i:], "@")
				fallthrough
			case '@':
				chars = append(chars, []byte(email[i:])...)
				break LOOP_BY_CHAR
			default:
				chars = append(chars, email[i])
			}
		}
		set[string(chars)] = struct{}{}
	}
	return len(set)
}

/*
 * Easy-to-understand solution
 */
// func numUniqueEmails(emails []string) int {
// 	set := make(map[string]struct{})
// 	for _, email := range emails {
// 		components := strings.Split(email, "@")
// 		components[0] = strings.Split(components[0], "+")[0]
// 		components[0] = strings.Replace(components[0], ".", "", -1)
// 		set[strings.Join(components, "@")] = struct{}{}
// 	}
// 	return len(set)
// }

/*
 * The first solution
 * > Runtime: 8 ms, faster than 77.69%
 * > Memory Usage: 5.9 MB, less than 80.00%
 */
// func numUniqueEmails(emails []string) int {
// 	uniq := map[string]bool{}
// 	for _, email := range emails {
// 		at := strings.Index(email, "@")
//
// 		localname := email[:at]
// 		domain := email[at:]
//
// 		if plus := strings.Index(localname, "+"); plus != -1 {
// 			localname = localname[:plus]
// 		}
//
// 		dot := strings.Index(localname, ".")
// 		for dot != -1 {
// 			localname = localname[:dot] + localname[dot+1:]
// 			dot = strings.Index(localname, ".")
// 		}
//
// 		uniq[localname+domain] = true
// 	}
// 	return len(uniq)
// }
