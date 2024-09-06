/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	j := 0
	for i, r := range s {
		for ; j < len(t); j++ {
			if r == rune(t[j]) {
				j += 1
				if i == len(s)-1 {
					return true
				}
				break
			}
		}
	}
	return false
}

// @lc code=end
