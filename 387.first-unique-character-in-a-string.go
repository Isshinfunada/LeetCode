/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
package main

type values struct {
	index int
	flag  bool
}

func firstUniqChar(s string) int {
	hm := make(map[rune]*values)

	for i, r := range s {
		if v, exists := hm[r]; exists {
			v.flag = true
		} else {
			hm[r] = &values{index: i, flag: false}
		}
	}

	result := -1
	for _, v := range s {
		if !hm[v].flag {
			result = hm[v].index
			break
		}
	}
	return result
}

// @lc code=end
