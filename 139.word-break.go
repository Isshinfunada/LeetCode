/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
package main

import "strings"

func wordBreak(s string, wordDict []string) bool {
	q := []string{s}
	for len(q) != 0 {
		remain := q[0]
		for _, w := range wordDict {
			if remain == w {
				return true
			}
			if strings.HasPrefix(remain, w) {
				q = append(q, remain[len(w):])
			}
		}
	}
	return false
}

// @lc code=end
