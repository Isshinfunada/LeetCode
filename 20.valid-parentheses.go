/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
package main

func isValid(s string) bool {
	cnt1 := 0
	cnt3 := 0
	cnt5 := 0
	hm := make(map[string]int)
	hm["("] = 1
	hm[")"] = 1
	hm["{"] = 3
	hm["}"] = 3
	hm["["] = 5
	hm["]"] = 5

	for _, r := range s {
		if hm[string(r)] == 1 {
			cnt1 += hm[string(r)]
		}
		if hm[string(r)] == 3 {
			cnt3 += hm[string(r)]
		}
		if hm[string(r)] == 5 {
			cnt5 += hm[string(r)]
		}
	}
	return cnt1%2 == 0 && cnt3%2 == 0 && cnt5%2 == 0
}

// @lc code=end
