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
	memo := make(map[string]bool)
	for len(q) != 0 {
		remaining := q[0]
		q = q[1:] // Dequeue
		if _, ok := memo[remaining]; ok {
			continue
		}
		for _, word := range wordDict {
			if word == remaining {
				return true
			}
			if strings.HasPrefix(remaining, word) {
				q = append(q, remaining[len(word):]) // Enqueue
				memo[remaining] = true
			}
		}
	}
	return false
}

// @lc code=end
