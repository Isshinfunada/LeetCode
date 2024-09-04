/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
package main

import "strings"

func numUniqueEmails(emails []string) int {
	d := make(map[string]bool)

	for _, email := range emails {
		parts := strings.Split(email, "@")
		local := strings.Split(parts[0], "+")
		check := strings.Replace(local[0], ".", "", -1) + "@" + parts[1]

		d[check] = true
	}

	return len(d)
}

// @lc code=end
