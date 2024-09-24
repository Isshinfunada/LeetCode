/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	hm := make(map[rune]bool)
	for _, str := range wordDict {
		fmt.Printf("str: %#v\n", str)
		for _, r := range str {
			fmt.Printf("r: %#v\n", string(r))
			hm[r] = true
		}
	}

	var check bool
	for _, c := range s {
		fmt.Printf("c: %#v\n", string(c))
		if _, ok := hm[c]; ok {
			continue
		} else {
			check = true
		}
	}
	fmt.Printf("hm: %#v\n", hm)
	return !check
}

// @lc code=end
