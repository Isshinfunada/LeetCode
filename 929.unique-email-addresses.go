/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
package main

import "fmt"

func numUniqueEmails(emails []string) int {
	list := normalize(emails)
	fmt.Printf("result: %#v\n", list)
	return len(list)
}

func normalize(emails []string) []string {
	var localname []string
	var domainname []string
	var result []string
	var after bool
	// まず@前後でわける
	for _, v := range emails {
		for _, r := range v {
			if r == '@' {
				after = true
				continue
			}
			if !after {
				if r == '.' {
					continue
				}
				if r == '+' {
					break
				}
				localname = append(localname, string(r))
				continue
			}
			domainname = append(domainname, string(r))
		}

		after = false
	}
	fmt.Printf("localname: %#v\n", localname)
	fmt.Printf("domainname: %#v\n", domainname)

	result = append(result, localname...)
	result = append(result, domainname...)
	return result
}

// @lc code=end
