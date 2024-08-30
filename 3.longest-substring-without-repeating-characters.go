/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
package main

import (
	"strconv"
)

type Lists struct {
	frequency int
	length    int
}

func lengthOfLongestSubstring(s string) int {
	var str []rune
	var moji string
	master := make(map[string]*Lists)

	for i, r := range s {
		moji = strconv.QuoteRune(r)
		master[moji] = &Lists{}
		if i != 0 {
			if str[i-1] != r {
				str = append(str, r)
				continue
			}
			moji = string(str)
		}
		if master[moji] == nil {
			master[moji].frequency += 1
			master[moji].length = len(str)
		}
		master[moji].frequency += 1
	}

	// for master[moji].frequency
	// 	for master[moji].length

	return master["abc"].length
}

// @lc code=end
